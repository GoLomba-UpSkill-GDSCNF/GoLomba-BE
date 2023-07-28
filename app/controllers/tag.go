package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
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
	database.DB.Db.First(&tag, id)

	if tag.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Tag not found",
		})
	}

	return c.JSON(tag)
}

// Update tag by id handler
func UpdateTag(c *fiber.Ctx) error {
	id := c.Params("id")

	// convert string to uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse ID",
		})
	}

	tag := models.Tag{
		ID: uint(idUint),
	}

	// find tag by ID
	result := database.DB.Db.First(&tag, tag.ID)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Tag not found",
		})
	}

	if err := c.BodyParser(&tag); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
		})
	}

	database.DB.Db.Model(&tag).Updates(tag)
	return c.JSON(tag)
}

// Delete tag by id handler
func DeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")
	var tag models.Tag
	database.DB.Db.First(&tag, id)
	if tag.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Tag not found",
		})
	}
	database.DB.Db.Delete(&tag)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Tag deleted",
	})
}
