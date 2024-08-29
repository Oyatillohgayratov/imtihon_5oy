// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: protos/hotel/hotel.proto

package hotelpb

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
	HotelService_AddHotel_FullMethodName              = "/HotelService/AddHotel"
	HotelService_GetHotels_FullMethodName             = "/HotelService/GetHotels"
	HotelService_GetHotelDetails_FullMethodName       = "/HotelService/GetHotelDetails"
	HotelService_CheckRoomAvailability_FullMethodName = "/HotelService/CheckRoomAvailability"
	HotelService_CheckHotelID_FullMethodName          = "/HotelService/CheckHotelID"
	HotelService_UdateRoomAvailability_FullMethodName = "/HotelService/UdateRoomAvailability"
	HotelService_GetRoomID_FullMethodName             = "/HotelService/GetRoomID"
)

// HotelServiceClient is the client API for HotelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HotelServiceClient interface {
	AddHotel(ctx context.Context, in *Hotel, opts ...grpc.CallOption) (*AddHotelResponse, error)
	GetHotels(ctx context.Context, in *GetHotelsRequest, opts ...grpc.CallOption) (*GetHotelsResponse, error)
	GetHotelDetails(ctx context.Context, in *GetHotelDetailsRequest, opts ...grpc.CallOption) (*GetHotelDetailsResponse, error)
	CheckRoomAvailability(ctx context.Context, in *CheckRoomAvailabilityRequest, opts ...grpc.CallOption) (*CheckRoomAvailabilityResponse, error)
	CheckHotelID(ctx context.Context, in *CheckHotelIDRequest, opts ...grpc.CallOption) (*CheckHotelIDResponse, error)
	UdateRoomAvailability(ctx context.Context, in *UdateRoomAvailabilityRequest, opts ...grpc.CallOption) (*UdateRoomAvailabilityResponse, error)
	GetRoomID(ctx context.Context, in *GetRoomIDRequest, opts ...grpc.CallOption) (*GetRoomIDResponse, error)
}

type hotelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHotelServiceClient(cc grpc.ClientConnInterface) HotelServiceClient {
	return &hotelServiceClient{cc}
}

