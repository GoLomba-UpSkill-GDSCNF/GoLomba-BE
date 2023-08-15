package utils

import (
	"math"

	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	Sort       string      `json:"sort"`
	TotalRows  int64       `json:"total_rows"`
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

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

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

func (cg *CompetitionGorm) ListCompetition(pagination *Pagination) (*Pagination, error) {
	var competitions []*models.Competition

	if err := cg.DB.Scopes(Paginate(competitions, pagination, cg.DB)).Find(&competitions).Error; err != nil {
		return nil, err
	}

	pagination.Rows = competitions
	return pagination, nil
}
