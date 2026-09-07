package auth

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type AuthResponse struct {
	ID entity.UserID
	FullName string
	Email string
	CreatedAt time.Time
}
