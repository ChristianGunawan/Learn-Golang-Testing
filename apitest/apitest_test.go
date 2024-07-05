package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users" {
			t.Errorf("Expected to request '/api/users', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"page": 2, "data": [{"id": 1, "name": "John"}]}`))
	}))

	defer ts.Close()

	originalBaseURL := baseURL
	baseURL = ts.URL
	defer func() {
		baseURL = originalBaseURL
	}()

	users, err := getUsers()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{"page": 2, "data": [{"id": 1, "name": "John"}]}`

	if string(users) != expected {
		t.Errorf("Expected %s, got %s", expected, string(users))
	}
}

func TestPostUsers(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users" {
			t.Errorf("Expected to request '/api/users', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "John", "job": "Developer"}`))
	}))
	defer ts.Close()

	originalBaseURL := baseURL
	baseURL = ts.URL
	defer func() {
		baseURL = originalBaseURL
	}()

	user, err := createUser("Jogn", "Developer")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{"name": "John", "job": "Developer"}`
	if string(user) != expected {
		t.Errorf("Expected %s, got %s", expected, string(user))
	}
}
