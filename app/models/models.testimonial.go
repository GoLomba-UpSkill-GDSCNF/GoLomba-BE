package models

import "gorm.io/gorm"

type Testimonial struct {
	gorm.Model
	UserID uint `json:"user_id" gorm:"type:bigint(20)"`
	Message string `json:"message" gorm:"type:text"`
	Stars uint `json:"stars" gorm:"type:tinyint(5)"`
	User User `gorm:"foreignKey:UserID"`
}

type TestimonialResponse struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	Message string `json:"message"`
	Stars uint `json:"stars"`
	FullName string `json:"user_full_name"`
	
}