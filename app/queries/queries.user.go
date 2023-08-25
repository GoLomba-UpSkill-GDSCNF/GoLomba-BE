package queries

import (
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
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

func GetUserById(id uint) (models.User, error) {
	var user models.User
	if err := database.DB.Db.Where("id = ?", id).Find(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetRoleById(id uint) (models.Role, error) {
	var role models.Role
	if err := database.DB.Db.Where("id = ?", id).Find(&role).Error; err != nil {
		return models.Role{}, err
	}

	return role, nil
}
