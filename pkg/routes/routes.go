package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/controllers"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
)

func SetupAuthRoutes(app *fiber.App) {
	userController := controllers.UserController{}
	testimonyController := controllers.TestimonyControllers{}
	app.Use(logger.New(logger.Config{}))

	appUser := app.Group("/user")
	appUser.Post("/register", middleware.ValidateField[models.User](), userController.Register)
	appUser.Post("/login", middleware.ValidateField[models.User](), userController.Login)
	appUser.Post("/testimonial", testimonyController.Add)

	app.Get("/testimonial", testimonyController.GetAll)
}
