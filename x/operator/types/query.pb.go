// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/operator/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	crypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// QueryOperatorInfoReq is the request to obtain the operator information.
type GetOperatorInfoReq struct {
	// operator_addr is the operator address,its type should be a sdk.AccAddress
	OperatorAddr string `protobuf:"bytes,1,opt,name=operator_addr,json=operatorAddr,proto3" json:"operator_addr,omitempty"`
}

func (m *GetOperatorInfoReq) Reset()         { *m = GetOperatorInfoReq{} }
func (m *GetOperatorInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetOperatorInfoReq) ProtoMessage()    {}
func (*GetOperatorInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91e795a3cecbdbf, []int{0}
}
func (m *GetOperatorInfoReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetOperatorInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetOperatorInfoReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetOperatorInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperatorInfoReq.Merge(m, src)
}
func (m *GetOperatorInfoReq) XXX_Size() int {
	return m.Size()
}
func (m *GetOperatorInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperatorInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperatorInfoReq proto.InternalMessageInfo

func (m *GetOperatorInfoReq) GetOperatorAddr() string {
	if m != nil {
		return m.OperatorAddr
	}
	return ""
}

// QueryOperatorConsKeyRequest is the request to obtain the consensus public key of the operator
type QueryOperatorConsKeyRequest struct {
	// addr is the ACC address of operator
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// chain_id is the id of the chain served by the operator
	ChainId string `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *QueryOperatorConsKeyRequest) Reset()         { *m = QueryOperatorConsKeyRequest{} }
func (m *QueryOperatorConsKeyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryOperatorConsKeyRequest) ProtoMessage()    {}
func (*QueryOperatorConsKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91e795a3cecbdbf, []int{1}
}
func (m *QueryOperatorConsKeyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOperatorConsKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOperatorConsKeyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOperatorConsKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOperatorConsKeyRequest.Merge(m, src)
}
func (m *QueryOperatorConsKeyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryOperatorConsKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOperatorConsKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOperatorConsKeyRequest proto.InternalMessageInfo

func (m *QueryOperatorConsKeyRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *QueryOperatorConsKeyRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

// QueryOperatorConsKeyResponse is the response for QueryOperatorConsKeyRequest
type QueryOperatorConsKeyResponse struct {
	// public_key is the consensus public key of the operator
	PublicKey crypto.PublicKey `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key"`
}

func (m *QueryOperatorConsKeyResponse) Reset()         { *m = QueryOperatorConsKeyResponse{} }
func (m *QueryOperatorConsKeyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryOperatorConsKeyResponse) ProtoMessage()    {}
func (*QueryOperatorConsKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91e795a3cecbdbf, []int{2}
}
func (m *QueryOperatorConsKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOperatorConsKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOperatorConsKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOperatorConsKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOperatorConsKeyResponse.Merge(m, src)
}
func (m *QueryOperatorConsKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryOperatorConsKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOperatorConsKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOperatorConsKeyResponse proto.InternalMessageInfo

func (m *QueryOperatorConsKeyResponse) GetPublicKey() crypto.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return crypto.PublicKey{}
}

func init() {
	proto.RegisterType((*GetOperatorInfoReq)(nil), "exocore.operator.v1.GetOperatorInfoReq")
	proto.RegisterType((*QueryOperatorConsKeyRequest)(nil), "exocore.operator.v1.QueryOperatorConsKeyRequest")
	proto.RegisterType((*QueryOperatorConsKeyResponse)(nil), "exocore.operator.v1.QueryOperatorConsKeyResponse")
}

func init() { proto.RegisterFile("exocore/operator/v1/query.proto", fileDescriptor_f91e795a3cecbdbf) }

var fileDescriptor_f91e795a3cecbdbf = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0xa3, 0xf2, 0xd1, 0x05, 0x84, 0xb4, 0xf4, 0x90, 0x86, 0xc8, 0x05, 0x1f, 0xa0, 0x17,
	0x76, 0x49, 0x38, 0x73, 0x48, 0xc2, 0x87, 0xa2, 0x56, 0x7c, 0xb8, 0x37, 0x2e, 0x96, 0x63, 0x0f,
	0xae, 0x95, 0x64, 0xc7, 0xd9, 0xdd, 0x94, 0x58, 0x55, 0x2f, 0xdc, 0x91, 0x90, 0xf8, 0x2b, 0xfc,
	0x88, 0x1e, 0x23, 0xb8, 0x70, 0x42, 0x28, 0xe1, 0xc4, 0xaf, 0x40, 0x5e, 0xdb, 0x0d, 0x0a, 0x16,
	0x12, 0x37, 0xef, 0xec, 0x9b, 0xf7, 0x66, 0xde, 0x5b, 0x93, 0x3d, 0x98, 0x63, 0x80, 0x12, 0x38,
	0x26, 0x20, 0x7d, 0x8d, 0x92, 0x9f, 0xb4, 0xf9, 0x74, 0x06, 0x32, 0x65, 0x89, 0x44, 0x8d, 0xf4,
	0x56, 0x01, 0x60, 0x25, 0x80, 0x9d, 0xb4, 0x9b, 0xbb, 0x01, 0xaa, 0x09, 0x2a, 0xcf, 0x40, 0x78,
	0x7e, 0xc8, 0xf1, 0xcd, 0x56, 0x15, 0xa1, 0x9e, 0x17, 0xb7, 0x3b, 0x11, 0x46, 0x98, 0x77, 0x65,
	0x5f, 0x65, 0x4f, 0x84, 0x18, 0x8d, 0x81, 0xfb, 0x49, 0xcc, 0x7d, 0x21, 0x50, 0xfb, 0x3a, 0x46,
	0x71, 0xc1, 0xa8, 0x41, 0x84, 0x20, 0x27, 0xb1, 0xd0, 0x3c, 0x90, 0x69, 0xa2, 0x91, 0x8f, 0x20,
	0x2d, 0x6e, 0x9d, 0x23, 0x42, 0x9f, 0x83, 0x7e, 0x59, 0x88, 0x0d, 0xc4, 0x5b, 0x74, 0x61, 0x4a,
	0x1f, 0x93, 0x1b, 0xa5, 0xbe, 0xe7, 0x87, 0xa1, 0x6c, 0x58, 0x77, 0xac, 0xfd, 0xed, 0x5e, 0xe3,
	0xcb, 0xe7, 0x07, 0x3b, 0xc5, 0xb8, 0xdd, 0x30, 0x94, 0xa0, 0xd4, 0x91, 0x96, 0xb1, 0x88, 0xdc,
	0xeb, 0x25, 0x3c, 0x2b, 0x3b, 0x87, 0xe4, 0xf6, 0xeb, 0xcc, 0x83, 0x92, 0xb6, 0x8f, 0x42, 0x1d,
	0x40, 0xea, 0xc2, 0x74, 0x06, 0x4a, 0x53, 0x4a, 0xb6, 0xd6, 0xa4, 0xae, 0xf9, 0xa6, 0xbb, 0xe4,
	0x6a, 0x70, 0xec, 0xc7, 0xc2, 0x8b, 0xc3, 0x46, 0xdd, 0xd4, 0xaf, 0x98, 0xf3, 0x20, 0x74, 0x7c,
	0xd2, 0xaa, 0x66, 0x53, 0x09, 0x0a, 0x05, 0xb4, 0x4b, 0x48, 0x32, 0x1b, 0x8e, 0xe3, 0xc0, 0x1b,
	0x41, 0x6a, 0x48, 0xaf, 0x75, 0x5a, 0x6c, 0xbd, 0x35, 0xcb, 0xb7, 0x66, 0xaf, 0x0c, 0xe8, 0x00,
	0xd2, 0xde, 0xd6, 0xf9, 0xf7, 0xbd, 0x9a, 0xbb, 0x9d, 0x94, 0x85, 0xce, 0xaf, 0x3a, 0xb9, 0x64,
	0x34, 0xe8, 0x07, 0x8b, 0xdc, 0xdc, 0x30, 0x84, 0xde, 0x67, 0x15, 0x21, 0xb2, 0xbf, 0x6d, 0x6b,
	0xde, 0xad, 0x04, 0xfe, 0x89, 0x72, 0xd8, 0xfb, 0xaf, 0x3f, 0x3f, 0xd5, 0xf7, 0xe9, 0x3d, 0x5e,
	0x06, 0x1d, 0xc2, 0x18, 0x22, 0x93, 0x58, 0x16, 0xf5, 0xa6, 0xf6, 0xc2, 0x22, 0x76, 0xd5, 0xf6,
	0xcf, 0x50, 0xf6, 0x8d, 0x3f, 0x4f, 0xe8, 0xc3, 0x4a, 0xd5, 0x7f, 0x04, 0xd0, 0x6c, 0xff, 0x47,
	0x47, 0x6e, 0xb2, 0x33, 0x30, 0x73, 0xf7, 0x69, 0x97, 0x6f, 0x3e, 0x50, 0x2f, 0xc8, 0x00, 0x42,
	0x6f, 0x4c, 0x5f, 0x10, 0xf0, 0xd3, 0x2c, 0xde, 0x33, 0x7e, 0x5a, 0xa6, 0x7b, 0xd6, 0x3b, 0x3c,
	0x5f, 0xda, 0xd6, 0x62, 0x69, 0x5b, 0x3f, 0x96, 0xb6, 0xf5, 0x71, 0x65, 0xd7, 0x16, 0x2b, 0xbb,
	0xf6, 0x6d, 0x65, 0xd7, 0xde, 0x74, 0xa2, 0x58, 0x1f, 0xcf, 0x86, 0x2c, 0xc0, 0x09, 0x7f, 0x9a,
	0xcb, 0xbc, 0x00, 0xfd, 0x0e, 0xe5, 0xe8, 0x42, 0x75, 0xbe, 0xfe, 0x31, 0x74, 0x9a, 0x80, 0x1a,
	0x5e, 0x36, 0xef, 0xf8, 0xd1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x5d, 0x1f, 0xfe, 0x8a,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// OperatorInfo queries the operator information.
	GetOperatorInfo(ctx context.Context, in *GetOperatorInfoReq, opts ...grpc.CallOption) (*OperatorInfo, error)
	// QueryOperatorConsKeyForChainID queries the consensus public key for the operator
	QueryOperatorConsKeyForChainID(ctx context.Context, in *QueryOperatorConsKeyRequest, opts ...grpc.CallOption) (*QueryOperatorConsKeyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetOperatorInfo(ctx context.Context, in *GetOperatorInfoReq, opts ...grpc.CallOption) (*OperatorInfo, error) {
	out := new(OperatorInfo)
	err := c.cc.Invoke(ctx, "/exocore.operator.v1.Query/GetOperatorInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryOperatorConsKeyForChainID(ctx context.Context, in *QueryOperatorConsKeyRequest, opts ...grpc.CallOption) (*QueryOperatorConsKeyResponse, error) {
	out := new(QueryOperatorConsKeyResponse)
	err := c.cc.Invoke(ctx, "/exocore.operator.v1.Query/QueryOperatorConsKeyForChainID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// OperatorInfo queries the operator information.
	GetOperatorInfo(context.Context, *GetOperatorInfoReq) (*OperatorInfo, error)
	// QueryOperatorConsKeyForChainID queries the consensus public key for the operator
	QueryOperatorConsKeyForChainID(context.Context, *QueryOperatorConsKeyRequest) (*QueryOperatorConsKeyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetOperatorInfo(ctx context.Context, req *GetOperatorInfoReq) (*OperatorInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatorInfo not implemented")
}
func (*UnimplementedQueryServer) QueryOperatorConsKeyForChainID(ctx context.Context, req *QueryOperatorConsKeyRequest) (*QueryOperatorConsKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOperatorConsKeyForChainID not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetOperatorInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatorInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetOperatorInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exocore.operator.v1.Query/GetOperatorInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetOperatorInfo(ctx, req.(*GetOperatorInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryOperatorConsKeyForChainID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOperatorConsKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryOperatorConsKeyForChainID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exocore.operator.v1.Query/QueryOperatorConsKeyForChainID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryOperatorConsKeyForChainID(ctx, req.(*QueryOperatorConsKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exocore.operator.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperatorInfo",
			Handler:    _Query_GetOperatorInfo_Handler,
		},
		{
			MethodName: "QueryOperatorConsKeyForChainID",
			Handler:    _Query_QueryOperatorConsKeyForChainID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exocore/operator/v1/query.proto",
}

func (m *GetOperatorInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetOperatorInfoReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetOperatorInfoReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OperatorAddr) > 0 {
		i -= len(m.OperatorAddr)
		copy(dAtA[i:], m.OperatorAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.OperatorAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryOperatorConsKeyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOperatorConsKeyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOperatorConsKeyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryOperatorConsKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOperatorConsKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOperatorConsKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PublicKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetOperatorInfoReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryOperatorConsKeyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryOperatorConsKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PublicKey.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetOperatorInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: GetOperatorInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetOperatorInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryOperatorConsKeyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryOperatorConsKeyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOperatorConsKeyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryOperatorConsKeyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryOperatorConsKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOperatorConsKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
