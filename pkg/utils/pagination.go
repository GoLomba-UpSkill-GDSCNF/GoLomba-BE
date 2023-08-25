package utils

import (
	"math"
	"strings"

	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	Sort       string      `json:"sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalData  int         `json:"total_data,omitempty"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		return 8
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		return "id asc"
	}
	return p.Sort
}

func PaginateCompetition(value interface{}, pagination *Pagination, db *gorm.DB, tags string, search string, eduLevels string) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	query := db.Model(value)

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if tags != "" {
		tagNames := strings.Split(tags, ",")
		subquery := db.Table("competition_tags").
			Select("competition_id, COUNT(*) AS tag_count").
			Joins("JOIN tags ON competition_tags.tag_id = tags.id").
			Where("tags.name IN (?)", tagNames).
			Group("competition_id").
			Having("tag_count = ?", len(tagNames))

		query = query.Joins("JOIN (?) AS matching_competitions ON competitions.id = matching_competitions.competition_id", subquery)
	}

	if eduLevels != "" {
		educationLevels := strings.Split(eduLevels, ",")
		subquery := db.Table("competition_education_levels").
			Select("competition_id, COUNT(*) AS education_level_count").
			Joins("JOIN education_levels ON competition_education_levels.education_level_id = education_levels.id").
			Where("education_levels.name IN (?)", educationLevels).
			Group("competition_id").
			Having("education_level_count = ?", len(educationLevels))

		query = query.Joins("JOIN (?) AS matching_competitions ON competitions.id = matching_competitions.competition_id", subquery)
	}

	query.Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

type CompetitionGorm struct {
	DB *gorm.DB
}

func (cg *CompetitionGorm) ListCompetition(pagination *Pagination, tags string, search string, eduLevels string) (*Pagination, error) {
	var competitions []*models.Competition

	query := cg.DB.Scopes(PaginateCompetition(competitions, pagination, cg.DB, tags, search, eduLevels))

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if tags != "" {
		tagNames := strings.Split(tags, ",")

		// subquery to get competition_id that has all tags
		subquery := cg.DB.Table("competition_tags").
			Select("competition_id, COUNT(*) AS tag_count").
			Joins("JOIN tags ON competition_tags.tag_id = tags.id").
			Where("tags.name IN (?)", tagNames).
			Group("competition_id").
			Having("tag_count = ?", len(tagNames))

		query = query.Joins("JOIN (?) AS matching_competitions ON competitions.id = matching_competitions.competition_id", subquery)
	}

	if eduLevels != "" {
		educationLevels := strings.Split(eduLevels, ",")

		// subquery to get competition_id that has all education levels
		subquery := cg.DB.Table("competition_education_levels").
			Select("competition_id, COUNT(*) AS education_level_count").
			Joins("JOIN education_levels ON competition_education_levels.education_level_id = education_levels.id").
			Where("education_levels.name IN (?)", educationLevels).
			Group("competition_id").
			Having("education_level_count = ?", len(educationLevels))

		query = query.Joins("JOIN (?) AS matching_competitions ON competitions.id = matching_competitions.competition_id", subquery)
	}

	if err := query.Find(&competitions).Error; err != nil {
		return nil, err
	}

	pagination.Rows = competitions
	return pagination, nil
}

func UserPaginateCompetitions(value interface{}, pagination *Pagination, db *gorm.DB, userID int) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	query := db.Model(value)

	// query get competitions that user joined
	query = query.Where("user_id = ?", userID)

	query.Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func (cg *CompetitionGorm) ListUserCompetition(pagination *Pagination, userID int) (*Pagination, error) {
	var competitions []*models.Competition

	query := cg.DB.Scopes(UserPaginateCompetitions(competitions, pagination, cg.DB, userID))

	// query get competitions that user joined
	query = query.Where("user_id = ?", userID)

	if err := query.Find(&competitions).Error; err != nil {
		return nil, err
	}

	pagination.Rows = competitions
	return pagination, nil
}
