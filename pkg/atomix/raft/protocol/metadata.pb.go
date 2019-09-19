// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/raft/protocol/metadata.proto

package protocol

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Raft system metadata
type RaftMetadata struct {
	Term Term     `protobuf:"varint,1,opt,name=term,proto3,casttype=Term" json:"term,omitempty"`
	Vote MemberID `protobuf:"bytes,2,opt,name=vote,proto3,casttype=MemberID" json:"vote,omitempty"`
}

func (m *RaftMetadata) Reset()         { *m = RaftMetadata{} }
func (m *RaftMetadata) String() string { return proto.CompactTextString(m) }
func (*RaftMetadata) ProtoMessage()    {}
func (*RaftMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c93df0fbe03b7c, []int{0}
}
func (m *RaftMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RaftMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RaftMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RaftMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftMetadata.Merge(m, src)
}
func (m *RaftMetadata) XXX_Size() int {
	return m.Size()
}
func (m *RaftMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RaftMetadata proto.InternalMessageInfo

func (m *RaftMetadata) GetTerm() Term {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RaftMetadata) GetVote() MemberID {
	if m != nil {
		return m.Vote
	}
	return ""
}

// Raft system configuration
type RaftConfiguration struct {
	Index     Index         `protobuf:"varint,1,opt,name=index,proto3,casttype=Index" json:"index,omitempty"`
	Term      Term          `protobuf:"varint,2,opt,name=term,proto3,casttype=Term" json:"term,omitempty"`
	Timestamp *time.Time    `protobuf:"bytes,3,opt,name=timestamp,proto3,stdtime" json:"timestamp,omitempty"`
	Members   []*RaftMember `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
}

func (m *RaftConfiguration) Reset()         { *m = RaftConfiguration{} }
func (m *RaftConfiguration) String() string { return proto.CompactTextString(m) }
func (*RaftConfiguration) ProtoMessage()    {}
func (*RaftConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c93df0fbe03b7c, []int{1}
}
func (m *RaftConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RaftConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RaftConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RaftConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftConfiguration.Merge(m, src)
}
func (m *RaftConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *RaftConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RaftConfiguration proto.InternalMessageInfo

func (m *RaftConfiguration) GetIndex() Index {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RaftConfiguration) GetTerm() Term {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RaftConfiguration) GetTimestamp() *time.Time {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *RaftConfiguration) GetMembers() []*RaftMember {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*RaftMetadata)(nil), "atomix.raft.protocol.RaftMetadata")
	proto.RegisterType((*RaftConfiguration)(nil), "atomix.raft.protocol.RaftConfiguration")
}

func init() {
	proto.RegisterFile("atomix/raft/protocol/metadata.proto", fileDescriptor_b1c93df0fbe03b7c)
}

var fileDescriptor_b1c93df0fbe03b7c = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x6e, 0xc2, 0x40,
	0x10, 0x45, 0x19, 0x30, 0x09, 0x2c, 0x34, 0xb1, 0x28, 0x2c, 0x14, 0xad, 0x2d, 0x92, 0xc2, 0xd5,
	0x5a, 0x22, 0x5d, 0x8a, 0x14, 0x4e, 0x1a, 0x0a, 0x52, 0x58, 0x5c, 0x60, 0x81, 0xb5, 0x65, 0x89,
	0x65, 0xd1, 0x32, 0x44, 0x1c, 0x83, 0x63, 0xe4, 0x08, 0x39, 0x42, 0xaa, 0x88, 0x32, 0x15, 0x49,
	0xcc, 0x25, 0x22, 0xaa, 0xc8, 0x5e, 0x4c, 0x1a, 0xba, 0xd1, 0x9f, 0x37, 0x7f, 0xe6, 0x0f, 0xb9,
	0xe1, 0xa8, 0x64, 0xba, 0x0e, 0x34, 0x8f, 0x31, 0x58, 0x68, 0x85, 0x6a, 0xa2, 0x66, 0x81, 0x14,
	0xc8, 0xa7, 0x1c, 0x39, 0x2b, 0x14, 0xbb, 0x63, 0x20, 0x96, 0x43, 0xac, 0x84, 0xba, 0xbd, 0xb3,
	0xa3, 0x93, 0xd9, 0x6a, 0x89, 0x42, 0x1b, 0xac, 0xeb, 0x26, 0x4a, 0x25, 0x33, 0x61, 0xda, 0xe3,
	0x55, 0x1c, 0x60, 0x2a, 0xc5, 0x12, 0xb9, 0x5c, 0x1c, 0x81, 0x4e, 0xa2, 0x12, 0x55, 0x94, 0x41,
	0x5e, 0x19, 0xb5, 0xf7, 0x4c, 0xda, 0x11, 0x8f, 0x71, 0x78, 0x3c, 0xc3, 0xbe, 0x26, 0x16, 0x0a,
	0x2d, 0x1d, 0xf0, 0xc0, 0xb7, 0xc2, 0xc6, 0x61, 0xe7, 0x5a, 0x23, 0xa1, 0x65, 0x54, 0xa8, 0xb6,
	0x47, 0xac, 0x17, 0x85, 0xc2, 0xa9, 0x7a, 0xe0, 0x37, 0xc3, 0xf6, 0x61, 0xe7, 0x36, 0x86, 0x42,
	0x8e, 0x85, 0x1e, 0x3c, 0x45, 0x45, 0xa7, 0xf7, 0x01, 0xe4, 0x2a, 0x37, 0x7c, 0x54, 0xf3, 0x38,
	0x4d, 0x56, 0x9a, 0x63, 0xaa, 0xe6, 0xb6, 0x4b, 0xea, 0xe9, 0x7c, 0x2a, 0xd6, 0x47, 0xdb, 0xe6,
	0x61, 0xe7, 0xd6, 0x07, 0xb9, 0x10, 0x19, 0xfd, 0xb4, 0xb6, 0x7a, 0x76, 0xed, 0x03, 0x69, 0x9e,
	0xd2, 0x38, 0x35, 0x0f, 0xfc, 0x56, 0xbf, 0xcb, 0x4c, 0x5e, 0x56, 0xe6, 0x65, 0xa3, 0x92, 0x08,
	0xad, 0xcd, 0x97, 0x0b, 0xd1, 0xff, 0x88, 0x7d, 0x4f, 0x2e, 0x65, 0x71, 0xe6, 0xd2, 0xb1, 0xbc,
	0x9a, 0xdf, 0xea, 0x7b, 0xec, 0xdc, 0x9f, 0x99, 0xf9, 0x44, 0x0e, 0x46, 0xe5, 0x40, 0x78, 0xfb,
	0xfb, 0x43, 0xe1, 0x35, 0xa3, 0xf0, 0x96, 0x51, 0x78, 0xcf, 0x28, 0x6c, 0x33, 0x0a, 0xdf, 0x19,
	0x85, 0xcd, 0x9e, 0x56, 0xb6, 0x7b, 0x5a, 0xf9, 0xdc, 0xd3, 0xca, 0xf8, 0xa2, 0xf0, 0xb8, 0xfb,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xcf, 0x2a, 0xa8, 0xe5, 0x01, 0x00, 0x00,
}

func (this *RaftMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RaftMetadata)
	if !ok {
		that2, ok := that.(RaftMetadata)
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
	if this.Term != that1.Term {
		return false
	}
	if this.Vote != that1.Vote {
		return false
	}
	return true
}
func (this *RaftConfiguration) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RaftConfiguration)
	if !ok {
		that2, ok := that.(RaftConfiguration)
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
	if this.Index != that1.Index {
		return false
	}
	if this.Term != that1.Term {
		return false
	}
	if that1.Timestamp == nil {
		if this.Timestamp != nil {
			return false
		}
	} else if !this.Timestamp.Equal(*that1.Timestamp) {
		return false
	}
	if len(this.Members) != len(that1.Members) {
		return false
	}
	for i := range this.Members {
		if !this.Members[i].Equal(that1.Members[i]) {
			return false
		}
	}
	return true
}
func (m *RaftMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RaftMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Vote) > 0 {
		i -= len(m.Vote)
		copy(dAtA[i:], m.Vote)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Vote)))
		i--
		dAtA[i] = 0x12
	}
	if m.Term != 0 {
		i = encodeVarintMetadata(dAtA, i, uint64(m.Term))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RaftConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RaftConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetadata(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Timestamp != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.Timestamp):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintMetadata(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x1a
	}
	if m.Term != 0 {
		i = encodeVarintMetadata(dAtA, i, uint64(m.Term))
		i--
		dAtA[i] = 0x10
	}
	if m.Index != 0 {
		i = encodeVarintMetadata(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedRaftMetadata(r randyMetadata, easy bool) *RaftMetadata {
	this := &RaftMetadata{}
	this.Term = Term(uint64(r.Uint32()))
	this.Vote = MemberID(randStringMetadata(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedRaftConfiguration(r randyMetadata, easy bool) *RaftConfiguration {
	this := &RaftConfiguration{}
	this.Index = Index(uint64(r.Uint32()))
	this.Term = Term(uint64(r.Uint32()))
	if r.Intn(5) != 0 {
		this.Timestamp = github_com_gogo_protobuf_types.NewPopulatedStdTime(r, easy)
	}
	if r.Intn(5) != 0 {
		v1 := r.Intn(5)
		this.Members = make([]*RaftMember, v1)
		for i := 0; i < v1; i++ {
			this.Members[i] = NewPopulatedRaftMember(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyMetadata interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMetadata(r randyMetadata) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMetadata(r randyMetadata) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneMetadata(r)
	}
	return string(tmps)
}
func randUnrecognizedMetadata(r randyMetadata, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMetadata(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMetadata(dAtA []byte, r randyMetadata, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMetadata(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateMetadata(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateMetadata(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMetadata(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMetadata(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMetadata(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMetadata(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *RaftMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Term != 0 {
		n += 1 + sovMetadata(uint64(m.Term))
	}
	l = len(m.Vote)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	return n
}

func (m *RaftConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovMetadata(uint64(m.Index))
	}
	if m.Term != 0 {
		n += 1 + sovMetadata(uint64(m.Term))
	}
	if m.Timestamp != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.Timestamp)
		n += 1 + l + sovMetadata(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	return n
}

func sovMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: RaftMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= Term(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vote = MemberID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func (m *RaftConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: RaftConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= Index(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= Term(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &RaftMember{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
				return 0, ErrInvalidLengthMetadata
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetadata
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
				next, err := skipMetadata(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMetadata
				}
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
	ErrInvalidLengthMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata   = fmt.Errorf("proto: integer overflow")
)