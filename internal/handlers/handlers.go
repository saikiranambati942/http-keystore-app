// Package handler provides handlers for different types of requests.
package handlers

import (
	"time"

	"github.com/gorilla/mux"
)

var m *expiryMap

func init() {
	m = &expiryMap{m: make(map[string]*item)}
	go func() {
		for now := range time.Tick(time.Second) {
			m.l.Lock()
			for k, v := range m.m {
				if now.Unix()-v.storedTime > int64(1800) {
					delete(m.m, k)
				}
			}
			m.l.Unlock()
		}
	}()
}

// Routes function routes requests to a specific handler based the request.
func Routes(r *mux.Router) {
	r.HandleFunc("/{key}", StoreHandler).Methods("POST")
	r.HandleFunc("/{key}", LoadHandler).Methods("GET")
}
