package user

import "github.com/Adejare77/go-BlogPost-API/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(userID entity.UserID) (*entity.User, error)
	FindAll() ([]entity.User, error)
	DeleteByID(userID entity.UserID) error
	Update(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
