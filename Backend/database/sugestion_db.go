package database

import (
	"backend/models"
)

func CreateSugestion(sugestion models.Sugestion) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Create(&sugestion).Error
}

func GetAllSugestions() ([]models.Sugestion, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var sugestions []models.Sugestion
	err = db.Find(&sugestions).Error
	if err != nil {
		return nil, err
	}
	return sugestions, nil
}
func GetSugestion(id int) (models.Sugestion, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Sugestion{}, err
	}
	var sugestion models.Sugestion
	err = db.First(&sugestion, id).Error
	if err != nil {
		return models.Sugestion{}, err
	}
	return sugestion, nil
}

func DeleteSugestions(id int) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Delete(&models.Sugestion{}, id).Error
}

func UpdateSugestion(sugestion models.Sugestion) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Save(&sugestion).Error
}

func DeleteSugestion(id int) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Delete(&models.Sugestion{}, id).Error
}
