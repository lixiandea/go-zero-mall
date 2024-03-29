// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: video_rpc.proto

package video_rpc

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

// VideoRpcClient is the client API for VideoRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoRpcClient interface {
	UpdateVideo(ctx context.Context, in *UpdateVideoRequest, opts ...grpc.CallOption) (*UpdateVideoResponse, error)
	DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*DeleteVideoResponse, error)
	AddVideo(ctx context.Context, in *AddVideoRequest, opts ...grpc.CallOption) (*AddVideoResponse, error)
	GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoResponse, error)
}

type videoRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoRpcClient(cc grpc.ClientConnInterface) VideoRpcClient {
	return &videoRpcClient{cc}
}

func (c *videoRpcClient) UpdateVideo(ctx context.Context, in *UpdateVideoRequest, opts ...grpc.CallOption) (*UpdateVideoResponse, error) {
	out := new(UpdateVideoResponse)
	err := c.cc.Invoke(ctx, "/video_rpc.video_rpc/UpdateVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*DeleteVideoResponse, error) {
	out := new(DeleteVideoResponse)
	err := c.cc.Invoke(ctx, "/video_rpc.video_rpc/DeleteVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) AddVideo(ctx context.Context, in *AddVideoRequest, opts ...grpc.CallOption) (*AddVideoResponse, error) {
	out := new(AddVideoResponse)
	err := c.cc.Invoke(ctx, "/video_rpc.video_rpc/AddVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoResponse, error) {
	out := new(GetVideoResponse)
	err := c.cc.Invoke(ctx, "/video_rpc.video_rpc/GetVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoRpcServer is the server API for VideoRpc service.
// All implementations must embed UnimplementedVideoRpcServer
// for forward compatibility
type VideoRpcServer interface {
	UpdateVideo(context.Context, *UpdateVideoRequest) (*UpdateVideoResponse, error)
	DeleteVideo(context.Context, *DeleteVideoRequest) (*DeleteVideoResponse, error)
	AddVideo(context.Context, *AddVideoRequest) (*AddVideoResponse, error)
	GetVideo(context.Context, *GetVideoRequest) (*GetVideoResponse, error)
	mustEmbedUnimplementedVideoRpcServer()
}

// UnimplementedVideoRpcServer must be embedded to have forward compatible implementations.
type UnimplementedVideoRpcServer struct {
}

func (UnimplementedVideoRpcServer) UpdateVideo(context.Context, *UpdateVideoRequest) (*UpdateVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVideo not implemented")
}
func (UnimplementedVideoRpcServer) DeleteVideo(context.Context, *DeleteVideoRequest) (*DeleteVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVideo not implemented")
}
func (UnimplementedVideoRpcServer) AddVideo(context.Context, *AddVideoRequest) (*AddVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVideo not implemented")
}
func (UnimplementedVideoRpcServer) GetVideo(context.Context, *GetVideoRequest) (*GetVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}
func (UnimplementedVideoRpcServer) mustEmbedUnimplementedVideoRpcServer() {}

// UnsafeVideoRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoRpcServer will
// result in compilation errors.
type UnsafeVideoRpcServer interface {
	mustEmbedUnimplementedVideoRpcServer()
}

func RegisterVideoRpcServer(s grpc.ServiceRegistrar, srv VideoRpcServer) {
	s.RegisterService(&VideoRpc_ServiceDesc, srv)
}

func _VideoRpc_UpdateVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).UpdateVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video_rpc.video_rpc/UpdateVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).UpdateVideo(ctx, req.(*UpdateVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_DeleteVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).DeleteVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video_rpc.video_rpc/DeleteVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).DeleteVideo(ctx, req.(*DeleteVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_AddVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).AddVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video_rpc.video_rpc/AddVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).AddVideo(ctx, req.(*AddVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video_rpc.video_rpc/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).GetVideo(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoRpc_ServiceDesc is the grpc.ServiceDesc for VideoRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video_rpc.video_rpc",
	HandlerType: (*VideoRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateVideo",
			Handler:    _VideoRpc_UpdateVideo_Handler,
		},
		{
			MethodName: "DeleteVideo",
			Handler:    _VideoRpc_DeleteVideo_Handler,
		},
		{
			MethodName: "AddVideo",
			Handler:    _VideoRpc_AddVideo_Handler,
		},
		{
			MethodName: "GetVideo",
			Handler:    _VideoRpc_GetVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video_rpc.proto",
}
