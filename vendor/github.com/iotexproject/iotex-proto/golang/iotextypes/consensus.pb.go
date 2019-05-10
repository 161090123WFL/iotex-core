// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/types/consensus.proto

package iotextypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ConsensusVote_Topic int32

const (
	ConsensusVote_PROPOSAL ConsensusVote_Topic = 0
	ConsensusVote_LOCK     ConsensusVote_Topic = 1
	ConsensusVote_COMMIT   ConsensusVote_Topic = 2
)

var ConsensusVote_Topic_name = map[int32]string{
	0: "PROPOSAL",
	1: "LOCK",
	2: "COMMIT",
}

var ConsensusVote_Topic_value = map[string]int32{
	"PROPOSAL": 0,
	"LOCK":     1,
	"COMMIT":   2,
}

func (x ConsensusVote_Topic) String() string {
	return proto.EnumName(ConsensusVote_Topic_name, int32(x))
}

func (ConsensusVote_Topic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2637092b19291c2e, []int{1, 0}
}

type BlockProposal struct {
	Block                *Block         `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Endorsements         []*Endorsement `protobuf:"bytes,2,rep,name=endorsements,proto3" json:"endorsements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BlockProposal) Reset()         { *m = BlockProposal{} }
func (m *BlockProposal) String() string { return proto.CompactTextString(m) }
func (*BlockProposal) ProtoMessage()    {}
func (*BlockProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_2637092b19291c2e, []int{0}
}

func (m *BlockProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockProposal.Unmarshal(m, b)
}
func (m *BlockProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockProposal.Marshal(b, m, deterministic)
}
func (m *BlockProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockProposal.Merge(m, src)
}
func (m *BlockProposal) XXX_Size() int {
	return xxx_messageInfo_BlockProposal.Size(m)
}
func (m *BlockProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockProposal.DiscardUnknown(m)
}

var xxx_messageInfo_BlockProposal proto.InternalMessageInfo

func (m *BlockProposal) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *BlockProposal) GetEndorsements() []*Endorsement {
	if m != nil {
		return m.Endorsements
	}
	return nil
}

type ConsensusVote struct {
	BlockHash            []byte              `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Topic                ConsensusVote_Topic `protobuf:"varint,2,opt,name=topic,proto3,enum=iotextypes.ConsensusVote_Topic" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ConsensusVote) Reset()         { *m = ConsensusVote{} }
func (m *ConsensusVote) String() string { return proto.CompactTextString(m) }
func (*ConsensusVote) ProtoMessage()    {}
func (*ConsensusVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_2637092b19291c2e, []int{1}
}

func (m *ConsensusVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusVote.Unmarshal(m, b)
}
func (m *ConsensusVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusVote.Marshal(b, m, deterministic)
}
func (m *ConsensusVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusVote.Merge(m, src)
}
func (m *ConsensusVote) XXX_Size() int {
	return xxx_messageInfo_ConsensusVote.Size(m)
}
func (m *ConsensusVote) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusVote.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusVote proto.InternalMessageInfo

func (m *ConsensusVote) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *ConsensusVote) GetTopic() ConsensusVote_Topic {
	if m != nil {
		return m.Topic
	}
	return ConsensusVote_PROPOSAL
}

type ConsensusMessage struct {
	Height      uint64       `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Endorsement *Endorsement `protobuf:"bytes,2,opt,name=endorsement,proto3" json:"endorsement,omitempty"`
	// Types that are valid to be assigned to Msg:
	//	*ConsensusMessage_BlockProposal
	//	*ConsensusMessage_Vote
	Msg                  isConsensusMessage_Msg `protobuf_oneof:"msg"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConsensusMessage) Reset()         { *m = ConsensusMessage{} }
func (m *ConsensusMessage) String() string { return proto.CompactTextString(m) }
func (*ConsensusMessage) ProtoMessage()    {}
func (*ConsensusMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2637092b19291c2e, []int{2}
}

func (m *ConsensusMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusMessage.Unmarshal(m, b)
}
func (m *ConsensusMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusMessage.Marshal(b, m, deterministic)
}
func (m *ConsensusMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusMessage.Merge(m, src)
}
func (m *ConsensusMessage) XXX_Size() int {
	return xxx_messageInfo_ConsensusMessage.Size(m)
}
func (m *ConsensusMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusMessage proto.InternalMessageInfo

func (m *ConsensusMessage) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ConsensusMessage) GetEndorsement() *Endorsement {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

type isConsensusMessage_Msg interface {
	isConsensusMessage_Msg()
}

type ConsensusMessage_BlockProposal struct {
	BlockProposal *BlockProposal `protobuf:"bytes,100,opt,name=blockProposal,proto3,oneof"`
}

type ConsensusMessage_Vote struct {
	Vote *ConsensusVote `protobuf:"bytes,101,opt,name=vote,proto3,oneof"`
}

func (*ConsensusMessage_BlockProposal) isConsensusMessage_Msg() {}

func (*ConsensusMessage_Vote) isConsensusMessage_Msg() {}

func (m *ConsensusMessage) GetMsg() isConsensusMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ConsensusMessage) GetBlockProposal() *BlockProposal {
	if x, ok := m.GetMsg().(*ConsensusMessage_BlockProposal); ok {
		return x.BlockProposal
	}
	return nil
}

func (m *ConsensusMessage) GetVote() *ConsensusVote {
	if x, ok := m.GetMsg().(*ConsensusMessage_Vote); ok {
		return x.Vote
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConsensusMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConsensusMessage_BlockProposal)(nil),
		(*ConsensusMessage_Vote)(nil),
	}
}

