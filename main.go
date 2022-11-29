package sql_db

import (
	"database/sql"

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

type Database struct {
	Db *sql.DB
}

const (
	username = "root"
	password = "toor"
	hostname = "127.0.0.1:3306"
	dbname   = "todosdb"
)

func startdb() {
	StartedDb := dbstart()
	defer StartedDb.Close()
	CrudDb = StartedDb
	crudin1()
}
