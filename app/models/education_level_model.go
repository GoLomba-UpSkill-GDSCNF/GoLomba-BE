package models

type EducationLevel struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique" validate:"required,min=2,max=20"`
}
