// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/delegation/v1/genesis.proto

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

// GenesisState defines the delegation module's state. It needs to encompass
// all of the state that is required to start the chain from the genesis
// or in the event of a restart. At this point, it is only built with
// the former in mind. There are no params in this module.
type GenesisState struct {
	// associations represents the association between a staker and an operator.
	Associations []StakerToOperator `protobuf:"bytes,1,rep,name=associations,proto3" json:"associations"`
	// delegation_states is a list of all delegation states.
	DelegationStates []DelegationStates `protobuf:"bytes,2,rep,name=delegation_states,json=delegationStates,proto3" json:"delegation_states"`
	// stakers_by_operator is a staker list for the operators
	StakersByOperator []StakersByOperator `protobuf:"bytes,3,rep,name=stakers_by_operator,json=stakersByOperator,proto3" json:"stakers_by_operator"`
	// undelegations is a list of all undelegations
	Undelegations []UndelegationAndHoldCount `protobuf:"bytes,4,rep,name=undelegations,proto3" json:"undelegations"`
	// last_undelegation_id is a globally incremented ID assigned to undelegations.
	LastUndelegationId uint64 `protobuf:"varint,5,opt,name=last_undelegation_id,json=lastUndelegationId,proto3" json:"last_undelegation_id,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26dd0d733927603, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetAssociations() []StakerToOperator {
	if m != nil {
		return m.Associations
	}
	return nil
}

func (m *GenesisState) GetDelegationStates() []DelegationStates {
	if m != nil {
		return m.DelegationStates
	}
	return nil
}

func (m *GenesisState) GetStakersByOperator() []StakersByOperator {
	if m != nil {
		return m.StakersByOperator
	}
	return nil
}

func (m *GenesisState) GetUndelegations() []UndelegationAndHoldCount {
	if m != nil {
		return m.Undelegations
	}
	return nil
}

func (m *GenesisState) GetLastUndelegationId() uint64 {
	if m != nil {
		return m.LastUndelegationId
	}
	return 0
}

// DelegationStates is a helper struct for the delegation state
// used to construct the genesis state
type DelegationStates struct {
	// key is used for storing the delegation states,
	// which is a combination of the staker ID, asset ID, and operator address.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// states is the value of undelegation state for the above key
	States DelegationAmounts `protobuf:"bytes,2,opt,name=states,proto3" json:"states"`
}

func (m *DelegationStates) Reset()         { *m = DelegationStates{} }
func (m *DelegationStates) String() string { return proto.CompactTextString(m) }
func (*DelegationStates) ProtoMessage()    {}
func (*DelegationStates) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26dd0d733927603, []int{1}
}
func (m *DelegationStates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegationStates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegationStates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegationStates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegationStates.Merge(m, src)
}
func (m *DelegationStates) XXX_Size() int {
	return m.Size()
}
func (m *DelegationStates) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegationStates.DiscardUnknown(m)
}

var xxx_messageInfo_DelegationStates proto.InternalMessageInfo

func (m *DelegationStates) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DelegationStates) GetStates() DelegationAmounts {
	if m != nil {
		return m.States
	}
	return DelegationAmounts{}
}

// StakersByOperator is a helper struct for an operator's staker list
// used to construct the genesis state
type StakersByOperator struct {
	// key is used for storing the staker list,
	// which is a combination of the operator address and the asset ID.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// stakers is the stakers list for the above key
	Stakers []string `protobuf:"bytes,2,rep,name=stakers,proto3" json:"stakers,omitempty"`
}

func (m *StakersByOperator) Reset()         { *m = StakersByOperator{} }
func (m *StakersByOperator) String() string { return proto.CompactTextString(m) }
func (*StakersByOperator) ProtoMessage()    {}
func (*StakersByOperator) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26dd0d733927603, []int{2}
}
func (m *StakersByOperator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakersByOperator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakersByOperator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakersByOperator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakersByOperator.Merge(m, src)
}
func (m *StakersByOperator) XXX_Size() int {
	return m.Size()
}
func (m *StakersByOperator) XXX_DiscardUnknown() {
	xxx_messageInfo_StakersByOperator.DiscardUnknown(m)
}

var xxx_messageInfo_StakersByOperator proto.InternalMessageInfo

func (m *StakersByOperator) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StakersByOperator) GetStakers() []string {
	if m != nil {
		return m.Stakers
	}
	return nil
}

// DelegationsByStaker is a list of delegations for a single staker.
type DelegationsByStaker struct {
	// staker_id is the staker's account address + _ + l0 chain id (hex).``
	StakerId string `protobuf:"bytes,1,opt,name=staker_id,json=stakerId,proto3" json:"staker_id,omitempty"`
	// delegations is the list of delegations for the staker, indexed by the
	// asset_id.
	Delegations []DelegatedSingleAssetInfo `protobuf:"bytes,2,rep,name=delegations,proto3" json:"delegations"`
}

func (m *DelegationsByStaker) Reset()         { *m = DelegationsByStaker{} }
func (m *DelegationsByStaker) String() string { return proto.CompactTextString(m) }
func (*DelegationsByStaker) ProtoMessage()    {}
func (*DelegationsByStaker) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26dd0d733927603, []int{3}
}
func (m *DelegationsByStaker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegationsByStaker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegationsByStaker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegationsByStaker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegationsByStaker.Merge(m, src)
}
func (m *DelegationsByStaker) XXX_Size() int {
	return m.Size()
}
func (m *DelegationsByStaker) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegationsByStaker.DiscardUnknown(m)
}

var xxx_messageInfo_DelegationsByStaker proto.InternalMessageInfo

func (m *DelegationsByStaker) GetStakerId() string {
	if m != nil {
		return m.StakerId
	}
	return ""
}

func (m *DelegationsByStaker) GetDelegations() []DelegatedSingleAssetInfo {
	if m != nil {
		return m.Delegations
	}
	return nil
}

// StakerToOperator is the association between a staker and an operator.
type StakerToOperator struct {
	// staker_id is the staker's account address + _ + l0 chain id (hex).
	StakerId string `protobuf:"bytes,1,opt,name=staker_id,json=stakerId,proto3" json:"staker_id,omitempty"`
	// operator is the bech32 address of the operator.
	Operator string `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *StakerToOperator) Reset()         { *m = StakerToOperator{} }
