// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: color.proto

package color // import "github.com/veith/protos/color"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gogo/protobuf/types"

import encoding_binary "encoding/binary"

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

// Represents a color in the RGBA color space. This representation is designed
// for simplicity of conversion to/from color representations in various
// languages over compactness; for example, the fields of this representation
// can be trivially provided to the constructor of "java.awt.Color" in Java; it
// can also be trivially provided to UIColor's "+colorWithRed:green:blue:alpha"
// method in iOS; and, with just a little work, it can be easily formatted into
// a CSS "rgba()" string in JavaScript, as well. Here are some examples:
//
// Example (Java):
//
//      import com.google.type.Color;
//
//      // ...
//      public static java.awt.Color fromProto(Color protocolor) {
//        float alpha = protocolor.hasAlpha()
//            ? protocolor.getAlpha().getValue()
//            : 1.0;
//
//        return new java.awt.Color(
//            protocolor.getRed(),
//            protocolor.getGreen(),
//            protocolor.getBlue(),
//            alpha);
//      }
//
//      public static Color toProto(java.awt.Color color) {
//        float red = (float) color.getRed();
//        float green = (float) color.getGreen();
//        float blue = (float) color.getBlue();
//        float denominator = 255.0;
//        Color.Builder resultBuilder =
//            Color
//                .newBuilder()
//                .setRed(red / denominator)
//                .setGreen(green / denominator)
//                .setBlue(blue / denominator);
//        int alpha = color.getAlpha();
//        if (alpha != 255) {
//          result.setAlpha(
//              FloatValue
//                  .newBuilder()
//                  .setValue(((float) alpha) / denominator)
//                  .build());
//        }
//        return resultBuilder.build();
//      }
//      // ...
//
// Example (iOS / Obj-C):
//
//      // ...
//      static UIColor* fromProto(Color* protocolor) {
//         float red = [protocolor red];
//         float green = [protocolor green];
//         float blue = [protocolor blue];
//         FloatValue* alpha_wrapper = [protocolor alpha];
//         float alpha = 1.0;
//         if (alpha_wrapper != nil) {
//           alpha = [alpha_wrapper value];
//         }
//         return [UIColor colorWithRed:red green:green blue:blue alpha:alpha];
//      }
//
//      static Color* toProto(UIColor* color) {
//          CGFloat red, green, blue, alpha;
//          if (![color getRed:&red green:&green blue:&blue alpha:&alpha]) {
//            return nil;
//          }
//          Color* result = [Color alloc] init];
//          [result setRed:red];
//          [result setGreen:green];
//          [result setBlue:blue];
//          if (alpha <= 0.9999) {
//            [result setAlpha:floatWrapperWithValue(alpha)];
//          }
//          [result autorelease];
//          return result;
//     }
//     // ...
//
//  Example (JavaScript):
//
//     // ...
//
//     var protoToCssColor = function(rgb_color) {
//        var redFrac = rgb_color.red || 0.0;
//        var greenFrac = rgb_color.green || 0.0;
//        var blueFrac = rgb_color.blue || 0.0;
//        var red = Math.floor(redFrac * 255);
//        var green = Math.floor(greenFrac * 255);
//        var blue = Math.floor(blueFrac * 255);
//
//        if (!('alpha' in rgb_color)) {
//           return rgbToCssColor_(red, green, blue);
//        }
//
//        var alphaFrac = rgb_color.alpha.value || 0.0;
//        var rgbParams = [red, green, blue].join(',');
//        return ['rgba(', rgbParams, ',', alphaFrac, ')'].join('');
//     };
//
//     var rgbToCssColor_ = function(red, green, blue) {
//       var rgbNumber = new Number((red << 16) | (green << 8) | blue);
//       var hexString = rgbNumber.toString(16);
//       var missingZeros = 6 - hexString.length;
//       var resultBuilder = ['#'];
//       for (var i = 0; i < missingZeros; i++) {
//          resultBuilder.push('0');
//       }
//       resultBuilder.push(hexString);
//       return resultBuilder.join('');
//     };
//
//     // ...
type Color struct {
	// The amount of red in the color as a value in the interval [0, 1].
	Red float32 `protobuf:"fixed32,1,opt,name=red,proto3" json:"red,omitempty"`
	// The amount of green in the color as a value in the interval [0, 1].
	Green float32 `protobuf:"fixed32,2,opt,name=green,proto3" json:"green,omitempty"`
	// The amount of blue in the color as a value in the interval [0, 1].
	Blue float32 `protobuf:"fixed32,3,opt,name=blue,proto3" json:"blue,omitempty"`
	// The fraction of this color that should be applied to the pixel. That is,
	// the final pixel color is defined by the equation:
	//
	//   pixel color = alpha * (this color) + (1.0 - alpha) * (background color)
	//
	// This means that a value of 1.0 corresponds to a solid color, whereas
	// a value of 0.0 corresponds to a completely transparent color. This
	// uses a wrapper message rather than a simple float scalar so that it is
	// possible to distinguish between a default value and the value being unset.
	// If omitted, this color object is to be rendered as a solid color
	// (as if the alpha value had been explicitly given with a value of 1.0).
	Alpha                *types.FloatValue `protobuf:"bytes,4,opt,name=alpha" json:"alpha,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Color) Reset()         { *m = Color{} }
func (m *Color) String() string { return proto.CompactTextString(m) }
func (*Color) ProtoMessage()    {}
func (*Color) Descriptor() ([]byte, []int) {
	return fileDescriptor_color_c69ffd7e5c6a3dd8, []int{0}
}
func (m *Color) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Color) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Color.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Color) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Color.Merge(dst, src)
}
func (m *Color) XXX_Size() int {
	return m.Size()
}
func (m *Color) XXX_DiscardUnknown() {
	xxx_messageInfo_Color.DiscardUnknown(m)
}

var xxx_messageInfo_Color proto.InternalMessageInfo

func (m *Color) GetRed() float32 {
	if m != nil {
		return m.Red
	}
	return 0
}

func (m *Color) GetGreen() float32 {
	if m != nil {
		return m.Green
	}
	return 0
}

func (m *Color) GetBlue() float32 {
	if m != nil {
		return m.Blue
	}
	return 0
}

func (m *Color) GetAlpha() *types.FloatValue {
	if m != nil {
		return m.Alpha
	}
	return nil
}

func init() {
	proto.RegisterType((*Color)(nil), "google.type.Color")
}
func (m *Color) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Color) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Red != 0 {
		dAtA[i] = 0xd
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Red))))
		i += 4
	}
	if m.Green != 0 {
		dAtA[i] = 0x15
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Green))))
		i += 4
	}
	if m.Blue != 0 {
		dAtA[i] = 0x1d
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Blue))))
		i += 4
	}
	if m.Alpha != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintColor(dAtA, i, uint64(m.Alpha.Size()))
		n1, err := m.Alpha.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintColor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Color) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Red != 0 {
		n += 5
	}
	if m.Green != 0 {
		n += 5
	}
	if m.Blue != 0 {
		n += 5
	}
	if m.Alpha != nil {
		l = m.Alpha.Size()
		n += 1 + l + sovColor(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovColor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozColor(x uint64) (n int) {
	return sovColor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Color) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowColor
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
			return fmt.Errorf("proto: Color: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Color: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Red", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Red = float32(math.Float32frombits(v))
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Green", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Green = float32(math.Float32frombits(v))
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blue", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Blue = float32(math.Float32frombits(v))
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alpha", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowColor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthColor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Alpha == nil {
				m.Alpha = &types.FloatValue{}
			}
			if err := m.Alpha.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipColor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthColor
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
func skipColor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowColor
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
					return 0, ErrIntOverflowColor
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
					return 0, ErrIntOverflowColor
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
				return 0, ErrInvalidLengthColor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowColor
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
				next, err := skipColor(dAtA[start:])
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
	ErrInvalidLengthColor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowColor   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("color.proto", fileDescriptor_color_c69ffd7e5c6a3dd8) }

var fileDescriptor_color_c69ffd7e5c6a3dd8 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xce, 0xcf, 0xc9,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5,
	0x2b, 0xa9, 0x2c, 0x48, 0x95, 0x92, 0x83, 0x70, 0xf4, 0xc1, 0x52, 0x49, 0xa5, 0x69, 0xfa, 0xe5,
	0x45, 0x89, 0x05, 0x05, 0xa9, 0x45, 0xc5, 0x10, 0xc5, 0x4a, 0x65, 0x5c, 0xac, 0xce, 0x20, 0xbd,
	0x42, 0x02, 0x5c, 0xcc, 0x45, 0xa9, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x4c, 0x41, 0x20, 0xa6,
	0x90, 0x08, 0x17, 0x6b, 0x7a, 0x51, 0x6a, 0x6a, 0x9e, 0x04, 0x13, 0x58, 0x0c, 0xc2, 0x11, 0x12,
	0xe2, 0x62, 0x49, 0xca, 0x29, 0x4d, 0x95, 0x60, 0x06, 0x0b, 0x82, 0xd9, 0x42, 0x86, 0x5c, 0xac,
	0x89, 0x39, 0x05, 0x19, 0x89, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xd2, 0x7a, 0x50, 0x17,
	0xc0, 0x2c, 0xd5, 0x73, 0xcb, 0xc9, 0x4f, 0x2c, 0x09, 0x4b, 0xcc, 0x29, 0x4d, 0x0d, 0x82, 0xa8,
	0x74, 0x0a, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0xb9,
	0xf8, 0x93, 0xf3, 0x73, 0xf5, 0x90, 0x9c, 0xed, 0xc4, 0x05, 0x76, 0x54, 0x00, 0x48, 0x7f, 0x00,
	0x63, 0x94, 0x72, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x59, 0x6a,
	0x66, 0x49, 0x06, 0xc4, 0x3b, 0xc5, 0xfa, 0x60, 0x6f, 0x5b, 0x83, 0xc9, 0x45, 0x4c, 0xcc, 0xee,
	0x21, 0x01, 0x49, 0x6c, 0x60, 0x19, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xe2, 0x1e,
	0x30, 0x12, 0x01, 0x00, 0x00,
}