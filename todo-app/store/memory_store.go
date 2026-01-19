package store

import (
	"sync"

	"github.com/rahulp18/todo/models"
)

type MemoryTaskStore struct {
	mu    sync.Mutex
	tasks map[string]models.Task
}

func NewMemoryTaskStore() *MemoryTaskStore {
	return &MemoryTaskStore{
		tasks: make(map[string]models.Task),
	}
}

func (m *MemoryTaskStore) Create(task models.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.tasks[task.ID] = task
	return nil
}
func (m *MemoryTaskStore) GetAll() ([]models.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	result := make([]models.Task, 0, len(m.tasks))
	for _, task := range m.tasks {
		result = append(result, task)
	}
	return result, nil
}
func (m *MemoryTaskStore) GetById(id string) (models.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	task, ok := m.tasks[id]
	if !ok {
		return models.Task{}, ErrorTaskNotFound
	}
	return task, nil
}
func (m *MemoryTaskStore) Update(id string, task models.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.tasks[id]
	if !ok {
		return ErrorTaskNotFound
	}
	m.tasks[id] = task
	return nil
}
func (m *MemoryTaskStore) Delete(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.tasks[id]
	if !ok {
		return ErrorTaskNotFound
	}
	delete(m.tasks, id)
	return nil
}
