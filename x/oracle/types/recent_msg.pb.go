// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imuachain/oracle/v1/recent_msg.proto

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

// RecentMsg represent the messages to be cached for recent blocks
type RecentMsg struct {
	// block height these messages from
	Block uint64 `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	// cached messages
	Msgs []*MsgItem `protobuf:"bytes,2,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (m *RecentMsg) Reset()         { *m = RecentMsg{} }
func (m *RecentMsg) String() string { return proto.CompactTextString(m) }
func (*RecentMsg) ProtoMessage()    {}
func (*RecentMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_49881b4e42685ec0, []int{0}
}
func (m *RecentMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecentMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecentMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecentMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecentMsg.Merge(m, src)
}
func (m *RecentMsg) XXX_Size() int {
	return m.Size()
}
func (m *RecentMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RecentMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RecentMsg proto.InternalMessageInfo

func (m *RecentMsg) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *RecentMsg) GetMsgs() []*MsgItem {
	if m != nil {
		return m.Msgs
	}
	return nil
}

// MsgItem represents the message info of createPrice
type MsgItem struct {
	// feeder_id tells of wich feeder this price if corresponding to
	FeederID uint64 `protobuf:"varint,2,opt,name=feeder_id,json=feederId,proto3" json:"feeder_id,omitempty"`
	// p_source price with its source info
	PSources []*PriceSource `protobuf:"bytes,3,rep,name=p_sources,json=pSources,proto3" json:"p_sources,omitempty"`
	// validator tells which validator create this price
	Validator string `protobuf:"bytes,4,opt,name=validator,proto3" json:"validator,omitempty"`
}

func (m *MsgItem) Reset()         { *m = MsgItem{} }
func (m *MsgItem) String() string { return proto.CompactTextString(m) }
func (*MsgItem) ProtoMessage()    {}
func (*MsgItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_49881b4e42685ec0, []int{1}
}
func (m *MsgItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgItem.Merge(m, src)
}
func (m *MsgItem) XXX_Size() int {
	return m.Size()
}
func (m *MsgItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgItem.DiscardUnknown(m)
}

var xxx_messageInfo_MsgItem proto.InternalMessageInfo

func (m *MsgItem) GetFeederID() uint64 {
	if m != nil {
		return m.FeederID
	}
	return 0
}

func (m *MsgItem) GetPSources() []*PriceSource {
	if m != nil {
		return m.PSources
	}
	return nil
}

func (m *MsgItem) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func init() {
	proto.RegisterType((*RecentMsg)(nil), "imuachain.oracle.v1.RecentMsg")
	proto.RegisterType((*MsgItem)(nil), "imuachain.oracle.v1.MsgItem")
}

func init() {
	proto.RegisterFile("imuachain/oracle/v1/recent_msg.proto", fileDescriptor_49881b4e42685ec0)
}

var fileDescriptor_49881b4e42685ec0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4f, 0x02, 0x31,
	0x1c, 0xc5, 0x29, 0xa0, 0x72, 0xd5, 0xe9, 0x64, 0xb8, 0x10, 0x72, 0x5c, 0x88, 0x03, 0x26, 0xda,
	0x8a, 0xce, 0x2e, 0xc4, 0x90, 0x30, 0x90, 0x98, 0xb2, 0xb9, 0x90, 0xa3, 0x57, 0x4b, 0x23, 0x47,
	0x2f, 0x6d, 0x21, 0xe0, 0xa7, 0xd0, 0x6f, 0xe5, 0xc8, 0xe8, 0x64, 0xcc, 0xf1, 0x45, 0xcc, 0xb5,
	0x2a, 0xcb, 0x6d, 0xff, 0xbe, 0xf7, 0xcb, 0x7b, 0xe9, 0x83, 0x17, 0x22, 0x5d, 0xc5, 0x74, 0x1e,
	0x8b, 0x25, 0x96, 0x2a, 0xa6, 0x0b, 0x86, 0xd7, 0x7d, 0xac, 0x18, 0x65, 0x4b, 0x33, 0x4d, 0x35,
	0x47, 0x99, 0x92, 0x46, 0xfa, 0xe7, 0xff, 0x14, 0x72, 0x14, 0x5a, 0xf7, 0x5b, 0x4d, 0x2e, 0xb9,
	0xb4, 0x3e, 0x2e, 0x2e, 0x87, 0xb6, 0x3a, 0x65, 0x81, 0x99, 0x12, 0x94, 0x39, 0xa0, 0x3b, 0x81,
	0x1e, 0xb1, 0xf9, 0x63, 0xcd, 0xfd, 0x26, 0x3c, 0x9a, 0x2d, 0x24, 0x7d, 0x09, 0x40, 0x04, 0x7a,
	0x75, 0xe2, 0x1e, 0xfe, 0x0d, 0xac, 0xa7, 0x9a, 0xeb, 0xa0, 0x1a, 0xd5, 0x7a, 0xa7, 0xb7, 0x6d,
	0x54, 0xd2, 0x8e, 0xc6, 0x9a, 0x8f, 0x0c, 0x4b, 0x89, 0x25, 0xbb, 0xef, 0x00, 0x9e, 0xfc, 0x2a,
	0xfe, 0x25, 0xf4, 0x9e, 0x19, 0x4b, 0x98, 0x9a, 0x8a, 0x24, 0xa8, 0x16, 0xb9, 0x83, 0xb3, 0xfc,
	0xab, 0xd3, 0x18, 0x5a, 0x71, 0xf4, 0x40, 0x1a, 0xce, 0x1e, 0x25, 0xfe, 0x3d, 0xf4, 0xb2, 0xa9,
	0x96, 0x2b, 0x45, 0x99, 0x0e, 0x6a, 0xb6, 0x2d, 0x2a, 0x6d, 0x7b, 0x2c, 0x3e, 0x30, 0xb1, 0x20,
	0x69, 0x64, 0xee, 0xd0, 0x7e, 0x1b, 0x7a, 0xeb, 0x78, 0x21, 0x92, 0xd8, 0x48, 0x15, 0xd4, 0x23,
	0xd0, 0xf3, 0xc8, 0x41, 0x18, 0x0c, 0x3f, 0xf2, 0x10, 0xec, 0xf2, 0x10, 0x7c, 0xe7, 0x21, 0x78,
	0xdb, 0x87, 0x95, 0xdd, 0x3e, 0xac, 0x7c, 0xee, 0xc3, 0xca, 0xd3, 0x15, 0x17, 0x66, 0xbe, 0x9a,
	0x21, 0x2a, 0x53, 0x5c, 0xb4, 0x5d, 0x6f, 0xb6, 0xaf, 0xf8, 0xb0, 0xdb, 0xe6, 0x6f, 0x39, 0xb3,
	0xcd, 0x98, 0x9e, 0x1d, 0xdb, 0xdd, 0xee, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x39, 0x54, 0x69,
	0x16, 0xab, 0x01, 0x00, 0x00,
}

