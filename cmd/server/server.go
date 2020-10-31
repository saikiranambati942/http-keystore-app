package main

import (
	"http-web-app/internal/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	handlers.Routes(r)
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		log.Fatal("Shutting down the application")
		os.Exit(1)
	}
}
