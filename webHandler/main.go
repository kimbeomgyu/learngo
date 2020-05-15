package main

import (
	"learngo/webHandler/myapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
