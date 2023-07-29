package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/routes"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.SetupTagRoutes(app)
	routes.SetupEducationLevelRoutes(app)

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}
