package like

import "github.com/Adejare77/go-BlogPost-API/internal/domain/user"

type LikeID string

type Like struct{
	ID LikeID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID user.UserID `gorm:"not null;uniqueIndex:idx_user_like"`
	LikeableID string `gorm:"not null;uniqueIndex:idx_user_like"`
	LikeableType string `gorm:"not null;uniqueIndex:idx_user_like"`
}
