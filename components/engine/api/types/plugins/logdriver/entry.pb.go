// Code generated by protoc-gen-gogo.
// source: entry.proto
// DO NOT EDIT!

/*
	Package logdriver is a generated protocol buffer package.

	It is generated from these files:
		entry.proto

	It has these top-level messages:
		LogEntry
		PartialLogEntryMetadata
*/
package logdriver

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type LogEntry struct {
	Source             string                   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	TimeNano           int64                    `protobuf:"varint,2,opt,name=time_nano,json=timeNano,proto3" json:"time_nano,omitempty"`
	Line               []byte                   `protobuf:"bytes,3,opt,name=line,proto3" json:"line,omitempty"`
	Partial            bool                     `protobuf:"varint,4,opt,name=partial,proto3" json:"partial,omitempty"`
	PartialLogMetadata *PartialLogEntryMetadata `protobuf:"bytes,5,opt,name=partial_log_metadata,json=partialLogMetadata" json:"partial_log_metadata,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptorEntry, []int{0} }

func (m *LogEntry) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *LogEntry) GetTimeNano() int64 {
	if m != nil {
		return m.TimeNano
	}
	return 0
}

func (m *LogEntry) GetLine() []byte {
	if m != nil {
		return m.Line
	}
	return nil
}

func (m *LogEntry) GetPartial() bool {
	if m != nil {
		return m.Partial
	}
	return false
}

func (m *LogEntry) GetPartialLogMetadata() *PartialLogEntryMetadata {
	if m != nil {
		return m.PartialLogMetadata
	}
	return nil
}

type PartialLogEntryMetadata struct {
	Last    bool   `protobuf:"varint,1,opt,name=last,proto3" json:"last,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Ordinal int32  `protobuf:"varint,3,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
}

func (m *PartialLogEntryMetadata) Reset()                    { *m = PartialLogEntryMetadata{} }
func (m *PartialLogEntryMetadata) String() string            { return proto.CompactTextString(m) }
func (*PartialLogEntryMetadata) ProtoMessage()               {}
func (*PartialLogEntryMetadata) Descriptor() ([]byte, []int) { return fileDescriptorEntry, []int{1} }

func (m *PartialLogEntryMetadata) GetLast() bool {
	if m != nil {
		return m.Last
	}
	return false
}

