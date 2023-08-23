package seeders

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/utils"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
)

func SeedCompetitions(c *fiber.Ctx) error {
	seed := c.Params("seed")

	seedInt, err := strconv.Atoi(seed)
	if err != nil {
		fmt.Println(err)
	}

	tokenJWT := ""

	if authHeader := c.Request().Header.Peek("Authorization"); len(authHeader) > 0 {
		tokenJWT = strings.Fields(string(authHeader))[1]
	}

	userID, _, err := middleware.CheckTokenValue(tokenJWT)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ServerError(err))
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
			UserID:              userID.(uint),
		}

		database.DB.Db.Create(&competition)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Competitions seeded successfully"})
}
