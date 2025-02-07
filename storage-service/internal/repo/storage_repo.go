package repo

import (
	"log/slog"

	"github.com/LavaJover/storage-master/storage-service/pkg/models"
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

// Get all storage instances by user_id
func (storageRepo *StorageRepo) GetStoragesByUserID (userID uint64) ([]models.Storage, error) {
	var storages []models.Storage
	err := storageRepo.DB.Where("user_id = ?", userID).Preload("Cells.Boxes").Find(&storages).Error
	return storages, err
}