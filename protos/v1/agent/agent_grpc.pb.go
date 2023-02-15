// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protos/agent/v1/agent.proto

package agent

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

// AgentRosBridgeClient is the client API for AgentRosBridge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentRosBridgeClient interface {
	// GetAgentConfiguration returns the Agent configuration.
	GetAgentConfiguration(ctx context.Context, in *GetAgentConfigurationRequest, opts ...grpc.CallOption) (*GetAgentConfigurationResponse, error)
	GetRosSubscriptionConfig(ctx context.Context, in *GetRosSubscriptionConfigRequest, opts ...grpc.CallOption) (*GetRosSubscriptionConfigResponse, error)
	GetRosCommandConfig(ctx context.Context, in *GetRosCommandConfigRequest, opts ...grpc.CallOption) (*GetRosCommandConfigResponse, error)
	// Health can be used to check if the Agent is running and its gRPC API is
	// available.
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	GetCommandRequestStream(ctx context.Context, in *GetCommandRequestStreamRequest, opts ...grpc.CallOption) (AgentRosBridge_GetCommandRequestStreamClient, error)
	GetSubscriptionRequestStream(ctx context.Context, opts ...grpc.CallOption) (AgentRosBridge_GetSubscriptionRequestStreamClient, error)
	// SendCommandResponse sends a response to a command request.
	SendCommandResponse(ctx context.Context, in *SendCommandResponseRequest, opts ...grpc.CallOption) (*SendCommandResponseResponse, error)
	SendCommandResponseStream(ctx context.Context, opts ...grpc.CallOption) (AgentRosBridge_SendCommandResponseStreamClient, error)
	PostData(ctx context.Context, in *model.Datapoint, opts ...grpc.CallOption) (*PostDataResponse, error)
	PostMultiData(ctx context.Context, in *PostMultiDataRequest, opts ...grpc.CallOption) (*PostMultiDataResponse, error)
	// StreamData accepts a stream of data points. See PostData for information on
	// expected error conditions and codes.
	StreamData(ctx context.Context, opts ...grpc.CallOption) (AgentRosBridge_StreamDataClient, error)
	ReceiveRosMessages(ctx context.Context, opts ...grpc.CallOption) (AgentRosBridge_ReceiveRosMessagesClient, error)
}

type agentRosBridgeClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentRosBridgeClient(cc grpc.ClientConnInterface) AgentRosBridgeClient {
	return &agentRosBridgeClient{cc}
}

