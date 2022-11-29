package sql_db

import (
	"database/sql"
	"fmt"
)

var CrudDb *sql.DB

func crudin1() {

	var i Input
	var r Returned
	InputTask(CrudDb, &i)
	AddTask(CrudDb, &i)
	GetTaskbyTitle(CrudDb, &i, &r)
	ReturnToInput(&r, &i)

	CompleteTask(CrudDb, &i)
	GetAllTasks(CrudDb)
	DeleteTask(CrudDb, &i)
	GetAllTasks(CrudDb)
}

func InputTask(db *sql.DB, i *Input) {
	fmt.Print("\nEnter title of task: ")
	fmt.Scanln(&i.Title)
}

func ReturnToInput(r *Returned, i *Input) {
	i.Id = r.Id
	i.Title = r.Title
	i.Completed = r.Completed
}

func GetTaskbyTitle(db *sql.DB, i *Input, r *Returned) {
	resultId, err := db.Query("SELECT * FROM tasks WHERE title = (?)", &i.Title)
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

func GetAllTasks(db *sql.DB) {
	results, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\nRetrieved these tasks:\n\n")
	for results.Next() {
		var task Task
		err = results.Scan(&task.Id, &task.Title, &task.Completed)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(task.Id, task.Title, task.Completed)
	}
	fmt.Print("\n")
}

func AddTask(db *sql.DB, i *Input) {
	insert, err := db.Query("INSERT INTO tasks (title,completed) VALUES(?,false)", i.Title)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Printf("\nSuccesfully inserted: %v\n", i.Title)
}

func DeleteTask(db *sql.DB, i *Input) {
	delete, err := db.Query("DELETE FROM tasks WHERE id = (?)", i.Id)
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
	fmt.Printf("\nSuccesfully completed task at ID: %v\n", i.Id)
}

func CompleteTask(db *sql.DB, i *Input) {
	i.Completed = true
	insert, err := db.Query("UPDATE tasks SET completed = (?) WHERE id = (?)", i.Completed, i.Id)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Printf("\nSuccesfully completed task at ID: %v\n", i.Id)
}
