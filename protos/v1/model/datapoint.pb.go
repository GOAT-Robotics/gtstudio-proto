// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protos/model/v1/datapoint.proto

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

type Datapoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream    string            `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	Timestamp int64             `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // unix epoch time in milliseconds
	Tags      map[string]string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to Data:
	//
	//	*Datapoint_Text
	//	*Datapoint_Numeric
	//	*Datapoint_NumericSet
	//	*Datapoint_Bitset
	//	*Datapoint_File
	//	*Datapoint_Image
	//	*Datapoint_PointCloud
	//	*Datapoint_Location
	//	*Datapoint_Localization
	//	*Datapoint_Health
	//	*Datapoint_Json
	//	*Datapoint_Battery
	//	*Datapoint_Video
	//	*Datapoint_TransformTree
	//	*Datapoint_Raw
	Data isDatapoint_Data `protobuf_oneof:"data"`
}

func (x *Datapoint) Reset() {
	*x = Datapoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_datapoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datapoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datapoint) ProtoMessage() {}

func (x *Datapoint) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_datapoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datapoint.ProtoReflect.Descriptor instead.
func (*Datapoint) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_datapoint_proto_rawDescGZIP(), []int{0}
}

func (x *Datapoint) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *Datapoint) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Datapoint) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (m *Datapoint) GetData() isDatapoint_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Datapoint) GetText() *Text {
	if x, ok := x.GetData().(*Datapoint_Text); ok {
		return x.Text
	}
	return nil
}

func (x *Datapoint) GetNumeric() *Numeric {
	if x, ok := x.GetData().(*Datapoint_Numeric); ok {
		return x.Numeric
	}
	return nil
}

func (x *Datapoint) GetNumericSet() *NumericSet {
	if x, ok := x.GetData().(*Datapoint_NumericSet); ok {
		return x.NumericSet
	}
	return nil
}

func (x *Datapoint) GetBitset() *Bitset {
	if x, ok := x.GetData().(*Datapoint_Bitset); ok {
		return x.Bitset
	}
	return nil
}

func (x *Datapoint) GetFile() *File {
	if x, ok := x.GetData().(*Datapoint_File); ok {
		return x.File
	}
	return nil
}

func (x *Datapoint) GetImage() *Image {
	if x, ok := x.GetData().(*Datapoint_Image); ok {
		return x.Image
	}
	return nil
}

func (x *Datapoint) GetPointCloud() *PointCloud {
	if x, ok := x.GetData().(*Datapoint_PointCloud); ok {
		return x.PointCloud
	}
	return nil
}

func (x *Datapoint) GetLocation() *Location {
	if x, ok := x.GetData().(*Datapoint_Location); ok {
		return x.Location
	}
	return nil
}

func (x *Datapoint) GetLocalization() *Localization {
	if x, ok := x.GetData().(*Datapoint_Localization); ok {
		return x.Localization
	}
	return nil
}

func (x *Datapoint) GetHealth() *Health {
	if x, ok := x.GetData().(*Datapoint_Health); ok {
		return x.Health
	}
	return nil
}

func (x *Datapoint) GetJson() *Json {
	if x, ok := x.GetData().(*Datapoint_Json); ok {
		return x.Json
	}
	return nil
}

func (x *Datapoint) GetBattery() *Battery {
	if x, ok := x.GetData().(*Datapoint_Battery); ok {
		return x.Battery
	}
	return nil
}

func (x *Datapoint) GetVideo() *Video {
	if x, ok := x.GetData().(*Datapoint_Video); ok {
		return x.Video
	}
	return nil
}

func (x *Datapoint) GetTransformTree() *TransformTree {
	if x, ok := x.GetData().(*Datapoint_TransformTree); ok {
		return x.TransformTree
	}
	return nil
}

func (x *Datapoint) GetRaw() []byte {
	if x, ok := x.GetData().(*Datapoint_Raw); ok {
		return x.Raw
	}
	return nil
}

