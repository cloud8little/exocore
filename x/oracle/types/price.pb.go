// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/oracle/price.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// token price with timestamp fetched from source
// {price:"12345",decimal:"2"}->price: 123.45 usdt
type PriceTimeDetID struct {
	// price at a specific point(timestamp of non-deterministic source, roundId of deteministic source)
	Price string `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	// decimal of the corresponding price
	Decimal int32 `protobuf:"varint,2,opt,name=decimal,proto3" json:"decimal,omitempty"`
	// timestamp when the price corresponding to
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// det_id is used for deterministic source to tell of which round from this source the price is corresponded
	DetID string `protobuf:"bytes,4,opt,name=det_id,json=detId,proto3" json:"det_id,omitempty"`
}

func (m *PriceTimeDetID) Reset()         { *m = PriceTimeDetID{} }
func (m *PriceTimeDetID) String() string { return proto.CompactTextString(m) }
func (*PriceTimeDetID) ProtoMessage()    {}
func (*PriceTimeDetID) Descriptor() ([]byte, []int) {
	return fileDescriptor_6755466c800b64fc, []int{0}
}
func (m *PriceTimeDetID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriceTimeDetID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceTimeDetID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriceTimeDetID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceTimeDetID.Merge(m, src)
}
func (m *PriceTimeDetID) XXX_Size() int {
	return m.Size()
}
func (m *PriceTimeDetID) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceTimeDetID.DiscardUnknown(m)
}

var xxx_messageInfo_PriceTimeDetID proto.InternalMessageInfo

func (m *PriceTimeDetID) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *PriceTimeDetID) GetDecimal() int32 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *PriceTimeDetID) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *PriceTimeDetID) GetDetID() string {
	if m != nil {
		return m.DetID
	}
	return ""
}

// price with its corresponding source
type PriceSource struct {
	// source_id refers to id from Params.SourceList, where this price fetched from, 0 is reserved for custom usage
	SourceID uint64 `protobuf:"varint,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	// if source is deteministic like chainlink with roundID, set this value with which returned from source
	// up to 3 values in case of the async of network, to give more time for oracle nodes(validators) get into consensus
	// eg.with deterministic source, this array will contian 3 continuous values up to latest
	// for non-deterministic source, it's a choice by v2 rules.
	Prices []*PriceTimeDetID `protobuf:"bytes,2,rep,name=prices,proto3" json:"prices,omitempty"`
	// used for 0-sourceID-customDefinedSource
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (m *PriceSource) Reset()         { *m = PriceSource{} }
func (m *PriceSource) String() string { return proto.CompactTextString(m) }
func (*PriceSource) ProtoMessage()    {}
func (*PriceSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_6755466c800b64fc, []int{1}
}
func (m *PriceSource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriceSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriceSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceSource.Merge(m, src)
}
func (m *PriceSource) XXX_Size() int {
	return m.Size()
}
func (m *PriceSource) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceSource.DiscardUnknown(m)
}

var xxx_messageInfo_PriceSource proto.InternalMessageInfo

func (m *PriceSource) GetSourceID() uint64 {
	if m != nil {
		return m.SourceID
	}
	return 0
}

func (m *PriceSource) GetPrices() []*PriceTimeDetID {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *PriceSource) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

// price with its specified timestamp and roundid(if from deteministic source)
type PriceTimeRound struct {
	// price
	Price string `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	// decimal of the corresponding price
	Decimal int32 `protobuf:"varint,2,opt,name=decimal,proto3" json:"decimal,omitempty"`
	// timestamp when the price is corresponded
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// roundid of the price if the source is deteministic
	RoundID uint64 `protobuf:"varint,4,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
}

func (m *PriceTimeRound) Reset()         { *m = PriceTimeRound{} }
func (m *PriceTimeRound) String() string { return proto.CompactTextString(m) }
func (*PriceTimeRound) ProtoMessage()    {}
func (*PriceTimeRound) Descriptor() ([]byte, []int) {
	return fileDescriptor_6755466c800b64fc, []int{2}
}
func (m *PriceTimeRound) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriceTimeRound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceTimeRound.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriceTimeRound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceTimeRound.Merge(m, src)
}
func (m *PriceTimeRound) XXX_Size() int {
	return m.Size()
}
func (m *PriceTimeRound) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceTimeRound.DiscardUnknown(m)
}

