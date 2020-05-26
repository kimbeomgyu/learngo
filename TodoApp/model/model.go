package model

import "time"

// Todo is Todo
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type dbHandler interface {
	getTodos() []*Todo
	addTodo(name string) *Todo
	removeTodo(id int) bool
	completeTodo(id int, complete bool) bool
}

var handler dbHandler

func init() {
	handler = newMemoryHandler()
}

// GetTodos is get todo list
func GetTodos() []*Todo {
	return handler.getTodos()
}

// AddTodo is add todo item
func AddTodo(name string) *Todo {
	return handler.addTodo(name)
}

// RemoveTodo is remove todo item
func RemoveTodo(id int) bool {
	return handler.removeTodo(id)
}

// CompleteTodo is complete todo item
func CompleteTodo(id int, complete bool) bool {
	return handler.completeTodo(id, complete)
}
