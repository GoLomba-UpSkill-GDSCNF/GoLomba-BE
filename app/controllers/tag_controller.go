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

// Create tag handler
func CreateTag(c *fiber.Ctx) error {
	var tag models.Tag

	if err := c.BodyParser(&tag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrParseJson())
	}

	// check if tag already exists, name must be unique
	existTag := queries.FindTagByName(tag.Name)
	if existTag.RowsAffected > 0 {
		return c.Status(fiber.StatusConflict).JSON(utils.DuplicateData("Tag"))
	}

	err := database.DB.Db.Create(&tag).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.SuccessCreated())
}

// Get all tags handler
func GetTags(c *fiber.Ctx) error {
	var tags []models.Tag

	// get all tags
	if err := database.DB.Db.Order("id asc").Find(&tags).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// check if tags is empty
	if len(tags) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.NotFound("Tags"))
	}

	return c.Status(fiber.StatusOK).JSON(tags)
}

// Get tag by id handler
func GetTag(c *fiber.Ctx) error {
	id := c.Params("id")

	// check if param is valid
	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam(id))
	}

	var tag models.Tag

	// check if tag exists, if true then assign to tag
	if err := queries.FindTagByID(id).Scan(&tag).Error; err != nil {

		// check if tag not found
		if gorm.ErrRecordNotFound.Error() == err.Error() {
			return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("Tag"))
		}

		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(tag)
}

// Update tag by id handler
func UpdateTag(c *fiber.Ctx) error {
	id := c.Params("id")

	// check if param is valid
	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam(id))
	}

	var tag models.Tag

	// check if tag exists, if true then assign to tag
	if err := queries.FindTagByID(id).Scan(&tag).Error; err != nil {

		// check if tag not found
		if gorm.ErrRecordNotFound.Error() == err.Error() {
			return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("Tag"))
		}

		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// parse body to tag
	if err := c.BodyParser(&tag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrParseJson())
	}

	// check if tag already exists, name must be unique
	existTag := queries.FindTagByName(tag.Name)
	if existTag.RowsAffected > 0 {
		return c.Status(fiber.StatusConflict).JSON(utils.DuplicateData("Tag"))
	}

	// update tag
	if err := database.DB.Db.Model(&tag).Updates(tag).Error; err != nil {

		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessUpdated())
}

// Delete tag by id handler
func DeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")

	// check if param is valid
	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam(id))
	}

	var tag models.Tag

	// check if tag exists
	if err := queries.FindTagByID(id).Scan(&tag).Error; err != nil {

		if gorm.ErrRecordNotFound.Error() == err.Error() {
			return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("Tag"))
		}

		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// delete tag
	if err := database.DB.Db.Delete(&tag).Error; err != nil {

		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessDeleted())
}
