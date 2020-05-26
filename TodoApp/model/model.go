package model

import (
	"time"
)

// Todo is Todo
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo

func init() {
	todoMap = make(map[int]*Todo)
}

// GetTodos is get todo list
func GetTodos() []*Todo {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	return list
}

// AddTodo is add todo item
func AddTodo(name string) *Todo {
	id := len(todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	todoMap[id] = todo
	return todo
}

// RemoveTodo is remove todo item
func RemoveTodo(id int) bool {
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		return true
	}
	return false
}

// CompleteTodo is complete todo item
func CompleteTodo(id int, complete bool) bool {
	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		return true
	}
	return false
}
