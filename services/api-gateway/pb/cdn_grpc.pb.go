// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: cdn.proto

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

// VideoStreamingServiceClient is the client API for VideoStreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoStreamingServiceClient interface {
	StreamVideoChunk(ctx context.Context, in *VideoChunkRequest, opts ...grpc.CallOption) (*VideoChunk, error)
	GetRecentVideos(ctx context.Context, in *RecentVideosRequest, opts ...grpc.CallOption) (VideoStreamingService_GetRecentVideosClient, error)
}

type videoStreamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoStreamingServiceClient(cc grpc.ClientConnInterface) VideoStreamingServiceClient {
	return &videoStreamingServiceClient{cc}
}

func (c *videoStreamingServiceClient) StreamVideoChunk(ctx context.Context, in *VideoChunkRequest, opts ...grpc.CallOption) (*VideoChunk, error) {
	out := new(VideoChunk)
	err := c.cc.Invoke(ctx, "/videoStreaming.VideoStreamingService/StreamVideoChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoStreamingServiceClient) GetRecentVideos(ctx context.Context, in *RecentVideosRequest, opts ...grpc.CallOption) (VideoStreamingService_GetRecentVideosClient, error) {
	stream, err := c.cc.NewStream(ctx, &VideoStreamingService_ServiceDesc.Streams[0], "/videoStreaming.VideoStreamingService/GetRecentVideos", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoStreamingServiceGetRecentVideosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VideoStreamingService_GetRecentVideosClient interface {
	Recv() (*RecentVideos, error)
	grpc.ClientStream
}

type videoStreamingServiceGetRecentVideosClient struct {
	grpc.ClientStream
}

func (x *videoStreamingServiceGetRecentVideosClient) Recv() (*RecentVideos, error) {
	m := new(RecentVideos)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VideoStreamingServiceServer is the server API for VideoStreamingService service.
// All implementations must embed UnimplementedVideoStreamingServiceServer
// for forward compatibility
type VideoStreamingServiceServer interface {
	StreamVideoChunk(context.Context, *VideoChunkRequest) (*VideoChunk, error)
	GetRecentVideos(*RecentVideosRequest, VideoStreamingService_GetRecentVideosServer) error
	mustEmbedUnimplementedVideoStreamingServiceServer()
}

// UnimplementedVideoStreamingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoStreamingServiceServer struct {
}

func (UnimplementedVideoStreamingServiceServer) StreamVideoChunk(context.Context, *VideoChunkRequest) (*VideoChunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamVideoChunk not implemented")
}
func (UnimplementedVideoStreamingServiceServer) GetRecentVideos(*RecentVideosRequest, VideoStreamingService_GetRecentVideosServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRecentVideos not implemented")
}
func (UnimplementedVideoStreamingServiceServer) mustEmbedUnimplementedVideoStreamingServiceServer() {}

// UnsafeVideoStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoStreamingServiceServer will
// result in compilation errors.
type UnsafeVideoStreamingServiceServer interface {
	mustEmbedUnimplementedVideoStreamingServiceServer()
}

func RegisterVideoStreamingServiceServer(s grpc.ServiceRegistrar, srv VideoStreamingServiceServer) {
	s.RegisterService(&VideoStreamingService_ServiceDesc, srv)
}

func _VideoStreamingService_StreamVideoChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoStreamingServiceServer).StreamVideoChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/videoStreaming.VideoStreamingService/StreamVideoChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoStreamingServiceServer).StreamVideoChunk(ctx, req.(*VideoChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoStreamingService_GetRecentVideos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RecentVideosRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoStreamingServiceServer).GetRecentVideos(m, &videoStreamingServiceGetRecentVideosServer{stream})
}

type VideoStreamingService_GetRecentVideosServer interface {
	Send(*RecentVideos) error
	grpc.ServerStream
}

type videoStreamingServiceGetRecentVideosServer struct {
	grpc.ServerStream
}

func (x *videoStreamingServiceGetRecentVideosServer) Send(m *RecentVideos) error {
	return x.ServerStream.SendMsg(m)
}

// VideoStreamingService_ServiceDesc is the grpc.ServiceDesc for VideoStreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoStreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "videoStreaming.VideoStreamingService",
	HandlerType: (*VideoStreamingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StreamVideoChunk",
			Handler:    _VideoStreamingService_StreamVideoChunk_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRecentVideos",
			Handler:       _VideoStreamingService_GetRecentVideos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cdn.proto",
}