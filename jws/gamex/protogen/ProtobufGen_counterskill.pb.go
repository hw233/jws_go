// Code generated by protoc-gen-go.
// source: ProtobufGen_counterskill.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type COUNTERSKILL struct {
	// *
	CounterSkillId *string `protobuf:"bytes,1,req,def=" json:"CounterSkillId,omitempty"`
	// *
	Name *string `protobuf:"bytes,2,opt,def=" json:"Name,omitempty"`
	// *
	Desc *string `protobuf:"bytes,3,opt,def=" json:"Desc,omitempty"`
	// *
	NextLevelDesc *string `protobuf:"bytes,4,opt,def=" json:"NextLevelDesc,omitempty"`
	// *
	Icon *string `protobuf:"bytes,13,opt,name=icon,def=" json:"icon,omitempty"`
	// *
	SmallIcon *string `protobuf:"bytes,21,opt,def=" json:"SmallIcon,omitempty"`
	// *
	Level *uint32 `protobuf:"varint,5,opt,def=0" json:"Level,omitempty"`
	// *
	Group *uint32 `protobuf:"varint,17,opt,def=0" json:"Group,omitempty"`
	// *
	EffectScope *uint32 `protobuf:"varint,18,opt,def=0" json:"EffectScope,omitempty"`
	// *
	TargetProp *string `protobuf:"bytes,6,opt,def=" json:"TargetProp,omitempty"`
	// *
	TargetPropValue *string `protobuf:"bytes,7,opt,def=" json:"TargetPropValue,omitempty"`
	// *
	SelfProp *string `protobuf:"bytes,19,opt,def=" json:"SelfProp,omitempty"`
	// *
	SlefPropValue *string `protobuf:"bytes,20,opt,def=" json:"SlefPropValue,omitempty"`
	// *
	ModifyType *string `protobuf:"bytes,8,opt,def=" json:"ModifyType,omitempty"`
	// *
	ModifyValueType *uint32 `protobuf:"varint,9,opt,def=0" json:"ModifyValueType,omitempty"`
	// *
	ModifyValue *float32 `protobuf:"fixed32,10,opt,def=0" json:"ModifyValue,omitempty"`
	// *
	CorrectionRatio *float32 `protobuf:"fixed32,11,opt,def=0" json:"CorrectionRatio,omitempty"`
	// *
	Counter3V3Target *uint32 `protobuf:"varint,12,opt,name=Counter3v3Target,def=0" json:"Counter3v3Target,omitempty"`
	// *
	ActivationItemId *string `protobuf:"bytes,14,opt,def=" json:"ActivationItemId,omitempty"`
	// *
	ActivationItemCount *float32 `protobuf:"fixed32,15,opt,def=0" json:"ActivationItemCount,omitempty"`
	// * 图标显示分组
	Show_Icon_Group  *int32 `protobuf:"varint,16,opt,def=0" json:"Show_Icon_Group,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *COUNTERSKILL) Reset()         { *m = COUNTERSKILL{} }
func (m *COUNTERSKILL) String() string { return proto.CompactTextString(m) }
func (*COUNTERSKILL) ProtoMessage()    {}

const Default_COUNTERSKILL_Level uint32 = 0
const Default_COUNTERSKILL_Group uint32 = 0
const Default_COUNTERSKILL_EffectScope uint32 = 0
const Default_COUNTERSKILL_ModifyValueType uint32 = 0
const Default_COUNTERSKILL_ModifyValue float32 = 0
const Default_COUNTERSKILL_CorrectionRatio float32 = 0
const Default_COUNTERSKILL_Counter3V3Target uint32 = 0
const Default_COUNTERSKILL_ActivationItemCount float32 = 0
const Default_COUNTERSKILL_Show_Icon_Group int32 = 0

func (m *COUNTERSKILL) GetCounterSkillId() string {
	if m != nil && m.CounterSkillId != nil {
		return *m.CounterSkillId
	}
	return ""
}

func (m *COUNTERSKILL) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *COUNTERSKILL) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

func (m *COUNTERSKILL) GetNextLevelDesc() string {
	if m != nil && m.NextLevelDesc != nil {
		return *m.NextLevelDesc
	}
	return ""
}

func (m *COUNTERSKILL) GetIcon() string {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return ""
}

func (m *COUNTERSKILL) GetSmallIcon() string {
	if m != nil && m.SmallIcon != nil {
		return *m.SmallIcon
	}
	return ""
}

func (m *COUNTERSKILL) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return Default_COUNTERSKILL_Level
}

func (m *COUNTERSKILL) GetGroup() uint32 {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return Default_COUNTERSKILL_Group
}

func (m *COUNTERSKILL) GetEffectScope() uint32 {
	if m != nil && m.EffectScope != nil {
		return *m.EffectScope
	}
	return Default_COUNTERSKILL_EffectScope
}

func (m *COUNTERSKILL) GetTargetProp() string {
	if m != nil && m.TargetProp != nil {
		return *m.TargetProp
	}
	return ""
}

func (m *COUNTERSKILL) GetTargetPropValue() string {
	if m != nil && m.TargetPropValue != nil {
		return *m.TargetPropValue
	}
	return ""
}

func (m *COUNTERSKILL) GetSelfProp() string {
	if m != nil && m.SelfProp != nil {
		return *m.SelfProp
	}
	return ""
}

func (m *COUNTERSKILL) GetSlefPropValue() string {
	if m != nil && m.SlefPropValue != nil {
		return *m.SlefPropValue
	}
	return ""
}

func (m *COUNTERSKILL) GetModifyType() string {
	if m != nil && m.ModifyType != nil {
		return *m.ModifyType
	}
	return ""
}

func (m *COUNTERSKILL) GetModifyValueType() uint32 {
	if m != nil && m.ModifyValueType != nil {
		return *m.ModifyValueType
	}
	return Default_COUNTERSKILL_ModifyValueType
}

func (m *COUNTERSKILL) GetModifyValue() float32 {
	if m != nil && m.ModifyValue != nil {
		return *m.ModifyValue
	}
	return Default_COUNTERSKILL_ModifyValue
}

func (m *COUNTERSKILL) GetCorrectionRatio() float32 {
	if m != nil && m.CorrectionRatio != nil {
		return *m.CorrectionRatio
	}
	return Default_COUNTERSKILL_CorrectionRatio
}

func (m *COUNTERSKILL) GetCounter3V3Target() uint32 {
	if m != nil && m.Counter3V3Target != nil {
		return *m.Counter3V3Target
	}
	return Default_COUNTERSKILL_Counter3V3Target
}

func (m *COUNTERSKILL) GetActivationItemId() string {
	if m != nil && m.ActivationItemId != nil {
		return *m.ActivationItemId
	}
	return ""
}

func (m *COUNTERSKILL) GetActivationItemCount() float32 {
	if m != nil && m.ActivationItemCount != nil {
		return *m.ActivationItemCount
	}
	return Default_COUNTERSKILL_ActivationItemCount
}

func (m *COUNTERSKILL) GetShow_Icon_Group() int32 {
	if m != nil && m.Show_Icon_Group != nil {
		return *m.Show_Icon_Group
	}
	return Default_COUNTERSKILL_Show_Icon_Group
}

type COUNTERSKILL_ARRAY struct {
	Items            []*COUNTERSKILL `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *COUNTERSKILL_ARRAY) Reset()         { *m = COUNTERSKILL_ARRAY{} }
func (m *COUNTERSKILL_ARRAY) String() string { return proto.CompactTextString(m) }
func (*COUNTERSKILL_ARRAY) ProtoMessage()    {}

func (m *COUNTERSKILL_ARRAY) GetItems() []*COUNTERSKILL {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
