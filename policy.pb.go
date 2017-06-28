// Code generated by protoc-gen-go. DO NOT EDIT.
// source: policy.proto

/*
Package protobufadapter is a generated protocol buffer package.

It is generated from these files:
	policy.proto

It has these top-level messages:
	Policy
*/
package protobufadapter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Policy struct {
	Rules []string `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Policy) GetRules() []string {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "protobufadapter.Policy")
}

func init() { proto.RegisterFile("policy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 81 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc8, 0xcf, 0xc9,
	0x4c, 0xae, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x07, 0x53, 0x49, 0xa5, 0x69, 0x89,
	0x29, 0x89, 0x05, 0x25, 0xa9, 0x45, 0x4a, 0x72, 0x5c, 0x6c, 0x01, 0x60, 0x05, 0x42, 0x22, 0x5c,
	0xac, 0x45, 0xa5, 0x39, 0xa9, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x10, 0x4e, 0x12,
	0x1b, 0x58, 0x83, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x88, 0x38, 0x61, 0x47, 0x00, 0x00,
	0x00,
}
