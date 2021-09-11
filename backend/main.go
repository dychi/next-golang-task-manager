package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type ResponseString string

func (rs ResponseString) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rs)
}

type Task struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
}

func main() {
	fmt.Println("HTTP Server starting...")
	// http.Handle("/api", ResponseString("Hello World!"))
	// http.ListenAndServe(":8080", nil)

	fmt.Println("MySQL Section started.")
	db, err := sql.Open("mysql", "apiserver:apipassword@tcp(127.0.0.1:3306)/task_mysql")

	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	fmt.Println("Connection started.")

	// Insert
	stmtInsert, err := db.Prepare("insert into task (name) values (?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsert.Close()
	insertResult, err := stmtInsert.Exec("task from statement")
	if err != nil {
		panic(err.Error())
	}
	lastInsertID, err := insertResult.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(lastInsertID)

	// Select
	results, err := db.Query("select id, name, created_at from task")
	if err != nil {
		log.Fatal(err.Error())
	}
	for results.Next() {
		var task Task
		err = results.Scan(&task.ID, &task.Name, &task.Created_at)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println(task)
	}
}
