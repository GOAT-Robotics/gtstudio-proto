// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: protos/model/v1/specifications.proto

package model

import (
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

type DeviceSpecifications struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processor          string `protobuf:"bytes,1,opt,name=processor,proto3" json:"processor,omitempty"`
	KernelArchitecture string `protobuf:"bytes,2,opt,name=kernel_architecture,json=kernelArchitecture,proto3" json:"kernel_architecture,omitempty"`
	KernelVersion      string `protobuf:"bytes,3,opt,name=kernel_version,json=kernelVersion,proto3" json:"kernel_version,omitempty"`
	CpuCores           uint32 `protobuf:"varint,4,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	StorageSize        uint32 `protobuf:"varint,5,opt,name=storage_size,json=storageSize,proto3" json:"storage_size,omitempty"`
	MemorySize         uint32 `protobuf:"varint,6,opt,name=memory_size,json=memorySize,proto3" json:"memory_size,omitempty"`
	MemoryType         string `protobuf:"bytes,7,opt,name=memory_type,json=memoryType,proto3" json:"memory_type,omitempty"`
	FirmwareVersion    string `protobuf:"bytes,8,opt,name=firmware_version,json=firmwareVersion,proto3" json:"firmware_version,omitempty"`
	IpAddress          string `protobuf:"bytes,9,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	UbuntuVersion      string `protobuf:"bytes,10,opt,name=ubuntu_version,json=ubuntuVersion,proto3" json:"ubuntu_version,omitempty"`
	OsArchitecture     string `protobuf:"bytes,11,opt,name=os_architecture,json=osArchitecture,proto3" json:"os_architecture,omitempty"`
	RosDistribution    string `protobuf:"bytes,12,opt,name=ros_distribution,json=rosDistribution,proto3" json:"ros_distribution,omitempty"`
	RosVersion         string `protobuf:"bytes,13,opt,name=ros_version,json=rosVersion,proto3" json:"ros_version,omitempty"`
	Firewall           string `protobuf:"bytes,14,opt,name=firewall,proto3" json:"firewall,omitempty"`
	AgentVersion       string `protobuf:"bytes,15,opt,name=agent_version,json=agentVersion,proto3" json:"agent_version,omitempty"`
	DeviceModel        string `protobuf:"bytes,16,opt,name=device_model,json=deviceModel,proto3" json:"device_model,omitempty"`
}

func (x *DeviceSpecifications) Reset() {
	*x = DeviceSpecifications{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_specifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceSpecifications) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceSpecifications) ProtoMessage() {}

func (x *DeviceSpecifications) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_specifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceSpecifications.ProtoReflect.Descriptor instead.
func (*DeviceSpecifications) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_specifications_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceSpecifications) GetProcessor() string {
	if x != nil {
		return x.Processor
	}
	return ""
}

func (x *DeviceSpecifications) GetKernelArchitecture() string {
	if x != nil {
		return x.KernelArchitecture
	}
	return ""
}

func (x *DeviceSpecifications) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

func (x *DeviceSpecifications) GetCpuCores() uint32 {
	if x != nil {
		return x.CpuCores
	}
	return 0
}

func (x *DeviceSpecifications) GetStorageSize() uint32 {
	if x != nil {
		return x.StorageSize
	}
	return 0
}

func (x *DeviceSpecifications) GetMemorySize() uint32 {
	if x != nil {
		return x.MemorySize
	}
	return 0
}

func (x *DeviceSpecifications) GetMemoryType() string {
	if x != nil {
		return x.MemoryType
	}
	return ""
}

func (x *DeviceSpecifications) GetFirmwareVersion() string {
	if x != nil {
		return x.FirmwareVersion
	}
	return ""
}

func (x *DeviceSpecifications) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *DeviceSpecifications) GetUbuntuVersion() string {
	if x != nil {
		return x.UbuntuVersion
	}
	return ""
}

func (x *DeviceSpecifications) GetOsArchitecture() string {
	if x != nil {
		return x.OsArchitecture
	}
	return ""
}

func (x *DeviceSpecifications) GetRosDistribution() string {
	if x != nil {
		return x.RosDistribution
	}
	return ""
}

func (x *DeviceSpecifications) GetRosVersion() string {
	if x != nil {
		return x.RosVersion
	}
	return ""
}

func (x *DeviceSpecifications) GetFirewall() string {
	if x != nil {
		return x.Firewall
	}
	return ""
}

func (x *DeviceSpecifications) GetAgentVersion() string {
	if x != nil {
		return x.AgentVersion
	}
	return ""
}

func (x *DeviceSpecifications) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

var File_protos_model_v1_specifications_proto protoreflect.FileDescriptor

var file_protos_model_v1_specifications_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x22, 0xd8, 0x04, 0x0a, 0x14, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x13, 0x6b, 0x65, 0x72, 0x6e, 0x65,
	0x6c, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6b, 0x65, 0x72, 0x6e,
	0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x72,
	0x6d, 0x77, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x75,
	0x62, 0x75, 0x6e, 0x74, 0x75, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x62, 0x75, 0x6e, 0x74, 0x75, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x73, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x73, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x6f, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x6f, 0x73, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x73, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x4f, 0x41, 0x54, 0x2d, 0x52,
	0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x78, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_v1_specifications_proto_rawDescOnce sync.Once
	file_protos_model_v1_specifications_proto_rawDescData = file_protos_model_v1_specifications_proto_rawDesc
)

func file_protos_model_v1_specifications_proto_rawDescGZIP() []byte {
	file_protos_model_v1_specifications_proto_rawDescOnce.Do(func() {
		file_protos_model_v1_specifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_v1_specifications_proto_rawDescData)
	})
	return file_protos_model_v1_specifications_proto_rawDescData
}

var file_protos_model_v1_specifications_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_model_v1_specifications_proto_goTypes = []any{
	(*DeviceSpecifications)(nil), // 0: v1.model.DeviceSpecifications
}
var file_protos_model_v1_specifications_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_model_v1_specifications_proto_init() }
func file_protos_model_v1_specifications_proto_init() {
	if File_protos_model_v1_specifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_model_v1_specifications_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceSpecifications); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_model_v1_specifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_v1_specifications_proto_goTypes,
		DependencyIndexes: file_protos_model_v1_specifications_proto_depIdxs,
		MessageInfos:      file_protos_model_v1_specifications_proto_msgTypes,
	}.Build()
	File_protos_model_v1_specifications_proto = out.File
	file_protos_model_v1_specifications_proto_rawDesc = nil
	file_protos_model_v1_specifications_proto_goTypes = nil
	file_protos_model_v1_specifications_proto_depIdxs = nil
}
