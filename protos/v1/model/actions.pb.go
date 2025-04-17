// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protos/model/v1/actions.proto

package model

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Command      string               `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Compressed   bool                 `protobuf:"varint,3,opt,name=compressed,proto3" json:"compressed,omitempty"`
	Data         string               `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ScrubberTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=scrubber_time,json=scrubberTime,proto3" json:"scrubber_time,omitempty"`
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_actions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_actions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRequest.ProtoReflect.Descriptor instead.
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_actions_proto_rawDescGZIP(), []int{0}
}

func (x *ActionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ActionRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ActionRequest) GetCompressed() bool {
	if x != nil {
		return x.Compressed
	}
	return false
}

func (x *ActionRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *ActionRequest) GetScrubberTime() *timestamp.Timestamp {
	if x != nil {
		return x.ScrubberTime
	}
	return nil
}

type ActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Status    bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	IsResult  bool   `protobuf:"varint,3,opt,name=isResult,proto3" json:"isResult,omitempty"`
	// Types that are assignable to Result:
	//
	//	*ActionResponse_Data
	//	*ActionResponse_Raw
	//	*ActionResponse_Error
	Result isActionResponse_Result `protobuf_oneof:"result"`
}

func (x *ActionResponse) Reset() {
	*x = ActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_actions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResponse) ProtoMessage() {}

func (x *ActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_actions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_actions_proto_rawDescGZIP(), []int{1}
}

func (x *ActionResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ActionResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ActionResponse) GetIsResult() bool {
	if x != nil {
		return x.IsResult
	}
	return false
}

func (m *ActionResponse) GetResult() isActionResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ActionResponse) GetData() string {
	if x, ok := x.GetResult().(*ActionResponse_Data); ok {
		return x.Data
	}
	return ""
}

func (x *ActionResponse) GetRaw() []byte {
	if x, ok := x.GetResult().(*ActionResponse_Raw); ok {
		return x.Raw
	}
	return nil
}

func (x *ActionResponse) GetError() string {
	if x, ok := x.GetResult().(*ActionResponse_Error); ok {
		return x.Error
	}
	return ""
}

type isActionResponse_Result interface {
	isActionResponse_Result()
}

type ActionResponse_Data struct {
	Data string `protobuf:"bytes,4,opt,name=data,proto3,oneof"`
}

type ActionResponse_Raw struct {
	Raw []byte `protobuf:"bytes,5,opt,name=raw,proto3,oneof"`
}

type ActionResponse_Error struct {
	Error string `protobuf:"bytes,6,opt,name=error,proto3,oneof"`
}

func (*ActionResponse_Data) isActionResponse_Result() {}

func (*ActionResponse_Raw) isActionResponse_Result() {}

func (*ActionResponse_Error) isActionResponse_Result() {}

var File_protos_model_v1_actions_proto protoreflect.FileDescriptor

var file_protos_model_v1_actions_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0d, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x63,
	0x72, 0x75, 0x62, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x73,
	0x63, 0x72, 0x75, 0x62, 0x62, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x0e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x16, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x39, 0x5a,
	0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x4f, 0x41, 0x54,
	0x2d, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x67, 0x74, 0x73, 0x74, 0x75, 0x64,
	0x69, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_v1_actions_proto_rawDescOnce sync.Once
	file_protos_model_v1_actions_proto_rawDescData = file_protos_model_v1_actions_proto_rawDesc
)

func file_protos_model_v1_actions_proto_rawDescGZIP() []byte {
	file_protos_model_v1_actions_proto_rawDescOnce.Do(func() {
		file_protos_model_v1_actions_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_v1_actions_proto_rawDescData)
	})
	return file_protos_model_v1_actions_proto_rawDescData
}

var file_protos_model_v1_actions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_model_v1_actions_proto_goTypes = []any{
	(*ActionRequest)(nil),       // 0: v1.model.ActionRequest
	(*ActionResponse)(nil),      // 1: v1.model.ActionResponse
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_protos_model_v1_actions_proto_depIdxs = []int32{
	2, // 0: v1.model.ActionRequest.scrubber_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_model_v1_actions_proto_init() }
func file_protos_model_v1_actions_proto_init() {
	if File_protos_model_v1_actions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_model_v1_actions_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_model_v1_actions_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_protos_model_v1_actions_proto_msgTypes[1].OneofWrappers = []any{
		(*ActionResponse_Data)(nil),
		(*ActionResponse_Raw)(nil),
		(*ActionResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_model_v1_actions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_v1_actions_proto_goTypes,
		DependencyIndexes: file_protos_model_v1_actions_proto_depIdxs,
		MessageInfos:      file_protos_model_v1_actions_proto_msgTypes,
	}.Build()
	File_protos_model_v1_actions_proto = out.File
	file_protos_model_v1_actions_proto_rawDesc = nil
	file_protos_model_v1_actions_proto_goTypes = nil
	file_protos_model_v1_actions_proto_depIdxs = nil
}
