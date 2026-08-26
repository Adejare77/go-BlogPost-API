package postgres

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	domainErrors "github.com/Adejare77/go-BlogPost-API/internal/domain/errors"
	"gorm.io/gorm"
)


type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (repo *RefreshTokenRepository) Create(token *entity.RefreshToken) error {
	return MapError(repo.db.Create(token).Error)
}

func (repo *RefreshTokenRepository) RevokeToken(tokenHash string) error {
	result := repo.db.Model(&entity.RefreshToken{}).
	Where("token_hash = ? AND revoked_at IS NULL", tokenHash).
	Update("revoked_at", time.Now())

	if result.RowsAffected == 0 {
		return domainErrors.ErrTokenRevoked
	}

	return nil
}

func (repo *RefreshTokenRepository) FindByTokenHash(tokenHash string) (*entity.RefreshToken, error) {
	var token entity.RefreshToken

	if err := repo.db.
	Where("token_hash = ?", tokenHash).
	First(&token).Error; err != nil {
		return nil, domainErrors.ErrInvalidToken
	}

	return &token, nil
}
