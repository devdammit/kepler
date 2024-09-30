// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: kepler/kepler/tx.proto

package kepler

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

const (
	Msg_UpdateParams_FullMethodName   = "/kepler.kepler.Msg/UpdateParams"
	Msg_CreateW3Func_FullMethodName   = "/kepler.kepler.Msg/CreateW3Func"
	Msg_W3FuncExecuted_FullMethodName = "/kepler.kepler.Msg/W3FuncExecuted"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateW3Func(ctx context.Context, in *MsgCreateW3Func, opts ...grpc.CallOption) (*MsgCreateW3FuncResponse, error)
	W3FuncExecuted(ctx context.Context, in *MsgW3FuncExecuted, opts ...grpc.CallOption) (*MsgW3FuncExecutedResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateW3Func(ctx context.Context, in *MsgCreateW3Func, opts ...grpc.CallOption) (*MsgCreateW3FuncResponse, error) {
	out := new(MsgCreateW3FuncResponse)
	err := c.cc.Invoke(ctx, Msg_CreateW3Func_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) W3FuncExecuted(ctx context.Context, in *MsgW3FuncExecuted, opts ...grpc.CallOption) (*MsgW3FuncExecutedResponse, error) {
	out := new(MsgW3FuncExecutedResponse)
	err := c.cc.Invoke(ctx, Msg_W3FuncExecuted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateW3Func(context.Context, *MsgCreateW3Func) (*MsgCreateW3FuncResponse, error)
	W3FuncExecuted(context.Context, *MsgW3FuncExecuted) (*MsgW3FuncExecutedResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateW3Func(context.Context, *MsgCreateW3Func) (*MsgCreateW3FuncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateW3Func not implemented")
}
func (UnimplementedMsgServer) W3FuncExecuted(context.Context, *MsgW3FuncExecuted) (*MsgW3FuncExecutedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method W3FuncExecuted not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateW3Func_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateW3Func)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateW3Func(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateW3Func_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateW3Func(ctx, req.(*MsgCreateW3Func))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_W3FuncExecuted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgW3FuncExecuted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).W3FuncExecuted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_W3FuncExecuted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).W3FuncExecuted(ctx, req.(*MsgW3FuncExecuted))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kepler.kepler.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateW3Func",
			Handler:    _Msg_CreateW3Func_Handler,
		},
		{
			MethodName: "W3FuncExecuted",
			Handler:    _Msg_W3FuncExecuted_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kepler/kepler/tx.proto",
}
