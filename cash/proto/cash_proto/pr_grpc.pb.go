// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cash_proto

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

// CashClient is the client API for Cash service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CashClient interface {
	CreateCash(ctx context.Context, in *CreateCashRequest, opts ...grpc.CallOption) (*CashProfile, error)
}

type cashClient struct {
	cc grpc.ClientConnInterface
}

func NewCashClient(cc grpc.ClientConnInterface) CashClient {
	return &cashClient{cc}
}

func (c *cashClient) CreateCash(ctx context.Context, in *CreateCashRequest, opts ...grpc.CallOption) (*CashProfile, error) {
	out := new(CashProfile)
	err := c.cc.Invoke(ctx, "/cash_proto.Cash/CreateCash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CashServer is the server API for Cash service.
// All implementations must embed UnimplementedCashServer
// for forward compatibility
type CashServer interface {
	CreateCash(context.Context, *CreateCashRequest) (*CashProfile, error)
	mustEmbedUnimplementedCashServer()
}

// UnimplementedCashServer must be embedded to have forward compatible implementations.
type UnimplementedCashServer struct {
}

func (UnimplementedCashServer) CreateCash(context.Context, *CreateCashRequest) (*CashProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCash not implemented")
}
func (UnimplementedCashServer) mustEmbedUnimplementedCashServer() {}

// UnsafeCashServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CashServer will
// result in compilation errors.
type UnsafeCashServer interface {
	mustEmbedUnimplementedCashServer()
}

func RegisterCashServer(s grpc.ServiceRegistrar, srv CashServer) {
	s.RegisterService(&Cash_ServiceDesc, srv)
}

func _Cash_CreateCash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CashServer).CreateCash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash_proto.Cash/CreateCash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CashServer).CreateCash(ctx, req.(*CreateCashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cash_ServiceDesc is the grpc.ServiceDesc for Cash service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cash_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cash_proto.Cash",
	HandlerType: (*CashServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCash",
			Handler:    _Cash_CreateCash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cash_proto/pr.proto",
}
