package main

import (
	"net/http"
)

type APIServer struct {
	listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

func (s *APIServer) RUN() {

}

func (s *APIServer) handleProjects(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetProject(w http.ResponseWriter, r *http.Request) error {
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
