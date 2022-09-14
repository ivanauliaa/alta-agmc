package database

import (
	"day2-dynamic/config"
	"day2-dynamic/models"
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

func GetUserById(id string) (interface{}, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserById(id string, userUpdate models.User) (interface{}, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	config.DB.Model(&user).Updates(userUpdate)

	return user, nil
}

func DeleteUserById(id string) error {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
