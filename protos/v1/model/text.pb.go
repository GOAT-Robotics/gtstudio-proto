// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: protos/model/v1/text.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Text struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Text) Reset() {
	*x = Text{}
	mi := &file_protos_model_v1_text_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Text) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text) ProtoMessage() {}

func (x *Text) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_text_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text.ProtoReflect.Descriptor instead.
func (*Text) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_text_proto_rawDescGZIP(), []int{0}
}

func (x *Text) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Json struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Json) Reset() {
	*x = Json{}
	mi := &file_protos_model_v1_text_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Json) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Json) ProtoMessage() {}

func (x *Json) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_text_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Json.ProtoReflect.Descriptor instead.
func (*Json) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_text_proto_rawDescGZIP(), []int{1}
}

func (x *Json) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_protos_model_v1_text_proto protoreflect.FileDescriptor

const file_protos_model_v1_text_proto_rawDesc = "" +
	"\n" +
	"\x1aprotos/model/v1/text.proto\x12\bv1.model\"\x1c\n" +
	"\x04Text\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05value\"\x1c\n" +
	"\x04Json\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05valueB9Z7github.com/GOAT-Robotics/gtstudio-proto/protos/v1/modelb\x06proto3"

var (
	file_protos_model_v1_text_proto_rawDescOnce sync.Once
	file_protos_model_v1_text_proto_rawDescData []byte
)

func file_protos_model_v1_text_proto_rawDescGZIP() []byte {
	file_protos_model_v1_text_proto_rawDescOnce.Do(func() {
		file_protos_model_v1_text_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_model_v1_text_proto_rawDesc), len(file_protos_model_v1_text_proto_rawDesc)))
	})
	return file_protos_model_v1_text_proto_rawDescData
}

var file_protos_model_v1_text_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_model_v1_text_proto_goTypes = []any{
	(*Text)(nil), // 0: v1.model.Text
	(*Json)(nil), // 1: v1.model.Json
}
var file_protos_model_v1_text_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_model_v1_text_proto_init() }
func file_protos_model_v1_text_proto_init() {
	if File_protos_model_v1_text_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_model_v1_text_proto_rawDesc), len(file_protos_model_v1_text_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_v1_text_proto_goTypes,
		DependencyIndexes: file_protos_model_v1_text_proto_depIdxs,
		MessageInfos:      file_protos_model_v1_text_proto_msgTypes,
	}.Build()
	File_protos_model_v1_text_proto = out.File
	file_protos_model_v1_text_proto_goTypes = nil
	file_protos_model_v1_text_proto_depIdxs = nil
}
