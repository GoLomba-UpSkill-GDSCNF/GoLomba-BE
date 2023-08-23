package seeders

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
)

func SeedRoles(c *fiber.Ctx) error {

	// if roles already seeded, return
	if database.DB.Db.First(&models.Role{}).Error == nil {
		return c.JSON(fiber.Map{
			"message": "roles already seeded",
		})
	}

	roles := []models.Role{
		{ID: 1, Name: "admin"},
		{ID: 2, Name: "user"},
	}

	database.DB.Db.Create(&roles)

	return c.JSON(fiber.Map{
		"message": "roles seeded",
	})

}
