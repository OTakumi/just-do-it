package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Task struct {
	ID   int
	Name string
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=myuser password=mypassword dbname=mydb sslmode=disable")
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to the database!")
	fmt.Println("Starting the application...")

	// Create a table
	createTable(db)

	// Insert a task
	task := Task{Name: "Learn Go"}
	insertTask(db, task)

	// Get all tasks
	tasks := getTasks(db)
	for _, task := range tasks {
		fmt.Println(task)
	}
}

func createTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS tasks (id SERIAL PRIMARY KEY, name TEXT)")
	if err != nil {
		panic(err.Error())
	}
}

func getTasks(db *sql.DB) []Task {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name)
		if err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, task)
	}

	return tasks
}

func insertTask(db *sql.DB, task Task) {
	_, err := db.Exec("INSERT INTO tasks(name) VALUES($1)", task.Name)
	if err != nil {
		panic(err.Error())
	}
}
