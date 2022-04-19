// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: api/comet/comet.proto

package comet

import (
	protocal "maoim/api/protocal"
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

type PushMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys  []string        `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Proto *protocal.Proto `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
}

func (x *PushMsgReq) Reset() {
	*x = PushMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReq) ProtoMessage() {}

func (x *PushMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReq.ProtoReflect.Descriptor instead.
func (*PushMsgReq) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{0}
}

func (x *PushMsgReq) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *PushMsgReq) GetProto() *protocal.Proto {
	if x != nil {
		return x.Proto
	}
	return nil
}

type PushMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushMsgReply) Reset() {
	*x = PushMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReply) ProtoMessage() {}

func (x *PushMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReply.ProtoReflect.Descriptor instead.
func (*PushMsgReply) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{1}
}

type NewFriendShipApplyNoticeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *NewFriendShipApplyNoticeReq) Reset() {
	*x = NewFriendShipApplyNoticeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewFriendShipApplyNoticeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewFriendShipApplyNoticeReq) ProtoMessage() {}

func (x *NewFriendShipApplyNoticeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewFriendShipApplyNoticeReq.ProtoReflect.Descriptor instead.
func (*NewFriendShipApplyNoticeReq) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{2}
}

func (x *NewFriendShipApplyNoticeReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type NewFriendShipApplyNoticeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewFriendShipApplyNoticeReply) Reset() {
	*x = NewFriendShipApplyNoticeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewFriendShipApplyNoticeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewFriendShipApplyNoticeReply) ProtoMessage() {}

func (x *NewFriendShipApplyNoticeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewFriendShipApplyNoticeReply.ProtoReflect.Descriptor instead.
func (*NewFriendShipApplyNoticeReply) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{3}
}

type FriendShipApplyPassReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FriendShipApplyPassReq) Reset() {
	*x = FriendShipApplyPassReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendShipApplyPassReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendShipApplyPassReq) ProtoMessage() {}

func (x *FriendShipApplyPassReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendShipApplyPassReq.ProtoReflect.Descriptor instead.
func (*FriendShipApplyPassReq) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{4}
}

func (x *FriendShipApplyPassReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type FriendShipApplyPassReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FriendShipApplyPassReply) Reset() {
	*x = FriendShipApplyPassReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendShipApplyPassReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendShipApplyPassReply) ProtoMessage() {}

func (x *FriendShipApplyPassReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendShipApplyPassReply.ProtoReflect.Descriptor instead.
func (*FriendShipApplyPassReply) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{5}
}

var File_api_comet_comet_proto protoreflect.FileDescriptor

var file_api_comet_comet_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x6f, 0x69, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x65, 0x74, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4d, 0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x61, 0x6f, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x36, 0x0a, 0x1b, 0x4e, 0x65, 0x77, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x69,
	0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x4e, 0x65, 0x77, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x69, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x31, 0x0a, 0x16, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x53, 0x68, 0x69, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x18,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x69, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x50,
	0x61, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa7, 0x02, 0x0a, 0x05, 0x43, 0x6f, 0x6d,
	0x65, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x2e,
	0x6d, 0x61, 0x6f, 0x69, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x6f, 0x69, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x72, 0x0a, 0x18, 0x4e, 0x65, 0x77, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x53, 0x68, 0x69, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12,
	0x28, 0x2e, 0x6d, 0x61, 0x6f, 0x69, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x4e, 0x65,
	0x77, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x69, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x6d, 0x61, 0x6f, 0x69,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x53, 0x68, 0x69, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x19, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x53, 0x68, 0x69, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x61, 0x73, 0x73, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x6d, 0x61, 0x6f, 0x69, 0x6d, 0x2e, 0x63, 0x6f, 0x6d,
	0x65, 0x74, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x69, 0x70, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x6d, 0x61, 0x6f, 0x69,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x68,
	0x69, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x3b,
	0x63, 0x6f, 0x6d, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_comet_comet_proto_rawDescOnce sync.Once
	file_api_comet_comet_proto_rawDescData = file_api_comet_comet_proto_rawDesc
)

func file_api_comet_comet_proto_rawDescGZIP() []byte {
	file_api_comet_comet_proto_rawDescOnce.Do(func() {
		file_api_comet_comet_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_comet_comet_proto_rawDescData)
	})
	return file_api_comet_comet_proto_rawDescData
}

var file_api_comet_comet_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_comet_comet_proto_goTypes = []interface{}{
	(*PushMsgReq)(nil),                    // 0: maoim.comet.PushMsgReq
	(*PushMsgReply)(nil),                  // 1: maoim.comet.PushMsgReply
	(*NewFriendShipApplyNoticeReq)(nil),   // 2: maoim.comet.NewFriendShipApplyNoticeReq
	(*NewFriendShipApplyNoticeReply)(nil), // 3: maoim.comet.NewFriendShipApplyNoticeReply
	(*FriendShipApplyPassReq)(nil),        // 4: maoim.comet.FriendShipApplyPassReq
	(*FriendShipApplyPassReply)(nil),      // 5: maoim.comet.FriendShipApplyPassReply
	(*protocal.Proto)(nil),                // 6: maoim.protocal.Proto
}
var file_api_comet_comet_proto_depIdxs = []int32{
	6, // 0: maoim.comet.PushMsgReq.proto:type_name -> maoim.protocal.Proto
	0, // 1: maoim.comet.Comet.PushMsg:input_type -> maoim.comet.PushMsgReq
	2, // 2: maoim.comet.Comet.NewFriendShipApplyNotice:input_type -> maoim.comet.NewFriendShipApplyNoticeReq
	4, // 3: maoim.comet.Comet.FriendShipApplyPassNotice:input_type -> maoim.comet.FriendShipApplyPassReq
	1, // 4: maoim.comet.Comet.PushMsg:output_type -> maoim.comet.PushMsgReply
	3, // 5: maoim.comet.Comet.NewFriendShipApplyNotice:output_type -> maoim.comet.NewFriendShipApplyNoticeReply
	5, // 6: maoim.comet.Comet.FriendShipApplyPassNotice:output_type -> maoim.comet.FriendShipApplyPassReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_comet_comet_proto_init() }
func file_api_comet_comet_proto_init() {
	if File_api_comet_comet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_comet_comet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgReq); i {
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
		file_api_comet_comet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgReply); i {
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
		file_api_comet_comet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewFriendShipApplyNoticeReq); i {
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
		file_api_comet_comet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewFriendShipApplyNoticeReply); i {
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
		file_api_comet_comet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendShipApplyPassReq); i {
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
		file_api_comet_comet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendShipApplyPassReply); i {
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
			RawDescriptor: file_api_comet_comet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_comet_comet_proto_goTypes,
		DependencyIndexes: file_api_comet_comet_proto_depIdxs,
		MessageInfos:      file_api_comet_comet_proto_msgTypes,
	}.Build()
	File_api_comet_comet_proto = out.File
	file_api_comet_comet_proto_rawDesc = nil
	file_api_comet_comet_proto_goTypes = nil
	file_api_comet_comet_proto_depIdxs = nil
}
