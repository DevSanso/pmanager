// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: private.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	request "pkg/message/request"
	response "pkg/message/response"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PrivateServiceClient is the client API for PrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivateServiceClient interface {
	ResetManagerClient(ctx context.Context, in *request.ResetManagerClientRequest, opts ...grpc.CallOption) (*response.ResetManagerClientResponse, error)
	SwitchServerOnoff(ctx context.Context, in *request.SwitchServerOnOffRequest, opts ...grpc.CallOption) (*response.SwitchServerOnOffResponse, error)
	SetPublicAccessInfo(ctx context.Context, in *request.SetPublicAccessInfoRequest, opts ...grpc.CallOption) (*response.SetPublicAccessInfoResponse, error)
	SetModuleRebootCond(ctx context.Context, in *request.SetModuleRebootCondRequest, opts ...grpc.CallOption) (*response.SetModuleRebootCondResponse, error)
	SetServerGroup(ctx context.Context, in *request.SetServerGroupRequest, opts ...grpc.CallOption) (*response.SetServerGroupResponse, error)
	SetService(ctx context.Context, in *request.SetServiceRequest, opts ...grpc.CallOption) (*response.SetServiceResponse, error)
	InitServer(ctx context.Context, in *request.InitServerRequest, opts ...grpc.CallOption) (*response.InitServerResponse, error)
	PushModule(ctx context.Context, in *request.PushModuleRequest, opts ...grpc.CallOption) (*response.PushModuleResponse, error)
	PopModule(ctx context.Context, in *request.PopModuleRequest, opts ...grpc.CallOption) (*response.PopModuleResponse, error)
}

type privateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivateServiceClient(cc grpc.ClientConnInterface) PrivateServiceClient {
	return &privateServiceClient{cc}
}

