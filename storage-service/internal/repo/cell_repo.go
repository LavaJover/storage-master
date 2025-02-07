package repo

import (
	"log/slog"

	"github.com/LavaJover/storage-master/storage-service/pkg/models"
	"gorm.io/gorm"
)

type CellRepo struct{
	*gorm.DB
}

// Insert new cell instance to database
func (cellRepo *CellRepo) InsertCell (cell *models.Cell) error{
	result := cellRepo.DB.Create(cell)

	if result.Error != nil{
		return result.Error
	}

	slog.Info("inserted new Cell instance", "ID", cell.ID)
	return nil
}

// Get all cell instances by storage_id
func (cellRepo *CellRepo) GetCellsByStorageID (storageID uint64) ([]models.Cell, error){
	var cells []models.Cell
	err := cellRepo.DB.Where("storage_id = ?", storageID).Preload("Boxes").Find(&cells).Error
	return cells, err
}