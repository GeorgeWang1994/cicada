// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc1
// source: proto/judge/judge.proto

package judge

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

// JudgeServiceClient is the client API for JudgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JudgeServiceClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	ReceiveEvent(ctx context.Context, in *ReceiveEventRequest, opts ...grpc.CallOption) (*Response, error)
}

type judgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJudgeServiceClient(cc grpc.ClientConnInterface) JudgeServiceClient {
	return &judgeServiceClient{cc}
}

func (c *judgeServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/event.JudgeService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeServiceClient) ReceiveEvent(ctx context.Context, in *ReceiveEventRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/event.JudgeService/ReceiveEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudgeServiceServer is the server API for JudgeService service.
// All implementations must embed UnimplementedJudgeServiceServer
// for forward compatibility
type JudgeServiceServer interface {
	Ping(context.Context, *Empty) (*Response, error)
	ReceiveEvent(context.Context, *ReceiveEventRequest) (*Response, error)
	mustEmbedUnimplementedJudgeServiceServer()
}

// UnimplementedJudgeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJudgeServiceServer struct {
}

func (UnimplementedJudgeServiceServer) Ping(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedJudgeServiceServer) ReceiveEvent(context.Context, *ReceiveEventRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveEvent not implemented")
}
func (UnimplementedJudgeServiceServer) mustEmbedUnimplementedJudgeServiceServer() {}

// UnsafeJudgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JudgeServiceServer will
// result in compilation errors.
type UnsafeJudgeServiceServer interface {
	mustEmbedUnimplementedJudgeServiceServer()
}

func RegisterJudgeServiceServer(s grpc.ServiceRegistrar, srv JudgeServiceServer) {
	s.RegisterService(&JudgeService_ServiceDesc, srv)
}

func _JudgeService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.JudgeService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _JudgeService_ReceiveEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServiceServer).ReceiveEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.JudgeService/ReceiveEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServiceServer).ReceiveEvent(ctx, req.(*ReceiveEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JudgeService_ServiceDesc is the grpc.ServiceDesc for JudgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JudgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.JudgeService",
	HandlerType: (*JudgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _JudgeService_Ping_Handler,
		},
		{
			MethodName: "ReceiveEvent",
			Handler:    _JudgeService_ReceiveEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/judge/judge.proto",
}
