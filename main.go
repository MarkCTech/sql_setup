package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Input struct {
	Task
}

type Returned struct {
	Task
}

func main() {

	// Opens up connection to already running MySQL database
	// Database is called todos
	// Table is called tasks
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todos")
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	defer db.Close()
	fmt.Println("\nSuccesfully connected to MySQL database on :3306")

	var i Input
	var r Returned
	InputTask(*db, &i)
	AddTask(*db, &i)
	GetTaskbyTitle(*db, &i, &r)
	ReturnToInput(&r, &i)

	CompleteTask(*db, &i)
	GetAllTasks(*db)
}

func InputTask(db sql.DB, i *Input) {
	fmt.Print("\nEnter title of task: ")
	fmt.Scanln(&i.Title)
}

func ReturnToInput(r *Returned, i *Input) {
	i.Id = r.Id
	i.Title = r.Title
	i.Completed = r.Completed
}

func GetTaskbyTitle(db sql.DB, i *Input, r *Returned) {
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

func GetAllTasks(db sql.DB) {
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

func AddTask(db sql.DB, i *Input) {
	insert, err := db.Query("INSERT INTO tasks (title,completed) VALUES(?,false)", i.Title)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Printf("\nSuccesfully inserted: %v\n", i.Title)
}

func CompleteTask(db sql.DB, i *Input) {
	i.Completed = true
	insert, err := db.Query("UPDATE tasks SET completed = (?) WHERE id = (?)", i.Completed, i.Id)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Printf("\nSuccesfully completed task at ID: %v\n", i.Id)
}
