package sql_db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func Dbstart() *sql.DB {
	db, err := sql.Open("mysql", Dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)

	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)

	}
	log.Printf("rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", Dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
	}
	log.Printf("Connected to DB %s successfully\n", dbname)
	CreateTable(*db, "tasks")
	return db
}

func CreateTable(db sql.DB, tableName string) {
	db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %v", tableName))
	db.Exec(fmt.Sprintf("CREATE TABLE %v (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, title VARCHAR(100) NOT NULL, completed BOOLEAN NOT NULL DEFAULT 0)", tableName))
	db.Query(fmt.Sprintf("INSERT INTO %v (title) VALUES ('test')", tableName))
}

func Dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
