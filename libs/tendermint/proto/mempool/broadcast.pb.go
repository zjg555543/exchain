// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: libs/tendermint/proto/mempool/broadcast.proto

package mempool

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/emptypb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TxRequest struct {
	Tx                   []byte   `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Txs                  [][]byte `protobuf:"bytes,3,rep,name=txs,proto3" json:"txs,omitempty"`
	FromAddrs            []string `protobuf:"bytes,4,rep,name=from_addrs,json=fromAddrs,proto3" json:"from_addrs,omitempty"`
	PeerId               uint32   `protobuf:"varint,5,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxRequest) Reset()         { *m = TxRequest{} }
func (m *TxRequest) String() string { return proto.CompactTextString(m) }
func (*TxRequest) ProtoMessage()    {}
func (*TxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1597b6564598034d, []int{0}
}
func (m *TxRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxRequest.Merge(m, src)
}
func (m *TxRequest) XXX_Size() int {
	return m.Size()
}
func (m *TxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TxRequest proto.InternalMessageInfo

func (m *TxRequest) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TxRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TxRequest) GetTxs() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *TxRequest) GetFromAddrs() []string {
	if m != nil {
		return m.FromAddrs
	}
	return nil
}

func (m *TxRequest) GetPeerId() uint32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

type ReceiverInfo struct {
	YourId               uint32   `protobuf:"varint,1,opt,name=your_id,json=yourId,proto3" json:"your_id,omitempty"`
	Port                 int64    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiverInfo) Reset()         { *m = ReceiverInfo{} }
func (m *ReceiverInfo) String() string { return proto.CompactTextString(m) }
func (*ReceiverInfo) ProtoMessage()    {}
func (*ReceiverInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1597b6564598034d, []int{1}
}
func (m *ReceiverInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReceiverInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReceiverInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReceiverInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiverInfo.Merge(m, src)
}
func (m *ReceiverInfo) XXX_Size() int {
	return m.Size()
}
func (m *ReceiverInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiverInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiverInfo proto.InternalMessageInfo

func (m *ReceiverInfo) GetYourId() uint32 {
	if m != nil {
		return m.YourId
	}
	return 0
}

func (m *ReceiverInfo) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*TxRequest)(nil), "mempool.TxRequest")
	proto.RegisterType((*ReceiverInfo)(nil), "mempool.ReceiverInfo")
}

func init() {
	proto.RegisterFile("libs/tendermint/proto/mempool/broadcast.proto", fileDescriptor_1597b6564598034d)
}

var fileDescriptor_1597b6564598034d = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x31, 0x4f, 0xf3, 0x30,
	0x14, 0xac, 0x9b, 0x7e, 0xed, 0x97, 0xa7, 0x80, 0xc0, 0x03, 0x44, 0x45, 0x8d, 0xa2, 0x4e, 0x59,
	0x48, 0x24, 0x18, 0x90, 0x80, 0xa5, 0x20, 0x86, 0x0e, 0x2c, 0x56, 0x27, 0x96, 0xaa, 0x89, 0x5f,
	0x4b, 0x44, 0x13, 0x07, 0xdb, 0x45, 0xc9, 0x3f, 0xe1, 0x17, 0x21, 0x46, 0x7e, 0x02, 0x2a, 0x7f,
	0x04, 0x39, 0x09, 0x8c, 0x48, 0xdd, 0xde, 0xbb, 0x77, 0x67, 0xdd, 0x9d, 0xe1, 0x74, 0x9d, 0xc6,
	0x2a, 0xd2, 0x98, 0x73, 0x94, 0x59, 0x9a, 0xeb, 0xa8, 0x90, 0x42, 0x8b, 0x28, 0xc3, 0xac, 0x10,
	0x62, 0x1d, 0xc5, 0x52, 0x2c, 0x78, 0xb2, 0x50, 0x3a, 0xac, 0x71, 0x3a, 0x68, 0x0f, 0xc3, 0x93,
	0x95, 0x10, 0xab, 0x35, 0x36, 0xf4, 0x78, 0xb3, 0x8c, 0x30, 0x2b, 0x74, 0xd5, 0xb0, 0xc6, 0x15,
	0xd8, 0xb3, 0x92, 0xe1, 0xf3, 0x06, 0x95, 0xa6, 0xfb, 0xd0, 0xd5, 0xa5, 0x4b, 0x7c, 0x12, 0x38,
	0xac, 0xab, 0x4b, 0x4a, 0xa1, 0xb7, 0x94, 0x22, 0x73, 0xbb, 0x3e, 0x09, 0x6c, 0x56, 0xcf, 0xf4,
	0x00, 0x2c, 0x5d, 0x2a, 0xd7, 0xf2, 0xad, 0xc0, 0x61, 0x66, 0xa4, 0x23, 0x00, 0x73, 0x99, 0x2f,
	0x38, 0x97, 0xca, 0xed, 0xf9, 0x56, 0x60, 0x33, 0xdb, 0x20, 0x13, 0x03, 0xd0, 0x63, 0x18, 0x14,
	0x88, 0x72, 0x9e, 0x72, 0xf7, 0x9f, 0x4f, 0x82, 0x3d, 0xd6, 0x37, 0xeb, 0x94, 0x8f, 0xaf, 0xc0,
	0x61, 0x98, 0x60, 0xfa, 0x82, 0x72, 0x9a, 0x2f, 0x85, 0x21, 0x56, 0x62, 0x53, 0x13, 0x49, 0x43,
	0x34, 0xeb, 0x94, 0x1b, 0x1b, 0x85, 0x90, 0xba, 0xb6, 0x61, 0xb1, 0x7a, 0x3e, 0x7b, 0x23, 0x70,
	0x78, 0xdf, 0x04, 0x34, 0xfe, 0x9b, 0x67, 0xe8, 0x05, 0x0c, 0x6e, 0x1f, 0x31, 0x79, 0x9a, 0x95,
	0x94, 0x86, 0x6d, 0xfe, 0xf0, 0x37, 0xdf, 0xf0, 0x28, 0x6c, 0xaa, 0x08, 0x7f, 0xaa, 0x08, 0xef,
	0x4c, 0x15, 0xe3, 0x0e, 0xbd, 0x06, 0xa7, 0x15, 0x4e, 0x54, 0x95, 0x27, 0x3b, 0xaa, 0x2f, 0xe1,
	0x7f, 0xab, 0x56, 0xbb, 0x29, 0x03, 0x72, 0x13, 0xbd, 0x6f, 0x3d, 0xf2, 0xb1, 0xf5, 0xc8, 0xe7,
	0xd6, 0x23, 0xaf, 0x5f, 0x5e, 0xe7, 0x61, 0xf4, 0xe7, 0x3f, 0xc7, 0xfd, 0x7a, 0x3d, 0xff, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xfa, 0xac, 0xb3, 0x5b, 0x0f, 0x02, 0x00, 0x00,
}

