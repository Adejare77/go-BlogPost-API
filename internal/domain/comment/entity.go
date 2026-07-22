package comment

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/like"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/post"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/user"
)


type CommentID string

type Comment struct {
	ID CommentID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	AuthorID user.User `gorm:"foreignKey;not null;type:int;constraint:onDelete:CASCADE"`
	PostID post.PostID `gorm:"not null;type:uuid;constraint:onDelete:CASCADE"`
	Content string `gorm:"not null"`
	Like []like.Like `gorm:"polymorphic:Likeable;polymorphicValue:Comment"`
	ParentID *string `gorm:"type:uuid;default:null"`
	Replies []Comment `gorm:"foreignKey:ParentID;constraint:onDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
