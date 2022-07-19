// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: protos/model/v1/math.proto

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

type Numeric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Numeric) Reset() {
	*x = Numeric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_math_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Numeric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Numeric) ProtoMessage() {}

func (x *Numeric) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_math_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Numeric.ProtoReflect.Descriptor instead.
func (*Numeric) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_math_proto_rawDescGZIP(), []int{0}
}

func (x *Numeric) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type NumericSetEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	Label string  `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Unit  string  `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
}

func (x *NumericSetEntry) Reset() {
	*x = NumericSetEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_math_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumericSetEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumericSetEntry) ProtoMessage() {}

func (x *NumericSetEntry) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_math_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumericSetEntry.ProtoReflect.Descriptor instead.
func (*NumericSetEntry) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_math_proto_rawDescGZIP(), []int{1}
}

func (x *NumericSetEntry) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *NumericSetEntry) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *NumericSetEntry) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

type NumericSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numerics []*NumericSetEntry `protobuf:"bytes,1,rep,name=numerics,proto3" json:"numerics,omitempty"`
}

func (x *NumericSet) Reset() {
	*x = NumericSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_math_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumericSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumericSet) ProtoMessage() {}

func (x *NumericSet) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_math_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumericSet.ProtoReflect.Descriptor instead.
func (*NumericSet) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_math_proto_rawDescGZIP(), []int{2}
}

func (x *NumericSet) GetNumerics() []*NumericSetEntry {
	if x != nil {
		return x.Numerics
	}
	return nil
}

type Bit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value bool   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Bit) Reset() {
	*x = Bit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_math_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bit) ProtoMessage() {}

func (x *Bit) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_math_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bit.ProtoReflect.Descriptor instead.
func (*Bit) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_math_proto_rawDescGZIP(), []int{3}
}

func (x *Bit) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Bit) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type Bitset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bits []*Bit `protobuf:"bytes,1,rep,name=bits,proto3" json:"bits,omitempty"`
}

func (x *Bitset) Reset() {
	*x = Bitset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_math_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bitset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bitset) ProtoMessage() {}

func (x *Bitset) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_math_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bitset.ProtoReflect.Descriptor instead.
func (*Bitset) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_math_proto_rawDescGZIP(), []int{4}
}

func (x *Bitset) GetBits() []*Bit {
	if x != nil {
		return x.Bits
	}
	return nil
}

type Twist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Linear  *Vector3 `protobuf:"bytes,1,opt,name=linear,proto3" json:"linear,omitempty"`
	Angular *Vector3 `protobuf:"bytes,2,opt,name=angular,proto3" json:"angular,omitempty"`
}

func (x *Twist) Reset() {
	*x = Twist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_math_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Twist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Twist) ProtoMessage() {}

func (x *Twist) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_math_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Twist.ProtoReflect.Descriptor instead.
func (*Twist) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_math_proto_rawDescGZIP(), []int{5}
}

func (x *Twist) GetLinear() *Vector3 {
	if x != nil {
		return x.Linear
	}
	return nil
}

func (x *Twist) GetAngular() *Vector3 {
	if x != nil {
		return x.Angular
	}
	return nil
}

type Transform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Translation *Vector3    `protobuf:"bytes,1,opt,name=translation,proto3" json:"translation,omitempty"`
	Rotation    *Quaternion `protobuf:"bytes,2,opt,name=rotation,proto3" json:"rotation,omitempty"`
}

func (x *Transform) Reset() {
	*x = Transform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_math_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transform) ProtoMessage() {}

func (x *Transform) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_math_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transform.ProtoReflect.Descriptor instead.
func (*Transform) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_math_proto_rawDescGZIP(), []int{6}
}

func (x *Transform) GetTranslation() *Vector3 {
	if x != nil {
		return x.Translation
	}
	return nil
}

func (x *Transform) GetRotation() *Quaternion {
	if x != nil {
		return x.Rotation
	}
	return nil
}

type TransformFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentFrame string     `protobuf:"bytes,1,opt,name=parent_frame,json=parentFrame,proto3" json:"parent_frame,omitempty"`
	ChildFrame  string     `protobuf:"bytes,2,opt,name=child_frame,json=childFrame,proto3" json:"child_frame,omitempty"`
	Transform   *Transform `protobuf:"bytes,3,opt,name=transform,proto3" json:"transform,omitempty"`
}

func (x *TransformFrame) Reset() {
	*x = TransformFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_math_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformFrame) ProtoMessage() {}

func (x *TransformFrame) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_math_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformFrame.ProtoReflect.Descriptor instead.
func (*TransformFrame) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_math_proto_rawDescGZIP(), []int{7}
}

func (x *TransformFrame) GetParentFrame() string {
	if x != nil {
		return x.ParentFrame
	}
	return ""
}

func (x *TransformFrame) GetChildFrame() string {
	if x != nil {
		return x.ChildFrame
	}
	return ""
}

func (x *TransformFrame) GetTransform() *Transform {
	if x != nil {
		return x.Transform
	}
	return nil
}

type Vector3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vector3) Reset() {
	*x = Vector3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_math_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector3) ProtoMessage() {}

func (x *Vector3) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_math_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector3.ProtoReflect.Descriptor instead.
func (*Vector3) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_math_proto_rawDescGZIP(), []int{8}
}

