// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: story/evmstaking/v1/types/evmstaking.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type WithdrawalType int32

const (
	WithdrawalType_WITHDRAWAL_TYPE_UNSTAKE WithdrawalType = 0
	WithdrawalType_WITHDRAWAL_TYPE_REWARD  WithdrawalType = 1
	WithdrawalType_WITHDRAWAL_TYPE_UBI     WithdrawalType = 2
)

var WithdrawalType_name = map[int32]string{
	0: "WITHDRAWAL_TYPE_UNSTAKE",
	1: "WITHDRAWAL_TYPE_REWARD",
	2: "WITHDRAWAL_TYPE_UBI",
}

var WithdrawalType_value = map[string]int32{
	"WITHDRAWAL_TYPE_UNSTAKE": 0,
	"WITHDRAWAL_TYPE_REWARD":  1,
	"WITHDRAWAL_TYPE_UBI":     2,
}

func (x WithdrawalType) String() string {
	return proto.EnumName(WithdrawalType_name, int32(x))
}

func (WithdrawalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fce68a5b3e5ef5ff, []int{0}
}

type Withdrawal struct {
	CreationHeight uint64 `protobuf:"varint,1,opt,name=creation_height,json=creationHeight,proto3" json:"creation_height,omitempty"`
	// TODO: use ethcommon.Address type
	ExecutionAddress string         `protobuf:"bytes,2,opt,name=execution_address,json=executionAddress,proto3" json:"execution_address,omitempty" yaml:"execution_address"`
	Amount           uint64         `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty" yaml:"amount"`
	WithdrawalType   WithdrawalType `protobuf:"varint,4,opt,name=withdrawal_type,json=withdrawalType,proto3,enum=story.evmstaking.v1.types.WithdrawalType" json:"withdrawal_type,omitempty" yaml:"withdrawal_type"`
}

func (m *Withdrawal) Reset()         { *m = Withdrawal{} }
func (m *Withdrawal) String() string { return proto.CompactTextString(m) }
func (*Withdrawal) ProtoMessage()    {}
func (*Withdrawal) Descriptor() ([]byte, []int) {
	return fileDescriptor_fce68a5b3e5ef5ff, []int{0}
}
func (m *Withdrawal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Withdrawal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Withdrawal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Withdrawal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Withdrawal.Merge(m, src)
}
func (m *Withdrawal) XXX_Size() int {
	return m.Size()
}
func (m *Withdrawal) XXX_DiscardUnknown() {
	xxx_messageInfo_Withdrawal.DiscardUnknown(m)
}

var xxx_messageInfo_Withdrawal proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("story.evmstaking.v1.types.WithdrawalType", WithdrawalType_name, WithdrawalType_value)
	proto.RegisterType((*Withdrawal)(nil), "story.evmstaking.v1.types.Withdrawal")
}

func init() {
	proto.RegisterFile("story/evmstaking/v1/types/evmstaking.proto", fileDescriptor_fce68a5b3e5ef5ff)
}

