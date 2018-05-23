// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/test.proto

package example

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

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}
var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}
func (FOO) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_test_f11af94dba7bc255, []int{0}
}

type Test struct {
	Label                *string             `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type                 *int32              `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps                 []int64             `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	Optionalgroup        *Test_OptionalGroup `protobuf:"group,4,opt,name=OptionalGroup,json=optionalgroup" json:"optionalgroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_f11af94dba7bc255, []int{0}
}
func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (dst *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(dst, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

const Default_Test_Type int32 = 77

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func (m *Test) GetOptionalgroup() *Test_OptionalGroup {
	if m != nil {
		return m.Optionalgroup
	}
	return nil
}

type Test_OptionalGroup struct {
	RequiredField        *string  `protobuf:"bytes,5,req,name=RequiredField" json:"RequiredField,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test_OptionalGroup) Reset()         { *m = Test_OptionalGroup{} }
func (m *Test_OptionalGroup) String() string { return proto.CompactTextString(m) }
func (*Test_OptionalGroup) ProtoMessage()    {}
func (*Test_OptionalGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_f11af94dba7bc255, []int{0, 0}
}
func (m *Test_OptionalGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test_OptionalGroup.Unmarshal(m, b)
}
func (m *Test_OptionalGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test_OptionalGroup.Marshal(b, m, deterministic)
}
func (dst *Test_OptionalGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test_OptionalGroup.Merge(dst, src)
}
func (m *Test_OptionalGroup) XXX_Size() int {
	return xxx_messageInfo_Test_OptionalGroup.Size(m)
}
func (m *Test_OptionalGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_Test_OptionalGroup.DiscardUnknown(m)
}

var xxx_messageInfo_Test_OptionalGroup proto.InternalMessageInfo

func (m *Test_OptionalGroup) GetRequiredField() string {
	if m != nil && m.RequiredField != nil {
		return *m.RequiredField
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "example.Test")
	proto.RegisterType((*Test_OptionalGroup)(nil), "example.Test.OptionalGroup")
	proto.RegisterEnum("example.FOO", FOO_name, FOO_value)
}

func init() { proto.RegisterFile("proto/test.proto", fileDescriptor_test_f11af94dba7bc255) }

var fileDescriptor_test_f11af94dba7bc255 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x03, 0x33, 0x85, 0xd8, 0x53, 0x2b, 0x12, 0x73, 0x0b,
	0x72, 0x52, 0x95, 0x0e, 0x31, 0x72, 0xb1, 0x84, 0xa4, 0x16, 0x97, 0x08, 0x89, 0x70, 0xb1, 0xe6,
	0x24, 0x26, 0xa5, 0xe6, 0x48, 0x30, 0x2a, 0x30, 0x69, 0x70, 0x06, 0x41, 0x38, 0x42, 0x62, 0x5c,
	0x2c, 0x25, 0x95, 0x05, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x56, 0x4c, 0xe6, 0xe6, 0x41,
	0x60, 0xbe, 0x90, 0x10, 0x17, 0x4b, 0x51, 0x6a, 0x41, 0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0x73,
	0x10, 0x98, 0x2d, 0xe4, 0xc8, 0xc5, 0x9b, 0x5f, 0x50, 0x92, 0x99, 0x9f, 0x97, 0x98, 0x93, 0x5e,
	0x94, 0x5f, 0x5a, 0x20, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x65, 0x24, 0xad, 0x07, 0xb5, 0x4b, 0x0f,
	0x64, 0x8f, 0x9e, 0x3f, 0x54, 0x89, 0x3b, 0x48, 0x49, 0x10, 0xaa, 0x0e, 0x29, 0x53, 0x2e, 0x5e,
	0x14, 0x79, 0x21, 0x15, 0x2e, 0xde, 0xa0, 0xd4, 0xc2, 0xd2, 0xcc, 0xa2, 0xd4, 0x14, 0xb7, 0xcc,
	0xd4, 0x9c, 0x14, 0x09, 0x56, 0xb0, 0xeb, 0x50, 0x05, 0xb5, 0x78, 0xb8, 0x98, 0xdd, 0xfc, 0xfd,
	0x85, 0x58, 0xb9, 0x18, 0x23, 0x04, 0x04, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x39, 0x8a, 0x13,
	0x3f, 0xee, 0x00, 0x00, 0x00,
}
