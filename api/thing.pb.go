// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/thing.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SetTimeSeriesDatumRequest struct {
	TenantId             int64                `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	SensorId             int64                `protobuf:"varint,2,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
	Value                float64              `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SetTimeSeriesDatumRequest) Reset()         { *m = SetTimeSeriesDatumRequest{} }
func (m *SetTimeSeriesDatumRequest) String() string { return proto.CompactTextString(m) }
func (*SetTimeSeriesDatumRequest) ProtoMessage()    {}
func (*SetTimeSeriesDatumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98595d89877c392b, []int{0}
}

func (m *SetTimeSeriesDatumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTimeSeriesDatumRequest.Unmarshal(m, b)
}
func (m *SetTimeSeriesDatumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTimeSeriesDatumRequest.Marshal(b, m, deterministic)
}
func (m *SetTimeSeriesDatumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTimeSeriesDatumRequest.Merge(m, src)
}
func (m *SetTimeSeriesDatumRequest) XXX_Size() int {
	return xxx_messageInfo_SetTimeSeriesDatumRequest.Size(m)
}
func (m *SetTimeSeriesDatumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTimeSeriesDatumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetTimeSeriesDatumRequest proto.InternalMessageInfo

func (m *SetTimeSeriesDatumRequest) GetTenantId() int64 {
	if m != nil {
		return m.TenantId
	}
	return 0
}

func (m *SetTimeSeriesDatumRequest) GetSensorId() int64 {
	if m != nil {
		return m.SensorId
	}
	return 0
}

func (m *SetTimeSeriesDatumRequest) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *SetTimeSeriesDatumRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type SetTimeSeriesDataRequest struct {
	Data                 []*SetTimeSeriesDatumRequest `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SetTimeSeriesDataRequest) Reset()         { *m = SetTimeSeriesDataRequest{} }
func (m *SetTimeSeriesDataRequest) String() string { return proto.CompactTextString(m) }
func (*SetTimeSeriesDataRequest) ProtoMessage()    {}
func (*SetTimeSeriesDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98595d89877c392b, []int{1}
}

func (m *SetTimeSeriesDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTimeSeriesDataRequest.Unmarshal(m, b)
}
func (m *SetTimeSeriesDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTimeSeriesDataRequest.Marshal(b, m, deterministic)
}
func (m *SetTimeSeriesDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTimeSeriesDataRequest.Merge(m, src)
}
func (m *SetTimeSeriesDataRequest) XXX_Size() int {
	return xxx_messageInfo_SetTimeSeriesDataRequest.Size(m)
}
func (m *SetTimeSeriesDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTimeSeriesDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetTimeSeriesDataRequest proto.InternalMessageInfo

func (m *SetTimeSeriesDataRequest) GetData() []*SetTimeSeriesDatumRequest {
	if m != nil {
		return m.Data
	}
	return nil
}

type SetTimeSeriesDataResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTimeSeriesDataResponse) Reset()         { *m = SetTimeSeriesDataResponse{} }
func (m *SetTimeSeriesDataResponse) String() string { return proto.CompactTextString(m) }
func (*SetTimeSeriesDataResponse) ProtoMessage()    {}
func (*SetTimeSeriesDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98595d89877c392b, []int{2}
}

