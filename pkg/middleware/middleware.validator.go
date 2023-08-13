package middleware

// type ErrorResponse struct {
// 	FailedField string
// 	Tag         string
// 	Value       string
// }

// func ValidateData(data interface{}) []*ErrorResponse {
// 	validate := validator.New()
// 	var errors []*ErrorResponse
// 	err := validate.Struct(data)

// 	if err != nil {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			var element ErrorResponse
// 			element.FailedField = err.StructNamespace()
// 			element.Tag = err.Tag()
// 			element.Value = err.Param()
// 			errors = append(errors, &element)
// 		}
// 	}

// 	return errors
// }

// func ValidateUser(c *fiber.Ctx) error {
// 	user := new(models.User)
// 	if err := c.BodyParser(user); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})

// 	}

// 	err := ValidateData(*user)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(err)
// 	}
// 	return c.Next()
// }
