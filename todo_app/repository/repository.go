package repository

import "time"

type Todo struct {
	Title       string    `json:"title"`
	Done        bool      `json:"done"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type repository struct {
	db []Todo
}

var (
	r *repository
)

// This function will run one time when package is imported before anything else.
func init() {
	r = &repository{
		db: []Todo{},
	}
}

// Repository, Create public interface
type Repository interface {
	ListTodo() []Todo
	AddTodo(todo *Todo) Todo
}

// Repository, Get instance of singleton repository
func NewRepository() Repository {
	return r
}

func (r *repository) ListTodo() []Todo {
	return r.db
}

func (r *repository) AddTodo(todo *Todo) Todo {
	todo.CreatedAt = time.Now()
	r.db = append(r.db, *todo)
	return *todo
}
