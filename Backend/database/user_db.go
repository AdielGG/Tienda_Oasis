package database

import (
	"backend/models"
)

func CreateUser(user models.User) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Create(&user).Error
}

func GetUser(id int) (models.User, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	err = db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUser(user models.User) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Save(&user).Error
}

func DeleteUser(id int) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Delete(&models.User{}, id).Error
}
