package service

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/rahulp18/todo/models"
	"github.com/rahulp18/todo/store"
)

type TaskService struct {
	store store.TaskStore
}

func NewTaskService(store store.TaskStore) *TaskService {
	return &TaskService{
		store: store,
	}
}
func (s *TaskService) CreateTask(title string, description sql.NullString, userID string) (models.Task, error) {
	task := models.Task{
		ID:          uuid.NewString(),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UserID:      userID,
	}
	err := s.store.Create(task)
	return task, err
}

func (s *TaskService) GetAllTasks(userID string) ([]models.Task, error) {
	return s.store.GetAll(userID)
}
func (s *TaskService) GetTaskById(id string) (models.Task, error) {
	return s.store.GetById(id)
}
func (s *TaskService) DeleteTask(id string) error {
	return s.store.Delete(id)
}
func (s *TaskService) UpdateTask(id string, completed bool) error {
	task, err := s.GetTaskById(id)
	if err != nil {
		return err
	}
	task.Completed = completed

	return s.store.Update(id, task)
}
