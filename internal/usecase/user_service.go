package usecase

import (
	"fmt"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo user.UserRepository
}


func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) FindAll() ([]user.UserDetail, error) {
	return s.repo.FindAll()
}

func (s *UserService) FindByID(targetUserID entity.UserID) (*user.UserDetail, error) {
	return s.repo.FindByID(targetUserID)
}

func (s *UserService) DeleteByID(targetUserID entity.UserID) error {
	return s.repo.DeleteByID(targetUserID)
}

func (s *UserService) Update(user *entity.User) (*user.UserDetail, error) {
	if user.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("hash password: %w", err)
		}
		password := string(hash)
		user.Password = &password
	}

	return s.repo.Update(user)
}

func (s *UserService) EnableByID(targetUserID entity.UserID) error {
	return s.repo.EnableByID(targetUserID)
}

func (s* UserService) DisableByID(targetUserID entity.UserID) error {
	return s.repo.DisableByID(targetUserID)
}
