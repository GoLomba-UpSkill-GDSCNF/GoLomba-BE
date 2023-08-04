package utils

import (
	"github.com/gofiber/fiber/v2"
)

type response struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SendResponse(c *fiber.Ctx, status int, errorMessage string, data interface{}) error {
	response := response{Message: errorMessage, Data: data}
	return c.Status(status).JSON(response)
}