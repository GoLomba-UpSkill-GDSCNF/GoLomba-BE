package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/queries"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

// Create tag handler
func CreateTag(c *fiber.Ctx) error {
	var tag models.Tag

	if err := c.BodyParser(&tag); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
		})
	}

	database.DB.Db.Create(&tag)

	return c.JSON(tag)
}

// Get all tags handler
func GetTags(c *fiber.Ctx) error {
	var tags []models.Tag
	database.DB.Db.Order("id asc").Find(&tags)
	return c.JSON(tags)
}

// Get tag by id handler
func GetTag(c *fiber.Ctx) error {
	id := c.Params("id")
	var tag models.Tag

	// check if param is valid
	if _, err := strconv.ParseUint(id, 10, 32); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse ID",
		})
	}

	// check if tag exists, if true then scan to tag
	if err := queries.FindTagByID(id).Scan(&tag).Error; err != nil {

		// check if tag not found
		if gorm.ErrRecordNotFound.Error() == err.Error() {
			return c.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "Tag not found",
			})
		}

		// if other error
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
		})
	}

	return c.JSON(tag)
}

// Update tag by id handler
func UpdateTag(c *fiber.Ctx) error {
	id := c.Params("id")

	// check if param is valid
	if _, err := strconv.ParseUint(id, 10, 32); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse ID",
		})
	}

	var tag models.Tag

	// check if tag exists, if true then scan to tag
	if err := queries.FindTagByID(id).Scan(&tag).Error; err != nil {

		// check if tag not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "Tag not found",
			})
		}

		// if other error
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
		})
	}

	// parse body to tag
	if err := c.BodyParser(&tag); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
		})
	}

	// update tag
	if err := database.DB.Db.Model(&tag).Updates(tag).Error; err != nil {

		// if other error
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
		})
	}

	return c.JSON(tag)
}

// Delete tag by id handler
func DeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")
	var tag models.Tag

	// check if param is valid
	if _, err := strconv.ParseUint(id, 10, 32); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse ID",
		})
	}

	// check if tag exists
	if err := queries.FindTagByID(id).Scan(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "Tag not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
		})
	}

	// delete tag
	if err := database.DB.Db.Delete(&tag).Error; err != nil {

		// if other error
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Tag deleted",
	})
}
