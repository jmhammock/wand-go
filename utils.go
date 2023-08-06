package main

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jmhammock/wand-go/models"
	"golang.org/x/crypto/bcrypt"
)

const secret = "SuperSecretKey"

func hashpw(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func genjwt(user *models.User) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   user.Id.String(),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validatejwt(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
