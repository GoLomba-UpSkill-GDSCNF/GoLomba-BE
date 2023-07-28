package utils

import (
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
)

func EmailChecker(email string) bool {
	var user models.User
	database.DB.Table("users").Where("email = ?", email).Find(&user)
	
	if user.ID != 0 {
		return true
	}

	return false
}