package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/queries"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/utils"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

// Create educationLevel handler
func CreateEducationLevel(c *fiber.Ctx) error {
	var educationLevel models.EducationLevel

	if err := c.BodyParser(&educationLevel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrParseJson())
	}

	// check if educationLevel already exists, name must be unique
	existEducationLevel := queries.FindEducationLevelByName(educationLevel.Name)
	if existEducationLevel.RowsAffected > 0 {
		return c.Status(fiber.StatusConflict).JSON(utils.DuplicateData("EducationLevel"))
	}

	err := database.DB.Db.Create(&educationLevel).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.SuccessCreated())
}

// Get all educationLevels handler
func GetEducationLevels(c *fiber.Ctx) error {
	var educationLevels []models.EducationLevel

	// get all educationLevels
	if err := database.DB.Db.Order("id asc").Find(&educationLevels).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// check if educationLevels is empty
	if len(educationLevels) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.NotFound("EducationLevels"))
	}

	return c.Status(fiber.StatusOK).JSON(educationLevels)
}

// Get educationLevel by id handler
func GetEducationLevel(c *fiber.Ctx) error {
	id := c.Params("id")

	// check if param is valid
	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("id"))
	}

	var educationLevel models.EducationLevel

	// check if tag exists, if true then assign to tag
	if err := queries.FindEducationLevelByID(id).Scan(&educationLevel).Error; err != nil {

		// check if educationLevel not found
		if gorm.ErrRecordNotFound.Error() == err.Error() {
			return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("EducationLevel"))
		}

		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(educationLevel)
}

// Update educationLevel by id handler
func UpdateEducationLevel(c *fiber.Ctx) error {
	id := c.Params("id")

	// check if param is valid
	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("id"))
	}

	var educationLevel models.EducationLevel

	// check if educationLevel exists, if true then assign to educationLevel
	if err := queries.FindEducationLevelByID(id).Scan(&educationLevel).Error; err != nil {

		// check if educationLevel not found
		if gorm.ErrRecordNotFound.Error() == err.Error() {
			return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("EducationLevel"))
		}

		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// parse body to educationLevel
	if err := c.BodyParser(&educationLevel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrParseJson())
	}

	// check if educationLevel already exists, name must be unique
	existEducationLevel := queries.FindEducationLevelByName(educationLevel.Name)
	if existEducationLevel.RowsAffected > 0 {
		return c.Status(fiber.StatusConflict).JSON(utils.DuplicateData("EducationLevel"))
	}

	// update educationLevel
	if err := database.DB.Db.Save(&educationLevel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessUpdated())
}

func DeleteEducationLevel(c *fiber.Ctx) error {
	id := c.Params("id")

	// check if param is valid
	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("id"))
	}

	var educationLevel models.EducationLevel

	// check if educationLevel exists, if true then assign to educationLevel
	if err := queries.FindEducationLevelByID(id).Scan(&educationLevel).Error; err != nil {

		// check if educationLevel not found
		if gorm.ErrRecordNotFound.Error() == err.Error() {
			return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("EducationLevel"))
		}

		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// delete educationLevel
	if err := database.DB.Db.Delete(&educationLevel).Error; err != nil {

		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessDeleted())
}
