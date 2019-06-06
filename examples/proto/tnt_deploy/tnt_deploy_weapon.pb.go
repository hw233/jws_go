// Code generated by protoc-gen-go.
// source: tnt_deploy_weapon.proto
// DO NOT EDIT!

/*
Package tnt_deploy is a generated protocol buffer package.

It is generated from these files:
	tnt_deploy_weapon.proto

It has these top-level messages:
	WEAPON
	WEAPON_ARRAY
*/
package tnt_deploy

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type WEAPON struct {
	// * 商品ID
	ID *string `protobuf:"bytes,1,req,def=" json:"ID,omitempty"`
	// *
	Type *string `protobuf:"bytes,2,opt,def=" json:"Type,omitempty"`
	// *
	Part *string `protobuf:"bytes,3,opt,def=" json:"Part,omitempty"`
	// *
	EnableLevel *int32 `protobuf:"varint,4,opt,def=0" json:"EnableLevel,omitempty"`
	// *
	Tier *int32 `protobuf:"varint,5,opt,def=0" json:"Tier,omitempty"`
	// *
	RareLevel *int32 `protobuf:"varint,6,opt,def=0" json:"RareLevel,omitempty"`
	// *
	FuseLevelLimit *int32 `protobuf:"varint,7,opt,def=0" json:"FuseLevelLimit,omitempty"`
	// *
	LootScore *int32 `protobuf:"varint,8,opt,def=0" json:"LootScore,omitempty"`
	// *
	SetGear *int32 `protobuf:"varint,9,opt,def=0" json:"SetGear,omitempty"`
	// *
	NameIDS *string `protobuf:"bytes,10,opt,def=" json:"NameIDS,omitempty"`
	// *
	DescriptionIDS *string `protobuf:"bytes,11,opt,def=" json:"DescriptionIDS,omitempty"`
	// *
	Icon *string `protobuf:"bytes,12,opt,def=" json:"Icon,omitempty"`
	// *
	Instance *string `protobuf:"bytes,13,opt,def=" json:"Instance,omitempty"`
	// *
	Attack *float32 `protobuf:"fixed32,14,opt,def=0" json:"Attack,omitempty"`
	// *
	Defense *float32 `protobuf:"fixed32,15,opt,def=0" json:"Defense,omitempty"`
	// *
	HP               *float32 `protobuf:"fixed32,16,opt,def=0" json:"HP,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *WEAPON) Reset()         { *m = WEAPON{} }
func (m *WEAPON) String() string { return proto.CompactTextString(m) }
func (*WEAPON) ProtoMessage()    {}

const Default_WEAPON_EnableLevel int32 = 0
const Default_WEAPON_Tier int32 = 0
const Default_WEAPON_RareLevel int32 = 0
const Default_WEAPON_FuseLevelLimit int32 = 0
const Default_WEAPON_LootScore int32 = 0
const Default_WEAPON_SetGear int32 = 0
const Default_WEAPON_Attack float32 = 0
const Default_WEAPON_Defense float32 = 0
const Default_WEAPON_HP float32 = 0

func (m *WEAPON) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *WEAPON) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *WEAPON) GetPart() string {
	if m != nil && m.Part != nil {
		return *m.Part
	}
	return ""
}

func (m *WEAPON) GetEnableLevel() int32 {
	if m != nil && m.EnableLevel != nil {
		return *m.EnableLevel
	}
	return Default_WEAPON_EnableLevel
}

func (m *WEAPON) GetTier() int32 {
	if m != nil && m.Tier != nil {
		return *m.Tier
	}
	return Default_WEAPON_Tier
}

func (m *WEAPON) GetRareLevel() int32 {
	if m != nil && m.RareLevel != nil {
		return *m.RareLevel
	}
	return Default_WEAPON_RareLevel
}

func (m *WEAPON) GetFuseLevelLimit() int32 {
	if m != nil && m.FuseLevelLimit != nil {
		return *m.FuseLevelLimit
	}
	return Default_WEAPON_FuseLevelLimit
}

func (m *WEAPON) GetLootScore() int32 {
	if m != nil && m.LootScore != nil {
		return *m.LootScore
	}
	return Default_WEAPON_LootScore
}

func (m *WEAPON) GetSetGear() int32 {
	if m != nil && m.SetGear != nil {
		return *m.SetGear
	}
	return Default_WEAPON_SetGear
}

func (m *WEAPON) GetNameIDS() string {
	if m != nil && m.NameIDS != nil {
		return *m.NameIDS
	}
	return ""
}

func (m *WEAPON) GetDescriptionIDS() string {
	if m != nil && m.DescriptionIDS != nil {
		return *m.DescriptionIDS
	}
	return ""
}

func (m *WEAPON) GetIcon() string {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return ""
}

func (m *WEAPON) GetInstance() string {
	if m != nil && m.Instance != nil {
		return *m.Instance
	}
	return ""
}

func (m *WEAPON) GetAttack() float32 {
	if m != nil && m.Attack != nil {
		return *m.Attack
	}
	return Default_WEAPON_Attack
}

func (m *WEAPON) GetDefense() float32 {
	if m != nil && m.Defense != nil {
		return *m.Defense
	}
	return Default_WEAPON_Defense
}

func (m *WEAPON) GetHP() float32 {
	if m != nil && m.HP != nil {
		return *m.HP
	}
	return Default_WEAPON_HP
}

type WEAPON_ARRAY struct {
	Items            []*WEAPON `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *WEAPON_ARRAY) Reset()         { *m = WEAPON_ARRAY{} }
func (m *WEAPON_ARRAY) String() string { return proto.CompactTextString(m) }
func (*WEAPON_ARRAY) ProtoMessage()    {}

func (m *WEAPON_ARRAY) GetItems() []*WEAPON {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}