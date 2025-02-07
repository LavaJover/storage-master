package server

import (
	"context"
	"log/slog"

	"github.com/LavaJover/storage-master/storage-service/pkg/models"
	"github.com/LavaJover/storage-master/storage-service/internal/service"
	storagepb "github.com/LavaJover/storage-master/storage-service/proto/gen"
)

type StorageServer struct{
	*service.StorageService
	storagepb.UnimplementedStorageServiceServer
}

// Storage methods
func (server *StorageServer) CreateStorage (ctx context.Context, r *storagepb.CreateStorageRequest) (*storagepb.CreateStorageResponse, error){
	newStorage := &models.Storage{
		Name: r.Name,
		UserID: r.UserId,
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

func (server *StorageServer) GetStorages (ctx context.Context, r *storagepb.GetStoragesRequest) (*storagepb.GetStoragesResponse, error){
	userID := r.UserId

	// Get storages from service layer
	storages, err := server.StorageService.GetStoragesByUserID(userID)
	if err != nil{
		slog.Error("failed to get storages by userID", "msg", err.Error())
		return nil, err
	}

	// Convert storages to rpc responseStorages
	var responseStorages []*storagepb.Storage
	for _, storage := range storages{
		responseStorages = append(responseStorages, 
			&storagepb.Storage{
				Id: uint64(storage.ID),
				Name: storage.Name,
				UserId: storage.UserID,
			},
		)
	}

	return &storagepb.GetStoragesResponse{
		Storages: responseStorages,
	}, nil
}

// Cell methods
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

func (server *StorageServer) GetCells (ctx context.Context, r *storagepb.GetCellsRequest) (*storagepb.GetCellsResponse, error){
	storageID := r.StorageId

	// Get storages from service layer
	cells, err := server.StorageService.GetCellsByStorageID(storageID)
	if err != nil{
		slog.Error("failed to get storages by userID", "msg", err.Error())
		return nil, err
	}

	// Convert storages to rpc responseStorages
	var responseCells []*storagepb.Cell
	for _, cell := range cells{
		responseCells = append(responseCells, 
			&storagepb.Cell{
				Id: uint64(cell.ID),
				Name: cell.Name,
				StorageId: uint64(cell.StorageID),
			},
		)
	}

	return &storagepb.GetCellsResponse{
		Cells: responseCells,
	}, nil
}

// Box methods
func (server *StorageServer) GetBoxes (ctx context.Context, r *storagepb.GetBoxesRequest) (*storagepb.GetBoxesResponse, error){
	cellID := r.CellId

	// Get storages from service layer
	boxes, err := server.StorageService.GetBoxesByCellID(cellID)
	if err != nil{
		slog.Error("failed to get storages by userID", "msg", err.Error())
		return nil, err
	}

	// Convert storages to rpc responseStorages
	var responseBoxes []*storagepb.Box
	for _, box := range boxes{
		responseBoxes = append(responseBoxes, 
			&storagepb.Box{
				Id: uint64(box.ID),
				Name: box.Name,
				CellId: uint64(box.CellID),
			},
		)
	}

	return &storagepb.GetBoxesResponse{
		Boxes: responseBoxes,
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