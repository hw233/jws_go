// Code generated by protoc-gen-go.
// source: ProtobufGen_firststagelimitreward.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type FIRSTSTAGELIMITREWARD struct {
	// * 关卡ID
	ID *string `protobuf:"bytes,1,req,def=" json:"ID,omitempty"`
	// * 战队经验
	CorpXP *uint32 `protobuf:"varint,2,opt,def=0" json:"CorpXP,omitempty"`
	// * 经验
	XP *uint32 `protobuf:"varint,3,opt,def=0" json:"XP,omitempty"`
	// * 金钱
	SC *uint32 `protobuf:"varint,4,opt,def=0" json:"SC,omitempty"`
	// * 手动经验
	ManualXP           *uint32                               `protobuf:"varint,5,opt,def=0" json:"ManualXP,omitempty"`
	SRewardLimit_Table []*FIRSTSTAGELIMITREWARD_SRewardLimit `protobuf:"bytes,6,rep" json:"SRewardLimit_Table,omitempty"`
	XXX_unrecognized   []byte                                `json:"-"`
}

func (m *FIRSTSTAGELIMITREWARD) Reset()         { *m = FIRSTSTAGELIMITREWARD{} }
func (m *FIRSTSTAGELIMITREWARD) String() string { return proto.CompactTextString(m) }
func (*FIRSTSTAGELIMITREWARD) ProtoMessage()    {}

const Default_FIRSTSTAGELIMITREWARD_CorpXP uint32 = 0
const Default_FIRSTSTAGELIMITREWARD_XP uint32 = 0
const Default_FIRSTSTAGELIMITREWARD_SC uint32 = 0
const Default_FIRSTSTAGELIMITREWARD_ManualXP uint32 = 0

func (m *FIRSTSTAGELIMITREWARD) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *FIRSTSTAGELIMITREWARD) GetCorpXP() uint32 {
	if m != nil && m.CorpXP != nil {
		return *m.CorpXP
	}
	return Default_FIRSTSTAGELIMITREWARD_CorpXP
}

func (m *FIRSTSTAGELIMITREWARD) GetXP() uint32 {
	if m != nil && m.XP != nil {
		return *m.XP
	}
	return Default_FIRSTSTAGELIMITREWARD_XP
}

func (m *FIRSTSTAGELIMITREWARD) GetSC() uint32 {
	if m != nil && m.SC != nil {
		return *m.SC
	}
	return Default_FIRSTSTAGELIMITREWARD_SC
}

func (m *FIRSTSTAGELIMITREWARD) GetManualXP() uint32 {
	if m != nil && m.ManualXP != nil {
		return *m.ManualXP
	}
	return Default_FIRSTSTAGELIMITREWARD_ManualXP
}

func (m *FIRSTSTAGELIMITREWARD) GetSRewardLimit_Table() []*FIRSTSTAGELIMITREWARD_SRewardLimit {
	if m != nil {
		return m.SRewardLimit_Table
	}
	return nil
}

type FIRSTSTAGELIMITREWARD_SRewardLimit struct {
	// * 物品组ID
	ItemGroupID *string `protobuf:"bytes,1,opt,def=" json:"ItemGroupID,omitempty"`
	// * （数量控制）掉落次数
	LootNum *int32 `protobuf:"varint,2,opt,def=1" json:"LootNum,omitempty"`
	// * （数量控制）随机区间
	LootSpace *int32 `protobuf:"varint,3,opt,def=0" json:"LootSpace,omitempty"`
	// * （数量控制）区间偏移量
	Offset *int32 `protobuf:"varint,4,opt,def=0" json:"Offset,omitempty"`
	// * 补偿物品组ID
	MItemGroupID     *string `protobuf:"bytes,5,opt,def=" json:"MItemGroupID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FIRSTSTAGELIMITREWARD_SRewardLimit) Reset()         { *m = FIRSTSTAGELIMITREWARD_SRewardLimit{} }
func (m *FIRSTSTAGELIMITREWARD_SRewardLimit) String() string { return proto.CompactTextString(m) }
func (*FIRSTSTAGELIMITREWARD_SRewardLimit) ProtoMessage()    {}

const Default_FIRSTSTAGELIMITREWARD_SRewardLimit_LootNum int32 = 1
const Default_FIRSTSTAGELIMITREWARD_SRewardLimit_LootSpace int32 = 0
const Default_FIRSTSTAGELIMITREWARD_SRewardLimit_Offset int32 = 0

func (m *FIRSTSTAGELIMITREWARD_SRewardLimit) GetItemGroupID() string {
	if m != nil && m.ItemGroupID != nil {
		return *m.ItemGroupID
	}
	return ""
}

func (m *FIRSTSTAGELIMITREWARD_SRewardLimit) GetLootNum() int32 {
	if m != nil && m.LootNum != nil {
		return *m.LootNum
	}
	return Default_FIRSTSTAGELIMITREWARD_SRewardLimit_LootNum
}

func (m *FIRSTSTAGELIMITREWARD_SRewardLimit) GetLootSpace() int32 {
	if m != nil && m.LootSpace != nil {
		return *m.LootSpace
	}
	return Default_FIRSTSTAGELIMITREWARD_SRewardLimit_LootSpace
}

func (m *FIRSTSTAGELIMITREWARD_SRewardLimit) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return Default_FIRSTSTAGELIMITREWARD_SRewardLimit_Offset
}

func (m *FIRSTSTAGELIMITREWARD_SRewardLimit) GetMItemGroupID() string {
	if m != nil && m.MItemGroupID != nil {
		return *m.MItemGroupID
	}
	return ""
}

type FIRSTSTAGELIMITREWARD_ARRAY struct {
	Items            []*FIRSTSTAGELIMITREWARD `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *FIRSTSTAGELIMITREWARD_ARRAY) Reset()         { *m = FIRSTSTAGELIMITREWARD_ARRAY{} }
func (m *FIRSTSTAGELIMITREWARD_ARRAY) String() string { return proto.CompactTextString(m) }
func (*FIRSTSTAGELIMITREWARD_ARRAY) ProtoMessage()    {}

func (m *FIRSTSTAGELIMITREWARD_ARRAY) GetItems() []*FIRSTSTAGELIMITREWARD {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}