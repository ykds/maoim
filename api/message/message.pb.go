// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: api/message/message.proto

package message

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

type AckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	MsgId    []string `protobuf:"bytes,3,rep,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
}

func (x *AckReq) Reset() {
	*x = AckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckReq) ProtoMessage() {}

func (x *AckReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_message_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckReq.ProtoReflect.Descriptor instead.
func (*AckReq) Descriptor() ([]byte, []int) {
	return file_api_message_message_proto_rawDescGZIP(), []int{0}
}

func (x *AckReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AckReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AckReq) GetMsgId() []string {
	if x != nil {
		return x.MsgId
	}
	return nil
}

type AckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AckReply) Reset() {
	*x = AckReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckReply) ProtoMessage() {}

func (x *AckReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_message_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckReply.ProtoReflect.Descriptor instead.
func (*AckReply) Descriptor() ([]byte, []int) {
	return file_api_message_message_proto_rawDescGZIP(), []int{1}
}

var File_api_message_message_proto protoreflect.FileDescriptor

var file_api_message_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x61, 0x6f,
	0x69, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x54, 0x0a, 0x06, 0x41, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64,
	0x22, 0x0a, 0x0a, 0x08, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x45, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x41, 0x63, 0x6b, 0x4d, 0x73,
	0x67, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x6f, 0x69, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x6f, 0x69, 0x6d,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_message_message_proto_rawDescOnce sync.Once
	file_api_message_message_proto_rawDescData = file_api_message_message_proto_rawDesc
)

func file_api_message_message_proto_rawDescGZIP() []byte {
	file_api_message_message_proto_rawDescOnce.Do(func() {
		file_api_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_message_message_proto_rawDescData)
	})
	return file_api_message_message_proto_rawDescData
}

var file_api_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_message_message_proto_goTypes = []interface{}{
	(*AckReq)(nil),   // 0: maoim.message.AckReq
	(*AckReply)(nil), // 1: maoim.message.AckReply
}
var file_api_message_message_proto_depIdxs = []int32{
	0, // 0: maoim.message.Message.AckMsg:input_type -> maoim.message.AckReq
	1, // 1: maoim.message.Message.AckMsg:output_type -> maoim.message.AckReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_message_message_proto_init() }
func file_api_message_message_proto_init() {
	if File_api_message_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_message_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckReq); i {
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
		file_api_message_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckReply); i {
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
			RawDescriptor: file_api_message_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_message_message_proto_goTypes,
		DependencyIndexes: file_api_message_message_proto_depIdxs,
		MessageInfos:      file_api_message_message_proto_msgTypes,
	}.Build()
	File_api_message_message_proto = out.File
	file_api_message_message_proto_rawDesc = nil
	file_api_message_message_proto_goTypes = nil
	file_api_message_message_proto_depIdxs = nil
}
