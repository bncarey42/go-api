package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type ProjectService struct {
	store Store
}

func NewProjectService(store Store) *ProjectService {
	return &ProjectService{store: store}
}

func (s *ProjectService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", s.handleCreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", s.handleGetProject).Methods("GET")
}

func (s *ProjectService) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	log.Println("call /projects", w, r)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request Payload"})
		return
	}

	defer r.Body.Close()

	var project *Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request Payload"})
		return
	}
	if err := validateProjectPayload(project); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
	}

	p, err := s.store.CreateProject(project)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating task"})
		return
	}

	WriteJSON(w, http.StatusCreated, p)
}

func validateProjectPayload(project *Project) error {
	if project.Name == "" {
		return errNameRequired
	}
	return nil
}

func (s *ProjectService) handleGetProject(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Path
	log.Println("PARAMS", params)
}
