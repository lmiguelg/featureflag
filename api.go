package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}

}

func (s *APIServer) RUN() {
	router := mux.NewRouter()

	router.HandleFunc("/project", decoratorHTTPHandler(s.handleProject))
	router.HandleFunc("/projects", decoratorHTTPHandler(s.handleGetProjects))
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

func (s *APIServer) handleGetProjects(w http.ResponseWriter, r *http.Request) error {
	projects, err := s.store.GetProjects()
	if err != nil {
		return err
	}

	return responseJSON(w, http.StatusOK, projects)

}

func (s *APIServer) handleGetProject(w http.ResponseWriter, r *http.Request) error {
	project := NewProject("TEST NEW FEATURE PROJECT", false)

	return responseJSON(w, http.StatusOK, project)

}

func (s *APIServer) handleCreateProject(w http.ResponseWriter, r *http.Request) error {
	createProjectReq := &CreateProjectRequest{}
	if err := json.NewDecoder(r.Body).Decode(createProjectReq); err != nil {
		return err
	}

	project := NewProject(createProjectReq.Description, createProjectReq.IsActive)
	if err := s.store.CreateProject(project); err != nil {
		return err
	}

	return responseJSON(w, http.StatusOK, project)
}

func (s *APIServer) handleUpdateProject(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteProject(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// decorator do add an error layer to every http request
func decoratorHTTPHandler(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			responseJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func responseJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type APIServer struct {
	listenAddress string
	store         Storage
}

type APIError struct {
	Error string
}

type ApiFunc func(http.ResponseWriter, *http.Request) error
