package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/routes"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database/seeders"
)

func main() {
	database.Connect()

	// seeders Roles
	rolesSeeder := seeders.SeedRoles()
	log.Println(rolesSeeder)

	app := fiber.New()

	routes.SetupSeedersRoutes(app)
	routes.SetupAuthRoutes(app)
	routes.SetupTagRoutes(app)
	routes.SetupEducationLevelRoutes(app)
	routes.SetupCompetitionRoutes(app)

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}
