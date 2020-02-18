// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/traffic_trace.proto

package v1alpha1

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

// TrafficTrace defines trace configuration for selected dataplanes.
type TrafficTrace struct {
	// List of selectors to match dataplanes.
	Selectors []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// Configuration of the tracing.
	Conf                 *TrafficTrace_Conf `protobuf:"bytes,3,opt,name=conf,proto3" json:"conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TrafficTrace) Reset()         { *m = TrafficTrace{} }
func (m *TrafficTrace) String() string { return proto.CompactTextString(m) }
func (*TrafficTrace) ProtoMessage()    {}
func (*TrafficTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc2b4b31d8d46cbb, []int{0}
}

func (m *TrafficTrace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficTrace.Unmarshal(m, b)
}
func (m *TrafficTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficTrace.Marshal(b, m, deterministic)
}
func (m *TrafficTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficTrace.Merge(m, src)
}
func (m *TrafficTrace) XXX_Size() int {
	return xxx_messageInfo_TrafficTrace.Size(m)
}
func (m *TrafficTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficTrace.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficTrace proto.InternalMessageInfo

func (m *TrafficTrace) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *TrafficTrace) GetConf() *TrafficTrace_Conf {
	if m != nil {
		return m.Conf
	}
	return nil
}

// Configuration defines settings of the tracing.
type TrafficTrace_Conf struct {
	// Backend defined in the Mesh entity.
	Backend              string   `protobuf:"bytes,1,opt,name=backend,proto3" json:"backend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficTrace_Conf) Reset()         { *m = TrafficTrace_Conf{} }
func (m *TrafficTrace_Conf) String() string { return proto.CompactTextString(m) }
func (*TrafficTrace_Conf) ProtoMessage()    {}
func (*TrafficTrace_Conf) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc2b4b31d8d46cbb, []int{0, 0}
}

func (m *TrafficTrace_Conf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficTrace_Conf.Unmarshal(m, b)
}
func (m *TrafficTrace_Conf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficTrace_Conf.Marshal(b, m, deterministic)
}
func (m *TrafficTrace_Conf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficTrace_Conf.Merge(m, src)
}
func (m *TrafficTrace_Conf) XXX_Size() int {
	return xxx_messageInfo_TrafficTrace_Conf.Size(m)
}
func (m *TrafficTrace_Conf) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficTrace_Conf.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficTrace_Conf proto.InternalMessageInfo

func (m *TrafficTrace_Conf) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func init() {
	proto.RegisterType((*TrafficTrace)(nil), "kuma.mesh.v1alpha1.TrafficTrace")
	proto.RegisterType((*TrafficTrace_Conf)(nil), "kuma.mesh.v1alpha1.TrafficTrace.Conf")
}

func init() { proto.RegisterFile("mesh/v1alpha1/traffic_trace.proto", fileDescriptor_bc2b4b31d8d46cbb) }

var fileDescriptor_bc2b4b31d8d46cbb = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0x4d, 0x2d, 0xce,
	0xd0, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x2f, 0x29, 0x4a, 0x4c, 0x4b, 0xcb,
	0x4c, 0x8e, 0x2f, 0x29, 0x4a, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xca,
	0x2e, 0xcd, 0x4d, 0xd4, 0x03, 0xa9, 0xd3, 0x83, 0xa9, 0x93, 0x92, 0x41, 0xd5, 0x56, 0x9c, 0x9a,
	0x93, 0x9a, 0x5c, 0x92, 0x5f, 0x04, 0xd1, 0xa1, 0xb4, 0x9c, 0x91, 0x8b, 0x27, 0x04, 0x62, 0x52,
	0x08, 0xc8, 0x20, 0x21, 0x2b, 0x2e, 0x4e, 0x98, 0x92, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e,
	0x23, 0x19, 0x3d, 0x4c, 0x63, 0xf5, 0x82, 0xa1, 0x8a, 0x82, 0x10, 0xca, 0x85, 0x2c, 0xb9, 0x58,
	0x92, 0xf3, 0xf3, 0xd2, 0x24, 0x98, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x54, 0xb1, 0x69, 0x43, 0xb6,
	0x4b, 0xcf, 0x39, 0x3f, 0x2f, 0x2d, 0x08, 0xac, 0x45, 0x4a, 0x81, 0x8b, 0x05, 0xc4, 0x13, 0x92,
	0xe0, 0x62, 0x4f, 0x4a, 0x4c, 0xce, 0x4e, 0xcd, 0x4b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x82, 0x71, 0x9d, 0xb8, 0xa2, 0x38, 0x60, 0xa6, 0x24, 0xb1, 0x81, 0x1d, 0x6f, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x2b, 0xc3, 0x68, 0xa5, 0x13, 0x01, 0x00, 0x00,
}
