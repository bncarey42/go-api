package main

type MockStore struct{}

func (m *MockStore) CreateUser() error {
	return nil
}
func (m *MockStore) CreateTask(t *Task) (*Task, error) {
	return nil, nil
}
func (m *MockStore) CreateProject(p *Project) (*Project, error) {
	return nil, nil
}
func (m *MockStore) GetTask(id string) (*Task, error) {
	return nil, nil
}