func (m *SetTimeSeriesDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTimeSeriesDataResponse.Unmarshal(m, b)
}
func (m *SetTimeSeriesDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTimeSeriesDataResponse.Marshal(b, m, deterministic)
}
func (m *SetTimeSeriesDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTimeSeriesDataResponse.Merge(m, src)
}
func (m *SetTimeSeriesDataResponse) XXX_Size() int {
	return xxx_messageInfo_SetTimeSeriesDataResponse.Size(m)
}
func (m *SetTimeSeriesDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTimeSeriesDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetTimeSeriesDataResponse proto.InternalMessageInfo

func (m *SetTimeSeriesDataResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *SetTimeSeriesDataResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SetSensorRequest struct {
	TenantId             int64    `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	ThingId              int64    `protobuf:"varint,2,opt,name=thing_id,json=thingId,proto3" json:"thing_id,omitempty"`
	TypeId               int32    `protobuf:"varint,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetSensorRequest) Reset()         { *m = SetSensorRequest{} }
func (m *SetSensorRequest) String() string { return proto.CompactTextString(m) }
func (*SetSensorRequest) ProtoMessage()    {}
func (*SetSensorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98595d89877c392b, []int{3}
}

func (m *SetSensorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetSensorRequest.Unmarshal(m, b)
}
func (m *SetSensorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetSensorRequest.Marshal(b, m, deterministic)
}
func (m *SetSensorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSensorRequest.Merge(m, src)
}
func (m *SetSensorRequest) XXX_Size() int {
	return xxx_messageInfo_SetSensorRequest.Size(m)
}
func (m *SetSensorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSensorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetSensorRequest proto.InternalMessageInfo

func (m *SetSensorRequest) GetTenantId() int64 {
	if m != nil {
		return m.TenantId
	}
	return 0
}

func (m *SetSensorRequest) GetThingId() int64 {
	if m != nil {
		return m.ThingId
	}
	return 0
}

func (m *SetSensorRequest) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type SetSensorResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetSensorResponse) Reset()         { *m = SetSensorResponse{} }
func (m *SetSensorResponse) String() string { return proto.CompactTextString(m) }
func (*SetSensorResponse) ProtoMessage()    {}
func (*SetSensorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98595d89877c392b, []int{4}
}

func (m *SetSensorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetSensorResponse.Unmarshal(m, b)
}
func (m *SetSensorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetSensorResponse.Marshal(b, m, deterministic)
}
func (m *SetSensorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSensorResponse.Merge(m, src)
}
func (m *SetSensorResponse) XXX_Size() int {
	return xxx_messageInfo_SetSensorResponse.Size(m)
}
func (m *SetSensorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSensorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetSensorResponse proto.InternalMessageInfo

func (m *SetSensorResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *SetSensorResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SetThingRequest struct {
	TenantId             int64    `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetThingRequest) Reset()         { *m = SetThingRequest{} }
func (m *SetThingRequest) String() string { return proto.CompactTextString(m) }
func (*SetThingRequest) ProtoMessage()    {}
func (*SetThingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98595d89877c392b, []int{5}
}

func (m *SetThingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetThingRequest.Unmarshal(m, b)
}
func (m *SetThingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetThingRequest.Marshal(b, m, deterministic)
}
func (m *SetThingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetThingRequest.Merge(m, src)
}
func (m *SetThingRequest) XXX_Size() int {
	return xxx_messageInfo_SetThingRequest.Size(m)
}
func (m *SetThingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetThingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetThingRequest proto.InternalMessageInfo

func (m *SetThingRequest) GetTenantId() int64 {
	if m != nil {
		return m.TenantId
	}
	return 0
}

func (m *SetThingRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SetThingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SetThingResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetThingResponse) Reset()         { *m = SetThingResponse{} }
func (m *SetThingResponse) String() string { return proto.CompactTextString(m) }
func (*SetThingResponse) ProtoMessage()    {}
func (*SetThingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98595d89877c392b, []int{6}
}

func (m *SetThingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetThingResponse.Unmarshal(m, b)
}
func (m *SetThingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetThingResponse.Marshal(b, m, deterministic)
}
func (m *SetThingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetThingResponse.Merge(m, src)
}
func (m *SetThingResponse) XXX_Size() int {
	return xxx_messageInfo_SetThingResponse.Size(m)
}
func (m *SetThingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetThingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetThingResponse proto.InternalMessageInfo

func (m *SetThingResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *SetThingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SetTimeSeriesDatumRequest)(nil), "api.SetTimeSeriesDatumRequest")
	proto.RegisterType((*SetTimeSeriesDataRequest)(nil), "api.SetTimeSeriesDataRequest")
	proto.RegisterType((*SetTimeSeriesDataResponse)(nil), "api.SetTimeSeriesDataResponse")
	proto.RegisterType((*SetSensorRequest)(nil), "api.SetSensorRequest")
	proto.RegisterType((*SetSensorResponse)(nil), "api.SetSensorResponse")
	proto.RegisterType((*SetThingRequest)(nil), "api.SetThingRequest")
	proto.RegisterType((*SetThingResponse)(nil), "api.SetThingResponse")
}

func init() { proto.RegisterFile("api/thing.proto", fileDescriptor_98595d89877c392b) }

var fileDescriptor_98595d89877c392b = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xec, 0x36, 0x3f, 0xd3, 0xc3, 0x82, 0xd5, 0x9f, 0x34, 0x88, 0x12, 0xe5, 0x94, 0x53,
	0x56, 0x0a, 0x17, 0x90, 0x38, 0x96, 0x43, 0x0e, 0xe5, 0xe0, 0xec, 0x8d, 0x03, 0x72, 0x9b, 0x21,
	0x58, 0x34, 0x89, 0x59, 0x3b, 0x48, 0x3c, 0x10, 0xef, 0xc5, 0xa3, 0x20, 0xdb, 0x75, 0xba, 0x5d,
	0xa8, 0x54, 0xf5, 0xe6, 0x6f, 0x66, 0xec, 0xf9, 0x7e, 0x0c, 0x2b, 0x2e, 0xc5, 0x5a, 0x7f, 0x13,
	0x43, 0x57, 0xca, 0xed, 0xa8, 0x47, 0xba, 0xe0, 0x52, 0xa4, 0x6f, 0xba, 0x71, 0xec, 0x6e, 0x70,
	0x6d, 0x4b, 0x57, 0xd3, 0xd7, 0xb5, 0x16, 0x3d, 0x2a, 0xcd, 0x7b, 0xe9, 0xa6, 0xf2, 0xdf, 0x04,
	0xce, 0x1a, 0xd4, 0x1b, 0xd1, 0x63, 0x83, 0x5b, 0x81, 0xea, 0x82, 0xeb, 0xa9, 0x67, 0xf8, 0x63,
	0x42, 0xa5, 0xe9, 0x2b, 0x88, 0x35, 0x0e, 0x7c, 0xd0, 0x5f, 0x44, 0x9b, 0x90, 0x8c, 0x14, 0x0b,
	0x16, 0xb9, 0x42, 0xdd, 0x9a, 0xa6, 0xc2, 0x41, 0x8d, 0x5b, 0xd3, 0x7c, 0xee, 0x9a, 0xae, 0x50,
	0xb7, 0xf4, 0x08, 0x0e, 0x7e, 0xf2, 0x9b, 0x09, 0x93, 0x45, 0x46, 0x0a, 0xc2, 0x1c, 0xa0, 0xef,
	0x20, 0x9e, 0x09, 0x24, 0xcb, 0x8c, 0x14, 0x87, 0x55, 0x5a, 0x3a, 0x8a, 0xa5, 0xa7, 0x58, 0x6e,
	0xfc, 0x04, 0xbb, 0x1b, 0xce, 0x3f, 0x41, 0xb2, 0x4f, 0x93, 0x7b, 0x96, 0x15, 0x2c, 0x5b, 0xae,
	0x79, 0x42, 0xb2, 0x45, 0x71, 0x58, 0x9d, 0x97, 0x5c, 0x8a, 0xf2, 0x41, 0x4d, 0xcc, 0xce, 0xe6,
	0x97, 0xff, 0xca, 0xe6, 0x0c, 0x95, 0x1c, 0x07, 0x85, 0xf4, 0x04, 0x02, 0xa5, 0xb9, 0x9e, 0x94,
	0xd5, 0x1c, 0xb1, 0x5b, 0x44, 0x13, 0x08, 0x7b, 0x54, 0x8a, 0x77, 0x68, 0xf5, 0xc6, 0xcc, 0xc3,
	0xfc, 0x1a, 0x5e, 0x34, 0xa8, 0x1b, 0xab, 0xfe, 0x51, 0xe6, 0x9d, 0x41, 0x64, 0xc3, 0xba, 0xf3,
	0x2e, 0xb4, 0xb8, 0x6e, 0xe9, 0x29, 0x84, 0xfa, 0x97, 0x44, 0xd3, 0x31, 0xe6, 0x1d, 0xb0, 0xc0,
	0xc0, 0xba, 0xcd, 0x3f, 0xc2, 0xcb, 0x9d, 0x25, 0x4f, 0xe6, 0xfa, 0x19, 0x56, 0x46, 0xba, 0xd9,
	0xf6, 0x28, 0xaa, 0xa7, 0x10, 0x4e, 0x0a, 0x77, 0x52, 0x0e, 0x0c, 0xac, 0x5b, 0x4a, 0x61, 0x39,
	0xf0, 0xde, 0x45, 0x1c, 0x33, 0x7b, 0xce, 0x2f, 0xac, 0x11, 0xb7, 0x8f, 0x3f, 0x95, 0x62, 0xf5,
	0x87, 0xc0, 0xea, 0x52, 0x7c, 0xe7, 0x72, 0x1c, 0xc4, 0xb5, 0xb2, 0xaf, 0xd1, 0x8d, 0x55, 0x7f,
	0x3f, 0x31, 0xfa, 0xfa, 0xbf, 0x61, 0xfb, 0x9f, 0x91, 0x9e, 0x3f, 0xd4, 0x76, 0xcc, 0xf2, 0x67,
	0xf4, 0x03, 0xc4, 0xb3, 0xa7, 0xf4, 0xd8, 0x8f, 0xdf, 0x0b, 0x32, 0x3d, 0xd9, 0x2f, 0xcf, 0xb7,
	0xdf, 0x43, 0xe4, 0xd5, 0xd2, 0xa3, 0x79, 0xd7, 0x8e, 0xb3, 0xe9, 0xf1, 0x5e, 0xd5, 0x5f, 0xbd,
	0x0a, 0xec, 0x7f, 0x7f, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0x88, 0x07, 0x9f, 0xff, 0xb8, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MikaponicsThingClient is the client API for MikaponicsThing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MikaponicsThingClient interface {
	SetTimeSeriesData(ctx context.Context, in *SetTimeSeriesDataRequest, opts ...grpc.CallOption) (*SetTimeSeriesDataResponse, error)
	SetSensor(ctx context.Context, in *SetSensorRequest, opts ...grpc.CallOption) (*SetSensorResponse, error)
	SetThing(ctx context.Context, in *SetThingRequest, opts ...grpc.CallOption) (*SetThingResponse, error)
}

type mikaponicsThingClient struct {
	cc *grpc.ClientConn
}

func NewMikaponicsThingClient(cc *grpc.ClientConn) MikaponicsThingClient {
	return &mikaponicsThingClient{cc}
}

func (c *mikaponicsThingClient) SetTimeSeriesData(ctx context.Context, in *SetTimeSeriesDataRequest, opts ...grpc.CallOption) (*SetTimeSeriesDataResponse, error) {
	out := new(SetTimeSeriesDataResponse)
	err := c.cc.Invoke(ctx, "/api.MikaponicsThing/SetTimeSeriesData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mikaponicsThingClient) SetSensor(ctx context.Context, in *SetSensorRequest, opts ...grpc.CallOption) (*SetSensorResponse, error) {
	out := new(SetSensorResponse)
	err := c.cc.Invoke(ctx, "/api.MikaponicsThing/SetSensor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mikaponicsThingClient) SetThing(ctx context.Context, in *SetThingRequest, opts ...grpc.CallOption) (*SetThingResponse, error) {
	out := new(SetThingResponse)
	err := c.cc.Invoke(ctx, "/api.MikaponicsThing/SetThing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MikaponicsThingServer is the server API for MikaponicsThing service.
type MikaponicsThingServer interface {
	SetTimeSeriesData(context.Context, *SetTimeSeriesDataRequest) (*SetTimeSeriesDataResponse, error)
	SetSensor(context.Context, *SetSensorRequest) (*SetSensorResponse, error)
	SetThing(context.Context, *SetThingRequest) (*SetThingResponse, error)
}

// UnimplementedMikaponicsThingServer can be embedded to have forward compatible implementations.
type UnimplementedMikaponicsThingServer struct {
}

func (*UnimplementedMikaponicsThingServer) SetTimeSeriesData(ctx context.Context, req *SetTimeSeriesDataRequest) (*SetTimeSeriesDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTimeSeriesData not implemented")
}
func (*UnimplementedMikaponicsThingServer) SetSensor(ctx context.Context, req *SetSensorRequest) (*SetSensorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSensor not implemented")
}
func (*UnimplementedMikaponicsThingServer) SetThing(ctx context.Context, req *SetThingRequest) (*SetThingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetThing not implemented")
}

func RegisterMikaponicsThingServer(s *grpc.Server, srv MikaponicsThingServer) {
	s.RegisterService(&_MikaponicsThing_serviceDesc, srv)
}

func _MikaponicsThing_SetTimeSeriesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTimeSeriesDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MikaponicsThingServer).SetTimeSeriesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MikaponicsThing/SetTimeSeriesData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MikaponicsThingServer).SetTimeSeriesData(ctx, req.(*SetTimeSeriesDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MikaponicsThing_SetSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MikaponicsThingServer).SetSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MikaponicsThing/SetSensor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MikaponicsThingServer).SetSensor(ctx, req.(*SetSensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MikaponicsThing_SetThing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetThingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MikaponicsThingServer).SetThing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MikaponicsThing/SetThing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MikaponicsThingServer).SetThing(ctx, req.(*SetThingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MikaponicsThing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MikaponicsThing",
	HandlerType: (*MikaponicsThingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetTimeSeriesData",
			Handler:    _MikaponicsThing_SetTimeSeriesData_Handler,
		},
		{
			MethodName: "SetSensor",
			Handler:    _MikaponicsThing_SetSensor_Handler,
		},
		{
			MethodName: "SetThing",
			Handler:    _MikaponicsThing_SetThing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/thing.proto",
}
