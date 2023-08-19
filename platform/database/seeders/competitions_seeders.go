package seeders

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
)

func SeedCompetitions(c *fiber.Ctx) error {
	seed := c.Params("seed")

	seedInt, err := strconv.Atoi(seed)
	if err != nil {
		fmt.Println(err)
	}

	tags := []models.Tag{
		{ID: 1},
		{ID: 2},
	}
	educationLevels := []models.EducationLevel{
		{ID: 1},
		{ID: 2},
	}

	for i := 0; i < seedInt; i++ {
		competition := models.Competition{
			Name:                "competition" + strconv.Itoa(i),
			Description:         "description example" + strconv.Itoa(i),
			Image:               "image" + strconv.Itoa(i),
			Tags:                tags,
			EducationLevels:     educationLevels,
			CompetitionURL:      "competition_url" + strconv.Itoa(i),
			EndRegistrationDate: time.Now(),
			UserID:              4,
		}

		database.DB.Db.Create(&competition)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Competitions seeded successfully"})
}
