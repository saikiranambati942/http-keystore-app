package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// value struct is used the unmarshal the value with respect to a key.
type value struct {
	Value string `json:"value"`
}

// load method is to retrieve a value of a given key.
func (m *expiryMap) load(k string) (v string, ok bool) {
	m.RLock()
	var it *item
	if it, ok = m.m[k]; ok {
		v = it.value
	}
	m.RUnlock()
	return
}

// LoadHandler is for handling the GET requests with respect to a key.
func LoadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	val, ok := m.load(key)
	if !ok {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		er := errorResponse{
			Error: "key not found"}
		json.NewEncoder(w).Encode(er)
		return
	}
	v := value{
		Value: val,
	}
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(v)
}
