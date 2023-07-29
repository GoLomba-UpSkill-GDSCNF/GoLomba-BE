package models

import (
	"time"

	"gorm.io/gorm"
)

type Testimonial struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Stars   uint   `json:"stars"`
	Message string `json:"message"`

	// Foreign Keys
	UserID uint `json:"user_id"`
	User   User `json:"user" gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