func (m *TxRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PeerId != 0 {
		i = encodeVarintBroadcast(dAtA, i, uint64(m.PeerId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.FromAddrs) > 0 {
		for iNdEx := len(m.FromAddrs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FromAddrs[iNdEx])
			copy(dAtA[i:], m.FromAddrs[iNdEx])
			i = encodeVarintBroadcast(dAtA, i, uint64(len(m.FromAddrs[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Txs[iNdEx])
			copy(dAtA[i:], m.Txs[iNdEx])
			i = encodeVarintBroadcast(dAtA, i, uint64(len(m.Txs[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintBroadcast(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = encodeVarintBroadcast(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReceiverInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReceiverInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReceiverInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Port != 0 {
		i = encodeVarintBroadcast(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x10
	}
	if m.YourId != 0 {
		i = encodeVarintBroadcast(dAtA, i, uint64(m.YourId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBroadcast(dAtA []byte, offset int, v uint64) int {
	offset -= sovBroadcast(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovBroadcast(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovBroadcast(uint64(l))
	}
	if len(m.Txs) > 0 {
		for _, b := range m.Txs {
			l = len(b)
			n += 1 + l + sovBroadcast(uint64(l))
		}
	}
	if len(m.FromAddrs) > 0 {
		for _, s := range m.FromAddrs {
			l = len(s)
			n += 1 + l + sovBroadcast(uint64(l))
		}
	}
	if m.PeerId != 0 {
		n += 1 + sovBroadcast(uint64(m.PeerId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ReceiverInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.YourId != 0 {
		n += 1 + sovBroadcast(uint64(m.YourId))
	}
	if m.Port != 0 {
		n += 1 + sovBroadcast(uint64(m.Port))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBroadcast(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBroadcast(x uint64) (n int) {
	return sovBroadcast(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBroadcast
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
			return fmt.Errorf("proto: TxRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroadcast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBroadcast
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBroadcast
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tx = append(m.Tx[:0], dAtA[iNdEx:postIndex]...)
			if m.Tx == nil {
				m.Tx = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroadcast
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
				return ErrInvalidLengthBroadcast
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroadcast
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroadcast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBroadcast
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBroadcast
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, make([]byte, postIndex-iNdEx))
			copy(m.Txs[len(m.Txs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddrs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroadcast
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
				return ErrInvalidLengthBroadcast
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBroadcast
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddrs = append(m.FromAddrs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			m.PeerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroadcast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBroadcast(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBroadcast
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReceiverInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBroadcast
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
			return fmt.Errorf("proto: ReceiverInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReceiverInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field YourId", wireType)
			}
			m.YourId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroadcast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.YourId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBroadcast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBroadcast(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBroadcast
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBroadcast(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBroadcast
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
					return 0, ErrIntOverflowBroadcast
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
					return 0, ErrIntOverflowBroadcast
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
				return 0, ErrInvalidLengthBroadcast
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBroadcast
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBroadcast
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBroadcast        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBroadcast          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBroadcast = fmt.Errorf("proto: unexpected end of group")
)
