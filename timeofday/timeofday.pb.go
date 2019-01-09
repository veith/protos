// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: timeofday.proto

package timeofday // import "github.com/veith/protos/timeofday"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Represents a time of day. The date and time zone are either not significant
// or are specified elsewhere. An API may chose to allow leap seconds. Related
// types are [google.type.Date][google.type.Date] and `google.protobuf.Timestamp`.
type TimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
	// to allow the value "24:00:00" for scenarios like business closing time.
	Hours int32 `protobuf:"varint,1,opt,name=hours,proto3" json:"hours,omitempty"`
	// Minutes of hour of day. Must be from 0 to 59.
	Minutes int32 `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may
	// allow the value 60 if it allows leap-seconds.
	Seconds int32 `protobuf:"varint,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Nanos                int32    `protobuf:"varint,4,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeOfDay) Reset()         { *m = TimeOfDay{} }
func (m *TimeOfDay) String() string { return proto.CompactTextString(m) }
func (*TimeOfDay) ProtoMessage()    {}
func (*TimeOfDay) Descriptor() ([]byte, []int) {
	return fileDescriptor_timeofday_2f811bd12be2d1e1, []int{0}
}
func (m *TimeOfDay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimeOfDay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimeOfDay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TimeOfDay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeOfDay.Merge(dst, src)
}
func (m *TimeOfDay) XXX_Size() int {
	return m.Size()
}
func (m *TimeOfDay) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeOfDay.DiscardUnknown(m)
}

var xxx_messageInfo_TimeOfDay proto.InternalMessageInfo

func (m *TimeOfDay) GetHours() int32 {
	if m != nil {
		return m.Hours
	}
	return 0
}

func (m *TimeOfDay) GetMinutes() int32 {
	if m != nil {
		return m.Minutes
	}
	return 0
}

func (m *TimeOfDay) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *TimeOfDay) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterType((*TimeOfDay)(nil), "google.type.TimeOfDay")
}
func (m *TimeOfDay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeOfDay) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Hours != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTimeofday(dAtA, i, uint64(m.Hours))
	}
	if m.Minutes != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTimeofday(dAtA, i, uint64(m.Minutes))
	}
	if m.Seconds != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTimeofday(dAtA, i, uint64(m.Seconds))
	}
	if m.Nanos != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintTimeofday(dAtA, i, uint64(m.Nanos))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTimeofday(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TimeOfDay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Hours != 0 {
		n += 1 + sovTimeofday(uint64(m.Hours))
	}
	if m.Minutes != 0 {
		n += 1 + sovTimeofday(uint64(m.Minutes))
	}
	if m.Seconds != 0 {
		n += 1 + sovTimeofday(uint64(m.Seconds))
	}
	if m.Nanos != 0 {
		n += 1 + sovTimeofday(uint64(m.Nanos))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTimeofday(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTimeofday(x uint64) (n int) {
	return sovTimeofday(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TimeOfDay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimeofday
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimeOfDay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeOfDay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hours", wireType)
			}
			m.Hours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeofday
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hours |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minutes", wireType)
			}
			m.Minutes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeofday
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Minutes |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeofday
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seconds |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nanos", wireType)
			}
			m.Nanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeofday
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nanos |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimeofday(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimeofday
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTimeofday(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimeofday
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTimeofday
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTimeofday
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTimeofday
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTimeofday
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTimeofday(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTimeofday = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimeofday   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("timeofday.proto", fileDescriptor_timeofday_2f811bd12be2d1e1) }

var fileDescriptor_timeofday_2f811bd12be2d1e1 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc9, 0xcc, 0x4d,
	0xcd, 0x4f, 0x4b, 0x49, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x2b, 0xa9, 0x2c, 0x48, 0x55, 0xca, 0xe6, 0xe2, 0x0c, 0xc9, 0xcc, 0x4d,
	0xf5, 0x4f, 0x73, 0x49, 0xac, 0x14, 0x12, 0xe1, 0x62, 0xcd, 0xc8, 0x2f, 0x2d, 0x2a, 0x96, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0x70, 0x84, 0x24, 0xb8, 0xd8, 0x73, 0x33, 0xf3, 0x4a, 0x4b,
	0x52, 0x8b, 0x25, 0x98, 0xc0, 0xe2, 0x30, 0x2e, 0x48, 0xa6, 0x38, 0x35, 0x39, 0x3f, 0x2f, 0xa5,
	0x58, 0x82, 0x19, 0x22, 0x03, 0xe5, 0x82, 0x4c, 0xca, 0x4b, 0xcc, 0xcb, 0x2f, 0x96, 0x60, 0x81,
	0x98, 0x04, 0xe6, 0x38, 0x25, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x5c, 0xfc, 0xc9, 0xf9, 0xb9, 0x7a, 0x48, 0x6e, 0x71, 0xe2, 0x83, 0xbb, 0x24, 0x00,
	0xe4, 0xd0, 0x00, 0xc6, 0x28, 0xed, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0xfd, 0xb2, 0xd4, 0xcc, 0x92, 0x0c, 0x7d, 0xb0, 0x17, 0x8a, 0xf5, 0xe1, 0x7e, 0xb2, 0x86, 0xb3,
	0x16, 0x31, 0x31, 0xbb, 0x87, 0x04, 0x24, 0xb1, 0x81, 0x55, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x2a, 0xad, 0xdc, 0xf6, 0xf7, 0x00, 0x00, 0x00,
}