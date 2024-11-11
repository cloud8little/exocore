// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/operator/v1/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	types1 "github.com/cosmos/cosmos-sdk/x/staking/types"
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

// Validator defines a validator, together with the total amount of the
// Validator's bond shares and their exchange rate to coins. Slashing results in
// a decrease in the exchange rate, allowing correct calculation of future
// undelegations without iterating over delegators. When coins are delegated to
// this validator, the validator is credited with a delegation whose number of
// bond shares is based on the amount of coins delegated divided by the current
// exchange rate. Voting power can be calculated as total bonded shares
// multiplied by exchange rate
type Validator struct {
	// earnoperator_earnings_addrings_addr is the earnings address.
	OperatorEarningsAddr string `protobuf:"bytes,1,opt,name=operator_earnings_addr,json=operatorEarningsAddr,proto3" json:"operator_earnings_addr,omitempty"`
	// operator_approve_addr is the approve address.
	OperatorApproveAddr string `protobuf:"bytes,2,opt,name=operator_approve_addr,json=operatorApproveAddr,proto3" json:"operator_approve_addr,omitempty"`
	// operator_meta_info is the operator meta info.
	OperatorMetaInfo string `protobuf:"bytes,3,opt,name=operator_meta_info,json=operatorMetaInfo,proto3" json:"operator_meta_info,omitempty"`
	// ConsAddress defines a wrapper around bytes meant to present a consensus node.
	// When marshaled to a string or JSON, it uses Bech32.
	ConsAddress string `protobuf:"bytes,4,opt,name=cons_address,json=consAddress,proto3" json:"cons_address,omitempty"`
	// consensus_pubkey is the consensus public key of the validator, as a Protobuf Any
	ConsensusPubkey *types.Any `protobuf:"bytes,5,opt,name=consensus_pubkey,json=consensusPubkey,proto3" json:"consensus_pubkey,omitempty"`
	// jailed defined whether the validator has been jailed from bonded status or not
	Jailed bool `protobuf:"varint,6,opt,name=jailed,proto3" json:"jailed,omitempty"`
	// status is the validator status (bonded/unbonding/unbonded)
	Status types1.BondStatus `protobuf:"varint,7,opt,name=status,proto3,enum=cosmos.staking.v1beta1.BondStatus" json:"status,omitempty"`
	// voting_power define the validator voting power
	VotingPower github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=voting_power,json=votingPower,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"voting_power"`
	// delegator_shares defines total shares issued to a validator's delegators
	DelegatorShares github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=delegator_shares,json=delegatorShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"delegator_shares"`
	// commission defines the commission parameters.
	Commission types1.Commission `protobuf:"bytes,10,opt,name=commission,proto3" json:"commission"`
	// delegator_tokens is the list of asset infos
	DelegatorTokens []DelegatorInfo `protobuf:"bytes,11,rep,name=delegator_tokens,json=delegatorTokens,proto3" json:"delegator_tokens"`
}

func (m *Validator) Reset()      { *m = Validator{} }
func (*Validator) ProtoMessage() {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8f7debe75430d24, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

// DelegatorInfo records the total opted-in USD value for the specified operator
type DelegatorInfo struct {
	// asset_id is the asset for which the query is made
	AssetID string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// symbol of the asset, like "USDT"
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// name of the asset, like "Tether USD"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// self_amount is the self amount of the asset which delegation
	SelfAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=self_amount,json=selfAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"self_amount"`
	// total_amount is the total amount of the asset which delegation
	TotalAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=total_amount,json=totalAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_amount"`
	// self_usd_value is the self delegation USD value for the validator
	SelfUSDValue github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=self_usd_value,json=selfUsdValue,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"self_usd_value"`
	// total_usd_value is the total delegation USD value for the validator
	TotalUSDValue github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=total_usd_value,json=totalUsdValue,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_usd_value"`
}

func (m *DelegatorInfo) Reset()         { *m = DelegatorInfo{} }
func (m *DelegatorInfo) String() string { return proto.CompactTextString(m) }
func (*DelegatorInfo) ProtoMessage()    {}
func (*DelegatorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8f7debe75430d24, []int{1}
}
func (m *DelegatorInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegatorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegatorInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegatorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatorInfo.Merge(m, src)
}
func (m *DelegatorInfo) XXX_Size() int {
	return m.Size()
}
func (m *DelegatorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatorInfo proto.InternalMessageInfo

func (m *DelegatorInfo) GetAssetID() string {
	if m != nil {
		return m.AssetID
	}
	return ""
}

func (m *DelegatorInfo) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *DelegatorInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Validator)(nil), "exocore.operator.v1.Validator")
	proto.RegisterType((*DelegatorInfo)(nil), "exocore.operator.v1.DelegatorInfo")
}

