package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type value struct {
	Value string `json:"value"`
}

func (m *TTLMap) Load(k string) (v string, ok bool) {
	m.l.RLock()
	var it *item
	if it, ok = m.m[k]; ok {
		v = it.value
	}
	m.l.RUnlock()
	return
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	val, ok := m.Load(key)
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
