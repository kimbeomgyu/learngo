package main

import (
	"learngo/RESTful_API/myapp"
	"learngo/webDecoratorHandler/decohandler"
	"log"
	"net/http"
	"time"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Print("[LOGGER1] Completed time: ", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Print("[LOGGER2] Completed time: ", time.Since(start).Milliseconds())
}

// NewHandler is NewHandler
func NewHandler() http.Handler {
	h := myapp.NewHandler()
	h = decohandler.NewDecoHandler(h, logger)
	h = decohandler.NewDecoHandler(h, logger)
	h = decohandler.NewDecoHandler(h, logger)
	h = decohandler.NewDecoHandler(h, logger2)
	return h
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}
