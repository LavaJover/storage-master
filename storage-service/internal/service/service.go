package service

import (
	"github.com/LavaJover/storage-master/pkg/models"
	"github.com/LavaJover/storage-master/internal/repo"
)


type StorageService struct{
	*repo.BoxRepo
	*repo.CellRepo
	*repo.StorageRepo
}

func (service *StorageService) CreateStorage (storage *models.Storage) error {
	return service.StorageRepo.InsertStorage(storage)
}

func (service *StorageService) AddCell (cell *models.Cell) error {
	return service.CellRepo.InsertCell(cell)
}

func (service *StorageService) AddBox (box *models.Box) error {
	return service.BoxRepo.InsertBox(box)
}