func (c *privateServiceClient) ResetManagerClient(ctx context.Context, in *request.ResetManagerClientRequest, opts ...grpc.CallOption) (*response.ResetManagerClientResponse, error) {
	out := new(response.ResetManagerClientResponse)
	err := c.cc.Invoke(ctx, "/grpc.PrivateService/ResetManagerClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) SwitchServerOnoff(ctx context.Context, in *request.SwitchServerOnOffRequest, opts ...grpc.CallOption) (*response.SwitchServerOnOffResponse, error) {
	out := new(response.SwitchServerOnOffResponse)
	err := c.cc.Invoke(ctx, "/grpc.PrivateService/SwitchServerOnoff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) SetPublicAccessInfo(ctx context.Context, in *request.SetPublicAccessInfoRequest, opts ...grpc.CallOption) (*response.SetPublicAccessInfoResponse, error) {
	out := new(response.SetPublicAccessInfoResponse)
	err := c.cc.Invoke(ctx, "/grpc.PrivateService/SetPublicAccessInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) SetModuleRebootCond(ctx context.Context, in *request.SetModuleRebootCondRequest, opts ...grpc.CallOption) (*response.SetModuleRebootCondResponse, error) {
	out := new(response.SetModuleRebootCondResponse)
	err := c.cc.Invoke(ctx, "/grpc.PrivateService/SetModuleRebootCond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) SetServerGroup(ctx context.Context, in *request.SetServerGroupRequest, opts ...grpc.CallOption) (*response.SetServerGroupResponse, error) {
	out := new(response.SetServerGroupResponse)
	err := c.cc.Invoke(ctx, "/grpc.PrivateService/SetServerGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) SetService(ctx context.Context, in *request.SetServiceRequest, opts ...grpc.CallOption) (*response.SetServiceResponse, error) {
	out := new(response.SetServiceResponse)
	err := c.cc.Invoke(ctx, "/grpc.PrivateService/SetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) InitServer(ctx context.Context, in *request.InitServerRequest, opts ...grpc.CallOption) (*response.InitServerResponse, error) {
	out := new(response.InitServerResponse)
	err := c.cc.Invoke(ctx, "/grpc.PrivateService/InitServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) PushModule(ctx context.Context, in *request.PushModuleRequest, opts ...grpc.CallOption) (*response.PushModuleResponse, error) {
	out := new(response.PushModuleResponse)
	err := c.cc.Invoke(ctx, "/grpc.PrivateService/PushModule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) PopModule(ctx context.Context, in *request.PopModuleRequest, opts ...grpc.CallOption) (*response.PopModuleResponse, error) {
	out := new(response.PopModuleResponse)
	err := c.cc.Invoke(ctx, "/grpc.PrivateService/PopModule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateServiceServer is the server API for PrivateService service.
// All implementations must embed UnimplementedPrivateServiceServer
// for forward compatibility
type PrivateServiceServer interface {
	ResetManagerClient(context.Context, *request.ResetManagerClientRequest) (*response.ResetManagerClientResponse, error)
	SwitchServerOnoff(context.Context, *request.SwitchServerOnOffRequest) (*response.SwitchServerOnOffResponse, error)
	SetPublicAccessInfo(context.Context, *request.SetPublicAccessInfoRequest) (*response.SetPublicAccessInfoResponse, error)
	SetModuleRebootCond(context.Context, *request.SetModuleRebootCondRequest) (*response.SetModuleRebootCondResponse, error)
	SetServerGroup(context.Context, *request.SetServerGroupRequest) (*response.SetServerGroupResponse, error)
	SetService(context.Context, *request.SetServiceRequest) (*response.SetServiceResponse, error)
	InitServer(context.Context, *request.InitServerRequest) (*response.InitServerResponse, error)
	PushModule(context.Context, *request.PushModuleRequest) (*response.PushModuleResponse, error)
	PopModule(context.Context, *request.PopModuleRequest) (*response.PopModuleResponse, error)
	mustEmbedUnimplementedPrivateServiceServer()
}

// UnimplementedPrivateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrivateServiceServer struct {
}

func (UnimplementedPrivateServiceServer) ResetManagerClient(context.Context, *request.ResetManagerClientRequest) (*response.ResetManagerClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetManagerClient not implemented")
}
func (UnimplementedPrivateServiceServer) SwitchServerOnoff(context.Context, *request.SwitchServerOnOffRequest) (*response.SwitchServerOnOffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchServerOnoff not implemented")
}
func (UnimplementedPrivateServiceServer) SetPublicAccessInfo(context.Context, *request.SetPublicAccessInfoRequest) (*response.SetPublicAccessInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPublicAccessInfo not implemented")
}
func (UnimplementedPrivateServiceServer) SetModuleRebootCond(context.Context, *request.SetModuleRebootCondRequest) (*response.SetModuleRebootCondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetModuleRebootCond not implemented")
}
func (UnimplementedPrivateServiceServer) SetServerGroup(context.Context, *request.SetServerGroupRequest) (*response.SetServerGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetServerGroup not implemented")
}
func (UnimplementedPrivateServiceServer) SetService(context.Context, *request.SetServiceRequest) (*response.SetServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetService not implemented")
}
func (UnimplementedPrivateServiceServer) InitServer(context.Context, *request.InitServerRequest) (*response.InitServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitServer not implemented")
}
func (UnimplementedPrivateServiceServer) PushModule(context.Context, *request.PushModuleRequest) (*response.PushModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushModule not implemented")
}
func (UnimplementedPrivateServiceServer) PopModule(context.Context, *request.PopModuleRequest) (*response.PopModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PopModule not implemented")
}
func (UnimplementedPrivateServiceServer) mustEmbedUnimplementedPrivateServiceServer() {}

// UnsafePrivateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivateServiceServer will
// result in compilation errors.
type UnsafePrivateServiceServer interface {
	mustEmbedUnimplementedPrivateServiceServer()
}

func RegisterPrivateServiceServer(s grpc.ServiceRegistrar, srv PrivateServiceServer) {
	s.RegisterService(&PrivateService_ServiceDesc, srv)
}

func _PrivateService_ResetManagerClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ResetManagerClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).ResetManagerClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PrivateService/ResetManagerClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).ResetManagerClient(ctx, req.(*request.ResetManagerClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_SwitchServerOnoff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SwitchServerOnOffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).SwitchServerOnoff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PrivateService/SwitchServerOnoff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).SwitchServerOnoff(ctx, req.(*request.SwitchServerOnOffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_SetPublicAccessInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SetPublicAccessInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).SetPublicAccessInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PrivateService/SetPublicAccessInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).SetPublicAccessInfo(ctx, req.(*request.SetPublicAccessInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_SetModuleRebootCond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SetModuleRebootCondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).SetModuleRebootCond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PrivateService/SetModuleRebootCond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).SetModuleRebootCond(ctx, req.(*request.SetModuleRebootCondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_SetServerGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SetServerGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).SetServerGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PrivateService/SetServerGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).SetServerGroup(ctx, req.(*request.SetServerGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_SetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).SetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PrivateService/SetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).SetService(ctx, req.(*request.SetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_InitServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.InitServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).InitServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PrivateService/InitServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).InitServer(ctx, req.(*request.InitServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_PushModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.PushModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).PushModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PrivateService/PushModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).PushModule(ctx, req.(*request.PushModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_PopModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.PopModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).PopModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PrivateService/PopModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).PopModule(ctx, req.(*request.PopModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrivateService_ServiceDesc is the grpc.ServiceDesc for PrivateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.PrivateService",
	HandlerType: (*PrivateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResetManagerClient",
			Handler:    _PrivateService_ResetManagerClient_Handler,
		},
		{
			MethodName: "SwitchServerOnoff",
			Handler:    _PrivateService_SwitchServerOnoff_Handler,
		},
		{
			MethodName: "SetPublicAccessInfo",
			Handler:    _PrivateService_SetPublicAccessInfo_Handler,
		},
		{
			MethodName: "SetModuleRebootCond",
			Handler:    _PrivateService_SetModuleRebootCond_Handler,
		},
		{
			MethodName: "SetServerGroup",
			Handler:    _PrivateService_SetServerGroup_Handler,
		},
		{
			MethodName: "SetService",
			Handler:    _PrivateService_SetService_Handler,
		},
		{
			MethodName: "InitServer",
			Handler:    _PrivateService_InitServer_Handler,
		},
		{
			MethodName: "PushModule",
			Handler:    _PrivateService_PushModule_Handler,
		},
		{
			MethodName: "PopModule",
			Handler:    _PrivateService_PopModule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "private.proto",
}
