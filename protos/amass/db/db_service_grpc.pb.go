// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protos/amass/db/db_service.proto

package db

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

// DBServiceClient is the client API for DBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBServiceClient interface {
	Run(ctx context.Context, in *DBRequest, opts ...grpc.CallOption) (*DBResponse, error)
}

type dBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDBServiceClient(cc grpc.ClientConnInterface) DBServiceClient {
	return &dBServiceClient{cc}
}

func (c *dBServiceClient) Run(ctx context.Context, in *DBRequest, opts ...grpc.CallOption) (*DBResponse, error) {
	out := new(DBResponse)
	err := c.cc.Invoke(ctx, "/db.DBService/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBServiceServer is the server API for DBService service.
// All implementations must embed UnimplementedDBServiceServer
// for forward compatibility
type DBServiceServer interface {
	Run(context.Context, *DBRequest) (*DBResponse, error)
	mustEmbedUnimplementedDBServiceServer()
}

// UnimplementedDBServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDBServiceServer struct {
}

func (UnimplementedDBServiceServer) Run(context.Context, *DBRequest) (*DBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedDBServiceServer) mustEmbedUnimplementedDBServiceServer() {}

// UnsafeDBServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DBServiceServer will
// result in compilation errors.
type UnsafeDBServiceServer interface {
	mustEmbedUnimplementedDBServiceServer()
}

func RegisterDBServiceServer(s grpc.ServiceRegistrar, srv DBServiceServer) {
	s.RegisterService(&DBService_ServiceDesc, srv)
}

func _DBService_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DBService/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).Run(ctx, req.(*DBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DBService_ServiceDesc is the grpc.ServiceDesc for DBService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DBService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db.DBService",
	HandlerType: (*DBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _DBService_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/amass/db/db_service.proto",
}