var xxx_messageInfo_PriceTimeRound proto.InternalMessageInfo

func (m *PriceTimeRound) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *PriceTimeRound) GetDecimal() int32 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *PriceTimeRound) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *PriceTimeRound) GetRoundID() uint64 {
	if m != nil {
		return m.RoundID
	}
	return 0
}

func init() {
	proto.RegisterType((*PriceTimeDetID)(nil), "exocore.oracle.PriceTimeDetID")
	proto.RegisterType((*PriceSource)(nil), "exocore.oracle.PriceSource")
	proto.RegisterType((*PriceTimeRound)(nil), "exocore.oracle.PriceTimeRound")
}

func init() { proto.RegisterFile("exocore/oracle/price.proto", fileDescriptor_6755466c800b64fc) }

var fileDescriptor_6755466c800b64fc = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xc7, 0xeb, 0x36, 0x69, 0x1b, 0x17, 0x75, 0xb0, 0x3a, 0x44, 0x15, 0x72, 0xa3, 0x0e, 0xa8,
	0x2c, 0x09, 0x02, 0x89, 0x07, 0x88, 0xc2, 0x10, 0x06, 0x84, 0x0c, 0x13, 0x0b, 0x6a, 0xed, 0x53,
	0x89, 0x68, 0x70, 0xe4, 0xb8, 0xa2, 0x6c, 0x0c, 0x88, 0x99, 0xc7, 0x62, 0xec, 0xc8, 0x54, 0xa1,
	0xf4, 0x45, 0x50, 0x9c, 0x46, 0xd0, 0x9d, 0xed, 0xce, 0xbf, 0xfb, 0xf8, 0xdf, 0xf9, 0xf0, 0x10,
	0x56, 0x92, 0x4b, 0x05, 0x81, 0x54, 0x53, 0xbe, 0x80, 0x20, 0x53, 0x09, 0x07, 0x3f, 0x53, 0x52,
	0x4b, 0xd2, 0xdf, 0x31, 0xbf, 0x62, 0xc3, 0xc1, 0x5c, 0xce, 0xa5, 0x41, 0x41, 0x69, 0x55, 0x51,
	0xe3, 0x57, 0x84, 0xfb, 0xd7, 0x65, 0xd6, 0x6d, 0x92, 0x42, 0x04, 0x3a, 0x8e, 0xc8, 0x00, 0xdb,
	0xa6, 0x8e, 0x8b, 0x3c, 0x34, 0x71, 0x58, 0xe5, 0x10, 0x17, 0x77, 0x04, 0xf0, 0x24, 0x9d, 0x2e,
	0xdc, 0xa6, 0x87, 0x26, 0x36, 0xab, 0x5d, 0x72, 0x88, 0x1d, 0x9d, 0xa4, 0x90, 0xeb, 0x69, 0x9a,
	0xb9, 0x2d, 0x93, 0xf3, 0xfb, 0x40, 0x3c, 0xdc, 0x16, 0xa0, 0xef, 0x13, 0xe1, 0x5a, 0x25, 0x0a,
	0x9d, 0x62, 0x33, 0xb2, 0x4d, 0x23, 0x66, 0x0b, 0xd0, 0xb1, 0x18, 0xbf, 0x21, 0xdc, 0x33, 0x12,
	0x6e, 0xe4, 0x52, 0x71, 0x20, 0xc7, 0xd8, 0xc9, 0x8d, 0x55, 0x26, 0x95, 0x1a, 0xac, 0xf0, 0xa0,
	0xd8, 0x8c, 0xba, 0x15, 0x8e, 0x23, 0xd6, 0xad, 0x70, 0x2c, 0xc8, 0x39, 0x6e, 0x1b, 0x75, 0xb9,
	0xdb, 0xf4, 0x5a, 0x93, 0xde, 0x29, 0xf5, 0xf7, 0x87, 0xf6, 0xf7, 0x47, 0x63, 0xbb, 0x68, 0x42,
	0xb0, 0x25, 0x20, 0xe7, 0x3b, 0xb5, 0xc6, 0x1e, 0xbf, 0xff, 0xdd, 0x04, 0x93, 0xcb, 0x27, 0xf1,
	0xcf, 0x9b, 0x38, 0xc2, 0x5d, 0x55, 0x96, 0xad, 0x77, 0x61, 0x85, 0xbd, 0x62, 0x33, 0xea, 0x98,
	0x56, 0x71, 0xc4, 0x3a, 0x06, 0xc6, 0x22, 0xbc, 0xfc, 0x2c, 0x28, 0x5a, 0x17, 0x14, 0x7d, 0x17,
	0x14, 0x7d, 0x6c, 0x69, 0x63, 0xbd, 0xa5, 0x8d, 0xaf, 0x2d, 0x6d, 0xdc, 0x9d, 0xcc, 0x13, 0xfd,
	0xb0, 0x9c, 0xf9, 0x5c, 0xa6, 0xc1, 0x45, 0x35, 0xe8, 0x15, 0xe8, 0x67, 0xa9, 0x1e, 0x83, 0xfa,
	0x10, 0x56, 0xf5, 0x29, 0xe8, 0x97, 0x0c, 0xf2, 0x59, 0xdb, 0xfc, 0xf2, 0xd9, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc3, 0xe6, 0x21, 0x72, 0x29, 0x02, 0x00, 0x00,
}

