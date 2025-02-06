package repo_test

import (
	"testing"

	"github.com/LavaJover/storage-master/storage-service/internal/config"
	"github.com/LavaJover/storage-master/storage-service/internal/db"
	"github.com/LavaJover/storage-master/storage-service/internal/models"
)

func TestInsertBox(t *testing.T){
	cfg := config.MustLoad()

	db, err := db.InitDB(cfg.Dsn)

	if err != nil{
		t.Fatalf("failed to init db: %v\n", err)
	}

	db.AutoMigrate(&models.Box{})
	result := db.Create(&models.Box{Name: "TestBoxName", CellID: 1})

	if result.Error != nil{
		t.Fatalf("failed to insert Box instance")
	}
}