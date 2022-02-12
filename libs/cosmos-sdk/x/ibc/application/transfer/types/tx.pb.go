// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/transfer/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	"github.com/okex/exchain/libs/cosmos-sdk/types"
	types1 "github.com/okex/exchain/libs/cosmos-sdk/x/ibc/core/02-client/types"
	//types "github.com/cosmos/cosmos-sdk/types"
	//types1 "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// MsgTransferAdapter defines a msg to transfer fungible tokens (i.e Coins) between
// ICS20 enabled chains. See ICS Spec here:
// https://github.com/cosmos/ics/tree/master/spec/ics-020-fungible-token-transfer#data-structures
type MsgTransferAdapter struct {
	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// the tokens to be transferred
	Token types.CoinAdapter `protobuf:"bytes,3,opt,name=token,proto3" json:"token"`
	// the sender address
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// the recipient address on the destination chain
	Receiver string `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types1.Height `protobuf:"bytes,6,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,7,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
}

func (m *MsgTransferAdapter) Reset()         { *m = MsgTransferAdapter{} }
func (m *MsgTransferAdapter) String() string { return proto.CompactTextString(m) }
func (*MsgTransferAdapter) ProtoMessage()    {}
func (*MsgTransferAdapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7401ed9bed2f8e09, []int{0}
}
func (m *MsgTransferAdapter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferAdapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferAdapter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferAdapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferAdapter.Merge(m, src)
}
func (m *MsgTransferAdapter) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferAdapter) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferAdapter.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferAdapter proto.InternalMessageInfo

// MsgTransferAdapterResponse defines the Msg/Transfer response type.
type MsgTransferAdapterResponse struct {
}

func (m *MsgTransferAdapterResponse) Reset()         { *m = MsgTransferAdapterResponse{} }
func (m *MsgTransferAdapterResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTransferAdapterResponse) ProtoMessage()    {}
func (*MsgTransferAdapterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7401ed9bed2f8e09, []int{1}
}
func (m *MsgTransferAdapterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferAdapterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferAdapterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferAdapterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferAdapterResponse.Merge(m, src)
}
func (m *MsgTransferAdapterResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferAdapterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferAdapterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferAdapterResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgTransferAdapter)(nil), "oec.ibc.applications.transfer.v1.MsgTransferAdapter")
	proto.RegisterType((*MsgTransferAdapterResponse)(nil), "oec.ibc.applications.transfer.v1.MsgTransferAdapterResponse")
}

func init() {
	proto.RegisterFile("ibc/applications/transfer/v1/tx.proto", fileDescriptor_7401ed9bed2f8e09)
}

