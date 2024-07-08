// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
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
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Agent_AuthenticateAgent_FullMethodName            = "/v1.rcc.Agent/AuthenticateAgent"
	Agent_GetDeviceSpecifications_FullMethodName      = "/v1.rcc.Agent/GetDeviceSpecifications"
	Agent_GetAgentConfiguration_FullMethodName        = "/v1.rcc.Agent/GetAgentConfiguration"
	Agent_Health_FullMethodName                       = "/v1.rcc.Agent/Health"
	Agent_GetCommandRequestStream_FullMethodName      = "/v1.rcc.Agent/GetCommandRequestStream"
	Agent_SendCommandResponse_FullMethodName          = "/v1.rcc.Agent/SendCommandResponse"
	Agent_GetSubscriptionRequestStream_FullMethodName = "/v1.rcc.Agent/GetSubscriptionRequestStream"
	Agent_SendSubscriptionResponse_FullMethodName     = "/v1.rcc.Agent/SendSubscriptionResponse"
	Agent_SendCommandResponseStream_FullMethodName    = "/v1.rcc.Agent/SendCommandResponseStream"
	Agent_PostData_FullMethodName                     = "/v1.rcc.Agent/PostData"
	Agent_PostMultiData_FullMethodName                = "/v1.rcc.Agent/PostMultiData"
	Agent_StreamData_FullMethodName                   = "/v1.rcc.Agent/StreamData"
	Agent_ReceiveRosMessages_FullMethodName           = "/v1.rcc.Agent/ReceiveRosMessages"
	Agent_GetRtcSignal_FullMethodName                 = "/v1.rcc.Agent/GetRtcSignal"
	Agent_SendSignalData_FullMethodName               = "/v1.rcc.Agent/SendSignalData"
)

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	// AuthenticateAgent service authenticates the agent and return a boolean
	// status which determines the authentication status of the agent
	AuthenticateAgent(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// GetDeviceSpecifications returns the device specification details.
	GetDeviceSpecifications(ctx context.Context, in *SpecsRequest, opts ...grpc.CallOption) (*SpecsResponse, error)
	// GetAgentConfiguration returns the Agent configuration.
	GetAgentConfiguration(ctx context.Context, in *GetAgentConfigurationRequest, opts ...grpc.CallOption) (*GetAgentConfigurationResponse, error)
	// Heath service is established between agent and server.When the connection
	// establises the status for the deivce changes
	// to ONLINE. This constantly checks for connection. If the connection is
	// disconnected the device status changes to OFFLINE.
	Health(ctx context.Context, opts ...grpc.CallOption) (Agent_HealthClient, error)
	// GetCommandRequestStream service receives command from the server both
	// "SERVICE" and "ACTION" commands. The Command should be configured in the
	// agent's config first. Unconfigured command is not accepted
	GetCommandRequestStream(ctx context.Context, in *GetCommandRequestStreamRequest, opts ...grpc.CallOption) (Agent_GetCommandRequestStreamClient, error)
	// SendCommandResponse service sends command response (Just response if it is
	// service call, if Action call it sends both result and feedback ) for a
	// particular command to the server.
	SendCommandResponse(ctx context.Context, in *SendCommandResponseRequest, opts ...grpc.CallOption) (*SendCommandResponseResponse, error)
	// GetSubscriptionRequestStream service recieves subscription related commands
	// such as "SUBSCRIBE", "UNSUBSCRIBE", "GET_TOPICS", "GET_SUBSCRIBED_TOPICS" and
	// sends back the response to the server. This is previously implemented for
	// node sdk with grpc, and not work with current implementation.
	GetSubscriptionRequestStream(ctx context.Context, in *GetSubscriptionRequestStreamRequest, opts ...grpc.CallOption) (Agent_GetSubscriptionRequestStreamClient, error)
	// SendSubscriptionResponse sends subscribed topic messages back to Rcc server
	SendSubscriptionResponse(ctx context.Context, in *SendSubscriptionResponseRequest, opts ...grpc.CallOption) (*SendSubscriptionResponseResponse, error)
	// (DEPRECATED) SendCommandResponseStream sends command response from agent to
	// rcc. It is deprecated using SendCommandResponse service instead.
	SendCommandResponseStream(ctx context.Context, opts ...grpc.CallOption) (Agent_SendCommandResponseStreamClient, error)
	// PostData service sends single datapoints from agent to rcc. These
	// points include ros messages or system stats which is then sent to influx at
	// the server.
	PostData(ctx context.Context, in *model.Datapoint, opts ...grpc.CallOption) (*PostDataResponse, error)
	// PostMultiData service sends multiple datapoints from agent to rcc. These
	// points include ros messages or system stats which is then sent to influx at
	// the server.
	PostMultiData(ctx context.Context, in *PostMultiDataRequest, opts ...grpc.CallOption) (*PostMultiDataResponse, error)
	// StreamData sends telemetry datapoints from rosnode to agent. This is used
	// for receiving telop data such as pose, realsense camera feed etc. from ros
	// node to agent and then agent sends the data to the client using webrtc
	// datachannel
	StreamData(ctx context.Context, opts ...grpc.CallOption) (Agent_StreamDataClient, error)
	// ReceiveRosMessages service is established between rosnode and agent. It
	// gets ros messages which needs to be published such as command velocity
	// messages, camera control messages from agent and then pubslishes to ros.
	ReceiveRosMessages(ctx context.Context, opts ...grpc.CallOption) (Agent_ReceiveRosMessagesClient, error)
	// Get rtc signal from Rcc server to agent
	GetRtcSignal(ctx context.Context, opts ...grpc.CallOption) (Agent_GetRtcSignalClient, error)
	// To transfer the signaling data between RCC server and agent
	SendSignalData(ctx context.Context, in *SignalDataRequest, opts ...grpc.CallOption) (*SignalDataResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) AuthenticateAgent(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, Agent_AuthenticateAgent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetDeviceSpecifications(ctx context.Context, in *SpecsRequest, opts ...grpc.CallOption) (*SpecsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SpecsResponse)
	err := c.cc.Invoke(ctx, Agent_GetDeviceSpecifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetAgentConfiguration(ctx context.Context, in *GetAgentConfigurationRequest, opts ...grpc.CallOption) (*GetAgentConfigurationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAgentConfigurationResponse)
	err := c.cc.Invoke(ctx, Agent_GetAgentConfiguration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Health(ctx context.Context, opts ...grpc.CallOption) (Agent_HealthClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[0], Agent_Health_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentHealthClient{ClientStream: stream}
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
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[1], Agent_GetCommandRequestStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentGetCommandRequestStreamClient{ClientStream: stream}
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

func (c *agentClient) SendCommandResponse(ctx context.Context, in *SendCommandResponseRequest, opts ...grpc.CallOption) (*SendCommandResponseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendCommandResponseResponse)
	err := c.cc.Invoke(ctx, Agent_SendCommandResponse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetSubscriptionRequestStream(ctx context.Context, in *GetSubscriptionRequestStreamRequest, opts ...grpc.CallOption) (Agent_GetSubscriptionRequestStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[2], Agent_GetSubscriptionRequestStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentGetSubscriptionRequestStreamClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Agent_GetSubscriptionRequestStreamClient interface {
	Recv() (*GetSubscriptionRequestStreamResponse, error)
	grpc.ClientStream
}

type agentGetSubscriptionRequestStreamClient struct {
	grpc.ClientStream
}

func (x *agentGetSubscriptionRequestStreamClient) Recv() (*GetSubscriptionRequestStreamResponse, error) {
	m := new(GetSubscriptionRequestStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) SendSubscriptionResponse(ctx context.Context, in *SendSubscriptionResponseRequest, opts ...grpc.CallOption) (*SendSubscriptionResponseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendSubscriptionResponseResponse)
	err := c.cc.Invoke(ctx, Agent_SendSubscriptionResponse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) SendCommandResponseStream(ctx context.Context, opts ...grpc.CallOption) (Agent_SendCommandResponseStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[3], Agent_SendCommandResponseStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentSendCommandResponseStreamClient{ClientStream: stream}
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
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostDataResponse)
	err := c.cc.Invoke(ctx, Agent_PostData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) PostMultiData(ctx context.Context, in *PostMultiDataRequest, opts ...grpc.CallOption) (*PostMultiDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostMultiDataResponse)
	err := c.cc.Invoke(ctx, Agent_PostMultiData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StreamData(ctx context.Context, opts ...grpc.CallOption) (Agent_StreamDataClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[4], Agent_StreamData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentStreamDataClient{ClientStream: stream}
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
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[5], Agent_ReceiveRosMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentReceiveRosMessagesClient{ClientStream: stream}
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
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[6], Agent_GetRtcSignal_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentGetRtcSignalClient{ClientStream: stream}
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

func (c *agentClient) SendSignalData(ctx context.Context, in *SignalDataRequest, opts ...grpc.CallOption) (*SignalDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignalDataResponse)
	err := c.cc.Invoke(ctx, Agent_SendSignalData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	// AuthenticateAgent service authenticates the agent and return a boolean
	// status which determines the authentication status of the agent
	AuthenticateAgent(context.Context, *AuthRequest) (*AuthResponse, error)
	// GetDeviceSpecifications returns the device specification details.
	GetDeviceSpecifications(context.Context, *SpecsRequest) (*SpecsResponse, error)
	// GetAgentConfiguration returns the Agent configuration.
	GetAgentConfiguration(context.Context, *GetAgentConfigurationRequest) (*GetAgentConfigurationResponse, error)
	// Heath service is established between agent and server.When the connection
	// establises the status for the deivce changes
	// to ONLINE. This constantly checks for connection. If the connection is
	// disconnected the device status changes to OFFLINE.
	Health(Agent_HealthServer) error
	// GetCommandRequestStream service receives command from the server both
	// "SERVICE" and "ACTION" commands. The Command should be configured in the
	// agent's config first. Unconfigured command is not accepted
	GetCommandRequestStream(*GetCommandRequestStreamRequest, Agent_GetCommandRequestStreamServer) error
	// SendCommandResponse service sends command response (Just response if it is
	// service call, if Action call it sends both result and feedback ) for a
	// particular command to the server.
	SendCommandResponse(context.Context, *SendCommandResponseRequest) (*SendCommandResponseResponse, error)
	// GetSubscriptionRequestStream service recieves subscription related commands
	// such as "SUBSCRIBE", "UNSUBSCRIBE", "GET_TOPICS", "GET_SUBSCRIBED_TOPICS" and
	// sends back the response to the server. This is previously implemented for
	// node sdk with grpc, and not work with current implementation.
	GetSubscriptionRequestStream(*GetSubscriptionRequestStreamRequest, Agent_GetSubscriptionRequestStreamServer) error
	// SendSubscriptionResponse sends subscribed topic messages back to Rcc server
	SendSubscriptionResponse(context.Context, *SendSubscriptionResponseRequest) (*SendSubscriptionResponseResponse, error)
	// (DEPRECATED) SendCommandResponseStream sends command response from agent to
	// rcc. It is deprecated using SendCommandResponse service instead.
	SendCommandResponseStream(Agent_SendCommandResponseStreamServer) error
	// PostData service sends single datapoints from agent to rcc. These
	// points include ros messages or system stats which is then sent to influx at
	// the server.
	PostData(context.Context, *model.Datapoint) (*PostDataResponse, error)
	// PostMultiData service sends multiple datapoints from agent to rcc. These
	// points include ros messages or system stats which is then sent to influx at
	// the server.
	PostMultiData(context.Context, *PostMultiDataRequest) (*PostMultiDataResponse, error)
	// StreamData sends telemetry datapoints from rosnode to agent. This is used
	// for receiving telop data such as pose, realsense camera feed etc. from ros
	// node to agent and then agent sends the data to the client using webrtc
	// datachannel
	StreamData(Agent_StreamDataServer) error
	// ReceiveRosMessages service is established between rosnode and agent. It
	// gets ros messages which needs to be published such as command velocity
	// messages, camera control messages from agent and then pubslishes to ros.
	ReceiveRosMessages(Agent_ReceiveRosMessagesServer) error
	// Get rtc signal from Rcc server to agent
	GetRtcSignal(Agent_GetRtcSignalServer) error
	// To transfer the signaling data between RCC server and agent
	SendSignalData(context.Context, *SignalDataRequest) (*SignalDataResponse, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) AuthenticateAgent(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateAgent not implemented")
}
func (UnimplementedAgentServer) GetDeviceSpecifications(context.Context, *SpecsRequest) (*SpecsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceSpecifications not implemented")
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
func (UnimplementedAgentServer) SendCommandResponse(context.Context, *SendCommandResponseRequest) (*SendCommandResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommandResponse not implemented")
}
func (UnimplementedAgentServer) GetSubscriptionRequestStream(*GetSubscriptionRequestStreamRequest, Agent_GetSubscriptionRequestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSubscriptionRequestStream not implemented")
}
func (UnimplementedAgentServer) SendSubscriptionResponse(context.Context, *SendSubscriptionResponseRequest) (*SendSubscriptionResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSubscriptionResponse not implemented")
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
func (UnimplementedAgentServer) SendSignalData(context.Context, *SignalDataRequest) (*SignalDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSignalData not implemented")
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
		FullMethod: Agent_AuthenticateAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).AuthenticateAgent(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetDeviceSpecifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpecsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetDeviceSpecifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetDeviceSpecifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetDeviceSpecifications(ctx, req.(*SpecsRequest))
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
		FullMethod: Agent_GetAgentConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetAgentConfiguration(ctx, req.(*GetAgentConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Health_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).Health(&agentHealthServer{ServerStream: stream})
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
	return srv.(AgentServer).GetCommandRequestStream(m, &agentGetCommandRequestStreamServer{ServerStream: stream})
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
		FullMethod: Agent_SendCommandResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SendCommandResponse(ctx, req.(*SendCommandResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetSubscriptionRequestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetSubscriptionRequestStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServer).GetSubscriptionRequestStream(m, &agentGetSubscriptionRequestStreamServer{ServerStream: stream})
}

type Agent_GetSubscriptionRequestStreamServer interface {
	Send(*GetSubscriptionRequestStreamResponse) error
	grpc.ServerStream
}

type agentGetSubscriptionRequestStreamServer struct {
	grpc.ServerStream
}

func (x *agentGetSubscriptionRequestStreamServer) Send(m *GetSubscriptionRequestStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Agent_SendSubscriptionResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSubscriptionResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SendSubscriptionResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_SendSubscriptionResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SendSubscriptionResponse(ctx, req.(*SendSubscriptionResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_SendCommandResponseStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).SendCommandResponseStream(&agentSendCommandResponseStreamServer{ServerStream: stream})
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
		FullMethod: Agent_PostData_FullMethodName,
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
		FullMethod: Agent_PostMultiData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).PostMultiData(ctx, req.(*PostMultiDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StreamData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).StreamData(&agentStreamDataServer{ServerStream: stream})
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
	return srv.(AgentServer).ReceiveRosMessages(&agentReceiveRosMessagesServer{ServerStream: stream})
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
	return srv.(AgentServer).GetRtcSignal(&agentGetRtcSignalServer{ServerStream: stream})
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

func _Agent_SendSignalData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SendSignalData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_SendSignalData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SendSignalData(ctx, req.(*SignalDataRequest))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "GetDeviceSpecifications",
			Handler:    _Agent_GetDeviceSpecifications_Handler,
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
			MethodName: "SendSubscriptionResponse",
			Handler:    _Agent_SendSubscriptionResponse_Handler,
		},
		{
			MethodName: "PostData",
			Handler:    _Agent_PostData_Handler,
		},
		{
			MethodName: "PostMultiData",
			Handler:    _Agent_PostMultiData_Handler,
		},
		{
			MethodName: "SendSignalData",
			Handler:    _Agent_SendSignalData_Handler,
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
