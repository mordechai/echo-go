package server

import (
	// "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestEcho tests the echo function
func TestEcho(t *testing.T) {
	s := NewServer()
	req, err := http.NewRequest("GET", "/echo", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.echo)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "echo!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
