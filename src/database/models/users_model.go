package models

import (
	"time"
)

// User represents the users table
type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null" json:"username"`
	Email     string    `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	Password  string    `gorm:"type:text;not null" json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
