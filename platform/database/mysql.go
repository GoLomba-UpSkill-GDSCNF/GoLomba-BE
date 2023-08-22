package database

import (
	"fmt"
	"log"
	"os"

	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Instance struct {
	Db *gorm.DB
}

var DB Instance

func Connect() {
	var DBConfig = configs.MySQLConfig

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.DBName)
	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Database connected!")
	dbConnection.Logger = logger.Default.LogMode(logger.Info)

	// Add migrations
	log.Println("Running migrations...")
	dbConnection.AutoMigrate(&models.User{}, &models.Role{}, &models.Tag{}, &models.Competition{}, &models.EducationLevel{}, &models.Testimonial{}, &models.User{}, &models.Testimonial{})

	DB = Instance{Db: dbConnection}
}
