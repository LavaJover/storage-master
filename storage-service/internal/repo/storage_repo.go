package repo

import (
	"log/slog"

	"github.com/LavaJover/storage-master/storage-service/internal/models"
	"gorm.io/gorm"
)

type StorageRepo struct{
	*gorm.DB
}

// Insert new storage instance to database
func (storageRepo *StorageRepo) InsertStorage (storage *models.Storage) error {
	result := storageRepo.DB.Create(storage)

	if result.Error != nil{
		return result.Error
	}

	slog.Info("inserted new Storage instance", "ID", storage.ID)

	return nil
}