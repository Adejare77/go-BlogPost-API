package entity

import "time"

type UserID int

type User struct {
	ID UserID `gorm:"primaryKey;"`
	Email string `gorm:"not null;uniqueIndex"`
	FullName string `gorm:"not null"`
	Password *string
	IsActive bool `gorm:"not null;default:true"`
	IsStaff bool `gorm:"not null;default:false"`
	CreatedAt time.Time
}
