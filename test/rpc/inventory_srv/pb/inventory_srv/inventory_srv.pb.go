// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: inventory_srv.proto

package inventory_srv

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

// 扣减库存
type DRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid   int64 `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *DRequest) Reset() {
	*x = DRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_srv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DRequest) ProtoMessage() {}

func (x *DRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_srv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DRequest.ProtoReflect.Descriptor instead.
func (*DRequest) Descriptor() ([]byte, []int) {
	return file_inventory_srv_proto_rawDescGZIP(), []int{0}
}

func (x *DRequest) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *DRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DResponse) Reset() {
	*x = DResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_srv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DResponse) ProtoMessage() {}

func (x *DResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_srv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DResponse.ProtoReflect.Descriptor instead.
func (*DResponse) Descriptor() ([]byte, []int) {
	return file_inventory_srv_proto_rawDescGZIP(), []int{1}
}

func (x *DResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

// 增加库存
type AResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid   int64 `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AResquest) Reset() {
	*x = AResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_srv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AResquest) ProtoMessage() {}

func (x *AResquest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_srv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AResquest.ProtoReflect.Descriptor instead.
func (*AResquest) Descriptor() ([]byte, []int) {
	return file_inventory_srv_proto_rawDescGZIP(), []int{2}
}

func (x *AResquest) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *AResquest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AResponse) Reset() {
	*x = AResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_srv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AResponse) ProtoMessage() {}

func (x *AResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_srv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AResponse.ProtoReflect.Descriptor instead.
func (*AResponse) Descriptor() ([]byte, []int) {
	return file_inventory_srv_proto_rawDescGZIP(), []int{3}
}

func (x *AResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_inventory_srv_proto protoreflect.FileDescriptor

var file_inventory_srv_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x72, 0x76, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x73, 0x72, 0x76, 0x22, 0x32, 0x0a, 0x08, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x09, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x33, 0x0a,
	0x09, 0x41, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x23, 0x0a, 0x09, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x97, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x72, 0x76, 0x12, 0x42, 0x0a, 0x0d, 0x44, 0x65, 0x73,
	0x63, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x73, 0x72, 0x76, 0x2e, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x0c, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x52,
	0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x73, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_srv_proto_rawDescOnce sync.Once
	file_inventory_srv_proto_rawDescData = file_inventory_srv_proto_rawDesc
)

func file_inventory_srv_proto_rawDescGZIP() []byte {
	file_inventory_srv_proto_rawDescOnce.Do(func() {
		file_inventory_srv_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_srv_proto_rawDescData)
	})
	return file_inventory_srv_proto_rawDescData
}

var file_inventory_srv_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_inventory_srv_proto_goTypes = []interface{}{
	(*DRequest)(nil),  // 0: inventory_srv.DRequest
	(*DResponse)(nil), // 1: inventory_srv.DResponse
	(*AResquest)(nil), // 2: inventory_srv.AResquest
	(*AResponse)(nil), // 3: inventory_srv.AResponse
}
var file_inventory_srv_proto_depIdxs = []int32{
	0, // 0: inventory_srv.Inventory_srv.DescInventory:input_type -> inventory_srv.DRequest
	2, // 1: inventory_srv.Inventory_srv.AddInventory:input_type -> inventory_srv.AResquest
	1, // 2: inventory_srv.Inventory_srv.DescInventory:output_type -> inventory_srv.DResponse
	3, // 3: inventory_srv.Inventory_srv.AddInventory:output_type -> inventory_srv.AResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_inventory_srv_proto_init() }
func file_inventory_srv_proto_init() {
	if File_inventory_srv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_srv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DRequest); i {
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
		file_inventory_srv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DResponse); i {
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
		file_inventory_srv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AResquest); i {
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
		file_inventory_srv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AResponse); i {
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
			RawDescriptor: file_inventory_srv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_srv_proto_goTypes,
		DependencyIndexes: file_inventory_srv_proto_depIdxs,
		MessageInfos:      file_inventory_srv_proto_msgTypes,
	}.Build()
	File_inventory_srv_proto = out.File
	file_inventory_srv_proto_rawDesc = nil
	file_inventory_srv_proto_goTypes = nil
	file_inventory_srv_proto_depIdxs = nil
}
