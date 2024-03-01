package main

type MockStore struct{}

func (m *MockStore) CreateUser(u *User) (*User, error) {
	return nil, nil
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
func (m *MockStore) GetProject(id string) (*Project, error) {
	return nil, nil
}

func (m *MockStore) GetUserByID(id string) (*User, error) {
	return nil, nil
}
