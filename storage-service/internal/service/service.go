package service

import (
	"github.com/LavaJover/storage-master/storage-service/pkg/models"
	"github.com/LavaJover/storage-master/storage-service/internal/repo"
)


type StorageService struct{
	*repo.BoxRepo
	*repo.CellRepo
	*repo.StorageRepo
}

// Storage service methods
func (service *StorageService) CreateStorage (storage *models.Storage) error {
	return service.StorageRepo.InsertStorage(storage)
}

func (service *StorageService) GetStorages (userID uint64) ([]models.Storage, error){
	return service.StorageRepo.GetStoragesByUserID(userID)
}

// Cell service methods
func (service *StorageService) AddCell (cell *models.Cell) error {
	return service.CellRepo.InsertCell(cell)
}

func (service *StorageService) GetCells (storageID uint64) ([]models.Cell, error) {
	return service.CellRepo.GetCellsByStorageID(storageID)
}

// Box service methods
func (service *StorageService) AddBox (box *models.Box) error {
	return service.BoxRepo.InsertBox(box)
}

func (service *StorageService) GetBoxes (cellID uint64) ([]models.Box, error){
	return service.BoxRepo.GetBoxesByCellID(cellID)
}