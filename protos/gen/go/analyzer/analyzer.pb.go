// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: analyzer/analyzer.proto

package analyzer_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// setting
type CreateSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker         string                 `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	StrategyName   string                 `protobuf:"bytes,2,opt,name=strategyName,proto3" json:"strategyName,omitempty"`
	Start          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End            *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	StartInsideDay *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=startInsideDay,proto3" json:"startInsideDay,omitempty"`
	EndInsideDay   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=endInsideDay,proto3" json:"endInsideDay,omitempty"`
}

func (x *CreateSettingRequest) Reset() {
	*x = CreateSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_analyzer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSettingRequest) ProtoMessage() {}

func (x *CreateSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_analyzer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSettingRequest.ProtoReflect.Descriptor instead.
func (*CreateSettingRequest) Descriptor() ([]byte, []int) {
	return file_analyzer_analyzer_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSettingRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *CreateSettingRequest) GetStrategyName() string {
	if x != nil {
		return x.StrategyName
	}
	return ""
}

func (x *CreateSettingRequest) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *CreateSettingRequest) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *CreateSettingRequest) GetStartInsideDay() *timestamppb.Timestamp {
	if x != nil {
		return x.StartInsideDay
	}
	return nil
}

func (x *CreateSettingRequest) GetEndInsideDay() *timestamppb.Timestamp {
	if x != nil {
		return x.EndInsideDay
	}
	return nil
}

type CreateSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateSettingResponse) Reset() {
	*x = CreateSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_analyzer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSettingResponse) ProtoMessage() {}

func (x *CreateSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_analyzer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSettingResponse.ProtoReflect.Descriptor instead.
func (*CreateSettingResponse) Descriptor() ([]byte, []int) {
	return file_analyzer_analyzer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSettingResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ActualSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Setting `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ActualSettingsResponse) Reset() {
	*x = ActualSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_analyzer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActualSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActualSettingsResponse) ProtoMessage() {}

func (x *ActualSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_analyzer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActualSettingsResponse.ProtoReflect.Descriptor instead.
func (*ActualSettingsResponse) Descriptor() ([]byte, []int) {
	return file_analyzer_analyzer_proto_rawDescGZIP(), []int{2}
}

func (x *ActualSettingsResponse) GetData() []*Setting {
	if x != nil {
		return x.Data
	}
	return nil
}

type Setting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ticker         string                 `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	StrategyName   string                 `protobuf:"bytes,3,opt,name=strategyName,proto3" json:"strategyName,omitempty"`
	Start          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	End            *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	StartInsideDay *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=startInsideDay,proto3" json:"startInsideDay,omitempty"`
	EndInsideDay   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=endInsideDay,proto3" json:"endInsideDay,omitempty"`
}

func (x *Setting) Reset() {
	*x = Setting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_analyzer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Setting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Setting) ProtoMessage() {}

func (x *Setting) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_analyzer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Setting.ProtoReflect.Descriptor instead.
func (*Setting) Descriptor() ([]byte, []int) {
	return file_analyzer_analyzer_proto_rawDescGZIP(), []int{3}
}

func (x *Setting) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Setting) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *Setting) GetStrategyName() string {
	if x != nil {
		return x.StrategyName
	}
	return ""
}

func (x *Setting) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Setting) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *Setting) GetStartInsideDay() *timestamppb.Timestamp {
	if x != nil {
		return x.StartInsideDay
	}
	return nil
}

func (x *Setting) GetEndInsideDay() *timestamppb.Timestamp {
	if x != nil {
		return x.EndInsideDay
	}
	return nil
}

type DeleteSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSettingRequest) Reset() {
	*x = DeleteSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_analyzer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSettingRequest) ProtoMessage() {}

