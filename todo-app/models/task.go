package models

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          string         `db:"id" json:"id"`
	Title       string         `db:"title" json:"title"`
	Description sql.NullString `db:"description" json:"description"`
	Completed   bool           `db:"completed" json:"completed"`
	UserID      string         `db:"user_id" json:"user_id"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
}
