// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: inboxapi/feed.proto

package inboxapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Feed_GetUserFeed_FullMethodName      = "/inboxapi.Feed/GetUserFeed"
	Feed_MarkAsRead_FullMethodName       = "/inboxapi.Feed/MarkAsRead"
	Feed_MarkAsUnread_FullMethodName     = "/inboxapi.Feed/MarkAsUnread"
	Feed_MarkAsArchived_FullMethodName   = "/inboxapi.Feed/MarkAsArchived"
	Feed_MarkAsUnarchived_FullMethodName = "/inboxapi.Feed/MarkAsUnarchived"
	Feed_UserSubscribe_FullMethodName    = "/inboxapi.Feed/UserSubscribe"
)

// FeedClient is the client API for Feed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedClient interface {
	GetUserFeed(ctx context.Context, in *GetUserFeedRequest, opts ...grpc.CallOption) (*FeedList, error)
	MarkAsRead(ctx context.Context, in *MarkAsReadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MarkAsUnread(ctx context.Context, in *MarkAsUnreadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MarkAsArchived(ctx context.Context, in *MarkAsArchivedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MarkAsUnarchived(ctx context.Context, in *MarkAsUnarchivedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UserSubscribe(ctx context.Context, in *UserSubscribeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type feedClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedClient(cc grpc.ClientConnInterface) FeedClient {
	return &feedClient{cc}
}

func (c *feedClient) GetUserFeed(ctx context.Context, in *GetUserFeedRequest, opts ...grpc.CallOption) (*FeedList, error) {
	out := new(FeedList)
	err := c.cc.Invoke(ctx, Feed_GetUserFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) MarkAsRead(ctx context.Context, in *MarkAsReadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Feed_MarkAsRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) MarkAsUnread(ctx context.Context, in *MarkAsUnreadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Feed_MarkAsUnread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) MarkAsArchived(ctx context.Context, in *MarkAsArchivedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Feed_MarkAsArchived_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) MarkAsUnarchived(ctx context.Context, in *MarkAsUnarchivedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Feed_MarkAsUnarchived_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) UserSubscribe(ctx context.Context, in *UserSubscribeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Feed_UserSubscribe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServer is the server API for Feed service.
// All implementations must embed UnimplementedFeedServer
// for forward compatibility
type FeedServer interface {
	GetUserFeed(context.Context, *GetUserFeedRequest) (*FeedList, error)
	MarkAsRead(context.Context, *MarkAsReadRequest) (*emptypb.Empty, error)
	MarkAsUnread(context.Context, *MarkAsUnreadRequest) (*emptypb.Empty, error)
	MarkAsArchived(context.Context, *MarkAsArchivedRequest) (*emptypb.Empty, error)
	MarkAsUnarchived(context.Context, *MarkAsUnarchivedRequest) (*emptypb.Empty, error)
	UserSubscribe(context.Context, *UserSubscribeRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedFeedServer()
}

// UnimplementedFeedServer must be embedded to have forward compatible implementations.
type UnimplementedFeedServer struct {
}

func (UnimplementedFeedServer) GetUserFeed(context.Context, *GetUserFeedRequest) (*FeedList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFeed not implemented")
}
func (UnimplementedFeedServer) MarkAsRead(context.Context, *MarkAsReadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsRead not implemented")
}
func (UnimplementedFeedServer) MarkAsUnread(context.Context, *MarkAsUnreadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsUnread not implemented")
}
func (UnimplementedFeedServer) MarkAsArchived(context.Context, *MarkAsArchivedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsArchived not implemented")
}
func (UnimplementedFeedServer) MarkAsUnarchived(context.Context, *MarkAsUnarchivedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsUnarchived not implemented")
}
func (UnimplementedFeedServer) UserSubscribe(context.Context, *UserSubscribeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSubscribe not implemented")
}
func (UnimplementedFeedServer) mustEmbedUnimplementedFeedServer() {}

// UnsafeFeedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServer will
// result in compilation errors.
type UnsafeFeedServer interface {
	mustEmbedUnimplementedFeedServer()
}

func RegisterFeedServer(s grpc.ServiceRegistrar, srv FeedServer) {
	s.RegisterService(&Feed_ServiceDesc, srv)
}

func _Feed_GetUserFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).GetUserFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_GetUserFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).GetUserFeed(ctx, req.(*GetUserFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_MarkAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAsReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).MarkAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_MarkAsRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).MarkAsRead(ctx, req.(*MarkAsReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_MarkAsUnread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAsUnreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).MarkAsUnread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_MarkAsUnread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).MarkAsUnread(ctx, req.(*MarkAsUnreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_MarkAsArchived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAsArchivedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).MarkAsArchived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_MarkAsArchived_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).MarkAsArchived(ctx, req.(*MarkAsArchivedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_MarkAsUnarchived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAsUnarchivedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).MarkAsUnarchived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_MarkAsUnarchived_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).MarkAsUnarchived(ctx, req.(*MarkAsUnarchivedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_UserSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).UserSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_UserSubscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).UserSubscribe(ctx, req.(*UserSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Feed_ServiceDesc is the grpc.ServiceDesc for Feed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Feed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inboxapi.Feed",
	HandlerType: (*FeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserFeed",
			Handler:    _Feed_GetUserFeed_Handler,
		},
		{
			MethodName: "MarkAsRead",
			Handler:    _Feed_MarkAsRead_Handler,
		},
		{
			MethodName: "MarkAsUnread",
			Handler:    _Feed_MarkAsUnread_Handler,
		},
		{
			MethodName: "MarkAsArchived",
			Handler:    _Feed_MarkAsArchived_Handler,
		},
		{
			MethodName: "MarkAsUnarchived",
			Handler:    _Feed_MarkAsUnarchived_Handler,
		},
		{
			MethodName: "UserSubscribe",
			Handler:    _Feed_UserSubscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inboxapi/feed.proto",
}
