package store

import (
	"errors"

	"github.com/rahulp18/todo/models"
)

var ErrorUserNotFound = errors.New("User not Found")

type UserStore interface {
	Create(user models.User) error
	GetByEmail(email string) (models.User, error)
}
