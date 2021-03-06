// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/region_code_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible region code errors.
type RegionCodeErrorEnum_RegionCodeError int32

const (
	// Enum unspecified.
	RegionCodeErrorEnum_UNSPECIFIED RegionCodeErrorEnum_RegionCodeError = 0
	// The received error code is not known in this version.
	RegionCodeErrorEnum_UNKNOWN RegionCodeErrorEnum_RegionCodeError = 1
	// Invalid region code.
	RegionCodeErrorEnum_INVALID_REGION_CODE RegionCodeErrorEnum_RegionCodeError = 2
)

var RegionCodeErrorEnum_RegionCodeError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_REGION_CODE",
}
var RegionCodeErrorEnum_RegionCodeError_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"INVALID_REGION_CODE": 2,
}

func (x RegionCodeErrorEnum_RegionCodeError) String() string {
	return proto.EnumName(RegionCodeErrorEnum_RegionCodeError_name, int32(x))
}
func (RegionCodeErrorEnum_RegionCodeError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_region_code_error_6722165fc661603e, []int{0, 0}
}

// Container for enum describing possible region code errors.
type RegionCodeErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegionCodeErrorEnum) Reset()         { *m = RegionCodeErrorEnum{} }
func (m *RegionCodeErrorEnum) String() string { return proto.CompactTextString(m) }
func (*RegionCodeErrorEnum) ProtoMessage()    {}
func (*RegionCodeErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_region_code_error_6722165fc661603e, []int{0}
}
func (m *RegionCodeErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegionCodeErrorEnum.Unmarshal(m, b)
}
func (m *RegionCodeErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegionCodeErrorEnum.Marshal(b, m, deterministic)
}
func (dst *RegionCodeErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionCodeErrorEnum.Merge(dst, src)
}
func (m *RegionCodeErrorEnum) XXX_Size() int {
	return xxx_messageInfo_RegionCodeErrorEnum.Size(m)
}
func (m *RegionCodeErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionCodeErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_RegionCodeErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegionCodeErrorEnum)(nil), "google.ads.googleads.v0.errors.RegionCodeErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.RegionCodeErrorEnum_RegionCodeError", RegionCodeErrorEnum_RegionCodeError_name, RegionCodeErrorEnum_RegionCodeError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/region_code_error.proto", fileDescriptor_region_code_error_6722165fc661603e)
}

var fileDescriptor_region_code_error_6722165fc661603e = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x5d, 0x05, 0x85, 0xec, 0xb0, 0xd2, 0x09, 0xde, 0x76, 0xe8, 0x03, 0xa4, 0x05, 0xc1,
	0x43, 0x3c, 0x65, 0x6d, 0xac, 0x45, 0x49, 0xcb, 0x64, 0x15, 0xa4, 0x50, 0xea, 0x12, 0xc2, 0x60,
	0xeb, 0x37, 0x12, 0xdd, 0x03, 0x79, 0xf4, 0x51, 0x7c, 0x14, 0x2f, 0xbe, 0x82, 0x34, 0x71, 0x3d,
	0x0c, 0xf4, 0x94, 0x3f, 0x7f, 0x7e, 0xff, 0x7c, 0xff, 0xef, 0x43, 0xd7, 0x0a, 0x40, 0x6d, 0x64,
	0xd4, 0x0a, 0x13, 0x39, 0xd9, 0xab, 0x7d, 0x1c, 0x49, 0xad, 0x41, 0x9b, 0x48, 0x4b, 0xb5, 0x86,
	0xae, 0x59, 0x81, 0x90, 0x8d, 0xb5, 0xf0, 0x4e, 0xc3, 0x2b, 0x04, 0x33, 0x07, 0xe3, 0x56, 0x18,
	0x3c, 0xe4, 0xf0, 0x3e, 0xc6, 0x2e, 0x17, 0x36, 0x68, 0xba, 0xb0, 0xd1, 0x04, 0x84, 0x64, 0xbd,
	0xc7, 0xba, 0xb7, 0x6d, 0x78, 0x87, 0x26, 0x47, 0x76, 0x30, 0x41, 0xe3, 0x25, 0x7f, 0x2c, 0x59,
	0x92, 0xdf, 0xe6, 0x2c, 0xf5, 0x4f, 0x82, 0x31, 0x3a, 0x5f, 0xf2, 0x7b, 0x5e, 0x3c, 0x71, 0x7f,
	0x14, 0x5c, 0xa2, 0x69, 0xce, 0x2b, 0xfa, 0x90, 0xa7, 0xcd, 0x82, 0x65, 0x79, 0xc1, 0x9b, 0xa4,
	0x48, 0x99, 0xef, 0xcd, 0xbf, 0x47, 0x28, 0x5c, 0xc1, 0x16, 0xff, 0xdf, 0x63, 0x7e, 0x71, 0x34,
	0xae, 0xec, 0xdb, 0x97, 0xa3, 0xe7, 0xf4, 0x37, 0xa7, 0x60, 0xd3, 0x76, 0x0a, 0x83, 0x56, 0x91,
	0x92, 0x9d, 0xdd, 0xed, 0x70, 0x87, 0xdd, 0xda, 0xfc, 0x75, 0x96, 0x1b, 0xf7, 0xbc, 0x7b, 0xa7,
	0x19, 0xa5, 0x1f, 0xde, 0x2c, 0x73, 0x9f, 0x51, 0x61, 0xb0, 0x93, 0xbd, 0xaa, 0x62, 0x6c, 0x47,
	0x9a, 0xcf, 0x03, 0x50, 0x53, 0x61, 0xea, 0x01, 0xa8, 0xab, 0xb8, 0x76, 0xc0, 0x97, 0x17, 0x3a,
	0x97, 0x10, 0x2a, 0x0c, 0x21, 0x03, 0x42, 0x48, 0x15, 0x13, 0xe2, 0xa0, 0x97, 0x33, 0xdb, 0xee,
	0xea, 0x27, 0x00, 0x00, 0xff, 0xff, 0x17, 0xe7, 0xf4, 0x96, 0xb3, 0x01, 0x00, 0x00,
}
