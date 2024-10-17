// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: client/x/evmengine/keeper/evmengine.proto

package keeper

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ExecutionHead defines the execution chain head.
// It is a singleton table; it only has a single row with ID==1.
type ExecutionHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // Auto-incremented ID (always and only 1).
	CreatedHeight uint64 `protobuf:"varint,2,opt,name=created_height,json=createdHeight,proto3" json:"created_height,omitempty"` // Consensus chain height this execution block was created in.
	BlockHeight   uint64 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`       // Execution block height.
	BlockHash     []byte `protobuf:"bytes,4,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`              // Execution block hash.
	BlockTime     uint64 `protobuf:"varint,5,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`             // Execution block time.
}

func (x *ExecutionHead) Reset() {
	*x = ExecutionHead{}
	mi := &file_client_x_evmengine_keeper_evmengine_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutionHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionHead) ProtoMessage() {}

func (x *ExecutionHead) ProtoReflect() protoreflect.Message {
	mi := &file_client_x_evmengine_keeper_evmengine_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionHead.ProtoReflect.Descriptor instead.
func (*ExecutionHead) Descriptor() ([]byte, []int) {
	return file_client_x_evmengine_keeper_evmengine_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutionHead) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExecutionHead) GetCreatedHeight() uint64 {
	if x != nil {
		return x.CreatedHeight
	}
	return 0
}

func (x *ExecutionHead) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *ExecutionHead) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *ExecutionHead) GetBlockTime() uint64 {
	if x != nil {
		return x.BlockTime
	}
	return 0
}

var File_client_x_evmengine_keeper_evmengine_proto protoreflect.FileDescriptor

var file_client_x_evmengine_keeper_evmengine_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x78, 0x2f, 0x65, 0x76, 0x6d, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x65, 0x76, 0x6d, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x78, 0x2e, 0x65, 0x76, 0x6d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x1a, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f,
	0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb9, 0x01, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x10, 0xf2, 0x9e, 0xd3, 0x8e, 0x03,
	0x0a, 0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x18, 0x01, 0x42, 0xeb, 0x01, 0x0a, 0x1d,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x78, 0x2e, 0x65, 0x76, 0x6d,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x42, 0x0e, 0x45,
	0x76, 0x6d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x70, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x78, 0x2f, 0x65, 0x76, 0x6d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0xa2, 0x02, 0x04, 0x43, 0x58, 0x45, 0x4b, 0xaa, 0x02, 0x19, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x58, 0x2e, 0x45, 0x76, 0x6d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0xca, 0x02, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5c,
	0x58, 0x5c, 0x45, 0x76, 0x6d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x4b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0xe2, 0x02, 0x25, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5c, 0x58, 0x5c, 0x45, 0x76,
	0x6d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x58, 0x3a, 0x3a, 0x45, 0x76, 0x6d, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x3a, 0x3a, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_client_x_evmengine_keeper_evmengine_proto_rawDescOnce sync.Once
	file_client_x_evmengine_keeper_evmengine_proto_rawDescData = file_client_x_evmengine_keeper_evmengine_proto_rawDesc
)

func file_client_x_evmengine_keeper_evmengine_proto_rawDescGZIP() []byte {
	file_client_x_evmengine_keeper_evmengine_proto_rawDescOnce.Do(func() {
		file_client_x_evmengine_keeper_evmengine_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_x_evmengine_keeper_evmengine_proto_rawDescData)
	})
	return file_client_x_evmengine_keeper_evmengine_proto_rawDescData
}

var file_client_x_evmengine_keeper_evmengine_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_client_x_evmengine_keeper_evmengine_proto_goTypes = []any{
	(*ExecutionHead)(nil), // 0: client.x.evmengine.keeper.ExecutionHead
}
var file_client_x_evmengine_keeper_evmengine_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_x_evmengine_keeper_evmengine_proto_init() }
func file_client_x_evmengine_keeper_evmengine_proto_init() {
	if File_client_x_evmengine_keeper_evmengine_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_x_evmengine_keeper_evmengine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_x_evmengine_keeper_evmengine_proto_goTypes,
		DependencyIndexes: file_client_x_evmengine_keeper_evmengine_proto_depIdxs,
		MessageInfos:      file_client_x_evmengine_keeper_evmengine_proto_msgTypes,
	}.Build()
	File_client_x_evmengine_keeper_evmengine_proto = out.File
	file_client_x_evmengine_keeper_evmengine_proto_rawDesc = nil
	file_client_x_evmengine_keeper_evmengine_proto_goTypes = nil
	file_client_x_evmengine_keeper_evmengine_proto_depIdxs = nil
}