func (m *PartialLogEntryMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PartialLogEntryMetadata) GetOrdinal() int32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "LogEntry")
	proto.RegisterType((*PartialLogEntryMetadata)(nil), "PartialLogEntryMetadata")
}
func (m *LogEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Source) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEntry(dAtA, i, uint64(len(m.Source)))
		i += copy(dAtA[i:], m.Source)
	}
	if m.TimeNano != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEntry(dAtA, i, uint64(m.TimeNano))
	}
	if len(m.Line) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEntry(dAtA, i, uint64(len(m.Line)))
		i += copy(dAtA[i:], m.Line)
	}
	if m.Partial {
		dAtA[i] = 0x20
		i++
		if m.Partial {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.PartialLogMetadata != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintEntry(dAtA, i, uint64(m.PartialLogMetadata.Size()))
		n1, err := m.PartialLogMetadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *PartialLogEntryMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartialLogEntryMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Last {
		dAtA[i] = 0x8
		i++
		if m.Last {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEntry(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Ordinal != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEntry(dAtA, i, uint64(m.Ordinal))
	}
	return i, nil
}

func encodeFixed64Entry(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Entry(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintEntry(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LogEntry) Size() (n int) {
	var l int
	_ = l
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovEntry(uint64(l))
	}
	if m.TimeNano != 0 {
		n += 1 + sovEntry(uint64(m.TimeNano))
	}
	l = len(m.Line)
	if l > 0 {
		n += 1 + l + sovEntry(uint64(l))
	}
	if m.Partial {
		n += 2
	}
	if m.PartialLogMetadata != nil {
		l = m.PartialLogMetadata.Size()
		n += 1 + l + sovEntry(uint64(l))
	}
	return n
}

func (m *PartialLogEntryMetadata) Size() (n int) {
	var l int
	_ = l
	if m.Last {
		n += 2
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEntry(uint64(l))
	}
	if m.Ordinal != 0 {
		n += 1 + sovEntry(uint64(m.Ordinal))
	}
	return n
}

func sovEntry(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEntry(x uint64) (n int) {
	return sovEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LogEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeNano", wireType)
			}
			m.TimeNano = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeNano |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEntry
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Line = append(m.Line[:0], dAtA[iNdEx:postIndex]...)
			if m.Line == nil {
				m.Line = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Partial", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Partial = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartialLogMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PartialLogMetadata == nil {
				m.PartialLogMetadata = &PartialLogEntryMetadata{}
			}
			if err := m.PartialLogMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntry
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
func (m *PartialLogEntryMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PartialLogEntryMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartialLogEntryMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Last", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Last = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ordinal", wireType)
			}
			m.Ordinal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ordinal |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntry
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
func skipEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntry
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
					return 0, ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEntry
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEntry
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEntry
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEntry(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEntry = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntry   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("entry.proto", fileDescriptorEntry) }

var fileDescriptorEntry = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xbd, 0x4a, 0x04, 0x31,
	0x14, 0x85, 0xb9, 0xb3, 0x3f, 0xce, 0xdc, 0x5d, 0x2c, 0x82, 0x68, 0x40, 0x18, 0xc2, 0x56, 0xa9,
	0xb6, 0xd0, 0x37, 0x10, 0x6c, 0x44, 0x45, 0xd2, 0x58, 0x0e, 0x57, 0x27, 0x2c, 0x81, 0xd9, 0xdc,
	0x21, 0x13, 0x0b, 0x1f, 0xcd, 0x37, 0xb0, 0xf4, 0x11, 0x64, 0x9e, 0x44, 0x26, 0x4e, 0xec, 0xec,
	0xce, 0x39, 0x5f, 0x8a, 0x2f, 0x17, 0x37, 0xd6, 0xc7, 0xf0, 0xbe, 0xef, 0x03, 0x47, 0xde, 0x7d,
	0x00, 0x96, 0xf7, 0x7c, 0xb8, 0x9d, 0x26, 0x71, 0x8e, 0xeb, 0x81, 0xdf, 0xc2, 0xab, 0x95, 0xa0,
	0x40, 0x57, 0x66, 0x6e, 0xe2, 0x12, 0xab, 0xe8, 0x8e, 0xb6, 0xf1, 0xe4, 0x59, 0x16, 0x0a, 0xf4,
	0xc2, 0x94, 0xd3, 0xf0, 0x48, 0x9e, 0x85, 0xc0, 0x65, 0xe7, 0xbc, 0x95, 0x0b, 0x05, 0x7a, 0x6b,
	0x52, 0x16, 0x12, 0x4f, 0x7a, 0x0a, 0xd1, 0x51, 0x27, 0x97, 0x0a, 0x74, 0x69, 0x72, 0x15, 0x77,
	0x78, 0x36, 0xc7, 0xa6, 0xe3, 0x43, 0x73, 0xb4, 0x91, 0x5a, 0x8a, 0x24, 0x57, 0x0a, 0xf4, 0xe6,
	0x4a, 0xee, 0x9f, 0x7e, 0x61, 0x56, 0x7a, 0x98, 0xb9, 0x11, 0xfd, 0x1f, 0xc8, 0xdb, 0xee, 0x19,
	0x2f, 0xfe, 0x79, 0x9e, 0xa4, 0x68, 0x88, 0xe9, 0x1f, 0xa5, 0x49, 0x59, 0x9c, 0x62, 0xe1, 0xda,
	0xa4, 0x5f, 0x99, 0xc2, 0xb5, 0x93, 0x24, 0x87, 0xd6, 0x79, 0xea, 0x92, 0xfb, 0xca, 0xe4, 0x7a,
	0xb3, 0xfd, 0x1c, 0x6b, 0xf8, 0x1a, 0x6b, 0xf8, 0x1e, 0x6b, 0x78, 0x59, 0xa7, 0x4b, 0x5d, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xed, 0x9f, 0xb6, 0x38, 0x01, 0x00, 0x00,
}
