package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Instance struct {
	Db *gorm.DB
}

var DB Instance

func Connect() {
	dsn := "root:root@tcp(localhost:3306)/golomba_db?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Database connected!")
	dbConnection.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations...")
	// TODO: Add migrations

	DB = Instance{Db: dbConnection}
}
