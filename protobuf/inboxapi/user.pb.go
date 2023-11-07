// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: inboxapi/user.proto

package inboxapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RecentlyViewedType int32

const (
	RecentlyViewedType_RECENTLY_VIEWED_TYPE_UNSPECIFIED RecentlyViewedType = 0
	RecentlyViewedType_RECENTLY_VIEWED_TYPE_DAO         RecentlyViewedType = 1
)

// Enum value maps for RecentlyViewedType.
var (
	RecentlyViewedType_name = map[int32]string{
		0: "RECENTLY_VIEWED_TYPE_UNSPECIFIED",
		1: "RECENTLY_VIEWED_TYPE_DAO",
	}
	RecentlyViewedType_value = map[string]int32{
		"RECENTLY_VIEWED_TYPE_UNSPECIFIED": 0,
		"RECENTLY_VIEWED_TYPE_DAO":         1,
	}
)

func (x RecentlyViewedType) Enum() *RecentlyViewedType {
	p := new(RecentlyViewedType)
	*p = x
	return p
}

func (x RecentlyViewedType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecentlyViewedType) Descriptor() protoreflect.EnumDescriptor {
	return file_inboxapi_user_proto_enumTypes[0].Descriptor()
}

func (RecentlyViewedType) Type() protoreflect.EnumType {
	return &file_inboxapi_user_proto_enumTypes[0]
}

func (x RecentlyViewedType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecentlyViewedType.Descriptor instead.
func (RecentlyViewedType) EnumDescriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{0}
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserID) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserByIDRequest) Reset() {
	*x = UserByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserByIDRequest) ProtoMessage() {}

func (x *UserByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserByIDRequest.ProtoReflect.Descriptor instead.
func (*UserByIDRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserByIDRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeviceUuid string                 `protobuf:"bytes,4,opt,name=device_uuid,json=deviceUuid,proto3" json:"device_uuid,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UserInfo) GetDeviceUuid() string {
	if x != nil {
		return x.DeviceUuid
	}
	return ""
}

type UserByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserInfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserByIDResponse) Reset() {
	*x = UserByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserByIDResponse) ProtoMessage() {}

func (x *UserByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserByIDResponse.ProtoReflect.Descriptor instead.
func (*UserByIDResponse) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserByIDResponse) GetUser() *UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

type UserByUuidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceUuid string `protobuf:"bytes,1,opt,name=device_uuid,json=deviceUuid,proto3" json:"device_uuid,omitempty"`
}

func (x *UserByUuidRequest) Reset() {
	*x = UserByUuidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserByUuidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserByUuidRequest) ProtoMessage() {}

func (x *UserByUuidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserByUuidRequest.ProtoReflect.Descriptor instead.
func (*UserByUuidRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserByUuidRequest) GetDeviceUuid() string {
	if x != nil {
		return x.DeviceUuid
	}
	return ""
}

type UserByUuidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserInfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserByUuidResponse) Reset() {
	*x = UserByUuidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserByUuidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserByUuidResponse) ProtoMessage() {}

func (x *UserByUuidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserByUuidResponse.ProtoReflect.Descriptor instead.
func (*UserByUuidResponse) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserByUuidResponse) GetUser() *UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

type UserCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceUuid string `protobuf:"bytes,1,opt,name=device_uuid,json=deviceUuid,proto3" json:"device_uuid,omitempty"`
}

func (x *UserCreateRequest) Reset() {
	*x = UserCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateRequest) ProtoMessage() {}

func (x *UserCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateRequest.ProtoReflect.Descriptor instead.
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserCreateRequest) GetDeviceUuid() string {
	if x != nil {
		return x.DeviceUuid
	}
	return ""
}

type UserCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserInfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserCreateResponse) Reset() {
	*x = UserCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateResponse) ProtoMessage() {}

func (x *UserCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateResponse.ProtoReflect.Descriptor instead.
func (*UserCreateResponse) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserCreateResponse) GetUser() *UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

type UserViewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string             `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Type   RecentlyViewedType `protobuf:"varint,2,opt,name=type,proto3,enum=inboxapi.RecentlyViewedType" json:"type,omitempty"`
	TypeId string             `protobuf:"bytes,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
}

func (x *UserViewRequest) Reset() {
	*x = UserViewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserViewRequest) ProtoMessage() {}

func (x *UserViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserViewRequest.ProtoReflect.Descriptor instead.
func (*UserViewRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserViewRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserViewRequest) GetType() RecentlyViewedType {
	if x != nil {
		return x.Type
	}
	return RecentlyViewedType_RECENTLY_VIEWED_TYPE_UNSPECIFIED
}

func (x *UserViewRequest) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

type UserLastViewedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string             `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Type   RecentlyViewedType `protobuf:"varint,2,opt,name=type,proto3,enum=inboxapi.RecentlyViewedType" json:"type,omitempty"`
	Limit  uint64             `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *UserLastViewedRequest) Reset() {
	*x = UserLastViewedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLastViewedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLastViewedRequest) ProtoMessage() {}

func (x *UserLastViewedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLastViewedRequest.ProtoReflect.Descriptor instead.
func (*UserLastViewedRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{9}
}

func (x *UserLastViewedRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserLastViewedRequest) GetType() RecentlyViewedType {
	if x != nil {
		return x.Type
	}
	return RecentlyViewedType_RECENTLY_VIEWED_TYPE_UNSPECIFIED
}

func (x *UserLastViewedRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type RecentlyViewed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Type      RecentlyViewedType     `protobuf:"varint,2,opt,name=type,proto3,enum=inboxapi.RecentlyViewedType" json:"type,omitempty"`
	TypeId    string                 `protobuf:"bytes,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
}

func (x *RecentlyViewed) Reset() {
	*x = RecentlyViewed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecentlyViewed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecentlyViewed) ProtoMessage() {}

func (x *RecentlyViewed) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecentlyViewed.ProtoReflect.Descriptor instead.
func (*RecentlyViewed) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{10}
}

