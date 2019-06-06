// Code generated by protoc-gen-go.
// source: ProtobufGen_herowingstar.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type HEROWINGSTAR struct {
	// * 翅膀星级
	HWStar *uint32 `protobuf:"varint,1,req,def=0" json:"HWStar,omitempty"`
	// * 攻击
	ATK *float32 `protobuf:"fixed32,2,opt,def=0" json:"ATK,omitempty"`
	// * 防御
	DEF *float32 `protobuf:"fixed32,3,opt,def=0" json:"DEF,omitempty"`
	// * 生命
	HP *float32 `protobuf:"fixed32,4,opt,def=0" json:"HP,omitempty"`
	// * 回退所需的钻石数量
	RebornCost        *uint32                  `protobuf:"varint,6,opt,def=0" json:"RebornCost,omitempty"`
	HWStarup_Template []*HEROWINGSTAR_HWStarup `protobuf:"bytes,5,rep" json:"HWStarup_Template,omitempty"`
	XXX_unrecognized  []byte                   `json:"-"`
}

func (m *HEROWINGSTAR) Reset()         { *m = HEROWINGSTAR{} }
func (m *HEROWINGSTAR) String() string { return proto.CompactTextString(m) }
func (*HEROWINGSTAR) ProtoMessage()    {}

const Default_HEROWINGSTAR_HWStar uint32 = 0
const Default_HEROWINGSTAR_ATK float32 = 0
const Default_HEROWINGSTAR_DEF float32 = 0
const Default_HEROWINGSTAR_HP float32 = 0
const Default_HEROWINGSTAR_RebornCost uint32 = 0

func (m *HEROWINGSTAR) GetHWStar() uint32 {
	if m != nil && m.HWStar != nil {
		return *m.HWStar
	}
	return Default_HEROWINGSTAR_HWStar
}

func (m *HEROWINGSTAR) GetATK() float32 {
	if m != nil && m.ATK != nil {
		return *m.ATK
	}
	return Default_HEROWINGSTAR_ATK
}

func (m *HEROWINGSTAR) GetDEF() float32 {
	if m != nil && m.DEF != nil {
		return *m.DEF
	}
	return Default_HEROWINGSTAR_DEF
}

func (m *HEROWINGSTAR) GetHP() float32 {
	if m != nil && m.HP != nil {
		return *m.HP
	}
	return Default_HEROWINGSTAR_HP
}

func (m *HEROWINGSTAR) GetRebornCost() uint32 {
	if m != nil && m.RebornCost != nil {
		return *m.RebornCost
	}
	return Default_HEROWINGSTAR_RebornCost
}

func (m *HEROWINGSTAR) GetHWStarup_Template() []*HEROWINGSTAR_HWStarup {
	if m != nil {
		return m.HWStarup_Template
	}
	return nil
}

type HEROWINGSTAR_HWStarup struct {
	// * 升星材料
	HWStarupMaterial *string `protobuf:"bytes,1,opt,def=" json:"HWStarupMaterial,omitempty"`
	// * 升星材料数量
	HWStarupMaterialCount *uint32 `protobuf:"varint,2,opt,def=1" json:"HWStarupMaterialCount,omitempty"`
	XXX_unrecognized      []byte  `json:"-"`
}

func (m *HEROWINGSTAR_HWStarup) Reset()         { *m = HEROWINGSTAR_HWStarup{} }
func (m *HEROWINGSTAR_HWStarup) String() string { return proto.CompactTextString(m) }
func (*HEROWINGSTAR_HWStarup) ProtoMessage()    {}

const Default_HEROWINGSTAR_HWStarup_HWStarupMaterialCount uint32 = 1

func (m *HEROWINGSTAR_HWStarup) GetHWStarupMaterial() string {
	if m != nil && m.HWStarupMaterial != nil {
		return *m.HWStarupMaterial
	}
	return ""
}

func (m *HEROWINGSTAR_HWStarup) GetHWStarupMaterialCount() uint32 {
	if m != nil && m.HWStarupMaterialCount != nil {
		return *m.HWStarupMaterialCount
	}
	return Default_HEROWINGSTAR_HWStarup_HWStarupMaterialCount
}

type HEROWINGSTAR_ARRAY struct {
	Items            []*HEROWINGSTAR `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *HEROWINGSTAR_ARRAY) Reset()         { *m = HEROWINGSTAR_ARRAY{} }
func (m *HEROWINGSTAR_ARRAY) String() string { return proto.CompactTextString(m) }
func (*HEROWINGSTAR_ARRAY) ProtoMessage()    {}

func (m *HEROWINGSTAR_ARRAY) GetItems() []*HEROWINGSTAR {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}