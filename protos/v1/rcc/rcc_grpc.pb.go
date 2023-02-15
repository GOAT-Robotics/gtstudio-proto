// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protos/rcc/v1/rcc.proto

package rcc

import (
	context "context"
	model "github.com/TeamDotworld/robotix-proto/protos/v1/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	AuthenticateAgent(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// GetAgentConfiguration returns the Agent configuration.
	GetAgentConfiguration(ctx context.Context, in *GetAgentConfigurationRequest, opts ...grpc.CallOption) (*GetAgentConfigurationResponse, error)
	// Health can be used to check if the Agent is running and its gRPC API is
	// available.
	Health(ctx context.Context, opts ...grpc.CallOption) (Agent_HealthClient, error)
	GetCommandRequestStream(ctx context.Context, in *GetCommandRequestStreamRequest, opts ...grpc.CallOption) (Agent_GetCommandRequestStreamClient, error)
	GetSubscriptionRequestStream(ctx context.Context, opts ...grpc.CallOption) (Agent_GetSubscriptionRequestStreamClient, error)
	// SendCommandResponse sends a response to a command request.
	SendCommandResponse(ctx context.Context, in *SendCommandResponseRequest, opts ...grpc.CallOption) (*SendCommandResponseResponse, error)
	SendCommandResponseStream(ctx context.Context, opts ...grpc.CallOption) (Agent_SendCommandResponseStreamClient, error)
	PostData(ctx context.Context, in *model.Datapoint, opts ...grpc.CallOption) (*PostDataResponse, error)
	PostMultiData(ctx context.Context, in *PostMultiDataRequest, opts ...grpc.CallOption) (*PostMultiDataResponse, error)
	// StreamData accepts a stream of data points. See PostData for information on
	// expected error conditions and codes.
	StreamData(ctx context.Context, opts ...grpc.CallOption) (Agent_StreamDataClient, error)
	ReceiveRosMessages(ctx context.Context, opts ...grpc.CallOption) (Agent_ReceiveRosMessagesClient, error)
	// Get rtc signal from Rcc server to agent
	GetRtcSignal(ctx context.Context, opts ...grpc.CallOption) (Agent_GetRtcSignalClient, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) AuthenticateAgent(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/v1.rcc.Agent/AuthenticateAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetAgentConfiguration(ctx context.Context, in *GetAgentConfigurationRequest, opts ...grpc.CallOption) (*GetAgentConfigurationResponse, error) {
	out := new(GetAgentConfigurationResponse)
	err := c.cc.Invoke(ctx, "/v1.rcc.Agent/GetAgentConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Health(ctx context.Context, opts ...grpc.CallOption) (Agent_HealthClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[0], "/v1.rcc.Agent/Health", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentHealthClient{stream}
	return x, nil
}

type Agent_HealthClient interface {
	Send(*HealthResponse) error
	Recv() (*HealthRequest, error)
	grpc.ClientStream
}

type agentHealthClient struct {
	grpc.ClientStream
}

func (x *agentHealthClient) Send(m *HealthResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentHealthClient) Recv() (*HealthRequest, error) {
	m := new(HealthRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) GetCommandRequestStream(ctx context.Context, in *GetCommandRequestStreamRequest, opts ...grpc.CallOption) (Agent_GetCommandRequestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[1], "/v1.rcc.Agent/GetCommandRequestStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentGetCommandRequestStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Agent_GetCommandRequestStreamClient interface {
	Recv() (*GetCommandRequestStreamResponse, error)
	grpc.ClientStream
}

type agentGetCommandRequestStreamClient struct {
	grpc.ClientStream
}

func (x *agentGetCommandRequestStreamClient) Recv() (*GetCommandRequestStreamResponse, error) {
	m := new(GetCommandRequestStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) GetSubscriptionRequestStream(ctx context.Context, opts ...grpc.CallOption) (Agent_GetSubscriptionRequestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[2], "/v1.rcc.Agent/GetSubscriptionRequestStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentGetSubscriptionRequestStreamClient{stream}
	return x, nil
}

type Agent_GetSubscriptionRequestStreamClient interface {
	Send(*GetSubscriptionRequestStreamResponse) error
	Recv() (*GetSubscriptionRequestStreamRequest, error)
	grpc.ClientStream
}

type agentGetSubscriptionRequestStreamClient struct {
	grpc.ClientStream
}

func (x *agentGetSubscriptionRequestStreamClient) Send(m *GetSubscriptionRequestStreamResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentGetSubscriptionRequestStreamClient) Recv() (*GetSubscriptionRequestStreamRequest, error) {
	m := new(GetSubscriptionRequestStreamRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) SendCommandResponse(ctx context.Context, in *SendCommandResponseRequest, opts ...grpc.CallOption) (*SendCommandResponseResponse, error) {
	out := new(SendCommandResponseResponse)
	err := c.cc.Invoke(ctx, "/v1.rcc.Agent/SendCommandResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) SendCommandResponseStream(ctx context.Context, opts ...grpc.CallOption) (Agent_SendCommandResponseStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[3], "/v1.rcc.Agent/SendCommandResponseStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentSendCommandResponseStreamClient{stream}
	return x, nil
}

type Agent_SendCommandResponseStreamClient interface {
	Send(*SendCommandResponseRequest) error
	Recv() (*SendCommandResponseResponse, error)
	grpc.ClientStream
}

type agentSendCommandResponseStreamClient struct {
	grpc.ClientStream
}

func (x *agentSendCommandResponseStreamClient) Send(m *SendCommandResponseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentSendCommandResponseStreamClient) Recv() (*SendCommandResponseResponse, error) {
	m := new(SendCommandResponseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) PostData(ctx context.Context, in *model.Datapoint, opts ...grpc.CallOption) (*PostDataResponse, error) {
	out := new(PostDataResponse)
	err := c.cc.Invoke(ctx, "/v1.rcc.Agent/PostData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) PostMultiData(ctx context.Context, in *PostMultiDataRequest, opts ...grpc.CallOption) (*PostMultiDataResponse, error) {
	out := new(PostMultiDataResponse)
	err := c.cc.Invoke(ctx, "/v1.rcc.Agent/PostMultiData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StreamData(ctx context.Context, opts ...grpc.CallOption) (Agent_StreamDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[4], "/v1.rcc.Agent/StreamData", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentStreamDataClient{stream}
	return x, nil
}

type Agent_StreamDataClient interface {
	Send(*model.Datapoint) error
	CloseAndRecv() (*StreamDataResponse, error)
	grpc.ClientStream
}

type agentStreamDataClient struct {
	grpc.ClientStream
}

func (x *agentStreamDataClient) Send(m *model.Datapoint) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentStreamDataClient) CloseAndRecv() (*StreamDataResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) ReceiveRosMessages(ctx context.Context, opts ...grpc.CallOption) (Agent_ReceiveRosMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[5], "/v1.rcc.Agent/ReceiveRosMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentReceiveRosMessagesClient{stream}
	return x, nil
}

type Agent_ReceiveRosMessagesClient interface {
	Send(*RecieveRosMessagesResponse) error
	Recv() (*RecieveRosMessagesRequest, error)
	grpc.ClientStream
}

type agentReceiveRosMessagesClient struct {
	grpc.ClientStream
}

func (x *agentReceiveRosMessagesClient) Send(m *RecieveRosMessagesResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentReceiveRosMessagesClient) Recv() (*RecieveRosMessagesRequest, error) {
	m := new(RecieveRosMessagesRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) GetRtcSignal(ctx context.Context, opts ...grpc.CallOption) (Agent_GetRtcSignalClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[6], "/v1.rcc.Agent/GetRtcSignal", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentGetRtcSignalClient{stream}
	return x, nil
}

type Agent_GetRtcSignalClient interface {
	Send(*GetRtcSignalResponse) error
	Recv() (*GetRtcSignalRequest, error)
	grpc.ClientStream
}

type agentGetRtcSignalClient struct {
	grpc.ClientStream
}

func (x *agentGetRtcSignalClient) Send(m *GetRtcSignalResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentGetRtcSignalClient) Recv() (*GetRtcSignalRequest, error) {
	m := new(GetRtcSignalRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	AuthenticateAgent(context.Context, *AuthRequest) (*AuthResponse, error)
	// GetAgentConfiguration returns the Agent configuration.
	GetAgentConfiguration(context.Context, *GetAgentConfigurationRequest) (*GetAgentConfigurationResponse, error)
	// Health can be used to check if the Agent is running and its gRPC API is
	// available.
	Health(Agent_HealthServer) error
	GetCommandRequestStream(*GetCommandRequestStreamRequest, Agent_GetCommandRequestStreamServer) error
	GetSubscriptionRequestStream(Agent_GetSubscriptionRequestStreamServer) error
	// SendCommandResponse sends a response to a command request.
	SendCommandResponse(context.Context, *SendCommandResponseRequest) (*SendCommandResponseResponse, error)
	SendCommandResponseStream(Agent_SendCommandResponseStreamServer) error
	PostData(context.Context, *model.Datapoint) (*PostDataResponse, error)
	PostMultiData(context.Context, *PostMultiDataRequest) (*PostMultiDataResponse, error)
	// StreamData accepts a stream of data points. See PostData for information on
	// expected error conditions and codes.
	StreamData(Agent_StreamDataServer) error
	ReceiveRosMessages(Agent_ReceiveRosMessagesServer) error
	// Get rtc signal from Rcc server to agent
	GetRtcSignal(Agent_GetRtcSignalServer) error
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) AuthenticateAgent(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateAgent not implemented")
}
func (UnimplementedAgentServer) GetAgentConfiguration(context.Context, *GetAgentConfigurationRequest) (*GetAgentConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgentConfiguration not implemented")
}
func (UnimplementedAgentServer) Health(Agent_HealthServer) error {
	return status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedAgentServer) GetCommandRequestStream(*GetCommandRequestStreamRequest, Agent_GetCommandRequestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCommandRequestStream not implemented")
}
func (UnimplementedAgentServer) GetSubscriptionRequestStream(Agent_GetSubscriptionRequestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSubscriptionRequestStream not implemented")
}
func (UnimplementedAgentServer) SendCommandResponse(context.Context, *SendCommandResponseRequest) (*SendCommandResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommandResponse not implemented")
}
func (UnimplementedAgentServer) SendCommandResponseStream(Agent_SendCommandResponseStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SendCommandResponseStream not implemented")
}
func (UnimplementedAgentServer) PostData(context.Context, *model.Datapoint) (*PostDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostData not implemented")
}
func (UnimplementedAgentServer) PostMultiData(context.Context, *PostMultiDataRequest) (*PostMultiDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMultiData not implemented")
}
func (UnimplementedAgentServer) StreamData(Agent_StreamDataServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamData not implemented")
}
func (UnimplementedAgentServer) ReceiveRosMessages(Agent_ReceiveRosMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveRosMessages not implemented")
}
func (UnimplementedAgentServer) GetRtcSignal(Agent_GetRtcSignalServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRtcSignal not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_AuthenticateAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).AuthenticateAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.rcc.Agent/AuthenticateAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).AuthenticateAgent(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetAgentConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetAgentConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.rcc.Agent/GetAgentConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetAgentConfiguration(ctx, req.(*GetAgentConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Health_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).Health(&agentHealthServer{stream})
}

type Agent_HealthServer interface {
	Send(*HealthRequest) error
	Recv() (*HealthResponse, error)
	grpc.ServerStream
}

type agentHealthServer struct {
	grpc.ServerStream
}

func (x *agentHealthServer) Send(m *HealthRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentHealthServer) Recv() (*HealthResponse, error) {
	m := new(HealthResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Agent_GetCommandRequestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCommandRequestStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServer).GetCommandRequestStream(m, &agentGetCommandRequestStreamServer{stream})
}

type Agent_GetCommandRequestStreamServer interface {
	Send(*GetCommandRequestStreamResponse) error
	grpc.ServerStream
}

type agentGetCommandRequestStreamServer struct {
	grpc.ServerStream
}

func (x *agentGetCommandRequestStreamServer) Send(m *GetCommandRequestStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Agent_GetSubscriptionRequestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).GetSubscriptionRequestStream(&agentGetSubscriptionRequestStreamServer{stream})
}

type Agent_GetSubscriptionRequestStreamServer interface {
	Send(*GetSubscriptionRequestStreamRequest) error
	Recv() (*GetSubscriptionRequestStreamResponse, error)
	grpc.ServerStream
}

type agentGetSubscriptionRequestStreamServer struct {
	grpc.ServerStream
}

func (x *agentGetSubscriptionRequestStreamServer) Send(m *GetSubscriptionRequestStreamRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentGetSubscriptionRequestStreamServer) Recv() (*GetSubscriptionRequestStreamResponse, error) {
	m := new(GetSubscriptionRequestStreamResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Agent_SendCommandResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCommandResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SendCommandResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.rcc.Agent/SendCommandResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SendCommandResponse(ctx, req.(*SendCommandResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_SendCommandResponseStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).SendCommandResponseStream(&agentSendCommandResponseStreamServer{stream})
}

type Agent_SendCommandResponseStreamServer interface {
	Send(*SendCommandResponseResponse) error
	Recv() (*SendCommandResponseRequest, error)
	grpc.ServerStream
}

type agentSendCommandResponseStreamServer struct {
	grpc.ServerStream
}

func (x *agentSendCommandResponseStreamServer) Send(m *SendCommandResponseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentSendCommandResponseStreamServer) Recv() (*SendCommandResponseRequest, error) {
	m := new(SendCommandResponseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Agent_PostData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Datapoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).PostData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.rcc.Agent/PostData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).PostData(ctx, req.(*model.Datapoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_PostMultiData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMultiDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).PostMultiData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.rcc.Agent/PostMultiData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).PostMultiData(ctx, req.(*PostMultiDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StreamData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).StreamData(&agentStreamDataServer{stream})
}

type Agent_StreamDataServer interface {
	SendAndClose(*StreamDataResponse) error
	Recv() (*model.Datapoint, error)
	grpc.ServerStream
}

type agentStreamDataServer struct {
	grpc.ServerStream
}

func (x *agentStreamDataServer) SendAndClose(m *StreamDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentStreamDataServer) Recv() (*model.Datapoint, error) {
	m := new(model.Datapoint)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Agent_ReceiveRosMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).ReceiveRosMessages(&agentReceiveRosMessagesServer{stream})
}

type Agent_ReceiveRosMessagesServer interface {
	Send(*RecieveRosMessagesRequest) error
	Recv() (*RecieveRosMessagesResponse, error)
	grpc.ServerStream
}

type agentReceiveRosMessagesServer struct {
	grpc.ServerStream
}

func (x *agentReceiveRosMessagesServer) Send(m *RecieveRosMessagesRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentReceiveRosMessagesServer) Recv() (*RecieveRosMessagesResponse, error) {
	m := new(RecieveRosMessagesResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Agent_GetRtcSignal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).GetRtcSignal(&agentGetRtcSignalServer{stream})
}

type Agent_GetRtcSignalServer interface {
	Send(*GetRtcSignalRequest) error
	Recv() (*GetRtcSignalResponse, error)
	grpc.ServerStream
}

type agentGetRtcSignalServer struct {
	grpc.ServerStream
}

func (x *agentGetRtcSignalServer) Send(m *GetRtcSignalRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentGetRtcSignalServer) Recv() (*GetRtcSignalResponse, error) {
	m := new(GetRtcSignalResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.rcc.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticateAgent",
			Handler:    _Agent_AuthenticateAgent_Handler,
		},
		{
			MethodName: "GetAgentConfiguration",
			Handler:    _Agent_GetAgentConfiguration_Handler,
		},
		{
			MethodName: "SendCommandResponse",
			Handler:    _Agent_SendCommandResponse_Handler,
		},
		{
			MethodName: "PostData",
			Handler:    _Agent_PostData_Handler,
		},
		{
			MethodName: "PostMultiData",
			Handler:    _Agent_PostMultiData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Health",
			Handler:       _Agent_Health_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetCommandRequestStream",
			Handler:       _Agent_GetCommandRequestStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSubscriptionRequestStream",
			Handler:       _Agent_GetSubscriptionRequestStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SendCommandResponseStream",
			Handler:       _Agent_SendCommandResponseStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamData",
			Handler:       _Agent_StreamData_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveRosMessages",
			Handler:       _Agent_ReceiveRosMessages_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetRtcSignal",
			Handler:       _Agent_GetRtcSignal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/rcc/v1/rcc.proto",
}
