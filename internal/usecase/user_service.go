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


func (s *UserService) FindAll() ([]entity.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) Create(user *entity.User) error {
	if user.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}
		password := string(hash)
		user.Password = &password
	}

	// log user created
	return s.repo.Create(user)
}

func (s *UserService) FindByID(userID entity.UserID) (*entity.User, error) {
	return s.repo.FindByID(userID)
}

func (s *UserService) DeleteByID(userID entity.UserID) error {
	// log user deletion
	return s.repo.DeleteByID(userID)
}

func (s *UserService) Update(user *entity.User) (*entity.User, error) {
	// log updated user
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
