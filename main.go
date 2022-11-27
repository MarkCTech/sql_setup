package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	Title     string `json:"title"`
	Completed string `json:"completed"`
}

type Input struct {
	Title string
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
	fmt.Println("Succesfully connected to MySQL database on :3306")

	var i Input
	InputTask(&i)
	AddTask(*db, i.Title)
	GetTasks(*db)
}

func InputTask(i *Input) {
	fmt.Println("Enter title of task:")
	fmt.Scanln(&i.Title)
}

func GetTasks(db sql.DB) {
	results, err := db.Query("SELECT title FROM tasks")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Retrieved these tasks:")
	for results.Next() {
		var task Task
		err = results.Scan(&task.Title)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(task.Title, task.Completed)
	}
}

func AddTask(db sql.DB, intitle string) {
	insert, err := db.Query("INSERT INTO tasks (title,completed) VALUES(?,false)", intitle)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Succesfully inserted into tasks")
}
