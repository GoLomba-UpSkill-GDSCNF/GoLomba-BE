package models

import (
	"time"

	"gorm.io/gorm"
)

type Competition struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
	Image       string `json:"image"`

	// Many to Many
	Tags           []Tag            `json:"tags" gorm:"many2many:competition_tags;"`
	EducationLevel []EducationLevel `json:"education_levels" gorm:"many2many:competition_education_levels;"`

	// Foreign Keys
	UserID uint `json:"user_id"`
	User   User `json:"user" gorm:"foreignKey:UserID"`

	EndRegistrationDate string `json:"end_registration_date"`
	CompetitionURL      string `json:"competition_url"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
