package entity

import (
	"time"
)

type CommentID string

type Comment struct {
	ID CommentID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	AuthorID UserID `gorm:"not null;type:int"`
	Author User `gorm:"foreignKey:AuthorID; constraint:OnDelete:CASCADE"`
	PostID PostID `gorm:"not null;type:uuid"`
	Content string `gorm:"not null"`
	ParentID *CommentID `gorm:"type:uuid;default:null"`
	Like []Like `gorm:"polymorphic:Likeable;polymorphicValue:Comment"`
	Replies []Comment `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
