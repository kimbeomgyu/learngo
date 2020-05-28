package main

import (
	"learngo/TodoApp/app"
	"log"
	"net/http"
)

func main() {
	m := app.MakeHandler("./TodoApp.db")
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":3000", m)
	if err != nil {
		panic(err)
	}
}
