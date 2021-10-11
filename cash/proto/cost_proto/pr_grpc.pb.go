// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cost_proto

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

// CostClient is the client API for Cost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CostClient interface {
	CreateCost(ctx context.Context, in *CreateCostRequest, opts ...grpc.CallOption) (*CostProfile, error)
}

type costClient struct {
	cc grpc.ClientConnInterface
}

func NewCostClient(cc grpc.ClientConnInterface) CostClient {
	return &costClient{cc}
}

func (c *costClient) CreateCost(ctx context.Context, in *CreateCostRequest, opts ...grpc.CallOption) (*CostProfile, error) {
	out := new(CostProfile)
	err := c.cc.Invoke(ctx, "/cost_proto.Cost/CreateCost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CostServer is the server API for Cost service.
// All implementations must embed UnimplementedCostServer
// for forward compatibility
type CostServer interface {
	CreateCost(context.Context, *CreateCostRequest) (*CostProfile, error)
	mustEmbedUnimplementedCostServer()
}

// UnimplementedCostServer must be embedded to have forward compatible implementations.
type UnimplementedCostServer struct {
}

func (UnimplementedCostServer) CreateCost(context.Context, *CreateCostRequest) (*CostProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCost not implemented")
}
func (UnimplementedCostServer) mustEmbedUnimplementedCostServer() {}

// UnsafeCostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CostServer will
// result in compilation errors.
type UnsafeCostServer interface {
	mustEmbedUnimplementedCostServer()
}

func RegisterCostServer(s grpc.ServiceRegistrar, srv CostServer) {
	s.RegisterService(&Cost_ServiceDesc, srv)
}

func _Cost_CreateCost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).CreateCost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cost_proto.Cost/CreateCost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).CreateCost(ctx, req.(*CreateCostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cost_ServiceDesc is the grpc.ServiceDesc for Cost service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cost_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cost_proto.Cost",
	HandlerType: (*CostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCost",
			Handler:    _Cost_CreateCost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pr.proto",
}