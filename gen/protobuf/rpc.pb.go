// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: rpc.proto

package protobuf

import (
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

type RpcType int32

const (
	RpcType_RPC_TYPE_IMAGE_PARSER     RpcType = 0
	RpcType_RPC_TYPE_LIST_VIEW_PARSER RpcType = 1
	RpcType_RPC_TYPE_GALLERY_PARSER   RpcType = 2
	RpcType_RPC_TYPE_AUTO_COMPLETE    RpcType = 3
)

// Enum value maps for RpcType.
var (
	RpcType_name = map[int32]string{
		0: "RPC_TYPE_IMAGE_PARSER",
		1: "RPC_TYPE_LIST_VIEW_PARSER",
		2: "RPC_TYPE_GALLERY_PARSER",
		3: "RPC_TYPE_AUTO_COMPLETE",
	}
	RpcType_value = map[string]int32{
		"RPC_TYPE_IMAGE_PARSER":     0,
		"RPC_TYPE_LIST_VIEW_PARSER": 1,
		"RPC_TYPE_GALLERY_PARSER":   2,
		"RPC_TYPE_AUTO_COMPLETE":    3,
	}
)

func (x RpcType) Enum() *RpcType {
	p := new(RpcType)
	*p = x
	return p
}

func (x RpcType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RpcType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[0].Descriptor()
}

func (RpcType) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[0]
}

func (x RpcType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RpcType.Descriptor instead.
func (RpcType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

type RpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       RpcType           `protobuf:"varint,1,opt,name=type,proto3,enum=RpcType" json:"type,omitempty"`
	Data       string            `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	ParserData []byte            `protobuf:"bytes,3,opt,name=parser_data,json=parserData,proto3" json:"parser_data,omitempty"`
	Env        map[string]string `protobuf:"bytes,4,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RpcRequest) Reset() {
	*x = RpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcRequest) ProtoMessage() {}

func (x *RpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcRequest.ProtoReflect.Descriptor instead.
func (*RpcRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *RpcRequest) GetType() RpcType {
	if x != nil {
		return x.Type
	}
	return RpcType_RPC_TYPE_IMAGE_PARSER
}

func (x *RpcRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *RpcRequest) GetParserData() []byte {
	if x != nil {
		return x.ParserData
	}
	return nil
}

func (x *RpcRequest) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

type RpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RpcResponse) Reset() {
	*x = RpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcResponse) ProtoMessage() {}

func (x *RpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcResponse.ProtoReflect.Descriptor instead.
func (*RpcResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *RpcResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RpcResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x0a,
	0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x52, 0x70, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a,
	0x03, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x52, 0x70, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x65, 0x6e, 0x76, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x37, 0x0a,
	0x0b, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x7c, 0x0a, 0x07, 0x52, 0x70, 0x63, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x50, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4d,
	0x41, 0x47, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19,
	0x52, 0x50, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x56, 0x49,
	0x45, 0x57, 0x5f, 0x50, 0x41, 0x52, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x52,
	0x50, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x41, 0x4c, 0x4c, 0x45, 0x52, 0x59, 0x5f,
	0x50, 0x41, 0x52, 0x53, 0x45, 0x52, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x50, 0x43, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x54, 0x45, 0x10, 0x03, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_proto_goTypes = []interface{}{
	(RpcType)(0),        // 0: RpcType
	(*RpcRequest)(nil),  // 1: RpcRequest
	(*RpcResponse)(nil), // 2: RpcResponse
	nil,                 // 3: RpcRequest.EnvEntry
}
var file_rpc_proto_depIdxs = []int32{
	0, // 0: RpcRequest.type:type_name -> RpcType
	3, // 1: RpcRequest.env:type_name -> RpcRequest.EnvEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		EnumInfos:         file_rpc_proto_enumTypes,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
