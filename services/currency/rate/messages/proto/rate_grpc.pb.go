// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: rate.proto

package proto

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
	RateService_GetRate_FullMethodName = "/messages.RateService/GetRate"
)

// RateServiceClient is the client API for RateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RateServiceClient interface {
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
}

type rateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRateServiceClient(cc grpc.ClientConnInterface) RateServiceClient {
	return &rateServiceClient{cc}
}

func (c *rateServiceClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, RateService_GetRate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateServiceServer is the server API for RateService service.
// All implementations should embed UnimplementedRateServiceServer
// for forward compatibility
type RateServiceServer interface {
	GetRate(context.Context, *RateRequest) (*RateResponse, error)
}

// UnimplementedRateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRateServiceServer struct {
}

func (UnimplementedRateServiceServer) GetRate(context.Context, *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}

// UnsafeRateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RateServiceServer will
// result in compilation error.
type UnsafeRateServiceServer interface {
	mustEmbedUnimplementedRateServiceServer()
}

func RegisterRateServiceServer(s grpc.ServiceRegistrar, srv RateServiceServer) {
	s.RegisterService(&RateService_ServiceDesc, srv)
}

func _RateService_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RateService_GetRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RateService_ServiceDesc is the grpc.ServiceDesc for RateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messages.RateService",
	HandlerType: (*RateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _RateService_GetRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rate.proto",
}
