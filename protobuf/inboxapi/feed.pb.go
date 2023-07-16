// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: inboxapi/feed.proto

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

type GetUserFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId    string `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	IncludeRead     bool   `protobuf:"varint,2,opt,name=include_read,json=includeRead,proto3" json:"include_read,omitempty"`
	IncludeArchived bool   `protobuf:"varint,3,opt,name=include_archived,json=includeArchived,proto3" json:"include_archived,omitempty"`
	Limit           uint32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset          uint32 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetUserFeedRequest) Reset() {
	*x = GetUserFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserFeedRequest) ProtoMessage() {}

func (x *GetUserFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserFeedRequest.ProtoReflect.Descriptor instead.
func (*GetUserFeedRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_feed_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserFeedRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *GetUserFeedRequest) GetIncludeRead() bool {
	if x != nil {
		return x.IncludeRead
	}
	return false
}

func (x *GetUserFeedRequest) GetIncludeArchived() bool {
	if x != nil {
		return x.IncludeArchived
	}
	return false
}

func (x *GetUserFeedRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUserFeedRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type MarkAsReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId string                 `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	Ids          []string               `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	Before       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=before,proto3,oneof" json:"before,omitempty"`
}

func (x *MarkAsReadRequest) Reset() {
	*x = MarkAsReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkAsReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAsReadRequest) ProtoMessage() {}

func (x *MarkAsReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAsReadRequest.ProtoReflect.Descriptor instead.
func (*MarkAsReadRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_feed_proto_rawDescGZIP(), []int{1}
}

func (x *MarkAsReadRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *MarkAsReadRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *MarkAsReadRequest) GetBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.Before
	}
	return nil
}

type MarkAsUnreadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId string   `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	Ids          []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *MarkAsUnreadRequest) Reset() {
	*x = MarkAsUnreadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkAsUnreadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAsUnreadRequest) ProtoMessage() {}

func (x *MarkAsUnreadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_feed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAsUnreadRequest.ProtoReflect.Descriptor instead.
func (*MarkAsUnreadRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_feed_proto_rawDescGZIP(), []int{2}
}

func (x *MarkAsUnreadRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *MarkAsUnreadRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type MarkAsArchivedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId string                 `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	Ids          []string               `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	Before       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=before,proto3,oneof" json:"before,omitempty"`
}

func (x *MarkAsArchivedRequest) Reset() {
	*x = MarkAsArchivedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_feed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkAsArchivedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAsArchivedRequest) ProtoMessage() {}

func (x *MarkAsArchivedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_feed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAsArchivedRequest.ProtoReflect.Descriptor instead.
func (*MarkAsArchivedRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_feed_proto_rawDescGZIP(), []int{3}
}

func (x *MarkAsArchivedRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *MarkAsArchivedRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *MarkAsArchivedRequest) GetBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.Before
	}
	return nil
}

type MarkAsUnrchivedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId string   `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	Ids          []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *MarkAsUnrchivedRequest) Reset() {
	*x = MarkAsUnrchivedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_feed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkAsUnrchivedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAsUnrchivedRequest) ProtoMessage() {}

func (x *MarkAsUnrchivedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_feed_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAsUnrchivedRequest.ProtoReflect.Descriptor instead.
func (*MarkAsUnrchivedRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_feed_proto_rawDescGZIP(), []int{4}
}

func (x *MarkAsUnrchivedRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *MarkAsUnrchivedRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type FeedList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List        []*FeedItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	TotalCount  uint32      `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	UnreadCount uint32      `protobuf:"varint,4,opt,name=unread_count,json=unreadCount,proto3" json:"unread_count,omitempty"`
}

func (x *FeedList) Reset() {
	*x = FeedList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_feed_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedList) ProtoMessage() {}

func (x *FeedList) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_feed_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedList.ProtoReflect.Descriptor instead.
func (*FeedList) Descriptor() ([]byte, []int) {
	return file_inboxapi_feed_proto_rawDescGZIP(), []int{5}
}

