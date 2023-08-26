package controllers

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/queries"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/middleware"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/utils"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
	"gorm.io/gorm"
)

// Create competition handler
func CreateCompetition(c *fiber.Ctx) error {
	var input models.CompetitionInput

	tokenJWT := ""

	if authHeader := c.Request().Header.Peek("Authorization"); len(authHeader) > 0 {
		tokenJWT = strings.Fields(string(authHeader))[1]
	}

	userID, _, err := middleware.CheckTokenValue(tokenJWT)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ServerError(err))
	}

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

	// assign input to competition struct
	competition := models.Competition{
		Name:                input.Name,
		Description:         input.Description,
		Image:               input.Image,
		Tags:                tags,
		EducationLevels:     educationLevels,
		UserID:              uint(userID.(float64)),
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

	// convert string to int
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrInvalidParam("page"))
	}
	sort := c.Query("sort", "asc")

	tags := c.Query("tags", "")
	search := c.Query("search", "")
	eduLevels := c.Query("edu_levels", "")

	// pagination
	pagination := &utils.Pagination{
		Limit: 8, // default limit is 8
		Page:  page,
		Sort:  "id " + sort,
	}
	cg := &utils.CompetitionGorm{
		DB: database.DB.Db,
	}
	if _, err := cg.ListCompetition(pagination, tags, search, eduLevels); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// check if competitions is empty in db
	if len(pagination.Rows.([]*models.Competition)) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.NotFound("Competitions"))
	}

	var competitionResponses []models.CompetitionResponse
	for _, competition := range pagination.Rows.([]*models.Competition) {

		// iterate through tags and find tags by id
		tagsResponse, err := queries.CompeFindTagsByNames(*competition)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
		}

		// iterate through education_levels and find education_levels by id
		eduLevelsResponse, err := queries.CompeFindEducationLevelsByNames(*competition)
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
	pagination.TotalData = len(pagination.Rows.([]*models.Competition))
	pagination.Rows = competitionResponses
	return c.Status(fiber.StatusOK).JSON(pagination)
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

	// iterate through tags and find tags by names in competition_tags table
	tagsResponse, err := queries.CompeFindTagsByNames(competition)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// iterate through education_levels and find education_levels by names in competition_education_levels table
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

	// return competition response
	return c.Status(fiber.StatusOK).JSON(competitionResponse)
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

	tokenJWT := ""

	if authHeader := c.Request().Header.Peek("Authorization"); len(authHeader) > 0 {
		tokenJWT = strings.Fields(string(authHeader))[1]
	}

	userID, _, err := middleware.CheckTokenValue(tokenJWT)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ServerError(err))
	}

	// check if user role is admin or not and if user id is same as competition user id
	if userID.(float64) != float64(competition.UserID) {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Unauthorized())
	} else if userID.(float64) != 1 {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Unauthorized())
	}

	var input models.CompetitionInput
	// parse body to competition
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrParseJson())
	}

	// find tags in competition_tags table by competition_id and delete all of them
	if err := queries.FindCompeDeleteTagsById(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// find competition_education_levels in competition_education_levels table by competition_id and delete all of them
	if err := queries.FindCompeDeleteEduLevelsByCompetitionID(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
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

	// assign input to competition struct
	competition = models.Competition{
		Name:        input.Name,
		Description: input.Description,
		Image:       input.Image,
		// Tags:                tags,            // uint id
		// EducationLevel:      educationLevels, // uint id
		UserID:              uint(userID.(float64)), // uint id
		EndRegistrationDate: endRegistrationDate,
		CompetitionURL:      input.CompetitionURL,
	}

	if err := queries.UpdateCompeTagsById(id, tags); err != nil {

		// if error is not found
		if gorm.ErrRecordNotFound.Error() != err.Error() {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.NotFound("Tags"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	if err := queries.UpdateCompeEduLevelsById(id, educationLevels); err != nil {

		// if error is not found
		if gorm.ErrRecordNotFound.Error() != err.Error() {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.NotFound("Education Levels"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
	}

	// update competition
	if err := database.DB.Db.Model(&competition).Where("id = ?", id).Updates(map[string]interface{}{
		"name":                  competition.Name,
		"description":           competition.Description,
		"image":                 competition.Image,
		"user_id":               competition.UserID,
		"end_registration_date": competition.EndRegistrationDate,
		"competition_url":       competition.CompetitionURL,
	}).Error; err != nil {
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

func UserCompetitions(c *fiber.Ctx) error {
	var tokenJWT string

	if authHeader := c.Request().Header.Peek("Authorization"); len(authHeader) > 0 {
		tokenJWT = strings.Fields(string(authHeader))[1]
	}

	userID, _, err := middleware.CheckTokenValue(tokenJWT)
	if err != nil {
		log.Println(err)
		return utils.SendResponse(c, fiber.StatusInternalServerError, "failed get user profile", nil)
	}

	user, err := queries.GetUserById(uint(userID.(float64)))
	if err != nil {
		log.Println(err)
		return utils.SendResponse(c, fiber.StatusInternalServerError, "failed get user profile", nil)
	}

	var competitions []models.Competition

	// if user role is admin get all competitions
	if user.RoleID == 1 {
		if err := database.DB.Db.Find(&competitions).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
		}
	} else {
		// if user role is not admin get competitions by user id
		competitions, err = queries.FindCompetitionByUserID(uint(userID.(float64)))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
		}
	}

	// check if competitions is empty in db
	if len(competitions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.NotFound("Competitions"))
	}

	var competitionResponses []models.CompetitionResponse
	for _, competition := range competitions {

		// iterate through tags and find tags by id
		tagsResponse, err := queries.CompeFindTagsByNames(competition)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerError(err))
		}

		// iterate through education_levels and find education_levels by id
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

	// return competition response
	return c.Status(fiber.StatusOK).JSON(competitionResponses)
}
