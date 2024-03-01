package main

import (
	"database/sql"
)

type Store interface {
	CreateUser(user *User) (*User, error)
	CreateTask(task *Task) (*Task, error)
	CreateProject(project *Project) (*Project, error)
	GetTask(id string) (*Task, error)
	GetProject(id string) (*Project, error)
	GetUserByID(id string) (*User, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateTask(task *Task) (*Task, error) {
	rows, err := s.db.Exec(
		"INSERT INTO projectmanager.tasks (name, status, projectId, assignedToID) values (?, ?, ?, ?)",
		task.Name, task.Status, task.ProjectID, task.AssignedToID)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	task.ID = id

	return task, nil
}

func (s *Storage) CreateProject(project *Project) (*Project, error) {
	rows, err := s.db.Exec(
		"insert into projectmanager.projects (name) values (?)", project.Name)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	project.ID = id
	return project, nil
}

func (s *Storage) GetTask(id string) (*Task, error) {
	var t Task
	err := s.db.QueryRow("select id, name, status, projectId, assignedToID, createAt from projectmanager.tasks where id = ?",
		id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedToID, &t.CreatedAt)
	return &t, err
}

func (s *Storage) CreateUser(user *User) (*User, error) {
	rows, err := s.db.Exec(
		"insert into projectmanager.users (firstname, lastname, password) values (?,?,?)", user.Firstname, user.Lastname, user.Password)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

func (s *Storage) GetUserByID(id string) (*User, error) {
	var u User
	err := s.db.QueryRow("SELECT id, firstname, lastname, password, createAt from projectmanager.users where id = ?",
		id).Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Password, &u.CreateAt)
	return &u, err
}

func (s *Storage) GetProject(id string) (*Project, error) {
	var project Project
	err := s.db.QueryRow("select id, name, createAt from projectmanager.projects where id = ?",
		id).Scan(&project.ID, &project.Name, &project.CreatedAt)
	return &project, err
}
