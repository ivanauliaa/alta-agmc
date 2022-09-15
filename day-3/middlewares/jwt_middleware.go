package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const jwtSecret = "ac78b74522535c33bcc9535e9ad217e74aafd0a29f8047853551130e2a853941b3bce9466adc35fbc846e89f8df504d24687a77b51fb1be4b825bef2e51199dc"

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func ExtractTokenUserId(c *echo.Context) uint {
	user := (*c).Get("user").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return uint(userId)
	}

	return 0
}

func GetJwtSecret() string {
	return jwtSecret
}
