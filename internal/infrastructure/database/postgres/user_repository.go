package postgres

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	domainErrors "github.com/Adejare77/go-BlogPost-API/internal/domain/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/user"
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
	return MapError(repo.db.Create(user).Error)
}

func (repo *UserRepository) FindByID(userID entity.UserID) (*user.UserDetail, error) {
	var user user.UserDetail

	if err := repo.db.First(&user, userID).Error; err != nil {
		return nil, MapError(err)
	}

	return &user, nil
}

func (repo *UserRepository) FindAll() ([]user.UserDetail, error) {
	var users []user.UserDetail

	return users, MapError(repo.db.Find(&users).Error)
}

func (repo *UserRepository) DeleteByID(userID entity.UserID) error {
	return MapError(repo.db.Delete(&entity.User{}, userID).Error)
}

func (repo *UserRepository) Update(user *entity.User) (*user.UserDetail, error) {

	result := repo.db.Model(&entity.User{}).
	Where("id = ?", user.ID).
	Updates(user)

	if result.Error != nil {
		return nil, MapError(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, domainErrors.ErrNotFound
	}

	return repo.FindByID(user.ID)
}

func (repo *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := repo.db.
	Where("email = ?", email).First(&user).Error; err != nil {
		return nil, MapError(err)
	}

	return &user, nil
}

func (repo *UserRepository) EnableByID(userID entity.UserID) (error) {
	return MapError(repo.db.Model(&entity.User{}).
	Where("id = ?", userID).
	Update("is_active", true).Error)
}

func (repo *UserRepository) DisableByID(userID entity.UserID) error {
	return MapError(repo.db.Model(&entity.User{}).
	Where("id = ?", userID).
	Update("is_active", false).Error)
}
