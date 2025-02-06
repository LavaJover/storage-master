package repo

import (
	"log/slog"

	"github.com/LavaJover/storage-master/storage-service/internal/models"
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