func (x *Vector3) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector3) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vector3) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_math_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_math_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_math_proto_rawDescGZIP(), []int{9}
}

func (x *Point) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Point) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Quaternion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
	W float64 `protobuf:"fixed64,4,opt,name=w,proto3" json:"w,omitempty"`
}

func (x *Quaternion) Reset() {
	*x = Quaternion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_math_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quaternion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quaternion) ProtoMessage() {}

func (x *Quaternion) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_math_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quaternion.ProtoReflect.Descriptor instead.
func (*Quaternion) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_math_proto_rawDescGZIP(), []int{10}
}

func (x *Quaternion) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Quaternion) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Quaternion) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Quaternion) GetW() float64 {
	if x != nil {
		return x.W
	}
	return 0
}

var File_protos_model_v1_math_proto protoreflect.FileDescriptor

var file_protos_model_v1_math_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x1f, 0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69,
	0x63, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x51, 0x0a, 0x0f, 0x4e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x43, 0x0a, 0x0a, 0x4e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x53, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x53, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x73, 0x22,
	0x2d, 0x0a, 0x03, 0x42, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2b,
	0x0a, 0x06, 0x42, 0x69, 0x74, 0x73, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x69, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x42, 0x69, 0x74, 0x52, 0x04, 0x62, 0x69, 0x74, 0x73, 0x22, 0x5f, 0x0a, 0x05, 0x54,
	0x77, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x12,
	0x2b, 0x0a, 0x07, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x33, 0x52, 0x07, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x22, 0x72, 0x0a, 0x09,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x33, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x33, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30,
	0x0a, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x51, 0x75, 0x61, 0x74,
	0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x87, 0x01, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52,
	0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x33, 0x0a, 0x07, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x33, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x7a, 0x22,
	0x31, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x01, 0x7a, 0x22, 0x44, 0x0a, 0x0a, 0x51, 0x75, 0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01,
	0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x7a, 0x12, 0x0c, 0x0a, 0x01, 0x77, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x77, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x6f, 0x74, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x78, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_v1_math_proto_rawDescOnce sync.Once
	file_protos_model_v1_math_proto_rawDescData = file_protos_model_v1_math_proto_rawDesc
)

func file_protos_model_v1_math_proto_rawDescGZIP() []byte {
	file_protos_model_v1_math_proto_rawDescOnce.Do(func() {
		file_protos_model_v1_math_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_v1_math_proto_rawDescData)
	})
	return file_protos_model_v1_math_proto_rawDescData
}

var file_protos_model_v1_math_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protos_model_v1_math_proto_goTypes = []interface{}{
	(*Numeric)(nil),         // 0: v1.model.Numeric
	(*NumericSetEntry)(nil), // 1: v1.model.NumericSetEntry
	(*NumericSet)(nil),      // 2: v1.model.NumericSet
	(*Bit)(nil),             // 3: v1.model.Bit
	(*Bitset)(nil),          // 4: v1.model.Bitset
	(*Twist)(nil),           // 5: v1.model.Twist
	(*Transform)(nil),       // 6: v1.model.Transform
	(*TransformFrame)(nil),  // 7: v1.model.TransformFrame
	(*Vector3)(nil),         // 8: v1.model.Vector3
	(*Point)(nil),           // 9: v1.model.Point
	(*Quaternion)(nil),      // 10: v1.model.Quaternion
}
var file_protos_model_v1_math_proto_depIdxs = []int32{
	1,  // 0: v1.model.NumericSet.numerics:type_name -> v1.model.NumericSetEntry
	3,  // 1: v1.model.Bitset.bits:type_name -> v1.model.Bit
	8,  // 2: v1.model.Twist.linear:type_name -> v1.model.Vector3
	8,  // 3: v1.model.Twist.angular:type_name -> v1.model.Vector3
	8,  // 4: v1.model.Transform.translation:type_name -> v1.model.Vector3
	10, // 5: v1.model.Transform.rotation:type_name -> v1.model.Quaternion
	6,  // 6: v1.model.TransformFrame.transform:type_name -> v1.model.Transform
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_protos_model_v1_math_proto_init() }
func file_protos_model_v1_math_proto_init() {
	if File_protos_model_v1_math_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_model_v1_math_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Numeric); i {
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
		file_protos_model_v1_math_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumericSetEntry); i {
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
		file_protos_model_v1_math_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumericSet); i {
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
		file_protos_model_v1_math_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bit); i {
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
		file_protos_model_v1_math_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bitset); i {
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
		file_protos_model_v1_math_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Twist); i {
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
		file_protos_model_v1_math_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transform); i {
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
		file_protos_model_v1_math_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformFrame); i {
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
		file_protos_model_v1_math_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector3); i {
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
		file_protos_model_v1_math_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_protos_model_v1_math_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quaternion); i {
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
			RawDescriptor: file_protos_model_v1_math_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_v1_math_proto_goTypes,
		DependencyIndexes: file_protos_model_v1_math_proto_depIdxs,
		MessageInfos:      file_protos_model_v1_math_proto_msgTypes,
	}.Build()
	File_protos_model_v1_math_proto = out.File
	file_protos_model_v1_math_proto_rawDesc = nil
	file_protos_model_v1_math_proto_goTypes = nil
	file_protos_model_v1_math_proto_depIdxs = nil
}
