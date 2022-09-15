package database

import (
	"day3-middleware/config"
	"day3-middleware/models"
	"fmt"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(user models.User) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserById(id uint) (interface{}, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserById(id uint, userUpdate models.User) (interface{}, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	config.DB.Model(&user).Updates(userUpdate)

	return user, nil
}

func DeleteUserById(id uint) error {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func LoginUser(userLogin models.UserLogin) (interface{}, error) {
	user := models.User{}
	if err := config.DB.Where("email = ? AND password = ?", userLogin.Email, userLogin.Password).
		First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func VerifyUserOwner(userId uint, owner uint) error {
	if userId != owner {
		return fmt.Errorf("restricted resource")
	}

	return nil
}
