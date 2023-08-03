package middleware

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
)

type Input interface {
	models.Tag | models.EducationLevel | models.CompetitionInput
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// to validate struct using validator package
func ValidateStruct[T Input](data T) []*ErrorResponse {
	validate := validator.New()
	var errors []*ErrorResponse
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

// to validate field from request body
func ValidateField[T Input]() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		data := new(T)

		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		errors := ValidateStruct(*data)
		if errors != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": errors,
			})
		}

		// if no error, continue
		return c.Next()
	}
}
