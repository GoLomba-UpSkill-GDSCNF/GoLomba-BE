package models

type Tag struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex" validate:"required,min=2,max=20"`
}

type TagInput struct {
	ID uint `json:"id" validate:"required"`
}
