package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/LavaJover/storage-master/storage-service/internal/config"
	"github.com/LavaJover/storage-master/storage-service/internal/db"
)

func main(){
	cfg := config.MustLoad()
	fmt.Println(cfg)

	_, err := db.InitDB(cfg.Dsn)

	if err != nil{
		log.Fatalf("failed to init db: %v\n", err)
	}

	slog.Info("connected to db!")
}