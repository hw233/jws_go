// Code generated by protoc-gen-go.
// source: ProtobufGen_guildactivity.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GUILDACTIVITY struct {
	// * 公会活动ID
	GuildActivityId *uint32 `protobuf:"varint,1,req,def=0" json:"GuildActivityId,omitempty"`
	// * 列表排序序号
	SortingNum *uint32 `protobuf:"varint,2,req,def=0" json:"SortingNum,omitempty"`
	// * 是否已开放
	IfActive *uint32 `protobuf:"varint,3,req,def=0" json:"IfActive,omitempty"`
	// * 军团活动名称IDS
	Name *string `protobuf:"bytes,4,req,def=" json:"Name,omitempty"`
	// * 军团活动图标
	Icon *string `protobuf:"bytes,5,req,def=" json:"Icon,omitempty"`
	// * 军团活动摘要IDS
	Tip *string `protobuf:"bytes,8,opt,def=" json:"Tip,omitempty"`
	// * 军团等级需求
	GuildLevelReq      *uint32                    `protobuf:"varint,6,req,def=0" json:"GuildLevelReq,omitempty"`
	RewardTip_Template []*GUILDACTIVITY_RewardTip `protobuf:"bytes,7,rep" json:"RewardTip_Template,omitempty"`
	XXX_unrecognized   []byte                     `json:"-"`
}

func (m *GUILDACTIVITY) Reset()         { *m = GUILDACTIVITY{} }
func (m *GUILDACTIVITY) String() string { return proto.CompactTextString(m) }
func (*GUILDACTIVITY) ProtoMessage()    {}

const Default_GUILDACTIVITY_GuildActivityId uint32 = 0
const Default_GUILDACTIVITY_SortingNum uint32 = 0
const Default_GUILDACTIVITY_IfActive uint32 = 0
const Default_GUILDACTIVITY_GuildLevelReq uint32 = 0

func (m *GUILDACTIVITY) GetGuildActivityId() uint32 {
	if m != nil && m.GuildActivityId != nil {
		return *m.GuildActivityId
	}
	return Default_GUILDACTIVITY_GuildActivityId
}

func (m *GUILDACTIVITY) GetSortingNum() uint32 {
	if m != nil && m.SortingNum != nil {
		return *m.SortingNum
	}
	return Default_GUILDACTIVITY_SortingNum
}

func (m *GUILDACTIVITY) GetIfActive() uint32 {
	if m != nil && m.IfActive != nil {
		return *m.IfActive
	}
	return Default_GUILDACTIVITY_IfActive
}

func (m *GUILDACTIVITY) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GUILDACTIVITY) GetIcon() string {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return ""
}

func (m *GUILDACTIVITY) GetTip() string {
	if m != nil && m.Tip != nil {
		return *m.Tip
	}
	return ""
}

func (m *GUILDACTIVITY) GetGuildLevelReq() uint32 {
	if m != nil && m.GuildLevelReq != nil {
		return *m.GuildLevelReq
	}
	return Default_GUILDACTIVITY_GuildLevelReq
}

func (m *GUILDACTIVITY) GetRewardTip_Template() []*GUILDACTIVITY_RewardTip {
	if m != nil {
		return m.RewardTip_Template
	}
	return nil
}

type GUILDACTIVITY_RewardTip struct {
	// * 活动奖励名称
	RewardTitle *string `protobuf:"bytes,1,opt,def=" json:"RewardTitle,omitempty"`
	// * 活动奖励内容
	RewardDesc       *string `protobuf:"bytes,2,opt,def=" json:"RewardDesc,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GUILDACTIVITY_RewardTip) Reset()         { *m = GUILDACTIVITY_RewardTip{} }
func (m *GUILDACTIVITY_RewardTip) String() string { return proto.CompactTextString(m) }
func (*GUILDACTIVITY_RewardTip) ProtoMessage()    {}

func (m *GUILDACTIVITY_RewardTip) GetRewardTitle() string {
	if m != nil && m.RewardTitle != nil {
		return *m.RewardTitle
	}
	return ""
}

func (m *GUILDACTIVITY_RewardTip) GetRewardDesc() string {
	if m != nil && m.RewardDesc != nil {
		return *m.RewardDesc
	}
	return ""
}

type GUILDACTIVITY_ARRAY struct {
	Items            []*GUILDACTIVITY `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GUILDACTIVITY_ARRAY) Reset()         { *m = GUILDACTIVITY_ARRAY{} }
func (m *GUILDACTIVITY_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GUILDACTIVITY_ARRAY) ProtoMessage()    {}

func (m *GUILDACTIVITY_ARRAY) GetItems() []*GUILDACTIVITY {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
