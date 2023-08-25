package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/controllers"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
)

func SetupCompetitionRoutes(app *fiber.App) {
	competition := app.Group("/competition")

	app.Get("/competitions", controllers.GetCompetitions)
	competition.Get("/:id", controllers.GetCompetition)

	competition.Use(middleware.JWTMiddleware())

	competition.Post("/", middleware.ValidateField[models.CompetitionInput](), controllers.CreateCompetition)
	competition.Put("/:id", middleware.ValidateField[models.CompetitionInput](), controllers.UpdateCompetition)
	competition.Delete("/:id", controllers.DeleteCompetition)

	app.Get("/user/competitions", middleware.JWTMiddleware(), controllers.UserCompetitions)
}
