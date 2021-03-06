// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/mutate_error.proto

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

// Enum describing possible mutate errors.
type MutateErrorEnum_MutateError int32

const (
	// Enum unspecified.
	MutateErrorEnum_UNSPECIFIED MutateErrorEnum_MutateError = 0
	// The received error code is not known in this version.
	MutateErrorEnum_UNKNOWN MutateErrorEnum_MutateError = 1
	// Requested resource was not found.
	MutateErrorEnum_RESOURCE_NOT_FOUND MutateErrorEnum_MutateError = 3
	// Cannot mutate the same resource twice in one request.
	MutateErrorEnum_ID_EXISTS_IN_MULTIPLE_MUTATES MutateErrorEnum_MutateError = 7
	// The field's contents don't match another field that represents the same
	// data.
	MutateErrorEnum_INCONSISTENT_FIELD_VALUES MutateErrorEnum_MutateError = 8
	// Mutates are not allowed for the requested resource.
	MutateErrorEnum_MUTATE_NOT_ALLOWED MutateErrorEnum_MutateError = 9
)

var MutateErrorEnum_MutateError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "RESOURCE_NOT_FOUND",
	7: "ID_EXISTS_IN_MULTIPLE_MUTATES",
	8: "INCONSISTENT_FIELD_VALUES",
	9: "MUTATE_NOT_ALLOWED",
}
var MutateErrorEnum_MutateError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"RESOURCE_NOT_FOUND":            3,
	"ID_EXISTS_IN_MULTIPLE_MUTATES": 7,
	"INCONSISTENT_FIELD_VALUES":     8,
	"MUTATE_NOT_ALLOWED":            9,
}

func (x MutateErrorEnum_MutateError) String() string {
	return proto.EnumName(MutateErrorEnum_MutateError_name, int32(x))
}
func (MutateErrorEnum_MutateError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mutate_error_662b71f580e25d75, []int{0, 0}
}

// Container for enum describing possible mutate errors.
type MutateErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateErrorEnum) Reset()         { *m = MutateErrorEnum{} }
func (m *MutateErrorEnum) String() string { return proto.CompactTextString(m) }
func (*MutateErrorEnum) ProtoMessage()    {}
func (*MutateErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_mutate_error_662b71f580e25d75, []int{0}
}
func (m *MutateErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateErrorEnum.Unmarshal(m, b)
}
func (m *MutateErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateErrorEnum.Marshal(b, m, deterministic)
}
func (dst *MutateErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateErrorEnum.Merge(dst, src)
}
func (m *MutateErrorEnum) XXX_Size() int {
	return xxx_messageInfo_MutateErrorEnum.Size(m)
}
func (m *MutateErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MutateErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MutateErrorEnum)(nil), "google.ads.googleads.v0.errors.MutateErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.MutateErrorEnum_MutateError", MutateErrorEnum_MutateError_name, MutateErrorEnum_MutateError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/mutate_error.proto", fileDescriptor_mutate_error_662b71f580e25d75)
}

var fileDescriptor_mutate_error_662b71f580e25d75 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x6e, 0xea, 0x30,
	0x18, 0x85, 0x6f, 0x40, 0xba, 0xdc, 0x6b, 0x06, 0x2c, 0x0f, 0x95, 0x3a, 0x80, 0x54, 0x1e, 0xc0,
	0x49, 0xd5, 0xb1, 0x93, 0x21, 0x06, 0x59, 0x0d, 0x4e, 0x84, 0x6d, 0xa8, 0xaa, 0x48, 0x56, 0xda,
	0x44, 0x51, 0x25, 0xc0, 0x28, 0x06, 0xde, 0xa6, 0x4b, 0xc7, 0x6e, 0x7d, 0x89, 0x0e, 0x7d, 0xaa,
	0x2a, 0x71, 0x41, 0x2c, 0xed, 0xe4, 0xe3, 0x5f, 0xdf, 0x39, 0xbf, 0xfe, 0x03, 0xae, 0x4b, 0x63,
	0xca, 0x55, 0xe1, 0x67, 0xb9, 0xf5, 0x9d, 0xac, 0xd5, 0x21, 0xf0, 0x8b, 0xaa, 0x32, 0x95, 0xf5,
	0xd7, 0xfb, 0x5d, 0xb6, 0x2b, 0x74, 0xf3, 0xc3, 0xdb, 0xca, 0xec, 0x0c, 0x1a, 0x38, 0x0e, 0x67,
	0xb9, 0xc5, 0x27, 0x0b, 0x3e, 0x04, 0xd8, 0x59, 0x86, 0xef, 0x1e, 0xe8, 0xcd, 0x1a, 0x1b, 0xad,
	0x07, 0x74, 0xb3, 0x5f, 0x0f, 0x5f, 0x3c, 0xd0, 0x3d, 0x9b, 0xa1, 0x1e, 0xe8, 0x2a, 0x2e, 0x12,
	0x3a, 0x66, 0x13, 0x46, 0x43, 0xf8, 0x07, 0x75, 0x41, 0x47, 0xf1, 0x3b, 0x1e, 0x2f, 0x39, 0xf4,
	0xd0, 0x05, 0x40, 0x73, 0x2a, 0x62, 0x35, 0x1f, 0x53, 0xcd, 0x63, 0xa9, 0x27, 0xb1, 0xe2, 0x21,
	0x6c, 0xa3, 0x2b, 0xd0, 0x67, 0xa1, 0xa6, 0xf7, 0x4c, 0x48, 0xa1, 0x19, 0xd7, 0x33, 0x15, 0x49,
	0x96, 0x44, 0x54, 0xcf, 0x94, 0x24, 0x92, 0x0a, 0xd8, 0x41, 0x7d, 0x70, 0xc9, 0xf8, 0x38, 0xe6,
	0x82, 0x09, 0x49, 0xb9, 0xd4, 0x13, 0x46, 0xa3, 0x50, 0x2f, 0x48, 0xa4, 0xa8, 0x80, 0xff, 0xea,
	0x64, 0xc7, 0x36, 0xb9, 0x24, 0x8a, 0xe2, 0x25, 0x0d, 0xe1, 0xff, 0xd1, 0x87, 0x07, 0x86, 0x4f,
	0x66, 0x8d, 0x7f, 0x3f, 0x6d, 0x04, 0xcf, 0x6e, 0x48, 0xea, 0x32, 0x12, 0xef, 0x21, 0xfc, 0xf6,
	0x94, 0x66, 0x95, 0x6d, 0x4a, 0x6c, 0xaa, 0xd2, 0x2f, 0x8b, 0x4d, 0x53, 0xd5, 0xb1, 0xd1, 0xed,
	0xb3, 0xfd, 0xa9, 0xe0, 0x5b, 0xf7, 0xbc, 0xb6, 0xda, 0x53, 0x42, 0xde, 0x5a, 0x83, 0xa9, 0x0b,
	0x23, 0xb9, 0xc5, 0x4e, 0xd6, 0x6a, 0x11, 0xe0, 0x66, 0xa5, 0xfd, 0x3c, 0x02, 0x29, 0xc9, 0x6d,
	0x7a, 0x02, 0xd2, 0x45, 0x90, 0x3a, 0xe0, 0xf1, 0x6f, 0xb3, 0xf8, 0xe6, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x79, 0xae, 0x21, 0x09, 0xd8, 0x01, 0x00, 0x00,
}
