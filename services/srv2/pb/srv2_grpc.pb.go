// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// Service2Client is the client API for Service2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Service2Client interface {
	Login(ctx context.Context, in *LoginMessage, opts ...grpc.CallOption) (*LoginResultMessage, error)
}

type service2Client struct {
	cc grpc.ClientConnInterface
}

func NewService2Client(cc grpc.ClientConnInterface) Service2Client {
	return &service2Client{cc}
}

func (c *service2Client) Login(ctx context.Context, in *LoginMessage, opts ...grpc.CallOption) (*LoginResultMessage, error) {
	out := new(LoginResultMessage)
	err := c.cc.Invoke(ctx, "/pb.Service2/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service2Server is the server API for Service2 service.
// All implementations must embed UnimplementedService2Server
// for forward compatibility
type Service2Server interface {
	Login(context.Context, *LoginMessage) (*LoginResultMessage, error)
	mustEmbedUnimplementedService2Server()
}

// UnimplementedService2Server must be embedded to have forward compatible implementations.
type UnimplementedService2Server struct {
}

func (*UnimplementedService2Server) Login(context.Context, *LoginMessage) (*LoginResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedService2Server) mustEmbedUnimplementedService2Server() {}

func RegisterService2Server(s *grpc.Server, srv Service2Server) {
	s.RegisterService(&_Service2_serviceDesc, srv)
}

func _Service2_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service2Server).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service2/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service2Server).Login(ctx, req.(*LoginMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Service2",
	HandlerType: (*Service2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Service2_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "srv2.proto",
}