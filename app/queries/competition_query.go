package queries

import (
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

func FindCompetitionByID(id string) *gorm.DB {
	return database.DB.Db.First(&models.Competition{}, id)
}

func FindCompetitionByName(name string) *gorm.DB {
	return database.DB.Db.Where("name = ?", name).First(&models.Competition{})
}
