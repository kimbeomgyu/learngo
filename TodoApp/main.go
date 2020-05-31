package main

import (
	"log"
	"net/http"
	"os"
	"todos/app"
)

func main() {
	m := app.MakeHandler(os.Getenv("DATABASE_URL"))
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), m)
	if err != nil {
		panic(err)
	}
}