func (x *RecentlyViewed) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RecentlyViewed) GetType() RecentlyViewedType {
	if x != nil {
		return x.Type
	}
	return RecentlyViewedType_RECENTLY_VIEWED_TYPE_UNSPECIFIED
}

func (x *RecentlyViewed) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

type UserLastViewedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*RecentlyViewed `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UserLastViewedResponse) Reset() {
	*x = UserLastViewedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLastViewedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLastViewedResponse) ProtoMessage() {}

func (x *UserLastViewedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLastViewedResponse.ProtoReflect.Descriptor instead.
func (*UserLastViewedResponse) Descriptor() ([]byte, []int) {
	return file_inboxapi_user_proto_rawDescGZIP(), []int{11}
}

func (x *UserLastViewedResponse) GetList() []*RecentlyViewed {
	if x != nil {
		return x.List
	}
	return nil
}

var File_inboxapi_user_proto protoreflect.FileDescriptor

var file_inboxapi_user_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x2a, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb1, 0x01, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x75, 0x69, 0x64,
	0x22, 0x3a, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x11,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x75,
	0x69, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x34, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x55, 0x75, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x6e, 0x62,
	0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x75, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x6c, 0x79, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x15, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6e,
	0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x56,
	0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x6c, 0x79, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x63,
	0x65, 0x6e, 0x74, 0x6c, 0x79, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x46,
	0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x2a, 0x58, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x6c, 0x79, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x20,
	0x52, 0x45, 0x43, 0x45, 0x4e, 0x54, 0x4c, 0x59, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x45, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x43, 0x45, 0x4e, 0x54, 0x4c, 0x59, 0x5f, 0x56,
	0x49, 0x45, 0x57, 0x45, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x4f, 0x10, 0x01,
	0x32, 0xe4, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x19, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e,
	0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x62,
	0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x19, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4f, 0x0a, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x69,
	0x65, 0x77, 0x65, 0x64, 0x12, 0x1f, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x69, 0x6e, 0x62,
	0x6f, 0x78, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inboxapi_user_proto_rawDescOnce sync.Once
	file_inboxapi_user_proto_rawDescData = file_inboxapi_user_proto_rawDesc
)

func file_inboxapi_user_proto_rawDescGZIP() []byte {
	file_inboxapi_user_proto_rawDescOnce.Do(func() {
		file_inboxapi_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_inboxapi_user_proto_rawDescData)
	})
	return file_inboxapi_user_proto_rawDescData
}

var file_inboxapi_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_inboxapi_user_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_inboxapi_user_proto_goTypes = []interface{}{
	(RecentlyViewedType)(0),        // 0: inboxapi.RecentlyViewedType
	(*UserID)(nil),                 // 1: inboxapi.UserID
	(*UserByIDRequest)(nil),        // 2: inboxapi.UserByIDRequest
	(*UserInfo)(nil),               // 3: inboxapi.UserInfo
	(*UserByIDResponse)(nil),       // 4: inboxapi.UserByIDResponse
	(*UserByUuidRequest)(nil),      // 5: inboxapi.UserByUuidRequest
	(*UserByUuidResponse)(nil),     // 6: inboxapi.UserByUuidResponse
	(*UserCreateRequest)(nil),      // 7: inboxapi.UserCreateRequest
	(*UserCreateResponse)(nil),     // 8: inboxapi.UserCreateResponse
	(*UserViewRequest)(nil),        // 9: inboxapi.UserViewRequest
	(*UserLastViewedRequest)(nil),  // 10: inboxapi.UserLastViewedRequest
	(*RecentlyViewed)(nil),         // 11: inboxapi.RecentlyViewed
	(*UserLastViewedResponse)(nil), // 12: inboxapi.UserLastViewedResponse
	(*timestamppb.Timestamp)(nil),  // 13: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 14: google.protobuf.Empty
}
var file_inboxapi_user_proto_depIdxs = []int32{
	13, // 0: inboxapi.UserInfo.created_at:type_name -> google.protobuf.Timestamp
	13, // 1: inboxapi.UserInfo.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 2: inboxapi.UserByIDResponse.user:type_name -> inboxapi.UserInfo
	3,  // 3: inboxapi.UserByUuidResponse.user:type_name -> inboxapi.UserInfo
	3,  // 4: inboxapi.UserCreateResponse.user:type_name -> inboxapi.UserInfo
	0,  // 5: inboxapi.UserViewRequest.type:type_name -> inboxapi.RecentlyViewedType
	0,  // 6: inboxapi.UserLastViewedRequest.type:type_name -> inboxapi.RecentlyViewedType
	13, // 7: inboxapi.RecentlyViewed.created_at:type_name -> google.protobuf.Timestamp
	0,  // 8: inboxapi.RecentlyViewed.type:type_name -> inboxapi.RecentlyViewedType
	11, // 9: inboxapi.UserLastViewedResponse.list:type_name -> inboxapi.RecentlyViewed
	2,  // 10: inboxapi.User.GetByID:input_type -> inboxapi.UserByIDRequest
	5,  // 11: inboxapi.User.GetByUuid:input_type -> inboxapi.UserByUuidRequest
	7,  // 12: inboxapi.User.Create:input_type -> inboxapi.UserCreateRequest
	9,  // 13: inboxapi.User.AddView:input_type -> inboxapi.UserViewRequest
	10, // 14: inboxapi.User.LastViewed:input_type -> inboxapi.UserLastViewedRequest
	4,  // 15: inboxapi.User.GetByID:output_type -> inboxapi.UserByIDResponse
	6,  // 16: inboxapi.User.GetByUuid:output_type -> inboxapi.UserByUuidResponse
	8,  // 17: inboxapi.User.Create:output_type -> inboxapi.UserCreateResponse
	14, // 18: inboxapi.User.AddView:output_type -> google.protobuf.Empty
	12, // 19: inboxapi.User.LastViewed:output_type -> inboxapi.UserLastViewedResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_inboxapi_user_proto_init() }
func file_inboxapi_user_proto_init() {
	if File_inboxapi_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inboxapi_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_inboxapi_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserByIDRequest); i {
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
		file_inboxapi_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_inboxapi_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserByIDResponse); i {
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
		file_inboxapi_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserByUuidRequest); i {
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
		file_inboxapi_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserByUuidResponse); i {
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
		file_inboxapi_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateRequest); i {
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
		file_inboxapi_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateResponse); i {
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
		file_inboxapi_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserViewRequest); i {
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
		file_inboxapi_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLastViewedRequest); i {
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
		file_inboxapi_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecentlyViewed); i {
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
		file_inboxapi_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLastViewedResponse); i {
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
			RawDescriptor: file_inboxapi_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inboxapi_user_proto_goTypes,
		DependencyIndexes: file_inboxapi_user_proto_depIdxs,
		EnumInfos:         file_inboxapi_user_proto_enumTypes,
		MessageInfos:      file_inboxapi_user_proto_msgTypes,
	}.Build()
	File_inboxapi_user_proto = out.File
	file_inboxapi_user_proto_rawDesc = nil
	file_inboxapi_user_proto_goTypes = nil
	file_inboxapi_user_proto_depIdxs = nil
}
