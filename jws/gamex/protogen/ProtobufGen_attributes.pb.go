// Code generated by protoc-gen-go.
// source: ProtobufGen_attributes.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ATTRIBUTES struct {
	// * 暴击率
	CritRate *float32 `protobuf:"fixed32,1,req,def=0" json:"CritRate,omitempty"`
	// * 暴击伤害
	CritValue *float32 `protobuf:"fixed32,2,req,def=0" json:"CritValue,omitempty"`
	// * 免暴率
	ResilienceRate *float32 `protobuf:"fixed32,3,req,def=0" json:"ResilienceRate,omitempty"`
	// * 免暴伤害
	ResilienceValue *float32 `protobuf:"fixed32,4,req,def=0" json:"ResilienceValue,omitempty"`
	// * 受伤值占总血量比例
	HurtRate *float32 `protobuf:"fixed32,5,req,def=0" json:"HurtRate,omitempty"`
	// * 闪避率
	DodgeRate *float32 `protobuf:"fixed32,24,req,def=0" json:"DodgeRate,omitempty"`
	// * 命中率
	HitRate *float32 `protobuf:"fixed32,25,req,def=0" json:"HitRate,omitempty"`
	// * 冰系攻击力
	IceDamage *int32 `protobuf:"varint,6,req,def=0" json:"IceDamage,omitempty"`
	// * 冰系防御力
	IceDefense *int32 `protobuf:"varint,7,req,def=0" json:"IceDefense,omitempty"`
	// * 冰系伤害加成
	IceBonus *float32 `protobuf:"fixed32,8,req,def=0" json:"IceBonus,omitempty"`
	// * 冰系抗性
	IceResist *float32 `protobuf:"fixed32,9,req,def=0" json:"IceResist,omitempty"`
	// * 火系攻击力
	FireDamage *int32 `protobuf:"varint,10,req,def=0" json:"FireDamage,omitempty"`
	// * 火系防御力
	FireDefense *int32 `protobuf:"varint,11,req,def=0" json:"FireDefense,omitempty"`
	// * 火系伤害加成
	FireBonus *float32 `protobuf:"fixed32,12,req,def=0" json:"FireBonus,omitempty"`
	// * 火系抗性
	FireResist *float32 `protobuf:"fixed32,13,req,def=0" json:"FireResist,omitempty"`
	// * 雷系攻击力
	LightingDamage *int32 `protobuf:"varint,14,req,def=0" json:"LightingDamage,omitempty"`
	// * 雷系防御力
	LightingDefense *int32 `protobuf:"varint,15,req,def=0" json:"LightingDefense,omitempty"`
	// * 雷系伤害加成
	LightingBonus *float32 `protobuf:"fixed32,16,req,def=0" json:"LightingBonus,omitempty"`
	// * 雷系抗性
	LightingResist *float32 `protobuf:"fixed32,17,req,def=0" json:"LightingResist,omitempty"`
	// * 毒系攻击力
	PoisonDamage *int32 `protobuf:"varint,18,req,def=0" json:"PoisonDamage,omitempty"`
	// * 毒系防御力
	PoisonDefense *int32 `protobuf:"varint,19,req,def=0" json:"PoisonDefense,omitempty"`
	// * 毒系伤害加成
	PoisonBonus *float32 `protobuf:"fixed32,20,req,def=0" json:"PoisonBonus,omitempty"`
	// * 毒系抗性
	PoisonResist *float32 `protobuf:"fixed32,21,req,def=0" json:"PoisonResist,omitempty"`
	// * 连击属性加成
	ComboRate        *float32 `protobuf:"fixed32,22,req,def=0" json:"ComboRate,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ATTRIBUTES) Reset()         { *m = ATTRIBUTES{} }
func (m *ATTRIBUTES) String() string { return proto.CompactTextString(m) }
func (*ATTRIBUTES) ProtoMessage()    {}

const Default_ATTRIBUTES_CritRate float32 = 0
const Default_ATTRIBUTES_CritValue float32 = 0
const Default_ATTRIBUTES_ResilienceRate float32 = 0
const Default_ATTRIBUTES_ResilienceValue float32 = 0
const Default_ATTRIBUTES_HurtRate float32 = 0
const Default_ATTRIBUTES_DodgeRate float32 = 0
const Default_ATTRIBUTES_HitRate float32 = 0
const Default_ATTRIBUTES_IceDamage int32 = 0
const Default_ATTRIBUTES_IceDefense int32 = 0
const Default_ATTRIBUTES_IceBonus float32 = 0
const Default_ATTRIBUTES_IceResist float32 = 0
const Default_ATTRIBUTES_FireDamage int32 = 0
const Default_ATTRIBUTES_FireDefense int32 = 0
const Default_ATTRIBUTES_FireBonus float32 = 0
const Default_ATTRIBUTES_FireResist float32 = 0
const Default_ATTRIBUTES_LightingDamage int32 = 0
const Default_ATTRIBUTES_LightingDefense int32 = 0
const Default_ATTRIBUTES_LightingBonus float32 = 0
const Default_ATTRIBUTES_LightingResist float32 = 0
const Default_ATTRIBUTES_PoisonDamage int32 = 0
const Default_ATTRIBUTES_PoisonDefense int32 = 0
const Default_ATTRIBUTES_PoisonBonus float32 = 0
const Default_ATTRIBUTES_PoisonResist float32 = 0
const Default_ATTRIBUTES_ComboRate float32 = 0

func (m *ATTRIBUTES) GetCritRate() float32 {
	if m != nil && m.CritRate != nil {
		return *m.CritRate
	}
	return Default_ATTRIBUTES_CritRate
}

func (m *ATTRIBUTES) GetCritValue() float32 {
	if m != nil && m.CritValue != nil {
		return *m.CritValue
	}
	return Default_ATTRIBUTES_CritValue
}

func (m *ATTRIBUTES) GetResilienceRate() float32 {
	if m != nil && m.ResilienceRate != nil {
		return *m.ResilienceRate
	}
	return Default_ATTRIBUTES_ResilienceRate
}

func (m *ATTRIBUTES) GetResilienceValue() float32 {
	if m != nil && m.ResilienceValue != nil {
		return *m.ResilienceValue
	}
	return Default_ATTRIBUTES_ResilienceValue
}

func (m *ATTRIBUTES) GetHurtRate() float32 {
	if m != nil && m.HurtRate != nil {
		return *m.HurtRate
	}
	return Default_ATTRIBUTES_HurtRate
}

func (m *ATTRIBUTES) GetDodgeRate() float32 {
	if m != nil && m.DodgeRate != nil {
		return *m.DodgeRate
	}
	return Default_ATTRIBUTES_DodgeRate
}

func (m *ATTRIBUTES) GetHitRate() float32 {
	if m != nil && m.HitRate != nil {
		return *m.HitRate
	}
	return Default_ATTRIBUTES_HitRate
}

func (m *ATTRIBUTES) GetIceDamage() int32 {
	if m != nil && m.IceDamage != nil {
		return *m.IceDamage
	}
	return Default_ATTRIBUTES_IceDamage
}

func (m *ATTRIBUTES) GetIceDefense() int32 {
	if m != nil && m.IceDefense != nil {
		return *m.IceDefense
	}
	return Default_ATTRIBUTES_IceDefense
}

func (m *ATTRIBUTES) GetIceBonus() float32 {
	if m != nil && m.IceBonus != nil {
		return *m.IceBonus
	}
	return Default_ATTRIBUTES_IceBonus
}

func (m *ATTRIBUTES) GetIceResist() float32 {
	if m != nil && m.IceResist != nil {
		return *m.IceResist
	}
	return Default_ATTRIBUTES_IceResist
}

func (m *ATTRIBUTES) GetFireDamage() int32 {
	if m != nil && m.FireDamage != nil {
		return *m.FireDamage
	}
	return Default_ATTRIBUTES_FireDamage
}

func (m *ATTRIBUTES) GetFireDefense() int32 {
	if m != nil && m.FireDefense != nil {
		return *m.FireDefense
	}
	return Default_ATTRIBUTES_FireDefense
}

func (m *ATTRIBUTES) GetFireBonus() float32 {
	if m != nil && m.FireBonus != nil {
		return *m.FireBonus
	}
	return Default_ATTRIBUTES_FireBonus
}

func (m *ATTRIBUTES) GetFireResist() float32 {
	if m != nil && m.FireResist != nil {
		return *m.FireResist
	}
	return Default_ATTRIBUTES_FireResist
}

func (m *ATTRIBUTES) GetLightingDamage() int32 {
	if m != nil && m.LightingDamage != nil {
		return *m.LightingDamage
	}
	return Default_ATTRIBUTES_LightingDamage
}

func (m *ATTRIBUTES) GetLightingDefense() int32 {
	if m != nil && m.LightingDefense != nil {
		return *m.LightingDefense
	}
	return Default_ATTRIBUTES_LightingDefense
}

func (m *ATTRIBUTES) GetLightingBonus() float32 {
	if m != nil && m.LightingBonus != nil {
		return *m.LightingBonus
	}
	return Default_ATTRIBUTES_LightingBonus
}

func (m *ATTRIBUTES) GetLightingResist() float32 {
	if m != nil && m.LightingResist != nil {
		return *m.LightingResist
	}
	return Default_ATTRIBUTES_LightingResist
}

func (m *ATTRIBUTES) GetPoisonDamage() int32 {
	if m != nil && m.PoisonDamage != nil {
		return *m.PoisonDamage
	}
	return Default_ATTRIBUTES_PoisonDamage
}

func (m *ATTRIBUTES) GetPoisonDefense() int32 {
	if m != nil && m.PoisonDefense != nil {
		return *m.PoisonDefense
	}
	return Default_ATTRIBUTES_PoisonDefense
}

func (m *ATTRIBUTES) GetPoisonBonus() float32 {
	if m != nil && m.PoisonBonus != nil {
		return *m.PoisonBonus
	}
	return Default_ATTRIBUTES_PoisonBonus
}

func (m *ATTRIBUTES) GetPoisonResist() float32 {
	if m != nil && m.PoisonResist != nil {
		return *m.PoisonResist
	}
	return Default_ATTRIBUTES_PoisonResist
}

func (m *ATTRIBUTES) GetComboRate() float32 {
	if m != nil && m.ComboRate != nil {
		return *m.ComboRate
	}
	return Default_ATTRIBUTES_ComboRate
}

type ATTRIBUTES_ARRAY struct {
	Items            []*ATTRIBUTES `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ATTRIBUTES_ARRAY) Reset()         { *m = ATTRIBUTES_ARRAY{} }
func (m *ATTRIBUTES_ARRAY) String() string { return proto.CompactTextString(m) }
func (*ATTRIBUTES_ARRAY) ProtoMessage()    {}

func (m *ATTRIBUTES_ARRAY) GetItems() []*ATTRIBUTES {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
