// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cyber/rank/v1beta1/pagination.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PageRequest struct {
	Page    uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage uint32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (m *PageRequest) Reset()         { *m = PageRequest{} }
func (m *PageRequest) String() string { return proto.CompactTextString(m) }
func (*PageRequest) ProtoMessage()    {}
func (*PageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6472f2c93262fe3, []int{0}
}

func (m *PageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *PageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *PageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageRequest.Merge(m, src)
}

func (m *PageRequest) XXX_Size() int {
	return m.Size()
}

func (m *PageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PageRequest proto.InternalMessageInfo

func (m *PageRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PageRequest) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

type PageResponse struct {
	Total uint32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (m *PageResponse) Reset()         { *m = PageResponse{} }
func (m *PageResponse) String() string { return proto.CompactTextString(m) }
func (*PageResponse) ProtoMessage()    {}
func (*PageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6472f2c93262fe3, []int{1}
}

func (m *PageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *PageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *PageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageResponse.Merge(m, src)
}

func (m *PageResponse) XXX_Size() int {
	return m.Size()
}

func (m *PageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PageResponse proto.InternalMessageInfo

func (m *PageResponse) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*PageRequest)(nil), "cyber.rank.v1beta1.PageRequest")
	proto.RegisterType((*PageResponse)(nil), "cyber.rank.v1beta1.PageResponse")
}

func init() {
	proto.RegisterFile("cyber/rank/v1beta1/pagination.proto", fileDescriptor_c6472f2c93262fe3)
}

var fileDescriptor_c6472f2c93262fe3 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8f, 0xbd, 0x4a, 0xc0, 0x30,
	0x14, 0x46, 0x1b, 0xf1, 0x8f, 0xa8, 0x4b, 0x70, 0xd0, 0x25, 0x48, 0x75, 0x70, 0x31, 0xa1, 0xb8,
	0x3a, 0x39, 0x39, 0x4a, 0x47, 0x17, 0x49, 0xca, 0x25, 0x16, 0x35, 0x37, 0x26, 0xb7, 0x62, 0xdf,
	0xc2, 0xc7, 0x72, 0xec, 0xe8, 0x28, 0xed, 0x8b, 0x48, 0x93, 0x6e, 0xc9, 0xc7, 0x39, 0x70, 0x0f,
	0xbf, 0xec, 0x46, 0x0b, 0x51, 0x47, 0xe3, 0x5f, 0xf5, 0x67, 0x63, 0x81, 0x4c, 0xa3, 0x83, 0x71,
	0xbd, 0x37, 0xd4, 0xa3, 0x57, 0x21, 0x22, 0xa1, 0x10, 0x19, 0x52, 0x2b, 0xa4, 0x36, 0xa8, 0xbe,
	0xe3, 0x47, 0x8f, 0xc6, 0x41, 0x0b, 0x1f, 0x03, 0x24, 0x12, 0x82, 0xef, 0x06, 0xe3, 0xe0, 0x8c,
	0x5d, 0xb0, 0xeb, 0x93, 0x36, 0xbf, 0xc5, 0x39, 0x3f, 0x0c, 0x10, 0x9f, 0xf3, 0xbe, 0x93, 0xf7,
	0x83, 0x00, 0x71, 0xb5, 0xea, 0x2b, 0x7e, 0x5c, 0xec, 0x14, 0xd0, 0x27, 0x10, 0xa7, 0x7c, 0x8f,
	0x90, 0xcc, 0xdb, 0xe6, 0x97, 0xcf, 0xfd, 0xc3, 0xcf, 0x2c, 0xd9, 0x34, 0x4b, 0xf6, 0x37, 0x4b,
	0xf6, 0xbd, 0xc8, 0x6a, 0x5a, 0x64, 0xf5, 0xbb, 0xc8, 0xea, 0x49, 0xb9, 0x9e, 0x5e, 0x06, 0xab,
	0x3a, 0x7c, 0xd7, 0xf9, 0xb8, 0x0e, 0xbd, 0x8b, 0x90, 0x92, 0x76, 0x78, 0x53, 0x92, 0xbe, 0x4a,
	0x14, 0x8d, 0x01, 0x92, 0xdd, 0xcf, 0x21, 0xb7, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0x7d,
	0x86, 0xe3, 0xef, 0x00, 0x00, 0x00,
}

func (m *PageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PerPage != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.PerPage))
		i--
		dAtA[i] = 0x10
	}
	if m.Page != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Total != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPagination(dAtA []byte, offset int, v uint64) int {
	offset -= sovPagination(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *PageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Page != 0 {
		n += 1 + sovPagination(uint64(m.Page))
	}
	if m.PerPage != 0 {
		n += 1 + sovPagination(uint64(m.PerPage))
	}
	return n
}

func (m *PageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Total != 0 {
		n += 1 + sovPagination(uint64(m.Total))
	}
	return n
}

func sovPagination(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozPagination(x uint64) (n int) {
	return sovPagination(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *PageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagination
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
			return fmt.Errorf("proto: PageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerPage", wireType)
			}
			m.PerPage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PerPage |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPagination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPagination
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

func (m *PageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagination
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
			return fmt.Errorf("proto: PageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPagination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPagination
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

func skipPagination(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPagination
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
					return 0, ErrIntOverflowPagination
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
					return 0, ErrIntOverflowPagination
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
				return 0, ErrInvalidLengthPagination
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPagination
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPagination
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPagination        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPagination          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPagination = fmt.Errorf("proto: unexpected end of group")
)