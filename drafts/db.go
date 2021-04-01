package repository

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func init() {
	os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records.
	// SQLite is a file based database.
	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	// defer sqliteDatabase.Close() // Defer Closing the database
	createTable(sqliteDatabase) // Create Database Tables
	r = &repository{
		db: sqliteDatabase,
	}
}

type repository struct {
	db *sql.DB
	mu sync.RWMutex
}

var (
	r *repository
)

func Repository() *repository {
	return r
}

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func createTable(db *sql.DB) {
	createTodoTableSQL := `CREATE TABLE todo (
		"idTodo" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"id" TEXT,
		"title" TEXT,
		"description" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Create Todo table...")
	statement, err := db.Prepare(createTodoTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Todo table created")
}

// We are passing db reference connection from main to our method with other parameters
func (r *repository) InsertTodo(t Todo) {
	log.Println("Inserting Todo record ...")
	insertTodoSQL := `INSERT INTO todo(id, title, description) VALUES (?, ?, ?)`
	statement, err := r.db.Prepare(insertTodoSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(t.Id, t.Title, t.Description)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (r *repository) ListTodos() []Todo {
	row, err := r.db.Query("SELECT * FROM todo ORDER BY title")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	todos := []Todo{}
	for row.Next() { // Iterate and fetch the records from result cursor
		var rowid string
		var id string
		var title string
		var description string
		row.Scan(&rowid, &id, &title, &description)
		todos = append(todos, Todo{id, title, description})
	}
	return todos
}
