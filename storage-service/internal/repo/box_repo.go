package repo

import (
	"log/slog"

	"github.com/LavaJover/storage-master/storage-service/pkg/models"
	"gorm.io/gorm"
)

type BoxRepo struct{
	*gorm.DB
}


// Insert new box instance to database
func (boxRepo *BoxRepo) InsertBox (box *models.Box) error{
	result := boxRepo.DB.Create(box)

	if result.Error != nil{
		return result.Error
	}

	slog.Info("inserted new Box instance", "ID", box.ID)
	return nil
}

// Get box instances from database by cell_id
func (boxRepo *BoxRepo) GetBoxesByCellID (cellID uint64) ([]models.Box, error){
	var boxes []models.Box
	err := boxRepo.DB.Where("cell_id = ?", cellID).Find(&boxes).Error
	return boxes, err
}