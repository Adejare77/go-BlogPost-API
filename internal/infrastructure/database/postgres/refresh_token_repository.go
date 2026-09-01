package postgres

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	domainErrors "github.com/Adejare77/go-BlogPost-API/internal/domain/errors"
	"github.com/google/uuid"
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

func (repo *RefreshTokenRepository) RevokeToken(tokenID uuid.UUID) error {
	result := repo.db.Model(&entity.RefreshToken{}).
	Where("id = ? AND revoked_at IS NULL", tokenID).
	Update("revoked_at", time.Now())

	if result.RowsAffected == 0 {
		return domainErrors.ErrTokenRevoked
	}

	return nil
}

func (repo *RefreshTokenRepository) FindByTokenID(tokenID uuid.UUID) (*entity.RefreshToken, error) {
	var tokenHash entity.RefreshToken

	if err := repo.db.
	Where("id = ?", tokenID).
	First(&tokenHash).Error; err != nil {
		return nil, domainErrors.ErrInvalidToken
	}

	return &tokenHash, nil
}
