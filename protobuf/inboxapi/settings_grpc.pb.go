// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Settings_AddPushToken_FullMethodName     = "/inboxapi.Settings/AddPushToken"
	Settings_RemovePushToken_FullMethodName  = "/inboxapi.Settings/RemovePushToken"
	Settings_PushTokenExists_FullMethodName  = "/inboxapi.Settings/PushTokenExists"
	Settings_GetPushToken_FullMethodName     = "/inboxapi.Settings/GetPushToken"
	Settings_GetPushTokenList_FullMethodName = "/inboxapi.Settings/GetPushTokenList"
	Settings_SetPushDetails_FullMethodName   = "/inboxapi.Settings/SetPushDetails"
	Settings_GetPushDetails_FullMethodName   = "/inboxapi.Settings/GetPushDetails"
	Settings_SetFeedSettings_FullMethodName  = "/inboxapi.Settings/SetFeedSettings"
	Settings_GetFeedSettings_FullMethodName  = "/inboxapi.Settings/GetFeedSettings"
)

// SettingsClient is the client API for Settings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingsClient interface {
	AddPushToken(ctx context.Context, in *AddPushTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemovePushToken(ctx context.Context, in *RemovePushTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PushTokenExists(ctx context.Context, in *PushTokenExistsRequest, opts ...grpc.CallOption) (*PushTokenExistsResponse, error)
	GetPushToken(ctx context.Context, in *GetPushTokenRequest, opts ...grpc.CallOption) (*PushTokenResponse, error)
	// GetPushTokenList returns list of saved pushed tokens
	GetPushTokenList(ctx context.Context, in *GetPushTokenListRequest, opts ...grpc.CallOption) (*PushTokenListResponse, error)
	SetPushDetails(ctx context.Context, in *SetPushDetailsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPushDetails(ctx context.Context, in *GetPushDetailsRequest, opts ...grpc.CallOption) (*GetPushDetailsResponse, error)
	// Feed settings section
	SetFeedSettings(ctx context.Context, in *SetFeedSettingsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetFeedSettings(ctx context.Context, in *GetFeedSettingsRequest, opts ...grpc.CallOption) (*GetFeedSettingsResponse, error)
}

type settingsClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsClient(cc grpc.ClientConnInterface) SettingsClient {
	return &settingsClient{cc}
}

func (c *settingsClient) AddPushToken(ctx context.Context, in *AddPushTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Settings_AddPushToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) RemovePushToken(ctx context.Context, in *RemovePushTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Settings_RemovePushToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) PushTokenExists(ctx context.Context, in *PushTokenExistsRequest, opts ...grpc.CallOption) (*PushTokenExistsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PushTokenExistsResponse)
	err := c.cc.Invoke(ctx, Settings_PushTokenExists_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetPushToken(ctx context.Context, in *GetPushTokenRequest, opts ...grpc.CallOption) (*PushTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PushTokenResponse)
	err := c.cc.Invoke(ctx, Settings_GetPushToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetPushTokenList(ctx context.Context, in *GetPushTokenListRequest, opts ...grpc.CallOption) (*PushTokenListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PushTokenListResponse)
	err := c.cc.Invoke(ctx, Settings_GetPushTokenList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) SetPushDetails(ctx context.Context, in *SetPushDetailsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Settings_SetPushDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetPushDetails(ctx context.Context, in *GetPushDetailsRequest, opts ...grpc.CallOption) (*GetPushDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPushDetailsResponse)
	err := c.cc.Invoke(ctx, Settings_GetPushDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) SetFeedSettings(ctx context.Context, in *SetFeedSettingsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Settings_SetFeedSettings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetFeedSettings(ctx context.Context, in *GetFeedSettingsRequest, opts ...grpc.CallOption) (*GetFeedSettingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFeedSettingsResponse)
	err := c.cc.Invoke(ctx, Settings_GetFeedSettings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServer is the server API for Settings service.
// All implementations must embed UnimplementedSettingsServer
// for forward compatibility.
type SettingsServer interface {
	AddPushToken(context.Context, *AddPushTokenRequest) (*emptypb.Empty, error)
	RemovePushToken(context.Context, *RemovePushTokenRequest) (*emptypb.Empty, error)
	PushTokenExists(context.Context, *PushTokenExistsRequest) (*PushTokenExistsResponse, error)
	GetPushToken(context.Context, *GetPushTokenRequest) (*PushTokenResponse, error)
	// GetPushTokenList returns list of saved pushed tokens
	GetPushTokenList(context.Context, *GetPushTokenListRequest) (*PushTokenListResponse, error)
	SetPushDetails(context.Context, *SetPushDetailsRequest) (*emptypb.Empty, error)
	GetPushDetails(context.Context, *GetPushDetailsRequest) (*GetPushDetailsResponse, error)
	// Feed settings section
	SetFeedSettings(context.Context, *SetFeedSettingsRequest) (*emptypb.Empty, error)
	GetFeedSettings(context.Context, *GetFeedSettingsRequest) (*GetFeedSettingsResponse, error)
	mustEmbedUnimplementedSettingsServer()
}

// UnimplementedSettingsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSettingsServer struct{}

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
func (UnimplementedSettingsServer) SetPushDetails(context.Context, *SetPushDetailsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPushDetails not implemented")
}
func (UnimplementedSettingsServer) GetPushDetails(context.Context, *GetPushDetailsRequest) (*GetPushDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPushDetails not implemented")
}
func (UnimplementedSettingsServer) SetFeedSettings(context.Context, *SetFeedSettingsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFeedSettings not implemented")
}
func (UnimplementedSettingsServer) GetFeedSettings(context.Context, *GetFeedSettingsRequest) (*GetFeedSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedSettings not implemented")
}
func (UnimplementedSettingsServer) mustEmbedUnimplementedSettingsServer() {}
func (UnimplementedSettingsServer) testEmbeddedByValue()                  {}

// UnsafeSettingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingsServer will
// result in compilation errors.
type UnsafeSettingsServer interface {
	mustEmbedUnimplementedSettingsServer()
}

func RegisterSettingsServer(s grpc.ServiceRegistrar, srv SettingsServer) {
	// If the following call pancis, it indicates UnimplementedSettingsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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

func _Settings_SetPushDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPushDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).SetPushDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Settings_SetPushDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).SetPushDetails(ctx, req.(*SetPushDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetPushDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPushDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetPushDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Settings_GetPushDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetPushDetails(ctx, req.(*GetPushDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_SetFeedSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFeedSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).SetFeedSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Settings_SetFeedSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).SetFeedSettings(ctx, req.(*SetFeedSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetFeedSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetFeedSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Settings_GetFeedSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetFeedSettings(ctx, req.(*GetFeedSettingsRequest))
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
		{
			MethodName: "SetPushDetails",
			Handler:    _Settings_SetPushDetails_Handler,
		},
		{
			MethodName: "GetPushDetails",
			Handler:    _Settings_GetPushDetails_Handler,
		},
		{
			MethodName: "SetFeedSettings",
			Handler:    _Settings_SetFeedSettings_Handler,
		},
		{
			MethodName: "GetFeedSettings",
			Handler:    _Settings_GetFeedSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inboxapi/settings.proto",
}
