package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {

	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	hnd := http.HandlerFunc(statusHandler)
	hnd.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("incorrect status code: got %v want %v", status, http.StatusNotImplemented)
	}
}
