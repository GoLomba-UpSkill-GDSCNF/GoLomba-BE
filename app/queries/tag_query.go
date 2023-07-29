package queries

import (
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

func FindTagByID(id string) *gorm.DB {
	var tag models.Tag

	return database.DB.Db.First(&tag, id)
}
