package queries

import (
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
)

func SaveUser(user *models.User) error {
	return database.DB.Save(user).Error
}

func GetUserByEmail(email string) (models.User) {
	var user models.User
	if err := database.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return models.User{}
	}

	return user
}

func GetUserHashedPassword(email string) (string, error) {
	var user models.User
	if err := database.DB.Table("users").Select("password").Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}
	
	return user.Password, nil
}

func CheckUser(id int) bool {
	var user models.User
	database.DB.Table("users").Find(&user, id)
	if user.ID == 0 {
		return false
	}

	return true
}