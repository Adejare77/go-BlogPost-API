package user

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type UserResponse struct {
	ID entity.UserID `json:"id"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
