package database

import (
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func InitDB() {
	dsn := "egi:@tcp(127.0.0.1:3306)/gdsc?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(models.User{}, models.Role{}, models.Testimonial{})
}