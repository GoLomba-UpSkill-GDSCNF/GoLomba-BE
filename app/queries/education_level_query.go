package queries

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/utils"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

func FindEducationLevelByID(id string) *gorm.DB {
	return database.DB.Db.First(&models.EducationLevel{}, id)
}

func FindEducationLevelByName(name string) *gorm.DB {
	return database.DB.Db.Where("name = ?", name).First(&models.EducationLevel{})
}

func FindEducationLevelsByIds(inputEducationLevels []models.EducationLevelInput) ([]models.EducationLevel, int, utils.ServerResponse) {
	var educationLevels []models.EducationLevel

	for _, educationLevelInput := range inputEducationLevels {

		// convert string to uint
		id, err := strconv.ParseUint(educationLevelInput.ID, 10, 32)
		if err != nil {
			return nil, fiber.StatusBadRequest, utils.InvalidData(err)
		}

		// check if there is id provided
		if id == 0 {
			return nil, fiber.StatusBadRequest, utils.CustomError("education_level_id is required")
		}
		var educationLevel models.EducationLevel

		// find single educationLevel by id
		if err := FindEducationLevelByID(educationLevelInput.ID).First(&educationLevel).Error; err != nil {
			// check if education level not found
			if gorm.ErrRecordNotFound.Error() == err.Error() {
				return nil, fiber.StatusNotFound, utils.IDNotFound("Education Level")
			}
			return nil, fiber.StatusInternalServerError, utils.ServerError(err)
		}
		educationLevels = append(educationLevels, educationLevel)
	}

	return educationLevels, 0, utils.ServerResponse{}
}
