// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: sistema.proto

package proto

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
	GobiernoService_ObtenerListaPiratas_FullMethodName = "/sistema.GobiernoService/ObtenerListaPiratas"
	GobiernoService_ConfirmarEntrega_FullMethodName    = "/sistema.GobiernoService/ConfirmarEntrega"
	GobiernoService_ConsultarReputacion_FullMethodName = "/sistema.GobiernoService/ConsultarReputacion"
	GobiernoService_ObtenerEstadoGlobal_FullMethodName = "/sistema.GobiernoService/ObtenerEstadoGlobal"
	GobiernoService_ReporteDeEstado_FullMethodName     = "/sistema.GobiernoService/ReporteDeEstado"
)

// GobiernoServiceClient is the client API for GobiernoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GobiernoServiceClient interface {
	ObtenerListaPiratas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListaPiratas, error)
	ConfirmarEntrega(ctx context.Context, in *ConfirmarEntregaRequest, opts ...grpc.CallOption) (*ResultadoOperacion, error)
	ConsultarReputacion(ctx context.Context, in *ReputacionRequest, opts ...grpc.CallOption) (*ReputacionResponse, error)
	ObtenerEstadoGlobal(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EstadoGlobalResponse, error)
	ReporteDeEstado(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (*ResultadoOperacion, error)
}

type gobiernoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGobiernoServiceClient(cc grpc.ClientConnInterface) GobiernoServiceClient {
	return &gobiernoServiceClient{cc}
}

