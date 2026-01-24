package store

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rahulp18/todo/models"
)

type PostgresUserStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db: db,
	}
}

func (us *PostgresUserStore) Create(user models.User) error {
	query := `INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`

	_, err := us.db.Exec(query, user.ID, user.Name, user.Email, user.Password)
	return err
}

func (us *PostgresUserStore) GetByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = $1`
	err := us.db.Get(&user, query, email)
	if err == sql.ErrNoRows {
		return models.User{}, ErrorUserNotFound
	}
	return user, err

}