func (m *StakerToOperator) String() string { return proto.CompactTextString(m) }
func (*StakerToOperator) ProtoMessage()    {}
func (*StakerToOperator) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26dd0d733927603, []int{4}
}
func (m *StakerToOperator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakerToOperator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakerToOperator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakerToOperator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakerToOperator.Merge(m, src)
}
func (m *StakerToOperator) XXX_Size() int {
	return m.Size()
}
func (m *StakerToOperator) XXX_DiscardUnknown() {
	xxx_messageInfo_StakerToOperator.DiscardUnknown(m)
}

var xxx_messageInfo_StakerToOperator proto.InternalMessageInfo

func (m *StakerToOperator) GetStakerId() string {
	if m != nil {
		return m.StakerId
	}
	return ""
}

func (m *StakerToOperator) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "exocore.delegation.v1.GenesisState")
	proto.RegisterType((*DelegationStates)(nil), "exocore.delegation.v1.DelegationStates")
	proto.RegisterType((*StakersByOperator)(nil), "exocore.delegation.v1.StakersByOperator")
	proto.RegisterType((*DelegationsByStaker)(nil), "exocore.delegation.v1.DelegationsByStaker")
	proto.RegisterType((*StakerToOperator)(nil), "exocore.delegation.v1.StakerToOperator")
}

func init() {
	proto.RegisterFile("exocore/delegation/v1/genesis.proto", fileDescriptor_c26dd0d733927603)
}

