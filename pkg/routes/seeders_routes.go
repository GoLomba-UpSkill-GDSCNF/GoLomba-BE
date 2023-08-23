package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database/seeders"
)

func SetupSeedersRoutes(app *fiber.App) {
	seeds := app.Group("/seeders")

	seeds.Use(middleware.JWTMiddleware(), middleware.IsAdmin())
	seeds.Get("/competitions/:seed", seeders.SeedCompetitions)
}
