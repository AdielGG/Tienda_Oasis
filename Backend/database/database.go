package database

import (
	cfg "backend/servercfg"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dbURL := "host=" + cfg.DbConfig.Host + " port=" + cfg.DbConfig.Port + " user=" + cfg.DbConfig.User + " password=" + cfg.DbConfig.Password + " dbname=" + cfg.DbConfig.Database

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
