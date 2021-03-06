// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dayofweek.proto

package dayofweek // import "github.com/veith/protos/dayofweek"

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Represents a day of week.
type DayOfWeek int32

const (
	// The unspecified day-of-week.
	DayOfWeek_DAY_OF_WEEK_UNSPECIFIED DayOfWeek = 0
	// The day-of-week of Monday.
	DayOfWeek_MONDAY DayOfWeek = 1
	// The day-of-week of Tuesday.
	DayOfWeek_TUESDAY DayOfWeek = 2
	// The day-of-week of Wednesday.
	DayOfWeek_WEDNESDAY DayOfWeek = 3
	// The day-of-week of Thursday.
	DayOfWeek_THURSDAY DayOfWeek = 4
	// The day-of-week of Friday.
	DayOfWeek_FRIDAY DayOfWeek = 5
	// The day-of-week of Saturday.
	DayOfWeek_SATURDAY DayOfWeek = 6
	// The day-of-week of Sunday.
	DayOfWeek_SUNDAY DayOfWeek = 7
)

var DayOfWeek_name = map[int32]string{
	0: "DAY_OF_WEEK_UNSPECIFIED",
	1: "MONDAY",
	2: "TUESDAY",
	3: "WEDNESDAY",
	4: "THURSDAY",
	5: "FRIDAY",
	6: "SATURDAY",
	7: "SUNDAY",
}
var DayOfWeek_value = map[string]int32{
	"DAY_OF_WEEK_UNSPECIFIED": 0,
	"MONDAY":                  1,
	"TUESDAY":                 2,
	"WEDNESDAY":               3,
	"THURSDAY":                4,
	"FRIDAY":                  5,
	"SATURDAY":                6,
	"SUNDAY":                  7,
}

func (x DayOfWeek) String() string {
	return proto.EnumName(DayOfWeek_name, int32(x))
}
func (DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dayofweek_d10a94494681a2d8, []int{0}
}

func init() {
	proto.RegisterEnum("google.type.DayOfWeek", DayOfWeek_name, DayOfWeek_value)
}

func init() { proto.RegisterFile("dayofweek.proto", fileDescriptor_dayofweek_d10a94494681a2d8) }

var fileDescriptor_dayofweek_d10a94494681a2d8 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x49, 0xac, 0xcc,
	0x4f, 0x2b, 0x4f, 0x4d, 0xcd, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x2b, 0xa9, 0x2c, 0x48, 0xd5, 0x6a, 0x61, 0xe4, 0xe2, 0x74, 0x49, 0xac,
	0xf4, 0x4f, 0x0b, 0x4f, 0x4d, 0xcd, 0x16, 0x92, 0xe6, 0x12, 0x77, 0x71, 0x8c, 0x8c, 0xf7, 0x77,
	0x8b, 0x0f, 0x77, 0x75, 0xf5, 0x8e, 0x0f, 0xf5, 0x0b, 0x0e, 0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74,
	0x75, 0x11, 0x60, 0x10, 0xe2, 0xe2, 0x62, 0xf3, 0xf5, 0xf7, 0x73, 0x71, 0x8c, 0x14, 0x60, 0x14,
	0xe2, 0xe6, 0x62, 0x0f, 0x09, 0x75, 0x0d, 0x06, 0x71, 0x98, 0x84, 0x78, 0xb9, 0x38, 0xc3, 0x5d,
	0x5d, 0xfc, 0x20, 0x5c, 0x66, 0x21, 0x1e, 0x2e, 0x8e, 0x10, 0x8f, 0xd0, 0x20, 0x30, 0x8f, 0x05,
	0xa4, 0xcb, 0x2d, 0xc8, 0x13, 0xc4, 0x66, 0x05, 0xc9, 0x04, 0x3b, 0x86, 0x84, 0x06, 0x81, 0x78,
	0x6c, 0x20, 0x99, 0xe0, 0x50, 0xb0, 0x79, 0xec, 0x4e, 0x49, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x23, 0x17, 0x7f, 0x72, 0x7e, 0xae, 0x1e, 0x92, 0x2b, 0x9d,
	0xf8, 0xe0, 0x4e, 0x0c, 0x00, 0x79, 0x21, 0x80, 0x31, 0x4a, 0x3b, 0x3d, 0xb3, 0x24, 0xa3, 0x34,
	0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x2c, 0x35, 0xb3, 0x24, 0x43, 0x1f, 0xec, 0xb9, 0x62, 0x7d,
	0xb8, 0x6f, 0xad, 0xe1, 0xac, 0x45, 0x4c, 0xcc, 0xee, 0x21, 0x01, 0x49, 0x6c, 0x60, 0x15, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x4b, 0xb3, 0x2b, 0x11, 0x01, 0x00, 0x00,
}