func init() {
	proto.RegisterFile("exocore/operator/v1/validator.proto", fileDescriptor_c8f7debe75430d24)
}

var fileDescriptor_c8f7debe75430d24 = []byte{
	// 769 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x3f, 0x6f, 0xfb, 0x44,
	0x18, 0x8e, 0x69, 0x9a, 0x34, 0x97, 0xf4, 0x0f, 0xd7, 0x50, 0x99, 0x0e, 0x49, 0x28, 0xa8, 0x8a,
	0x10, 0xb5, 0x95, 0xc0, 0x54, 0xb1, 0xc4, 0xa4, 0x43, 0x44, 0x8b, 0x2a, 0xa7, 0xad, 0x04, 0x12,
	0xb2, 0x2e, 0xf6, 0xc5, 0x35, 0xb1, 0xef, 0x2c, 0xdf, 0x39, 0x6d, 0xbe, 0x01, 0x13, 0x62, 0x64,
	0xec, 0xc8, 0xc8, 0xd0, 0x0f, 0x51, 0x31, 0x55, 0x4c, 0x88, 0x21, 0x42, 0xe9, 0x00, 0x2b, 0xdf,
	0x00, 0xdd, 0xf9, 0x9c, 0x06, 0x44, 0x87, 0x8a, 0xdf, 0x92, 0xdc, 0x7b, 0xef, 0xfb, 0x3c, 0xcf,
	0x7b, 0xef, 0x3d, 0x3e, 0xf0, 0x3e, 0xbe, 0xa5, 0x2e, 0x4d, 0xb0, 0x49, 0x63, 0x9c, 0x20, 0x4e,
	0x13, 0x73, 0xda, 0x31, 0xa7, 0x28, 0x0c, 0x3c, 0x11, 0x18, 0x71, 0x42, 0x39, 0x85, 0xbb, 0xaa,
	0xc8, 0xc8, 0x8b, 0x8c, 0x69, 0x67, 0xff, 0x6d, 0x14, 0x05, 0x84, 0x9a, 0xf2, 0x37, 0xab, 0xdb,
	0xff, 0xc0, 0xa5, 0x2c, 0xa2, 0xcc, 0x64, 0x1c, 0x4d, 0x02, 0xe2, 0x9b, 0xd3, 0xce, 0x08, 0x73,
	0xd4, 0xc9, 0x63, 0x55, 0xf5, 0x6e, 0x56, 0xe5, 0xc8, 0xc8, 0xcc, 0x02, 0x95, 0xaa, 0xfb, 0xd4,
	0xa7, 0xd9, 0xbe, 0x58, 0xe5, 0x00, 0x9f, 0x52, 0x3f, 0xc4, 0xa6, 0x8c, 0x46, 0xe9, 0xd8, 0x44,
	0x64, 0x96, 0xa5, 0x0e, 0xfe, 0x5a, 0x07, 0x95, 0xab, 0xbc, 0x5b, 0xf8, 0x09, 0xd8, 0xcb, 0x3b,
	0x74, 0x30, 0x4a, 0x48, 0x40, 0x7c, 0xe6, 0x20, 0xcf, 0x4b, 0x74, 0xad, 0xa5, 0xb5, 0x2b, 0x76,
	0x3d, 0xcf, 0x9e, 0xa8, 0x64, 0xcf, 0xf3, 0x12, 0xd8, 0x05, 0xef, 0x2c, 0x51, 0x28, 0x8e, 0x13,
	0x3a, 0xc5, 0x19, 0xe8, 0x2d, 0x09, 0xda, 0xcd, 0x93, 0xbd, 0x2c, 0x27, 0x31, 0x1f, 0x01, 0xb8,
	0xc4, 0x44, 0x98, 0x23, 0x27, 0x20, 0x63, 0xaa, 0xaf, 0x49, 0xc0, 0x4e, 0x9e, 0x39, 0xc3, 0x1c,
	0x0d, 0xc8, 0x98, 0xc2, 0xf7, 0x40, 0xcd, 0xa5, 0x24, 0x6b, 0x05, 0x33, 0xa6, 0x17, 0x65, 0x5d,
	0x55, 0xec, 0xf5, 0xb2, 0x2d, 0xf8, 0x25, 0xd8, 0x11, 0x21, 0x26, 0x2c, 0x65, 0x4e, 0x9c, 0x8e,
	0x26, 0x78, 0xa6, 0xaf, 0xb7, 0xb4, 0x76, 0xb5, 0x5b, 0x37, 0xb2, 0xe3, 0x1b, 0xf9, 0xf1, 0x8d,
	0x1e, 0x99, 0x59, 0xfa, 0xcf, 0xf7, 0x47, 0x75, 0x35, 0x3b, 0x37, 0x99, 0xc5, 0x9c, 0x1a, 0xe7,
	0xe9, 0xe8, 0x73, 0x3c, 0xb3, 0xb7, 0x97, 0x3c, 0xe7, 0x92, 0x06, 0xee, 0x81, 0xd2, 0x37, 0x28,
	0x08, 0xb1, 0xa7, 0x97, 0x5a, 0x5a, 0x7b, 0xc3, 0x56, 0x11, 0x3c, 0x06, 0x25, 0xc6, 0x11, 0x4f,
	0x99, 0x5e, 0x6e, 0x69, 0xed, 0xad, 0xee, 0x81, 0xa1, 0xf8, 0xf2, 0xeb, 0x52, 0xd7, 0x67, 0x58,
	0x94, 0x78, 0x43, 0x59, 0x69, 0x2b, 0x04, 0x74, 0x40, 0x6d, 0x4a, 0x79, 0x40, 0x7c, 0x27, 0xa6,
	0x37, 0x38, 0xd1, 0x37, 0xc4, 0x89, 0xac, 0x4f, 0x1f, 0xe6, 0xcd, 0xc2, 0x6f, 0xf3, 0xe6, 0xa1,
	0x1f, 0xf0, 0xeb, 0x74, 0x64, 0xb8, 0x34, 0x52, 0xf7, 0xab, 0xfe, 0x8e, 0x98, 0x37, 0x31, 0xf9,
	0x2c, 0xc6, 0xcc, 0xe8, 0x63, 0xf7, 0x97, 0xfb, 0x23, 0xa0, 0x24, 0xfb, 0xd8, 0xb5, 0xab, 0x19,
	0xe3, 0xb9, 0x20, 0x84, 0x3e, 0xd8, 0xf1, 0x70, 0x88, 0x7d, 0x39, 0x61, 0x76, 0x8d, 0x12, 0xcc,
	0xf4, 0xca, 0xab, 0x45, 0x06, 0x84, 0xaf, 0x88, 0x0c, 0x08, 0xb7, 0xb7, 0x97, 0xac, 0x43, 0x49,
	0x0a, 0xcf, 0x00, 0x70, 0x69, 0x14, 0x05, 0x8c, 0x05, 0x94, 0xe8, 0x40, 0x8e, 0xfc, 0xc5, 0x49,
	0x7c, 0xb6, 0xac, 0xb4, 0x2a, 0xa2, 0x8d, 0x1f, 0xff, 0xf8, 0xe9, 0x43, 0xcd, 0x5e, 0x21, 0x80,
	0xc3, 0xd5, 0xbe, 0x39, 0x9d, 0x60, 0xc2, 0xf4, 0x6a, 0x6b, 0x4d, 0x92, 0xfe, 0xc7, 0x57, 0x64,
	0xf4, 0xf3, 0x62, 0x61, 0x14, 0xab, 0x28, 0x48, 0x57, 0x7a, 0xbc, 0x90, 0x04, 0xc7, 0xb5, 0x6f,
	0xef, 0x9a, 0x85, 0x1f, 0xee, 0x9a, 0x85, 0x3f, 0xef, 0x9a, 0x85, 0x83, 0xef, 0x8a, 0x60, 0xf3,
	0x1f, 0x30, 0x78, 0x08, 0x36, 0x10, 0x63, 0x98, 0x3b, 0x81, 0x97, 0x39, 0xdd, 0xaa, 0x2e, 0xe6,
	0xcd, 0x72, 0x4f, 0xec, 0x0d, 0xfa, 0x76, 0x59, 0x26, 0x07, 0x9e, 0x70, 0x02, 0x9b, 0x45, 0x23,
	0x1a, 0x2a, 0x6b, 0xab, 0x08, 0x42, 0x50, 0x24, 0x28, 0xc2, 0xca, 0xbf, 0x72, 0x0d, 0xbf, 0x06,
	0x55, 0x86, 0xc3, 0xb1, 0x83, 0x22, 0x9a, 0x12, 0x9e, 0x59, 0xf6, 0x7f, 0xce, 0x1e, 0x08, 0xc2,
	0x9e, 0xe4, 0x13, 0x06, 0xe2, 0x94, 0xa3, 0x30, 0xe7, 0x5f, 0x7f, 0x03, 0xfc, 0x55, 0xc9, 0xa8,
	0x04, 0x12, 0xb0, 0x25, 0xfb, 0x4f, 0x99, 0xe7, 0x4c, 0x51, 0x98, 0x62, 0xe9, 0xfe, 0x8a, 0x75,
	0xfa, 0x3a, 0x8f, 0x2e, 0xe6, 0xcd, 0xda, 0x10, 0x87, 0xe3, 0xcb, 0x61, 0xff, 0x4a, 0xb0, 0xfc,
	0xcb, 0xb3, 0x35, 0xa1, 0x71, 0xc9, 0x3c, 0x99, 0x83, 0x29, 0xd8, 0xce, 0x0e, 0xf5, 0x2c, 0x5a,
	0x96, 0xa2, 0x67, 0xaf, 0x16, 0xdd, 0xbc, 0x10, 0x44, 0x2f, 0xa8, 0x6e, 0x4a, 0x95, 0x5c, 0xd6,
	0x3a, 0x7d, 0x58, 0x34, 0xb4, 0xc7, 0x45, 0x43, 0xfb, 0x7d, 0xd1, 0xd0, 0xbe, 0x7f, 0x6a, 0x14,
	0x1e, 0x9f, 0x1a, 0x85, 0x5f, 0x9f, 0x1a, 0x85, 0xaf, 0xba, 0x2b, 0x7a, 0x27, 0x99, 0xfb, 0xbe,
	0xc0, 0xfc, 0x86, 0x26, 0x13, 0x33, 0x7f, 0xf7, 0x6f, 0x9f, 0x5f, 0x7e, 0xa9, 0x3f, 0x2a, 0xc9,
	0x77, 0xe6, 0xe3, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x93, 0x73, 0xe8, 0x9b, 0x1a, 0x06, 0x00,
	0x00,
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DelegatorTokens) > 0 {
		for iNdEx := len(m.DelegatorTokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatorTokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintValidator(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	{
		size, err := m.Commission.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.DelegatorShares.Size()
		i -= size
		if _, err := m.DelegatorShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.VotingPower.Size()
		i -= size
		if _, err := m.VotingPower.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.Status != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.ConsensusPubkey != nil {
		{
			size, err := m.ConsensusPubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ConsAddress) > 0 {
		i -= len(m.ConsAddress)
		copy(dAtA[i:], m.ConsAddress)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.ConsAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OperatorMetaInfo) > 0 {
		i -= len(m.OperatorMetaInfo)
		copy(dAtA[i:], m.OperatorMetaInfo)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.OperatorMetaInfo)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OperatorApproveAddr) > 0 {
		i -= len(m.OperatorApproveAddr)
		copy(dAtA[i:], m.OperatorApproveAddr)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.OperatorApproveAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OperatorEarningsAddr) > 0 {
		i -= len(m.OperatorEarningsAddr)
		copy(dAtA[i:], m.OperatorEarningsAddr)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.OperatorEarningsAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DelegatorInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatorInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegatorInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalUSDValue.Size()
		i -= size
		if _, err := m.TotalUSDValue.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.SelfUSDValue.Size()
		i -= size
		if _, err := m.SelfUSDValue.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.TotalAmount.Size()
		i -= size
		if _, err := m.TotalAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.SelfAmount.Size()
		i -= size
		if _, err := m.SelfAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AssetID) > 0 {
		i -= len(m.AssetID)
		copy(dAtA[i:], m.AssetID)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.AssetID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorEarningsAddr)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.OperatorApproveAddr)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.OperatorMetaInfo)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.ConsAddress)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.ConsensusPubkey != nil {
		l = m.ConsensusPubkey.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.Jailed {
		n += 2
	}
	if m.Status != 0 {
		n += 1 + sovValidator(uint64(m.Status))
	}
	l = m.VotingPower.Size()
	n += 1 + l + sovValidator(uint64(l))
	l = m.DelegatorShares.Size()
	n += 1 + l + sovValidator(uint64(l))
	l = m.Commission.Size()
	n += 1 + l + sovValidator(uint64(l))
	if len(m.DelegatorTokens) > 0 {
		for _, e := range m.DelegatorTokens {
			l = e.Size()
			n += 1 + l + sovValidator(uint64(l))
		}
	}
	return n
}

func (m *DelegatorInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AssetID)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = m.SelfAmount.Size()
	n += 1 + l + sovValidator(uint64(l))
	l = m.TotalAmount.Size()
	n += 1 + l + sovValidator(uint64(l))
	l = m.SelfUSDValue.Size()
	n += 1 + l + sovValidator(uint64(l))
	l = m.TotalUSDValue.Size()
	n += 1 + l + sovValidator(uint64(l))
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorEarningsAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorEarningsAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorApproveAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorApproveAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorMetaInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorMetaInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusPubkey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsensusPubkey == nil {
				m.ConsensusPubkey = &types.Any{}
			}
			if err := m.ConsensusPubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types1.BondStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VotingPower.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegatorShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Commission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorTokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorTokens = append(m.DelegatorTokens, DelegatorInfo{})
			if err := m.DelegatorTokens[len(m.DelegatorTokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *DelegatorInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: DelegatorInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatorInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelfAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SelfAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelfUSDValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SelfUSDValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalUSDValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalUSDValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
