package queries

import (
	"errors"
	"strconv"

	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

func SaveUser(user *models.User) error {
	return database.DB.Db.Save(user).Error
}

func GetUserByEmail(email string) models.User {
	var user models.User
	if err := database.DB.Db.Where("email = ?", email).Find(&user).Error; err != nil {
		return models.User{}
	}

	return user
}

func GetUserHashedPassword(email string) (string, error) {
	var user models.User
	if err := database.DB.Db.Table("users").Select("password").Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}

	return user.Password, nil
}

func GetUserById(id string) (uint, bool, error) {
	// convert string to uint
	userId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, false, err
	}

	var user models.User
	result := database.DB.Db.First(&user, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, false, nil
		} else {
			panic(result.Error)
		}
	}
	return uint(userId), true, nil // return true if user exists
}
