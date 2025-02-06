// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.21.12
// source: storage.proto

package storagepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateStorageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateStorageRequest) Reset() {
	*x = CreateStorageRequest{}
	mi := &file_storage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStorageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStorageRequest) ProtoMessage() {}

func (x *CreateStorageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStorageRequest.ProtoReflect.Descriptor instead.
func (*CreateStorageRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStorageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateStorageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateStorageResponse) Reset() {
	*x = CreateStorageResponse{}
	mi := &file_storage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStorageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStorageResponse) ProtoMessage() {}

func (x *CreateStorageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStorageResponse.ProtoReflect.Descriptor instead.
func (*CreateStorageResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStorageResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateStorageResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddCellRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StorageId     uint64                 `protobuf:"varint,2,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddCellRequest) Reset() {
	*x = AddCellRequest{}
	mi := &file_storage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddCellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCellRequest) ProtoMessage() {}

func (x *AddCellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCellRequest.ProtoReflect.Descriptor instead.
func (*AddCellRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{2}
}

func (x *AddCellRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddCellRequest) GetStorageId() uint64 {
	if x != nil {
		return x.StorageId
	}
	return 0
}

type AddCellResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddCellResponse) Reset() {
	*x = AddCellResponse{}
	mi := &file_storage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddCellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCellResponse) ProtoMessage() {}

func (x *AddCellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCellResponse.ProtoReflect.Descriptor instead.
func (*AddCellResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{3}
}

func (x *AddCellResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddCellResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddBoxRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CellId        uint64                 `protobuf:"varint,2,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddBoxRequest) Reset() {
	*x = AddBoxRequest{}
	mi := &file_storage_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddBoxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBoxRequest) ProtoMessage() {}

func (x *AddBoxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBoxRequest.ProtoReflect.Descriptor instead.
func (*AddBoxRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{4}
}

func (x *AddBoxRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddBoxRequest) GetCellId() uint64 {
	if x != nil {
		return x.CellId
	}
	return 0
}

type AddBoxResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddBoxResponse) Reset() {
	*x = AddBoxResponse{}
	mi := &file_storage_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddBoxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBoxResponse) ProtoMessage() {}

func (x *AddBoxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBoxResponse.ProtoReflect.Descriptor instead.
func (*AddBoxResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{5}
}

func (x *AddBoxResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddBoxResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_storage_proto protoreflect.FileDescriptor

var file_storage_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x43, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x65, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a,
	0x0d, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x32, 0xd9, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x65, 0x6c, 0x6c, 0x12,
	0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x65, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x78, 0x12, 0x16, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41,
	0x64, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a,
	0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x76, 0x61,
	0x4a, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2d, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_storage_proto_rawDescOnce sync.Once
	file_storage_proto_rawDescData []byte
)

func file_storage_proto_rawDescGZIP() []byte {
	file_storage_proto_rawDescOnce.Do(func() {
		file_storage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_proto_rawDesc), len(file_storage_proto_rawDesc)))
	})
	return file_storage_proto_rawDescData
}

var file_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_storage_proto_goTypes = []any{
	(*CreateStorageRequest)(nil),  // 0: storage.CreateStorageRequest
	(*CreateStorageResponse)(nil), // 1: storage.CreateStorageResponse
	(*AddCellRequest)(nil),        // 2: storage.AddCellRequest
	(*AddCellResponse)(nil),       // 3: storage.AddCellResponse
	(*AddBoxRequest)(nil),         // 4: storage.AddBoxRequest
	(*AddBoxResponse)(nil),        // 5: storage.AddBoxResponse
}
var file_storage_proto_depIdxs = []int32{
	0, // 0: storage.StorageService.CreateStorage:input_type -> storage.CreateStorageRequest
	2, // 1: storage.StorageService.AddCell:input_type -> storage.AddCellRequest
	4, // 2: storage.StorageService.AddBox:input_type -> storage.AddBoxRequest
	1, // 3: storage.StorageService.CreateStorage:output_type -> storage.CreateStorageResponse
	3, // 4: storage.StorageService.AddCell:output_type -> storage.AddCellResponse
	5, // 5: storage.StorageService.AddBox:output_type -> storage.AddBoxResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_storage_proto_init() }
func file_storage_proto_init() {
	if File_storage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_proto_rawDesc), len(file_storage_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_proto_goTypes,
		DependencyIndexes: file_storage_proto_depIdxs,
		MessageInfos:      file_storage_proto_msgTypes,
	}.Build()
	File_storage_proto = out.File
	file_storage_proto_goTypes = nil
	file_storage_proto_depIdxs = nil
}
