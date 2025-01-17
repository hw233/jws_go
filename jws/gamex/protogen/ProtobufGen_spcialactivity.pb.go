// Code generated by protoc-gen-go.
// source: ProtobufGen_spcialactivity.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SPCIALACTIVITY struct {
	// * 特殊活动ID,0=手机绑定活动
	SpecialActivityID  *uint32                     `protobuf:"varint,1,req,def=0" json:"SpecialActivityID,omitempty"`
	GoalAward_Template []*SPCIALACTIVITY_GoalAward `protobuf:"bytes,2,rep" json:"GoalAward_Template,omitempty"`
	XXX_unrecognized   []byte                      `json:"-"`
}

func (m *SPCIALACTIVITY) Reset()         { *m = SPCIALACTIVITY{} }
func (m *SPCIALACTIVITY) String() string { return proto.CompactTextString(m) }
func (*SPCIALACTIVITY) ProtoMessage()    {}

const Default_SPCIALACTIVITY_SpecialActivityID uint32 = 0

func (m *SPCIALACTIVITY) GetSpecialActivityID() uint32 {
	if m != nil && m.SpecialActivityID != nil {
		return *m.SpecialActivityID
	}
	return Default_SPCIALACTIVITY_SpecialActivityID
}

func (m *SPCIALACTIVITY) GetGoalAward_Template() []*SPCIALACTIVITY_GoalAward {
	if m != nil {
		return m.GoalAward_Template
	}
	return nil
}

type SPCIALACTIVITY_GoalAward struct {
	// * 奖励ID
	Reward *string `protobuf:"bytes,1,opt,def=" json:"Reward,omitempty"`
	// * 奖励数目
	Count            *uint32 `protobuf:"varint,2,opt,def=0" json:"Count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SPCIALACTIVITY_GoalAward) Reset()         { *m = SPCIALACTIVITY_GoalAward{} }
func (m *SPCIALACTIVITY_GoalAward) String() string { return proto.CompactTextString(m) }
func (*SPCIALACTIVITY_GoalAward) ProtoMessage()    {}

const Default_SPCIALACTIVITY_GoalAward_Count uint32 = 0

func (m *SPCIALACTIVITY_GoalAward) GetReward() string {
	if m != nil && m.Reward != nil {
		return *m.Reward
	}
	return ""
}

func (m *SPCIALACTIVITY_GoalAward) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return Default_SPCIALACTIVITY_GoalAward_Count
}

type SPCIALACTIVITY_ARRAY struct {
	Items            []*SPCIALACTIVITY `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SPCIALACTIVITY_ARRAY) Reset()         { *m = SPCIALACTIVITY_ARRAY{} }
func (m *SPCIALACTIVITY_ARRAY) String() string { return proto.CompactTextString(m) }
func (*SPCIALACTIVITY_ARRAY) ProtoMessage()    {}

func (m *SPCIALACTIVITY_ARRAY) GetItems() []*SPCIALACTIVITY {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