type isDatapoint_Data interface {
	isDatapoint_Data()
}

type Datapoint_Text struct {
	Text *Text `protobuf:"bytes,4,opt,name=text,proto3,oneof"`
}

type Datapoint_Numeric struct {
	Numeric *Numeric `protobuf:"bytes,5,opt,name=numeric,proto3,oneof"`
}

type Datapoint_NumericSet struct {
	NumericSet *NumericSet `protobuf:"bytes,17,opt,name=numeric_set,json=numericSet,proto3,oneof"`
}

type Datapoint_Bitset struct {
	Bitset *Bitset `protobuf:"bytes,7,opt,name=bitset,proto3,oneof"`
}

type Datapoint_File struct {
	File *File `protobuf:"bytes,8,opt,name=file,proto3,oneof"`
}

type Datapoint_Image struct {
	Image *Image `protobuf:"bytes,9,opt,name=image,proto3,oneof"`
}

type Datapoint_PointCloud struct {
	PointCloud *PointCloud `protobuf:"bytes,10,opt,name=point_cloud,json=pointCloud,proto3,oneof"`
}

type Datapoint_Location struct {
	Location *Location `protobuf:"bytes,11,opt,name=location,proto3,oneof"`
}

type Datapoint_Localization struct {
	Localization *Localization `protobuf:"bytes,12,opt,name=localization,proto3,oneof"`
}

type Datapoint_Health struct {
	Health *Health `protobuf:"bytes,13,opt,name=health,proto3,oneof"`
}

type Datapoint_Json struct {
	Json *Json `protobuf:"bytes,14,opt,name=json,proto3,oneof"`
}

type Datapoint_Battery struct {
	Battery *Battery `protobuf:"bytes,15,opt,name=battery,proto3,oneof"`
}

type Datapoint_Video struct {
	Video *Video `protobuf:"bytes,16,opt,name=video,proto3,oneof"`
}

type Datapoint_TransformTree struct {
	TransformTree *TransformTree `protobuf:"bytes,18,opt,name=transform_tree,json=transformTree,proto3,oneof"`
}

type Datapoint_Raw struct {
	Raw []byte `protobuf:"bytes,19,opt,name=raw,proto3,oneof"`
}

func (*Datapoint_Text) isDatapoint_Data() {}

func (*Datapoint_Numeric) isDatapoint_Data() {}

func (*Datapoint_NumericSet) isDatapoint_Data() {}

func (*Datapoint_Bitset) isDatapoint_Data() {}

func (*Datapoint_File) isDatapoint_Data() {}

func (*Datapoint_Image) isDatapoint_Data() {}

func (*Datapoint_PointCloud) isDatapoint_Data() {}

func (*Datapoint_Location) isDatapoint_Data() {}

func (*Datapoint_Localization) isDatapoint_Data() {}

func (*Datapoint_Health) isDatapoint_Data() {}

func (*Datapoint_Json) isDatapoint_Data() {}

func (*Datapoint_Battery) isDatapoint_Data() {}

func (*Datapoint_Video) isDatapoint_Data() {}

func (*Datapoint_TransformTree) isDatapoint_Data() {}

func (*Datapoint_Raw) isDatapoint_Data() {}

// ControlDatapoint contains data sent from Rcc teleop.
type ControlDatapoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream    string            `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	Timestamp int64             `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // unix epoch time in milliseconds
	Tags      map[string]string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to Data:
	//
	//	*ControlDatapoint_Text
	//	*ControlDatapoint_Numeric
	//	*ControlDatapoint_Bitset
	//	*ControlDatapoint_Twist
	//	*ControlDatapoint_Pose
	//	*ControlDatapoint_InitialPose
	//	*ControlDatapoint_Point
	//	*ControlDatapoint_GoalId
	Data isControlDatapoint_Data `protobuf_oneof:"data"`
}

func (x *ControlDatapoint) Reset() {
	*x = ControlDatapoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_datapoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlDatapoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlDatapoint) ProtoMessage() {}

