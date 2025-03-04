package database

import (
	"backend/models"
)

func CreateProduct(product models.Product) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Create(&product).Error
}

func GetProduct(id int) (models.Product, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Product{}, err
	}
	var product models.Product
	err = db.First(&product, id).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func GetAllProducts() ([]models.Product, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var products []models.Product
	err = db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func UpdateProduct(product models.Product) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Save(&product).Error
}

func DeleteProduct(id int) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Delete(&models.Product{}, id).Error
}
