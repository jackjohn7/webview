// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: video_processing.proto

package pb

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

// VideoProcessingServiceClient is the client API for VideoProcessingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoProcessingServiceClient interface {
	ProcessNewVideo(ctx context.Context, opts ...grpc.CallOption) (VideoProcessingService_ProcessNewVideoClient, error)
	DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*DeleteVideoResponse, error)
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type videoProcessingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoProcessingServiceClient(cc grpc.ClientConnInterface) VideoProcessingServiceClient {
	return &videoProcessingServiceClient{cc}
}

func (c *videoProcessingServiceClient) ProcessNewVideo(ctx context.Context, opts ...grpc.CallOption) (VideoProcessingService_ProcessNewVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &VideoProcessingService_ServiceDesc.Streams[0], "/videoProcessing.VideoProcessingService/ProcessNewVideo", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoProcessingServiceProcessNewVideoClient{stream}
	return x, nil
}

type VideoProcessingService_ProcessNewVideoClient interface {
	Send(*ProcessVideoRequest) error
	CloseAndRecv() (*ProcessedVideoData, error)
	grpc.ClientStream
}

type videoProcessingServiceProcessNewVideoClient struct {
	grpc.ClientStream
}

func (x *videoProcessingServiceProcessNewVideoClient) Send(m *ProcessVideoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *videoProcessingServiceProcessNewVideoClient) CloseAndRecv() (*ProcessedVideoData, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ProcessedVideoData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *videoProcessingServiceClient) DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*DeleteVideoResponse, error) {
	out := new(DeleteVideoResponse)
	err := c.cc.Invoke(ctx, "/videoProcessing.VideoProcessingService/DeleteVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoProcessingServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/videoProcessing.VideoProcessingService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoProcessingServiceServer is the server API for VideoProcessingService service.
// All implementations must embed UnimplementedVideoProcessingServiceServer
// for forward compatibility
type VideoProcessingServiceServer interface {
	ProcessNewVideo(VideoProcessingService_ProcessNewVideoServer) error
	DeleteVideo(context.Context, *DeleteVideoRequest) (*DeleteVideoResponse, error)
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedVideoProcessingServiceServer()
}

// UnimplementedVideoProcessingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoProcessingServiceServer struct {
}

func (UnimplementedVideoProcessingServiceServer) ProcessNewVideo(VideoProcessingService_ProcessNewVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessNewVideo not implemented")
}
func (UnimplementedVideoProcessingServiceServer) DeleteVideo(context.Context, *DeleteVideoRequest) (*DeleteVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVideo not implemented")
}
func (UnimplementedVideoProcessingServiceServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedVideoProcessingServiceServer) mustEmbedUnimplementedVideoProcessingServiceServer() {
}

// UnsafeVideoProcessingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoProcessingServiceServer will
// result in compilation errors.
type UnsafeVideoProcessingServiceServer interface {
	mustEmbedUnimplementedVideoProcessingServiceServer()
}

func RegisterVideoProcessingServiceServer(s grpc.ServiceRegistrar, srv VideoProcessingServiceServer) {
	s.RegisterService(&VideoProcessingService_ServiceDesc, srv)
}

func _VideoProcessingService_ProcessNewVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoProcessingServiceServer).ProcessNewVideo(&videoProcessingServiceProcessNewVideoServer{stream})
}

type VideoProcessingService_ProcessNewVideoServer interface {
	SendAndClose(*ProcessedVideoData) error
	Recv() (*ProcessVideoRequest, error)
	grpc.ServerStream
}

type videoProcessingServiceProcessNewVideoServer struct {
	grpc.ServerStream
}

func (x *videoProcessingServiceProcessNewVideoServer) SendAndClose(m *ProcessedVideoData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *videoProcessingServiceProcessNewVideoServer) Recv() (*ProcessVideoRequest, error) {
	m := new(ProcessVideoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VideoProcessingService_DeleteVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoProcessingServiceServer).DeleteVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/videoProcessing.VideoProcessingService/DeleteVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoProcessingServiceServer).DeleteVideo(ctx, req.(*DeleteVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoProcessingService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoProcessingServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/videoProcessing.VideoProcessingService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoProcessingServiceServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoProcessingService_ServiceDesc is the grpc.ServiceDesc for VideoProcessingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoProcessingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "videoProcessing.VideoProcessingService",
	HandlerType: (*VideoProcessingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteVideo",
			Handler:    _VideoProcessingService_DeleteVideo_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _VideoProcessingService_HealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessNewVideo",
			Handler:       _VideoProcessingService_ProcessNewVideo_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "video_processing.proto",
}
