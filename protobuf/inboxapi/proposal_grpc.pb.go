// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: inboxapi/proposal.proto

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
	Proposal_GetFeaturedProposals_FullMethodName = "/inboxapi.Proposal/GetFeaturedProposals"
	Proposal_GetAISummary_FullMethodName         = "/inboxapi.Proposal/GetAISummary"
)

// ProposalClient is the client API for Proposal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProposalClient interface {
	GetFeaturedProposals(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFeaturedProposalsResponse, error)
	// GetAISummary returns AI summary text by the provided proposal identifier
	GetAISummary(ctx context.Context, in *GetAISummaryRequest, opts ...grpc.CallOption) (*GetAISummaryResponse, error)
}

type proposalClient struct {
	cc grpc.ClientConnInterface
}

func NewProposalClient(cc grpc.ClientConnInterface) ProposalClient {
	return &proposalClient{cc}
}

func (c *proposalClient) GetFeaturedProposals(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFeaturedProposalsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFeaturedProposalsResponse)
	err := c.cc.Invoke(ctx, Proposal_GetFeaturedProposals_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proposalClient) GetAISummary(ctx context.Context, in *GetAISummaryRequest, opts ...grpc.CallOption) (*GetAISummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAISummaryResponse)
	err := c.cc.Invoke(ctx, Proposal_GetAISummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProposalServer is the server API for Proposal service.
// All implementations must embed UnimplementedProposalServer
// for forward compatibility.
type ProposalServer interface {
	GetFeaturedProposals(context.Context, *emptypb.Empty) (*GetFeaturedProposalsResponse, error)
	// GetAISummary returns AI summary text by the provided proposal identifier
	GetAISummary(context.Context, *GetAISummaryRequest) (*GetAISummaryResponse, error)
	mustEmbedUnimplementedProposalServer()
}

// UnimplementedProposalServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProposalServer struct{}

func (UnimplementedProposalServer) GetFeaturedProposals(context.Context, *emptypb.Empty) (*GetFeaturedProposalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeaturedProposals not implemented")
}
func (UnimplementedProposalServer) GetAISummary(context.Context, *GetAISummaryRequest) (*GetAISummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAISummary not implemented")
}
func (UnimplementedProposalServer) mustEmbedUnimplementedProposalServer() {}
func (UnimplementedProposalServer) testEmbeddedByValue()                  {}

// UnsafeProposalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProposalServer will
// result in compilation errors.
type UnsafeProposalServer interface {
	mustEmbedUnimplementedProposalServer()
}

func RegisterProposalServer(s grpc.ServiceRegistrar, srv ProposalServer) {
	// If the following call pancis, it indicates UnimplementedProposalServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Proposal_ServiceDesc, srv)
}

func _Proposal_GetFeaturedProposals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProposalServer).GetFeaturedProposals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Proposal_GetFeaturedProposals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProposalServer).GetFeaturedProposals(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proposal_GetAISummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAISummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProposalServer).GetAISummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Proposal_GetAISummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProposalServer).GetAISummary(ctx, req.(*GetAISummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Proposal_ServiceDesc is the grpc.ServiceDesc for Proposal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Proposal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inboxapi.Proposal",
	HandlerType: (*ProposalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeaturedProposals",
			Handler:    _Proposal_GetFeaturedProposals_Handler,
		},
		{
			MethodName: "GetAISummary",
			Handler:    _Proposal_GetAISummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inboxapi/proposal.proto",
}
