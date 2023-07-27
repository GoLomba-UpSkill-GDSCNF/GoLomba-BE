package models

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique"`
}

type Competition struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Tags        []Tag  `json:"tags" gorm:"many2many:competition_tags;"`

	// Foreign Keys
	EducationLevelID uint           `json:"education_level_id"`
	EducationLevel   EducationLevel `json:"education_level" gorm:"foreignKey:EducationLevelID"`
	UserID           uint           `json:"user_id"`
	User             User           `json:"user" gorm:"foreignKey:UserID"`

	EndRegistrationDate string `json:"end_registration_date"`
	CompetitionURL      string `json:"competition_url"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
