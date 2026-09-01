package entity

import (
	"database/sql/driver"
	"encoding/json"
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
	Like []Like `gorm:"constraint:OnDelete:CASCADE;polymorphic:Likeable;polymorphicValue:post"`
	IsPublished bool `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}

func (id *PostID) Scan(value any) error {
	var u uuid.UUID

	if err := u.Scan(value); err != nil {
		return err
	}

	*id = PostID(u)

	return nil
}

func (id PostID) Value() (driver.Value, error) {
	return uuid.UUID(id).Value()
}

func (id PostID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(id).String())
}
