// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: hotel.proto

package hotel_v1

import (
	context "context"
	"fmt"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	HotelV1_GetInfo_FullMethodName = "/hotel_v1.HotelV1/GetInfo"
)

// HotelV1Client is the client API for HotelV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HotelV1Client interface {
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
}

type hotelV1Client struct {
	cc grpc.ClientConnInterface
}

func NewHotelV1Client(cc grpc.ClientConnInterface) HotelV1Client {
	return &hotelV1Client{cc}
}

func (c *hotelV1Client) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInfoResponse)
	fmt.Println(out)
	err := c.cc.Invoke(ctx, HotelV1_GetInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelV1Server is the server API for HotelV1 service.
// All implementations must embed UnimplementedHotelV1Server
// for forward compatibility
type HotelV1Server interface {
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
	mustEmbedUnimplementedHotelV1Server()
}

// UnimplementedHotelV1Server must be embedded to have forward compatible implementations.
type UnimplementedHotelV1Server struct {
}

func (UnimplementedHotelV1Server) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedHotelV1Server) mustEmbedUnimplementedHotelV1Server() {}

// UnsafeHotelV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HotelV1Server will
// result in compilation errors.
type UnsafeHotelV1Server interface {
	mustEmbedUnimplementedHotelV1Server()
}

func RegisterHotelV1Server(s grpc.ServiceRegistrar, srv HotelV1Server) {
	s.RegisterService(&HotelV1_ServiceDesc, srv)
}

func _HotelV1_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelV1Server).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelV1_GetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelV1Server).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HotelV1_ServiceDesc is the grpc.ServiceDesc for HotelV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HotelV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hotel_v1.HotelV1",
	HandlerType: (*HotelV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _HotelV1_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hotel.proto",
}
