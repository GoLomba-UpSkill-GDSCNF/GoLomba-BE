package models

import (
	"time"

	"gorm.io/gorm"
)

type Competition struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" gorm:"type:text" validate:"required,min=3,max=510" `
	Image       string `json:"image" gorm:"type:varchar(255)" validate:"required,min=3,max=255" `

	// Many to Many
	Tags           []Tag            `json:"tags" gorm:"many2many:competition_tags;" validate:"required,min=1,max=255"`
	EducationLevel []EducationLevel `json:"education_levels" gorm:"many2many:competition_education_levels;" validate:"required,min=1,max=255"`

	// Foreign Keys
	UserID uint `json:"user_id" gorm:"not null" validate:"required"`
	User   User `json:"user" gorm:"foreignKey:UserID" validate:"required"`

	EndRegistrationDate time.Time `json:"end_registration_date" validate:"required,min=3,max=255"`
	CompetitionURL      string    `json:"competition_url" gorm:"type:varchar(255)" validate:"required,min=3,max=255"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

type CompetitionInput struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" gorm:"type:text" validate:"required,min=3,max=510" `
	Image       string `json:"image" gorm:"type:varchar(255)" validate:"required,min=3,max=255" `

	// Many to Many
	Tags            []TagInput            `json:"tags" gorm:"many2many:competition_tags;" validate:"required,min=1,max=255"`
	EducationLevels []EducationLevelInput `json:"education_levels" gorm:"many2many:competition_education_levels;" validate:"required,min=1,max=255"`

	// Foreign Keys
	UserID uint `json:"user_id" gorm:"not null" validate:"required"`

	EndRegistrationDate string `json:"end_registration_date" validate:"required,min=3,max=255"`
	CompetitionURL      string `json:"competition_url" gorm:"type:varchar(255)" validate:"required,min=3,max=255"`
}
