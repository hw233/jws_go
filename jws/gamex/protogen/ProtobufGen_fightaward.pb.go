// Code generated by protoc-gen-go.
// source: ProtobufGen_fightaward.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type FIGHTAWARD struct {
	// * 战力
	Fight               *uint32                  `protobuf:"varint,1,req,def=0" json:"Fight,omitempty"`
	AwardItems_Template []*FIGHTAWARD_AwardItems `protobuf:"bytes,2,rep" json:"AwardItems_Template,omitempty"`
	XXX_unrecognized    []byte                   `json:"-"`
}

func (m *FIGHTAWARD) Reset()         { *m = FIGHTAWARD{} }
func (m *FIGHTAWARD) String() string { return proto.CompactTextString(m) }
func (*FIGHTAWARD) ProtoMessage()    {}

const Default_FIGHTAWARD_Fight uint32 = 0

func (m *FIGHTAWARD) GetFight() uint32 {
	if m != nil && m.Fight != nil {
		return *m.Fight
	}
	return Default_FIGHTAWARD_Fight
}

func (m *FIGHTAWARD) GetAwardItems_Template() []*FIGHTAWARD_AwardItems {
	if m != nil {
		return m.AwardItems_Template
	}
	return nil
}

type FIGHTAWARD_AwardItems struct {
	// * 物品编号
	ItemID *string `protobuf:"bytes,1,opt,def=" json:"ItemID,omitempty"`
	// * 奖励数量
	Amount           *uint32 `protobuf:"varint,2,opt,def=0" json:"Amount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FIGHTAWARD_AwardItems) Reset()         { *m = FIGHTAWARD_AwardItems{} }
func (m *FIGHTAWARD_AwardItems) String() string { return proto.CompactTextString(m) }
func (*FIGHTAWARD_AwardItems) ProtoMessage()    {}

const Default_FIGHTAWARD_AwardItems_Amount uint32 = 0

func (m *FIGHTAWARD_AwardItems) GetItemID() string {
	if m != nil && m.ItemID != nil {
		return *m.ItemID
	}
	return ""
}

func (m *FIGHTAWARD_AwardItems) GetAmount() uint32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return Default_FIGHTAWARD_AwardItems_Amount
}

type FIGHTAWARD_ARRAY struct {
	Items            []*FIGHTAWARD `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *FIGHTAWARD_ARRAY) Reset()         { *m = FIGHTAWARD_ARRAY{} }
func (m *FIGHTAWARD_ARRAY) String() string { return proto.CompactTextString(m) }
func (*FIGHTAWARD_ARRAY) ProtoMessage()    {}

func (m *FIGHTAWARD_ARRAY) GetItems() []*FIGHTAWARD {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
