// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: inboxapi/subscription.proto

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

// SubscriptionClient is the client API for Subscription service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscriptionInfo, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListSubscriptions(ctx context.Context, in *ListSubscriptionRequest, opts ...grpc.CallOption) (*ListSubscriptionResponse, error)
	GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*SubscriptionInfo, error)
	FindSubscribers(ctx context.Context, in *FindSubscribersRequest, opts ...grpc.CallOption) (*UserList, error)
}

type subscriptionClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionClient(cc grpc.ClientConnInterface) SubscriptionClient {
	return &subscriptionClient{cc}
}

func (c *subscriptionClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscriptionInfo, error) {
	out := new(SubscriptionInfo)
	err := c.cc.Invoke(ctx, "/internalapi.Subscription/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internalapi.Subscription/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionClient) ListSubscriptions(ctx context.Context, in *ListSubscriptionRequest, opts ...grpc.CallOption) (*ListSubscriptionResponse, error) {
	out := new(ListSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/internalapi.Subscription/ListSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionClient) GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*SubscriptionInfo, error) {
	out := new(SubscriptionInfo)
	err := c.cc.Invoke(ctx, "/internalapi.Subscription/GetSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionClient) FindSubscribers(ctx context.Context, in *FindSubscribersRequest, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/internalapi.Subscription/FindSubscribers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServer is the server API for Subscription service.
// All implementations must embed UnimplementedSubscriptionServer
// for forward compatibility
type SubscriptionServer interface {
	Subscribe(context.Context, *SubscribeRequest) (*SubscriptionInfo, error)
	Unsubscribe(context.Context, *UnsubscribeRequest) (*emptypb.Empty, error)
	ListSubscriptions(context.Context, *ListSubscriptionRequest) (*ListSubscriptionResponse, error)
	GetSubscription(context.Context, *GetSubscriptionRequest) (*SubscriptionInfo, error)
	FindSubscribers(context.Context, *FindSubscribersRequest) (*UserList, error)
	mustEmbedUnimplementedSubscriptionServer()
}

// UnimplementedSubscriptionServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServer struct {
}

func (UnimplementedSubscriptionServer) Subscribe(context.Context, *SubscribeRequest) (*SubscriptionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSubscriptionServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedSubscriptionServer) ListSubscriptions(context.Context, *ListSubscriptionRequest) (*ListSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscriptions not implemented")
}
func (UnimplementedSubscriptionServer) GetSubscription(context.Context, *GetSubscriptionRequest) (*SubscriptionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscription not implemented")
}
func (UnimplementedSubscriptionServer) FindSubscribers(context.Context, *FindSubscribersRequest) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSubscribers not implemented")
}
func (UnimplementedSubscriptionServer) mustEmbedUnimplementedSubscriptionServer() {}

// UnsafeSubscriptionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionServer will
// result in compilation errors.
type UnsafeSubscriptionServer interface {
	mustEmbedUnimplementedSubscriptionServer()
}

func RegisterSubscriptionServer(s grpc.ServiceRegistrar, srv SubscriptionServer) {
	s.RegisterService(&Subscription_ServiceDesc, srv)
}

func _Subscription_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalapi.Subscription/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscription_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalapi.Subscription/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscription_ListSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServer).ListSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalapi.Subscription/ListSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServer).ListSubscriptions(ctx, req.(*ListSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscription_GetSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServer).GetSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalapi.Subscription/GetSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServer).GetSubscription(ctx, req.(*GetSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscription_FindSubscribers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSubscribersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServer).FindSubscribers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalapi.Subscription/FindSubscribers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServer).FindSubscribers(ctx, req.(*FindSubscribersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Subscription_ServiceDesc is the grpc.ServiceDesc for Subscription service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Subscription_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internalapi.Subscription",
	HandlerType: (*SubscriptionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _Subscription_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _Subscription_Unsubscribe_Handler,
		},
		{
			MethodName: "ListSubscriptions",
			Handler:    _Subscription_ListSubscriptions_Handler,
		},
		{
			MethodName: "GetSubscription",
			Handler:    _Subscription_GetSubscription_Handler,
		},
		{
			MethodName: "FindSubscribers",
			Handler:    _Subscription_FindSubscribers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inboxapi/subscription.proto",
}
