package helpers

import (
	"day3-middleware/config"
	"day3-middleware/middlewares"
	"day3-middleware/models"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func AddUserWithId(id uint) {
	user := models.User{
		Model: gorm.Model{
			ID: id,
		},
		Name:     "alta",
		Email:    "alta@gmail.com",
		Password: "alta",
	}

	config.DB.Create(&user)
}

func AddUserWithEmailPassword(email, password string) {
	user := models.User{
		Name:     "alta",
		Email:    email,
		Password: password,
	}

	config.DB.Create(&user)
}

func AddUserWithIdEmailPassword(id uint, email, password string) {
	user := models.User{
		Model: gorm.Model{
			ID: id,
		},
		Name:     "alta",
		Email:    email,
		Password: password,
	}

	config.DB.Create(&user)
}

func CleanUsersTable() {
	config.DB.Exec("DELETE FROM users")
}

func GetJwtToken(email, password string) string {
	user := models.User{}
	config.DB.Where("email = ? AND password = ?", email, password).
		First(&user)

	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["userId"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte(middlewares.GetJwtSecret()))
	return tokenStr
}
