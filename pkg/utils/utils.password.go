package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPass,err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}

func ComparePassword(password, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password)); err != nil {
		return err
	}

	return nil
}