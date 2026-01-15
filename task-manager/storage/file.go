package storage

import (
	"encoding/json"
	"os"

	"github.com/rahulp18/task-manager/task"
)

type FileStorage struct {
	Filename string
}

func (f FileStorage) Load() ([]task.Task, error) {
	data, err := os.ReadFile(f.Filename)
	if err != nil {
		return []task.Task{}, nil
	}
	var tasks []task.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func (f FileStorage) Save(tasks []task.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(f.Filename, data, 0644)
}
