package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type message struct {
	Message string `json:"message"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type item struct {
	value      string
	storedTime int64
}
type TTLMap struct {
	m map[string]*item
	l sync.RWMutex
}

func (m *TTLMap) Store(k, v string) {
	m.l.Lock()
	it, ok := m.m[k]
	if !ok {
		it = &item{value: v}
		m.m[k] = it
	}
	it.storedTime = time.Now().Unix()
	m.l.Unlock()
}

func storeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	var v value
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		er := errorResponse{
			Error: "error occured while decoding json input"}
		json.NewEncoder(w).Encode(er)
		return
	}
	m.Store(key, v.Value)
	m := message{Message: "key stored successfully"}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}
