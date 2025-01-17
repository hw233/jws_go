// Code generated by protoc-gen-go.
// source: ProtobufGen_template.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type TEMPLATE struct {
	// * 掉落组ID
	ID               *string                 `protobuf:"bytes,1,req,def=" json:"ID,omitempty"`
	Rules            []*TEMPLATE_GroupWeight `protobuf:"bytes,2,rep" json:"Rules,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *TEMPLATE) Reset()         { *m = TEMPLATE{} }
func (m *TEMPLATE) String() string { return proto.CompactTextString(m) }
func (*TEMPLATE) ProtoMessage()    {}

func (m *TEMPLATE) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *TEMPLATE) GetRules() []*TEMPLATE_GroupWeight {
	if m != nil {
		return m.Rules
	}
	return nil
}

type TEMPLATE_GroupWeight struct {
	// * 物品组ID
	ItemGroupID *string `protobuf:"bytes,1,opt,def=" json:"ItemGroupID,omitempty"`
	// * 物品组权重
	Weight           *int32 `protobuf:"varint,2,opt,def=0" json:"Weight,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TEMPLATE_GroupWeight) Reset()         { *m = TEMPLATE_GroupWeight{} }
func (m *TEMPLATE_GroupWeight) String() string { return proto.CompactTextString(m) }
func (*TEMPLATE_GroupWeight) ProtoMessage()    {}

const Default_TEMPLATE_GroupWeight_Weight int32 = 0

func (m *TEMPLATE_GroupWeight) GetItemGroupID() string {
	if m != nil && m.ItemGroupID != nil {
		return *m.ItemGroupID
	}
	return ""
}

func (m *TEMPLATE_GroupWeight) GetWeight() int32 {
	if m != nil && m.Weight != nil {
		return *m.Weight
	}
	return Default_TEMPLATE_GroupWeight_Weight
}

type TEMPLATE_ARRAY struct {
	Items            []*TEMPLATE `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TEMPLATE_ARRAY) Reset()         { *m = TEMPLATE_ARRAY{} }
func (m *TEMPLATE_ARRAY) String() string { return proto.CompactTextString(m) }
func (*TEMPLATE_ARRAY) ProtoMessage()    {}

func (m *TEMPLATE_ARRAY) GetItems() []*TEMPLATE {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
