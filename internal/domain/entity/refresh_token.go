package entity

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID UserID
	IsStaff bool
	TokenHash string `gorm:"not null;index"`
	RevokedAt *time.Time
	ExpiresAt time.Time
	CreatedAt time.Time
}
