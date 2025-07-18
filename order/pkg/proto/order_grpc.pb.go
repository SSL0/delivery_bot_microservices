// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: pkg/proto/order.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Order_CreateOrderByCart_FullMethodName = "/order.Order/CreateOrderByCart"
)

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	CreateOrderByCart(ctx context.Context, in *CreateOrderByCartRequest, opts ...grpc.CallOption) (*CreateOrderByCartResponse, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) CreateOrderByCart(ctx context.Context, in *CreateOrderByCartRequest, opts ...grpc.CallOption) (*CreateOrderByCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrderByCartResponse)
	err := c.cc.Invoke(ctx, Order_CreateOrderByCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	CreateOrderByCart(context.Context, *CreateOrderByCartRequest) (*CreateOrderByCartResponse, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) CreateOrderByCart(context.Context, *CreateOrderByCartRequest) (*CreateOrderByCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderByCart not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_CreateOrderByCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderByCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrderByCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CreateOrderByCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrderByCart(ctx, req.(*CreateOrderByCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrderByCart",
			Handler:    _Order_CreateOrderByCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/order.proto",
}
