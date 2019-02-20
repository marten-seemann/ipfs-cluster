// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package api_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Pin_PinType int32

const (
	Pin_BadType        Pin_PinType = 0
	Pin_DataType       Pin_PinType = 1
	Pin_MetaType       Pin_PinType = 2
	Pin_ClusterDAGType Pin_PinType = 3
	Pin_ShardType      Pin_PinType = 4
)

var Pin_PinType_name = map[int32]string{
	0: "BadType",
	1: "DataType",
	2: "MetaType",
	3: "ClusterDAGType",
	4: "ShardType",
}

var Pin_PinType_value = map[string]int32{
	"BadType":        0,
	"DataType":       1,
	"MetaType":       2,
	"ClusterDAGType": 3,
	"ShardType":      4,
}

func (x Pin_PinType) String() string {
	return proto.EnumName(Pin_PinType_name, int32(x))
}

func (Pin_PinType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0, 0}
}

type Pin struct {
	Cid                  []byte            `protobuf:"bytes,1,opt,name=Cid,proto3" json:"Cid,omitempty"`
	Type                 Pin_PinType       `protobuf:"varint,2,opt,name=Type,proto3,enum=api.pb.Pin_PinType" json:"Type,omitempty"`
	Allocations          [][]byte          `protobuf:"bytes,3,rep,name=Allocations,proto3" json:"Allocations,omitempty"`
	MaxDepth             int32             `protobuf:"zigzag32,4,opt,name=MaxDepth,proto3" json:"MaxDepth,omitempty"`
	Reference            []byte            `protobuf:"bytes,5,opt,name=Reference,proto3" json:"Reference,omitempty"`
	ReplicationFactorMin int32             `protobuf:"zigzag32,8,opt,name=ReplicationFactorMin,proto3" json:"ReplicationFactorMin,omitempty"`
	ReplicationFactorMax int32             `protobuf:"zigzag32,9,opt,name=ReplicationFactorMax,proto3" json:"ReplicationFactorMax,omitempty"`
	Name                 string            `protobuf:"bytes,10,opt,name=Name,proto3" json:"Name,omitempty"`
	ShardSize            uint64            `protobuf:"varint,11,opt,name=ShardSize,proto3" json:"ShardSize,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,12,rep,name=Metadata,proto3" json:"Metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Pin) Reset()         { *m = Pin{} }
func (m *Pin) String() string { return proto.CompactTextString(m) }
func (*Pin) ProtoMessage()    {}
func (*Pin) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *Pin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pin.Unmarshal(m, b)
}
func (m *Pin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pin.Marshal(b, m, deterministic)
}
func (m *Pin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pin.Merge(m, src)
}
func (m *Pin) XXX_Size() int {
	return xxx_messageInfo_Pin.Size(m)
}
func (m *Pin) XXX_DiscardUnknown() {
	xxx_messageInfo_Pin.DiscardUnknown(m)
}

var xxx_messageInfo_Pin proto.InternalMessageInfo

func (m *Pin) GetCid() []byte {
	if m != nil {
		return m.Cid
	}
	return nil
}

func (m *Pin) GetType() Pin_PinType {
	if m != nil {
		return m.Type
	}
	return Pin_BadType
}

func (m *Pin) GetAllocations() [][]byte {
	if m != nil {
		return m.Allocations
	}
	return nil
}

func (m *Pin) GetMaxDepth() int32 {
	if m != nil {
		return m.MaxDepth
	}
	return 0
}

func (m *Pin) GetReference() []byte {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *Pin) GetReplicationFactorMin() int32 {
	if m != nil {
		return m.ReplicationFactorMin
	}
	return 0
}

func (m *Pin) GetReplicationFactorMax() int32 {
	if m != nil {
		return m.ReplicationFactorMax
	}
	return 0
}

func (m *Pin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pin) GetShardSize() uint64 {
	if m != nil {
		return m.ShardSize
	}
	return 0
}

func (m *Pin) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.pb.Pin_PinType", Pin_PinType_name, Pin_PinType_value)
	proto.RegisterType((*Pin)(nil), "api.pb.Pin")
	proto.RegisterMapType((map[string]string)(nil), "api.pb.Pin.MetadataEntry")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x0f, 0xd2, 0x30,
	0x18, 0xb6, 0xac, 0xb0, 0xed, 0xdd, 0x20, 0xf3, 0x95, 0x43, 0x25, 0x1e, 0x1a, 0x2e, 0xee, 0xb4,
	0x03, 0xc6, 0xc4, 0xe8, 0x09, 0x41, 0x4d, 0x48, 0x30, 0xa4, 0xe8, 0x0f, 0x28, 0xac, 0x86, 0xc6,
	0xb9, 0x2d, 0xa3, 0x18, 0xe6, 0x7f, 0xf0, 0x3f, 0x9b, 0x76, 0x08, 0x98, 0xe0, 0xa1, 0xc9, 0xf3,
	0xf1, 0x7e, 0xb4, 0x4f, 0x0a, 0x91, 0x69, 0x6b, 0x75, 0xcc, 0xea, 0xa6, 0x32, 0x15, 0x0e, 0x64,
	0xad, 0xb3, 0x7a, 0x37, 0xfd, 0x4d, 0xc1, 0xdb, 0xe8, 0x12, 0x13, 0xf0, 0x16, 0x3a, 0x67, 0x84,
	0x93, 0x34, 0x16, 0x16, 0xe2, 0x4b, 0xa0, 0x5f, 0xda, 0x5a, 0xb1, 0x1e, 0x27, 0xe9, 0x68, 0xf6,
	0x2c, 0xeb, 0x1a, 0xb2, 0x8d, 0x2e, 0xed, 0xb1, 0x96, 0x70, 0x05, 0xc8, 0x21, 0x9a, 0x17, 0x45,
	0xb5, 0x97, 0x46, 0x57, 0xe5, 0x91, 0x79, 0xdc, 0x4b, 0x63, 0x71, 0x2f, 0xe1, 0x04, 0x82, 0xb5,
	0x3c, 0x2f, 0x55, 0x6d, 0x0e, 0x8c, 0x72, 0x92, 0x3e, 0x15, 0x57, 0x8e, 0x2f, 0x20, 0x14, 0xea,
	0x9b, 0x6a, 0x54, 0xb9, 0x57, 0xac, 0xef, 0xd6, 0xdf, 0x04, 0x9c, 0xc1, 0x58, 0xa8, 0xba, 0xd0,
	0xdd, 0xa4, 0x8f, 0x72, 0x6f, 0xaa, 0x66, 0xad, 0x4b, 0x16, 0xb8, 0x29, 0x0f, 0xbd, 0xc7, 0x3d,
	0xf2, 0xcc, 0xc2, 0xff, 0xf5, 0xc8, 0x33, 0x22, 0xd0, 0xcf, 0xf2, 0x87, 0x62, 0xc0, 0x49, 0x1a,
	0x0a, 0x87, 0xed, 0xcd, 0xb6, 0x07, 0xd9, 0xe4, 0x5b, 0xfd, 0x4b, 0xb1, 0x88, 0x93, 0x94, 0x8a,
	0x9b, 0x80, 0xaf, 0x21, 0x58, 0x2b, 0x23, 0x73, 0x69, 0x24, 0x8b, 0xb9, 0x97, 0x46, 0xb3, 0xe7,
	0xf7, 0x11, 0xfd, 0xf5, 0x3e, 0x94, 0xa6, 0x69, 0xc5, 0xb5, 0x74, 0xf2, 0x0e, 0x86, 0xff, 0x58,
	0x36, 0xf8, 0xef, 0xaa, 0x75, 0xc1, 0x87, 0xc2, 0x42, 0x1c, 0x43, 0xff, 0xa7, 0x2c, 0x4e, 0x5d,
	0xf2, 0xa1, 0xe8, 0xc8, 0xdb, 0xde, 0x1b, 0x32, 0xfd, 0x0a, 0xfe, 0x25, 0x7a, 0x8c, 0xc0, 0x7f,
	0x2f, 0x73, 0x0b, 0x93, 0x27, 0x18, 0x43, 0xb0, 0x94, 0x46, 0x3a, 0x46, 0x2c, 0xb3, 0x2b, 0x1c,
	0xeb, 0x21, 0xc2, 0x68, 0x51, 0x9c, 0x8e, 0x46, 0x35, 0xcb, 0xf9, 0x27, 0xa7, 0x79, 0x38, 0xbc,
	0xbc, 0xcc, 0x51, 0xba, 0xa2, 0xc1, 0x20, 0xf1, 0x57, 0x34, 0xf0, 0x93, 0x60, 0x37, 0x70, 0xdf,
	0xe3, 0xd5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x03, 0x06, 0x76, 0x2d, 0x02, 0x00, 0x00,
}
