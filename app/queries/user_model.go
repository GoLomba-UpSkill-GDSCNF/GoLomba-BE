package queries

import (
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

func UserExistById(id uint) bool {
	var user models.User
	result := database.DB.Db.First(&user, id)
	return !(gorm.ErrRecordNotFound.Error() == result.Error.Error())
}
