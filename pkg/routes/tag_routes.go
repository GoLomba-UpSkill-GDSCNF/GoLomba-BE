package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/controllers"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
)

func SetupTagRoutes(app *fiber.App) {
	tag := app.Group("/tag")

	app.Get("/tags", controllers.GetTags)
	tag.Get("/:id", controllers.GetTag)

	tag.Use(middleware.JWTMiddleware(), middleware.IsAdmin())
	tag.Post("/", middleware.ValidateField[models.Tag](), controllers.CreateTag)
	tag.Put("/:id", middleware.ValidateField[models.Tag](), controllers.UpdateTag)
	tag.Delete("/:id", controllers.DeleteTag)
}
