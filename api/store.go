package main

import "database/sql"

type Store interface {
	CreateUser() error
	CreateTask(task *Task) (*Task, error)
	CreateProject(project *Project) (*Project, error)
	GetTask(id string) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateUser() error {
	return nil
}

func (s *Storage) CreateTask(task *Task) (*Task, error) {
	rows, err := s.db.Exec(
		"INSERT INTO projectmanager.tasks (name, status, project_id, assigned_to) values (?, ?, ?, ?)",
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
	err := s.db.QueryRow("select id, name, status, projectId, assignedToID, createAt from projectmanager.tasks where id = ?", id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedToID, &t.CreatedAt)
	return &t, err
}
