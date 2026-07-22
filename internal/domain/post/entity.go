package post

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/comment"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/like"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/user"
)

type PostID string

type Post struct {
	ID PostID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	AuthorID user.UserID `gorm:"not null;index"`
	Author user.User `gorm:"foreignKey:AuthorID;constraint:onDelete:CASCADE"`
	Title string `gorm:"not null;size:100"`
	Content string `gorm:"not null"`
	Comments []comment.Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
	Like like.Like `gorm:"constraint:OnDelete:CASCADE;polymorphic:Likeable;polymorphicValue:Post"`
	IsPublished bool `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}
