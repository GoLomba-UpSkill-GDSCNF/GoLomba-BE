package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/controllers"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
)

func New() *fiber.App {
	userController := controllers.UserController{}
	app := fiber.New()
	app.Use(logger.New(logger.Config{}))

	appUser := app.Group("/user")
	appUser.Get("/register", middleware.ValidateUser, userController.Register)
	appUser.Get("/login", middleware.ValidateUser, userController.Login)

	return app
}