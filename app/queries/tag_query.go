package queries

import (
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

func FindTagByID(id string) *gorm.DB {
	return database.DB.Db.First(&models.Tag{}, id)
}

func FindTagByName(name string) *gorm.DB {
	return database.DB.Db.Where("name = ?", name).First(&models.Tag{})
}
