// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: inboxapi/achievement.proto

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
	Achievement_GetUserAchievementList_FullMethodName = "/inboxapi.Achievement/GetUserAchievementList"
	Achievement_MarkAsViewed_FullMethodName           = "/inboxapi.Achievement/MarkAsViewed"
)

// AchievementClient is the client API for Achievement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AchievementClient interface {
	GetUserAchievementList(ctx context.Context, in *GetUserAchievementListRequest, opts ...grpc.CallOption) (*AchievementList, error)
	MarkAsViewed(ctx context.Context, in *MarkAsViewedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type achievementClient struct {
	cc grpc.ClientConnInterface
}

func NewAchievementClient(cc grpc.ClientConnInterface) AchievementClient {
	return &achievementClient{cc}
}

func (c *achievementClient) GetUserAchievementList(ctx context.Context, in *GetUserAchievementListRequest, opts ...grpc.CallOption) (*AchievementList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AchievementList)
	err := c.cc.Invoke(ctx, Achievement_GetUserAchievementList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *achievementClient) MarkAsViewed(ctx context.Context, in *MarkAsViewedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Achievement_MarkAsViewed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AchievementServer is the server API for Achievement service.
// All implementations must embed UnimplementedAchievementServer
// for forward compatibility.
type AchievementServer interface {
	GetUserAchievementList(context.Context, *GetUserAchievementListRequest) (*AchievementList, error)
	MarkAsViewed(context.Context, *MarkAsViewedRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAchievementServer()
}

// UnimplementedAchievementServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAchievementServer struct{}

func (UnimplementedAchievementServer) GetUserAchievementList(context.Context, *GetUserAchievementListRequest) (*AchievementList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAchievementList not implemented")
}
func (UnimplementedAchievementServer) MarkAsViewed(context.Context, *MarkAsViewedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsViewed not implemented")
}
func (UnimplementedAchievementServer) mustEmbedUnimplementedAchievementServer() {}
func (UnimplementedAchievementServer) testEmbeddedByValue()                     {}

// UnsafeAchievementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AchievementServer will
// result in compilation errors.
type UnsafeAchievementServer interface {
	mustEmbedUnimplementedAchievementServer()
}

func RegisterAchievementServer(s grpc.ServiceRegistrar, srv AchievementServer) {
	// If the following call pancis, it indicates UnimplementedAchievementServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Achievement_ServiceDesc, srv)
}

func _Achievement_GetUserAchievementList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAchievementListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AchievementServer).GetUserAchievementList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Achievement_GetUserAchievementList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AchievementServer).GetUserAchievementList(ctx, req.(*GetUserAchievementListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Achievement_MarkAsViewed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAsViewedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AchievementServer).MarkAsViewed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Achievement_MarkAsViewed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AchievementServer).MarkAsViewed(ctx, req.(*MarkAsViewedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Achievement_ServiceDesc is the grpc.ServiceDesc for Achievement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Achievement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inboxapi.Achievement",
	HandlerType: (*AchievementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserAchievementList",
			Handler:    _Achievement_GetUserAchievementList_Handler,
		},
		{
			MethodName: "MarkAsViewed",
			Handler:    _Achievement_MarkAsViewed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inboxapi/achievement.proto",
}
