package main

import (
	"learngo/TodoApp/app"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler("./TodoApp.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
