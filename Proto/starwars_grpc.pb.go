// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Proto

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

// InformanteClient is the client API for Informante service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InformanteClient interface {
}

type informanteClient struct {
	cc grpc.ClientConnInterface
}

func NewInformanteClient(cc grpc.ClientConnInterface) InformanteClient {
	return &informanteClient{cc}
}

// InformanteServer is the server API for Informante service.
// All implementations must embed UnimplementedInformanteServer
// for forward compatibility
type InformanteServer interface {
	mustEmbedUnimplementedInformanteServer()
}

// UnimplementedInformanteServer must be embedded to have forward compatible implementations.
type UnimplementedInformanteServer struct {
}

func (UnimplementedInformanteServer) mustEmbedUnimplementedInformanteServer() {}

// UnsafeInformanteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InformanteServer will
// result in compilation errors.
type UnsafeInformanteServer interface {
	mustEmbedUnimplementedInformanteServer()
}

func RegisterInformanteServer(s grpc.ServiceRegistrar, srv InformanteServer) {
	s.RegisterService(&Informante_ServiceDesc, srv)
}

// Informante_ServiceDesc is the grpc.ServiceDesc for Informante service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Informante_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.Informante",
	HandlerType: (*InformanteServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "Proto/starwars.proto",
}

// MosEisleyClient is the client API for MosEisley service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MosEisleyClient interface {
	RecibirComando(ctx context.Context, in *SolicitudRecibirComando, opts ...grpc.CallOption) (*RespuestaRecibirComandoMosEisley, error)
	GetNumberRebelds(ctx context.Context, in *SolicitudGetNumberRebelds, opts ...grpc.CallOption) (*RespuestaGetNumberRebelds, error)
}

type mosEisleyClient struct {
	cc grpc.ClientConnInterface
}

func NewMosEisleyClient(cc grpc.ClientConnInterface) MosEisleyClient {
	return &mosEisleyClient{cc}
}

func (c *mosEisleyClient) RecibirComando(ctx context.Context, in *SolicitudRecibirComando, opts ...grpc.CallOption) (*RespuestaRecibirComandoMosEisley, error) {
	out := new(RespuestaRecibirComandoMosEisley)
	err := c.cc.Invoke(ctx, "/Proto.MosEisley/RecibirComando", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mosEisleyClient) GetNumberRebelds(ctx context.Context, in *SolicitudGetNumberRebelds, opts ...grpc.CallOption) (*RespuestaGetNumberRebelds, error) {
	out := new(RespuestaGetNumberRebelds)
	err := c.cc.Invoke(ctx, "/Proto.MosEisley/GetNumberRebelds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MosEisleyServer is the server API for MosEisley service.
// All implementations must embed UnimplementedMosEisleyServer
// for forward compatibility
type MosEisleyServer interface {
	RecibirComando(context.Context, *SolicitudRecibirComando) (*RespuestaRecibirComandoMosEisley, error)
	GetNumberRebelds(context.Context, *SolicitudGetNumberRebelds) (*RespuestaGetNumberRebelds, error)
	mustEmbedUnimplementedMosEisleyServer()
}

// UnimplementedMosEisleyServer must be embedded to have forward compatible implementations.
type UnimplementedMosEisleyServer struct {
}

func (UnimplementedMosEisleyServer) RecibirComando(context.Context, *SolicitudRecibirComando) (*RespuestaRecibirComandoMosEisley, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecibirComando not implemented")
}
func (UnimplementedMosEisleyServer) GetNumberRebelds(context.Context, *SolicitudGetNumberRebelds) (*RespuestaGetNumberRebelds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberRebelds not implemented")
}
func (UnimplementedMosEisleyServer) mustEmbedUnimplementedMosEisleyServer() {}

// UnsafeMosEisleyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MosEisleyServer will
// result in compilation errors.
type UnsafeMosEisleyServer interface {
	mustEmbedUnimplementedMosEisleyServer()
}

func RegisterMosEisleyServer(s grpc.ServiceRegistrar, srv MosEisleyServer) {
	s.RegisterService(&MosEisley_ServiceDesc, srv)
}

func _MosEisley_RecibirComando_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudRecibirComando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MosEisleyServer).RecibirComando(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.MosEisley/RecibirComando",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MosEisleyServer).RecibirComando(ctx, req.(*SolicitudRecibirComando))
	}
	return interceptor(ctx, in, info, handler)
}

func _MosEisley_GetNumberRebelds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudGetNumberRebelds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MosEisleyServer).GetNumberRebelds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.MosEisley/GetNumberRebelds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MosEisleyServer).GetNumberRebelds(ctx, req.(*SolicitudGetNumberRebelds))
	}
	return interceptor(ctx, in, info, handler)
}

