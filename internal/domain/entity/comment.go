package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type CommentID uuid.UUID

type Comment struct {
	ID CommentID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	AuthorID UserID `gorm:"not null;type:int"`
	Author User `gorm:"foreignKey:AuthorID; constraint:OnDelete:CASCADE"`
	PostID PostID `gorm:"not null;type:uuid"`
	Content string `gorm:"not null"`
	ParentID *CommentID `gorm:"type:uuid;default:null"`
	Like []Like `gorm:"polymorphic:Likeable;polymorphicValue:comment"`
	Replies []Comment `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


func (id *CommentID) Scan(value any) error {
	var u uuid.UUID

	if err := u.Scan(value); err != nil {
		return err
	}

	*id = CommentID(u)

	return nil
}

func (id CommentID) Value() (driver.Value, error) {
	return uuid.UUID(id).Value()
}

func (id CommentID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(id).String())
}
