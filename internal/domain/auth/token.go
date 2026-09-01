package auth

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type TokenService interface {
	GenerateAccessToken(userID entity.UserID, isStaff bool) (string, error)
	GenerateRefreshToken() (string, error)
	Validate(tokenString string) (any, error)
}
