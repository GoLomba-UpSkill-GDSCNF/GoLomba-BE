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

// find all tags by competition_id in competition_tags table, return tags name
func CompeFindTagsByNames(Competition models.Competition) ([]models.TagResponse, error) {
	var tags []models.Tag
	var tagsResponse []models.TagResponse

	err := database.DB.Db.Select("tags.name").Joins("JOIN competition_tags ON competition_tags.tag_id = tags.id").Where("competition_tags.competition_id = ?", Competition.ID).Find(&tags).Scan(&tagsResponse).Error
	if err != nil {
		return []models.TagResponse{}, err
	}

	return tagsResponse, nil
}

// find all competition_tags by competition_id in competition_tags table and delete all of them
func FindCompeDeleteTagsById(id string) error {

	// check if competition_id is valid
	if _, err := strconv.ParseUint(id, 10, 32); err != nil {
		return err
	}

	// check if competition_id is exist in competition_tags table
	if err := database.DB.Db.Table("competition_tags").Where("competition_id = ?", id).Error; err != nil {
		return err
	}

	// delete all rows in competition_tags table by competition_id
	if err := database.DB.Db.Table("competition_tags").Where("competition_id = ?", id).Delete(&models.Tag{}).Error; err != nil {
		return err
	}
	return nil

}

// Update tags from competition_tags table, by first delete all existing tags associated with the competition, then add new tags to the competition_tags table
func UpdateCompeTagsById(id string, tags []models.Tag) error {
	// delete all existing tags associated with the competition
	if err := FindCompeDeleteTagsById(id); err != nil {
		return err
	}

	// add new tags to the competition
	competition := &models.Competition{}
	if err := database.DB.Db.First(&competition, id).Error; err != nil {
		return err
	}
	if err := database.DB.Db.Model(&competition).Association("Tags").Append(tags); err != nil {
		return err
	}
	return nil
}
