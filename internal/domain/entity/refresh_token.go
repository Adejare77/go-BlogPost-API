package entity

import (
	"time"
)

type RefreshTokenID int

type RefreshToken struct {
	ID RefreshTokenID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID UserID `gorm:"type:uuid"`
	TokenHash string `gorm:"not null;index"`
	RevokedAt *time.Time
	ExpiresAt time.Time
	CreatedAt time.Time
}
