package store

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/rahulp18/todo/models"
)

type FileTaskStore struct {
	mu       sync.Mutex
	filePath string
}

func NewFileTaskStore(path string) *FileTaskStore {
	// Ensure file exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.WriteFile(path, []byte("{}"), 0644)
	}
	return &FileTaskStore{
		filePath: path,
	}
}
func (fs *FileTaskStore) Read() (map[string]models.Task, error) {
	data, err := os.ReadFile(fs.filePath)
	if err != nil {
		return nil, err
	}
	var tasks map[string]models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
func (fs *FileTaskStore) Write(tasks map[string]models.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.filePath, data, 0644)
}

// Implement interfaces Now
func (fs *FileTaskStore) Create(task models.Task) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	tasks, err := fs.Read()
	if err != nil {
		return err
	}
	tasks[task.ID] = task
	return fs.Write(tasks)
}
func (fs *FileTaskStore) GetAll() ([]models.Task, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	tasks, err := fs.Read()
	if err != nil {
		return nil, err
	}
	result := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, task)
	}
	return result, nil
}

func (fs *FileTaskStore) GetById(id string) (models.Task, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	tasks, err := fs.Read()
	if err != nil {
		return models.Task{}, err
	}
	task, ok := tasks[id]
	if !ok {
		return models.Task{}, ErrorTaskNotFound
	}
	return task, nil
}

func (fs *FileTaskStore) Update(id string, task models.Task) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	tasks, err := fs.Read()
	if err != nil {
		return err
	}
	if _, ok := tasks[id]; !ok {
		return ErrorTaskNotFound
	}
	tasks[id] = task
	return fs.Write(tasks)

}
func (fs *FileTaskStore) Delete(id string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	tasks, err := fs.Read()
	if err != nil {
		return err
	}
	if _, ok := tasks[id]; !ok {
		return ErrorTaskNotFound
	}
	delete(tasks, id)
	return fs.Write(tasks)
}
