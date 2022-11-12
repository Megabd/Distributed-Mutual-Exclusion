// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: grpc/interface.proto

package ping

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_grpc_interface_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_grpc_interface_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ReturnInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TimesAccessed int32 `protobuf:"varint,2,opt,name=timesAccessed,proto3" json:"timesAccessed,omitempty"`
	Wanted        bool  `protobuf:"varint,3,opt,name=wanted,proto3" json:"wanted,omitempty"`
	Held          bool  `protobuf:"varint,4,opt,name=held,proto3" json:"held,omitempty"`
}

func (x *ReturnInfoReply) Reset() {
	*x = ReturnInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnInfoReply) ProtoMessage() {}

func (x *ReturnInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnInfoReply.ProtoReflect.Descriptor instead.
func (*ReturnInfoReply) Descriptor() ([]byte, []int) {
	return file_grpc_interface_proto_rawDescGZIP(), []int{2}
}

func (x *ReturnInfoReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReturnInfoReply) GetTimesAccessed() int32 {
	if x != nil {
		return x.TimesAccessed
	}
	return 0
}

func (x *ReturnInfoReply) GetWanted() bool {
	if x != nil {
		return x.Wanted
	}
	return false
}

func (x *ReturnInfoReply) GetHeld() bool {
	if x != nil {
		return x.Held
	}
	return false
}

var File_grpc_interface_proto protoreflect.FileDescriptor

var file_grpc_interface_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x19, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x73, 0x0a, 0x0f, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x6c,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x64, 0x32, 0x3c, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0d, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x61, 0x64, 0x64, 0x69, 0x4e, 0x61, 0x64, 0x6a, 0x61, 0x2f, 0x70, 0x65, 0x65,
	0x72, 0x2d, 0x74, 0x6f, 0x2d, 0x70, 0x65, 0x65, 0x72, 0x3b, 0x70, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_interface_proto_rawDescOnce sync.Once
	file_grpc_interface_proto_rawDescData = file_grpc_interface_proto_rawDesc
)

func file_grpc_interface_proto_rawDescGZIP() []byte {
	file_grpc_interface_proto_rawDescOnce.Do(func() {
		file_grpc_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_interface_proto_rawDescData)
	})
	return file_grpc_interface_proto_rawDescData
}

var file_grpc_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_interface_proto_goTypes = []interface{}{
	(*Request)(nil),         // 0: ping.Request
	(*Reply)(nil),           // 1: ping.Reply
	(*ReturnInfoReply)(nil), // 2: ping.returnInfoReply
}
var file_grpc_interface_proto_depIdxs = []int32{
	0, // 0: ping.Ping.returnInfo:input_type -> ping.Request
	2, // 1: ping.Ping.returnInfo:output_type -> ping.returnInfoReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_interface_proto_init() }
func file_grpc_interface_proto_init() {
	if File_grpc_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_grpc_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_grpc_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnInfoReply); i {
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
			RawDescriptor: file_grpc_interface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_interface_proto_goTypes,
		DependencyIndexes: file_grpc_interface_proto_depIdxs,
		MessageInfos:      file_grpc_interface_proto_msgTypes,
	}.Build()
	File_grpc_interface_proto = out.File
	file_grpc_interface_proto_rawDesc = nil
	file_grpc_interface_proto_goTypes = nil
	file_grpc_interface_proto_depIdxs = nil
}
