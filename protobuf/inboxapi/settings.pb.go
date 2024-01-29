// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: inboxapi/settings.proto

package inboxapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddPushTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AddPushTokenRequest) Reset() {
	*x = AddPushTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPushTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPushTokenRequest) ProtoMessage() {}

func (x *AddPushTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPushTokenRequest.ProtoReflect.Descriptor instead.
func (*AddPushTokenRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_settings_proto_rawDescGZIP(), []int{0}
}

func (x *AddPushTokenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddPushTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RemovePushTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RemovePushTokenRequest) Reset() {
	*x = RemovePushTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePushTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePushTokenRequest) ProtoMessage() {}

func (x *RemovePushTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePushTokenRequest.ProtoReflect.Descriptor instead.
func (*RemovePushTokenRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_settings_proto_rawDescGZIP(), []int{1}
}

func (x *RemovePushTokenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type PushTokenExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *PushTokenExistsRequest) Reset() {
	*x = PushTokenExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTokenExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTokenExistsRequest) ProtoMessage() {}

func (x *PushTokenExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTokenExistsRequest.ProtoReflect.Descriptor instead.
func (*PushTokenExistsRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_settings_proto_rawDescGZIP(), []int{2}
}

func (x *PushTokenExistsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type PushTokenExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *PushTokenExistsResponse) Reset() {
	*x = PushTokenExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTokenExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTokenExistsResponse) ProtoMessage() {}

func (x *PushTokenExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTokenExistsResponse.ProtoReflect.Descriptor instead.
func (*PushTokenExistsResponse) Descriptor() ([]byte, []int) {
	return file_inboxapi_settings_proto_rawDescGZIP(), []int{3}
}

func (x *PushTokenExistsResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type GetPushTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetPushTokenRequest) Reset() {
	*x = GetPushTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_settings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPushTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPushTokenRequest) ProtoMessage() {}

func (x *GetPushTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_settings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPushTokenRequest.ProtoReflect.Descriptor instead.
func (*GetPushTokenRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_settings_proto_rawDescGZIP(), []int{4}
}

func (x *GetPushTokenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type PushTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PushTokenResponse) Reset() {
	*x = PushTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_settings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTokenResponse) ProtoMessage() {}

func (x *PushTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_settings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTokenResponse.ProtoReflect.Descriptor instead.
func (*PushTokenResponse) Descriptor() ([]byte, []int) {
	return file_inboxapi_settings_proto_rawDescGZIP(), []int{5}
}

func (x *PushTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_inboxapi_settings_proto protoreflect.FileDescriptor

var file_inboxapi_settings_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6e, 0x62, 0x6f, 0x78,
	0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x44, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x31, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x16, 0x50, 0x75, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x17,
	0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22,
	0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x29, 0x0a, 0x11, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc2, 0x02, 0x0a, 0x08, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x45, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x50, 0x75,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4b,
	0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x20, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x56, 0x0a, 0x0f, 0x50,
	0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x20,
	0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inboxapi_settings_proto_rawDescOnce sync.Once
	file_inboxapi_settings_proto_rawDescData = file_inboxapi_settings_proto_rawDesc
)

func file_inboxapi_settings_proto_rawDescGZIP() []byte {
	file_inboxapi_settings_proto_rawDescOnce.Do(func() {
		file_inboxapi_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_inboxapi_settings_proto_rawDescData)
	})
	return file_inboxapi_settings_proto_rawDescData
}

var file_inboxapi_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_inboxapi_settings_proto_goTypes = []interface{}{
	(*AddPushTokenRequest)(nil),     // 0: inboxapi.AddPushTokenRequest
	(*RemovePushTokenRequest)(nil),  // 1: inboxapi.RemovePushTokenRequest
	(*PushTokenExistsRequest)(nil),  // 2: inboxapi.PushTokenExistsRequest
	(*PushTokenExistsResponse)(nil), // 3: inboxapi.PushTokenExistsResponse
	(*GetPushTokenRequest)(nil),     // 4: inboxapi.GetPushTokenRequest
	(*PushTokenResponse)(nil),       // 5: inboxapi.PushTokenResponse
	(*emptypb.Empty)(nil),           // 6: google.protobuf.Empty
}
var file_inboxapi_settings_proto_depIdxs = []int32{
	0, // 0: inboxapi.Settings.AddPushToken:input_type -> inboxapi.AddPushTokenRequest
	1, // 1: inboxapi.Settings.RemovePushToken:input_type -> inboxapi.RemovePushTokenRequest
	2, // 2: inboxapi.Settings.PushTokenExists:input_type -> inboxapi.PushTokenExistsRequest
	4, // 3: inboxapi.Settings.GetPushToken:input_type -> inboxapi.GetPushTokenRequest
	6, // 4: inboxapi.Settings.AddPushToken:output_type -> google.protobuf.Empty
	6, // 5: inboxapi.Settings.RemovePushToken:output_type -> google.protobuf.Empty
	3, // 6: inboxapi.Settings.PushTokenExists:output_type -> inboxapi.PushTokenExistsResponse
	5, // 7: inboxapi.Settings.GetPushToken:output_type -> inboxapi.PushTokenResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_inboxapi_settings_proto_init() }
func file_inboxapi_settings_proto_init() {
	if File_inboxapi_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inboxapi_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPushTokenRequest); i {
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
		file_inboxapi_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePushTokenRequest); i {
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
		file_inboxapi_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTokenExistsRequest); i {
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
		file_inboxapi_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTokenExistsResponse); i {
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
		file_inboxapi_settings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPushTokenRequest); i {
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
		file_inboxapi_settings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTokenResponse); i {
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
			RawDescriptor: file_inboxapi_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inboxapi_settings_proto_goTypes,
		DependencyIndexes: file_inboxapi_settings_proto_depIdxs,
		MessageInfos:      file_inboxapi_settings_proto_msgTypes,
	}.Build()
	File_inboxapi_settings_proto = out.File
	file_inboxapi_settings_proto_rawDesc = nil
	file_inboxapi_settings_proto_goTypes = nil
	file_inboxapi_settings_proto_depIdxs = nil
}
