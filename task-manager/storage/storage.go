package storage

import (
	"github.com/rahulp18/task-manager/task"
)

type Storage interface {
	Load() ([]task.Task, error)
	Save([]task.Task) error
}