var fileDescriptor_c26dd0d733927603 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdf, 0x8e, 0xd2, 0x40,
	0x14, 0xc6, 0x19, 0xc0, 0x75, 0x39, 0xbb, 0x26, 0x30, 0xbb, 0x26, 0x0d, 0x26, 0x15, 0xeb, 0x85,
	0x5c, 0xb5, 0xae, 0xfa, 0x02, 0xa0, 0xab, 0x12, 0x93, 0x35, 0x82, 0xc6, 0x64, 0x4d, 0x6c, 0x0a,
	0x1d, 0x6b, 0x43, 0xb7, 0x83, 0x3d, 0xc3, 0x4a, 0x9f, 0xc1, 0x1b, 0x1f, 0xc4, 0x07, 0xd9, 0xcb,
	0xbd, 0xf4, 0xca, 0x18, 0x78, 0x11, 0xd3, 0xce, 0x40, 0x4b, 0x65, 0x7b, 0x37, 0x73, 0xfe, 0xfc,
	0xce, 0xd7, 0xaf, 0x73, 0xe0, 0x21, 0x5b, 0xf0, 0x09, 0x8f, 0x98, 0xe5, 0xb2, 0x80, 0x79, 0x8e,
	0xf0, 0x79, 0x68, 0x5d, 0x9e, 0x58, 0x1e, 0x0b, 0x19, 0xfa, 0x68, 0xce, 0x22, 0x2e, 0x38, 0xbd,
	0xab, 0x8a, 0xcc, 0xac, 0xc8, 0xbc, 0x3c, 0x69, 0x3f, 0xd8, 0xdd, 0xfb, 0x6d, 0xce, 0xa2, 0x58,
	0x76, 0xb6, 0xf5, 0xdd, 0x25, 0x62, 0xa1, 0xf2, 0xc7, 0x1e, 0xf7, 0x78, 0x7a, 0xb4, 0x92, 0x93,
	0x8c, 0x1a, 0xbf, 0x6a, 0x70, 0xf8, 0x4a, 0x2a, 0x18, 0x09, 0x47, 0x30, 0xfa, 0x0e, 0x0e, 0x1d,
	0x44, 0x3e, 0xf1, 0x53, 0x02, 0x6a, 0xa4, 0x53, 0xeb, 0x1e, 0x3c, 0x79, 0x64, 0xee, 0xd4, 0x65,
	0x8e, 0x84, 0x33, 0x65, 0xd1, 0x7b, 0xfe, 0x76, 0xc6, 0x22, 0x47, 0xf0, 0xa8, 0x5f, 0xbf, 0xfa,
	0x73, 0xbf, 0x32, 0xdc, 0x42, 0xd0, 0x73, 0x68, 0x65, 0x5d, 0x36, 0x26, 0x63, 0x50, 0xab, 0x96,
	0x72, 0x5f, 0x6c, 0x6e, 0xa9, 0x2a, 0x54, 0xdc, 0xa6, 0x5b, 0x88, 0xd3, 0xcf, 0x70, 0x84, 0xa9,
	0x06, 0xb4, 0xc7, 0xb1, 0xcd, 0x95, 0x0c, 0xad, 0x96, 0xd2, 0xbb, 0xa5, 0xaa, 0xb1, 0x1f, 0x17,
	0x64, 0xb7, 0xb0, 0x98, 0xa0, 0x9f, 0xe0, 0xce, 0x3c, 0xcc, 0xba, 0x51, 0xab, 0xa7, 0x64, 0xeb,
	0x06, 0xf2, 0x87, 0x5c, 0x6d, 0x2f, 0x74, 0x5f, 0xf3, 0xc0, 0x7d, 0xce, 0xe7, 0xa1, 0x50, 0x03,
	0xb6, 0x59, 0xf4, 0x31, 0x1c, 0x07, 0x0e, 0x0a, 0x3b, 0x1f, 0xb5, 0x7d, 0x57, 0xbb, 0xd5, 0x21,
	0xdd, 0xfa, 0x90, 0x26, 0xb9, 0x3c, 0x70, 0xe0, 0x1a, 0x01, 0x34, 0x8b, 0xd6, 0xd0, 0x26, 0xd4,
	0xa6, 0x2c, 0xd6, 0x48, 0x87, 0x74, 0x1b, 0xc3, 0xe4, 0x48, 0x5f, 0xc2, 0xde, 0xc6, 0x65, 0x52,
	0xe2, 0x43, 0x86, 0xea, 0x5d, 0x24, 0x2a, 0xd7, 0x36, 0xab, 0x6e, 0xe3, 0x14, 0x5a, 0xff, 0x59,
	0xb5, 0x63, 0x9c, 0x0e, 0xb7, 0x95, 0x71, 0xe9, 0x5f, 0x6d, 0x28, 0xca, 0x3a, 0x68, 0xfc, 0x20,
	0x70, 0x94, 0x8d, 0xc2, 0x7e, 0x2c, 0xa1, 0xf4, 0x1e, 0x34, 0x64, 0x49, 0xf2, 0xcd, 0x92, 0xb7,
	0x2f, 0x03, 0x03, 0x97, 0x7e, 0x84, 0x83, 0xbc, 0xed, 0xd5, 0x52, 0xdb, 0x15, 0x9d, 0xb9, 0x23,
	0x3f, 0xf4, 0x02, 0xd6, 0x43, 0x64, 0x62, 0x10, 0x7e, 0xe1, 0x4a, 0x49, 0x9e, 0x64, 0xbc, 0x81,
	0x66, 0xf1, 0xd5, 0x96, 0x2b, 0x69, 0xc3, 0xfe, 0xe6, 0x5d, 0x55, 0x65, 0x6e, 0x7d, 0xef, 0x9f,
	0x5d, 0x2d, 0x75, 0x72, 0xbd, 0xd4, 0xc9, 0xdf, 0xa5, 0x4e, 0x7e, 0xae, 0xf4, 0xca, 0xf5, 0x4a,
	0xaf, 0xfc, 0x5e, 0xe9, 0x95, 0xf3, 0x67, 0x9e, 0x2f, 0xbe, 0xce, 0xc7, 0xe6, 0x84, 0x5f, 0x58,
	0xa7, 0x52, 0xf4, 0x19, 0x13, 0xdf, 0x79, 0x34, 0xb5, 0xd6, 0x8b, 0xba, 0xc8, 0xaf, 0xaa, 0x88,
	0x67, 0x0c, 0xc7, 0x7b, 0xe9, 0x56, 0x3e, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x9d, 0xf5,
	0xbf, 0x2c, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastUndelegationId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastUndelegationId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Undelegations) > 0 {
		for iNdEx := len(m.Undelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Undelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.StakersByOperator) > 0 {
		for iNdEx := len(m.StakersByOperator) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StakersByOperator[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DelegationStates) > 0 {
		for iNdEx := len(m.DelegationStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Associations) > 0 {
		for iNdEx := len(m.Associations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Associations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DelegationStates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegationStates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegationStates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.States.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StakersByOperator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakersByOperator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakersByOperator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Stakers) > 0 {
		for iNdEx := len(m.Stakers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Stakers[iNdEx])
			copy(dAtA[i:], m.Stakers[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.Stakers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DelegationsByStaker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegationsByStaker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegationsByStaker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegations) > 0 {
		for iNdEx := len(m.Delegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Delegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.StakerId) > 0 {
		i -= len(m.StakerId)
		copy(dAtA[i:], m.StakerId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.StakerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StakerToOperator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakerToOperator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakerToOperator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.StakerId) > 0 {
		i -= len(m.StakerId)
		copy(dAtA[i:], m.StakerId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.StakerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Associations) > 0 {
		for _, e := range m.Associations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DelegationStates) > 0 {
		for _, e := range m.DelegationStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StakersByOperator) > 0 {
		for _, e := range m.StakersByOperator {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Undelegations) > 0 {
		for _, e := range m.Undelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastUndelegationId != 0 {
		n += 1 + sovGenesis(uint64(m.LastUndelegationId))
	}
	return n
}

func (m *DelegationStates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.States.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *StakersByOperator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Stakers) > 0 {
		for _, s := range m.Stakers {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *DelegationsByStaker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StakerId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Delegations) > 0 {
		for _, e := range m.Delegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *StakerToOperator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StakerId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Associations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Associations = append(m.Associations, StakerToOperator{})
			if err := m.Associations[len(m.Associations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegationStates = append(m.DelegationStates, DelegationStates{})
			if err := m.DelegationStates[len(m.DelegationStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakersByOperator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakersByOperator = append(m.StakersByOperator, StakersByOperator{})
			if err := m.StakersByOperator[len(m.StakersByOperator)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Undelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Undelegations = append(m.Undelegations, UndelegationAndHoldCount{})
			if err := m.Undelegations[len(m.Undelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUndelegationId", wireType)
			}
			m.LastUndelegationId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastUndelegationId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *DelegationStates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: DelegationStates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegationStates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field States", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.States.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *StakersByOperator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: StakersByOperator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakersByOperator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stakers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stakers = append(m.Stakers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *DelegationsByStaker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: DelegationsByStaker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegationsByStaker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegations = append(m.Delegations, DelegatedSingleAssetInfo{})
			if err := m.Delegations[len(m.Delegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *StakerToOperator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: StakerToOperator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakerToOperator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
