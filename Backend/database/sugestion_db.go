package database

import (
	"backend/models"
)

func CreateSuggestion(suggestion models.Suggestion) error {
	return DB.Create(&suggestion).Error
}

func GetAllSuggestions() ([]models.Suggestion, error) {
	var suggestions []models.Suggestion
	err := DB.Find(&suggestions).Error
	if err != nil {
		return nil, err
	}
	return suggestions, nil
}
func GetSuggestion(id int) (models.Suggestion, error) {
	var suggestion models.Suggestion
	err := DB.First(&suggestion, id).Error
	if err != nil {
		return models.Suggestion{}, err
	}
	return suggestion, nil
}

func DeleteSuggestions(id int) error {
	return DB.Delete(&models.Suggestion{}, id).Error
}

func UpdateSuggestion(suggestion models.Suggestion) error {
	return DB.Save(&suggestion).Error
}

func DeleteSuggestion(id int) error {
	return DB.Delete(&models.Suggestion{}, id).Error
}
