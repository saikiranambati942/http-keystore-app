// Command server ...
package main

import (
	"http-keystore-app/internal/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// main function is the application starting point.
func main() {
	r := mux.NewRouter()
	handlers.Routes(r)
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		log.Fatal("Shutting down the application")
		os.Exit(1)
	}
}
