package service

import (
	"github.com/danush754/Task-Manager/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) HashPassword(password string) (hashedPassword string, err error) {

	hashValue, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashValue), err
}

func (s *UserService) CheckPassword(password, hashValue string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashValue), []byte(password)) == nil
}
