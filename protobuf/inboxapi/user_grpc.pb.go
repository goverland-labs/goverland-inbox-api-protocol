// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: inboxapi/user.proto

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	GetByID(ctx context.Context, in *UserByIDRequest, opts ...grpc.CallOption) (*UserByIDResponse, error)
	GetByUuid(ctx context.Context, in *UserByUuidRequest, opts ...grpc.CallOption) (*UserByUuidResponse, error)
	Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	AddView(ctx context.Context, in *UserViewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LastViewed(ctx context.Context, in *UserLastViewedRequest, opts ...grpc.CallOption) (*UserLastViewedResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetByID(ctx context.Context, in *UserByIDRequest, opts ...grpc.CallOption) (*UserByIDResponse, error) {
	out := new(UserByIDResponse)
	err := c.cc.Invoke(ctx, "/inboxapi.User/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetByUuid(ctx context.Context, in *UserByUuidRequest, opts ...grpc.CallOption) (*UserByUuidResponse, error) {
	out := new(UserByUuidResponse)
	err := c.cc.Invoke(ctx, "/inboxapi.User/GetByUuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, "/inboxapi.User/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddView(ctx context.Context, in *UserViewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/inboxapi.User/AddView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LastViewed(ctx context.Context, in *UserLastViewedRequest, opts ...grpc.CallOption) (*UserLastViewedResponse, error) {
	out := new(UserLastViewedResponse)
	err := c.cc.Invoke(ctx, "/inboxapi.User/LastViewed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	GetByID(context.Context, *UserByIDRequest) (*UserByIDResponse, error)
	GetByUuid(context.Context, *UserByUuidRequest) (*UserByUuidResponse, error)
	Create(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	AddView(context.Context, *UserViewRequest) (*emptypb.Empty, error)
	LastViewed(context.Context, *UserLastViewedRequest) (*UserLastViewedResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GetByID(context.Context, *UserByIDRequest) (*UserByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedUserServer) GetByUuid(context.Context, *UserByUuidRequest) (*UserByUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUuid not implemented")
}
func (UnimplementedUserServer) Create(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServer) AddView(context.Context, *UserViewRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddView not implemented")
}
func (UnimplementedUserServer) LastViewed(context.Context, *UserLastViewedRequest) (*UserLastViewedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastViewed not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inboxapi.User/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetByID(ctx, req.(*UserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserByUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inboxapi.User/GetByUuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetByUuid(ctx, req.(*UserByUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inboxapi.User/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Create(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inboxapi.User/AddView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddView(ctx, req.(*UserViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LastViewed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLastViewedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LastViewed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inboxapi.User/LastViewed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LastViewed(ctx, req.(*UserLastViewedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inboxapi.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _User_GetByID_Handler,
		},
		{
			MethodName: "GetByUuid",
			Handler:    _User_GetByUuid_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _User_Create_Handler,
		},
		{
			MethodName: "AddView",
			Handler:    _User_AddView_Handler,
		},
		{
			MethodName: "LastViewed",
			Handler:    _User_LastViewed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inboxapi/user.proto",
}