func (x *ControlDatapoint) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_datapoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlDatapoint.ProtoReflect.Descriptor instead.
func (*ControlDatapoint) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_datapoint_proto_rawDescGZIP(), []int{1}
}

func (x *ControlDatapoint) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *ControlDatapoint) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ControlDatapoint) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (m *ControlDatapoint) GetData() isControlDatapoint_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ControlDatapoint) GetText() *Text {
	if x, ok := x.GetData().(*ControlDatapoint_Text); ok {
		return x.Text
	}
	return nil
}

func (x *ControlDatapoint) GetNumeric() *Numeric {
	if x, ok := x.GetData().(*ControlDatapoint_Numeric); ok {
		return x.Numeric
	}
	return nil
}

func (x *ControlDatapoint) GetBitset() *Bitset {
	if x, ok := x.GetData().(*ControlDatapoint_Bitset); ok {
		return x.Bitset
	}
	return nil
}

func (x *ControlDatapoint) GetTwist() *Twist {
	if x, ok := x.GetData().(*ControlDatapoint_Twist); ok {
		return x.Twist
	}
	return nil
}

func (x *ControlDatapoint) GetPose() []byte {
	if x, ok := x.GetData().(*ControlDatapoint_Pose); ok {
		return x.Pose
	}
	return nil
}

func (x *ControlDatapoint) GetInitialPose() []byte {
	if x, ok := x.GetData().(*ControlDatapoint_InitialPose); ok {
		return x.InitialPose
	}
	return nil
}

func (x *ControlDatapoint) GetPoint() *Point {
	if x, ok := x.GetData().(*ControlDatapoint_Point); ok {
		return x.Point
	}
	return nil
}

func (x *ControlDatapoint) GetGoalId() *GoalID {
	if x, ok := x.GetData().(*ControlDatapoint_GoalId); ok {
		return x.GoalId
	}
	return nil
}

type isControlDatapoint_Data interface {
	isControlDatapoint_Data()
}

type ControlDatapoint_Text struct {
	Text *Text `protobuf:"bytes,4,opt,name=text,proto3,oneof"`
}

type ControlDatapoint_Numeric struct {
	Numeric *Numeric `protobuf:"bytes,5,opt,name=numeric,proto3,oneof"`
}

type ControlDatapoint_Bitset struct {
	Bitset *Bitset `protobuf:"bytes,6,opt,name=bitset,proto3,oneof"`
}

type ControlDatapoint_Twist struct {
	Twist *Twist `protobuf:"bytes,7,opt,name=twist,proto3,oneof"`
}

type ControlDatapoint_Pose struct {
	Pose []byte `protobuf:"bytes,8,opt,name=pose,proto3,oneof"`
}

type ControlDatapoint_InitialPose struct {
	InitialPose []byte `protobuf:"bytes,9,opt,name=initial_pose,json=initialPose,proto3,oneof"`
}

type ControlDatapoint_Point struct {
	Point *Point `protobuf:"bytes,10,opt,name=point,proto3,oneof"`
}

type ControlDatapoint_GoalId struct {
	GoalId *GoalID `protobuf:"bytes,11,opt,name=goal_id,json=goalId,proto3,oneof"`
}

func (*ControlDatapoint_Text) isControlDatapoint_Data() {}

func (*ControlDatapoint_Numeric) isControlDatapoint_Data() {}

func (*ControlDatapoint_Bitset) isControlDatapoint_Data() {}

func (*ControlDatapoint_Twist) isControlDatapoint_Data() {}

func (*ControlDatapoint_Pose) isControlDatapoint_Data() {}

func (*ControlDatapoint_InitialPose) isControlDatapoint_Data() {}

func (*ControlDatapoint_Point) isControlDatapoint_Data() {}

func (*ControlDatapoint_GoalId) isControlDatapoint_Data() {}

var File_protos_model_v1_datapoint_proto protoreflect.FileDescriptor

