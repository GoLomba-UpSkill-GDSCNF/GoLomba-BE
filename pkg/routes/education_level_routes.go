package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/controllers"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
)

func SetupEducationLevelRoutes(app *fiber.App) {
	eduLevel := app.Group("/edu-level")

	app.Get("/edu-levels", controllers.GetEducationLevels)
	eduLevel.Get("/:id", controllers.GetEducationLevel)

	eduLevel.Use(middleware.JWTMiddleware(), middleware.IsAdmin())
	eduLevel.Post("/", middleware.ValidateField[models.EducationLevel](), controllers.CreateEducationLevel)
	eduLevel.Put("/:id", middleware.ValidateField[models.EducationLevel](), controllers.UpdateEducationLevel)
	eduLevel.Delete("/:id", controllers.DeleteEducationLevel)
}
