package xrequestid

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNonExistXRequestIDInHeader(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	middleware := New(16)
	middleware.Generate = func(n int) (string, error) { return "test-id", nil }
	middleware.ServeHTTP(recorder, req, func(w http.ResponseWriter, r *http.Request) {})

	if id := req.Header.Get("X-Request-ID"); id != "test-id" {
		t.Fatalf("Expected request X-Request-Id to be `test-id`, got `%v`", id)
	}

	if responseID := recorder.HeaderMap.Get("X-Request-ID"); responseID != "test-id" {
		t.Fatalf("Expected response X-Request-Id to be `test-id`, got `%v`", responseID)
	}
}

func TestExistXRequestIDInHeader(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Request-ID", "test-id")

	middleware := New(16)
	middleware.ServeHTTP(recorder, req, func(w http.ResponseWriter, r *http.Request) {})

	if id := req.Header.Get("X-Request-ID"); id != "test-id" {
		t.Fatalf("Expected request X-Request-Id to be `test-id`, got `%v`", id)
	}

	if responseID := recorder.HeaderMap.Get("X-Request-ID"); responseID != "test-id" {
		t.Fatalf("Expected response X-Request-Id to be `test-id`, got `%v`", responseID)
	}
}
