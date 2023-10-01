package utils

import (
	"errors"
	"os"
	"server/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	User models.User
	jwt.StandardClaims
}

func GenerateJwt(user *models.User) (string, error) {

	claims := CustomClaims{
		*user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			Issuer:    "server",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("APP_JWT")))
}

func CheckJwt(jwtString string) error {

	claims, err := jwt.ParseWithClaims(jwtString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("APP_JWT")), nil
	})

	if !claims.Valid {
		return errors.New("token not valid")
	}

	return err
}

func GetClaims(jwtString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(jwtString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("APP_JWT")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); claims != nil && ok {
		return claims, nil
	} else {
		return nil, nil
	}
}
