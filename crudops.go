package sql_db

import (
	"database/sql"
	"fmt"
	"log"
)

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Crudin1(CrudDb *sql.DB) {
	var i Task
	var r Task
	InputTask(&i)
	AddTask(CrudDb, &i)
	GetTaskbyTitle(CrudDb, &i, &r)
	ReturnToInput(&r, &i)

	CompleteTask(CrudDb, &i)
	GetAllTasks(CrudDb)
	DeleteTask(CrudDb, &i)
	GetAllTasks(CrudDb)
}

func InputTask(i *Task) {
	fmt.Print("\nEnter title of task: ")
	fmt.Scanln(&i.Title)
}

func ReturnToInput(r *Task, i *Task) {
	i.Id = r.Id
	i.Title = r.Title
	i.Completed = r.Completed
}

func Count(db *sql.DB) int {
	rows, err := db.Query("SELECT COUNT(*) FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Number of rows are %v\n", count)
	return count
}

func GetTaskbyTitle(db *sql.DB, i *Task, r *Task) {
	resultId, err := db.Query("SELECT * FROM tasks WHERE title IS NOT NULL AND title = (?)", &i.Title)
	if err != nil {
		panic(err.Error())
	}
	for resultId.Next() {
		err = resultId.Scan(&r.Id, &r.Title, &r.Completed)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Printf("\nRetrieved these tasks where title is: %s\n\n", i.Title)
	fmt.Println(r.Id, r.Title, r.Completed)
}

func GetAllTasks(db *sql.DB) []Task {
	results, err := db.Query("SELECT * FROM tasks WHERE id IS NOT NULL")
	if err != nil {
		panic(err.Error())
	}
	// count := Count(db)
	fmt.Printf("\nRetrieved these tasks:\n\n")
	var alltasks []Task
	for results.Next() {
		var task Task
		err = results.Scan(&task.Id, &task.Title, &task.Completed)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(task.Id, task.Title, task.Completed)
		alltasks = append(alltasks, task)
	}
	fmt.Print("\n")
	return alltasks
}

func AddTask(db *sql.DB, i *Task) {
	insert, err := db.Query("INSERT INTO tasks (title,completed) VALUES(?,false)", i.Title)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Printf("\nSuccesfully inserted: %v\n", i.Title)
}

func DeleteTask(db *sql.DB, i *Task) {
	delete, err := db.Query("DELETE FROM tasks WHERE id = (?)", i.Id)
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
	fmt.Printf("\nSuccesfully completed task at ID: %v\n", i.Id)
}

func CompleteTask(db *sql.DB, i *Task) {
	i.Completed = true
	insert, err := db.Query("UPDATE tasks SET completed = (?) WHERE id = (?)", i.Completed, i.Id)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Printf("\nSuccesfully completed task at ID: %v\n", i.Id)
}
