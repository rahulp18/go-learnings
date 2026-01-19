package store

import (
	"errors"

	"github.com/rahulp18/todo/models"
)

var ErrorTaskNotFound = errors.New("Task not Found")

type TaskStore interface {
	Create(task models.Task) error
	GetAll() ([]models.Task, error)
	GetById(id string) (models.Task, error)
	Update(id string, task models.Task) error
	Delete(id string) error
}
