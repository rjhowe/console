// Code generated by protoc-gen-gogo.
// source: common/bool_filter.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/bool_filter.proto
	common/timestamp.proto
	common/wrappers.proto

It has these top-level messages:
	Timestamp
	DoubleValue
	FloatValue
	Int64Value
	UInt64Value
	Int32Value
	UInt32Value
	BoolValue
	StringValue
	BytesValue
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type BoolFilter int32

const (
	BoolFilter_ALL   BoolFilter = 0
	BoolFilter_TRUE  BoolFilter = 1
	BoolFilter_FALSE BoolFilter = 2
)

var BoolFilter_name = map[int32]string{
	0: "ALL",
	1: "TRUE",
	2: "FALSE",
}
var BoolFilter_value = map[string]int32{
	"ALL":   0,
	"TRUE":  1,
	"FALSE": 2,
}

func (x BoolFilter) String() string {
	return proto.EnumName(BoolFilter_name, int32(x))
}
func (BoolFilter) EnumDescriptor() ([]byte, []int) { return fileDescriptorBoolFilter, []int{0} }

func init() {
	proto.RegisterEnum("common.BoolFilter", BoolFilter_name, BoolFilter_value)
}

var fileDescriptorBoolFilter = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x4f, 0xca, 0xcf, 0xcf, 0x89, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xc8, 0x68, 0x69, 0x71, 0x71, 0x39, 0x01, 0x25, 0xdd,
	0xc0, 0x72, 0x42, 0xec, 0x5c, 0xcc, 0x8e, 0x3e, 0x3e, 0x02, 0x0c, 0x42, 0x1c, 0x5c, 0x2c, 0x21,
	0x41, 0xa1, 0xae, 0x02, 0x8c, 0x42, 0x9c, 0x5c, 0xac, 0x6e, 0x8e, 0x3e, 0xc1, 0xae, 0x02, 0x4c,
	0x4e, 0x02, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x00, 0xe2, 0x07, 0x40, 0x3c, 0xe3, 0xb1, 0x1c, 0x43,
	0x12, 0x1b, 0xd8, 0x30, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x15, 0xe9, 0x07, 0x68,
	0x00, 0x00, 0x00,
}
