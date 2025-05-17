package database

import (
	"backend/models"
	"fmt"
)

func CreateUser(user models.User) error {
	return DB.Create(&user).Error
}

func GetUser(id int) (models.User, error) {
	var user models.User
	err := DB.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUserByUserName(username string) (models.User, error) {
	var user models.User

	err := DB.Where(fmt.Sprintf("username = '%s'", username)).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUser(user models.User) error {
	return DB.Save(&user).Error
}

func DeleteUser(id int) error {
	return DB.Delete(&models.User{}, id).Error
}
