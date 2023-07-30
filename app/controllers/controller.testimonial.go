package controllers

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/queries"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/utils"
)

type TestimonyControllers struct{}

func (t *TestimonyControllers) Add(c *fiber.Ctx) error {
	var testimony models.Testimonial
	tokenJWT := strings.Fields(string(c.Request().Header.Peek("Authorization")))[1]

	if err = c.BodyParser(&testimony); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "invalid input data", nil)
	}
	
	userID, _, err :=  middleware.CheckTokenValue(tokenJWT)
	if err != nil {
		log.Println(err)
		return utils.SendResponse(c, fiber.StatusBadRequest, "invalid or expired token", nil)
	}
	
	
	if isUserExist := queries.CheckUser(int(userID.(float64))); !isUserExist {
		return utils.SendResponse(c, fiber.StatusNotFound, "user with the provided user_id does not exist.", nil)
	}
	
	testimony.UserID = uint(userID.(float64))
	if err = queries.SaveTestimony(&testimony); err != nil {
		log.Println(err)
		return utils.SendResponse(c, fiber.StatusInternalServerError, "failed add testimony", nil)
	}

	return utils.SendResponse(c, fiber.StatusCreated, "success add testimony", nil)
}

func (t *TestimonyControllers) GetAll(c *fiber.Ctx) error {
	testimonials, err := queries.GetAllTestimony()
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "failed get testimonials",nil)
	}
	
	return utils.SendResponse(c, fiber.StatusOK, "success get testimonials", testimonials)
}