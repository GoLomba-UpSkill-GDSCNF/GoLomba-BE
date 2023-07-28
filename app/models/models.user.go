package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	FullName string `json:"username" gorm:"type:varchar(255)"`
	Email string `json:"email" gorm:"type:varchar(255);unique" validate:"required,email"`
	Password string `json:"password" gorm:"type:text" validate:"required"`
	RoleID uint `gorm:"type:int"`
	Role Role `json:"role" gorm:"foreignKey:RoleID"`
}

