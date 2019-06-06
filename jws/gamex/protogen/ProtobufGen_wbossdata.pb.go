// Code generated by protoc-gen-go.
// source: ProtobufGen_wbossdata.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type WBOSSDATA struct {
	// * boss等级
	BossLevel *uint32 `protobuf:"varint,1,req,def=0" json:"BossLevel,omitempty"`
	// * 标准生命值
	HitPoint *int64 `protobuf:"varint,2,opt,def=0" json:"HitPoint,omitempty"`
	// * 标准伤害
	PhysicalDamage *uint32 `protobuf:"varint,3,opt,def=0" json:"PhysicalDamage,omitempty"`
	// * 标准防御
	PhysicalResist   *uint32               `protobuf:"varint,4,opt,def=0" json:"PhysicalResist,omitempty"`
	Loot_Table       []*WBOSSDATA_LootRule `protobuf:"bytes,5,rep" json:"Loot_Table,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *WBOSSDATA) Reset()         { *m = WBOSSDATA{} }
func (m *WBOSSDATA) String() string { return proto.CompactTextString(m) }
func (*WBOSSDATA) ProtoMessage()    {}

const Default_WBOSSDATA_BossLevel uint32 = 0
const Default_WBOSSDATA_HitPoint int64 = 0
const Default_WBOSSDATA_PhysicalDamage uint32 = 0
const Default_WBOSSDATA_PhysicalResist uint32 = 0

func (m *WBOSSDATA) GetBossLevel() uint32 {
	if m != nil && m.BossLevel != nil {
		return *m.BossLevel
	}
	return Default_WBOSSDATA_BossLevel
}

func (m *WBOSSDATA) GetHitPoint() int64 {
	if m != nil && m.HitPoint != nil {
		return *m.HitPoint
	}
	return Default_WBOSSDATA_HitPoint
}

func (m *WBOSSDATA) GetPhysicalDamage() uint32 {
	if m != nil && m.PhysicalDamage != nil {
		return *m.PhysicalDamage
	}
	return Default_WBOSSDATA_PhysicalDamage
}

func (m *WBOSSDATA) GetPhysicalResist() uint32 {
	if m != nil && m.PhysicalResist != nil {
		return *m.PhysicalResist
	}
	return Default_WBOSSDATA_PhysicalResist
}

func (m *WBOSSDATA) GetLoot_Table() []*WBOSSDATA_LootRule {
	if m != nil {
		return m.Loot_Table
	}
	return nil
}

type WBOSSDATA_LootRule struct {
	// * 奖励1
	ItemID *string `protobuf:"bytes,1,opt,def=" json:"ItemID,omitempty"`
	// * 奖励数量
	ItemNum          *uint32 `protobuf:"varint,2,opt,def=0" json:"ItemNum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WBOSSDATA_LootRule) Reset()         { *m = WBOSSDATA_LootRule{} }
func (m *WBOSSDATA_LootRule) String() string { return proto.CompactTextString(m) }
func (*WBOSSDATA_LootRule) ProtoMessage()    {}

const Default_WBOSSDATA_LootRule_ItemNum uint32 = 0

func (m *WBOSSDATA_LootRule) GetItemID() string {
	if m != nil && m.ItemID != nil {
		return *m.ItemID
	}
	return ""
}

func (m *WBOSSDATA_LootRule) GetItemNum() uint32 {
	if m != nil && m.ItemNum != nil {
		return *m.ItemNum
	}
	return Default_WBOSSDATA_LootRule_ItemNum
}

type WBOSSDATA_ARRAY struct {
	Items            []*WBOSSDATA `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WBOSSDATA_ARRAY) Reset()         { *m = WBOSSDATA_ARRAY{} }
func (m *WBOSSDATA_ARRAY) String() string { return proto.CompactTextString(m) }
func (*WBOSSDATA_ARRAY) ProtoMessage()    {}

func (m *WBOSSDATA_ARRAY) GetItems() []*WBOSSDATA {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}