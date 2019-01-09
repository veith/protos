// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: date.proto

package date // import "github.com/veith/protos/date"

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

// Represents a whole calendar date, e.g. date of birth. The time of day and
// time zone are either specified elsewhere or are not significant. The date
// is relative to the Proleptic Gregorian Calendar. The day may be 0 to
// represent a year and month where the day is not significant, e.g. credit card
// expiration date. The year may be 0 to represent a month and day independent
// of year, e.g. anniversary date. Related types are [google.type.TimeOfDay][google.type.TimeOfDay]
// and `google.protobuf.Timestamp`.
type Date struct {
	// Year of date. Must be from 1 to 9999, or 0 if specifying a date without
	// a year.
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Month of year. Must be from 1 to 12.
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// Day of month. Must be from 1 to 31 and valid for the year and month, or 0
	// if specifying a year/month where the day is not significant.
	Day                  int32    `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_date_da3470df9312eea8, []int{0}
}
func (m *Date) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Date.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(dst, src)
}
func (m *Date) XXX_Size() int {
	return m.Size()
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func init() {
	proto.RegisterType((*Date)(nil), "furo.type.Date")
}
func (m *Date) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Date) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Year != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDate(dAtA, i, uint64(m.Year))
	}
	if m.Month != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDate(dAtA, i, uint64(m.Month))
	}
	if m.Day != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintDate(dAtA, i, uint64(m.Day))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDate(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Date) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Year != 0 {
		n += 1 + sovDate(uint64(m.Year))
	}
	if m.Month != 0 {
		n += 1 + sovDate(uint64(m.Month))
	}
	if m.Day != 0 {
		n += 1 + sovDate(uint64(m.Day))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDate(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDate(x uint64) (n int) {
	return sovDate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Date) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDate
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
			return fmt.Errorf("proto: Date: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Date: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Year", wireType)
			}
			m.Year = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Year |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Month", wireType)
			}
			m.Month = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Month |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Day", wireType)
			}
			m.Day = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Day |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDate
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
func skipDate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDate
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
					return 0, ErrIntOverflowDate
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
					return 0, ErrIntOverflowDate
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
				return 0, ErrInvalidLengthDate
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDate
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
				next, err := skipDate(dAtA[start:])
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
	ErrInvalidLengthDate = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDate   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("date.proto", fileDescriptor_date_da3470df9312eea8) }

var fileDescriptor_date_da3470df9312eea8 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x2b, 0x2d, 0xca, 0xd7, 0x2b, 0xa9, 0x2c,
	0x48, 0x55, 0x72, 0xe2, 0x62, 0x71, 0x49, 0x2c, 0x49, 0x15, 0x12, 0xe2, 0x62, 0xa9, 0x4c, 0x4d,
	0x2c, 0x92, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0x73, 0xf3,
	0xf3, 0x4a, 0x32, 0x24, 0x98, 0xc0, 0x82, 0x10, 0x8e, 0x90, 0x00, 0x17, 0x73, 0x4a, 0x62, 0xa5,
	0x04, 0x33, 0x58, 0x0c, 0xc4, 0x74, 0x0a, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0xb9, 0xf8, 0x93, 0xf3, 0x73, 0xf5, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52,
	0xc1, 0x56, 0x38, 0x71, 0x82, 0x2c, 0x08, 0x00, 0x59, 0x1c, 0xc0, 0x18, 0xa5, 0x98, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x96, 0x9a, 0x59, 0x92, 0xa1, 0x0f, 0x76,
	0x52, 0xb1, 0x3e, 0xc8, 0x79, 0xd6, 0x20, 0xe2, 0x07, 0x23, 0xe3, 0x22, 0x26, 0x66, 0xf7, 0x90,
	0x80, 0x24, 0x36, 0xb0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x41, 0xde, 0xe3, 0xbf, 0xbb,
	0x00, 0x00, 0x00,
}
