package auth

import "github.com/Adejare77/go-BlogPost-API/internal/domain/entity"

type RefreshTokenRepository interface {
	Create(*entity.RefreshToken) error
	RevokeToken(token string) error
	FindByTokenHash(tokenHash string) (*entity.RefreshToken, error)
}
