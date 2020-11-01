package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{key}", StoreHandler).Methods("POST")
	r.HandleFunc("/{key}", LoadHandler).Methods("GET")
	return r
}

func TestStoreHandlerPositive(t *testing.T) {
	var m message
	x := `{
		"value":"v1"
	}`

	r, _ := http.NewRequest(http.MethodPost, "/k1", strings.NewReader(x))
	w := httptest.NewRecorder()
	Router().ServeHTTP(w, r)
	if w.Code != 200 {
		t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w.Code)
	}
	json.Unmarshal(w.Body.Bytes(), &m)
	assert.Equal(t, "key stored successfully", m.Message)
}

func TestStoreHandlerNegative(t *testing.T) {
	var e errorResponse
	x := `{
		"value":"v1" @
	}`

	r, _ := http.NewRequest(http.MethodPost, "/k1", strings.NewReader(x))
	w := httptest.NewRecorder()
	Router().ServeHTTP(w, r)
	assert.Equal(t, 400, w.Code)
	json.Unmarshal(w.Body.Bytes(), &e)
	assert.Equal(t, "error occured while decoding json input", e.Error)

}
