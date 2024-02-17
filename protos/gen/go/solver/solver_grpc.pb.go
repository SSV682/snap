// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: solver/solver.proto

package solver_v1

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
	Solver_MakeDecision_FullMethodName = "/solver.Solver/MakeDecision"
)

// SolverClient is the client API for Solver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SolverClient interface {
	MakeDecision(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type solverClient struct {
	cc grpc.ClientConnInterface
}

func NewSolverClient(cc grpc.ClientConnInterface) SolverClient {
	return &solverClient{cc}
}

func (c *solverClient) MakeDecision(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Solver_MakeDecision_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolverServer is the server API for Solver service.
// All implementations must embed UnimplementedSolverServer
// for forward compatibility
type SolverServer interface {
	MakeDecision(context.Context, *EventRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSolverServer()
}

// UnimplementedSolverServer must be embedded to have forward compatible implementations.
type UnimplementedSolverServer struct {
}

func (UnimplementedSolverServer) MakeDecision(context.Context, *EventRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeDecision not implemented")
}
func (UnimplementedSolverServer) mustEmbedUnimplementedSolverServer() {}

// UnsafeSolverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SolverServer will
// result in compilation errors.
type UnsafeSolverServer interface {
	mustEmbedUnimplementedSolverServer()
}

func RegisterSolverServer(s grpc.ServiceRegistrar, srv SolverServer) {
	s.RegisterService(&Solver_ServiceDesc, srv)
}

func _Solver_MakeDecision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolverServer).MakeDecision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Solver_MakeDecision_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolverServer).MakeDecision(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Solver_ServiceDesc is the grpc.ServiceDesc for Solver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Solver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "solver.Solver",
	HandlerType: (*SolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeDecision",
			Handler:    _Solver_MakeDecision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "solver/solver.proto",
}
