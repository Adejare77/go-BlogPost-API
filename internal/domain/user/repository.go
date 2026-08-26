package user

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(userID entity.UserID) (*UserDetail, error)
	FindAll() ([]UserDetail, error)
	DeleteByID(userID entity.UserID) error
	Update(user *entity.User) (*UserDetail, error)
	EnableByID(user entity.UserID) error
	DisableByID(user entity.UserID) error
	FindByEmail(email string) (*entity.User, error)
}
