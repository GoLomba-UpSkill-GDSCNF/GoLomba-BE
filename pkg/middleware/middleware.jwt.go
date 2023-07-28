package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/repository"
)

func CreateToken(id, role uint) (string, error) {
	if id == 0 {
		return "", errors.New("failed create token")
	}
	
	claims := jwt.MapClaims{
		"id":id,
		"role":role,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(repository.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return t, nil
}