func (m *PriceTimeDetID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceTimeDetID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceTimeDetID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DetID) > 0 {
		i -= len(m.DetID)
		copy(dAtA[i:], m.DetID)
		i = encodeVarintPrice(dAtA, i, uint64(len(m.DetID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Timestamp) > 0 {
		i -= len(m.Timestamp)
		copy(dAtA[i:], m.Timestamp)
		i = encodeVarintPrice(dAtA, i, uint64(len(m.Timestamp)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Decimal != 0 {
		i = encodeVarintPrice(dAtA, i, uint64(m.Decimal))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintPrice(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PriceSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceSource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceSource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintPrice(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Prices) > 0 {
		for iNdEx := len(m.Prices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPrice(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.SourceID != 0 {
		i = encodeVarintPrice(dAtA, i, uint64(m.SourceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PriceTimeRound) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceTimeRound) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceTimeRound) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RoundID != 0 {
		i = encodeVarintPrice(dAtA, i, uint64(m.RoundID))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Timestamp) > 0 {
		i -= len(m.Timestamp)
		copy(dAtA[i:], m.Timestamp)
		i = encodeVarintPrice(dAtA, i, uint64(len(m.Timestamp)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Decimal != 0 {
		i = encodeVarintPrice(dAtA, i, uint64(m.Decimal))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintPrice(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPrice(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PriceTimeDetID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovPrice(uint64(l))
	}
	if m.Decimal != 0 {
		n += 1 + sovPrice(uint64(m.Decimal))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovPrice(uint64(l))
	}
	l = len(m.DetID)
	if l > 0 {
		n += 1 + l + sovPrice(uint64(l))
	}
	return n
}

func (m *PriceSource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceID != 0 {
		n += 1 + sovPrice(uint64(m.SourceID))
	}
	if len(m.Prices) > 0 {
		for _, e := range m.Prices {
			l = e.Size()
			n += 1 + l + sovPrice(uint64(l))
		}
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovPrice(uint64(l))
	}
	return n
}

func (m *PriceTimeRound) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovPrice(uint64(l))
	}
	if m.Decimal != 0 {
		n += 1 + sovPrice(uint64(m.Decimal))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovPrice(uint64(l))
	}
	if m.RoundID != 0 {
		n += 1 + sovPrice(uint64(m.RoundID))
	}
	return n
}

func sovPrice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrice(x uint64) (n int) {
	return sovPrice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PriceTimeDetID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrice
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
			return fmt.Errorf("proto: PriceTimeDetID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceTimeDetID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimal |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DetID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DetID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrice
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
func (m *PriceSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrice
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
			return fmt.Errorf("proto: PriceSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceID", wireType)
			}
			m.SourceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SourceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prices = append(m.Prices, &PriceTimeDetID{})
			if err := m.Prices[len(m.Prices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrice
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
func (m *PriceTimeRound) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrice
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
			return fmt.Errorf("proto: PriceTimeRound: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceTimeRound: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimal |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoundID", wireType)
			}
			m.RoundID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoundID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrice
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
func skipPrice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrice
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
					return 0, ErrIntOverflowPrice
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
					return 0, ErrIntOverflowPrice
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
				return 0, ErrInvalidLengthPrice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrice = fmt.Errorf("proto: unexpected end of group")
)