func (c *agentRosBridgeClient) GetAgentConfiguration(ctx context.Context, in *GetAgentConfigurationRequest, opts ...grpc.CallOption) (*GetAgentConfigurationResponse, error) {
	out := new(GetAgentConfigurationResponse)
	err := c.cc.Invoke(ctx, "/v1.agent.AgentRosBridge/GetAgentConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) GetRosSubscriptionConfig(ctx context.Context, in *GetRosSubscriptionConfigRequest, opts ...grpc.CallOption) (*GetRosSubscriptionConfigResponse, error) {
	out := new(GetRosSubscriptionConfigResponse)
	err := c.cc.Invoke(ctx, "/v1.agent.AgentRosBridge/GetRosSubscriptionConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) GetRosCommandConfig(ctx context.Context, in *GetRosCommandConfigRequest, opts ...grpc.CallOption) (*GetRosCommandConfigResponse, error) {
	out := new(GetRosCommandConfigResponse)
	err := c.cc.Invoke(ctx, "/v1.agent.AgentRosBridge/GetRosCommandConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/v1.agent.AgentRosBridge/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) GetCommandRequestStream(ctx context.Context, in *GetCommandRequestStreamRequest, opts ...grpc.CallOption) (AgentRosBridge_GetCommandRequestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentRosBridge_ServiceDesc.Streams[0], "/v1.agent.AgentRosBridge/GetCommandRequestStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentRosBridgeGetCommandRequestStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentRosBridge_GetCommandRequestStreamClient interface {
	Recv() (*GetCommandRequestStreamResponse, error)
	grpc.ClientStream
}

type agentRosBridgeGetCommandRequestStreamClient struct {
	grpc.ClientStream
}

func (x *agentRosBridgeGetCommandRequestStreamClient) Recv() (*GetCommandRequestStreamResponse, error) {
	m := new(GetCommandRequestStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentRosBridgeClient) GetSubscriptionRequestStream(ctx context.Context, opts ...grpc.CallOption) (AgentRosBridge_GetSubscriptionRequestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentRosBridge_ServiceDesc.Streams[1], "/v1.agent.AgentRosBridge/GetSubscriptionRequestStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentRosBridgeGetSubscriptionRequestStreamClient{stream}
	return x, nil
}

type AgentRosBridge_GetSubscriptionRequestStreamClient interface {
	Send(*GetSubscriptionRequestStreamResponse) error
	Recv() (*GetSubscriptionRequestStreamRequest, error)
	grpc.ClientStream
}

type agentRosBridgeGetSubscriptionRequestStreamClient struct {
	grpc.ClientStream
}

func (x *agentRosBridgeGetSubscriptionRequestStreamClient) Send(m *GetSubscriptionRequestStreamResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentRosBridgeGetSubscriptionRequestStreamClient) Recv() (*GetSubscriptionRequestStreamRequest, error) {
	m := new(GetSubscriptionRequestStreamRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentRosBridgeClient) SendCommandResponse(ctx context.Context, in *SendCommandResponseRequest, opts ...grpc.CallOption) (*SendCommandResponseResponse, error) {
	out := new(SendCommandResponseResponse)
	err := c.cc.Invoke(ctx, "/v1.agent.AgentRosBridge/SendCommandResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) SendCommandResponseStream(ctx context.Context, opts ...grpc.CallOption) (AgentRosBridge_SendCommandResponseStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentRosBridge_ServiceDesc.Streams[2], "/v1.agent.AgentRosBridge/SendCommandResponseStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentRosBridgeSendCommandResponseStreamClient{stream}
	return x, nil
}

type AgentRosBridge_SendCommandResponseStreamClient interface {
	Send(*SendCommandResponseRequest) error
	Recv() (*SendCommandResponseResponse, error)
	grpc.ClientStream
}

type agentRosBridgeSendCommandResponseStreamClient struct {
	grpc.ClientStream
}

func (x *agentRosBridgeSendCommandResponseStreamClient) Send(m *SendCommandResponseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentRosBridgeSendCommandResponseStreamClient) Recv() (*SendCommandResponseResponse, error) {
	m := new(SendCommandResponseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentRosBridgeClient) PostData(ctx context.Context, in *model.Datapoint, opts ...grpc.CallOption) (*PostDataResponse, error) {
	out := new(PostDataResponse)
	err := c.cc.Invoke(ctx, "/v1.agent.AgentRosBridge/PostData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) PostMultiData(ctx context.Context, in *PostMultiDataRequest, opts ...grpc.CallOption) (*PostMultiDataResponse, error) {
	out := new(PostMultiDataResponse)
	err := c.cc.Invoke(ctx, "/v1.agent.AgentRosBridge/PostMultiData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) StreamData(ctx context.Context, opts ...grpc.CallOption) (AgentRosBridge_StreamDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentRosBridge_ServiceDesc.Streams[3], "/v1.agent.AgentRosBridge/StreamData", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentRosBridgeStreamDataClient{stream}
	return x, nil
}

type AgentRosBridge_StreamDataClient interface {
	Send(*StreamDataResponse) error
	Recv() (*StreamDataRequest, error)
	grpc.ClientStream
}

type agentRosBridgeStreamDataClient struct {
	grpc.ClientStream
}

func (x *agentRosBridgeStreamDataClient) Send(m *StreamDataResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentRosBridgeStreamDataClient) Recv() (*StreamDataRequest, error) {
	m := new(StreamDataRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentRosBridgeClient) ReceiveRosMessages(ctx context.Context, opts ...grpc.CallOption) (AgentRosBridge_ReceiveRosMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentRosBridge_ServiceDesc.Streams[4], "/v1.agent.AgentRosBridge/ReceiveRosMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentRosBridgeReceiveRosMessagesClient{stream}
	return x, nil
}

type AgentRosBridge_ReceiveRosMessagesClient interface {
	Send(*RecieveRosMessagesResponse) error
	Recv() (*RecieveRosMessagesRequest, error)
	grpc.ClientStream
}

type agentRosBridgeReceiveRosMessagesClient struct {
	grpc.ClientStream
}

func (x *agentRosBridgeReceiveRosMessagesClient) Send(m *RecieveRosMessagesResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentRosBridgeReceiveRosMessagesClient) Recv() (*RecieveRosMessagesRequest, error) {
	m := new(RecieveRosMessagesRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentRosBridgeServer is the server API for AgentRosBridge service.
// All implementations must embed UnimplementedAgentRosBridgeServer
// for forward compatibility
type AgentRosBridgeServer interface {
	// GetAgentConfiguration returns the Agent configuration.
	GetAgentConfiguration(context.Context, *GetAgentConfigurationRequest) (*GetAgentConfigurationResponse, error)
	GetRosSubscriptionConfig(context.Context, *GetRosSubscriptionConfigRequest) (*GetRosSubscriptionConfigResponse, error)
	GetRosCommandConfig(context.Context, *GetRosCommandConfigRequest) (*GetRosCommandConfigResponse, error)
	// Health can be used to check if the Agent is running and its gRPC API is
	// available.
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	GetCommandRequestStream(*GetCommandRequestStreamRequest, AgentRosBridge_GetCommandRequestStreamServer) error
	GetSubscriptionRequestStream(AgentRosBridge_GetSubscriptionRequestStreamServer) error
	// SendCommandResponse sends a response to a command request.
	SendCommandResponse(context.Context, *SendCommandResponseRequest) (*SendCommandResponseResponse, error)
	SendCommandResponseStream(AgentRosBridge_SendCommandResponseStreamServer) error
	PostData(context.Context, *model.Datapoint) (*PostDataResponse, error)
	PostMultiData(context.Context, *PostMultiDataRequest) (*PostMultiDataResponse, error)
	// StreamData accepts a stream of data points. See PostData for information on
	// expected error conditions and codes.
	StreamData(AgentRosBridge_StreamDataServer) error
	ReceiveRosMessages(AgentRosBridge_ReceiveRosMessagesServer) error
	mustEmbedUnimplementedAgentRosBridgeServer()
}

// UnimplementedAgentRosBridgeServer must be embedded to have forward compatible implementations.
type UnimplementedAgentRosBridgeServer struct {
}

func (UnimplementedAgentRosBridgeServer) GetAgentConfiguration(context.Context, *GetAgentConfigurationRequest) (*GetAgentConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgentConfiguration not implemented")
}
func (UnimplementedAgentRosBridgeServer) GetRosSubscriptionConfig(context.Context, *GetRosSubscriptionConfigRequest) (*GetRosSubscriptionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRosSubscriptionConfig not implemented")
}
func (UnimplementedAgentRosBridgeServer) GetRosCommandConfig(context.Context, *GetRosCommandConfigRequest) (*GetRosCommandConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRosCommandConfig not implemented")
}
func (UnimplementedAgentRosBridgeServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedAgentRosBridgeServer) GetCommandRequestStream(*GetCommandRequestStreamRequest, AgentRosBridge_GetCommandRequestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCommandRequestStream not implemented")
}
func (UnimplementedAgentRosBridgeServer) GetSubscriptionRequestStream(AgentRosBridge_GetSubscriptionRequestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSubscriptionRequestStream not implemented")
}
func (UnimplementedAgentRosBridgeServer) SendCommandResponse(context.Context, *SendCommandResponseRequest) (*SendCommandResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommandResponse not implemented")
}
func (UnimplementedAgentRosBridgeServer) SendCommandResponseStream(AgentRosBridge_SendCommandResponseStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SendCommandResponseStream not implemented")
}
func (UnimplementedAgentRosBridgeServer) PostData(context.Context, *model.Datapoint) (*PostDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostData not implemented")
}
func (UnimplementedAgentRosBridgeServer) PostMultiData(context.Context, *PostMultiDataRequest) (*PostMultiDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMultiData not implemented")
}
func (UnimplementedAgentRosBridgeServer) StreamData(AgentRosBridge_StreamDataServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamData not implemented")
}
func (UnimplementedAgentRosBridgeServer) ReceiveRosMessages(AgentRosBridge_ReceiveRosMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveRosMessages not implemented")
}
func (UnimplementedAgentRosBridgeServer) mustEmbedUnimplementedAgentRosBridgeServer() {}

// UnsafeAgentRosBridgeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentRosBridgeServer will
// result in compilation errors.
type UnsafeAgentRosBridgeServer interface {
	mustEmbedUnimplementedAgentRosBridgeServer()
}

func RegisterAgentRosBridgeServer(s grpc.ServiceRegistrar, srv AgentRosBridgeServer) {
	s.RegisterService(&AgentRosBridge_ServiceDesc, srv)
}

func _AgentRosBridge_GetAgentConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).GetAgentConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.agent.AgentRosBridge/GetAgentConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).GetAgentConfiguration(ctx, req.(*GetAgentConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_GetRosSubscriptionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRosSubscriptionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).GetRosSubscriptionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.agent.AgentRosBridge/GetRosSubscriptionConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).GetRosSubscriptionConfig(ctx, req.(*GetRosSubscriptionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_GetRosCommandConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRosCommandConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).GetRosCommandConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.agent.AgentRosBridge/GetRosCommandConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).GetRosCommandConfig(ctx, req.(*GetRosCommandConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.agent.AgentRosBridge/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_GetCommandRequestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCommandRequestStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentRosBridgeServer).GetCommandRequestStream(m, &agentRosBridgeGetCommandRequestStreamServer{stream})
}

type AgentRosBridge_GetCommandRequestStreamServer interface {
	Send(*GetCommandRequestStreamResponse) error
	grpc.ServerStream
}

type agentRosBridgeGetCommandRequestStreamServer struct {
	grpc.ServerStream
}

func (x *agentRosBridgeGetCommandRequestStreamServer) Send(m *GetCommandRequestStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentRosBridge_GetSubscriptionRequestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentRosBridgeServer).GetSubscriptionRequestStream(&agentRosBridgeGetSubscriptionRequestStreamServer{stream})
}

type AgentRosBridge_GetSubscriptionRequestStreamServer interface {
	Send(*GetSubscriptionRequestStreamRequest) error
	Recv() (*GetSubscriptionRequestStreamResponse, error)
	grpc.ServerStream
}

type agentRosBridgeGetSubscriptionRequestStreamServer struct {
	grpc.ServerStream
}

func (x *agentRosBridgeGetSubscriptionRequestStreamServer) Send(m *GetSubscriptionRequestStreamRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentRosBridgeGetSubscriptionRequestStreamServer) Recv() (*GetSubscriptionRequestStreamResponse, error) {
	m := new(GetSubscriptionRequestStreamResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AgentRosBridge_SendCommandResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCommandResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).SendCommandResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.agent.AgentRosBridge/SendCommandResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).SendCommandResponse(ctx, req.(*SendCommandResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_SendCommandResponseStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentRosBridgeServer).SendCommandResponseStream(&agentRosBridgeSendCommandResponseStreamServer{stream})
}

type AgentRosBridge_SendCommandResponseStreamServer interface {
	Send(*SendCommandResponseResponse) error
	Recv() (*SendCommandResponseRequest, error)
	grpc.ServerStream
}

type agentRosBridgeSendCommandResponseStreamServer struct {
	grpc.ServerStream
}

func (x *agentRosBridgeSendCommandResponseStreamServer) Send(m *SendCommandResponseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentRosBridgeSendCommandResponseStreamServer) Recv() (*SendCommandResponseRequest, error) {
	m := new(SendCommandResponseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AgentRosBridge_PostData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Datapoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).PostData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.agent.AgentRosBridge/PostData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).PostData(ctx, req.(*model.Datapoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_PostMultiData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMultiDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).PostMultiData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.agent.AgentRosBridge/PostMultiData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).PostMultiData(ctx, req.(*PostMultiDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_StreamData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentRosBridgeServer).StreamData(&agentRosBridgeStreamDataServer{stream})
}

type AgentRosBridge_StreamDataServer interface {
	Send(*StreamDataRequest) error
	Recv() (*StreamDataResponse, error)
	grpc.ServerStream
}

type agentRosBridgeStreamDataServer struct {
	grpc.ServerStream
}

func (x *agentRosBridgeStreamDataServer) Send(m *StreamDataRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentRosBridgeStreamDataServer) Recv() (*StreamDataResponse, error) {
	m := new(StreamDataResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AgentRosBridge_ReceiveRosMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentRosBridgeServer).ReceiveRosMessages(&agentRosBridgeReceiveRosMessagesServer{stream})
}

type AgentRosBridge_ReceiveRosMessagesServer interface {
	Send(*RecieveRosMessagesRequest) error
	Recv() (*RecieveRosMessagesResponse, error)
	grpc.ServerStream
}

type agentRosBridgeReceiveRosMessagesServer struct {
	grpc.ServerStream
}

func (x *agentRosBridgeReceiveRosMessagesServer) Send(m *RecieveRosMessagesRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentRosBridgeReceiveRosMessagesServer) Recv() (*RecieveRosMessagesResponse, error) {
	m := new(RecieveRosMessagesResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentRosBridge_ServiceDesc is the grpc.ServiceDesc for AgentRosBridge service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentRosBridge_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.agent.AgentRosBridge",
	HandlerType: (*AgentRosBridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgentConfiguration",
			Handler:    _AgentRosBridge_GetAgentConfiguration_Handler,
		},
		{
			MethodName: "GetRosSubscriptionConfig",
			Handler:    _AgentRosBridge_GetRosSubscriptionConfig_Handler,
		},
		{
			MethodName: "GetRosCommandConfig",
			Handler:    _AgentRosBridge_GetRosCommandConfig_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _AgentRosBridge_Health_Handler,
		},
		{
			MethodName: "SendCommandResponse",
			Handler:    _AgentRosBridge_SendCommandResponse_Handler,
		},
		{
			MethodName: "PostData",
			Handler:    _AgentRosBridge_PostData_Handler,
		},
		{
			MethodName: "PostMultiData",
			Handler:    _AgentRosBridge_PostMultiData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCommandRequestStream",
			Handler:       _AgentRosBridge_GetCommandRequestStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSubscriptionRequestStream",
			Handler:       _AgentRosBridge_GetSubscriptionRequestStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SendCommandResponseStream",
			Handler:       _AgentRosBridge_SendCommandResponseStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamData",
			Handler:       _AgentRosBridge_StreamData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveRosMessages",
			Handler:       _AgentRosBridge_ReceiveRosMessages_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/agent/v1/agent.proto",
}
