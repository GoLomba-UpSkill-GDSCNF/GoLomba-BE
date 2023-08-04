package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/controllers"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
)

func New() *fiber.App {
	userController := controllers.UserController{}
	testimonyController := controllers.TestimonyControllers{}
	app := fiber.New()
	app.Use(logger.New(logger.Config{}))

	appUser := app.Group("/user")
	appUser.Post("/register", middleware.ValidateUser, userController.Register)
	appUser.Post("/login", middleware.ValidateUser, userController.Login)
	appUser.Post("/testimonial", testimonyController.Add)
	
	app.Get("/testimonial", testimonyController.GetAll)

	return app
}
