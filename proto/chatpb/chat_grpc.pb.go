// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto/chatpb/chat.proto

package chatpb

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	SentSingleIssue(ctx context.Context, in *SentSingleIssueReq, opts ...grpc.CallOption) (*SentSingleIssueResp, error)
	SentProjectIssue(ctx context.Context, in *SentProjectIssueReq, opts ...grpc.CallOption) (*SentProjectIssueResp, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) SentSingleIssue(ctx context.Context, in *SentSingleIssueReq, opts ...grpc.CallOption) (*SentSingleIssueResp, error) {
	out := new(SentSingleIssueResp)
	err := c.cc.Invoke(ctx, "/chatpb.Chat/SentSingleIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SentProjectIssue(ctx context.Context, in *SentProjectIssueReq, opts ...grpc.CallOption) (*SentProjectIssueResp, error) {
	out := new(SentProjectIssueResp)
	err := c.cc.Invoke(ctx, "/chatpb.Chat/SentProjectIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	SentSingleIssue(context.Context, *SentSingleIssueReq) (*SentSingleIssueResp, error)
	SentProjectIssue(context.Context, *SentProjectIssueReq) (*SentProjectIssueResp, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) SentSingleIssue(context.Context, *SentSingleIssueReq) (*SentSingleIssueResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SentSingleIssue not implemented")
}
func (UnimplementedChatServer) SentProjectIssue(context.Context, *SentProjectIssueReq) (*SentProjectIssueResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SentProjectIssue not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_SentSingleIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentSingleIssueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SentSingleIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatpb.Chat/SentSingleIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SentSingleIssue(ctx, req.(*SentSingleIssueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SentProjectIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentProjectIssueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SentProjectIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatpb.Chat/SentProjectIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SentProjectIssue(ctx, req.(*SentProjectIssueReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatpb.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SentSingleIssue",
			Handler:    _Chat_SentSingleIssue_Handler,
		},
		{
			MethodName: "SentProjectIssue",
			Handler:    _Chat_SentProjectIssue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/chatpb/chat.proto",
}
