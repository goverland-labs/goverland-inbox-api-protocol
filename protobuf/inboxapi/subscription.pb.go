// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: inboxapi/subscription.proto

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

type FindSubscribersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DaoId string `protobuf:"bytes,1,opt,name=dao_id,json=daoId,proto3" json:"dao_id,omitempty"`
}

func (x *FindSubscribersRequest) Reset() {
	*x = FindSubscribersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindSubscribersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindSubscribersRequest) ProtoMessage() {}

func (x *FindSubscribersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindSubscribersRequest.ProtoReflect.Descriptor instead.
func (*FindSubscribersRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *FindSubscribersRequest) GetDaoId() string {
	if x != nil {
		return x.DaoId
	}
	return ""
}

type UserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserID `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserList) Reset() {
	*x = UserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_inboxapi_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *UserList) GetUsers() []*UserID {
	if x != nil {
		return x.Users
	}
	return nil
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId string `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	DaoId        string `protobuf:"bytes,2,opt,name=dao_id,json=daoId,proto3" json:"dao_id,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_subscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_subscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *SubscribeRequest) GetDaoId() string {
	if x != nil {
		return x.DaoId
	}
	return ""
}

type SubscriptionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string                 `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	SubscriberId   string                 `protobuf:"bytes,2,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	DaoId          string                 `protobuf:"bytes,3,opt,name=dao_id,json=daoId,proto3" json:"dao_id,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *SubscriptionInfo) Reset() {
	*x = SubscriptionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_subscription_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionInfo) ProtoMessage() {}

func (x *SubscriptionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_subscription_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionInfo.ProtoReflect.Descriptor instead.
func (*SubscriptionInfo) Descriptor() ([]byte, []int) {
	return file_inboxapi_subscription_proto_rawDescGZIP(), []int{3}
}

func (x *SubscriptionInfo) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *SubscriptionInfo) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *SubscriptionInfo) GetDaoId() string {
	if x != nil {
		return x.DaoId
	}
	return ""
}

func (x *SubscriptionInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ListSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId string  `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	Limit        *uint64 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset       *uint64 `protobuf:"varint,3,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *ListSubscriptionRequest) Reset() {
	*x = ListSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_subscription_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionRequest) ProtoMessage() {}

func (x *ListSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_subscription_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*ListSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_subscription_proto_rawDescGZIP(), []int{4}
}

func (x *ListSubscriptionRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *ListSubscriptionRequest) GetLimit() uint64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *ListSubscriptionRequest) GetOffset() uint64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

type ListSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*SubscriptionInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	TotalCount uint64              `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ListSubscriptionResponse) Reset() {
	*x = ListSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_subscription_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionResponse) ProtoMessage() {}

func (x *ListSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_subscription_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*ListSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_inboxapi_subscription_proto_rawDescGZIP(), []int{5}
}

func (x *ListSubscriptionResponse) GetItems() []*SubscriptionInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListSubscriptionResponse) GetTotalCount() uint64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_subscription_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_subscription_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_subscription_proto_rawDescGZIP(), []int{6}
}

func (x *UnsubscribeRequest) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

type GetSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *GetSubscriptionRequest) Reset() {
	*x = GetSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_subscription_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionRequest) ProtoMessage() {}

func (x *GetSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_subscription_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_subscription_proto_rawDescGZIP(), []int{7}
}

func (x *GetSubscriptionRequest) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

var File_inboxapi_subscription_proto protoreflect.FileDescriptor

var file_inboxapi_subscription_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61,
	0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a,
	0x16, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x61, 0x6f, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x6f, 0x49, 0x64, 0x22, 0x32,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e, 0x62, 0x6f,
	0x78, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x4e, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64,
	0x61, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x6f,
	0x49, 0x64, 0x22, 0xb2, 0x01, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x61, 0x6f, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x6f, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x70, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3d, 0x0a, 0x12, 0x55, 0x6e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xa9, 0x03, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x09, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x1f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x60, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x24, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x55, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4d, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x69, 0x6e, 0x62, 0x6f, 0x78,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inboxapi_subscription_proto_rawDescOnce sync.Once
	file_inboxapi_subscription_proto_rawDescData = file_inboxapi_subscription_proto_rawDesc
)

func file_inboxapi_subscription_proto_rawDescGZIP() []byte {
	file_inboxapi_subscription_proto_rawDescOnce.Do(func() {
		file_inboxapi_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_inboxapi_subscription_proto_rawDescData)
	})
	return file_inboxapi_subscription_proto_rawDescData
}

var file_inboxapi_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_inboxapi_subscription_proto_goTypes = []interface{}{
	(*FindSubscribersRequest)(nil),   // 0: internalapi.FindSubscribersRequest
	(*UserList)(nil),                 // 1: internalapi.UserList
	(*SubscribeRequest)(nil),         // 2: internalapi.SubscribeRequest
	(*SubscriptionInfo)(nil),         // 3: internalapi.SubscriptionInfo
	(*ListSubscriptionRequest)(nil),  // 4: internalapi.ListSubscriptionRequest
	(*ListSubscriptionResponse)(nil), // 5: internalapi.ListSubscriptionResponse
	(*UnsubscribeRequest)(nil),       // 6: internalapi.UnsubscribeRequest
	(*GetSubscriptionRequest)(nil),   // 7: internalapi.GetSubscriptionRequest
	(*UserID)(nil),                   // 8: inboxapi.UserID
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 10: google.protobuf.Empty
}
var file_inboxapi_subscription_proto_depIdxs = []int32{
	8,  // 0: internalapi.UserList.users:type_name -> inboxapi.UserID
	9,  // 1: internalapi.SubscriptionInfo.created_at:type_name -> google.protobuf.Timestamp
	3,  // 2: internalapi.ListSubscriptionResponse.items:type_name -> internalapi.SubscriptionInfo
	2,  // 3: internalapi.Subscription.Subscribe:input_type -> internalapi.SubscribeRequest
	6,  // 4: internalapi.Subscription.Unsubscribe:input_type -> internalapi.UnsubscribeRequest
	4,  // 5: internalapi.Subscription.ListSubscriptions:input_type -> internalapi.ListSubscriptionRequest
	7,  // 6: internalapi.Subscription.GetSubscription:input_type -> internalapi.GetSubscriptionRequest
	0,  // 7: internalapi.Subscription.FindSubscribers:input_type -> internalapi.FindSubscribersRequest
	3,  // 8: internalapi.Subscription.Subscribe:output_type -> internalapi.SubscriptionInfo
	10, // 9: internalapi.Subscription.Unsubscribe:output_type -> google.protobuf.Empty
	5,  // 10: internalapi.Subscription.ListSubscriptions:output_type -> internalapi.ListSubscriptionResponse
	3,  // 11: internalapi.Subscription.GetSubscription:output_type -> internalapi.SubscriptionInfo
	1,  // 12: internalapi.Subscription.FindSubscribers:output_type -> internalapi.UserList
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_inboxapi_subscription_proto_init() }
func file_inboxapi_subscription_proto_init() {
	if File_inboxapi_subscription_proto != nil {
		return
	}
	file_inboxapi_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_inboxapi_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindSubscribersRequest); i {
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
		file_inboxapi_subscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserList); i {
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
		file_inboxapi_subscription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_inboxapi_subscription_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionInfo); i {
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
		file_inboxapi_subscription_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionRequest); i {
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
		file_inboxapi_subscription_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionResponse); i {
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
		file_inboxapi_subscription_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeRequest); i {
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
		file_inboxapi_subscription_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionRequest); i {
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
	file_inboxapi_subscription_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inboxapi_subscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inboxapi_subscription_proto_goTypes,
		DependencyIndexes: file_inboxapi_subscription_proto_depIdxs,
		MessageInfos:      file_inboxapi_subscription_proto_msgTypes,
	}.Build()
	File_inboxapi_subscription_proto = out.File
	file_inboxapi_subscription_proto_rawDesc = nil
	file_inboxapi_subscription_proto_goTypes = nil
	file_inboxapi_subscription_proto_depIdxs = nil
}