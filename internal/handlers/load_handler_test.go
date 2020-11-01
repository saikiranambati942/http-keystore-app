package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadHandlerPositive(t *testing.T) {
	var v value
	x := `{
		"value":"v1"
	}`

	r := httptest.NewRequest(http.MethodPost, "/k1", strings.NewReader(x))
	w := httptest.NewRecorder()
	Router().ServeHTTP(w, r)
	r1 := httptest.NewRequest(http.MethodGet, "/k1", nil)
	w1 := httptest.NewRecorder()
	Router().ServeHTTP(w1, r1)
	json.Unmarshal(w1.Body.Bytes(), &v)
	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, "v1", v.Value)
}

func TestLoadHandlerNegative(t *testing.T) {
	var e errorResponse
	r := httptest.NewRequest(http.MethodGet, "/k2", nil)
	w := httptest.NewRecorder()
	Router().ServeHTTP(w, r)
	json.Unmarshal(w.Body.Bytes(), &e)
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "key not found", e.Error)
}
