package types
//
//import (
//	fmt "fmt"
//	_ "github.com/gogo/protobuf/gogoproto"
//	proto "github.com/gogo/protobuf/proto"
//	_ "github.com/gogo/protobuf/types"
//	//github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
//	//crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
//	//version "github.com/tendermint/tendermint/proto/tendermint/version"
//	io "io"
//	//math "math"
//	//math_bits "math/bits"
//	//time "time"
//)
//
//// IBCBlockID
//type IBCBlockID struct {
//	Hash             []byte           `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
//	IBCPartSetHeader IBCPartSetHeader `protobuf:"bytes,2,opt,name=part_set_header,json=PartSetHeader,proto3" json:"part_set_header"`
//}
//
//func (m *IBCBlockID) Reset()         { *m = IBCBlockID{} }
//func (m *IBCBlockID) String() string { return proto.CompactTextString(m) }
//func (*IBCBlockID) ProtoMessage()    {}
//func (*IBCBlockID) Descriptor() ([]byte, []int) {
//	return fileDescriptor_d3a6e55e2345de56, []int{2}
//}
//func (m *IBCBlockID) XXX_Unmarshal(b []byte) error {
//	return m.Unmarshal(b)
//}
//func (m *IBCBlockID) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}
//func (m *IBCBlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
//	if deterministic {
//		return xxx_messageInfo_IBCBlockID.Marshal(b, m, deterministic)
//	} else {
//		b = b[:cap(b)]
//		n, err := m.MarshalToSizedBuffer(b)
//		if err != nil {
//			return nil, err
//		}
//		return b[:n], nil
//	}
//}
//func (m *IBCBlockID) XXX_Merge(src proto.Message) {
//	xxx_messageInfo_IBCBlockID.Merge(m, src)
//}
//func (m *IBCBlockID) XXX_Size() int {
//	return m.Size()
//}
//func (m *IBCBlockID) XXX_DiscardUnknown() {
//	xxx_messageInfo_IBCBlockID.DiscardUnknown(m)
//}
//
//var xxx_messageInfo_IBCBlockID proto.InternalMessageInfo
//
//var fileDescriptor_d3a6e55e2345de56 = []byte{
//	// 1314 bytes of a gzipped FileDescriptorProto
//	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x4f, 0x6f, 0x1b, 0x45,
//	0x14, 0xcf, 0xda, 0x9b, 0xd8, 0x7e, 0xb6, 0x13, 0x67, 0x95, 0xb6, 0xae, 0xdb, 0x38, 0x2b, 0x23,
//	0x20, 0x2d, 0x68, 0x53, 0x52, 0xc4, 0x9f, 0x03, 0x07, 0xdb, 0x49, 0x5b, 0xab, 0x89, 0x63, 0xd6,
//	0x6e, 0x11, 0x5c, 0x56, 0x6b, 0xef, 0xd4, 0x5e, 0xba, 0xde, 0x59, 0xed, 0x8c, 0x43, 0xd2, 0x4f,
//	0x80, 0x72, 0xea, 0x89, 0x5b, 0x4e, 0x70, 0xe0, 0xce, 0x17, 0x40, 0x9c, 0x7a, 0xec, 0x0d, 0x2e,
//	0x14, 0x94, 0x4a, 0x88, 0x8f, 0x81, 0xe6, 0x8f, 0xd7, 0xeb, 0x38, 0x86, 0xaa, 0xaa, 0xb8, 0x58,
//	0x3b, 0xef, 0xfd, 0xde, 0xcc, 0x7b, 0xbf, 0xf7, 0x9b, 0x3f, 0x86, 0xeb, 0x14, 0xf9, 0x0e, 0x0a,
//	0x87, 0xae, 0x4f, 0xb7, 0xe8, 0x71, 0x80, 0x88, 0xf8, 0x35, 0x82, 0x10, 0x53, 0xac, 0x15, 0x26,
//	0x5e, 0x83, 0xdb, 0x4b, 0x6b, 0x7d, 0xdc, 0xc7, 0xdc, 0xb9, 0xc5, 0xbe, 0x04, 0xae, 0xb4, 0xd1,
//	0xc7, 0xb8, 0xef, 0xa1, 0x2d, 0x3e, 0xea, 0x8e, 0x1e, 0x6d, 0x51, 0x77, 0x88, 0x08, 0xb5, 0x87,
//	0x81, 0x04, 0xac, 0xc7, 0x96, 0xe9, 0x85, 0xc7, 0x01, 0xc5, 0x0c, 0x8b, 0x1f, 0x49, 0x77, 0x39,
//	0xe6, 0x3e, 0x44, 0x21, 0x71, 0xb1, 0x1f, 0xcf, 0xa3, 0xa4, 0xcf, 0x64, 0x79, 0x68, 0x7b, 0xae,
//	0x63, 0x53, 0x1c, 0x0a, 0x44, 0xe5, 0x53, 0xc8, 0xb7, 0xec, 0x90, 0xb6, 0x11, 0xbd, 0x87, 0x6c,
//	0x07, 0x85, 0xda, 0x1a, 0x2c, 0x52, 0x4c, 0x6d, 0xaf, 0xa8, 0xe8, 0xca, 0x66, 0xde, 0x14, 0x03,
//	0x4d, 0x03, 0x75, 0x60, 0x93, 0x41, 0x31, 0xa1, 0x2b, 0x9b, 0x39, 0x93, 0x7f, 0x57, 0x06, 0xa0,
//	0xb2, 0x50, 0x16, 0xe1, 0xfa, 0x0e, 0x3a, 0x1a, 0x47, 0xf0, 0x01, 0xb3, 0x76, 0x8f, 0x29, 0x22,
//	0x32, 0x44, 0x0c, 0xb4, 0x0f, 0x61, 0x91, 0xe7, 0x5f, 0x4c, 0xea, 0xca, 0x66, 0x76, 0xbb, 0x68,
//	0xc4, 0x88, 0x12, 0xf5, 0x19, 0x2d, 0xe6, 0xaf, 0xa9, 0xcf, 0x5e, 0x6c, 0x2c, 0x98, 0x02, 0x5c,
//	0xf1, 0x20, 0x55, 0xf3, 0x70, 0xef, 0x71, 0x63, 0x27, 0x4a, 0x44, 0x99, 0x24, 0xa2, 0xed, 0xc3,
//	0x4a, 0x60, 0x87, 0xd4, 0x22, 0x88, 0x5a, 0x03, 0x5e, 0x05, 0x5f, 0x34, 0xbb, 0xbd, 0x61, 0x9c,
//	0xef, 0x83, 0x31, 0x55, 0xac, 0x5c, 0x25, 0x1f, 0xc4, 0x8d, 0x95, 0xbf, 0x54, 0x58, 0x92, 0x64,
//	0x7c, 0x06, 0x29, 0x49, 0x2b, 0x5f, 0x30, 0xbb, 0xbd, 0x1e, 0x9f, 0x51, 0xba, 0x8c, 0x3a, 0xf6,
//	0x09, 0xf2, 0xc9, 0x88, 0xc8, 0xf9, 0xc6, 0x31, 0xda, 0x3b, 0x90, 0xee, 0x0d, 0x6c, 0xd7, 0xb7,
//	0x5c, 0x87, 0x67, 0x94, 0xa9, 0x65, 0xcf, 0x5e, 0x6c, 0xa4, 0xea, 0xcc, 0xd6, 0xd8, 0x31, 0x53,
//	0xdc, 0xd9, 0x70, 0xb4, 0xcb, 0xb0, 0x34, 0x40, 0x6e, 0x7f, 0x40, 0x39, 0x2d, 0x49, 0x53, 0x8e,
//	0xb4, 0x4f, 0x40, 0x65, 0x82, 0x28, 0xaa, 0x7c, 0xed, 0x92, 0x21, 0xd4, 0x62, 0x8c, 0xd5, 0x62,
//	0x74, 0xc6, 0x6a, 0xa9, 0xa5, 0xd9, 0xc2, 0x4f, 0xff, 0xd8, 0x50, 0x4c, 0x1e, 0xa1, 0xd5, 0x21,
//	0xef, 0xd9, 0x84, 0x5a, 0x5d, 0x46, 0x1b, 0x5b, 0x7e, 0x91, 0x4f, 0x71, 0x75, 0x96, 0x10, 0x49,
//	0xac, 0x4c, 0x3d, 0xcb, 0xa2, 0x84, 0xc9, 0xd1, 0x36, 0xa1, 0xc0, 0x27, 0xe9, 0xe1, 0xe1, 0xd0,
//	0xa5, 0x16, 0xe7, 0x7d, 0x89, 0xf3, 0xbe, 0xcc, 0xec, 0x75, 0x6e, 0xbe, 0xc7, 0x3a, 0x70, 0x0d,
//	0x32, 0x8e, 0x4d, 0x6d, 0x01, 0x49, 0x71, 0x48, 0x9a, 0x19, 0xb8, 0xf3, 0x5d, 0x58, 0x89, 0x54,
//	0x47, 0x04, 0x24, 0x2d, 0x66, 0x99, 0x98, 0x39, 0xf0, 0x16, 0xac, 0xf9, 0xe8, 0x88, 0x5a, 0xe7,
//	0xd1, 0x19, 0x8e, 0xd6, 0x98, 0xef, 0xe1, 0x74, 0xc4, 0xdb, 0xb0, 0xdc, 0x1b, 0x93, 0x2f, 0xb0,
//	0xc0, 0xb1, 0xf9, 0xc8, 0xca, 0x61, 0x57, 0x21, 0x6d, 0x07, 0x81, 0x00, 0x64, 0x39, 0x20, 0x65,
//	0x07, 0x01, 0x77, 0xdd, 0x84, 0x55, 0x5e, 0x63, 0x88, 0xc8, 0xc8, 0xa3, 0x72, 0x92, 0x1c, 0xc7,
//	0xac, 0x30, 0x87, 0x29, 0xec, 0x1c, 0xfb, 0x16, 0xe4, 0xd1, 0xa1, 0xeb, 0x20, 0xbf, 0x87, 0x04,
//	0x2e, 0xcf, 0x71, 0xb9, 0xb1, 0x91, 0x83, 0x6e, 0x40, 0x21, 0x08, 0x71, 0x80, 0x09, 0x0a, 0x2d,
//	0xdb, 0x71, 0x42, 0x44, 0x48, 0x71, 0x59, 0xcc, 0x37, 0xb6, 0x57, 0x85, 0xb9, 0x52, 0x04, 0x75,
//	0xc7, 0xa6, 0xb6, 0x56, 0x80, 0x24, 0x3d, 0x22, 0x45, 0x45, 0x4f, 0x6e, 0xe6, 0x4c, 0xf6, 0x59,
//	0xf9, 0x3b, 0x01, 0xea, 0x43, 0x4c, 0x91, 0x76, 0x1b, 0x54, 0xd6, 0x26, 0xae, 0xbe, 0xe5, 0x8b,
//	0xf4, 0xdc, 0x76, 0xfb, 0x3e, 0x72, 0xf6, 0x49, 0xbf, 0x73, 0x1c, 0x20, 0x93, 0x83, 0x63, 0x72,
//	0x4a, 0x4c, 0xc9, 0x69, 0x0d, 0x16, 0x43, 0x3c, 0xf2, 0x1d, 0xae, 0xb2, 0x45, 0x53, 0x0c, 0xb4,
//	0x5d, 0x48, 0x47, 0x2a, 0x51, 0xff, 0x4b, 0x25, 0x2b, 0x4c, 0x25, 0x4c, 0xc3, 0xd2, 0x60, 0xa6,
//	0xba, 0x52, 0x2c, 0x35, 0xc8, 0x44, 0x87, 0x97, 0x54, 0xdb, 0xab, 0x09, 0x76, 0x12, 0xa6, 0xbd,
//	0x07, 0xab, 0x51, 0xef, 0x23, 0xf2, 0x84, 0xe2, 0x0a, 0x91, 0x43, 0xb2, 0x37, 0x25, 0x2b, 0x4b,
//	0x1c, 0x40, 0x29, 0x5e, 0xd7, 0x44, 0x56, 0x0d, 0x7e, 0x12, 0x5d, 0x87, 0x0c, 0x71, 0xfb, 0xbe,
//	0x4d, 0x47, 0x21, 0x92, 0xca, 0x9b, 0x18, 0x2a, 0x3f, 0x2b, 0xb0, 0x24, 0x94, 0x1c, 0xe3, 0x4d,
//	0xb9, 0x98, 0xb7, 0xc4, 0x3c, 0xde, 0x92, 0xaf, 0xcf, 0x5b, 0x15, 0x20, 0x4a, 0x86, 0x14, 0x55,
//	0x3d, 0xb9, 0x99, 0xdd, 0xbe, 0x36, 0x3b, 0x91, 0x48, 0xb1, 0xed, 0xf6, 0xe5, 0x46, 0x8d, 0x05,
//	0x55, 0x7e, 0x57, 0x20, 0x13, 0xf9, 0xb5, 0x2a, 0xe4, 0xc7, 0x79, 0x59, 0x8f, 0x3c, 0xbb, 0x2f,
//	0xb5, 0xb3, 0x3e, 0x37, 0xb9, 0x3b, 0x9e, 0xdd, 0x37, 0xb3, 0x32, 0x1f, 0x36, 0xb8, 0xb8, 0x0f,
//	0x89, 0x39, 0x7d, 0x98, 0x6a, 0x7c, 0xf2, 0xf5, 0x1a, 0x3f, 0xd5, 0x22, 0xf5, 0x7c, 0x8b, 0x7e,
//	0x4a, 0x40, 0xba, 0xc5, 0xf7, 0x8e, 0xed, 0xfd, 0x1f, 0x3b, 0xe2, 0x1a, 0x64, 0x02, 0xec, 0x59,
//	0xc2, 0xa3, 0x72, 0x4f, 0x3a, 0xc0, 0x9e, 0x39, 0xd3, 0xf6, 0xc5, 0x37, 0xb4, 0x5d, 0x96, 0xde,
//	0x00, 0x6b, 0xa9, 0xf3, 0xac, 0x85, 0x90, 0x13, 0x54, 0xc8, 0xbb, 0xec, 0x16, 0xe3, 0x80, 0x5f,
//	0x8e, 0xca, 0xec, 0xdd, 0x2b, 0xd2, 0x16, 0x48, 0x53, 0xe2, 0x58, 0x84, 0x38, 0xfa, 0xe5, 0x75,
//	0x5a, 0x9c, 0x27, 0x4b, 0x53, 0xe2, 0x2a, 0xdf, 0x29, 0x00, 0x7b, 0x8c, 0x59, 0x5e, 0x2f, 0xbb,
//	0x85, 0x08, 0x4f, 0xc1, 0x9a, 0x5a, 0xb9, 0x3c, 0xaf, 0x69, 0x72, 0xfd, 0x1c, 0x89, 0xe7, 0x5d,
//	0x87, 0xfc, 0x44, 0x8c, 0x04, 0x8d, 0x93, 0xb9, 0x60, 0x92, 0xe8, 0x72, 0x68, 0x23, 0x6a, 0xe6,
//	0x0e, 0x63, 0xa3, 0xca, 0x2f, 0x0a, 0x64, 0x78, 0x4e, 0xfb, 0x88, 0xda, 0x53, 0x3d, 0x54, 0x5e,
//	0xbf, 0x87, 0xeb, 0x00, 0x62, 0x1a, 0xe2, 0x3e, 0x41, 0x52, 0x59, 0x19, 0x6e, 0x69, 0xbb, 0x4f,
//	0x90, 0xf6, 0x51, 0x44, 0x78, 0xf2, 0xdf, 0x09, 0x97, 0x5b, 0x7a, 0x4c, 0xfb, 0x15, 0x48, 0xf9,
//	0xa3, 0xa1, 0xc5, 0xae, 0x04, 0x55, 0xa8, 0xd5, 0x1f, 0x0d, 0x3b, 0x47, 0xa4, 0xf2, 0x35, 0xa4,
//	0x3a, 0x47, 0xfc, 0x79, 0xc4, 0x24, 0x1a, 0x62, 0x2c, 0xef, 0x64, 0xf1, 0x16, 0x4a, 0x33, 0x03,
//	0xbf, 0x82, 0x34, 0x50, 0xd9, 0xe5, 0x3b, 0x7e, 0xac, 0xb1, 0x6f, 0xcd, 0x78, 0xc5, 0x87, 0x97,
//	0x7c, 0x72, 0xdd, 0xfc, 0x55, 0x81, 0x6c, 0xec, 0x7c, 0xd0, 0x3e, 0x80, 0x4b, 0xb5, 0xbd, 0x83,
//	0xfa, 0x7d, 0xab, 0xb1, 0x63, 0xdd, 0xd9, 0xab, 0xde, 0xb5, 0x1e, 0x34, 0xef, 0x37, 0x0f, 0xbe,
//	0x68, 0x16, 0x16, 0x4a, 0x97, 0x4f, 0x4e, 0x75, 0x2d, 0x86, 0x7d, 0xe0, 0x3f, 0xf6, 0xf1, 0x37,
//	0xbe, 0xb6, 0x05, 0x6b, 0xd3, 0x21, 0xd5, 0x5a, 0x7b, 0xb7, 0xd9, 0x29, 0x28, 0xa5, 0x4b, 0x27,
//	0xa7, 0xfa, 0x6a, 0x2c, 0xa2, 0xda, 0x25, 0xc8, 0xa7, 0xb3, 0x01, 0xf5, 0x83, 0xfd, 0xfd, 0x46,
//	0xa7, 0x90, 0x98, 0x09, 0x90, 0x07, 0xf6, 0x0d, 0x58, 0x9d, 0x0e, 0x68, 0x36, 0xf6, 0x0a, 0xc9,
//	0x92, 0x76, 0x72, 0xaa, 0x2f, 0xc7, 0xd0, 0x4d, 0xd7, 0x2b, 0xa5, 0xbf, 0xfd, 0xbe, 0xbc, 0xf0,
//	0xe3, 0x0f, 0x65, 0x85, 0x55, 0x96, 0x9f, 0x3a, 0x23, 0xb4, 0xf7, 0xe1, 0x4a, 0xbb, 0x71, 0xb7,
//	0xb9, 0xbb, 0x63, 0xed, 0xb7, 0xef, 0x5a, 0x9d, 0x2f, 0x5b, 0xbb, 0xb1, 0xea, 0x56, 0x4e, 0x4e,
//	0xf5, 0xac, 0x2c, 0x69, 0x1e, 0xba, 0x65, 0xee, 0x3e, 0x3c, 0xe8, 0xec, 0x16, 0x14, 0x81, 0x6e,
//	0x85, 0xe8, 0x10, 0x53, 0xc4, 0xd1, 0xb7, 0xe0, 0xea, 0x05, 0xe8, 0xa8, 0xb0, 0xd5, 0x93, 0x53,
//	0x3d, 0xdf, 0x0a, 0x91, 0xd8, 0x3f, 0x3c, 0xc2, 0x80, 0xe2, 0x6c, 0xc4, 0x41, 0xeb, 0xa0, 0x5d,
//	0xdd, 0x2b, 0xe8, 0xa5, 0xc2, 0xc9, 0xa9, 0x9e, 0x1b, 0x1f, 0x86, 0x0c, 0x3f, 0xa9, 0xac, 0xf6,
//	0xf9, 0xb3, 0xb3, 0xb2, 0xf2, 0xfc, 0xac, 0xac, 0xfc, 0x79, 0x56, 0x56, 0x9e, 0xbe, 0x2c, 0x2f,
//	0x3c, 0x7f, 0x59, 0x5e, 0xf8, 0xed, 0x65, 0x79, 0xe1, 0xab, 0x8f, 0xfb, 0x2e, 0x1d, 0x8c, 0xba,
//	0x46, 0x0f, 0x0f, 0xb7, 0xe2, 0x7f, 0x09, 0x26, 0x9f, 0xe2, 0xaf, 0xc9, 0xf9, 0xbf, 0x0b, 0xdd,
//	0x25, 0x6e, 0xbf, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x78, 0x43, 0xdf, 0xef, 0x0c,
//	0x00, 0x00,
//}
//
//func (m *IBCBlockID) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowTypes
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: BlockID: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: BlockID: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
//			}
//			var byteLen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowTypes
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				byteLen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if byteLen < 0 {
//				return ErrInvalidLengthTypes
//			}
//			postIndex := iNdEx + byteLen
//			if postIndex < 0 {
//				return ErrInvalidLengthTypes
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
//			if m.Hash == nil {
//				m.Hash = []byte{}
//			}
//			iNdEx = postIndex
//		case 2:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field IBCPartSetHeader", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowTypes
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthTypes
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthTypes
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if err := m.IBCPartSetHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		default:
//			iNdEx = preIndex
//			skippy, err := skipTypes(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthTypes
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
//func (m *IBCBlockID) Size() (n int) {
//	if m == nil {
//		return 0
//	}
//	var l int
//	_ = l
//	l = len(m.Hash)
//	if l > 0 {
//		n += 1 + l + sovTypes(uint64(l))
//	}
//	l = m.IBCPartSetHeader.Size()
//	n += 1 + l + sovTypes(uint64(l))
//	return n
//}
//func (m *IBCBlockID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
//	i := len(dAtA)
//	_ = i
//	var l int
//	_ = l
//	{
//		size, err := m.IBCPartSetHeader.MarshalToSizedBuffer(dAtA[:i])
//		if err != nil {
//			return 0, err
//		}
//		i -= size
//		i = encodeVarintTypes(dAtA, i, uint64(size))
//	}
//	i--
//	dAtA[i] = 0x12
//	if len(m.Hash) > 0 {
//		i -= len(m.Hash)
//		copy(dAtA[i:], m.Hash)
//		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
//		i--
//		dAtA[i] = 0xa
//	}
//	return len(dAtA) - i, nil
//}
//
/////
//// IBCPartSetHeader
//type IBCPartSetHeader struct {
//	Total uint32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
//	Hash  []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
//}
//
//func (m *IBCPartSetHeader) Reset()         { *m = IBCPartSetHeader{} }
//func (m *IBCPartSetHeader) String() string { return proto.CompactTextString(m) }
//func (*IBCPartSetHeader) ProtoMessage()    {}
//func (*IBCPartSetHeader) Descriptor() ([]byte, []int) {
//	return fileDescriptor_d3a6e55e2345de56, []int{0}
//}
//func (m *IBCPartSetHeader) XXX_Unmarshal(b []byte) error {
//	return m.Unmarshal(b)
//}
//func (m *IBCPartSetHeader) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowTypes
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: PartSetHeader: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: PartSetHeader: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
//			}
//			m.Total = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowTypes
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Total |= uint32(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 2:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
//			}
//			var byteLen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowTypes
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				byteLen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if byteLen < 0 {
//				return ErrInvalidLengthTypes
//			}
//			postIndex := iNdEx + byteLen
//			if postIndex < 0 {
//				return ErrInvalidLengthTypes
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
//			if m.Hash == nil {
//				m.Hash = []byte{}
//			}
//			iNdEx = postIndex
//		default:
//			iNdEx = preIndex
//			skippy, err := skipTypes(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthTypes
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
//func (m *IBCPartSetHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
//	if deterministic {
//		return xxx_messageInfo_IBCPartSetHeader.Marshal(b, m, deterministic)
//	} else {
//		b = b[:cap(b)]
//		n, err := m.MarshalToSizedBuffer(b)
//		if err != nil {
//			return nil, err
//		}
//		return b[:n], nil
//	}
//}
//
//var xxx_messageInfo_IBCPartSetHeader proto.InternalMessageInfo
//
//func (m *IBCPartSetHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
//	i := len(dAtA)
//	_ = i
//	var l int
//	_ = l
//	if len(m.Hash) > 0 {
//		i -= len(m.Hash)
//		copy(dAtA[i:], m.Hash)
//		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
//		i--
//		dAtA[i] = 0x12
//	}
//	if m.Total != 0 {
//		i = encodeVarintTypes(dAtA, i, uint64(m.Total))
//		i--
//		dAtA[i] = 0x8
//	}
//	return len(dAtA) - i, nil
//}
//func (m *IBCPartSetHeader) XXX_Merge(src proto.Message) {
//	xxx_messageInfo_IBCPartSetHeader.Merge(m, src)
//}
//func (m *IBCPartSetHeader) Size() (n int) {
//	if m == nil {
//		return 0
//	}
//	var l int
//	_ = l
//	if m.Total != 0 {
//		n += 1 + sovTypes(uint64(m.Total))
//	}
//	l = len(m.Hash)
//	if l > 0 {
//		n += 1 + l + sovTypes(uint64(l))
//	}
//	return n
//}
//func (m *IBCPartSetHeader) XXX_Size() int {
//	return m.Size()
//}
//func (m *IBCPartSetHeader) XXX_DiscardUnknown() {
//	xxx_messageInfo_IBCPartSetHeader.DiscardUnknown(m)
//}
