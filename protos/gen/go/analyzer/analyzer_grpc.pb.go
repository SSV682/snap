// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: analyzer/analyzer.proto

package analyzer_v1

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
	Analyzer_CreateSetting_FullMethodName      = "/analyzer.Analyzer/CreateSetting"
	Analyzer_ListActualSettings_FullMethodName = "/analyzer.Analyzer/ListActualSettings"
	Analyzer_DeleteSetting_FullMethodName      = "/analyzer.Analyzer/DeleteSetting"
)

// AnalyzerClient is the client API for Analyzer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyzerClient interface {
	CreateSetting(ctx context.Context, in *CreateSettingRequest, opts ...grpc.CallOption) (*CreateSettingResponse, error)
	ListActualSettings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ActualSettingsResponse, error)
	DeleteSetting(ctx context.Context, in *DeleteSettingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type analyzerClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyzerClient(cc grpc.ClientConnInterface) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) CreateSetting(ctx context.Context, in *CreateSettingRequest, opts ...grpc.CallOption) (*CreateSettingResponse, error) {
	out := new(CreateSettingResponse)
	err := c.cc.Invoke(ctx, Analyzer_CreateSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) ListActualSettings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ActualSettingsResponse, error) {
	out := new(ActualSettingsResponse)
	err := c.cc.Invoke(ctx, Analyzer_ListActualSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) DeleteSetting(ctx context.Context, in *DeleteSettingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Analyzer_DeleteSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyzerServer is the server API for Analyzer service.
// All implementations must embed UnimplementedAnalyzerServer
// for forward compatibility
type AnalyzerServer interface {
	CreateSetting(context.Context, *CreateSettingRequest) (*CreateSettingResponse, error)
	ListActualSettings(context.Context, *emptypb.Empty) (*ActualSettingsResponse, error)
	DeleteSetting(context.Context, *DeleteSettingRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAnalyzerServer()
}

// UnimplementedAnalyzerServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyzerServer struct {
}

func (UnimplementedAnalyzerServer) CreateSetting(context.Context, *CreateSettingRequest) (*CreateSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSetting not implemented")
}
func (UnimplementedAnalyzerServer) ListActualSettings(context.Context, *emptypb.Empty) (*ActualSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActualSettings not implemented")
}
func (UnimplementedAnalyzerServer) DeleteSetting(context.Context, *DeleteSettingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSetting not implemented")
}
func (UnimplementedAnalyzerServer) mustEmbedUnimplementedAnalyzerServer() {}

// UnsafeAnalyzerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyzerServer will
// result in compilation errors.
type UnsafeAnalyzerServer interface {
	mustEmbedUnimplementedAnalyzerServer()
}

func RegisterAnalyzerServer(s grpc.ServiceRegistrar, srv AnalyzerServer) {
	s.RegisterService(&Analyzer_ServiceDesc, srv)
}

func _Analyzer_CreateSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).CreateSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Analyzer_CreateSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).CreateSetting(ctx, req.(*CreateSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_ListActualSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).ListActualSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Analyzer_ListActualSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).ListActualSettings(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_DeleteSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).DeleteSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Analyzer_DeleteSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).DeleteSetting(ctx, req.(*DeleteSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Analyzer_ServiceDesc is the grpc.ServiceDesc for Analyzer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Analyzer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analyzer.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSetting",
			Handler:    _Analyzer_CreateSetting_Handler,
		},
		{
			MethodName: "ListActualSettings",
			Handler:    _Analyzer_ListActualSettings_Handler,
		},
		{
			MethodName: "DeleteSetting",
			Handler:    _Analyzer_DeleteSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyzer/analyzer.proto",
}
