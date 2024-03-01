package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//registering services
	tasksService := NewTaskService(s.store)
	tasksService.RegisterRoutes(subrouter)
	projectService := NewProjectService(s.store)
	projectService.RegisterRoutes(subrouter)

	log.Println("Starting the API Server a port", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
