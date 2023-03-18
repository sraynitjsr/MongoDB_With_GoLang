package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	// Create a request to GET all users
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the getUsers function with the GET request and ResponseRecorder
	handler := http.HandlerFunc(getUsers)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("getUsers returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[]`
	if rr.Body.String() != expected {
		t.Errorf("getUsers returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
