// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/oracle/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgCreatePrice provide the price updating message
type MsgCreatePrice struct {
	// creator tells which is the message sender and should sign this message
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// refer to id from Params.TokenFeeders, 0 is reserved, invalid to use
	FeederID uint64 `protobuf:"varint,2,opt,name=feeder_id,json=feederId,proto3" json:"feeder_id,omitempty"`
	// prices price with its corresponding source
	Prices []*PriceSource `protobuf:"bytes,3,rep,name=prices,proto3" json:"prices,omitempty"`
	// on which block commit does this message be built on
	BasedBlock uint64 `protobuf:"varint,4,opt,name=based_block,json=basedBlock,proto3" json:"based_block,omitempty"`
	// nonce represents the unique number to disginguish duplicated messages
	Nonce int32 `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *MsgCreatePrice) Reset()         { *m = MsgCreatePrice{} }
func (m *MsgCreatePrice) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePrice) ProtoMessage()    {}
func (*MsgCreatePrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d4ae88d3a4a094d, []int{0}
}
func (m *MsgCreatePrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePrice.Merge(m, src)
}
func (m *MsgCreatePrice) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePrice) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePrice.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePrice proto.InternalMessageInfo

func (m *MsgCreatePrice) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreatePrice) GetFeederID() uint64 {
	if m != nil {
		return m.FeederID
	}
	return 0
}

func (m *MsgCreatePrice) GetPrices() []*PriceSource {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *MsgCreatePrice) GetBasedBlock() uint64 {
	if m != nil {
		return m.BasedBlock
	}
	return 0
}

func (m *MsgCreatePrice) GetNonce() int32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

// MsgCreatePriceResponse
type MsgCreatePriceResponse struct {
}

func (m *MsgCreatePriceResponse) Reset()         { *m = MsgCreatePriceResponse{} }
func (m *MsgCreatePriceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePriceResponse) ProtoMessage()    {}
func (*MsgCreatePriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d4ae88d3a4a094d, []int{1}
}
func (m *MsgCreatePriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePriceResponse.Merge(m, src)
}
func (m *MsgCreatePriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePriceResponse proto.InternalMessageInfo

// MsgUpdateParms
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the x/staking parameters to update.
	//
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d4ae88d3a4a094d, []int{2}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d4ae88d3a4a094d, []int{3}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreatePrice)(nil), "exocore.oracle.v1.MsgCreatePrice")
	proto.RegisterType((*MsgCreatePriceResponse)(nil), "exocore.oracle.v1.MsgCreatePriceResponse")
	proto.RegisterType((*MsgUpdateParams)(nil), "exocore.oracle.v1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "exocore.oracle.v1.MsgUpdateParamsResponse")
}

func init() { proto.RegisterFile("exocore/oracle/v1/tx.proto", fileDescriptor_0d4ae88d3a4a094d) }

var fileDescriptor_0d4ae88d3a4a094d = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xce, 0x91, 0x26, 0x34, 0x97, 0x0a, 0xd4, 0x53, 0x44, 0x9d, 0x48, 0x75, 0x52, 0xb3, 0xa4,
	0x91, 0x62, 0xd3, 0x20, 0x65, 0x00, 0x16, 0xcc, 0x0f, 0xa9, 0x48, 0x41, 0xc8, 0x15, 0x0c, 0x20,
	0x11, 0x39, 0xf6, 0xe1, 0x5a, 0x89, 0xfd, 0xac, 0x3b, 0xa7, 0xa4, 0x2b, 0x23, 0x13, 0x7f, 0x06,
	0x63, 0x86, 0xee, 0x48, 0x4c, 0x1d, 0xab, 0x4e, 0x4c, 0x15, 0x4a, 0x86, 0xfc, 0x1b, 0xc8, 0x77,
	0x36, 0x6d, 0xda, 0x48, 0x5d, 0x2c, 0xbf, 0xf7, 0x7d, 0xef, 0xc7, 0xf7, 0x3d, 0x1b, 0xd7, 0xe8,
	0x04, 0x1c, 0x60, 0xd4, 0x00, 0x66, 0x3b, 0x23, 0x6a, 0x1c, 0xed, 0x19, 0xf1, 0x44, 0x8f, 0x18,
	0xc4, 0x40, 0x36, 0x53, 0x4c, 0x97, 0x98, 0x7e, 0xb4, 0x57, 0xdb, 0xb4, 0x03, 0x3f, 0x04, 0x43,
	0x3c, 0x25, 0xab, 0xb6, 0xe5, 0x00, 0x0f, 0x80, 0x1b, 0x01, 0xf7, 0x92, 0xea, 0x80, 0x7b, 0x29,
	0x50, 0x95, 0x40, 0x5f, 0x44, 0x86, 0x0c, 0x52, 0x48, 0xbd, 0x39, 0x35, 0xb2, 0x99, 0x1d, 0x64,
	0xf8, 0xf6, 0x0a, 0x9c, 0xf9, 0x0e, 0x4d, 0xe1, 0x8a, 0x07, 0x1e, 0xc8, 0xb6, 0xc9, 0x9b, 0xcc,
	0x6a, 0x0b, 0x84, 0xef, 0xf5, 0xb8, 0xf7, 0x82, 0x51, 0x3b, 0xa6, 0xef, 0x12, 0x3a, 0x79, 0x8a,
	0xef, 0x3a, 0x49, 0x08, 0x4c, 0x41, 0x0d, 0xd4, 0x2c, 0x99, 0x3b, 0xe7, 0x27, 0xed, 0xed, 0x74,
	0x95, 0x0f, 0xf6, 0xc8, 0x77, 0x13, 0xec, 0xb9, 0xeb, 0x32, 0xca, 0xf9, 0x41, 0xcc, 0xfc, 0xd0,
	0xb3, 0xb2, 0x0a, 0xb2, 0x8b, 0x4b, 0x5f, 0x28, 0x75, 0x29, 0xeb, 0xfb, 0xae, 0x72, 0xa7, 0x81,
	0x9a, 0x6b, 0xe6, 0xc6, 0xec, 0xa2, 0xbe, 0xfe, 0x5a, 0x24, 0xf7, 0x5f, 0x5a, 0xeb, 0x12, 0xde,
	0x77, 0x49, 0x17, 0x17, 0xc5, 0x7e, 0x5c, 0xc9, 0x37, 0xf2, 0xcd, 0x72, 0x47, 0xd5, 0x6f, 0x58,
	0xa7, 0x8b, 0x8d, 0x0e, 0x60, 0xcc, 0x1c, 0x6a, 0xa5, 0x6c, 0x52, 0xc7, 0xe5, 0x81, 0xcd, 0xa9,
	0xdb, 0x1f, 0x8c, 0xc0, 0x19, 0x2a, 0x6b, 0xc9, 0x10, 0x0b, 0x8b, 0x94, 0x99, 0x64, 0x48, 0x05,
	0x17, 0x42, 0x08, 0x1d, 0xaa, 0x14, 0x1a, 0xa8, 0x59, 0xb0, 0x64, 0xa0, 0x29, 0xf8, 0xc1, 0xb2,
	0x50, 0x8b, 0xf2, 0x08, 0x42, 0x4e, 0xb5, 0x5f, 0x08, 0xdf, 0xef, 0x71, 0xef, 0x7d, 0xe4, 0x26,
	0x90, 0xb0, 0x94, 0x74, 0x71, 0xc9, 0x1e, 0xc7, 0x87, 0xc0, 0xfc, 0xf8, 0x38, 0xb5, 0x41, 0x39,
	0x3f, 0x69, 0x57, 0x52, 0x1b, 0x96, 0xd5, 0x5f, 0x52, 0xc9, 0x33, 0x5c, 0x94, 0x47, 0x11, 0xe2,
	0xcb, 0x9d, 0xea, 0x2a, 0x51, 0x82, 0x60, 0x96, 0x4e, 0x2f, 0xea, 0xb9, 0x9f, 0x8b, 0x69, 0x0b,
	0x59, 0x69, 0xcd, 0x93, 0xee, 0xb7, 0xc5, 0xb4, 0x75, 0xd9, 0xed, 0xfb, 0x62, 0xda, 0x7a, 0x28,
	0x27, 0xb6, 0xb9, 0x3b, 0x34, 0x26, 0xd9, 0x69, 0xaf, 0x6d, 0xab, 0x55, 0xf1, 0xd6, 0xb5, 0x54,
	0x26, 0xae, 0xf3, 0x1b, 0xe1, 0x7c, 0x8f, 0x7b, 0xe4, 0x13, 0x2e, 0x5f, 0x3d, 0xf2, 0xce, 0x8a,
	0xbd, 0x96, 0xed, 0xa9, 0xed, 0xde, 0x4a, 0xc9, 0x86, 0x90, 0xcf, 0x78, 0x63, 0xc9, 0x3d, 0x6d,
	0x75, 0xe9, 0x55, 0x4e, 0xad, 0x75, 0x3b, 0x27, 0xeb, 0x6f, 0xbe, 0x39, 0x9d, 0xa9, 0xe8, 0x6c,
	0xa6, 0xa2, 0xbf, 0x33, 0x15, 0xfd, 0x98, 0xab, 0xb9, 0xb3, 0xb9, 0x9a, 0xfb, 0x33, 0x57, 0x73,
	0x1f, 0x1f, 0x79, 0x7e, 0x7c, 0x38, 0x1e, 0xe8, 0x0e, 0x04, 0xc6, 0x2b, 0xd9, 0xef, 0x2d, 0x8d,
	0xbf, 0x02, 0x1b, 0x1a, 0xd9, 0xef, 0xf0, 0xdf, 0xb5, 0xf8, 0x38, 0xa2, 0x7c, 0x50, 0x14, 0x1f,
	0xfe, 0xe3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xc0, 0xcc, 0xd8, 0xc5, 0x03, 0x00, 0x00,
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
	// CreatePrice creates price for a new oracle round
	CreatePrice(ctx context.Context, in *MsgCreatePrice, opts ...grpc.CallOption) (*MsgCreatePriceResponse, error)
	// UpdateParams update params value
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePrice(ctx context.Context, in *MsgCreatePrice, opts ...grpc.CallOption) (*MsgCreatePriceResponse, error) {
	out := new(MsgCreatePriceResponse)
	err := c.cc.Invoke(ctx, "/exocore.oracle.v1.Msg/CreatePrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/exocore.oracle.v1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreatePrice creates price for a new oracle round
	CreatePrice(context.Context, *MsgCreatePrice) (*MsgCreatePriceResponse, error)
	// UpdateParams update params value
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreatePrice(ctx context.Context, req *MsgCreatePrice) (*MsgCreatePriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrice not implemented")
}
func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreatePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exocore.oracle.v1.Msg/CreatePrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePrice(ctx, req.(*MsgCreatePrice))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exocore.oracle.v1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exocore.oracle.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePrice",
			Handler:    _Msg_CreatePrice_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exocore/oracle/v1/tx.proto",
}

func (m *MsgCreatePrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nonce != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x28
	}
	if m.BasedBlock != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BasedBlock))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Prices) > 0 {
		for iNdEx := len(m.Prices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.FeederID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.FeederID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreatePrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.FeederID != 0 {
		n += 1 + sovTx(uint64(m.FeederID))
	}
	if len(m.Prices) > 0 {
		for _, e := range m.Prices {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.BasedBlock != 0 {
		n += 1 + sovTx(uint64(m.BasedBlock))
	}
	if m.Nonce != 0 {
		n += 1 + sovTx(uint64(m.Nonce))
	}
	return n
}

func (m *MsgCreatePriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
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
func (m *MsgCreatePrice) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreatePrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeederID", wireType)
			}
			m.FeederID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
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
			m.Prices = append(m.Prices, &PriceSource{})
			if err := m.Prices[len(m.Prices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasedBlock", wireType)
			}
			m.BasedBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BasedBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgCreatePriceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreatePriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