func init() {
	proto.RegisterEnum("iotextypes.ConsensusVote_Topic", ConsensusVote_Topic_name, ConsensusVote_Topic_value)
	proto.RegisterType((*BlockProposal)(nil), "iotextypes.BlockProposal")
	proto.RegisterType((*ConsensusVote)(nil), "iotextypes.ConsensusVote")
	proto.RegisterType((*ConsensusMessage)(nil), "iotextypes.ConsensusMessage")
}

func init() { proto.RegisterFile("proto/types/consensus.proto", fileDescriptor_2637092b19291c2e) }

var fileDescriptor_2637092b19291c2e = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6a, 0xe2, 0x40,
	0x14, 0xc6, 0x8d, 0x1a, 0x71, 0x8f, 0xba, 0x64, 0xe7, 0x62, 0x37, 0xeb, 0xba, 0xac, 0xe4, 0x66,
	0x85, 0xd2, 0x04, 0x2c, 0x52, 0x4a, 0xaf, 0x54, 0x0a, 0x96, 0x2a, 0x91, 0xa9, 0xf4, 0xa2, 0x77,
	0x49, 0x1c, 0x92, 0xb4, 0x9a, 0x13, 0x32, 0x63, 0x69, 0x1f, 0xa3, 0x6f, 0xd8, 0x47, 0x29, 0x4e,
	0x6a, 0x33, 0x41, 0xe8, 0xe5, 0x99, 0xef, 0xf7, 0x7d, 0xe7, 0x0f, 0x03, 0x7f, 0xd2, 0x0c, 0x05,
	0x3a, 0xe2, 0x25, 0x65, 0xdc, 0x09, 0x30, 0xe1, 0x2c, 0xe1, 0x3b, 0x6e, 0xcb, 0x57, 0x02, 0x31,
	0x0a, 0xf6, 0x2c, 0xb5, 0x6e, 0x4f, 0x05, 0xfd, 0x0d, 0x06, 0x8f, 0x41, 0xe4, 0xc5, 0x49, 0x4e,
	0x76, 0xff, 0xaa, 0x2a, 0x4b, 0xd6, 0x98, 0x71, 0xb6, 0x65, 0x89, 0xc8, 0x65, 0x6b, 0x07, 0x9d,
	0xc9, 0xde, 0xb2, 0xcc, 0x30, 0x45, 0xee, 0x6d, 0xc8, 0x7f, 0xd0, 0x65, 0x86, 0xa9, 0xf5, 0xb5,
	0x41, 0x6b, 0xf8, 0xc3, 0x2e, 0x3a, 0xd9, 0x92, 0xa4, 0xb9, 0x4e, 0x2e, 0xa1, 0xad, 0xc4, 0x71,
	0xb3, 0xda, 0xaf, 0x0d, 0x5a, 0xc3, 0x5f, 0x2a, 0x7f, 0x55, 0xe8, 0xb4, 0x04, 0x5b, 0xaf, 0x1a,
	0x74, 0xa6, 0x87, 0x9d, 0xee, 0x50, 0x30, 0xd2, 0x83, 0x6f, 0x32, 0x77, 0xe6, 0xf1, 0x48, 0xf6,
	0x6e, 0xd3, 0xe2, 0x81, 0x8c, 0x40, 0x17, 0x98, 0xc6, 0x81, 0x59, 0xed, 0x6b, 0x83, 0xef, 0xc3,
	0x7f, 0x6a, 0x97, 0x52, 0x8e, 0xbd, 0xda, 0x63, 0x34, 0xa7, 0xad, 0x13, 0xd0, 0x65, 0x4d, 0xda,
	0xd0, 0x5c, 0x52, 0x77, 0xe9, 0xde, 0x8e, 0xe7, 0x46, 0x85, 0x34, 0xa1, 0x3e, 0x77, 0xa7, 0x37,
	0x86, 0x46, 0x00, 0x1a, 0x53, 0x77, 0xb1, 0xb8, 0x5e, 0x19, 0x55, 0xeb, 0x4d, 0x03, 0xe3, 0x33,
	0x6b, 0xc1, 0x38, 0xf7, 0x42, 0x46, 0x7e, 0x42, 0x23, 0x62, 0x71, 0x18, 0x09, 0x39, 0x53, 0x9d,
	0x7e, 0x54, 0xe4, 0x02, 0x5a, 0xca, 0x42, 0x72, 0xac, 0x2f, 0x96, 0x57, 0x59, 0x32, 0x86, 0x8e,
	0xaf, 0x9e, 0xdc, 0x5c, 0x4b, 0xf3, 0xef, 0xa3, 0x4b, 0x1f, 0x80, 0x59, 0x85, 0x96, 0x1d, 0xc4,
	0x81, 0xfa, 0x13, 0x0a, 0x66, 0xb2, 0x63, 0x67, 0xe9, 0x1a, 0xb3, 0x0a, 0x95, 0xe0, 0x44, 0x87,
	0xda, 0x96, 0x87, 0x93, 0xf3, 0xfb, 0x51, 0x18, 0x8b, 0x68, 0xe7, 0xdb, 0x01, 0x6e, 0x1d, 0xe9,
	0x4a, 0x33, 0x7c, 0x60, 0x81, 0xc8, 0x8b, 0xd3, 0xfc, 0xb3, 0x84, 0xb8, 0xf1, 0x92, 0xd0, 0x29,
	0x52, 0xfd, 0x86, 0x14, 0xce, 0xde, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xcf, 0xe0, 0x44, 0x95,
	0x02, 0x00, 0x00,
}