package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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
		Value: val.(string),
	}
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(v)
}
