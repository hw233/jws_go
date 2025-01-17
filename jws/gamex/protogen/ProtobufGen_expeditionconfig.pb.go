// Code generated by protoc-gen-go.
// source: ProtobufGen_expeditionconfig.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type EXPEDITIONCONFIG struct {
	// * 进入主将等级
	HeroLevel *string `protobuf:"bytes,1,req,def=" json:"HeroLevel,omitempty"`
	// * 扫荡战力差
	GSLevelGap *float32 `protobuf:"fixed32,2,req,def=0" json:"GSLevelGap,omitempty"`
	// * 伤害加成
	AddDamage *uint32 `protobuf:"varint,3,req,def=0" json:"AddDamage,omitempty"`
	// * 暴击概率提升至%
	CritRateUp *float32 `protobuf:"fixed32,4,req,def=0" json:"CritRateUp,omitempty"`
	// * 暴击倍率提升至%
	CritDamageUp *float32 `protobuf:"fixed32,5,req,def=0" json:"CritDamageUp,omitempty"`
	// * 玩家方战力调整系数（1为不变）
	GSBuff *float32 `protobuf:"fixed32,6,req,def=0" json:"GSBuff,omitempty"`
	// * 消耗扫荡道具个数
	SweepCost *uint32 `protobuf:"varint,8,req,def=0" json:"SweepCost,omitempty"`
	// * 购买扫荡券用到的货币
	BuySweepItemUse *string `protobuf:"bytes,9,req,def=" json:"BuySweepItemUse,omitempty"`
	// * 购买扫荡券的单价
	BuySweepItemPrice *uint32                   `protobuf:"varint,10,req,def=0" json:"BuySweepItemPrice,omitempty"`
	Unlock_Table      []*EXPEDITIONCONFIG_Level `protobuf:"bytes,7,rep" json:"Unlock_Table,omitempty"`
	XXX_unrecognized  []byte                    `json:"-"`
}

func (m *EXPEDITIONCONFIG) Reset()         { *m = EXPEDITIONCONFIG{} }
func (m *EXPEDITIONCONFIG) String() string { return proto.CompactTextString(m) }
func (*EXPEDITIONCONFIG) ProtoMessage()    {}

const Default_EXPEDITIONCONFIG_GSLevelGap float32 = 0
const Default_EXPEDITIONCONFIG_AddDamage uint32 = 0
const Default_EXPEDITIONCONFIG_CritRateUp float32 = 0
const Default_EXPEDITIONCONFIG_CritDamageUp float32 = 0
const Default_EXPEDITIONCONFIG_GSBuff float32 = 0
const Default_EXPEDITIONCONFIG_SweepCost uint32 = 0
const Default_EXPEDITIONCONFIG_BuySweepItemPrice uint32 = 0

func (m *EXPEDITIONCONFIG) GetHeroLevel() string {
	if m != nil && m.HeroLevel != nil {
		return *m.HeroLevel
	}
	return ""
}

func (m *EXPEDITIONCONFIG) GetGSLevelGap() float32 {
	if m != nil && m.GSLevelGap != nil {
		return *m.GSLevelGap
	}
	return Default_EXPEDITIONCONFIG_GSLevelGap
}

func (m *EXPEDITIONCONFIG) GetAddDamage() uint32 {
	if m != nil && m.AddDamage != nil {
		return *m.AddDamage
	}
	return Default_EXPEDITIONCONFIG_AddDamage
}

func (m *EXPEDITIONCONFIG) GetCritRateUp() float32 {
	if m != nil && m.CritRateUp != nil {
		return *m.CritRateUp
	}
	return Default_EXPEDITIONCONFIG_CritRateUp
}

func (m *EXPEDITIONCONFIG) GetCritDamageUp() float32 {
	if m != nil && m.CritDamageUp != nil {
		return *m.CritDamageUp
	}
	return Default_EXPEDITIONCONFIG_CritDamageUp
}

func (m *EXPEDITIONCONFIG) GetGSBuff() float32 {
	if m != nil && m.GSBuff != nil {
		return *m.GSBuff
	}
	return Default_EXPEDITIONCONFIG_GSBuff
}

func (m *EXPEDITIONCONFIG) GetSweepCost() uint32 {
	if m != nil && m.SweepCost != nil {
		return *m.SweepCost
	}
	return Default_EXPEDITIONCONFIG_SweepCost
}

func (m *EXPEDITIONCONFIG) GetBuySweepItemUse() string {
	if m != nil && m.BuySweepItemUse != nil {
		return *m.BuySweepItemUse
	}
	return ""
}

func (m *EXPEDITIONCONFIG) GetBuySweepItemPrice() uint32 {
	if m != nil && m.BuySweepItemPrice != nil {
		return *m.BuySweepItemPrice
	}
	return Default_EXPEDITIONCONFIG_BuySweepItemPrice
}

func (m *EXPEDITIONCONFIG) GetUnlock_Table() []*EXPEDITIONCONFIG_Level {
	if m != nil {
		return m.Unlock_Table
	}
	return nil
}

type EXPEDITIONCONFIG_Level struct {
	// * 解锁等级
	UnlockLevel *uint32 `protobuf:"varint,1,opt,def=0" json:"UnlockLevel,omitempty"`
	// * 第X关前都可扫荡（包含）
	LevelID          *uint32 `protobuf:"varint,2,opt,def=0" json:"LevelID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EXPEDITIONCONFIG_Level) Reset()         { *m = EXPEDITIONCONFIG_Level{} }
func (m *EXPEDITIONCONFIG_Level) String() string { return proto.CompactTextString(m) }
func (*EXPEDITIONCONFIG_Level) ProtoMessage()    {}

const Default_EXPEDITIONCONFIG_Level_UnlockLevel uint32 = 0
const Default_EXPEDITIONCONFIG_Level_LevelID uint32 = 0

func (m *EXPEDITIONCONFIG_Level) GetUnlockLevel() uint32 {
	if m != nil && m.UnlockLevel != nil {
		return *m.UnlockLevel
	}
	return Default_EXPEDITIONCONFIG_Level_UnlockLevel
}

func (m *EXPEDITIONCONFIG_Level) GetLevelID() uint32 {
	if m != nil && m.LevelID != nil {
		return *m.LevelID
	}
	return Default_EXPEDITIONCONFIG_Level_LevelID
}

type EXPEDITIONCONFIG_ARRAY struct {
	Items            []*EXPEDITIONCONFIG `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *EXPEDITIONCONFIG_ARRAY) Reset()         { *m = EXPEDITIONCONFIG_ARRAY{} }
func (m *EXPEDITIONCONFIG_ARRAY) String() string { return proto.CompactTextString(m) }
func (*EXPEDITIONCONFIG_ARRAY) ProtoMessage()    {}

func (m *EXPEDITIONCONFIG_ARRAY) GetItems() []*EXPEDITIONCONFIG {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
