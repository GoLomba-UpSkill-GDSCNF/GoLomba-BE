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

// find all educationLevels by competition_id in competition_educationLevels table, return educationLevels name
func CompeFindEducationLevelsByNames(Competition models.Competition) ([]models.EducationLevelResponse, error) {
	var educationLevels []models.EducationLevel
	var educationLevelsResponse []models.EducationLevelResponse

	err := database.DB.Db.Select("education_levels.name").Joins("JOIN competition_education_levels ON competition_education_levels.education_level_id = education_levels.id").Where("competition_education_levels.competition_id = ?", Competition.ID).Find(&educationLevels).Scan(&educationLevelsResponse).Error
	if err != nil {
		return []models.EducationLevelResponse{}, err
	}
	return educationLevelsResponse, nil
}

// find all educationLevels by competition_id in competition_educationLevels table amd delete them
func FindCompeDeleteEduLevelsByCompetitionID(id string) error {

	// check if competition_id is valid
	if _, err := strconv.ParseUint(id, 10, 32); err != nil {
		return err
	}

	// check if competition_id exist in competition_education_levels table
	if err := database.DB.Db.Table("competition_education_levels").Where("competition_id = ?", id).Error; err != nil {
		return err
	}

	// delete all rows with competition_education_levels table by competition_id
	if err := database.DB.Db.Table("competition_education_levels").Where("competition_id = ?", id).Delete(&models.EducationLevel{}).Error; err != nil {
		return err
	}
	return nil

}

// Update education levels from competition_education_levels table, by first delete all existing education levels associated with the competition, then add new education levels to the competition_education_levels table
func UpdateCompeEduLevelsById(id string, eduLevels []models.EducationLevel) error {
	// delete all existing education levels associated with the competition
	if err := FindCompeDeleteEduLevelsByCompetitionID(id); err != nil {
		return err
	}

	// add new education levels to the competition
	competition := &models.Competition{}
	if err := database.DB.Db.First(&competition, id).Error; err != nil {
		return err
	}
	if err := database.DB.Db.Model(&competition).Association("EducationLevels").Append(eduLevels); err != nil {
		return err
	}
	return nil
}
