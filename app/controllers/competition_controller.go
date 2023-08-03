package controllers

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/queries"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/utils"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

// Create competition handler
func CreateCompetition(c *fiber.Ctx) error {
	var input models.CompetitionInput

	if err := c.BodyParser(&input); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(utils.InvalidData(err))
	}

	// get end_registration_date from form
	endRegistrationDateStr := "2022-01-02"
	endRegistrationDate, err := time.Parse("2006-01-02", endRegistrationDateStr) // convert string to time.Time
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("end_registration_date"))
	}

	// find Tags by id
	var tags []models.Tag
	for _, tag := range input.Tags {

		// check if there is id provided
		if tag.ID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(utils.CustomError("tag_id is required"))
		}

		// check if the input is uint type, even though it's already validated in pkg\middleware\validator.go
		if reflect.TypeOf(tag.ID).Kind() != reflect.Uint {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("tag_id"))
		}

		tagID := tag.ID
		var tag models.Tag

		// convert input to string
		tagIDStr := fmt.Sprint(tagID)

		// check if tag exists
		if err := queries.FindTagByID(tagIDStr).Scan(&tag).Error; err != nil {
			// check if tag not found
			if gorm.ErrRecordNotFound.Error() == err.Error() {
				return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("Tag"))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
		}
		tags = append(tags, tag)
	}

	// find EducationLevel by id
	var educationLevels []models.EducationLevel
	for _, educationLevel := range input.EducationLevels {

		// check if there is id provided
		if educationLevel.ID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(utils.CustomError("education_level_id is required"))
		}

		// check if the input is uint type, even though it's already validated in pkg\middleware\validator.go
		if reflect.TypeOf(educationLevel.ID).Kind() != reflect.Uint {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("education_level_id"))
		}

		educationLevelID := educationLevel.ID
		var educationLevel models.EducationLevel

		// convert input to string
		educationLevelIDStr := fmt.Sprint(educationLevelID)

		// check if education level exists
		if err := queries.FindEducationLevelByID(educationLevelIDStr).Scan(&educationLevel).Error; err != nil {
			// check if education level not found
			if gorm.ErrRecordNotFound.Error() == err.Error() {
				return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("Education Level"))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
		}
		educationLevels = append(educationLevels, educationLevel)
	}

	// check/find if user exists

	// userID := input.UserID
	// if !queries.UserExistById(userID) {
	// 	return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("User"))
	// }

	// assign input to competition struct
	competition := models.Competition{
		Name:                input.Name,
		Description:         input.Description,
		Image:               input.Image,
		Tags:                tags,
		EducationLevel:      educationLevels,
		UserID:              input.UserID,
		EndRegistrationDate: endRegistrationDate,
		CompetitionURL:      input.CompetitionURL,
	}

	// res := database.DB.Db.Create(&competition).Error

	// if res != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	// }

	return c.Status(fiber.StatusCreated).JSON(utils.SuccessCreated(competition))
}

// Get all competitions handler
func GetCompetitions(c *fiber.Ctx) error {
	var competitions []models.Competition

	// get all competitions
	if err := database.DB.Db.Order("id asc").Find(&competitions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// check if competitions is empty
	if len(competitions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.NotFound("Competitions"))
	}

	return c.Status(fiber.StatusOK).JSON(competitions)
}

// Get competition by id handler
func GetCompetition(c *fiber.Ctx) error {
	id := c.Params("id")

	// check if param is valid
	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("id"))
	}

	var competition models.Competition

	// check if competition exists, if true then assign to competition variable
	if err := queries.FindCompetitionByID(id).Scan(&competition).Error; err != nil {

		// check if competition not found
		if gorm.ErrRecordNotFound.Error() == err.Error() {
			return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("Competition"))
		}
		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(competition)
}

// Update competition by id handler
func UpdateCompetition(c *fiber.Ctx) error {
	id := c.Params("id")
	var competition models.Competition
	// check if param is valid, able to convert to uint
	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("id"))
	}

	// check if competition exists, if true then assign to competition variable
	if err := queries.FindCompetitionByID(id).Scan(&competition).Error; err != nil {

		// check if competition not found
		if gorm.ErrRecordNotFound.Error() == err.Error() {
			return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("Competition"))
		}
		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	var input models.CompetitionInput
	// parse body to competition
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrParseJson())
	}

	// get end_registration_date from form
	endRegistrationDateStr := "2022-01-02"
	endRegistrationDate, err := time.Parse("2006-01-02", endRegistrationDateStr) // convert string to time.Time
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("end_registration_date"))
	}

	// find Tags by id
	var tags []models.Tag
	for _, tag := range input.Tags {

		// check if there is id provided
		if tag.ID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(utils.CustomError("tag_id is required"))
		}

		// check if the input is uint type, even though it's already validated in pkg\middleware\validator.go
		if reflect.TypeOf(tag.ID).Kind() != reflect.Uint {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("tag_id"))
		}

		tagID := tag.ID
		var tag models.Tag

		// convert input to string
		tagIDStr := fmt.Sprint(tagID)

		// check if tag exists
		if err := queries.FindTagByID(tagIDStr).Scan(&tag).Error; err != nil {
			// check if tag not found
			if gorm.ErrRecordNotFound.Error() == err.Error() {
				return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("Tag"))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
		}
		tags = append(tags, tag)
	}

	// find EducationLevel by id
	var educationLevels []models.EducationLevel
	for _, educationLevel := range input.EducationLevels {

		// check if there is id provided
		if educationLevel.ID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(utils.CustomError("education_level_id is required"))
		}

		// check if the input is uint type, even though it's already validated in pkg\middleware\validator.go
		if reflect.TypeOf(educationLevel.ID).Kind() != reflect.Uint {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("education_level_id"))
		}

		educationLevelID := educationLevel.ID
		var educationLevel models.EducationLevel

		// convert input to string
		educationLevelIDStr := fmt.Sprint(educationLevelID)

		// check if education level exists
		if err := queries.FindEducationLevelByID(educationLevelIDStr).Scan(&educationLevel).Error; err != nil {
			// check if education level not found
			if gorm.ErrRecordNotFound.Error() == err.Error() {
				return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("Education Level"))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
		}
		educationLevels = append(educationLevels, educationLevel)
	}

	// assign input to competition struct
	competition = models.Competition{
		Name:                input.Name,
		Description:         input.Description,
		Image:               input.Image,
		Tags:                tags,            // uint id
		EducationLevel:      educationLevels, // uint id
		UserID:              input.UserID,    // uint id
		EndRegistrationDate: endRegistrationDate,
		CompetitionURL:      input.CompetitionURL,
	}

	// update competition in db
	if err := database.DB.Db.Model(&competition).Where("id = ?", id).Updates(competition).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessUpdated(competition))
}

// Delete competition by id handler
func DeleteCompetition(c *fiber.Ctx) error {
	id := c.Params("id")

	// check if param is valid, able to convert to uint
	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("id"))
	}

	var competition models.Competition

	// check if competition exists, if true then assign to competition variable
	if err := queries.FindCompetitionByID(id).Scan(&competition).Error; err != nil {

		// check if competition not found
		if gorm.ErrRecordNotFound.Error() == err.Error() {
			return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("Competition"))
		}
		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// delete competition
	if err := database.DB.Db.Delete(&competition).Error; err != nil {

		// if other error
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessDeleted())
}
