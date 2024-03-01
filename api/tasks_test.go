package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRoutePOST(payload *Task, serviceTask func(w http.ResponseWriter, r *http.Request)) (*httptest.ResponseRecorder, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))

	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc("/tasks", serviceTask)
	router.ServeHTTP(rr, req)
	return rr, nil
}
func testRouteGET(serviceTask func(w http.ResponseWriter, r *http.Request)) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(http.MethodGet, "/tasks/42", nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc("/tasks/{id}", serviceTask)
	router.ServeHTTP(rr, req)
	return rr, nil
}

func TestCreateTask(t *testing.T) {
	ms := &MockStore{}

	service := NewTaskService(ms)

	t.Run("Should create task", func(t *testing.T) {
		rr, err := testRoutePOST(
			&Task{Name: "Hello", ProjectID: 5, AssignedToID: 42},
			service.handleCreateTask,
		)
		if err != nil {
			t.Fatal(err)
		}
		if rr.Code != http.StatusCreated {
			t.Error("Invalid status code, it should create")
		}
	})

	t.Run("Should return an error if name is empty", func(t *testing.T) {
		rr, err := testRoutePOST(
			&Task{Name: ""},
			service.handleCreateTask,
		)
		if err != nil {
			t.Fatal(err)
		}

		if rr.Code != http.StatusBadRequest {
			t.Error("Invalid status code, it should fail")
		}
	})
}

func TestGetTask(t *testing.T) {
	ms := &MockStore{}
	service := NewTaskService(ms)

	t.Run("Should return the task", func(t *testing.T) {
		rr, err := testRouteGET(service.handleGetTask)
		if err != nil {
			t.Fatal(err)
		}
		if rr.Code != http.StatusOK {
			t.Error("invalid status code")
		}
	})

}