func (c *gobiernoServiceClient) ObtenerListaPiratas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListaPiratas, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListaPiratas)
	err := c.cc.Invoke(ctx, GobiernoService_ObtenerListaPiratas_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gobiernoServiceClient) ConfirmarEntrega(ctx context.Context, in *ConfirmarEntregaRequest, opts ...grpc.CallOption) (*ResultadoOperacion, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResultadoOperacion)
	err := c.cc.Invoke(ctx, GobiernoService_ConfirmarEntrega_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gobiernoServiceClient) ConsultarReputacion(ctx context.Context, in *ReputacionRequest, opts ...grpc.CallOption) (*ReputacionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReputacionResponse)
	err := c.cc.Invoke(ctx, GobiernoService_ConsultarReputacion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gobiernoServiceClient) ObtenerEstadoGlobal(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EstadoGlobalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EstadoGlobalResponse)
	err := c.cc.Invoke(ctx, GobiernoService_ObtenerEstadoGlobal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gobiernoServiceClient) ReporteDeEstado(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (*ResultadoOperacion, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResultadoOperacion)
	err := c.cc.Invoke(ctx, GobiernoService_ReporteDeEstado_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GobiernoServiceServer is the server API for GobiernoService service.
// All implementations must embed UnimplementedGobiernoServiceServer
// for forward compatibility.
type GobiernoServiceServer interface {
	ObtenerListaPiratas(context.Context, *Empty) (*ListaPiratas, error)
	ConfirmarEntrega(context.Context, *ConfirmarEntregaRequest) (*ResultadoOperacion, error)
	ConsultarReputacion(context.Context, *ReputacionRequest) (*ReputacionResponse, error)
	ObtenerEstadoGlobal(context.Context, *Empty) (*EstadoGlobalResponse, error)
	ReporteDeEstado(context.Context, *EstadoRequest) (*ResultadoOperacion, error)
	mustEmbedUnimplementedGobiernoServiceServer()
}

// UnimplementedGobiernoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGobiernoServiceServer struct{}

func (UnimplementedGobiernoServiceServer) ObtenerListaPiratas(context.Context, *Empty) (*ListaPiratas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerListaPiratas not implemented")
}
func (UnimplementedGobiernoServiceServer) ConfirmarEntrega(context.Context, *ConfirmarEntregaRequest) (*ResultadoOperacion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmarEntrega not implemented")
}
func (UnimplementedGobiernoServiceServer) ConsultarReputacion(context.Context, *ReputacionRequest) (*ReputacionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultarReputacion not implemented")
}
func (UnimplementedGobiernoServiceServer) ObtenerEstadoGlobal(context.Context, *Empty) (*EstadoGlobalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerEstadoGlobal not implemented")
}
func (UnimplementedGobiernoServiceServer) ReporteDeEstado(context.Context, *EstadoRequest) (*ResultadoOperacion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReporteDeEstado not implemented")
}
func (UnimplementedGobiernoServiceServer) mustEmbedUnimplementedGobiernoServiceServer() {}
func (UnimplementedGobiernoServiceServer) testEmbeddedByValue()                         {}

// UnsafeGobiernoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GobiernoServiceServer will
// result in compilation errors.
type UnsafeGobiernoServiceServer interface {
	mustEmbedUnimplementedGobiernoServiceServer()
}

func RegisterGobiernoServiceServer(s grpc.ServiceRegistrar, srv GobiernoServiceServer) {
	// If the following call pancis, it indicates UnimplementedGobiernoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GobiernoService_ServiceDesc, srv)
}

func _GobiernoService_ObtenerListaPiratas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GobiernoServiceServer).ObtenerListaPiratas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GobiernoService_ObtenerListaPiratas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GobiernoServiceServer).ObtenerListaPiratas(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GobiernoService_ConfirmarEntrega_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmarEntregaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GobiernoServiceServer).ConfirmarEntrega(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GobiernoService_ConfirmarEntrega_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GobiernoServiceServer).ConfirmarEntrega(ctx, req.(*ConfirmarEntregaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GobiernoService_ConsultarReputacion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReputacionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GobiernoServiceServer).ConsultarReputacion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GobiernoService_ConsultarReputacion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GobiernoServiceServer).ConsultarReputacion(ctx, req.(*ReputacionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GobiernoService_ObtenerEstadoGlobal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GobiernoServiceServer).ObtenerEstadoGlobal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GobiernoService_ObtenerEstadoGlobal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GobiernoServiceServer).ObtenerEstadoGlobal(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GobiernoService_ReporteDeEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstadoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GobiernoServiceServer).ReporteDeEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GobiernoService_ReporteDeEstado_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GobiernoServiceServer).ReporteDeEstado(ctx, req.(*EstadoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GobiernoService_ServiceDesc is the grpc.ServiceDesc for GobiernoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GobiernoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sistema.GobiernoService",
	HandlerType: (*GobiernoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ObtenerListaPiratas",
			Handler:    _GobiernoService_ObtenerListaPiratas_Handler,
		},
		{
			MethodName: "ConfirmarEntrega",
			Handler:    _GobiernoService_ConfirmarEntrega_Handler,
		},
		{
			MethodName: "ConsultarReputacion",
			Handler:    _GobiernoService_ConsultarReputacion_Handler,
		},
		{
			MethodName: "ObtenerEstadoGlobal",
			Handler:    _GobiernoService_ObtenerEstadoGlobal_Handler,
		},
		{
			MethodName: "ReporteDeEstado",
			Handler:    _GobiernoService_ReporteDeEstado_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sistema.proto",
}

const (
	MarinaService_EntregarPirataMarina_FullMethodName = "/sistema.MarinaService/EntregarPirataMarina"
	MarinaService_AlertaTraficoIlegal_FullMethodName  = "/sistema.MarinaService/AlertaTraficoIlegal"
)

// MarinaServiceClient is the client API for MarinaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarinaServiceClient interface {
	EntregarPirataMarina(ctx context.Context, in *EntregaRequest, opts ...grpc.CallOption) (*ResultadoOperacion, error)
	AlertaTraficoIlegal(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResultadoOperacion, error)
}

type marinaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarinaServiceClient(cc grpc.ClientConnInterface) MarinaServiceClient {
	return &marinaServiceClient{cc}
}

func (c *marinaServiceClient) EntregarPirataMarina(ctx context.Context, in *EntregaRequest, opts ...grpc.CallOption) (*ResultadoOperacion, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResultadoOperacion)
	err := c.cc.Invoke(ctx, MarinaService_EntregarPirataMarina_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marinaServiceClient) AlertaTraficoIlegal(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResultadoOperacion, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResultadoOperacion)
	err := c.cc.Invoke(ctx, MarinaService_AlertaTraficoIlegal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarinaServiceServer is the server API for MarinaService service.
// All implementations must embed UnimplementedMarinaServiceServer
// for forward compatibility.
type MarinaServiceServer interface {
	EntregarPirataMarina(context.Context, *EntregaRequest) (*ResultadoOperacion, error)
	AlertaTraficoIlegal(context.Context, *Empty) (*ResultadoOperacion, error)
	mustEmbedUnimplementedMarinaServiceServer()
}

// UnimplementedMarinaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMarinaServiceServer struct{}

func (UnimplementedMarinaServiceServer) EntregarPirataMarina(context.Context, *EntregaRequest) (*ResultadoOperacion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntregarPirataMarina not implemented")
}
func (UnimplementedMarinaServiceServer) AlertaTraficoIlegal(context.Context, *Empty) (*ResultadoOperacion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlertaTraficoIlegal not implemented")
}
func (UnimplementedMarinaServiceServer) mustEmbedUnimplementedMarinaServiceServer() {}
func (UnimplementedMarinaServiceServer) testEmbeddedByValue()                       {}

// UnsafeMarinaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarinaServiceServer will
// result in compilation errors.
type UnsafeMarinaServiceServer interface {
	mustEmbedUnimplementedMarinaServiceServer()
}

func RegisterMarinaServiceServer(s grpc.ServiceRegistrar, srv MarinaServiceServer) {
	// If the following call pancis, it indicates UnimplementedMarinaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MarinaService_ServiceDesc, srv)
}

func _MarinaService_EntregarPirataMarina_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntregaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarinaServiceServer).EntregarPirataMarina(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarinaService_EntregarPirataMarina_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarinaServiceServer).EntregarPirataMarina(ctx, req.(*EntregaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarinaService_AlertaTraficoIlegal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarinaServiceServer).AlertaTraficoIlegal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarinaService_AlertaTraficoIlegal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarinaServiceServer).AlertaTraficoIlegal(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MarinaService_ServiceDesc is the grpc.ServiceDesc for MarinaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarinaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sistema.MarinaService",
	HandlerType: (*MarinaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EntregarPirataMarina",
			Handler:    _MarinaService_EntregarPirataMarina_Handler,
		},
		{
			MethodName: "AlertaTraficoIlegal",
			Handler:    _MarinaService_AlertaTraficoIlegal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sistema.proto",
}

const (
	SubmundoService_EntregarPirataSubmundo_FullMethodName = "/sistema.SubmundoService/EntregarPirataSubmundo"
)

// SubmundoServiceClient is the client API for SubmundoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmundoServiceClient interface {
	EntregarPirataSubmundo(ctx context.Context, in *VentaRequest, opts ...grpc.CallOption) (*ResultadoOperacion, error)
}

type submundoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmundoServiceClient(cc grpc.ClientConnInterface) SubmundoServiceClient {
	return &submundoServiceClient{cc}
}

func (c *submundoServiceClient) EntregarPirataSubmundo(ctx context.Context, in *VentaRequest, opts ...grpc.CallOption) (*ResultadoOperacion, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResultadoOperacion)
	err := c.cc.Invoke(ctx, SubmundoService_EntregarPirataSubmundo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmundoServiceServer is the server API for SubmundoService service.
// All implementations must embed UnimplementedSubmundoServiceServer
// for forward compatibility.
type SubmundoServiceServer interface {
	EntregarPirataSubmundo(context.Context, *VentaRequest) (*ResultadoOperacion, error)
	mustEmbedUnimplementedSubmundoServiceServer()
}

// UnimplementedSubmundoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSubmundoServiceServer struct{}

func (UnimplementedSubmundoServiceServer) EntregarPirataSubmundo(context.Context, *VentaRequest) (*ResultadoOperacion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntregarPirataSubmundo not implemented")
}
func (UnimplementedSubmundoServiceServer) mustEmbedUnimplementedSubmundoServiceServer() {}
func (UnimplementedSubmundoServiceServer) testEmbeddedByValue()                         {}

// UnsafeSubmundoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubmundoServiceServer will
// result in compilation errors.
type UnsafeSubmundoServiceServer interface {
	mustEmbedUnimplementedSubmundoServiceServer()
}

func RegisterSubmundoServiceServer(s grpc.ServiceRegistrar, srv SubmundoServiceServer) {
	// If the following call pancis, it indicates UnimplementedSubmundoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SubmundoService_ServiceDesc, srv)
}

func _SubmundoService_EntregarPirataSubmundo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VentaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmundoServiceServer).EntregarPirataSubmundo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmundoService_EntregarPirataSubmundo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmundoServiceServer).EntregarPirataSubmundo(ctx, req.(*VentaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubmundoService_ServiceDesc is the grpc.ServiceDesc for SubmundoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubmundoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sistema.SubmundoService",
	HandlerType: (*SubmundoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EntregarPirataSubmundo",
			Handler:    _SubmundoService_EntregarPirataSubmundo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sistema.proto",
}
