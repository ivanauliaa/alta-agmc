package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
}

func ExtractPayload(token string) (uint, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	return uint(claims["userId"].(float64)), nil
}
