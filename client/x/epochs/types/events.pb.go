// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/x/epochs/types/events.proto

package types

import (
	fmt "fmt"
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

// EventEpochEnd is an event emitted when an epoch end.
type EventEpochEnd struct {
	EpochNumber int64 `protobuf:"varint,1,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
}

func (m *EventEpochEnd) Reset()         { *m = EventEpochEnd{} }
func (m *EventEpochEnd) String() string { return proto.CompactTextString(m) }
func (*EventEpochEnd) ProtoMessage()    {}
func (*EventEpochEnd) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d70eeda0e3075, []int{0}
}
func (m *EventEpochEnd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEpochEnd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEpochEnd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEpochEnd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEpochEnd.Merge(m, src)
}
func (m *EventEpochEnd) XXX_Size() int {
	return m.Size()
}
func (m *EventEpochEnd) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEpochEnd.DiscardUnknown(m)
}

var xxx_messageInfo_EventEpochEnd proto.InternalMessageInfo

func (m *EventEpochEnd) GetEpochNumber() int64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

// EventEpochStart is an event emitted when an epoch start.
type EventEpochStart struct {
	EpochNumber    int64 `protobuf:"varint,1,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	EpochStartTime int64 `protobuf:"varint,2,opt,name=epoch_start_time,json=epochStartTime,proto3" json:"epoch_start_time,omitempty"`
}

func (m *EventEpochStart) Reset()         { *m = EventEpochStart{} }
func (m *EventEpochStart) String() string { return proto.CompactTextString(m) }
func (*EventEpochStart) ProtoMessage()    {}
func (*EventEpochStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d70eeda0e3075, []int{1}
}
func (m *EventEpochStart) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEpochStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEpochStart.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEpochStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEpochStart.Merge(m, src)
}
func (m *EventEpochStart) XXX_Size() int {
	return m.Size()
}
func (m *EventEpochStart) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEpochStart.DiscardUnknown(m)
}

var xxx_messageInfo_EventEpochStart proto.InternalMessageInfo

func (m *EventEpochStart) GetEpochNumber() int64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func (m *EventEpochStart) GetEpochStartTime() int64 {
	if m != nil {
		return m.EpochStartTime
	}
	return 0
}

func init() {
	proto.RegisterType((*EventEpochEnd)(nil), "client.x.epochs.types.EventEpochEnd")
	proto.RegisterType((*EventEpochStart)(nil), "client.x.epochs.types.EventEpochStart")
}

func init() {
	proto.RegisterFile("client/x/epochs/types/events.proto", fileDescriptor_ee2d70eeda0e3075)
}

var fileDescriptor_ee2d70eeda0e3075 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0xaf, 0xd0, 0x4f, 0x2d, 0xc8, 0x4f, 0xce, 0x28, 0xd6, 0x2f, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x85, 0xa8, 0xd1, 0xab, 0xd0, 0x83, 0xa8, 0xd1, 0x03, 0xab, 0x51, 0x32, 0xe2, 0xe2, 0x75, 0x05,
	0x29, 0x73, 0x05, 0x09, 0xba, 0xe6, 0xa5, 0x08, 0x29, 0x72, 0xf1, 0x80, 0x15, 0xc4, 0xe7, 0x95,
	0xe6, 0x26, 0xa5, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x71, 0x83, 0xc5, 0xfc, 0xc0,
	0x42, 0x4a, 0x71, 0x5c, 0xfc, 0x08, 0x3d, 0xc1, 0x25, 0x89, 0x45, 0x25, 0x44, 0xe8, 0x12, 0xd2,
	0xe0, 0x12, 0x80, 0x28, 0x29, 0x06, 0xe9, 0x88, 0x2f, 0xc9, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x2b,
	0xe3, 0x4b, 0x85, 0x1b, 0x14, 0x92, 0x99, 0x9b, 0xea, 0xa4, 0x7f, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3,
	0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xa2, 0x58, 0x3d, 0x9a, 0xc4, 0x06, 0xf6, 0xa2, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xa9, 0x19, 0xc8, 0xb7, 0x08, 0x01, 0x00, 0x00,
}

func (m *EventEpochEnd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEpochEnd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEpochEnd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventEpochStart) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEpochStart) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEpochStart) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochStartTime != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.EpochStartTime))
		i--
		dAtA[i] = 0x10
	}
	if m.EpochNumber != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventEpochEnd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNumber != 0 {
		n += 1 + sovEvents(uint64(m.EpochNumber))
	}
	return n
}

func (m *EventEpochStart) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNumber != 0 {
		n += 1 + sovEvents(uint64(m.EpochNumber))
	}
	if m.EpochStartTime != 0 {
		n += 1 + sovEvents(uint64(m.EpochStartTime))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventEpochEnd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventEpochEnd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEpochEnd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventEpochStart) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventEpochStart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEpochStart: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochStartTime", wireType)
			}
			m.EpochStartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochStartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)