var file_protos_model_v1_datapoint_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x06, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x31, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x54,
	0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x24,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x12, 0x37, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x73,
	0x65, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x53, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x53, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x06,
	0x62, 0x69, 0x74, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x69, 0x74, 0x73, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x62, 0x69, 0x74, 0x73, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x27,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x12, 0x30, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x04,
	0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x6a, 0x73,
	0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x48, 0x00, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x40, 0x0a, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x72, 0x65, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x72, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x03,
	0x72, 0x61, 0x77, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77,
	0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0xfe, 0x03, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x24, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e,
	0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69,
	0x63, 0x12, 0x2a, 0x0a, 0x06, 0x62, 0x69, 0x74, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x69, 0x74,
	0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x62, 0x69, 0x74, 0x73, 0x65, 0x74, 0x12, 0x27, 0x0a,
	0x05, 0x74, 0x77, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x77, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x05, 0x74, 0x77, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0c,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x48, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x67, 0x6f,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x49, 0x44, 0x48, 0x00, 0x52,
	0x06, 0x67, 0x6f, 0x61, 0x6c, 0x49, 0x64, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x4f, 0x41, 0x54, 0x2d, 0x52, 0x6f, 0x62, 0x6f,
	0x74, 0x69, 0x63, 0x73, 0x2f, 0x67, 0x74, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_v1_datapoint_proto_rawDescOnce sync.Once
	file_protos_model_v1_datapoint_proto_rawDescData = file_protos_model_v1_datapoint_proto_rawDesc
)

func file_protos_model_v1_datapoint_proto_rawDescGZIP() []byte {
	file_protos_model_v1_datapoint_proto_rawDescOnce.Do(func() {
		file_protos_model_v1_datapoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_v1_datapoint_proto_rawDescData)
	})
	return file_protos_model_v1_datapoint_proto_rawDescData
}