func (x *DeleteSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_analyzer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSettingRequest.ProtoReflect.Descriptor instead.
func (*DeleteSettingRequest) Descriptor() ([]byte, []int) {
	return file_analyzer_analyzer_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSettingRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// back tests
type BackTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker       string                 `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	StrategyName string                 `protobuf:"bytes,2,opt,name=strategyName,proto3" json:"strategyName,omitempty"`
	Start        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *BackTestRequest) Reset() {
	*x = BackTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_analyzer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackTestRequest) ProtoMessage() {}

func (x *BackTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_analyzer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackTestRequest.ProtoReflect.Descriptor instead.
func (*BackTestRequest) Descriptor() ([]byte, []int) {
	return file_analyzer_analyzer_proto_rawDescGZIP(), []int{5}
}

func (x *BackTestRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *BackTestRequest) GetStrategyName() string {
	if x != nil {
		return x.StrategyName
	}
	return ""
}

func (x *BackTestRequest) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *BackTestRequest) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type BackTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberDials int64   `protobuf:"varint,1,opt,name=NumberDials,proto3" json:"NumberDials,omitempty"`
	PNL         float32 `protobuf:"fixed32,2,opt,name=PNL,proto3" json:"PNL,omitempty"`
	Dials       []*Dial `protobuf:"bytes,3,rep,name=dials,proto3" json:"dials,omitempty"`
}

func (x *BackTestResponse) Reset() {
	*x = BackTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_analyzer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackTestResponse) ProtoMessage() {}

func (x *BackTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_analyzer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackTestResponse.ProtoReflect.Descriptor instead.
func (*BackTestResponse) Descriptor() ([]byte, []int) {
	return file_analyzer_analyzer_proto_rawDescGZIP(), []int{6}
}

func (x *BackTestResponse) GetNumberDials() int64 {
	if x != nil {
		return x.NumberDials
	}
	return 0
}

func (x *BackTestResponse) GetPNL() float32 {
	if x != nil {
		return x.PNL
	}
	return 0
}

func (x *BackTestResponse) GetDials() []*Dial {
	if x != nil {
		return x.Dials
	}
	return nil
}

type Dial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buy    float32 `protobuf:"fixed32,1,opt,name=Buy,proto3" json:"Buy,omitempty"`
	Sell   float32 `protobuf:"fixed32,2,opt,name=Sell,proto3" json:"Sell,omitempty"`
	PNL    float32 `protobuf:"fixed32,3,opt,name=PNL,proto3" json:"PNL,omitempty"`
	Period int64   `protobuf:"varint,4,opt,name=Period,proto3" json:"Period,omitempty"`
}

func (x *Dial) Reset() {
	*x = Dial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_analyzer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dial) ProtoMessage() {}

func (x *Dial) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_analyzer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dial.ProtoReflect.Descriptor instead.
func (*Dial) Descriptor() ([]byte, []int) {
	return file_analyzer_analyzer_proto_rawDescGZIP(), []int{7}
}

func (x *Dial) GetBuy() float32 {
	if x != nil {
		return x.Buy
	}
	return 0
}

func (x *Dial) GetSell() float32 {
	if x != nil {
		return x.Sell
	}
	return 0
}

func (x *Dial) GetPNL() float32 {
	if x != nil {
		return x.PNL
	}
	return 0
}

func (x *Dial) GetPeriod() int64 {
	if x != nil {
		return x.Period
	}
	return 0
}

var File_analyzer_analyzer_proto protoreflect.FileDescriptor

var file_analyzer_analyzer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x7a, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x42, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49,
	0x6e, 0x73, 0x69, 0x64, 0x65, 0x44, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x49, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x44, 0x61, 0x79, 0x12, 0x3e, 0x0a, 0x0c, 0x65, 0x6e,
	0x64, 0x49, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x44, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x65, 0x6e,
	0x64, 0x49, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x44, 0x61, 0x79, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x16, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xb9, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x42, 0x0a, 0x0e,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x44, 0x61, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x44, 0x61, 0x79,
	0x12, 0x3e, 0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x44, 0x61, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x44, 0x61, 0x79,
	0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x0f, 0x42, 0x61, 0x63,
	0x6b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x6c, 0x0a, 0x10, 0x42, 0x61, 0x63, 0x6b,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x50, 0x4e, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x50, 0x4e, 0x4c,
	0x12, 0x24, 0x0a, 0x05, 0x64, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x52,
	0x05, 0x64, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x56, 0x0a, 0x04, 0x44, 0x69, 0x61, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x42, 0x75, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x42, 0x75, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x65, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04,
	0x53, 0x65, 0x6c, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x4e, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x50, 0x4e, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x32, 0xc0,
	0x02, 0x0a, 0x08, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x50, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x08,
	0x42, 0x61, 0x63, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x7a, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x42,
	0x61, 0x63, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analyzer_analyzer_proto_rawDescOnce sync.Once
	file_analyzer_analyzer_proto_rawDescData = file_analyzer_analyzer_proto_rawDesc
)

func file_analyzer_analyzer_proto_rawDescGZIP() []byte {
	file_analyzer_analyzer_proto_rawDescOnce.Do(func() {
		file_analyzer_analyzer_proto_rawDescData = protoimpl.X.CompressGZIP(file_analyzer_analyzer_proto_rawDescData)
	})
	return file_analyzer_analyzer_proto_rawDescData
}

var file_analyzer_analyzer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_analyzer_analyzer_proto_goTypes = []interface{}{
	(*CreateSettingRequest)(nil),   // 0: analyzer.CreateSettingRequest
	(*CreateSettingResponse)(nil),  // 1: analyzer.CreateSettingResponse
	(*ActualSettingsResponse)(nil), // 2: analyzer.ActualSettingsResponse
	(*Setting)(nil),                // 3: analyzer.Setting
	(*DeleteSettingRequest)(nil),   // 4: analyzer.DeleteSettingRequest
	(*BackTestRequest)(nil),        // 5: analyzer.BackTestRequest
	(*BackTestResponse)(nil),       // 6: analyzer.BackTestResponse
	(*Dial)(nil),                   // 7: analyzer.Dial
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_analyzer_analyzer_proto_depIdxs = []int32{
	8,  // 0: analyzer.CreateSettingRequest.start:type_name -> google.protobuf.Timestamp
	8,  // 1: analyzer.CreateSettingRequest.end:type_name -> google.protobuf.Timestamp
	8,  // 2: analyzer.CreateSettingRequest.startInsideDay:type_name -> google.protobuf.Timestamp
	8,  // 3: analyzer.CreateSettingRequest.endInsideDay:type_name -> google.protobuf.Timestamp
	3,  // 4: analyzer.ActualSettingsResponse.data:type_name -> analyzer.Setting
	8,  // 5: analyzer.Setting.start:type_name -> google.protobuf.Timestamp
	8,  // 6: analyzer.Setting.end:type_name -> google.protobuf.Timestamp
	8,  // 7: analyzer.Setting.startInsideDay:type_name -> google.protobuf.Timestamp
	8,  // 8: analyzer.Setting.endInsideDay:type_name -> google.protobuf.Timestamp
	8,  // 9: analyzer.BackTestRequest.start:type_name -> google.protobuf.Timestamp
	8,  // 10: analyzer.BackTestRequest.end:type_name -> google.protobuf.Timestamp
	7,  // 11: analyzer.BackTestResponse.dials:type_name -> analyzer.Dial
	0,  // 12: analyzer.Analyzer.CreateSetting:input_type -> analyzer.CreateSettingRequest
	9,  // 13: analyzer.Analyzer.ListActualSettings:input_type -> google.protobuf.Empty
	4,  // 14: analyzer.Analyzer.DeleteSetting:input_type -> analyzer.DeleteSettingRequest
	5,  // 15: analyzer.Analyzer.BackTest:input_type -> analyzer.BackTestRequest
	1,  // 16: analyzer.Analyzer.CreateSetting:output_type -> analyzer.CreateSettingResponse
	2,  // 17: analyzer.Analyzer.ListActualSettings:output_type -> analyzer.ActualSettingsResponse
	9,  // 18: analyzer.Analyzer.DeleteSetting:output_type -> google.protobuf.Empty
	6,  // 19: analyzer.Analyzer.BackTest:output_type -> analyzer.BackTestResponse
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_analyzer_analyzer_proto_init() }
func file_analyzer_analyzer_proto_init() {
	if File_analyzer_analyzer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_analyzer_analyzer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSettingRequest); i {
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
		file_analyzer_analyzer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSettingResponse); i {
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
		file_analyzer_analyzer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActualSettingsResponse); i {
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
		file_analyzer_analyzer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Setting); i {
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
		file_analyzer_analyzer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSettingRequest); i {
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
		file_analyzer_analyzer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackTestRequest); i {
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
		file_analyzer_analyzer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackTestResponse); i {
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
		file_analyzer_analyzer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dial); i {
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
			RawDescriptor: file_analyzer_analyzer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analyzer_analyzer_proto_goTypes,
		DependencyIndexes: file_analyzer_analyzer_proto_depIdxs,
		MessageInfos:      file_analyzer_analyzer_proto_msgTypes,
	}.Build()
	File_analyzer_analyzer_proto = out.File
	file_analyzer_analyzer_proto_rawDesc = nil
	file_analyzer_analyzer_proto_goTypes = nil
	file_analyzer_analyzer_proto_depIdxs = nil
}
