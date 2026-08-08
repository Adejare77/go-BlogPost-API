package postgres

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(user *entity.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) FindByID(userID entity.UserID) (*entity.User, error) {
	var user entity.User

	err := repo.db.First(&user, userID).Error

	return &user, err
}

func (repo *UserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User

	err := repo.db.Find(&users).Error

	return users, err
}

func (repo *UserRepository) DeleteByID(userID entity.UserID) error {
	var user entity.User
	return repo.db.Delete(&user, userID).Error
}

func (repo *UserRepository) Update(user *entity.User) (*entity.User, error) {
	var updateUser entity.User

	if err := repo.db.Model(&entity.User{}).
	Where("id = ?", user.ID).
	Updates(user).Error; err != nil {
		return nil, err
	}

	if err := repo.db.First(&updateUser, user.ID).Error; err != nil {
		return nil, err
	}

	return &updateUser, nil
}
