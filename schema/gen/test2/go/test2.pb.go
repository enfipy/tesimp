// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test2.proto

package test2

import (
	fmt "fmt"
	_go "github.com/enfipy/tesimp/schema/gen/common/go"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Test2 struct {
	Test                 *_go.TestCommon `protobuf:"bytes,1,opt,name=test,proto3" json:"test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Test2) Reset()         { *m = Test2{} }
func (m *Test2) String() string { return proto.CompactTextString(m) }
func (*Test2) ProtoMessage()    {}
func (*Test2) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e5535eff96cd18, []int{0}
}

func (m *Test2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test2.Unmarshal(m, b)
}
func (m *Test2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test2.Marshal(b, m, deterministic)
}
func (m *Test2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test2.Merge(m, src)
}
func (m *Test2) XXX_Size() int {
	return xxx_messageInfo_Test2.Size(m)
}
func (m *Test2) XXX_DiscardUnknown() {
	xxx_messageInfo_Test2.DiscardUnknown(m)
}

var xxx_messageInfo_Test2 proto.InternalMessageInfo

func (m *Test2) GetTest() *_go.TestCommon {
	if m != nil {
		return m.Test
	}
	return nil
}

func init() {
	proto.RegisterType((*Test2)(nil), "test2.Test2")
}

func init() { proto.RegisterFile("test2.proto", fileDescriptor_61e5535eff96cd18) }

var fileDescriptor_61e5535eff96cd18 = []byte{
	// 92 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x49, 0x2d, 0x2e,
	0x31, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x78, 0x92, 0xf3, 0x73,
	0x73, 0xf3, 0xf3, 0x20, 0x82, 0x4a, 0xfa, 0x5c, 0xac, 0x21, 0x20, 0x61, 0x21, 0x35, 0x2e, 0x16,
	0x90, 0xbc, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x90, 0x1e, 0x54, 0x15, 0x48, 0xd2, 0x19,
	0xcc, 0x0c, 0x02, 0xcb, 0x3b, 0xb1, 0x47, 0x41, 0xcc, 0x49, 0x62, 0x03, 0x1b, 0x60, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0x90, 0xf3, 0xb1, 0x64, 0x00, 0x00, 0x00,
}