var file_protos_model_v1_datapoint_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_model_v1_datapoint_proto_goTypes = []any{
	(*Datapoint)(nil),        // 0: v1.model.Datapoint
	(*ControlDatapoint)(nil), // 1: v1.model.ControlDatapoint
	nil,                      // 2: v1.model.Datapoint.TagsEntry
	nil,                      // 3: v1.model.ControlDatapoint.TagsEntry
	(*Text)(nil),             // 4: v1.model.Text
	(*Numeric)(nil),          // 5: v1.model.Numeric
	(*NumericSet)(nil),       // 6: v1.model.NumericSet
	(*Bitset)(nil),           // 7: v1.model.Bitset
	(*File)(nil),             // 8: v1.model.File
	(*Image)(nil),            // 9: v1.model.Image
	(*PointCloud)(nil),       // 10: v1.model.PointCloud
	(*Location)(nil),         // 11: v1.model.Location
	(*Localization)(nil),     // 12: v1.model.Localization
	(*Health)(nil),           // 13: v1.model.Health
	(*Json)(nil),             // 14: v1.model.Json
	(*Battery)(nil),          // 15: v1.model.Battery
	(*Video)(nil),            // 16: v1.model.Video
	(*TransformTree)(nil),    // 17: v1.model.TransformTree
	(*Twist)(nil),            // 18: v1.model.Twist
	(*Point)(nil),            // 19: v1.model.Point
	(*GoalID)(nil),           // 20: v1.model.GoalID
}
var file_protos_model_v1_datapoint_proto_depIdxs = []int32{
	2,  // 0: v1.model.Datapoint.tags:type_name -> v1.model.Datapoint.TagsEntry
	4,  // 1: v1.model.Datapoint.text:type_name -> v1.model.Text
	5,  // 2: v1.model.Datapoint.numeric:type_name -> v1.model.Numeric
	6,  // 3: v1.model.Datapoint.numeric_set:type_name -> v1.model.NumericSet
	7,  // 4: v1.model.Datapoint.bitset:type_name -> v1.model.Bitset
	8,  // 5: v1.model.Datapoint.file:type_name -> v1.model.File
	9,  // 6: v1.model.Datapoint.image:type_name -> v1.model.Image
	10, // 7: v1.model.Datapoint.point_cloud:type_name -> v1.model.PointCloud
	11, // 8: v1.model.Datapoint.location:type_name -> v1.model.Location
	12, // 9: v1.model.Datapoint.localization:type_name -> v1.model.Localization
	13, // 10: v1.model.Datapoint.health:type_name -> v1.model.Health
	14, // 11: v1.model.Datapoint.json:type_name -> v1.model.Json
	15, // 12: v1.model.Datapoint.battery:type_name -> v1.model.Battery
	16, // 13: v1.model.Datapoint.video:type_name -> v1.model.Video
	17, // 14: v1.model.Datapoint.transform_tree:type_name -> v1.model.TransformTree
	3,  // 15: v1.model.ControlDatapoint.tags:type_name -> v1.model.ControlDatapoint.TagsEntry
	4,  // 16: v1.model.ControlDatapoint.text:type_name -> v1.model.Text
	5,  // 17: v1.model.ControlDatapoint.numeric:type_name -> v1.model.Numeric
	7,  // 18: v1.model.ControlDatapoint.bitset:type_name -> v1.model.Bitset
	18, // 19: v1.model.ControlDatapoint.twist:type_name -> v1.model.Twist
	19, // 20: v1.model.ControlDatapoint.point:type_name -> v1.model.Point
	20, // 21: v1.model.ControlDatapoint.goal_id:type_name -> v1.model.GoalID
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_protos_model_v1_datapoint_proto_init() }
func file_protos_model_v1_datapoint_proto_init() {
	if File_protos_model_v1_datapoint_proto != nil {
		return
	}
	file_protos_model_v1_file_proto_init()
	file_protos_model_v1_health_proto_init()
	file_protos_model_v1_math_proto_init()
	file_protos_model_v1_navigation_proto_init()
	file_protos_model_v1_text_proto_init()
	file_protos_model_v1_media_proto_init()
	file_protos_model_v1_ros_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_model_v1_datapoint_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Datapoint); i {
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
		file_protos_model_v1_datapoint_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ControlDatapoint); i {
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
	file_protos_model_v1_datapoint_proto_msgTypes[0].OneofWrappers = []any{
		(*Datapoint_Text)(nil),
		(*Datapoint_Numeric)(nil),
		(*Datapoint_NumericSet)(nil),
		(*Datapoint_Bitset)(nil),
		(*Datapoint_File)(nil),
		(*Datapoint_Image)(nil),
		(*Datapoint_PointCloud)(nil),
		(*Datapoint_Location)(nil),
		(*Datapoint_Localization)(nil),
		(*Datapoint_Health)(nil),
		(*Datapoint_Json)(nil),
		(*Datapoint_Battery)(nil),
		(*Datapoint_Video)(nil),
		(*Datapoint_TransformTree)(nil),
		(*Datapoint_Raw)(nil),
	}
	file_protos_model_v1_datapoint_proto_msgTypes[1].OneofWrappers = []any{
		(*ControlDatapoint_Text)(nil),
		(*ControlDatapoint_Numeric)(nil),
		(*ControlDatapoint_Bitset)(nil),
		(*ControlDatapoint_Twist)(nil),
		(*ControlDatapoint_Pose)(nil),
		(*ControlDatapoint_InitialPose)(nil),
		(*ControlDatapoint_Point)(nil),
		(*ControlDatapoint_GoalId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_model_v1_datapoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_v1_datapoint_proto_goTypes,
		DependencyIndexes: file_protos_model_v1_datapoint_proto_depIdxs,
		MessageInfos:      file_protos_model_v1_datapoint_proto_msgTypes,
	}.Build()
	File_protos_model_v1_datapoint_proto = out.File
	file_protos_model_v1_datapoint_proto_rawDesc = nil
	file_protos_model_v1_datapoint_proto_goTypes = nil
	file_protos_model_v1_datapoint_proto_depIdxs = nil
}
