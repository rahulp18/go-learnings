package service

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/rahulp18/todo/models"
	"github.com/rahulp18/todo/store"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userStore store.UserStore
}

func NewAuthService(us store.UserStore) *AuthService {
	return &AuthService{
		userStore: us,
	}

}
func (s *AuthService) Register(name, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		ID:       uuid.NewString(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}

	return s.userStore.Create(user)
}

func (s *AuthService) Login(email, password string) (models.User, error) {
	user, err := s.userStore.GetByEmail(email)
	if err != nil {
		return models.User{}, errors.New("Invalid credentials")
	}
	log.Println(user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("Invalid credentials")
	}
	return user, nil
}
