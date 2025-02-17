// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/uuid.proto

package validator_examples

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/oaago/go-proto-validators"
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

type UUIDMsg struct {
	// user_id must be a valid version 4 UUID.
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUIDMsg) Reset()         { *m = UUIDMsg{} }
func (m *UUIDMsg) String() string { return proto.CompactTextString(m) }
func (*UUIDMsg) ProtoMessage()    {}
func (*UUIDMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_0029f00507e892b3, []int{0}
}

func (m *UUIDMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUIDMsg.Unmarshal(m, b)
}
func (m *UUIDMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUIDMsg.Marshal(b, m, deterministic)
}
func (m *UUIDMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUIDMsg.Merge(m, src)
}
func (m *UUIDMsg) XXX_Size() int {
	return xxx_messageInfo_UUIDMsg.Size(m)
}
func (m *UUIDMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UUIDMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UUIDMsg proto.InternalMessageInfo

func (m *UUIDMsg) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*UUIDMsg)(nil), "validator.examples.UUIDMsg")
}

func init() { proto.RegisterFile("examples/uuid.proto", fileDescriptor_0029f00507e892b3) }

var fileDescriptor_0029f00507e892b3 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x2d, 0xcd, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x2a, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x83, 0x49, 0x4b, 0x99, 0xa5,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xe7, 0x96, 0x67, 0x96, 0x64, 0xe7,
	0x97, 0xeb, 0xa7, 0xe7, 0xeb, 0x82, 0x35, 0xe8, 0xc2, 0xd5, 0x17, 0xeb, 0x23, 0xb4, 0x82, 0xa5,
	0x94, 0x74, 0xb9, 0xd8, 0x43, 0x43, 0x3d, 0x5d, 0x7c, 0x8b, 0xd3, 0x85, 0x94, 0xb8, 0xd8, 0x4b,
	0x8b, 0x53, 0x8b, 0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0x38, 0x1f, 0xdd,
	0x97, 0x67, 0x8d, 0x60, 0x9c, 0xc0, 0xc8, 0x12, 0xc4, 0x06, 0x92, 0xf1, 0x4c, 0x49, 0x62, 0x03,
	0xeb, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x56, 0xba, 0x22, 0x98, 0x00, 0x00, 0x00,
}
