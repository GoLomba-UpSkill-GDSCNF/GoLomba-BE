package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/queries"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/utils"
)

type UserController struct{}
var err error

func (u *UserController) Register(c *fiber.Ctx) error {
	var user models.User

	if err = c.BodyParser(&user); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if isEmailRegistered := utils.EmailChecker(user.Email); isEmailRegistered {
		return utils.SendResponse(c, fiber.StatusBadRequest, "email already registered", nil)
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return utils.SendResponse(c, fiber.StatusInternalServerError, "registration failed", nil)
	}

	user.Password = hashedPassword
	user.RoleID = 2 // peserta

	if err = queries.SaveUser(&user); err != nil {
		log.Println(err)
		return utils.SendResponse(c, fiber.StatusInternalServerError, "registration failed", nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "success create user", nil)
}

func (u *UserController) Login(c *fiber.Ctx) error {
	var user models.User

	if err = c.BodyParser(&user); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if isEmailRegistered := utils.EmailChecker(user.Email); !isEmailRegistered {
		return utils.SendResponse(c, fiber.StatusBadRequest, "email or password wrong", nil)
	}


	hashedPassword, err := queries.GetUserHashedPassword(user.Email)
	if err != nil {
		log.Println(err)
		return utils.SendResponse(c, fiber.StatusInternalServerError, "login failed", nil)
	}

	if err = utils.ComparePassword(user.Password, hashedPassword); err != nil {
		log.Println(err.Error())
		return utils.SendResponse(c, fiber.StatusBadRequest, "email or password wrong", nil)
	}

	dataUser := queries.GetUserByEmail(user.Email)
	token ,err := middleware.CreateToken(dataUser.ID, dataUser.RoleID)
	if err != nil {
		log.Println(err)
		return utils.SendResponse(c, fiber.StatusInternalServerError, "failed login", nil)
	}
	
	data := map[string]string{
		"token": token,
	}

	return utils.SendResponse(c, fiber.StatusOK, "success login", data)
}