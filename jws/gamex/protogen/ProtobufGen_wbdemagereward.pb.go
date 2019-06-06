// Code generated by protoc-gen-go.
// source: ProtobufGen_wbdemagereward.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type WBDEMAGEREWARD struct {
	// * 主键ID
	ID *uint32 `protobuf:"varint,1,req,def=0" json:"ID,omitempty"`
	// * 要求达到的伤害量
	NeedDemage       *uint32                    `protobuf:"varint,2,req,def=0" json:"NeedDemage,omitempty"`
	Loot_Table       []*WBDEMAGEREWARD_LootRule `protobuf:"bytes,3,rep" json:"Loot_Table,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *WBDEMAGEREWARD) Reset()         { *m = WBDEMAGEREWARD{} }
func (m *WBDEMAGEREWARD) String() string { return proto.CompactTextString(m) }
func (*WBDEMAGEREWARD) ProtoMessage()    {}

const Default_WBDEMAGEREWARD_ID uint32 = 0
const Default_WBDEMAGEREWARD_NeedDemage uint32 = 0

func (m *WBDEMAGEREWARD) GetID() uint32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return Default_WBDEMAGEREWARD_ID
}

func (m *WBDEMAGEREWARD) GetNeedDemage() uint32 {
	if m != nil && m.NeedDemage != nil {
		return *m.NeedDemage
	}
	return Default_WBDEMAGEREWARD_NeedDemage
}

func (m *WBDEMAGEREWARD) GetLoot_Table() []*WBDEMAGEREWARD_LootRule {
	if m != nil {
		return m.Loot_Table
	}
	return nil
}

type WBDEMAGEREWARD_LootRule struct {
	// * 奖励1
	ItemID *string `protobuf:"bytes,1,opt,def=" json:"ItemID,omitempty"`
	// * 奖励数量
	ItemNum          *uint32 `protobuf:"varint,2,opt,def=0" json:"ItemNum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WBDEMAGEREWARD_LootRule) Reset()         { *m = WBDEMAGEREWARD_LootRule{} }
func (m *WBDEMAGEREWARD_LootRule) String() string { return proto.CompactTextString(m) }
func (*WBDEMAGEREWARD_LootRule) ProtoMessage()    {}

const Default_WBDEMAGEREWARD_LootRule_ItemNum uint32 = 0

func (m *WBDEMAGEREWARD_LootRule) GetItemID() string {
	if m != nil && m.ItemID != nil {
		return *m.ItemID
	}
	return ""
}

func (m *WBDEMAGEREWARD_LootRule) GetItemNum() uint32 {
	if m != nil && m.ItemNum != nil {
		return *m.ItemNum
	}
	return Default_WBDEMAGEREWARD_LootRule_ItemNum
}

type WBDEMAGEREWARD_ARRAY struct {
	Items            []*WBDEMAGEREWARD `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *WBDEMAGEREWARD_ARRAY) Reset()         { *m = WBDEMAGEREWARD_ARRAY{} }
func (m *WBDEMAGEREWARD_ARRAY) String() string { return proto.CompactTextString(m) }
func (*WBDEMAGEREWARD_ARRAY) ProtoMessage()    {}

func (m *WBDEMAGEREWARD_ARRAY) GetItems() []*WBDEMAGEREWARD {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}