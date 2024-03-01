package main

import (
	"errors"
	"time"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Task struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ProjectID    int64     `json:"projectID"`
	AssignedToID int64     `json:"assignedTo"`
	CreatedAt    time.Time `json:"createdAt"`
}

type User struct {
	ID        int64     `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Password  string    `json:"password"`
	CreateAt  time.Time `json:"createAt"`
}

type Project struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

var errNameRequired = errors.New("name is required")
var errProjectIDRequired = errors.New("projectID is required")
var errUserIDRequired = errors.New("userID is required")
var errPasswordRequired = errors.New("password is required")
