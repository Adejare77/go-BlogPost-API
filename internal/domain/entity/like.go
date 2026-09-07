package entity

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type LikeID uuid.UUID

type Like struct{
	ID LikeID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID UserID `gorm:"not null;uniqueIndex:idx_user_like"`
	LikeableID LikeID `gorm:"not null;uniqueIndex:idx_user_like"`
	LikeableType string `gorm:"not null;uniqueIndex:idx_user_like"`
}


func (id *LikeID) Scan(value any) error{
	var v uuid.UUID

	if err := v.Scan(value); err != nil {
		return err
	}

	*id = LikeID(v)

	return nil
}

func (id  LikeID) Value() (driver.Value, error) {
	return uuid.UUID(id).Value()
}
