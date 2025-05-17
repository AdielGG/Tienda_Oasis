package database

import (
	"backend/models"
)

func CreateProduct(product models.Product) error {
	return DB.Create(&product).Error
}

func GetProduct(id int) (models.Product, error) {

	var product models.Product

	err := DB.First(&product, id).Error

	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	err := DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func UpdateProduct(product models.Product) error {

	return DB.Save(&product).Error
}

func DeleteProduct(id int) error {

	return DB.Delete(&models.Product{}, id).Error
}

func GetProductByType(Type string) (models.Product, error) {

	var product models.Product

	err := DB.Where("type = ?", Type).First(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func GetProductByCategory(Category string) (models.Product, error) {

	var product models.Product

	err := DB.Where("category = ?", Category).First(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
