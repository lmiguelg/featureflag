package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func responseJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type APIServer struct {
	listenAddress string
}

type APIError struct {
	Error string
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

// decorator do add an error layer to every http request
func decoratorHTTPHandler(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			responseJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

func (s *APIServer) RUN() {
	router := mux.NewRouter()

	router.HandleFunc("/project", decoratorHTTPHandler(s.handleProject))

	http.ListenAndServe(s.listenAddress, router)
}

func (s *APIServer) handleProject(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetProject(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateProject(w, r)
	}
	if r.Method == "PUT" {
		return s.handleUpdateProject(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteProject(w, r)
	}
	return fmt.Errorf("Method not allowed %s", r.Method)
}

func (s *APIServer) handleGetProject(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("asdasdas")

	return nil
}

func (s *APIServer) handleCreateProject(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleUpdateProject(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteProject(w http.ResponseWriter, r *http.Request) error {
	return nil
}
