// Code generated by protoc-gen-go.
// source: ProtobufGen_level_trial.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type LEVEL_TRIAL struct {
	// * 关卡ID
	LevelID *int32 `protobuf:"varint,1,req,name=levelID,def=0" json:"levelID,omitempty"`
	// * 推荐战力
	LevelGS *uint32 `protobuf:"varint,2,opt,def=0" json:"LevelGS,omitempty"`
	// * 扫荡时间（秒）
	Time *int32 `protobuf:"varint,3,opt,name=time,def=0" json:"time,omitempty"`
	// * 对应的关卡
	LevelDemo *string `protobuf:"bytes,4,opt,def=" json:"LevelDemo,omitempty"`
	// * 关卡限时
	TimeLimit *int32 `protobuf:"varint,5,opt,name=timeLimit,def=0" json:"timeLimit,omitempty"`
	// * 层数排序
	TrialIndex *int32 `protobuf:"varint,6,opt,def=0" json:"TrialIndex,omitempty"`
	// * 首通奖励SC数量
	FirstSC *int32 `protobuf:"varint,7,opt,def=0" json:"FirstSC,omitempty"`
	// * 首通奖励DC数量
	FirstDC *int32 `protobuf:"varint,8,opt,def=0" json:"FirstDC,omitempty"`
	// * 首通奖励FI数量
	FirstFI *int32 `protobuf:"varint,9,opt,def=0" json:"FirstFI,omitempty"`
	// * 首通奖励SB数量
	FirstSB *int32 `protobuf:"varint,10,opt,def=0" json:"FirstSB,omitempty"`
	// * 是否有宝箱
	Bonus *int32 `protobuf:"varint,11,opt,def=0" json:"Bonus,omitempty"`
	// * 宝箱奖励SC数量
	BonusSC *int32 `protobuf:"varint,12,opt,def=0" json:"BonusSC,omitempty"`
	// * 宝箱奖励DC数量
	BonusDC *int32 `protobuf:"varint,13,opt,def=0" json:"BonusDC,omitempty"`
	// * 宝箱奖励展示物品ID
	BonusDropShow *string `protobuf:"bytes,14,opt,def=" json:"BonusDropShow,omitempty"`
	// * 宝箱奖励掉落组ID
	BonusDropGroup *string `protobuf:"bytes,15,opt,def=" json:"BonusDropGroup,omitempty"`
	// * 宝箱奖励掉落组ID
	BonusDropGroup1 *string `protobuf:"bytes,16,opt,def=" json:"BonusDropGroup1,omitempty"`
	// * 宝箱奖励掉落组ID
	BonusDropGroup2 *string `protobuf:"bytes,17,opt,def=" json:"BonusDropGroup2,omitempty"`
	// * 非首通奖励SC数量
	RewardSC *int32 `protobuf:"varint,18,opt,def=0" json:"RewardSC,omitempty"`
	// * 非首通奖励DC数量
	RewardDC *int32 `protobuf:"varint,19,opt,def=0" json:"RewardDC,omitempty"`
	// * 非首通奖励FI数量
	RewardFI *int32 `protobuf:"varint,20,opt,def=0" json:"RewardFI,omitempty"`
	// * 非首通奖励SB数量
	RewardSB *int32 `protobuf:"varint,26,opt,def=0" json:"RewardSB,omitempty"`
	// * 非首通奖励展示物品ID
	DropGroupShow *string `protobuf:"bytes,21,opt,def=" json:"DropGroupShow,omitempty"`
	// * 非首通奖励组掉落概率
	LootProbability *float32 `protobuf:"fixed32,22,opt,def=0" json:"LootProbability,omitempty"`
	// * 非首通奖励掉落组ID
	DropGroup *string `protobuf:"bytes,23,opt,def=" json:"DropGroup,omitempty"`
	// * 非首通奖励掉落组ID
	DropGroup1 *string `protobuf:"bytes,24,opt,def=" json:"DropGroup1,omitempty"`
	// * 非首通奖励掉落组ID
	DropGroup2       *string `protobuf:"bytes,25,opt,def=" json:"DropGroup2,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LEVEL_TRIAL) Reset()         { *m = LEVEL_TRIAL{} }
func (m *LEVEL_TRIAL) String() string { return proto.CompactTextString(m) }
func (*LEVEL_TRIAL) ProtoMessage()    {}

const Default_LEVEL_TRIAL_LevelID int32 = 0
const Default_LEVEL_TRIAL_LevelGS uint32 = 0
const Default_LEVEL_TRIAL_Time int32 = 0
const Default_LEVEL_TRIAL_TimeLimit int32 = 0
const Default_LEVEL_TRIAL_TrialIndex int32 = 0
const Default_LEVEL_TRIAL_FirstSC int32 = 0
const Default_LEVEL_TRIAL_FirstDC int32 = 0
const Default_LEVEL_TRIAL_FirstFI int32 = 0
const Default_LEVEL_TRIAL_FirstSB int32 = 0
const Default_LEVEL_TRIAL_Bonus int32 = 0
const Default_LEVEL_TRIAL_BonusSC int32 = 0
const Default_LEVEL_TRIAL_BonusDC int32 = 0
const Default_LEVEL_TRIAL_RewardSC int32 = 0
const Default_LEVEL_TRIAL_RewardDC int32 = 0
const Default_LEVEL_TRIAL_RewardFI int32 = 0
const Default_LEVEL_TRIAL_RewardSB int32 = 0
const Default_LEVEL_TRIAL_LootProbability float32 = 0

func (m *LEVEL_TRIAL) GetLevelID() int32 {
	if m != nil && m.LevelID != nil {
		return *m.LevelID
	}
	return Default_LEVEL_TRIAL_LevelID
}

func (m *LEVEL_TRIAL) GetLevelGS() uint32 {
	if m != nil && m.LevelGS != nil {
		return *m.LevelGS
	}
	return Default_LEVEL_TRIAL_LevelGS
}

func (m *LEVEL_TRIAL) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return Default_LEVEL_TRIAL_Time
}

func (m *LEVEL_TRIAL) GetLevelDemo() string {
	if m != nil && m.LevelDemo != nil {
		return *m.LevelDemo
	}
	return ""
}

func (m *LEVEL_TRIAL) GetTimeLimit() int32 {
	if m != nil && m.TimeLimit != nil {
		return *m.TimeLimit
	}
	return Default_LEVEL_TRIAL_TimeLimit
}

func (m *LEVEL_TRIAL) GetTrialIndex() int32 {
	if m != nil && m.TrialIndex != nil {
		return *m.TrialIndex
	}
	return Default_LEVEL_TRIAL_TrialIndex
}

func (m *LEVEL_TRIAL) GetFirstSC() int32 {
	if m != nil && m.FirstSC != nil {
		return *m.FirstSC
	}
	return Default_LEVEL_TRIAL_FirstSC
}

func (m *LEVEL_TRIAL) GetFirstDC() int32 {
	if m != nil && m.FirstDC != nil {
		return *m.FirstDC
	}
	return Default_LEVEL_TRIAL_FirstDC
}

func (m *LEVEL_TRIAL) GetFirstFI() int32 {
	if m != nil && m.FirstFI != nil {
		return *m.FirstFI
	}
	return Default_LEVEL_TRIAL_FirstFI
}

func (m *LEVEL_TRIAL) GetFirstSB() int32 {
	if m != nil && m.FirstSB != nil {
		return *m.FirstSB
	}
	return Default_LEVEL_TRIAL_FirstSB
}

func (m *LEVEL_TRIAL) GetBonus() int32 {
	if m != nil && m.Bonus != nil {
		return *m.Bonus
	}
	return Default_LEVEL_TRIAL_Bonus
}

func (m *LEVEL_TRIAL) GetBonusSC() int32 {
	if m != nil && m.BonusSC != nil {
		return *m.BonusSC
	}
	return Default_LEVEL_TRIAL_BonusSC
}

func (m *LEVEL_TRIAL) GetBonusDC() int32 {
	if m != nil && m.BonusDC != nil {
		return *m.BonusDC
	}
	return Default_LEVEL_TRIAL_BonusDC
}

func (m *LEVEL_TRIAL) GetBonusDropShow() string {
	if m != nil && m.BonusDropShow != nil {
		return *m.BonusDropShow
	}
	return ""
}

func (m *LEVEL_TRIAL) GetBonusDropGroup() string {
	if m != nil && m.BonusDropGroup != nil {
		return *m.BonusDropGroup
	}
	return ""
}

func (m *LEVEL_TRIAL) GetBonusDropGroup1() string {
	if m != nil && m.BonusDropGroup1 != nil {
		return *m.BonusDropGroup1
	}
	return ""
}

func (m *LEVEL_TRIAL) GetBonusDropGroup2() string {
	if m != nil && m.BonusDropGroup2 != nil {
		return *m.BonusDropGroup2
	}
	return ""
}

func (m *LEVEL_TRIAL) GetRewardSC() int32 {
	if m != nil && m.RewardSC != nil {
		return *m.RewardSC
	}
	return Default_LEVEL_TRIAL_RewardSC
}

func (m *LEVEL_TRIAL) GetRewardDC() int32 {
	if m != nil && m.RewardDC != nil {
		return *m.RewardDC
	}
	return Default_LEVEL_TRIAL_RewardDC
}

func (m *LEVEL_TRIAL) GetRewardFI() int32 {
	if m != nil && m.RewardFI != nil {
		return *m.RewardFI
	}
	return Default_LEVEL_TRIAL_RewardFI
}

func (m *LEVEL_TRIAL) GetRewardSB() int32 {
	if m != nil && m.RewardSB != nil {
		return *m.RewardSB
	}
	return Default_LEVEL_TRIAL_RewardSB
}

func (m *LEVEL_TRIAL) GetDropGroupShow() string {
	if m != nil && m.DropGroupShow != nil {
		return *m.DropGroupShow
	}
	return ""
}

func (m *LEVEL_TRIAL) GetLootProbability() float32 {
	if m != nil && m.LootProbability != nil {
		return *m.LootProbability
	}
	return Default_LEVEL_TRIAL_LootProbability
}

func (m *LEVEL_TRIAL) GetDropGroup() string {
	if m != nil && m.DropGroup != nil {
		return *m.DropGroup
	}
	return ""
}

func (m *LEVEL_TRIAL) GetDropGroup1() string {
	if m != nil && m.DropGroup1 != nil {
		return *m.DropGroup1
	}
	return ""
}

func (m *LEVEL_TRIAL) GetDropGroup2() string {
	if m != nil && m.DropGroup2 != nil {
		return *m.DropGroup2
	}
	return ""
}

type LEVEL_TRIAL_ARRAY struct {
	Items            []*LEVEL_TRIAL `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *LEVEL_TRIAL_ARRAY) Reset()         { *m = LEVEL_TRIAL_ARRAY{} }
func (m *LEVEL_TRIAL_ARRAY) String() string { return proto.CompactTextString(m) }
func (*LEVEL_TRIAL_ARRAY) ProtoMessage()    {}

func (m *LEVEL_TRIAL_ARRAY) GetItems() []*LEVEL_TRIAL {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}