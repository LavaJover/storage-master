package server

import (
	"context"
	"log/slog"

	"github.com/LavaJover/storage-master/pkg/models"
	"github.com/LavaJover/storage-master/internal/service"
	storagepb "github.com/LavaJover/storage-master/proto/gen"
)

type StorageServer struct{
	*service.StorageService
	storagepb.UnimplementedStorageServiceServer
}

func (server *StorageServer) CreateStorage (ctx context.Context, r *storagepb.CreateStorageRequest) (*storagepb.CreateStorageResponse, error){
	newStorage := &models.Storage{
		Name: r.Name,
	}

	err := server.StorageService.CreateStorage(newStorage)

	if err != nil{
		slog.Error("failed to create Storage instance", "error", err.Error())
		return nil, err
	}

	return &storagepb.CreateStorageResponse{
		Id: uint64(newStorage.ID),
		Name: newStorage.Name,
	}, nil
}

func (server *StorageServer) AddCell (ctx context.Context, r *storagepb.AddCellRequest) (*storagepb.AddCellResponse, error){
	newCell := &models.Cell{
		Name: r.Name,
		StorageID: uint(r.StorageId),
	}

	err := server.StorageService.AddCell(newCell)

	if err != nil{
		slog.Error("failed to create Cell instance", "error", err.Error())
		return nil, err
	}

	return &storagepb.AddCellResponse{
		Id: uint64(newCell.ID),
		Name: newCell.Name,
		StorageId: uint64(newCell.StorageID),
	}, nil
}

func (server *StorageServer) AddBox (ctx context.Context, r *storagepb.AddBoxRequest) (*storagepb.AddBoxResponse, error){
	newBox := &models.Box{
		Name: r.Name,
		CellID: uint(r.CellId),
	}

	err := server.StorageService.AddBox(newBox)

	if err != nil{
		slog.Error("failed to create Box instance", "error", err.Error())
		return nil, err
	}

	return &storagepb.AddBoxResponse{
		Id: uint64(newBox.ID),
		Name: newBox.Name,
		CellId: uint64(newBox.CellID),
	}, nil
}