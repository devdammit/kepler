// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kepler/kepler/w_3_func.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type W3Func struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Cond       string `protobuf:"bytes,2,opt,name=cond,proto3" json:"cond,omitempty"`
	Code       string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Counter    uint64 `protobuf:"varint,4,opt,name=counter,proto3" json:"counter,omitempty"`
	ExecutedAt string `protobuf:"bytes,5,opt,name=executedAt,proto3" json:"executedAt,omitempty"`
	Id         uint64 `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *W3Func) Reset()         { *m = W3Func{} }
func (m *W3Func) String() string { return proto.CompactTextString(m) }
func (*W3Func) ProtoMessage()    {}
func (*W3Func) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2816027511e78a1, []int{0}
}
func (m *W3Func) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *W3Func) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_W3Func.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *W3Func) XXX_Merge(src proto.Message) {
	xxx_messageInfo_W3Func.Merge(m, src)
}
func (m *W3Func) XXX_Size() int {
	return m.Size()
}
func (m *W3Func) XXX_DiscardUnknown() {
	xxx_messageInfo_W3Func.DiscardUnknown(m)
}

var xxx_messageInfo_W3Func proto.InternalMessageInfo

func (m *W3Func) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *W3Func) GetCond() string {
	if m != nil {
		return m.Cond
	}
	return ""
}

func (m *W3Func) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *W3Func) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *W3Func) GetExecutedAt() string {
	if m != nil {
		return m.ExecutedAt
	}
	return ""
}

func (m *W3Func) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*W3Func)(nil), "kepler.kepler.W3Func")
}

func init() { proto.RegisterFile("kepler/kepler/w_3_func.proto", fileDescriptor_c2816027511e78a1) }

var fileDescriptor_c2816027511e78a1 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0x4e, 0x2d, 0xc8,
	0x49, 0x2d, 0xd2, 0x87, 0x52, 0xe5, 0xf1, 0xc6, 0xf1, 0x69, 0xa5, 0x79, 0xc9, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x10, 0x61, 0x3d, 0x08, 0xa5, 0x34, 0x85, 0x91, 0x8b, 0x2d, 0xdc,
	0xd8, 0xad, 0x34, 0x2f, 0x59, 0x48, 0x82, 0x8b, 0x3d, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0xbf, 0x48,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x12, 0xe2, 0x62, 0x49, 0xce, 0xcf, 0x4b,
	0x91, 0x60, 0x02, 0x0b, 0x83, 0xd9, 0x10, 0xb1, 0x94, 0x54, 0x09, 0x66, 0x98, 0x58, 0x4a, 0x2a,
	0xd8, 0x84, 0xfc, 0xd2, 0xbc, 0x92, 0xd4, 0x22, 0x09, 0x16, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x18,
	0x57, 0x48, 0x8e, 0x8b, 0x2b, 0xb5, 0x22, 0x35, 0xb9, 0xb4, 0x24, 0x35, 0xc5, 0xb1, 0x44, 0x82,
	0x15, 0xac, 0x07, 0x49, 0x44, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x0d, 0xac, 0x89, 0x29,
	0x33, 0xc5, 0xc9, 0xe5, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0xd2,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x53, 0x52, 0xcb, 0x52, 0x12, 0x73,
	0x73, 0x33, 0x4b, 0x60, 0x7e, 0xad, 0x80, 0x31, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0,
	0x5e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x34, 0x79, 0xbf, 0x9b, 0x12, 0x01, 0x00, 0x00,
}

func (m *W3Func) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *W3Func) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *W3Func) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintW_3Func(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ExecutedAt) > 0 {
		i -= len(m.ExecutedAt)
		copy(dAtA[i:], m.ExecutedAt)
		i = encodeVarintW_3Func(dAtA, i, uint64(len(m.ExecutedAt)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Counter != 0 {
		i = encodeVarintW_3Func(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintW_3Func(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Cond) > 0 {
		i -= len(m.Cond)
		copy(dAtA[i:], m.Cond)
		i = encodeVarintW_3Func(dAtA, i, uint64(len(m.Cond)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintW_3Func(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintW_3Func(dAtA []byte, offset int, v uint64) int {
	offset -= sovW_3Func(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *W3Func) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovW_3Func(uint64(l))
	}
	l = len(m.Cond)
	if l > 0 {
		n += 1 + l + sovW_3Func(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovW_3Func(uint64(l))
	}
	if m.Counter != 0 {
		n += 1 + sovW_3Func(uint64(m.Counter))
	}
	l = len(m.ExecutedAt)
	if l > 0 {
		n += 1 + l + sovW_3Func(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovW_3Func(uint64(m.Id))
	}
	return n
}

func sovW_3Func(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozW_3Func(x uint64) (n int) {
	return sovW_3Func(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *W3Func) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowW_3Func
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: W3Func: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: W3Func: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowW_3Func
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthW_3Func
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthW_3Func
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cond", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowW_3Func
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthW_3Func
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthW_3Func
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cond = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowW_3Func
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthW_3Func
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthW_3Func
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowW_3Func
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowW_3Func
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthW_3Func
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthW_3Func
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowW_3Func
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipW_3Func(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthW_3Func
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipW_3Func(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowW_3Func
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
					return 0, ErrIntOverflowW_3Func
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowW_3Func
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
			if length < 0 {
				return 0, ErrInvalidLengthW_3Func
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupW_3Func
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthW_3Func
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthW_3Func        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowW_3Func          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupW_3Func = fmt.Errorf("proto: unexpected end of group")
)
