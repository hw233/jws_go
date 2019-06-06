// Code generated by protoc-gen-go.
// source: ProtobufGen_dailygift.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type DAILYGIFT struct {
	// * 索引
	Index *uint32 `protobuf:"varint,1,req,def=0" json:"Index,omitempty"`
	// * 活动ID
	ActivityID *uint32 `protobuf:"varint,2,req,def=0" json:"ActivityID,omitempty"`
	// * 活动天数索引
	ActivityIndex *uint32 `protobuf:"varint,3,opt,def=0" json:"ActivityIndex,omitempty"`
	// * 角色限制
	RoleLimit *uint32 `protobuf:"varint,4,opt,def=0" json:"RoleLimit,omitempty"`
	// * 是否显示流光特效
	GoodsEffects *uint32 `protobuf:"varint,15,opt,def=0" json:"GoodsEffects,omitempty"`
	// * 奖励ID
	AWardID *string `protobuf:"bytes,5,opt,def=" json:"AWardID,omitempty"`
	// * 奖励数量
	Count *uint32 `protobuf:"varint,6,opt,def=0" json:"Count,omitempty"`
	// * 奖励ID
	AWardID2 *string `protobuf:"bytes,7,opt,def=" json:"AWardID2,omitempty"`
	// * 奖励数量
	Count2 *uint32 `protobuf:"varint,8,opt,def=0" json:"Count2,omitempty"`
	// * 奖励ID
	AWardID3 *string `protobuf:"bytes,9,opt,def=" json:"AWardID3,omitempty"`
	// * 奖励数量
	Count3 *uint32 `protobuf:"varint,10,opt,def=0" json:"Count3,omitempty"`
	// * 奖励ID
	AWardID4 *string `protobuf:"bytes,11,opt,def=" json:"AWardID4,omitempty"`
	// * 奖励数量
	Count4 *uint32 `protobuf:"varint,12,opt,def=0" json:"Count4,omitempty"`
	// * VIP等级下限
	VIPBase *uint32 `protobuf:"varint,13,opt,def=0" json:"VIPBase,omitempty"`
	// * VIP数量倍数
	VIPBonus         *float32 `protobuf:"fixed32,14,opt,def=1" json:"VIPBonus,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DAILYGIFT) Reset()         { *m = DAILYGIFT{} }
func (m *DAILYGIFT) String() string { return proto.CompactTextString(m) }
func (*DAILYGIFT) ProtoMessage()    {}

const Default_DAILYGIFT_Index uint32 = 0
const Default_DAILYGIFT_ActivityID uint32 = 0
const Default_DAILYGIFT_ActivityIndex uint32 = 0
const Default_DAILYGIFT_RoleLimit uint32 = 0
const Default_DAILYGIFT_GoodsEffects uint32 = 0
const Default_DAILYGIFT_Count uint32 = 0
const Default_DAILYGIFT_Count2 uint32 = 0
const Default_DAILYGIFT_Count3 uint32 = 0
const Default_DAILYGIFT_Count4 uint32 = 0
const Default_DAILYGIFT_VIPBase uint32 = 0
const Default_DAILYGIFT_VIPBonus float32 = 1

func (m *DAILYGIFT) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return Default_DAILYGIFT_Index
}

func (m *DAILYGIFT) GetActivityID() uint32 {
	if m != nil && m.ActivityID != nil {
		return *m.ActivityID
	}
	return Default_DAILYGIFT_ActivityID
}

func (m *DAILYGIFT) GetActivityIndex() uint32 {
	if m != nil && m.ActivityIndex != nil {
		return *m.ActivityIndex
	}
	return Default_DAILYGIFT_ActivityIndex
}

func (m *DAILYGIFT) GetRoleLimit() uint32 {
	if m != nil && m.RoleLimit != nil {
		return *m.RoleLimit
	}
	return Default_DAILYGIFT_RoleLimit
}

func (m *DAILYGIFT) GetGoodsEffects() uint32 {
	if m != nil && m.GoodsEffects != nil {
		return *m.GoodsEffects
	}
	return Default_DAILYGIFT_GoodsEffects
}

func (m *DAILYGIFT) GetAWardID() string {
	if m != nil && m.AWardID != nil {
		return *m.AWardID
	}
	return ""
}

func (m *DAILYGIFT) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return Default_DAILYGIFT_Count
}

func (m *DAILYGIFT) GetAWardID2() string {
	if m != nil && m.AWardID2 != nil {
		return *m.AWardID2
	}
	return ""
}

func (m *DAILYGIFT) GetCount2() uint32 {
	if m != nil && m.Count2 != nil {
		return *m.Count2
	}
	return Default_DAILYGIFT_Count2
}

func (m *DAILYGIFT) GetAWardID3() string {
	if m != nil && m.AWardID3 != nil {
		return *m.AWardID3
	}
	return ""
}

func (m *DAILYGIFT) GetCount3() uint32 {
	if m != nil && m.Count3 != nil {
		return *m.Count3
	}
	return Default_DAILYGIFT_Count3
}

func (m *DAILYGIFT) GetAWardID4() string {
	if m != nil && m.AWardID4 != nil {
		return *m.AWardID4
	}
	return ""
}

func (m *DAILYGIFT) GetCount4() uint32 {
	if m != nil && m.Count4 != nil {
		return *m.Count4
	}
	return Default_DAILYGIFT_Count4
}

func (m *DAILYGIFT) GetVIPBase() uint32 {
	if m != nil && m.VIPBase != nil {
		return *m.VIPBase
	}
	return Default_DAILYGIFT_VIPBase
}

func (m *DAILYGIFT) GetVIPBonus() float32 {
	if m != nil && m.VIPBonus != nil {
		return *m.VIPBonus
	}
	return Default_DAILYGIFT_VIPBonus
}

type DAILYGIFT_ARRAY struct {
	Items            []*DAILYGIFT `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DAILYGIFT_ARRAY) Reset()         { *m = DAILYGIFT_ARRAY{} }
func (m *DAILYGIFT_ARRAY) String() string { return proto.CompactTextString(m) }
func (*DAILYGIFT_ARRAY) ProtoMessage()    {}

func (m *DAILYGIFT_ARRAY) GetItems() []*DAILYGIFT {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}