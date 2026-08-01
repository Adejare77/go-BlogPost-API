package entity

import (
	"time"

	"github.com/google/uuid"
)

type PostID uuid.UUID

type Post struct {
	ID PostID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	AuthorID UserID `gorm:"not null;index"`
	Author User `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Title string `gorm:"not null;size:100"`
	Content string `gorm:"not null"`
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
	Like []Like `gorm:"constraint:OnDelete:CASCADE;polymorphic:Likeable;polymorphicValue:Post"`
	IsPublished bool `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}
