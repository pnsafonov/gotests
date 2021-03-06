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

// Service1Client is the client API for Service1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Service1Client interface {
	DoCall1(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg0, error)
	DoCallMsg1(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg1, error)
	DoCallMsg2(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg2, error)
	DoCallMsg3(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg3, error)
	DoCallMsg4(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg4, error)
	DoCallMsg5(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg5, error)
}

type service1Client struct {
	cc grpc.ClientConnInterface
}

func NewService1Client(cc grpc.ClientConnInterface) Service1Client {
	return &service1Client{cc}
}

func (c *service1Client) DoCall1(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg0, error) {
	out := new(Msg0)
	err := c.cc.Invoke(ctx, "/pb.Service1/DoCall1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service1Client) DoCallMsg1(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg1, error) {
	out := new(Msg1)
	err := c.cc.Invoke(ctx, "/pb.Service1/DoCallMsg1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service1Client) DoCallMsg2(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg2, error) {
	out := new(Msg2)
	err := c.cc.Invoke(ctx, "/pb.Service1/DoCallMsg2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service1Client) DoCallMsg3(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg3, error) {
	out := new(Msg3)
	err := c.cc.Invoke(ctx, "/pb.Service1/DoCallMsg3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service1Client) DoCallMsg4(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg4, error) {
	out := new(Msg4)
	err := c.cc.Invoke(ctx, "/pb.Service1/DoCallMsg4", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service1Client) DoCallMsg5(ctx context.Context, in *Msg0, opts ...grpc.CallOption) (*Msg5, error) {
	out := new(Msg5)
	err := c.cc.Invoke(ctx, "/pb.Service1/DoCallMsg5", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service1Server is the server API for Service1 service.
// All implementations must embed UnimplementedService1Server
// for forward compatibility
type Service1Server interface {
	DoCall1(context.Context, *Msg0) (*Msg0, error)
	DoCallMsg1(context.Context, *Msg0) (*Msg1, error)
	DoCallMsg2(context.Context, *Msg0) (*Msg2, error)
	DoCallMsg3(context.Context, *Msg0) (*Msg3, error)
	DoCallMsg4(context.Context, *Msg0) (*Msg4, error)
	DoCallMsg5(context.Context, *Msg0) (*Msg5, error)
	mustEmbedUnimplementedService1Server()
}

// UnimplementedService1Server must be embedded to have forward compatible implementations.
type UnimplementedService1Server struct {
}

func (*UnimplementedService1Server) DoCall1(context.Context, *Msg0) (*Msg0, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoCall1 not implemented")
}
func (*UnimplementedService1Server) DoCallMsg1(context.Context, *Msg0) (*Msg1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoCallMsg1 not implemented")
}
func (*UnimplementedService1Server) DoCallMsg2(context.Context, *Msg0) (*Msg2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoCallMsg2 not implemented")
}
func (*UnimplementedService1Server) DoCallMsg3(context.Context, *Msg0) (*Msg3, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoCallMsg3 not implemented")
}
func (*UnimplementedService1Server) DoCallMsg4(context.Context, *Msg0) (*Msg4, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoCallMsg4 not implemented")
}
func (*UnimplementedService1Server) DoCallMsg5(context.Context, *Msg0) (*Msg5, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoCallMsg5 not implemented")
}
func (*UnimplementedService1Server) mustEmbedUnimplementedService1Server() {}

func RegisterService1Server(s *grpc.Server, srv Service1Server) {
	s.RegisterService(&_Service1_serviceDesc, srv)
}

func _Service1_DoCall1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg0)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service1Server).DoCall1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service1/DoCall1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service1Server).DoCall1(ctx, req.(*Msg0))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service1_DoCallMsg1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg0)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service1Server).DoCallMsg1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service1/DoCallMsg1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service1Server).DoCallMsg1(ctx, req.(*Msg0))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service1_DoCallMsg2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg0)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service1Server).DoCallMsg2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service1/DoCallMsg2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service1Server).DoCallMsg2(ctx, req.(*Msg0))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service1_DoCallMsg3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg0)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service1Server).DoCallMsg3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service1/DoCallMsg3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service1Server).DoCallMsg3(ctx, req.(*Msg0))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service1_DoCallMsg4_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg0)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service1Server).DoCallMsg4(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service1/DoCallMsg4",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service1Server).DoCallMsg4(ctx, req.(*Msg0))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service1_DoCallMsg5_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg0)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service1Server).DoCallMsg5(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service1/DoCallMsg5",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service1Server).DoCallMsg5(ctx, req.(*Msg0))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Service1",
	HandlerType: (*Service1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoCall1",
			Handler:    _Service1_DoCall1_Handler,
		},
		{
			MethodName: "DoCallMsg1",
			Handler:    _Service1_DoCallMsg1_Handler,
		},
		{
			MethodName: "DoCallMsg2",
			Handler:    _Service1_DoCallMsg2_Handler,
		},
		{
			MethodName: "DoCallMsg3",
			Handler:    _Service1_DoCallMsg3_Handler,
		},
		{
			MethodName: "DoCallMsg4",
			Handler:    _Service1_DoCallMsg4_Handler,
		},
		{
			MethodName: "DoCallMsg5",
			Handler:    _Service1_DoCallMsg5_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "srv1.proto",
}
