// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: inboxapi/delegate.proto

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
	Delegate_GetAllowedDaos_FullMethodName    = "/inboxapi.Delegate/GetAllowedDaos"
	Delegate_StoreDelegated_FullMethodName    = "/inboxapi.Delegate/StoreDelegated"
	Delegate_GetLastDelegation_FullMethodName = "/inboxapi.Delegate/GetLastDelegation"
)

// DelegateClient is the client API for Delegate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DelegateClient interface {
	GetAllowedDaos(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllowedDaosResponse, error)
	StoreDelegated(ctx context.Context, in *StoreDelegatedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetLastDelegation(ctx context.Context, in *GetLastDelegationRequest, opts ...grpc.CallOption) (*GetLastDelegationResponse, error)
}

type delegateClient struct {
	cc grpc.ClientConnInterface
}

func NewDelegateClient(cc grpc.ClientConnInterface) DelegateClient {
	return &delegateClient{cc}
}

func (c *delegateClient) GetAllowedDaos(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllowedDaosResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllowedDaosResponse)
	err := c.cc.Invoke(ctx, Delegate_GetAllowedDaos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delegateClient) StoreDelegated(ctx context.Context, in *StoreDelegatedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Delegate_StoreDelegated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delegateClient) GetLastDelegation(ctx context.Context, in *GetLastDelegationRequest, opts ...grpc.CallOption) (*GetLastDelegationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLastDelegationResponse)
	err := c.cc.Invoke(ctx, Delegate_GetLastDelegation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DelegateServer is the server API for Delegate service.
// All implementations must embed UnimplementedDelegateServer
// for forward compatibility.
type DelegateServer interface {
	GetAllowedDaos(context.Context, *emptypb.Empty) (*GetAllowedDaosResponse, error)
	StoreDelegated(context.Context, *StoreDelegatedRequest) (*emptypb.Empty, error)
	GetLastDelegation(context.Context, *GetLastDelegationRequest) (*GetLastDelegationResponse, error)
	mustEmbedUnimplementedDelegateServer()
}

// UnimplementedDelegateServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDelegateServer struct{}

func (UnimplementedDelegateServer) GetAllowedDaos(context.Context, *emptypb.Empty) (*GetAllowedDaosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllowedDaos not implemented")
}
func (UnimplementedDelegateServer) StoreDelegated(context.Context, *StoreDelegatedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreDelegated not implemented")
}
func (UnimplementedDelegateServer) GetLastDelegation(context.Context, *GetLastDelegationRequest) (*GetLastDelegationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastDelegation not implemented")
}
func (UnimplementedDelegateServer) mustEmbedUnimplementedDelegateServer() {}
func (UnimplementedDelegateServer) testEmbeddedByValue()                  {}

// UnsafeDelegateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DelegateServer will
// result in compilation errors.
type UnsafeDelegateServer interface {
	mustEmbedUnimplementedDelegateServer()
}

func RegisterDelegateServer(s grpc.ServiceRegistrar, srv DelegateServer) {
	// If the following call pancis, it indicates UnimplementedDelegateServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Delegate_ServiceDesc, srv)
}

func _Delegate_GetAllowedDaos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegateServer).GetAllowedDaos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Delegate_GetAllowedDaos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegateServer).GetAllowedDaos(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Delegate_StoreDelegated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDelegatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegateServer).StoreDelegated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Delegate_StoreDelegated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegateServer).StoreDelegated(ctx, req.(*StoreDelegatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Delegate_GetLastDelegation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastDelegationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegateServer).GetLastDelegation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Delegate_GetLastDelegation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegateServer).GetLastDelegation(ctx, req.(*GetLastDelegationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Delegate_ServiceDesc is the grpc.ServiceDesc for Delegate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Delegate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inboxapi.Delegate",
	HandlerType: (*DelegateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllowedDaos",
			Handler:    _Delegate_GetAllowedDaos_Handler,
		},
		{
			MethodName: "StoreDelegated",
			Handler:    _Delegate_StoreDelegated_Handler,
		},
		{
			MethodName: "GetLastDelegation",
			Handler:    _Delegate_GetLastDelegation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inboxapi/delegate.proto",
}