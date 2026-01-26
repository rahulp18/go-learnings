package store

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rahulp18/todo/models"
)

type PostgresStore struct {
	dbCon *sqlx.DB
}

func NewPgTaskStore(db *sqlx.DB) *PostgresStore {
	return &PostgresStore{
		dbCon: db,
	}
}
func (s *PostgresStore) Create(task models.Task) error {
	query := `INSERT INTO tasks(id,title,description,completed,user_id) VALUES ($1,$2,$3,$4,$5)`
	_, err := s.dbCon.Exec(
		query,
		task.ID,
		task.Title,
		task.Description,
		task.Completed,
		task.UserID,
	)
	return err
}

func (s *PostgresStore) GetAll(userID string) ([]models.Task, error) {

	var tasks []models.Task
	query := `SELECT *
	FROM tasks
	ORDER BY created_at DESC`
	err := s.dbCon.Select(&tasks, query)
	if err != nil {
		return nil, err
	}
	return tasks, nil

}

func (s *PostgresStore) GetById(id string) (models.Task, error) {
	var task models.Task
	query := `SELECT * FROM tasks WHERE id = $1`
	err := s.dbCon.Get(&task, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Task{}, ErrorTaskNotFound
		}
		return models.Task{}, err
	}
	return task, nil
}

func (s *PostgresStore) Update(id string, task models.Task) error {

	query := `UPDATE tasks SET title = $1, completed = $2 WHERE id= $3`

	result, err := s.dbCon.Exec(query,
		task.Title,
		task.Completed,
		id,
	)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrorTaskNotFound
	}
	return nil
}
func (s *PostgresStore) Delete(id string) error {
	query := `DELETE FROM tasks WHERE id=$1`

	result, err := s.dbCon.Exec(query, id)

	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrorTaskNotFound
	}
	return nil
}
