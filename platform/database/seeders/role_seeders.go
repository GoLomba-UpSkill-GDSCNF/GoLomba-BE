package seeders

import (
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
)

func SeedRoles() string {

	// if roles already seeded, return
	var roles []models.Role
	database.DB.Db.Find(&roles)
	if len(roles) > 0 {
		return "Roles already seeded"
	}

	seedRoles := []models.Role{
		{ID: 1, Name: "admin"},
		{ID: 2, Name: "user"},
	}

	database.DB.Db.Create(&seedRoles)

	return "Roles seeded successfully"

}