func (x *FeedList) GetList() []*FeedItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *FeedList) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *FeedList) GetUnreadCount() uint32 {
	if x != nil {
		return x.UnreadCount
	}
	return 0
}

type FeedItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ReadAt       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=read_at,json=readAt,proto3" json:"read_at,omitempty"`
	ArchivedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=archived_at,json=archivedAt,proto3" json:"archived_at,omitempty"`
	DaoId        string                 `protobuf:"bytes,6,opt,name=dao_id,json=daoId,proto3" json:"dao_id,omitempty"`
	ProposalId   *string                `protobuf:"bytes,7,opt,name=proposal_id,json=proposalId,proto3,oneof" json:"proposal_id,omitempty"`
	DiscussionId *string                `protobuf:"bytes,8,opt,name=discussion_id,json=discussionId,proto3,oneof" json:"discussion_id,omitempty"`
	Type         string                 `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	Action       string                 `protobuf:"bytes,10,opt,name=action,proto3" json:"action,omitempty"`
	Snapshot     []byte                 `protobuf:"bytes,11,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	Timeline     []byte                 `protobuf:"bytes,12,opt,name=timeline,proto3" json:"timeline,omitempty"`
}

func (x *FeedItem) Reset() {
	*x = FeedItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_feed_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItem) ProtoMessage() {}

func (x *FeedItem) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_feed_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItem.ProtoReflect.Descriptor instead.
func (*FeedItem) Descriptor() ([]byte, []int) {
	return file_inboxapi_feed_proto_rawDescGZIP(), []int{6}
}

func (x *FeedItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FeedItem) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *FeedItem) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *FeedItem) GetReadAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReadAt
	}
	return nil
}

func (x *FeedItem) GetArchivedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ArchivedAt
	}
	return nil
}

func (x *FeedItem) GetDaoId() string {
	if x != nil {
		return x.DaoId
	}
	return ""
}

func (x *FeedItem) GetProposalId() string {
	if x != nil && x.ProposalId != nil {
		return *x.ProposalId
	}
	return ""
}

func (x *FeedItem) GetDiscussionId() string {
	if x != nil && x.DiscussionId != nil {
		return *x.DiscussionId
	}
	return ""
}

func (x *FeedItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FeedItem) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *FeedItem) GetSnapshot() []byte {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

func (x *FeedItem) GetTimeline() []byte {
	if x != nil {
		return x.Timeline
	}
	return nil
}

var File_inboxapi_feed_proto protoreflect.FileDescriptor

var file_inboxapi_feed_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x65, 0x61, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x11, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x73,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x12, 0x37, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00,
	0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x22, 0x4c, 0x0a, 0x13, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x73,
	0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x15, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x73, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x37, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x22, 0x4f, 0x0a, 0x16, 0x4d, 0x61, 0x72,
	0x6b, 0x41, 0x73, 0x55, 0x6e, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x76, 0x0a, 0x08, 0x46, 0x65,
	0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e,
	0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xef, 0x03, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x06, 0x72, 0x65, 0x61, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x61, 0x6f, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x6f, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x64,
	0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x32, 0xea, 0x02, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x3f, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x12, 0x1c, 0x2e, 0x69,
	0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x69, 0x6e, 0x62,
	0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41,
	0x0a, 0x0a, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x69,
	0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x73, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x45, 0x0a, 0x0c, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x73, 0x55, 0x6e, 0x72, 0x65, 0x61,
	0x64, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x41, 0x73, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x0e, 0x4d, 0x61, 0x72, 0x6b,
	0x41, 0x73, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1f, 0x2e, 0x69, 0x6e, 0x62,
	0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x73, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x10, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x73, 0x55, 0x6e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x12, 0x20, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61,
	0x70, 0x69, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x73, 0x55, 0x6e, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inboxapi_feed_proto_rawDescOnce sync.Once
	file_inboxapi_feed_proto_rawDescData = file_inboxapi_feed_proto_rawDesc
)

