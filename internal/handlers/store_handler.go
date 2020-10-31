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

type value struct {
	Value string `json:"value"`
}

var m sync.Map

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
	f := func() {
		m.Delete(key)
	}
	time.AfterFunc(30*time.Minute, f)
	m := message{Message: "key stored successfully"}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}
