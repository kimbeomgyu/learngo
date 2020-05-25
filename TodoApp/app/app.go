package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

// MakeHandler is app handler
func MakeHandler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	return r
}