func file_inboxapi_feed_proto_rawDescGZIP() []byte {
	file_inboxapi_feed_proto_rawDescOnce.Do(func() {
		file_inboxapi_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_inboxapi_feed_proto_rawDescData)
	})
	return file_inboxapi_feed_proto_rawDescData
}

var file_inboxapi_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_inboxapi_feed_proto_goTypes = []interface{}{
	(*GetUserFeedRequest)(nil),     // 0: inboxapi.GetUserFeedRequest
	(*MarkAsReadRequest)(nil),      // 1: inboxapi.MarkAsReadRequest
	(*MarkAsUnreadRequest)(nil),    // 2: inboxapi.MarkAsUnreadRequest
	(*MarkAsArchivedRequest)(nil),  // 3: inboxapi.MarkAsArchivedRequest
	(*MarkAsUnrchivedRequest)(nil), // 4: inboxapi.MarkAsUnrchivedRequest
	(*FeedList)(nil),               // 5: inboxapi.FeedList
	(*FeedItem)(nil),               // 6: inboxapi.FeedItem
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 8: google.protobuf.Empty
}
var file_inboxapi_feed_proto_depIdxs = []int32{
	7,  // 0: inboxapi.MarkAsReadRequest.before:type_name -> google.protobuf.Timestamp
	7,  // 1: inboxapi.MarkAsArchivedRequest.before:type_name -> google.protobuf.Timestamp
	6,  // 2: inboxapi.FeedList.list:type_name -> inboxapi.FeedItem
	7,  // 3: inboxapi.FeedItem.created_at:type_name -> google.protobuf.Timestamp
	7,  // 4: inboxapi.FeedItem.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 5: inboxapi.FeedItem.read_at:type_name -> google.protobuf.Timestamp
	7,  // 6: inboxapi.FeedItem.archived_at:type_name -> google.protobuf.Timestamp
	0,  // 7: inboxapi.Feed.GetUserFeed:input_type -> inboxapi.GetUserFeedRequest
	1,  // 8: inboxapi.Feed.MarkAsRead:input_type -> inboxapi.MarkAsReadRequest
	2,  // 9: inboxapi.Feed.MarkAsUnread:input_type -> inboxapi.MarkAsUnreadRequest
	3,  // 10: inboxapi.Feed.MarkAsArchived:input_type -> inboxapi.MarkAsArchivedRequest
	4,  // 11: inboxapi.Feed.MarkAsUnarchived:input_type -> inboxapi.MarkAsUnrchivedRequest
	5,  // 12: inboxapi.Feed.GetUserFeed:output_type -> inboxapi.FeedList
	8,  // 13: inboxapi.Feed.MarkAsRead:output_type -> google.protobuf.Empty
	8,  // 14: inboxapi.Feed.MarkAsUnread:output_type -> google.protobuf.Empty
	8,  // 15: inboxapi.Feed.MarkAsArchived:output_type -> google.protobuf.Empty
	8,  // 16: inboxapi.Feed.MarkAsUnarchived:output_type -> google.protobuf.Empty
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_inboxapi_feed_proto_init() }
func file_inboxapi_feed_proto_init() {
	if File_inboxapi_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inboxapi_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserFeedRequest); i {
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
		file_inboxapi_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkAsReadRequest); i {
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
		file_inboxapi_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkAsUnreadRequest); i {
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
		file_inboxapi_feed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkAsArchivedRequest); i {
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
		file_inboxapi_feed_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkAsUnrchivedRequest); i {
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
		file_inboxapi_feed_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedList); i {
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
		file_inboxapi_feed_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedItem); i {
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
	file_inboxapi_feed_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_inboxapi_feed_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_inboxapi_feed_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inboxapi_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inboxapi_feed_proto_goTypes,
		DependencyIndexes: file_inboxapi_feed_proto_depIdxs,
		MessageInfos:      file_inboxapi_feed_proto_msgTypes,
	}.Build()
	File_inboxapi_feed_proto = out.File
	file_inboxapi_feed_proto_rawDesc = nil
	file_inboxapi_feed_proto_goTypes = nil
	file_inboxapi_feed_proto_depIdxs = nil
}
