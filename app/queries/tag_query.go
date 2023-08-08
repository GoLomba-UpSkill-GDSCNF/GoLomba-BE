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

		// check if there is id provided
		if tagInput.ID == 0 {
			return nil, fiber.StatusBadRequest, utils.CustomError("tag_id is required")
		}
		var tag models.Tag

		// convert id to string
		id := strconv.Itoa(int(tagInput.ID))

		// find single tag by id
		if err := FindTagByID(id).First(&tag).Error; err != nil {

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
