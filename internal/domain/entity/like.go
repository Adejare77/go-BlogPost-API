package entity

import "github.com/google/uuid"

type LikeID uuid.UUID

type Like struct{
	ID LikeID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID UserID `gorm:"not null;uniqueIndex:idx_user_like"`
	LikeableID LikeID `gorm:"not null;uniqueIndex:idx_user_like"`
	LikeableType string `gorm:"not null;uniqueIndex:idx_user_like"`
}
