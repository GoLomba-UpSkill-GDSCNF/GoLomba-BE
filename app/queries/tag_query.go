package queries

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/utils"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

func FindTagByID(id string) *gorm.DB {
	return database.DB.Db.First(&models.Tag{}, id)
}

func FindTagByName(name string) *gorm.DB {
	return database.DB.Db.Where("name = ?", name).First(&models.Tag{})
}

func FindTagsByIds(inputTags []models.TagInput) ([]models.Tag, int, utils.ServerResponse) {
	var tags []models.Tag

	for _, tagInput := range inputTags {

		// convert string to uint
		id, err := strconv.ParseUint(tagInput.ID, 10, 32)
		if err != nil {
			return nil, fiber.StatusBadRequest, utils.InvalidData(err)
		}

		// check if there is id provided
		if id == 0 {
			return nil, fiber.StatusBadRequest, utils.CustomError("tag_id is required")
		}
		var tag models.Tag

		// find single tag by id
		if err := FindTagByID(tagInput.ID).First(&tag).Error; err != nil {

			// if tag not found
			if gorm.ErrRecordNotFound == err {
				return nil, fiber.StatusNotFound, utils.IDNotFound("Tag")
			}
			return nil, fiber.StatusInternalServerError, utils.ServerError(err)
		}
		tags = append(tags, tag)
	}

	return tags, 0, utils.ServerResponse{}
}
