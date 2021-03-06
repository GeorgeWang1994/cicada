// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc1
// source: proto/alarm/alarm.proto

package alarm

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

// AlarmServiceClient is the client API for AlarmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlarmServiceClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	ReceiveEvent(ctx context.Context, in *ReceiveEventRequest, opts ...grpc.CallOption) (*Response, error)
}

type alarmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlarmServiceClient(cc grpc.ClientConnInterface) AlarmServiceClient {
	return &alarmServiceClient{cc}
}

func (c *alarmServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/event.AlarmService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmServiceClient) ReceiveEvent(ctx context.Context, in *ReceiveEventRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/event.AlarmService/ReceiveEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlarmServiceServer is the server API for AlarmService service.
// All implementations must embed UnimplementedAlarmServiceServer
// for forward compatibility
type AlarmServiceServer interface {
	Ping(context.Context, *Empty) (*Response, error)
	ReceiveEvent(context.Context, *ReceiveEventRequest) (*Response, error)
	mustEmbedUnimplementedAlarmServiceServer()
}

// UnimplementedAlarmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAlarmServiceServer struct {
}

func (UnimplementedAlarmServiceServer) Ping(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAlarmServiceServer) ReceiveEvent(context.Context, *ReceiveEventRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveEvent not implemented")
}
func (UnimplementedAlarmServiceServer) mustEmbedUnimplementedAlarmServiceServer() {}

// UnsafeAlarmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlarmServiceServer will
// result in compilation errors.
type UnsafeAlarmServiceServer interface {
	mustEmbedUnimplementedAlarmServiceServer()
}

func RegisterAlarmServiceServer(s grpc.ServiceRegistrar, srv AlarmServiceServer) {
	s.RegisterService(&AlarmService_ServiceDesc, srv)
}

func _AlarmService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.AlarmService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmService_ReceiveEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServiceServer).ReceiveEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.AlarmService/ReceiveEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServiceServer).ReceiveEvent(ctx, req.(*ReceiveEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlarmService_ServiceDesc is the grpc.ServiceDesc for AlarmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlarmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.AlarmService",
	HandlerType: (*AlarmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AlarmService_Ping_Handler,
		},
		{
			MethodName: "ReceiveEvent",
			Handler:    _AlarmService_ReceiveEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/alarm/alarm.proto",
}
