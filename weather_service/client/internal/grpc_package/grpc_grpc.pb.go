// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: grpc.proto

package grpc_package

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
	WeatherService_GetWeatherUpdates_FullMethodName = "/grpc_package.WeatherService/GetWeatherUpdates"
)

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис данных о погоде
type WeatherServiceClient interface {
	GetWeatherUpdates(ctx context.Context, in *Request, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Response], error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) GetWeatherUpdates(ctx context.Context, in *Request, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Response], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WeatherService_ServiceDesc.Streams[0], WeatherService_GetWeatherUpdates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Request, Response]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WeatherService_GetWeatherUpdatesClient = grpc.ServerStreamingClient[Response]

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility.
//
// Сервис данных о погоде
type WeatherServiceServer interface {
	GetWeatherUpdates(*Request, grpc.ServerStreamingServer[Response]) error
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWeatherServiceServer struct{}

func (UnimplementedWeatherServiceServer) GetWeatherUpdates(*Request, grpc.ServerStreamingServer[Response]) error {
	return status.Errorf(codes.Unimplemented, "method GetWeatherUpdates not implemented")
}
func (UnimplementedWeatherServiceServer) mustEmbedUnimplementedWeatherServiceServer() {}
func (UnimplementedWeatherServiceServer) testEmbeddedByValue()                        {}

// UnsafeWeatherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServiceServer will
// result in compilation errors.
type UnsafeWeatherServiceServer interface {
	mustEmbedUnimplementedWeatherServiceServer()
}

func RegisterWeatherServiceServer(s grpc.ServiceRegistrar, srv WeatherServiceServer) {
	// If the following call pancis, it indicates UnimplementedWeatherServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WeatherService_ServiceDesc, srv)
}

func _WeatherService_GetWeatherUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WeatherServiceServer).GetWeatherUpdates(m, &grpc.GenericServerStream[Request, Response]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WeatherService_GetWeatherUpdatesServer = grpc.ServerStreamingServer[Response]

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_package.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetWeatherUpdates",
			Handler:       _WeatherService_GetWeatherUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc.proto",
}
