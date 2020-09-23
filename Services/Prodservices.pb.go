// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: Prodservices.proto

package Services

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prodservices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_Prodservices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_Prodservices_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Rsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Test `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Rsp) Reset() {
	*x = Rsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prodservices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rsp) ProtoMessage() {}

func (x *Rsp) ProtoReflect() protoreflect.Message {
	mi := &file_Prodservices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rsp.ProtoReflect.Descriptor instead.
func (*Rsp) Descriptor() ([]byte, []int) {
	return file_Prodservices_proto_rawDescGZIP(), []int{1}
}

func (x *Rsp) GetData() []*Test {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_Prodservices_proto protoreflect.FileDescriptor

var file_Prodservices_proto_rawDesc = []byte{
	0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x0c,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x03,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x29, 0x0a, 0x03, 0x52, 0x73, 0x70, 0x12, 0x22,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x36, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x27, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x12, 0x0d, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Prodservices_proto_rawDescOnce sync.Once
	file_Prodservices_proto_rawDescData = file_Prodservices_proto_rawDesc
)

func file_Prodservices_proto_rawDescGZIP() []byte {
	file_Prodservices_proto_rawDescOnce.Do(func() {
		file_Prodservices_proto_rawDescData = protoimpl.X.CompressGZIP(file_Prodservices_proto_rawDescData)
	})
	return file_Prodservices_proto_rawDescData
}

var file_Prodservices_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Prodservices_proto_goTypes = []interface{}{
	(*Req)(nil),  // 0: Services.Req
	(*Rsp)(nil),  // 1: Services.Rsp
	(*Test)(nil), // 2: Services.Test
}
var file_Prodservices_proto_depIdxs = []int32{
	2, // 0: Services.Rsp.data:type_name -> Services.Test
	0, // 1: Services.ProdService.getProd:input_type -> Services.Req
	1, // 2: Services.ProdService.getProd:output_type -> Services.Rsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Prodservices_proto_init() }
func file_Prodservices_proto_init() {
	if File_Prodservices_proto != nil {
		return
	}
	file_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Prodservices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_Prodservices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rsp); i {
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
			RawDescriptor: file_Prodservices_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Prodservices_proto_goTypes,
		DependencyIndexes: file_Prodservices_proto_depIdxs,
		MessageInfos:      file_Prodservices_proto_msgTypes,
	}.Build()
	File_Prodservices_proto = out.File
	file_Prodservices_proto_rawDesc = nil
	file_Prodservices_proto_goTypes = nil
	file_Prodservices_proto_depIdxs = nil
}
