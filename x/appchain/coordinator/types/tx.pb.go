// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imuachain/appchain/coordinator/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/imua-xyz/imuachain/x/appchain/common/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// RegisterSubscriberChainRequest is the request type for the
// RegisterSubscriberChain message.
type RegisterSubscriberChainRequest struct {
	// from_address is the address of the transaction signer. any transactions
	// originating from this address may be used to edit the chain. at some point
	// in the future this will be offloaded to the governance module on the
	// subscriber chain. (TODO)
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	// chain_id is the unique identifier for the chain, serving as the primary key.
	ChainID string `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// epoch_identifier specifies the unit of epoch (week, hour, day). It must be
	// registered in the x/epochs module.
	// This epoch is the identifier used by the coordinator to send validator set
	// updates to the subscriber at the end of each epoch. The subscriber chain's
	// genesis is made available at the end of the current epoch
	// (marked by this identifier).
	EpochIdentifier string `protobuf:"bytes,3,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty"`
	// asset_ids lists the IDs of assets accepted by the subscriber chain.
	AssetIDs []string `protobuf:"bytes,4,rep,name=asset_ids,json=assetIds,proto3" json:"asset_ids,omitempty"`
	// min_self_delegation_usd is the minimum self-delegation in USD required to
	// be a validator on the chain.
	MinSelfDelegationUsd uint64 `protobuf:"varint,5,opt,name=min_self_delegation_usd,json=minSelfDelegationUsd,proto3" json:"min_self_delegation_usd,omitempty"`
	// max_validators is the maximum number of validators allowed on the chain.
	MaxValidators uint32 `protobuf:"varint,6,opt,name=max_validators,json=maxValidators,proto3" json:"max_validators,omitempty"`
	// subscriber_params are the parameters used by the subscriber module
	// on the subscriber chain.
	SubscriberParams types.SubscriberParams `protobuf:"bytes,7,opt,name=subscriber_params,json=subscriberParams,proto3" json:"subscriber_params"`
}

func (m *RegisterSubscriberChainRequest) Reset()         { *m = RegisterSubscriberChainRequest{} }
func (m *RegisterSubscriberChainRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterSubscriberChainRequest) ProtoMessage()    {}
func (*RegisterSubscriberChainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c979f131765613e4, []int{0}
}
func (m *RegisterSubscriberChainRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterSubscriberChainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterSubscriberChainRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterSubscriberChainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterSubscriberChainRequest.Merge(m, src)
}
func (m *RegisterSubscriberChainRequest) XXX_Size() int {
	return m.Size()
}
func (m *RegisterSubscriberChainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterSubscriberChainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterSubscriberChainRequest proto.InternalMessageInfo

func (m *RegisterSubscriberChainRequest) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *RegisterSubscriberChainRequest) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *RegisterSubscriberChainRequest) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *RegisterSubscriberChainRequest) GetAssetIDs() []string {
	if m != nil {
		return m.AssetIDs
	}
	return nil
}

func (m *RegisterSubscriberChainRequest) GetMinSelfDelegationUsd() uint64 {
	if m != nil {
		return m.MinSelfDelegationUsd
	}
	return 0
}

func (m *RegisterSubscriberChainRequest) GetMaxValidators() uint32 {
	if m != nil {
		return m.MaxValidators
	}
	return 0
}

func (m *RegisterSubscriberChainRequest) GetSubscriberParams() types.SubscriberParams {
	if m != nil {
		return m.SubscriberParams
	}
	return types.SubscriberParams{}
}

// RegisterSubscriberChainResponse defines the response structure for executing a
// RegisterSubscriberChain message.
type RegisterSubscriberChainResponse struct {
}

func (m *RegisterSubscriberChainResponse) Reset()         { *m = RegisterSubscriberChainResponse{} }
func (m *RegisterSubscriberChainResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterSubscriberChainResponse) ProtoMessage()    {}
func (*RegisterSubscriberChainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c979f131765613e4, []int{1}
}
func (m *RegisterSubscriberChainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterSubscriberChainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterSubscriberChainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterSubscriberChainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterSubscriberChainResponse.Merge(m, src)
}
func (m *RegisterSubscriberChainResponse) XXX_Size() int {
	return m.Size()
}
func (m *RegisterSubscriberChainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterSubscriberChainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterSubscriberChainResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterSubscriberChainRequest)(nil), "imuachain.appchain.coordinator.v1.RegisterSubscriberChainRequest")
	proto.RegisterType((*RegisterSubscriberChainResponse)(nil), "imuachain.appchain.coordinator.v1.RegisterSubscriberChainResponse")
}

func init() {
	proto.RegisterFile("imuachain/appchain/coordinator/v1/tx.proto", fileDescriptor_c979f131765613e4)
}

var fileDescriptor_c979f131765613e4 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xcf, 0x36, 0xfd, 0x97, 0x49, 0xab, 0x71, 0x09, 0x64, 0x0d, 0xb2, 0x49, 0x03, 0x4a, 0x12,
	0xe8, 0x2e, 0xa9, 0x78, 0xa9, 0x88, 0x24, 0x46, 0x21, 0x07, 0x41, 0x36, 0xea, 0xc1, 0xcb, 0x32,
	0xc9, 0x4c, 0x36, 0x03, 0x99, 0x9d, 0x75, 0xdf, 0x24, 0xa4, 0x9e, 0xa4, 0x9f, 0x40, 0xf0, 0x7b,
	0x48, 0x0f, 0x7e, 0x88, 0x1e, 0x8b, 0x5e, 0x3c, 0x05, 0x49, 0x84, 0x7c, 0x03, 0xcf, 0xb2, 0xb3,
	0x69, 0x5a, 0x24, 0x51, 0xc1, 0xdb, 0xcc, 0xef, 0xcf, 0xfc, 0x1e, 0xef, 0xbd, 0x41, 0x55, 0xc6,
	0x87, 0xb8, 0xdb, 0xc7, 0xcc, 0xb7, 0x71, 0x10, 0xc4, 0x87, 0xae, 0x10, 0x21, 0x61, 0x3e, 0x96,
	0x22, 0xb4, 0x47, 0x35, 0x5b, 0x8e, 0xad, 0x20, 0x14, 0x52, 0xe8, 0x07, 0x4b, 0xad, 0x75, 0xa9,
	0xb5, 0xae, 0x69, 0xad, 0x51, 0x2d, 0x9f, 0xeb, 0x0a, 0xe0, 0x02, 0x6c, 0x0e, 0x5e, 0x64, 0xe5,
	0xe0, 0xc5, 0xde, 0xfc, 0xed, 0x98, 0x70, 0xd5, 0xcd, 0x8e, 0x2f, 0x0b, 0x2a, 0xeb, 0x09, 0x4f,
	0xc4, 0x78, 0x74, 0x5a, 0xa0, 0x77, 0x3c, 0x21, 0xbc, 0x01, 0xb5, 0x71, 0xc0, 0x6c, 0xec, 0xfb,
	0x42, 0x62, 0xc9, 0x84, 0x7f, 0xe9, 0xa9, 0xac, 0x2c, 0x9b, 0x73, 0xe1, 0x47, 0xb1, 0xf1, 0x29,
	0x96, 0x96, 0x3e, 0x25, 0x91, 0xe9, 0x50, 0x8f, 0x81, 0xa4, 0x61, 0x7b, 0xd8, 0x81, 0x6e, 0xc8,
	0x3a, 0x34, 0x7c, 0x12, 0x59, 0x1c, 0xfa, 0x76, 0x48, 0x41, 0xea, 0x0f, 0xd1, 0x5e, 0x2f, 0x14,
	0xdc, 0xc5, 0x84, 0x84, 0x14, 0xc0, 0xd0, 0x8a, 0x5a, 0x39, 0xd5, 0x30, 0xbe, 0x7c, 0x3e, 0xcc,
	0x2e, 0x2a, 0xad, 0xc7, 0x4c, 0x5b, 0x86, 0xcc, 0xf7, 0x9c, 0x74, 0xa4, 0x5e, 0x40, 0xfa, 0x3d,
	0xb4, 0xab, 0xf2, 0x5d, 0x46, 0x8c, 0x0d, 0x65, 0x4c, 0x4f, 0x27, 0x85, 0x1d, 0x15, 0xd0, 0x6a,
	0x3a, 0x3b, 0x8a, 0x6c, 0x11, 0xbd, 0x82, 0x32, 0x34, 0x10, 0xdd, 0xbe, 0xcb, 0x08, 0xf5, 0x25,
	0xeb, 0x31, 0x1a, 0x1a, 0xc9, 0x48, 0xef, 0xdc, 0x54, 0x78, 0x6b, 0x09, 0xeb, 0x15, 0x94, 0xc2,
	0x00, 0x54, 0xba, 0x8c, 0x80, 0xb1, 0x59, 0x4c, 0x96, 0x53, 0x8d, 0xbd, 0xe9, 0xa4, 0xb0, 0x5b,
	0x8f, 0xc0, 0x56, 0x13, 0x9c, 0x5d, 0x45, 0xb7, 0x08, 0xe8, 0x0f, 0x50, 0x8e, 0x33, 0xdf, 0x05,
	0x3a, 0xe8, 0xb9, 0x84, 0x0e, 0xa8, 0xa7, 0xda, 0xe4, 0x0e, 0x81, 0x18, 0x5b, 0x45, 0xad, 0xbc,
	0xe9, 0x64, 0x39, 0xf3, 0xdb, 0x74, 0xd0, 0x6b, 0x2e, 0xc9, 0x57, 0x40, 0xf4, 0xbb, 0xe8, 0x06,
	0xc7, 0x63, 0x77, 0x84, 0x07, 0x8c, 0x44, 0xb3, 0x03, 0x63, 0xbb, 0xa8, 0x95, 0xf7, 0x9d, 0x7d,
	0x8e, 0xc7, 0xaf, 0x97, 0xa0, 0x8e, 0xd1, 0x2d, 0x58, 0xb6, 0xcc, 0x0d, 0x70, 0x88, 0x39, 0x18,
	0x3b, 0x45, 0xad, 0x9c, 0x3e, 0xb2, 0xac, 0x95, 0xdb, 0xa0, 0x1a, 0x3f, 0xaa, 0x59, 0x57, 0x9d,
	0x7e, 0xa1, 0x5c, 0x8d, 0xcd, 0xf3, 0x49, 0x21, 0xe1, 0x64, 0xe0, 0x37, 0xfc, 0x38, 0x73, 0x3a,
	0x3f, 0xab, 0xa6, 0x9f, 0x5d, 0x35, 0xb4, 0x74, 0x80, 0x0a, 0x6b, 0xe7, 0x05, 0x81, 0xf0, 0x81,
	0x1e, 0xfd, 0xd4, 0x50, 0xf2, 0x39, 0x78, 0xfa, 0x5c, 0x43, 0xb9, 0x35, 0x5a, 0xbd, 0x6e, 0xfd,
	0x75, 0x5d, 0xad, 0x3f, 0xef, 0x45, 0xbe, 0xf1, 0x3f, 0x4f, 0xc4, 0xa5, 0x96, 0x9e, 0x9e, 0x7e,
	0xfd, 0xf1, 0x71, 0xe3, 0x71, 0xe9, 0x91, 0xfd, 0x2f, 0x3f, 0xcd, 0x5e, 0xf3, 0x5c, 0x7e, 0xeb,
	0xfd, 0xfc, 0xac, 0xaa, 0x35, 0x5e, 0x9e, 0x4f, 0x4d, 0xed, 0x62, 0x6a, 0x6a, 0xdf, 0xa7, 0xa6,
	0xf6, 0x61, 0x66, 0x26, 0x2e, 0x66, 0x66, 0xe2, 0xdb, 0xcc, 0x4c, 0xbc, 0x39, 0xf6, 0x98, 0xec,
	0x0f, 0x3b, 0xd1, 0x08, 0x54, 0xd2, 0xe1, 0xf8, 0xe4, 0xdd, 0xb5, 0xc8, 0xf1, 0xea, 0x50, 0x79,
	0x12, 0x50, 0xe8, 0x6c, 0xab, 0x9f, 0x72, 0xff, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xf7,
	0xe8, 0x6f, 0x0d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// RegisterSubscriberChain registers a subscriber chain with the coordinator.
	// By default, it is activated at the next epoch.
	RegisterSubscriberChain(ctx context.Context, in *RegisterSubscriberChainRequest, opts ...grpc.CallOption) (*RegisterSubscriberChainResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterSubscriberChain(ctx context.Context, in *RegisterSubscriberChainRequest, opts ...grpc.CallOption) (*RegisterSubscriberChainResponse, error) {
	out := new(RegisterSubscriberChainResponse)
	err := c.cc.Invoke(ctx, "/imuachain.appchain.coordinator.v1.Msg/RegisterSubscriberChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RegisterSubscriberChain registers a subscriber chain with the coordinator.
	// By default, it is activated at the next epoch.
	RegisterSubscriberChain(context.Context, *RegisterSubscriberChainRequest) (*RegisterSubscriberChainResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterSubscriberChain(ctx context.Context, req *RegisterSubscriberChainRequest) (*RegisterSubscriberChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSubscriberChain not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterSubscriberChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterSubscriberChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterSubscriberChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imuachain.appchain.coordinator.v1.Msg/RegisterSubscriberChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterSubscriberChain(ctx, req.(*RegisterSubscriberChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imuachain.appchain.coordinator.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterSubscriberChain",
			Handler:    _Msg_RegisterSubscriberChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imuachain/appchain/coordinator/v1/tx.proto",
}

func (m *RegisterSubscriberChainRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterSubscriberChainRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterSubscriberChainRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SubscriberParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.MaxValidators != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MaxValidators))
		i--
		dAtA[i] = 0x30
	}
	if m.MinSelfDelegationUsd != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MinSelfDelegationUsd))
		i--
		dAtA[i] = 0x28
	}
	if len(m.AssetIDs) > 0 {
		for iNdEx := len(m.AssetIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AssetIDs[iNdEx])
			copy(dAtA[i:], m.AssetIDs[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.AssetIDs[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterSubscriberChainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterSubscriberChainResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterSubscriberChainResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterSubscriberChainRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.AssetIDs) > 0 {
		for _, s := range m.AssetIDs {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.MinSelfDelegationUsd != 0 {
		n += 1 + sovTx(uint64(m.MinSelfDelegationUsd))
	}
	if m.MaxValidators != 0 {
		n += 1 + sovTx(uint64(m.MaxValidators))
	}
	l = m.SubscriberParams.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *RegisterSubscriberChainResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterSubscriberChainRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: RegisterSubscriberChainRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterSubscriberChainRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetIDs = append(m.AssetIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSelfDelegationUsd", wireType)
			}
			m.MinSelfDelegationUsd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSelfDelegationUsd |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxValidators", wireType)
			}
			m.MaxValidators = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxValidators |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriberParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SubscriberParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *RegisterSubscriberChainResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: RegisterSubscriberChainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterSubscriberChainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
