package auth

import "github.com/Adejare77/go-BlogPost-API/internal/domain/entity"

type AuthResult struct {
	AccessToken string
	RefreshToken string
	UserID entity.UserID
}
