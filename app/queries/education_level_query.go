package queries

import (
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

func FindEducationLevelByID(id string) *gorm.DB {
	return database.DB.Db.First(&models.EducationLevel{}, id)
}

func FindEducationLevelByName(name string) *gorm.DB {
	return database.DB.Db.Where("name = ?", name).First(&models.EducationLevel{})
}
