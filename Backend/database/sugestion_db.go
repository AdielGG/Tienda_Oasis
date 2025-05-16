package database

import (
	"backend/models"
)

func CreateSuggestion(suggestion models.Suggestion) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Create(&suggestion).Error
}

func GetAllSuggestions() ([]models.Suggestion, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var suggestions []models.Suggestion
	err = db.Find(&suggestions).Error
	if err != nil {
		return nil, err
	}
	return suggestions, nil
}
func GetSuggestion(id int) (models.Suggestion, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Suggestion{}, err
	}
	var suggestion models.Suggestion
	err = db.First(&suggestion, id).Error
	if err != nil {
		return models.Suggestion{}, err
	}
	return suggestion, nil
}

func DeleteSuggestions(id int) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Delete(&models.Suggestion{}, id).Error
}

func UpdateSuggestion(suggestion models.Suggestion) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Save(&suggestion).Error
}

func DeleteSuggestion(id int) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Delete(&models.Suggestion{}, id).Error
}
