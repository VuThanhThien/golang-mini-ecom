// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserGrpc_ReadUser_FullMethodName = "/pb.UserGrpc/ReadUser"
)

// UserGrpcClient is the client API for UserGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserGrpcClient interface {
	ReadUser(ctx context.Context, in *ReadUserRequest, opts ...grpc.CallOption) (*User, error)
}

type userGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewUserGrpcClient(cc grpc.ClientConnInterface) UserGrpcClient {
	return &userGrpcClient{cc}
}

func (c *userGrpcClient) ReadUser(ctx context.Context, in *ReadUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserGrpc_ReadUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGrpcServer is the server API for UserGrpc service.
// All implementations must embed UnimplementedUserGrpcServer
// for forward compatibility.
type UserGrpcServer interface {
	ReadUser(context.Context, *ReadUserRequest) (*User, error)
	mustEmbedUnimplementedUserGrpcServer()
}

// UnimplementedUserGrpcServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserGrpcServer struct{}

func (UnimplementedUserGrpcServer) ReadUser(context.Context, *ReadUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUser not implemented")
}
func (UnimplementedUserGrpcServer) mustEmbedUnimplementedUserGrpcServer() {}
func (UnimplementedUserGrpcServer) testEmbeddedByValue()                  {}

// UnsafeUserGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserGrpcServer will
// result in compilation errors.
type UnsafeUserGrpcServer interface {
	mustEmbedUnimplementedUserGrpcServer()
}

func RegisterUserGrpcServer(s grpc.ServiceRegistrar, srv UserGrpcServer) {
	// If the following call pancis, it indicates UnimplementedUserGrpcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserGrpc_ServiceDesc, srv)
}

func _UserGrpc_ReadUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServer).ReadUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserGrpc_ReadUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServer).ReadUser(ctx, req.(*ReadUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserGrpc_ServiceDesc is the grpc.ServiceDesc for UserGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserGrpc",
	HandlerType: (*UserGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadUser",
			Handler:    _UserGrpc_ReadUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
