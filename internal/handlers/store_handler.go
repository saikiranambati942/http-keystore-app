package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// message struct is used to encode a successful message string.
type message struct {
	Message string `json:"message"`
}

// errorResponse struct is used to encode an error message string.
type errorResponse struct {
	Error string `json:"error"`
}

// item struct is used to store the value of a key with its stored time.
type item struct {
	value      string
	storedTime int64
}

// expiryMap is a struct which contains a map with RWMutex embedded in to it for concurrent safety access.
type expiryMap struct {
	m map[string]*item
	sync.RWMutex
}

// Store method is to store a value with respect to a key.
func (m *expiryMap) store(k, v string) {
	m.Lock()
	it, ok := m.m[k]
	if !ok {
		it = &item{value: v}
		m.m[k] = it
	}
	it.storedTime = time.Now().Unix()
	m.Unlock()
}

// StoreHandler is for handling the POST requests to store a value with respect to a key.
func StoreHandler(w http.ResponseWriter, r *http.Request) {
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
	m.store(key, v.Value)
	m := message{Message: "key stored successfully"}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}
