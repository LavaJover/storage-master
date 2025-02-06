package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/LavaJover/storage-master/internal/config"
	"github.com/LavaJover/storage-master/internal/db"
	"github.com/LavaJover/storage-master/internal/repo"
	"github.com/LavaJover/storage-master/internal/server"
	"github.com/LavaJover/storage-master/internal/service"
	storagepb "github.com/LavaJover/storage-master/proto/gen"
	"google.golang.org/grpc"
)

func main(){
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// Initialize db layer
	storageDB, err := db.InitDB(cfg.Dsn)
	if err != nil{
		log.Fatalf("failed to init db: %v\n", err)
	}
	slog.Info("connected to db!")

	// Initialize repository layer
	boxRepo, cellRepo, storageRepo := repo.BoxRepo{storageDB}, repo.CellRepo{storageDB}, repo.StorageRepo{storageDB}

	// Initialize service layer
	storageService := service.StorageService{
		BoxRepo: &boxRepo,
		CellRepo: &cellRepo,
		StorageRepo: &storageRepo,
	}

	// Initialize server layer
	storageServer := server.StorageServer{
		StorageService: &storageService,
	}
	grpcServer := grpc.NewServer()
	storagepb.RegisterStorageServiceServer(grpcServer, &storageServer)

	// Starting server
	listener, err := net.Listen("tcp", cfg.Host+":"+cfg.Port)

	if err != nil{
		log.Fatalf("failed to listen %s:%s\n", cfg.Host, cfg.Port)
	}

	slog.Info("grpc StorageServer running on " + cfg.Host + ":" + cfg.Port)

	if err := grpcServer.Serve(listener); err != nil{
		log.Fatalf("failed to serve grpc server: %v", err)
	}
	
}