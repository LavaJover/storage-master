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

// Cell service methods
func (service *StorageService) AddCell (cell *models.Cell) error {
	return service.CellRepo.InsertCell(cell)
}

// Box service methods
func (service *StorageService) AddBox (box *models.Box) error {
	return service.BoxRepo.InsertBox(box)
}