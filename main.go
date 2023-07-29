package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/controllers"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
)

func SetupTagRoutes(app *fiber.App) {
	app.Post("/tag", middleware.ValidateField[models.Tag](), controllers.CreateTag)
	app.Get("/tags", controllers.GetTags)
	app.Get("/tag/:id", controllers.GetTag)
	app.Put("/tag/:id", middleware.ValidateField[models.Tag](), controllers.UpdateTag)
	app.Delete("/tag/:id", controllers.DeleteTag)
}

func main() {
	database.Connect()

	app := fiber.New()

	SetupTagRoutes(app)

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}
