// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: dashboard.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DashboardService_UpdateDashboard_FullMethodName = "/dashboard.DashboardService/UpdateDashboard"
)

// DashboardServiceClient is the client API for DashboardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DashboardServiceClient interface {
	// RPC to update dashboard information
	UpdateDashboard(ctx context.Context, in *DashboardRequest, opts ...grpc.CallOption) (*DashboardResponse, error)
}

type dashboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardServiceClient(cc grpc.ClientConnInterface) DashboardServiceClient {
	return &dashboardServiceClient{cc}
}

func (c *dashboardServiceClient) UpdateDashboard(ctx context.Context, in *DashboardRequest, opts ...grpc.CallOption) (*DashboardResponse, error) {
	out := new(DashboardResponse)
	err := c.cc.Invoke(ctx, DashboardService_UpdateDashboard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardServiceServer is the server API for DashboardService service.
// All implementations must embed UnimplementedDashboardServiceServer
// for forward compatibility
type DashboardServiceServer interface {
	// RPC to update dashboard information
	UpdateDashboard(context.Context, *DashboardRequest) (*DashboardResponse, error)
	mustEmbedUnimplementedDashboardServiceServer()
}

// UnimplementedDashboardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDashboardServiceServer struct {
}

func (UnimplementedDashboardServiceServer) UpdateDashboard(context.Context, *DashboardRequest) (*DashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDashboard not implemented")
}
func (UnimplementedDashboardServiceServer) mustEmbedUnimplementedDashboardServiceServer() {}

// UnsafeDashboardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DashboardServiceServer will
// result in compilation errors.
type UnsafeDashboardServiceServer interface {
	mustEmbedUnimplementedDashboardServiceServer()
}

func RegisterDashboardServiceServer(s grpc.ServiceRegistrar, srv DashboardServiceServer) {
	s.RegisterService(&DashboardService_ServiceDesc, srv)
}

func _DashboardService_UpdateDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).UpdateDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardService_UpdateDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).UpdateDashboard(ctx, req.(*DashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DashboardService_ServiceDesc is the grpc.ServiceDesc for DashboardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DashboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dashboard.DashboardService",
	HandlerType: (*DashboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateDashboard",
			Handler:    _DashboardService_UpdateDashboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dashboard.proto",
}
