package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{key}", postHandler).Methods("POST")
	r.HandleFunc("/{key}", getHandler).Methods("GET")
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		log.Fatal("Shutting down the application")
		os.Exit(1)
	}
}

type value struct {
	Value string `json:"value"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	val, ok := m[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	v := value{
		Value: val,
	}
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(v)
}

var m = make(map[string]string)

func postHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	var v value
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m[key] = v.Value
	f := func() {
		delete(m, key)
	}
	time.AfterFunc(30*time.Second, f)

}
