// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: storage.proto

package storagepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StorageService_CreateStorage_FullMethodName = "/storage.StorageService/CreateStorage"
	StorageService_AddCell_FullMethodName       = "/storage.StorageService/AddCell"
	StorageService_AddBox_FullMethodName        = "/storage.StorageService/AddBox"
	StorageService_GetStorages_FullMethodName   = "/storage.StorageService/GetStorages"
	StorageService_GetCells_FullMethodName      = "/storage.StorageService/GetCells"
	StorageService_GetBoxes_FullMethodName      = "/storage.StorageService/GetBoxes"
)

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	CreateStorage(ctx context.Context, in *CreateStorageRequest, opts ...grpc.CallOption) (*CreateStorageResponse, error)
	AddCell(ctx context.Context, in *AddCellRequest, opts ...grpc.CallOption) (*AddCellResponse, error)
	AddBox(ctx context.Context, in *AddBoxRequest, opts ...grpc.CallOption) (*AddBoxResponse, error)
	GetStorages(ctx context.Context, in *GetStoragesRequest, opts ...grpc.CallOption) (*GetStoragesResponse, error)
	GetCells(ctx context.Context, in *GetCellsRequest, opts ...grpc.CallOption) (*GetCellsResponse, error)
	GetBoxes(ctx context.Context, in *GetBoxesRequest, opts ...grpc.CallOption) (*GetBoxesResponse, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) CreateStorage(ctx context.Context, in *CreateStorageRequest, opts ...grpc.CallOption) (*CreateStorageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateStorageResponse)
	err := c.cc.Invoke(ctx, StorageService_CreateStorage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) AddCell(ctx context.Context, in *AddCellRequest, opts ...grpc.CallOption) (*AddCellResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCellResponse)
	err := c.cc.Invoke(ctx, StorageService_AddCell_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) AddBox(ctx context.Context, in *AddBoxRequest, opts ...grpc.CallOption) (*AddBoxResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddBoxResponse)
	err := c.cc.Invoke(ctx, StorageService_AddBox_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetStorages(ctx context.Context, in *GetStoragesRequest, opts ...grpc.CallOption) (*GetStoragesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStoragesResponse)
	err := c.cc.Invoke(ctx, StorageService_GetStorages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetCells(ctx context.Context, in *GetCellsRequest, opts ...grpc.CallOption) (*GetCellsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCellsResponse)
	err := c.cc.Invoke(ctx, StorageService_GetCells_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetBoxes(ctx context.Context, in *GetBoxesRequest, opts ...grpc.CallOption) (*GetBoxesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBoxesResponse)
	err := c.cc.Invoke(ctx, StorageService_GetBoxes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations must embed UnimplementedStorageServiceServer
// for forward compatibility.
type StorageServiceServer interface {
	CreateStorage(context.Context, *CreateStorageRequest) (*CreateStorageResponse, error)
	AddCell(context.Context, *AddCellRequest) (*AddCellResponse, error)
	AddBox(context.Context, *AddBoxRequest) (*AddBoxResponse, error)
	GetStorages(context.Context, *GetStoragesRequest) (*GetStoragesResponse, error)
	GetCells(context.Context, *GetCellsRequest) (*GetCellsResponse, error)
	GetBoxes(context.Context, *GetBoxesRequest) (*GetBoxesResponse, error)
	mustEmbedUnimplementedStorageServiceServer()
}

// UnimplementedStorageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStorageServiceServer struct{}

func (UnimplementedStorageServiceServer) CreateStorage(context.Context, *CreateStorageRequest) (*CreateStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorage not implemented")
}
func (UnimplementedStorageServiceServer) AddCell(context.Context, *AddCellRequest) (*AddCellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCell not implemented")
}
func (UnimplementedStorageServiceServer) AddBox(context.Context, *AddBoxRequest) (*AddBoxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBox not implemented")
}
func (UnimplementedStorageServiceServer) GetStorages(context.Context, *GetStoragesRequest) (*GetStoragesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorages not implemented")
}
func (UnimplementedStorageServiceServer) GetCells(context.Context, *GetCellsRequest) (*GetCellsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCells not implemented")
}
func (UnimplementedStorageServiceServer) GetBoxes(context.Context, *GetBoxesRequest) (*GetBoxesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoxes not implemented")
}
func (UnimplementedStorageServiceServer) mustEmbedUnimplementedStorageServiceServer() {}
func (UnimplementedStorageServiceServer) testEmbeddedByValue()                        {}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	// If the following call pancis, it indicates UnimplementedStorageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_CreateStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).CreateStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageService_CreateStorage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).CreateStorage(ctx, req.(*CreateStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_AddCell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).AddCell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageService_AddCell_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).AddCell(ctx, req.(*AddCellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_AddBox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBoxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).AddBox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageService_AddBox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).AddBox(ctx, req.(*AddBoxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetStorages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoragesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetStorages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageService_GetStorages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetStorages(ctx, req.(*GetStoragesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetCells_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCellsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetCells(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageService_GetCells_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetCells(ctx, req.(*GetCellsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetBoxes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoxesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetBoxes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageService_GetBoxes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetBoxes(ctx, req.(*GetBoxesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStorage",
			Handler:    _StorageService_CreateStorage_Handler,
		},
		{
			MethodName: "AddCell",
			Handler:    _StorageService_AddCell_Handler,
		},
		{
			MethodName: "AddBox",
			Handler:    _StorageService_AddBox_Handler,
		},
		{
			MethodName: "GetStorages",
			Handler:    _StorageService_GetStorages_Handler,
		},
		{
			MethodName: "GetCells",
			Handler:    _StorageService_GetCells_Handler,
		},
		{
			MethodName: "GetBoxes",
			Handler:    _StorageService_GetBoxes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}
