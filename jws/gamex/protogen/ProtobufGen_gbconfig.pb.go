// Code generated by protoc-gen-go.
// source: ProtobufGen_gbconfig.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GBCONFIG struct {
	// * 小boss购买次数花费
	LittleBossCostHC *uint32 `protobuf:"varint,1,opt,def=0" json:"LittleBossCostHC,omitempty"`
	// * 大boss购买次数花费
	BigBossCostHC *uint32 `protobuf:"varint,2,opt,def=0" json:"BigBossCostHC,omitempty"`
	// * 活动结束时间
	FinishTime *string `protobuf:"bytes,3,opt,def=" json:"FinishTime,omitempty"`
	// * 活动重置开启时间
	RestartTime *string `protobuf:"bytes,4,opt,def=" json:"RestartTime,omitempty"`
	// * Loading阶段保护时间（秒）
	LoadingSafeTime *uint32 `protobuf:"varint,5,opt,def=0" json:"LoadingSafeTime,omitempty"`
	// * 功能解锁需要的公会等级
	OpenGuildLevel *uint32 `protobuf:"varint,6,opt,def=0" json:"OpenGuildLevel,omitempty"`
	// * 剩余挑战小boss次数补偿军魂系数
	OffsetMiniBossFreeGB *uint32 `protobuf:"varint,7,opt,def=0" json:"OffsetMiniBossFreeGB,omitempty"`
	// * 可付费小boss次数补偿军魂系数
	OffsetMiniBossCostGB *uint32 `protobuf:"varint,8,opt,def=0" json:"OffsetMiniBossCostGB,omitempty"`
	// * 小boss付费补偿单价
	OffsetMiniBossCostHC *uint32 `protobuf:"varint,9,opt,def=0" json:"OffsetMiniBossCostHC,omitempty"`
	// * 剩余挑战大boss次数补偿军魂系数
	OffsetBigBossFreeGB *uint32 `protobuf:"varint,10,opt,def=0" json:"OffsetBigBossFreeGB,omitempty"`
	// * 可付费大boss次数补偿军魂系数
	OffsetBigBossCostGB *uint32 `protobuf:"varint,11,opt,def=0" json:"OffsetBigBossCostGB,omitempty"`
	// * 大boss付费补偿单价
	OffsetBigBossCostHC *uint32 `protobuf:"varint,12,opt,def=0" json:"OffsetBigBossCostHC,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *GBCONFIG) Reset()         { *m = GBCONFIG{} }
func (m *GBCONFIG) String() string { return proto.CompactTextString(m) }
func (*GBCONFIG) ProtoMessage()    {}

const Default_GBCONFIG_LittleBossCostHC uint32 = 0
const Default_GBCONFIG_BigBossCostHC uint32 = 0
const Default_GBCONFIG_LoadingSafeTime uint32 = 0
const Default_GBCONFIG_OpenGuildLevel uint32 = 0
const Default_GBCONFIG_OffsetMiniBossFreeGB uint32 = 0
const Default_GBCONFIG_OffsetMiniBossCostGB uint32 = 0
const Default_GBCONFIG_OffsetMiniBossCostHC uint32 = 0
const Default_GBCONFIG_OffsetBigBossFreeGB uint32 = 0
const Default_GBCONFIG_OffsetBigBossCostGB uint32 = 0
const Default_GBCONFIG_OffsetBigBossCostHC uint32 = 0

func (m *GBCONFIG) GetLittleBossCostHC() uint32 {
	if m != nil && m.LittleBossCostHC != nil {
		return *m.LittleBossCostHC
	}
	return Default_GBCONFIG_LittleBossCostHC
}

func (m *GBCONFIG) GetBigBossCostHC() uint32 {
	if m != nil && m.BigBossCostHC != nil {
		return *m.BigBossCostHC
	}
	return Default_GBCONFIG_BigBossCostHC
}

func (m *GBCONFIG) GetFinishTime() string {
	if m != nil && m.FinishTime != nil {
		return *m.FinishTime
	}
	return ""
}

func (m *GBCONFIG) GetRestartTime() string {
	if m != nil && m.RestartTime != nil {
		return *m.RestartTime
	}
	return ""
}

func (m *GBCONFIG) GetLoadingSafeTime() uint32 {
	if m != nil && m.LoadingSafeTime != nil {
		return *m.LoadingSafeTime
	}
	return Default_GBCONFIG_LoadingSafeTime
}

func (m *GBCONFIG) GetOpenGuildLevel() uint32 {
	if m != nil && m.OpenGuildLevel != nil {
		return *m.OpenGuildLevel
	}
	return Default_GBCONFIG_OpenGuildLevel
}

func (m *GBCONFIG) GetOffsetMiniBossFreeGB() uint32 {
	if m != nil && m.OffsetMiniBossFreeGB != nil {
		return *m.OffsetMiniBossFreeGB
	}
	return Default_GBCONFIG_OffsetMiniBossFreeGB
}

func (m *GBCONFIG) GetOffsetMiniBossCostGB() uint32 {
	if m != nil && m.OffsetMiniBossCostGB != nil {
		return *m.OffsetMiniBossCostGB
	}
	return Default_GBCONFIG_OffsetMiniBossCostGB
}

func (m *GBCONFIG) GetOffsetMiniBossCostHC() uint32 {
	if m != nil && m.OffsetMiniBossCostHC != nil {
		return *m.OffsetMiniBossCostHC
	}
	return Default_GBCONFIG_OffsetMiniBossCostHC
}

func (m *GBCONFIG) GetOffsetBigBossFreeGB() uint32 {
	if m != nil && m.OffsetBigBossFreeGB != nil {
		return *m.OffsetBigBossFreeGB
	}
	return Default_GBCONFIG_OffsetBigBossFreeGB
}

func (m *GBCONFIG) GetOffsetBigBossCostGB() uint32 {
	if m != nil && m.OffsetBigBossCostGB != nil {
		return *m.OffsetBigBossCostGB
	}
	return Default_GBCONFIG_OffsetBigBossCostGB
}

func (m *GBCONFIG) GetOffsetBigBossCostHC() uint32 {
	if m != nil && m.OffsetBigBossCostHC != nil {
		return *m.OffsetBigBossCostHC
	}
	return Default_GBCONFIG_OffsetBigBossCostHC
}

type GBCONFIG_ARRAY struct {
	Items            []*GBCONFIG `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *GBCONFIG_ARRAY) Reset()         { *m = GBCONFIG_ARRAY{} }
func (m *GBCONFIG_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GBCONFIG_ARRAY) ProtoMessage()    {}

func (m *GBCONFIG_ARRAY) GetItems() []*GBCONFIG {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
