package controllers

import (
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
	endRegistrationDateStr := input.EndRegistrationDate                          // "2021-10-10"
	endRegistrationDate, err := time.Parse("2006-01-02", endRegistrationDateStr) // convert string to time.Time
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("end_registration_date"))
	}

	tags, statusCode, serverResponse := queries.FindTagsByIds(input.Tags) // iterate through tags and find tags by id
	if statusCode != 0 {
		return c.Status(statusCode).JSON(serverResponse)
	}

	educationLevels, statusCode, serverResponse := queries.FindEducationLevelsByIds(input.EducationLevels) // iterate through education_levels and find education_levels by id
	if statusCode != 0 {
		return c.Status(statusCode).JSON(serverResponse)
	}

	// check/find if user exists

	userId, exists, err := queries.GetUserById(input.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrInvalidParam("user_id"))
	}
	if !exists || userId == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("User"))
	}

	// assign input to competition struct
	competition := models.Competition{
		Name:                input.Name,
		Description:         input.Description,
		Image:               input.Image,
		Tags:                tags,
		EducationLevel:      educationLevels,
		UserID:              userId,
		EndRegistrationDate: endRegistrationDate,
		CompetitionURL:      input.CompetitionURL,
	}

	res := database.DB.Db.Create(&competition).Error

	if res != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.SuccessCreated())
}

// Get all competitions handler
func GetCompetitions(c *fiber.Ctx) error {
	var competitions []models.Competition

	// get query parameters for pagination and filtering
	page, err := strconv.Atoi(c.Query("page", "8"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CustomError("Invalid page number"))
	}
	// pageSize, err := strconv.Atoi(c.Query("page_size", "10"))
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(utils.CustomError("Invalid page size"))
	// }
	// tags := c.Query("tags")
	// search := c.Query("search")
	// eduLevels := c.Query("edu_levels")

	// page size must not be greater than 20
	// if pageSize > 20 {
	// 	return c.Status(fiber.StatusBadRequest).JSON(utils.CustomError("Page size must not be greater than 20"))
	// }
	pageSize := 8

	// calculate offset and limit for pagination
	offset := (page - 1) * pageSize
	limit := pageSize

	// get all competitions with pagination and filtering
	db := database.DB.Db
	// if tags != "" {
	// 	db = db.Joins("JOIN competition_tags ON competition_tags.competition_id = competitions.id").
	// 		Joins("JOIN tags ON tags.id = competition_tags.tag_id").
	// 		Where("tags.name IN (?)", strings.Split(tags, ","))
	// }
	// if search != "" {
	// 	db = db.Where("competitions.name LIKE ?", "%"+search+"%")
	// }
	// if eduLevels != "" {
	// 	db = db.Joins("JOIN education_levels ON education_levels.id = competitions.education_level_id").
	// 		Where("education_levels.name IN (?)", strings.Split(eduLevels, ","))
	// }

	if err = db.Offset(offset).Limit(limit).Find(&competitions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	var total int64
	if err := db.Model(&models.Competition{}).Count(&total).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// check if competitions is empty
	if len(competitions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.NotFound("Competitions"))
	}

	var competitionResponses []models.CompetitionResponse
	for _, competition := range competitions {

		// check competition_tags db and find tag name by id
		tagsResponse, err := queries.CompeFindTagsByNames(competition)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
		}

		// check competition_education_levels db and find education_level name by id
		eduLevelsResponse, err := queries.CompeFindEducationLevelsByNames(competition)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
		}

		// convert userId to string
		userId := strconv.FormatUint(uint64(competition.UserID), 10)

		// create competition response
		competitionResponse := models.CompetitionResponse{
			ID:                  competition.ID,
			Name:                competition.Name,
			Description:         competition.Description,
			Image:               competition.Image,
			Tags:                tagsResponse,
			EducationLevels:     eduLevelsResponse,
			UserID:              userId,
			EndRegistrationDate: competition.EndRegistrationDate,
			CompetitionURL:      competition.CompetitionURL,
		}

		// add competition response to slice
		competitionResponses = append(competitionResponses, competitionResponse)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": competitionResponses,
		"meta": fiber.Map{
			"total_data": total,
			"page_data":  len(competitions),
			"page":       page,
			"last_page":  total/int64(limit) + 1,
		},
	})
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

	tags, status, serverResponse := queries.FindTagsByIds(input.Tags) // iterate through tags and find tags by id
	if status != 0 {
		return c.Status(status).JSON(serverResponse)
	}

	educationLevels, status, serverResponse := queries.FindEducationLevelsByIds(input.EducationLevels) // iterate through education_levels and find education_levels by id
	if status != 0 {
		return c.Status(status).JSON(serverResponse)
	}

	// check/find if user exists
	userId, exists, err := queries.GetUserById(input.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrInvalidParam("user_id"))
	}
	if !exists || userId == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("User"))
	}

	// check/find if user exists
	userId, exists, err = queries.GetUserById(input.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrInvalidParam("user_id"))
	}
	if !exists || userId == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.IDNotFound("User"))
	}

	// assign input to competition struct
	competition = models.Competition{
		Name:                input.Name,
		Description:         input.Description,
		Image:               input.Image,
		Tags:                tags,            // uint id
		EducationLevel:      educationLevels, // uint id
		UserID:              userId,          // uint id
		EndRegistrationDate: endRegistrationDate,
		CompetitionURL:      input.CompetitionURL,
	}

	// update competition in db
	if err := database.DB.Db.Model(&competition).Updates(competition).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessUpdated())
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
