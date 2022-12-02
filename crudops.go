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

func CrudTest(CrudDb *sql.DB) {
	var testTask Task
	InputTask(&testTask)
	AddTask(CrudDb, &testTask)
	returnedArray := *GetTaskByTitle(CrudDb, &testTask)
	testTask = returnedArray[0]
	CompleteTask(CrudDb, &testTask)
	returnedArray = *GetAllTasks(CrudDb)
	printArray(returnedArray)
	DeleteTask(CrudDb, &testTask)
	returnedArray = *GetAllTasks(CrudDb)
	printArray(returnedArray)
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

func printArray(array []Task) {
	fmt.Println("printArray: ")
	for _, task := range array {
		fmt.Println(task.Id, task.Title, task.Completed)
	}
}

func queryToArr(results *sql.Rows) *[]Task {
	fmt.Printf("\nLog: Added *sql.Rows to array: \n\n")
	var allTasks []Task
	for results.Next() {
		var task Task
		err := results.Scan(&task.Id, &task.Title, &task.Completed)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(task.Id, task.Title, task.Completed)
		allTasks = append(allTasks, task)
	}
	fmt.Print("\n")
	return &allTasks
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

func GetTaskByTitle(db *sql.DB, i *Task) *[]Task {
	results, err := db.Query("SELECT * FROM tasks WHERE title IS NOT NULL AND title = (?)", &i.Title)
	if err != nil {
		panic(err.Error())
	}
	return queryToArr(results)
}

func GetTaskById(db *sql.DB, i *Task) *[]Task {
	results, err := db.Query("SELECT * FROM tasks WHERE id IS NOT NULL AND id = (?)", &i.Id)
	if err != nil {
		panic(err.Error())
	}
	return queryToArr(results)
}

func GetAllTasks(db *sql.DB) *[]Task {
	results, err := db.Query("SELECT * FROM tasks WHERE id IS NOT NULL")
	if err != nil {
		panic(err.Error())
	}
	return queryToArr(results)
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
	fmt.Printf("\nSuccesfully deleted task at ID: %v\n", i.Id)
}

func CompleteTask(db *sql.DB, i *Task) {
	i.Completed = !i.Completed
	insert, err := db.Query("UPDATE tasks SET completed = (?) WHERE id = (?)", i.Completed, i.Id)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Printf("\nSuccesfully updated completion of task at ID: %v\n", i.Id)
}
