// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imuachain/oracle/v1/index_recent_params.proto

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

// index for the cached recent params
type IndexRecentParams struct {
	// index list
	Index []uint64 `protobuf:"varint,1,rep,packed,name=index,proto3" json:"index,omitempty"`
}

func (m *IndexRecentParams) Reset()         { *m = IndexRecentParams{} }
func (m *IndexRecentParams) String() string { return proto.CompactTextString(m) }
func (*IndexRecentParams) ProtoMessage()    {}
func (*IndexRecentParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65dead6eda6af21, []int{0}
}
func (m *IndexRecentParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexRecentParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexRecentParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexRecentParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexRecentParams.Merge(m, src)
}
func (m *IndexRecentParams) XXX_Size() int {
	return m.Size()
}
func (m *IndexRecentParams) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexRecentParams.DiscardUnknown(m)
}

var xxx_messageInfo_IndexRecentParams proto.InternalMessageInfo

func (m *IndexRecentParams) GetIndex() []uint64 {
	if m != nil {
		return m.Index
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexRecentParams)(nil), "imuachain.oracle.v1.IndexRecentParams")
}

func init() {
	proto.RegisterFile("imuachain/oracle/v1/index_recent_params.proto", fileDescriptor_f65dead6eda6af21)
}

var fileDescriptor_f65dead6eda6af21 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcd, 0xcc, 0x2d, 0x4d,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f, 0x33, 0xd4,
	0xcf, 0xcc, 0x4b, 0x49, 0xad, 0x88, 0x2f, 0x4a, 0x4d, 0x4e, 0xcd, 0x2b, 0x89, 0x2f, 0x48, 0x2c,
	0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0x2b, 0xd7, 0x83, 0x28,
	0xd7, 0x2b, 0x33, 0x54, 0xd2, 0xe4, 0x12, 0xf4, 0x04, 0xe9, 0x08, 0x02, 0x6b, 0x08, 0x00, 0xab,
	0x17, 0x12, 0xe1, 0x62, 0x05, 0x1b, 0x23, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x12, 0x04, 0xe1, 0x38,
	0xb9, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x4e, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc8, 0x12, 0xdd, 0x8a, 0xca, 0x2a, 0x7d, 0x84,
	0xe3, 0x2a, 0x60, 0xce, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x3b, 0xc7, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x93, 0x3b, 0xce, 0xb1, 0xbf, 0x00, 0x00, 0x00,
}

func (m *IndexRecentParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexRecentParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexRecentParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		dAtA2 := make([]byte, len(m.Index)*10)
		var j1 int
		for _, num := range m.Index {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintIndexRecentParams(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIndexRecentParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovIndexRecentParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IndexRecentParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Index) > 0 {
		l = 0
		for _, e := range m.Index {
			l += sovIndexRecentParams(uint64(e))
		}
		n += 1 + sovIndexRecentParams(uint64(l)) + l
	}
	return n
}

func sovIndexRecentParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIndexRecentParams(x uint64) (n int) {
	return sovIndexRecentParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexRecentParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexRecentParams
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
			return fmt.Errorf("proto: IndexRecentParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexRecentParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIndexRecentParams
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Index = append(m.Index, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIndexRecentParams
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthIndexRecentParams
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthIndexRecentParams
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Index) == 0 {
					m.Index = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowIndexRecentParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Index = append(m.Index, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIndexRecentParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndexRecentParams
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
func skipIndexRecentParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndexRecentParams
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
					return 0, ErrIntOverflowIndexRecentParams
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
					return 0, ErrIntOverflowIndexRecentParams
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
				return 0, ErrInvalidLengthIndexRecentParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIndexRecentParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIndexRecentParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIndexRecentParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndexRecentParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIndexRecentParams = fmt.Errorf("proto: unexpected end of group")
)
