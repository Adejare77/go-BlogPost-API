package auth

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	UserID entity.UserID `json:"user_id"`
	Email string `json:"email"`
}
