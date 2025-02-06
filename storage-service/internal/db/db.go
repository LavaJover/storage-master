package db

import (
	"github.com/LavaJover/storage-master/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDB(dsn string) (*gorm.DB, error){
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		return nil, err
	}

	db.AutoMigrate(&models.Storage{})
	db.AutoMigrate(&models.Cell{})
	db.AutoMigrate(&models.Box{})

	return db, nil
}