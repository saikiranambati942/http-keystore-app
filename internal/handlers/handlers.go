package handlers

import (
	"github.com/gorilla/mux"
)

// Routes function routes requests to a specific handler based the requestendpoint
func Routes(r *mux.Router) {
	r.HandleFunc("/{key}", storeHandler).Methods("POST")
	r.HandleFunc("/{key}", loadHandler).Methods("GET")
}
