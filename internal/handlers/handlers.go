// Package handler provides handlers for different types of requests.
package handlers

import (
	"time"

	"github.com/gorilla/mux"
)

var m *expiryMap

// init function returns address of expiryMap struct by initializing it and also starts a goroutine which continuously checks
// for every second whether a key value in map has expired i.e, expiry time is set to 30 minutes.
func init() {
	m = &expiryMap{m: make(map[string]*item)}
	go func() {
		for now := range time.Tick(time.Second) {
			m.Lock()
			for k, v := range m.m {
				if now.Unix()-v.storedTime > int64(1800) {
					delete(m.m, k)
				}
			}
			m.Unlock()
		}
	}()
}

// Routes function routes requests to a specific handler based the request.
func Routes(r *mux.Router) {
	r.HandleFunc("/{key}", StoreHandler).Methods("POST")
	r.HandleFunc("/{key}", LoadHandler).Methods("GET")
}
