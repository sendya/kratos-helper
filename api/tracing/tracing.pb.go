// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: tracing/tracing.proto

package tracing

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

type Tracing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint   string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	CustomName string `protobuf:"bytes,2,opt,name=custom_name,json=customName,proto3" json:"custom_name,omitempty"`
}

func (x *Tracing) Reset() {
	*x = Tracing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracing_tracing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tracing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tracing) ProtoMessage() {}

func (x *Tracing) ProtoReflect() protoreflect.Message {
	mi := &file_tracing_tracing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tracing.ProtoReflect.Descriptor instead.
func (*Tracing) Descriptor() ([]byte, []int) {
	return file_tracing_tracing_proto_rawDescGZIP(), []int{0}
}

func (x *Tracing) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Tracing) GetCustomName() string {
	if x != nil {
		return x.CustomName
	}
	return ""
}

var File_tracing_tracing_proto protoreflect.FileDescriptor

var file_tracing_tracing_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22, 0x46, 0x0a, 0x07, 0x54, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x42, 0x6d, 0x0a, 0x1c, 0x73, 0x65, 0x6e, 0x64, 0x79, 0x61, 0x2e, 0x6d, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x42, 0x0e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x56, 0x31, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x65, 0x6e, 0x64, 0x79, 0x61, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x68,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x3b, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracing_tracing_proto_rawDescOnce sync.Once
	file_tracing_tracing_proto_rawDescData = file_tracing_tracing_proto_rawDesc
)

func file_tracing_tracing_proto_rawDescGZIP() []byte {
	file_tracing_tracing_proto_rawDescOnce.Do(func() {
		file_tracing_tracing_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracing_tracing_proto_rawDescData)
	})
	return file_tracing_tracing_proto_rawDescData
}

var file_tracing_tracing_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tracing_tracing_proto_goTypes = []interface{}{
	(*Tracing)(nil), // 0: protos.tracing.v1.Tracing
}
var file_tracing_tracing_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tracing_tracing_proto_init() }
func file_tracing_tracing_proto_init() {
	if File_tracing_tracing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracing_tracing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tracing); i {
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
			RawDescriptor: file_tracing_tracing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tracing_tracing_proto_goTypes,
		DependencyIndexes: file_tracing_tracing_proto_depIdxs,
		MessageInfos:      file_tracing_tracing_proto_msgTypes,
	}.Build()
	File_tracing_tracing_proto = out.File
	file_tracing_tracing_proto_rawDesc = nil
	file_tracing_tracing_proto_goTypes = nil
	file_tracing_tracing_proto_depIdxs = nil
}
