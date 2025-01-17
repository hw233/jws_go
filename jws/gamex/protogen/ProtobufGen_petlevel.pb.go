// Code generated by protoc-gen-go.
// source: ProtobufGen_petlevel.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type PETLEVEL struct {
	// * 灵宠的ID
	PetID *uint32 `protobuf:"varint,6,req,def=0" json:"PetID,omitempty"`
	// * 灵宠等级
	Petlevel *uint32 `protobuf:"varint,1,req,def=0" json:"Petlevel,omitempty"`
	// * 攻击提升
	PetATK *uint32 `protobuf:"varint,2,opt,def=0" json:"PetATK,omitempty"`
	// * 防御提升
	PetDEF *uint32 `protobuf:"varint,3,opt,def=0" json:"PetDEF,omitempty"`
	// * 生命提升
	PetHP            *uint32              `protobuf:"varint,4,opt,def=0" json:"PetHP,omitempty"`
	Material_Table   []*PETLEVEL_Material `protobuf:"bytes,5,rep" json:"Material_Table,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PETLEVEL) Reset()         { *m = PETLEVEL{} }
func (m *PETLEVEL) String() string { return proto.CompactTextString(m) }
func (*PETLEVEL) ProtoMessage()    {}

const Default_PETLEVEL_PetID uint32 = 0
const Default_PETLEVEL_Petlevel uint32 = 0
const Default_PETLEVEL_PetATK uint32 = 0
const Default_PETLEVEL_PetDEF uint32 = 0
const Default_PETLEVEL_PetHP uint32 = 0

func (m *PETLEVEL) GetPetID() uint32 {
	if m != nil && m.PetID != nil {
		return *m.PetID
	}
	return Default_PETLEVEL_PetID
}

func (m *PETLEVEL) GetPetlevel() uint32 {
	if m != nil && m.Petlevel != nil {
		return *m.Petlevel
	}
	return Default_PETLEVEL_Petlevel
}

func (m *PETLEVEL) GetPetATK() uint32 {
	if m != nil && m.PetATK != nil {
		return *m.PetATK
	}
	return Default_PETLEVEL_PetATK
}

func (m *PETLEVEL) GetPetDEF() uint32 {
	if m != nil && m.PetDEF != nil {
		return *m.PetDEF
	}
	return Default_PETLEVEL_PetDEF
}

func (m *PETLEVEL) GetPetHP() uint32 {
	if m != nil && m.PetHP != nil {
		return *m.PetHP
	}
	return Default_PETLEVEL_PetHP
}

func (m *PETLEVEL) GetMaterial_Table() []*PETLEVEL_Material {
	if m != nil {
		return m.Material_Table
	}
	return nil
}

type PETLEVEL_Material struct {
	// * 消耗材料类型
	MaterialID *string `protobuf:"bytes,1,opt,def=" json:"MaterialID,omitempty"`
	// * 消耗材料数量
	MaterialAmount   *uint32 `protobuf:"varint,2,opt,def=0" json:"MaterialAmount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PETLEVEL_Material) Reset()         { *m = PETLEVEL_Material{} }
func (m *PETLEVEL_Material) String() string { return proto.CompactTextString(m) }
func (*PETLEVEL_Material) ProtoMessage()    {}

const Default_PETLEVEL_Material_MaterialAmount uint32 = 0

func (m *PETLEVEL_Material) GetMaterialID() string {
	if m != nil && m.MaterialID != nil {
		return *m.MaterialID
	}
	return ""
}

func (m *PETLEVEL_Material) GetMaterialAmount() uint32 {
	if m != nil && m.MaterialAmount != nil {
		return *m.MaterialAmount
	}
	return Default_PETLEVEL_Material_MaterialAmount
}

type PETLEVEL_ARRAY struct {
	Items            []*PETLEVEL `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PETLEVEL_ARRAY) Reset()         { *m = PETLEVEL_ARRAY{} }
func (m *PETLEVEL_ARRAY) String() string { return proto.CompactTextString(m) }
func (*PETLEVEL_ARRAY) ProtoMessage()    {}

func (m *PETLEVEL_ARRAY) GetItems() []*PETLEVEL {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