var fileDescriptor_fce68a5b3e5ef5ff = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0x2e, 0xc9, 0x2f,
	0xaa, 0xd4, 0x4f, 0x2d, 0xcb, 0x2d, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0xd4,
	0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x46, 0x12, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x04,
	0xab, 0xd5, 0x43, 0x12, 0x2f, 0x33, 0xd4, 0x03, 0xab, 0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0xab, 0xd2, 0x07, 0xb1, 0x20, 0x1a, 0xa4, 0x24, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xe3, 0x21,
	0x12, 0x10, 0x0e, 0x44, 0x4a, 0x69, 0x2f, 0x13, 0x17, 0x57, 0x78, 0x66, 0x49, 0x46, 0x4a, 0x51,
	0x62, 0x79, 0x62, 0x8e, 0x90, 0x3a, 0x17, 0x7f, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x66, 0x7e, 0x5e,
	0x7c, 0x46, 0x6a, 0x66, 0x7a, 0x46, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x1f, 0x4c,
	0xd8, 0x03, 0x2c, 0x2a, 0x94, 0xc8, 0x25, 0x98, 0x5a, 0x91, 0x9a, 0x5c, 0x0a, 0x56, 0x99, 0x98,
	0x92, 0x52, 0x94, 0x5a, 0x5c, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe9, 0x64, 0xf2, 0xe9, 0x9e,
	0xbc, 0x44, 0x65, 0x62, 0x6e, 0x8e, 0x95, 0x12, 0x86, 0x12, 0xa5, 0x4b, 0x5b, 0x74, 0x45, 0xa0,
	0x0e, 0x70, 0x84, 0x08, 0x05, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x07, 0x09, 0xc0, 0xd5, 0x42, 0xc5,
	0x85, 0x34, 0xb9, 0xd8, 0x12, 0x73, 0xf3, 0x4b, 0xf3, 0x4a, 0x24, 0x98, 0x41, 0x4e, 0x70, 0x12,
	0xfc, 0x74, 0x4f, 0x9e, 0x17, 0x62, 0x2e, 0x44, 0x5c, 0x29, 0x08, 0xaa, 0x40, 0x28, 0x8f, 0x8b,
	0xbf, 0x1c, 0xee, 0x89, 0x78, 0x50, 0x50, 0x48, 0xb0, 0x28, 0x30, 0x6a, 0xf0, 0x19, 0x69, 0xea,
	0xe1, 0x0c, 0x2b, 0x3d, 0x84, 0xb7, 0x43, 0x2a, 0x0b, 0x52, 0x9d, 0xa4, 0x3e, 0xdd, 0x93, 0x17,
	0x83, 0x18, 0x8f, 0x66, 0x96, 0x52, 0x10, 0x5f, 0x39, 0x8a, 0x5a, 0x2b, 0x8e, 0x8e, 0x05, 0xf2,
	0x0c, 0x2f, 0x16, 0xc8, 0x33, 0x6a, 0x25, 0x71, 0xf1, 0xa1, 0x9a, 0x23, 0x24, 0xcd, 0x25, 0x1e,
	0xee, 0x19, 0xe2, 0xe1, 0x12, 0xe4, 0x18, 0xee, 0xe8, 0x13, 0x1f, 0x12, 0x19, 0xe0, 0x1a, 0x1f,
	0xea, 0x17, 0x1c, 0xe2, 0xe8, 0xed, 0x2a, 0xc0, 0x20, 0x24, 0xc5, 0x25, 0x86, 0x2e, 0x19, 0xe4,
	0x1a, 0xee, 0x18, 0xe4, 0x22, 0xc0, 0x28, 0x24, 0xce, 0x25, 0x8c, 0xa1, 0xd1, 0xc9, 0x53, 0x80,
	0xc9, 0xc9, 0xe7, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x8c, 0xd2, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x0b, 0x32, 0x0b, 0x72, 0x12, 0x93, 0x8a,
	0xf5, 0x21, 0x09, 0x29, 0x39, 0x27, 0x33, 0x35, 0xaf, 0x44, 0xbf, 0x02, 0x39, 0x45, 0x81, 0xbd,
	0x9d, 0xc4, 0x06, 0x8e, 0x78, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x08, 0x03, 0xc9,
	0x72, 0x02, 0x00, 0x00,
}

func (this *Withdrawal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Withdrawal)
	if !ok {
		that2, ok := that.(Withdrawal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CreationHeight != that1.CreationHeight {
		return false
	}
	if this.ExecutionAddress != that1.ExecutionAddress {
		return false
	}
	if this.Amount != that1.Amount {
		return false
	}
	if this.WithdrawalType != that1.WithdrawalType {
		return false
	}
	return true
}
func (m *Withdrawal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Withdrawal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Withdrawal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WithdrawalType != 0 {
		i = encodeVarintEvmstaking(dAtA, i, uint64(m.WithdrawalType))
		i--
		dAtA[i] = 0x20
	}
	if m.Amount != 0 {
		i = encodeVarintEvmstaking(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ExecutionAddress) > 0 {
		i -= len(m.ExecutionAddress)
		copy(dAtA[i:], m.ExecutionAddress)
		i = encodeVarintEvmstaking(dAtA, i, uint64(len(m.ExecutionAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.CreationHeight != 0 {
		i = encodeVarintEvmstaking(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvmstaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvmstaking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Withdrawal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreationHeight != 0 {
		n += 1 + sovEvmstaking(uint64(m.CreationHeight))
	}
	l = len(m.ExecutionAddress)
	if l > 0 {
		n += 1 + l + sovEvmstaking(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovEvmstaking(uint64(m.Amount))
	}
	if m.WithdrawalType != 0 {
		n += 1 + sovEvmstaking(uint64(m.WithdrawalType))
	}
	return n
}

func sovEvmstaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvmstaking(x uint64) (n int) {
	return sovEvmstaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Withdrawal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvmstaking
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
			return fmt.Errorf("proto: Withdrawal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Withdrawal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmstaking
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
				return ErrInvalidLengthEvmstaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvmstaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutionAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalType", wireType)
			}
			m.WithdrawalType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawalType |= WithdrawalType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvmstaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvmstaking
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
func skipEvmstaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvmstaking
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
					return 0, ErrIntOverflowEvmstaking
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
					return 0, ErrIntOverflowEvmstaking
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
				return 0, ErrInvalidLengthEvmstaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvmstaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvmstaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvmstaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvmstaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvmstaking = fmt.Errorf("proto: unexpected end of group")
)
