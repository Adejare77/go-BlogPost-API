package auth

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/google/uuid"
)

type RefreshTokenRepository interface {
	Create(*entity.RefreshToken) error
	RevokeToken(tokenID uuid.UUID) error
	FindByTokenID(tokenID uuid.UUID) (*entity.RefreshToken, error)
}
