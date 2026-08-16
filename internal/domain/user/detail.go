package user

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type UserDetail struct {
	ID entity.UserID
	Email string
	FullName string
	CreatedAt time.Time
}