var fileDescriptor_7401ed9bed2f8e09 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x13, 0xd6, 0x95, 0xe2, 0x6a, 0x13, 0x18, 0x36, 0x65, 0xd5, 0x48, 0xaa, 0x48, 0x48,
	0xe5, 0x80, 0xad, 0x0c, 0x21, 0xa4, 0x1d, 0x10, 0xca, 0x2e, 0x70, 0x98, 0x84, 0xa2, 0x1d, 0x10,
	0x97, 0x91, 0x78, 0x26, 0xb1, 0xd6, 0xd8, 0x91, 0xed, 0x46, 0xdb, 0x37, 0xe0, 0xc8, 0x47, 0xd8,
	0x99, 0x4f, 0xb2, 0xe3, 0x8e, 0x9c, 0x2a, 0xd4, 0x5e, 0x38, 0xf7, 0x13, 0xa0, 0xc4, 0x6e, 0x69,
	0x0f, 0x20, 0x4e, 0xf1, 0x7b, 0xff, 0xdf, 0xf3, 0x5f, 0xcf, 0xef, 0x05, 0x3c, 0x63, 0x19, 0xc1,
	0x69, 0x55, 0x8d, 0x19, 0x49, 0x35, 0x13, 0x5c, 0x61, 0x2d, 0x53, 0xae, 0xbe, 0x50, 0x89, 0xeb,
	0x08, 0xeb, 0x2b, 0x54, 0x49, 0xa1, 0x05, 0x3c, 0x64, 0x19, 0x41, 0xeb, 0x18, 0x5a, 0x62, 0xa8,
	0x8e, 0x06, 0x4f, 0x72, 0x91, 0x8b, 0x16, 0xc4, 0xcd, 0xc9, 0xd4, 0x0c, 0x7c, 0x22, 0x54, 0x29,
	0x14, 0xce, 0x52, 0x45, 0x71, 0x1d, 0x65, 0x54, 0xa7, 0x11, 0x26, 0x82, 0x71, 0xab, 0x07, 0x8d,
	0x35, 0x11, 0x92, 0x62, 0x32, 0x66, 0x94, 0xeb, 0xc6, 0xd0, 0x9c, 0x0c, 0x10, 0x7e, 0xdf, 0x02,
	0xfd, 0x53, 0x95, 0x9f, 0x59, 0x27, 0xf8, 0x1a, 0xf4, 0x95, 0x98, 0x48, 0x42, 0xcf, 0x2b, 0x21,
	0xb5, 0xe7, 0x0e, 0xdd, 0xd1, 0x83, 0x78, 0x7f, 0x31, 0x0d, 0xe0, 0x75, 0x5a, 0x8e, 0x8f, 0xc3,
	0x35, 0x31, 0x4c, 0x80, 0x89, 0x3e, 0x08, 0xa9, 0xe1, 0x5b, 0xb0, 0x6b, 0x35, 0x52, 0xa4, 0x9c,
	0xd3, 0xb1, 0x77, 0xaf, 0xad, 0x3d, 0x58, 0x4c, 0x83, 0xbd, 0x8d, 0x5a, 0xab, 0x87, 0xc9, 0x8e,
	0x49, 0x9c, 0x98, 0x18, 0xbe, 0x02, 0xdb, 0x5a, 0x5c, 0x52, 0xee, 0x6d, 0x0d, 0xdd, 0x51, 0xff,
	0xe8, 0x00, 0x99, 0xde, 0x50, 0xd3, 0x1b, 0xb2, 0xbd, 0xa1, 0x13, 0xc1, 0x78, 0xdc, 0xb9, 0x9d,
	0x06, 0x4e, 0x62, 0x68, 0xb8, 0x0f, 0xba, 0x8a, 0xf2, 0x0b, 0x2a, 0xbd, 0x4e, 0x63, 0x98, 0xd8,
	0x08, 0x0e, 0x40, 0x4f, 0x52, 0x42, 0x59, 0x4d, 0xa5, 0xb7, 0xdd, 0x2a, 0xab, 0x18, 0x7e, 0x06,
	0xbb, 0x9a, 0x95, 0x54, 0x4c, 0xf4, 0x79, 0x41, 0x59, 0x5e, 0x68, 0xaf, 0xdb, 0x7a, 0x0e, 0x50,
	0x33, 0x83, 0xe6, 0xbd, 0x90, 0x7d, 0xa5, 0x3a, 0x42, 0xef, 0x5a, 0x22, 0x7e, 0xda, 0x98, 0xfe,
	0x69, 0x66, 0xb3, 0x3e, 0x4c, 0x76, 0x6c, 0xc2, 0xd0, 0xf0, 0x3d, 0x78, 0xb4, 0x24, 0x9a, 0xaf,
	0xd2, 0x69, 0x59, 0x79, 0xf7, 0x87, 0xee, 0xa8, 0x13, 0x1f, 0x2e, 0xa6, 0x81, 0xb7, 0x79, 0xc9,
	0x0a, 0x09, 0x93, 0x87, 0x36, 0x77, 0xb6, 0x4c, 0x1d, 0xf7, 0xbe, 0xde, 0x04, 0xce, 0xaf, 0x9b,
	0xc0, 0x09, 0xf7, 0xc0, 0xe3, 0xb5, 0x59, 0x25, 0x54, 0x55, 0x82, 0x2b, 0x7a, 0x24, 0xc0, 0xd6,
	0xa9, 0xca, 0x61, 0x01, 0x7a, 0xab, 0x31, 0x3e, 0x47, 0xff, 0x5a, 0x26, 0xb4, 0x76, 0xcb, 0x20,
	0xfa, 0x6f, 0x74, 0x69, 0x18, 0x7f, 0xbc, 0x9d, 0xf9, 0xee, 0xdd, 0xcc, 0x77, 0x7f, 0xce, 0x7c,
	0xf7, 0xdb, 0xdc, 0x77, 0xee, 0xe6, 0xbe, 0xf3, 0x63, 0xee, 0x3b, 0x9f, 0xde, 0xe4, 0x4c, 0x17,
	0x93, 0x0c, 0x11, 0x51, 0x62, 0xbb, 0x9a, 0xe6, 0xf3, 0x42, 0x5d, 0x5c, 0xe2, 0x2b, 0xfc, 0xf7,
	0x3f, 0x41, 0x5f, 0x57, 0x54, 0x65, 0xdd, 0x76, 0x2b, 0x5f, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x26, 0x76, 0x5b, 0xfa, 0x33, 0x03, 0x00, 0x00,
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
	// Transfer defines a rpc handler method for MsgTransferAdapter.
	Transfer(ctx context.Context, in *MsgTransferAdapter, opts ...grpc.CallOption) (*MsgTransferAdapterResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Transfer(ctx context.Context, in *MsgTransferAdapter, opts ...grpc.CallOption) (*MsgTransferAdapterResponse, error) {
	out := new(MsgTransferAdapterResponse)
	err := c.cc.Invoke(ctx, "/oec.ibc.applications.transfer.v1.Msg/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Transfer defines a rpc handler method for MsgTransferAdapter.
	Transfer(context.Context, *MsgTransferAdapter) (*MsgTransferAdapterResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Transfer(ctx context.Context, req *MsgTransferAdapter) (*MsgTransferAdapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferAdapter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oec.ibc.applications.transfer.v1.Msg/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Transfer(ctx, req.(*MsgTransferAdapter))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oec.ibc.applications.transfer.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _Msg_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/transfer/v1/tx.proto",
}

func (m *MsgTransferAdapter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferAdapter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferAdapter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransferAdapterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferAdapterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferAdapterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgTransferAdapter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Token.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	return n
}

func (m *MsgTransferAdapterResponse) Size() (n int) {
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
func (m *MsgTransferAdapter) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgTransferAdapter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferAdapter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
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
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
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
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
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
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
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
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
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
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
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
func (m *MsgTransferAdapterResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgTransferAdapterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferAdapterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