// MosEisley_ServiceDesc is the grpc.ServiceDesc for MosEisley service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MosEisley_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.MosEisley",
	HandlerType: (*MosEisleyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecibirComando",
			Handler:    _MosEisley_RecibirComando_Handler,
		},
		{
			MethodName: "GetNumberRebelds",
			Handler:    _MosEisley_GetNumberRebelds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/starwars.proto",
}

// FulcrumClient is the client API for Fulcrum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FulcrumClient interface {
	RecibirComando(ctx context.Context, in *SolicitudRecibirComando, opts ...grpc.CallOption) (*RespuestaRecibirComandoFulcrum, error)
	GetNumberRebelds(ctx context.Context, in *SolicitudGetNumberRebelds, opts ...grpc.CallOption) (*RespuestaGetNumberRebelds, error)
}

type fulcrumClient struct {
	cc grpc.ClientConnInterface
}

func NewFulcrumClient(cc grpc.ClientConnInterface) FulcrumClient {
	return &fulcrumClient{cc}
}

func (c *fulcrumClient) RecibirComando(ctx context.Context, in *SolicitudRecibirComando, opts ...grpc.CallOption) (*RespuestaRecibirComandoFulcrum, error) {
	out := new(RespuestaRecibirComandoFulcrum)
	err := c.cc.Invoke(ctx, "/Proto.Fulcrum/RecibirComando", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) GetNumberRebelds(ctx context.Context, in *SolicitudGetNumberRebelds, opts ...grpc.CallOption) (*RespuestaGetNumberRebelds, error) {
	out := new(RespuestaGetNumberRebelds)
	err := c.cc.Invoke(ctx, "/Proto.Fulcrum/GetNumberRebelds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FulcrumServer is the server API for Fulcrum service.
// All implementations must embed UnimplementedFulcrumServer
// for forward compatibility
type FulcrumServer interface {
	RecibirComando(context.Context, *SolicitudRecibirComando) (*RespuestaRecibirComandoFulcrum, error)
	GetNumberRebelds(context.Context, *SolicitudGetNumberRebelds) (*RespuestaGetNumberRebelds, error)
	mustEmbedUnimplementedFulcrumServer()
}

// UnimplementedFulcrumServer must be embedded to have forward compatible implementations.
type UnimplementedFulcrumServer struct {
}

func (UnimplementedFulcrumServer) RecibirComando(context.Context, *SolicitudRecibirComando) (*RespuestaRecibirComandoFulcrum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecibirComando not implemented")
}
func (UnimplementedFulcrumServer) GetNumberRebelds(context.Context, *SolicitudGetNumberRebelds) (*RespuestaGetNumberRebelds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberRebelds not implemented")
}
func (UnimplementedFulcrumServer) mustEmbedUnimplementedFulcrumServer() {}

// UnsafeFulcrumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FulcrumServer will
// result in compilation errors.
type UnsafeFulcrumServer interface {
	mustEmbedUnimplementedFulcrumServer()
}

func RegisterFulcrumServer(s grpc.ServiceRegistrar, srv FulcrumServer) {
	s.RegisterService(&Fulcrum_ServiceDesc, srv)
}

func _Fulcrum_RecibirComando_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudRecibirComando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).RecibirComando(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.Fulcrum/RecibirComando",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).RecibirComando(ctx, req.(*SolicitudRecibirComando))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_GetNumberRebelds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudGetNumberRebelds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).GetNumberRebelds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.Fulcrum/GetNumberRebelds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).GetNumberRebelds(ctx, req.(*SolicitudGetNumberRebelds))
	}
	return interceptor(ctx, in, info, handler)
}

// Fulcrum_ServiceDesc is the grpc.ServiceDesc for Fulcrum service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fulcrum_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.Fulcrum",
	HandlerType: (*FulcrumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecibirComando",
			Handler:    _Fulcrum_RecibirComando_Handler,
		},
		{
			MethodName: "GetNumberRebelds",
			Handler:    _Fulcrum_GetNumberRebelds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/starwars.proto",
}
