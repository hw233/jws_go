// Code generated by protoc-gen-go.
// source: ProtobufGen_roleweight.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ROLEWEIGHT struct {
	// * 出战角色权重
	RoleFightWeight *uint32 `protobuf:"varint,1,req,def=0" json:"RoleFightWeight,omitempty"`
	// * 不出战角色权重
	RoleRestWeight   *uint32 `protobuf:"varint,2,req,def=0" json:"RoleRestWeight,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ROLEWEIGHT) Reset()         { *m = ROLEWEIGHT{} }
func (m *ROLEWEIGHT) String() string { return proto.CompactTextString(m) }
func (*ROLEWEIGHT) ProtoMessage()    {}

const Default_ROLEWEIGHT_RoleFightWeight uint32 = 0
const Default_ROLEWEIGHT_RoleRestWeight uint32 = 0

func (m *ROLEWEIGHT) GetRoleFightWeight() uint32 {
	if m != nil && m.RoleFightWeight != nil {
		return *m.RoleFightWeight
	}
	return Default_ROLEWEIGHT_RoleFightWeight
}

func (m *ROLEWEIGHT) GetRoleRestWeight() uint32 {
	if m != nil && m.RoleRestWeight != nil {
		return *m.RoleRestWeight
	}
	return Default_ROLEWEIGHT_RoleRestWeight
}

type ROLEWEIGHT_ARRAY struct {
	Items            []*ROLEWEIGHT `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ROLEWEIGHT_ARRAY) Reset()         { *m = ROLEWEIGHT_ARRAY{} }
func (m *ROLEWEIGHT_ARRAY) String() string { return proto.CompactTextString(m) }
func (*ROLEWEIGHT_ARRAY) ProtoMessage()    {}

func (m *ROLEWEIGHT_ARRAY) GetItems() []*ROLEWEIGHT {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}