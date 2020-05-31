package model

import "time"

// Todo is Todo
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// DBHandler is DB Handler
type DBHandler interface {
	GetTodos(sessionID string) []*Todo
	AddTodo(sessionID string, name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close()
}

var handler DBHandler

// NewDBHandler is New DB Handler
func NewDBHandler(dbConn string) DBHandler {
	// handler = newMemoryHandler()
	// return newSqliteHandler(filepath)
	return newPQHandler(dbConn)
}
