package handlers

import (
	"time"

	"github.com/gorilla/mux"
)

var m *TTLMap

func init() {
	m = &TTLMap{m: make(map[string]*item)}
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

// Routes function routes requests to a specific handler based the requestendpoint
func Routes(r *mux.Router) {
	r.HandleFunc("/{key}", storeHandler).Methods("POST")
	r.HandleFunc("/{key}", loadHandler).Methods("GET")
}
