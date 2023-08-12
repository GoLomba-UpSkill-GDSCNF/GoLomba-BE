package queries

import (
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/app/models"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
)

func SaveTestimony(data *models.Testimonial) error {
	return database.DB.Db.Save(data).Error
}

func GetAllTestimony() ([]models.TestimonialResponse, error) {
	var testimonials []models.TestimonialResponse
	if err := database.DB.Db.Table("testimonials").Select("testimonials.id, testimonials.user_id, testimonials.message, testimonials.stars, users.full_name").Joins("left join users ON users.id = testimonials.user_id").Scan(&testimonials).Error; err != nil {
		return nil, err
	}

	return testimonials, nil
}
