package auth

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	UserID entity.UserID `json:"user_id"`
}

type AuthResponse struct {
	ID entity.UserID `json:"id"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
