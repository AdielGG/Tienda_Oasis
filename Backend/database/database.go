package database

import (
	cfg "backend/config"
	"backend/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var er error

func InitDB() {
	DB, er = CreateDB()

	if er != nil {
		panic(er)
	}
}

func CreateDB() (*gorm.DB, error) {

	dbURL := "host=" + cfg.DbConfig.Host + " port=" + cfg.DbConfig.Port + " user=" + cfg.DbConfig.User + " password=" + cfg.DbConfig.Password + " dbname=" + cfg.DbConfig.Database
	fmt.Println(dbURL + " \n\n\n")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Suggestion{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