func (c *hotelServiceClient) AddHotel(ctx context.Context, in *Hotel, opts ...grpc.CallOption) (*AddHotelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddHotelResponse)
	err := c.cc.Invoke(ctx, HotelService_AddHotel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetHotels(ctx context.Context, in *GetHotelsRequest, opts ...grpc.CallOption) (*GetHotelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHotelsResponse)
	err := c.cc.Invoke(ctx, HotelService_GetHotels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetHotelDetails(ctx context.Context, in *GetHotelDetailsRequest, opts ...grpc.CallOption) (*GetHotelDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHotelDetailsResponse)
	err := c.cc.Invoke(ctx, HotelService_GetHotelDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) CheckRoomAvailability(ctx context.Context, in *CheckRoomAvailabilityRequest, opts ...grpc.CallOption) (*CheckRoomAvailabilityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckRoomAvailabilityResponse)
	err := c.cc.Invoke(ctx, HotelService_CheckRoomAvailability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) CheckHotelID(ctx context.Context, in *CheckHotelIDRequest, opts ...grpc.CallOption) (*CheckHotelIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckHotelIDResponse)
	err := c.cc.Invoke(ctx, HotelService_CheckHotelID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) UdateRoomAvailability(ctx context.Context, in *UdateRoomAvailabilityRequest, opts ...grpc.CallOption) (*UdateRoomAvailabilityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UdateRoomAvailabilityResponse)
	err := c.cc.Invoke(ctx, HotelService_UdateRoomAvailability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetRoomID(ctx context.Context, in *GetRoomIDRequest, opts ...grpc.CallOption) (*GetRoomIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRoomIDResponse)
	err := c.cc.Invoke(ctx, HotelService_GetRoomID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelServiceServer is the server API for HotelService service.
// All implementations must embed UnimplementedHotelServiceServer
// for forward compatibility
type HotelServiceServer interface {
	AddHotel(context.Context, *Hotel) (*AddHotelResponse, error)
	GetHotels(context.Context, *GetHotelsRequest) (*GetHotelsResponse, error)
	GetHotelDetails(context.Context, *GetHotelDetailsRequest) (*GetHotelDetailsResponse, error)
	CheckRoomAvailability(context.Context, *CheckRoomAvailabilityRequest) (*CheckRoomAvailabilityResponse, error)
	CheckHotelID(context.Context, *CheckHotelIDRequest) (*CheckHotelIDResponse, error)
	UdateRoomAvailability(context.Context, *UdateRoomAvailabilityRequest) (*UdateRoomAvailabilityResponse, error)
	GetRoomID(context.Context, *GetRoomIDRequest) (*GetRoomIDResponse, error)
	mustEmbedUnimplementedHotelServiceServer()
}

// UnimplementedHotelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHotelServiceServer struct {
}

func (UnimplementedHotelServiceServer) AddHotel(context.Context, *Hotel) (*AddHotelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHotel not implemented")
}
func (UnimplementedHotelServiceServer) GetHotels(context.Context, *GetHotelsRequest) (*GetHotelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotels not implemented")
}
func (UnimplementedHotelServiceServer) GetHotelDetails(context.Context, *GetHotelDetailsRequest) (*GetHotelDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotelDetails not implemented")
}
func (UnimplementedHotelServiceServer) CheckRoomAvailability(context.Context, *CheckRoomAvailabilityRequest) (*CheckRoomAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRoomAvailability not implemented")
}
func (UnimplementedHotelServiceServer) CheckHotelID(context.Context, *CheckHotelIDRequest) (*CheckHotelIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHotelID not implemented")
}
func (UnimplementedHotelServiceServer) UdateRoomAvailability(context.Context, *UdateRoomAvailabilityRequest) (*UdateRoomAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UdateRoomAvailability not implemented")
}
func (UnimplementedHotelServiceServer) GetRoomID(context.Context, *GetRoomIDRequest) (*GetRoomIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomID not implemented")
}
func (UnimplementedHotelServiceServer) mustEmbedUnimplementedHotelServiceServer() {}

// UnsafeHotelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HotelServiceServer will
// result in compilation errors.
type UnsafeHotelServiceServer interface {
	mustEmbedUnimplementedHotelServiceServer()
}

func RegisterHotelServiceServer(s grpc.ServiceRegistrar, srv HotelServiceServer) {
	s.RegisterService(&HotelService_ServiceDesc, srv)
}

func _HotelService_AddHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hotel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).AddHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_AddHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).AddHotel(ctx, req.(*Hotel))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetHotels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetHotels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetHotels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetHotels(ctx, req.(*GetHotelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetHotelDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetHotelDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetHotelDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetHotelDetails(ctx, req.(*GetHotelDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_CheckRoomAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRoomAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).CheckRoomAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_CheckRoomAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).CheckRoomAvailability(ctx, req.(*CheckRoomAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_CheckHotelID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckHotelIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).CheckHotelID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_CheckHotelID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).CheckHotelID(ctx, req.(*CheckHotelIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_UdateRoomAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UdateRoomAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).UdateRoomAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_UdateRoomAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).UdateRoomAvailability(ctx, req.(*UdateRoomAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetRoomID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetRoomID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetRoomID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetRoomID(ctx, req.(*GetRoomIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HotelService_ServiceDesc is the grpc.ServiceDesc for HotelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HotelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HotelService",
	HandlerType: (*HotelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddHotel",
			Handler:    _HotelService_AddHotel_Handler,
		},
		{
			MethodName: "GetHotels",
			Handler:    _HotelService_GetHotels_Handler,
		},
		{
			MethodName: "GetHotelDetails",
			Handler:    _HotelService_GetHotelDetails_Handler,
		},
		{
			MethodName: "CheckRoomAvailability",
			Handler:    _HotelService_CheckRoomAvailability_Handler,
		},
		{
			MethodName: "CheckHotelID",
			Handler:    _HotelService_CheckHotelID_Handler,
		},
		{
			MethodName: "UdateRoomAvailability",
			Handler:    _HotelService_UdateRoomAvailability_Handler,
		},
		{
			MethodName: "GetRoomID",
			Handler:    _HotelService_GetRoomID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/hotel/hotel.proto",
}
