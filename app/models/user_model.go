package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	Email     string `json:"email" gorm:"unique"`
	RoleID    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleID"`
}
