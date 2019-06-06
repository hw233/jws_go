// Code generated by protoc-gen-go.
// source: ProtobufGen_materialenhance.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MATERIALENHANCE struct {
	// * 索引
	Index *string `protobuf:"bytes,1,req,def=" json:"Index,omitempty"`
	// * 进化等级
	EnhanceLevel *uint32 `protobuf:"varint,2,req,def=0" json:"EnhanceLevel,omitempty"`
	// * 部位
	Part *string `protobuf:"bytes,3,req,def=" json:"Part,omitempty"`
	// * 精炼SC
	SC *uint32 `protobuf:"varint,4,opt,def=0" json:"SC,omitempty"`
	// * 攻击力
	TotalATK *float32 `protobuf:"fixed32,5,opt,def=0" json:"TotalATK,omitempty"`
	// * 防御力
	TotalDEF *float32 `protobuf:"fixed32,6,opt,def=0" json:"TotalDEF,omitempty"`
	// * 生命值
	TotalHP          *float32                     `protobuf:"fixed32,7,opt,def=0" json:"TotalHP,omitempty"`
	Materials_Table  []*MATERIALENHANCE_Materials `protobuf:"bytes,8,rep" json:"Materials_Table,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *MATERIALENHANCE) Reset()         { *m = MATERIALENHANCE{} }
func (m *MATERIALENHANCE) String() string { return proto.CompactTextString(m) }
func (*MATERIALENHANCE) ProtoMessage()    {}

const Default_MATERIALENHANCE_EnhanceLevel uint32 = 0
const Default_MATERIALENHANCE_SC uint32 = 0
const Default_MATERIALENHANCE_TotalATK float32 = 0
const Default_MATERIALENHANCE_TotalDEF float32 = 0
const Default_MATERIALENHANCE_TotalHP float32 = 0

func (m *MATERIALENHANCE) GetIndex() string {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return ""
}

func (m *MATERIALENHANCE) GetEnhanceLevel() uint32 {
	if m != nil && m.EnhanceLevel != nil {
		return *m.EnhanceLevel
	}
	return Default_MATERIALENHANCE_EnhanceLevel
}

func (m *MATERIALENHANCE) GetPart() string {
	if m != nil && m.Part != nil {
		return *m.Part
	}
	return ""
}

func (m *MATERIALENHANCE) GetSC() uint32 {
	if m != nil && m.SC != nil {
		return *m.SC
	}
	return Default_MATERIALENHANCE_SC
}

func (m *MATERIALENHANCE) GetTotalATK() float32 {
	if m != nil && m.TotalATK != nil {
		return *m.TotalATK
	}
	return Default_MATERIALENHANCE_TotalATK
}

func (m *MATERIALENHANCE) GetTotalDEF() float32 {
	if m != nil && m.TotalDEF != nil {
		return *m.TotalDEF
	}
	return Default_MATERIALENHANCE_TotalDEF
}

func (m *MATERIALENHANCE) GetTotalHP() float32 {
	if m != nil && m.TotalHP != nil {
		return *m.TotalHP
	}
	return Default_MATERIALENHANCE_TotalHP
}

func (m *MATERIALENHANCE) GetMaterials_Table() []*MATERIALENHANCE_Materials {
	if m != nil {
		return m.Materials_Table
	}
	return nil
}

type MATERIALENHANCE_Materials struct {
	// * 材料ID
	MaterialsID *string `protobuf:"bytes,1,opt,def=" json:"MaterialsID,omitempty"`
	// * 材料数量
	MaterialsCount *uint32 `protobuf:"varint,2,opt,def=0" json:"MaterialsCount,omitempty"`
	// * 攻击力
	ATK *float32 `protobuf:"fixed32,3,opt,def=0" json:"ATK,omitempty"`
	// * 防御力
	DEF *float32 `protobuf:"fixed32,4,opt,def=0" json:"DEF,omitempty"`
	// * 生命值
	HP               *float32 `protobuf:"fixed32,5,opt,def=0" json:"HP,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MATERIALENHANCE_Materials) Reset()         { *m = MATERIALENHANCE_Materials{} }
func (m *MATERIALENHANCE_Materials) String() string { return proto.CompactTextString(m) }
func (*MATERIALENHANCE_Materials) ProtoMessage()    {}

const Default_MATERIALENHANCE_Materials_MaterialsCount uint32 = 0
const Default_MATERIALENHANCE_Materials_ATK float32 = 0
const Default_MATERIALENHANCE_Materials_DEF float32 = 0
const Default_MATERIALENHANCE_Materials_HP float32 = 0

func (m *MATERIALENHANCE_Materials) GetMaterialsID() string {
	if m != nil && m.MaterialsID != nil {
		return *m.MaterialsID
	}
	return ""
}

func (m *MATERIALENHANCE_Materials) GetMaterialsCount() uint32 {
	if m != nil && m.MaterialsCount != nil {
		return *m.MaterialsCount
	}
	return Default_MATERIALENHANCE_Materials_MaterialsCount
}

func (m *MATERIALENHANCE_Materials) GetATK() float32 {
	if m != nil && m.ATK != nil {
		return *m.ATK
	}
	return Default_MATERIALENHANCE_Materials_ATK
}

func (m *MATERIALENHANCE_Materials) GetDEF() float32 {
	if m != nil && m.DEF != nil {
		return *m.DEF
	}
	return Default_MATERIALENHANCE_Materials_DEF
}

func (m *MATERIALENHANCE_Materials) GetHP() float32 {
	if m != nil && m.HP != nil {
		return *m.HP
	}
	return Default_MATERIALENHANCE_Materials_HP
}

type MATERIALENHANCE_ARRAY struct {
	Items            []*MATERIALENHANCE `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *MATERIALENHANCE_ARRAY) Reset()         { *m = MATERIALENHANCE_ARRAY{} }
func (m *MATERIALENHANCE_ARRAY) String() string { return proto.CompactTextString(m) }
func (*MATERIALENHANCE_ARRAY) ProtoMessage()    {}

func (m *MATERIALENHANCE_ARRAY) GetItems() []*MATERIALENHANCE {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}