func (m *RecentMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecentMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecentMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRecentMsg(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Block != 0 {
		i = encodeVarintRecentMsg(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintRecentMsg(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PSources) > 0 {
		for iNdEx := len(m.PSources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PSources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRecentMsg(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.FeederID != 0 {
		i = encodeVarintRecentMsg(dAtA, i, uint64(m.FeederID))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func encodeVarintRecentMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovRecentMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RecentMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != 0 {
		n += 1 + sovRecentMsg(uint64(m.Block))
	}
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovRecentMsg(uint64(l))
		}
	}
	return n
}

func (m *MsgItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FeederID != 0 {
		n += 1 + sovRecentMsg(uint64(m.FeederID))
	}
	if len(m.PSources) > 0 {
		for _, e := range m.PSources {
			l = e.Size()
			n += 1 + l + sovRecentMsg(uint64(l))
		}
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovRecentMsg(uint64(l))
	}
	return n
}

func sovRecentMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRecentMsg(x uint64) (n int) {
	return sovRecentMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RecentMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecentMsg
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
			return fmt.Errorf("proto: RecentMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecentMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecentMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecentMsg
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
				return ErrInvalidLengthRecentMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecentMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, &MsgItem{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecentMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecentMsg
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
func (m *MsgItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecentMsg
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
			return fmt.Errorf("proto: MsgItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeederID", wireType)
			}
			m.FeederID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecentMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeederID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PSources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecentMsg
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
				return ErrInvalidLengthRecentMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecentMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PSources = append(m.PSources, &PriceSource{})
			if err := m.PSources[len(m.PSources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecentMsg
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
				return ErrInvalidLengthRecentMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecentMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecentMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecentMsg
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
func skipRecentMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecentMsg
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
					return 0, ErrIntOverflowRecentMsg
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
					return 0, ErrIntOverflowRecentMsg
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
				return 0, ErrInvalidLengthRecentMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRecentMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRecentMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRecentMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecentMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRecentMsg = fmt.Errorf("proto: unexpected end of group")
)
