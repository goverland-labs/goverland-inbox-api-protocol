// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: inboxapi/settings.proto

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
	Settings_AddPushToken_FullMethodName     = "/inboxapi.Settings/AddPushToken"
	Settings_RemovePushToken_FullMethodName  = "/inboxapi.Settings/RemovePushToken"
	Settings_PushTokenExists_FullMethodName  = "/inboxapi.Settings/PushTokenExists"
	Settings_GetPushToken_FullMethodName     = "/inboxapi.Settings/GetPushToken"
	Settings_GetPushTokenList_FullMethodName = "/inboxapi.Settings/GetPushTokenList"
)

// SettingsClient is the client API for Settings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingsClient interface {
	AddPushToken(ctx context.Context, in *AddPushTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemovePushToken(ctx context.Context, in *RemovePushTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PushTokenExists(ctx context.Context, in *PushTokenExistsRequest, opts ...grpc.CallOption) (*PushTokenExistsResponse, error)
	// Deprecated: Do not use.
	// Deprecated: use GetPushTokenList instead
	GetPushToken(ctx context.Context, in *GetPushTokenRequest, opts ...grpc.CallOption) (*PushTokenResponse, error)
	// GetPushTokenList returns list of saved pushed tokens
	GetPushTokenList(ctx context.Context, in *GetPushTokenListRequest, opts ...grpc.CallOption) (*PushTokenListResponse, error)
}

type settingsClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsClient(cc grpc.ClientConnInterface) SettingsClient {
	return &settingsClient{cc}
}

func (c *settingsClient) AddPushToken(ctx context.Context, in *AddPushTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Settings_AddPushToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) RemovePushToken(ctx context.Context, in *RemovePushTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Settings_RemovePushToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) PushTokenExists(ctx context.Context, in *PushTokenExistsRequest, opts ...grpc.CallOption) (*PushTokenExistsResponse, error) {
	out := new(PushTokenExistsResponse)
	err := c.cc.Invoke(ctx, Settings_PushTokenExists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *settingsClient) GetPushToken(ctx context.Context, in *GetPushTokenRequest, opts ...grpc.CallOption) (*PushTokenResponse, error) {
	out := new(PushTokenResponse)
	err := c.cc.Invoke(ctx, Settings_GetPushToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetPushTokenList(ctx context.Context, in *GetPushTokenListRequest, opts ...grpc.CallOption) (*PushTokenListResponse, error) {
	out := new(PushTokenListResponse)
	err := c.cc.Invoke(ctx, Settings_GetPushTokenList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServer is the server API for Settings service.
// All implementations must embed UnimplementedSettingsServer
// for forward compatibility
type SettingsServer interface {
	AddPushToken(context.Context, *AddPushTokenRequest) (*emptypb.Empty, error)
	RemovePushToken(context.Context, *RemovePushTokenRequest) (*emptypb.Empty, error)
	PushTokenExists(context.Context, *PushTokenExistsRequest) (*PushTokenExistsResponse, error)
	// Deprecated: Do not use.
	// Deprecated: use GetPushTokenList instead
	GetPushToken(context.Context, *GetPushTokenRequest) (*PushTokenResponse, error)
	// GetPushTokenList returns list of saved pushed tokens
	GetPushTokenList(context.Context, *GetPushTokenListRequest) (*PushTokenListResponse, error)
	mustEmbedUnimplementedSettingsServer()
}

// UnimplementedSettingsServer must be embedded to have forward compatible implementations.
type UnimplementedSettingsServer struct {
}

func (UnimplementedSettingsServer) AddPushToken(context.Context, *AddPushTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPushToken not implemented")
}
func (UnimplementedSettingsServer) RemovePushToken(context.Context, *RemovePushTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePushToken not implemented")
}
func (UnimplementedSettingsServer) PushTokenExists(context.Context, *PushTokenExistsRequest) (*PushTokenExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushTokenExists not implemented")
}
func (UnimplementedSettingsServer) GetPushToken(context.Context, *GetPushTokenRequest) (*PushTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPushToken not implemented")
}
func (UnimplementedSettingsServer) GetPushTokenList(context.Context, *GetPushTokenListRequest) (*PushTokenListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPushTokenList not implemented")
}
func (UnimplementedSettingsServer) mustEmbedUnimplementedSettingsServer() {}

// UnsafeSettingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingsServer will
// result in compilation errors.
type UnsafeSettingsServer interface {
	mustEmbedUnimplementedSettingsServer()
}

func RegisterSettingsServer(s grpc.ServiceRegistrar, srv SettingsServer) {
	s.RegisterService(&Settings_ServiceDesc, srv)
}

func _Settings_AddPushToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPushTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).AddPushToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Settings_AddPushToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).AddPushToken(ctx, req.(*AddPushTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_RemovePushToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePushTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).RemovePushToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Settings_RemovePushToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).RemovePushToken(ctx, req.(*RemovePushTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_PushTokenExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushTokenExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).PushTokenExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Settings_PushTokenExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).PushTokenExists(ctx, req.(*PushTokenExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetPushToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPushTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetPushToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Settings_GetPushToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetPushToken(ctx, req.(*GetPushTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetPushTokenList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPushTokenListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetPushTokenList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Settings_GetPushTokenList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetPushTokenList(ctx, req.(*GetPushTokenListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Settings_ServiceDesc is the grpc.ServiceDesc for Settings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Settings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inboxapi.Settings",
	HandlerType: (*SettingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPushToken",
			Handler:    _Settings_AddPushToken_Handler,
		},
		{
			MethodName: "RemovePushToken",
			Handler:    _Settings_RemovePushToken_Handler,
		},
		{
			MethodName: "PushTokenExists",
			Handler:    _Settings_PushTokenExists_Handler,
		},
		{
			MethodName: "GetPushToken",
			Handler:    _Settings_GetPushToken_Handler,
		},
		{
			MethodName: "GetPushTokenList",
			Handler:    _Settings_GetPushTokenList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inboxapi/settings.proto",
}
