// Code generated by protoc-gen-go.
// source: ProtobufGen_fenghuomain.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type FENGHUOMAIN struct {
	// * 战役ID
	BattleHard *uint32 `protobuf:"varint,1,req,def=0" json:"BattleHard,omitempty"`
	// * 从属战役
	BattleName *string `protobuf:"bytes,2,opt,def=" json:"BattleName,omitempty"`
	// * 战役大类（1官渡，2赤壁，3南蛮）
	BattleEpic *uint32 `protobuf:"varint,24,opt,def=0" json:"BattleEpic,omitempty"`
	// * 从属章节
	BattleCharper *uint32 `protobuf:"varint,3,opt,def=0" json:"BattleCharper,omitempty"`
	// * 章节名
	CharperName *string `protobuf:"bytes,15,opt,def=" json:"CharperName,omitempty"`
	// * 最低开启等级
	NeedLevel *uint32 `protobuf:"varint,4,opt,def=0" json:"NeedLevel,omitempty"`
	// * 每小关金币消耗
	SubLevelCostSC *uint32 `protobuf:"varint,17,opt,def=0" json:"SubLevelCostSC,omitempty"`
	// * 免费加倍的每小关补偿奖励
	Compensation *string `protobuf:"bytes,23,opt,def=" json:"Compensation,omitempty"`
	// * 组队额外1倍小关钻石消耗（1倍不消耗钻）
	TeamSubLevelCostHC *uint32 `protobuf:"varint,19,opt,def=0" json:"TeamSubLevelCostHC,omitempty"`
	// * 一轮最终奖励
	FinalLoot *string `protobuf:"bytes,12,opt,def=" json:"FinalLoot,omitempty"`
	// * 单人扫荡每轮消耗钻石
	SoloSweepEachRound *uint32 `protobuf:"varint,13,opt,def=0" json:"SoloSweepEachRound,omitempty"`
	// * 单人高级扫荡每轮消耗钻石
	SoloVIPSweepCostHCEachRound *uint32 `protobuf:"varint,14,opt,def=0" json:"SoloVIPSweepCostHCEachRound,omitempty"`
	// * 关卡展示图
	Texture *string `protobuf:"bytes,5,opt,def=" json:"Texture,omitempty"`
	// * 方图
	SquareTexture *string `protobuf:"bytes,6,opt,def=" json:"SquareTexture,omitempty"`
	// * 房间名称IDS
	TeamRoomIDS *string `protobuf:"bytes,7,opt,def=" json:"TeamRoomIDS,omitempty"`
	// * 战役中小关卡数量
	SubLevelNum          *uint32                      `protobuf:"varint,20,opt,def=0" json:"SubLevelNum,omitempty"`
	DoubleBossData_Table *FENGHUOMAIN_DoubleBossData  `protobuf:"bytes,21,opt" json:"DoubleBossData_Table,omitempty"`
	BigBossData_Table    *FENGHUOMAIN_BigBossData     `protobuf:"bytes,22,opt" json:"BigBossData_Table,omitempty"`
	LevelINfoData_Table  []*FENGHUOMAIN_LevelInfoData `protobuf:"bytes,9,rep" json:"LevelINfoData_Table,omitempty"`
	XXX_unrecognized     []byte                       `json:"-"`
}

func (m *FENGHUOMAIN) Reset()         { *m = FENGHUOMAIN{} }
func (m *FENGHUOMAIN) String() string { return proto.CompactTextString(m) }
func (*FENGHUOMAIN) ProtoMessage()    {}

const Default_FENGHUOMAIN_BattleHard uint32 = 0
const Default_FENGHUOMAIN_BattleEpic uint32 = 0
const Default_FENGHUOMAIN_BattleCharper uint32 = 0
const Default_FENGHUOMAIN_NeedLevel uint32 = 0
const Default_FENGHUOMAIN_SubLevelCostSC uint32 = 0
const Default_FENGHUOMAIN_TeamSubLevelCostHC uint32 = 0
const Default_FENGHUOMAIN_SoloSweepEachRound uint32 = 0
const Default_FENGHUOMAIN_SoloVIPSweepCostHCEachRound uint32 = 0
const Default_FENGHUOMAIN_SubLevelNum uint32 = 0

func (m *FENGHUOMAIN) GetBattleHard() uint32 {
	if m != nil && m.BattleHard != nil {
		return *m.BattleHard
	}
	return Default_FENGHUOMAIN_BattleHard
}

func (m *FENGHUOMAIN) GetBattleName() string {
	if m != nil && m.BattleName != nil {
		return *m.BattleName
	}
	return ""
}

func (m *FENGHUOMAIN) GetBattleEpic() uint32 {
	if m != nil && m.BattleEpic != nil {
		return *m.BattleEpic
	}
	return Default_FENGHUOMAIN_BattleEpic
}

func (m *FENGHUOMAIN) GetBattleCharper() uint32 {
	if m != nil && m.BattleCharper != nil {
		return *m.BattleCharper
	}
	return Default_FENGHUOMAIN_BattleCharper
}

func (m *FENGHUOMAIN) GetCharperName() string {
	if m != nil && m.CharperName != nil {
		return *m.CharperName
	}
	return ""
}

func (m *FENGHUOMAIN) GetNeedLevel() uint32 {
	if m != nil && m.NeedLevel != nil {
		return *m.NeedLevel
	}
	return Default_FENGHUOMAIN_NeedLevel
}

func (m *FENGHUOMAIN) GetSubLevelCostSC() uint32 {
	if m != nil && m.SubLevelCostSC != nil {
		return *m.SubLevelCostSC
	}
	return Default_FENGHUOMAIN_SubLevelCostSC
}

func (m *FENGHUOMAIN) GetCompensation() string {
	if m != nil && m.Compensation != nil {
		return *m.Compensation
	}
	return ""
}

func (m *FENGHUOMAIN) GetTeamSubLevelCostHC() uint32 {
	if m != nil && m.TeamSubLevelCostHC != nil {
		return *m.TeamSubLevelCostHC
	}
	return Default_FENGHUOMAIN_TeamSubLevelCostHC
}

func (m *FENGHUOMAIN) GetFinalLoot() string {
	if m != nil && m.FinalLoot != nil {
		return *m.FinalLoot
	}
	return ""
}

func (m *FENGHUOMAIN) GetSoloSweepEachRound() uint32 {
	if m != nil && m.SoloSweepEachRound != nil {
		return *m.SoloSweepEachRound
	}
	return Default_FENGHUOMAIN_SoloSweepEachRound
}

func (m *FENGHUOMAIN) GetSoloVIPSweepCostHCEachRound() uint32 {
	if m != nil && m.SoloVIPSweepCostHCEachRound != nil {
		return *m.SoloVIPSweepCostHCEachRound
	}
	return Default_FENGHUOMAIN_SoloVIPSweepCostHCEachRound
}

func (m *FENGHUOMAIN) GetTexture() string {
	if m != nil && m.Texture != nil {
		return *m.Texture
	}
	return ""
}

func (m *FENGHUOMAIN) GetSquareTexture() string {
	if m != nil && m.SquareTexture != nil {
		return *m.SquareTexture
	}
	return ""
}

func (m *FENGHUOMAIN) GetTeamRoomIDS() string {
	if m != nil && m.TeamRoomIDS != nil {
		return *m.TeamRoomIDS
	}
	return ""
}

func (m *FENGHUOMAIN) GetSubLevelNum() uint32 {
	if m != nil && m.SubLevelNum != nil {
		return *m.SubLevelNum
	}
	return Default_FENGHUOMAIN_SubLevelNum
}

func (m *FENGHUOMAIN) GetDoubleBossData_Table() *FENGHUOMAIN_DoubleBossData {
	if m != nil {
		return m.DoubleBossData_Table
	}
	return nil
}

func (m *FENGHUOMAIN) GetBigBossData_Table() *FENGHUOMAIN_BigBossData {
	if m != nil {
		return m.BigBossData_Table
	}
	return nil
}

func (m *FENGHUOMAIN) GetLevelINfoData_Table() []*FENGHUOMAIN_LevelInfoData {
	if m != nil {
		return m.LevelINfoData_Table
	}
	return nil
}

type FENGHUOMAIN_DoubleBossData struct {
	// * 双王关
	LevelInfoID *string `protobuf:"bytes,1,req,def=" json:"LevelInfoID,omitempty"`
	// * 最高出现次数
	MaxNum *uint32 `protobuf:"varint,2,opt,def=0" json:"MaxNum,omitempty"`
	// * 普通掉落（普通扫荡也用此列）
	NormalDrop *string `protobuf:"bytes,3,opt,def=" json:"NormalDrop,omitempty"`
	// * 高级扫荡掉落
	SeniorSweepDrop *string `protobuf:"bytes,4,opt,def=" json:"SeniorSweepDrop,omitempty"`
	// * 关卡类型（1=大boss，2=双子boss，3=小怪
	LevelType        *uint32 `protobuf:"varint,5,opt,def=0" json:"LevelType,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FENGHUOMAIN_DoubleBossData) Reset()         { *m = FENGHUOMAIN_DoubleBossData{} }
func (m *FENGHUOMAIN_DoubleBossData) String() string { return proto.CompactTextString(m) }
func (*FENGHUOMAIN_DoubleBossData) ProtoMessage()    {}

const Default_FENGHUOMAIN_DoubleBossData_MaxNum uint32 = 0
const Default_FENGHUOMAIN_DoubleBossData_LevelType uint32 = 0

func (m *FENGHUOMAIN_DoubleBossData) GetLevelInfoID() string {
	if m != nil && m.LevelInfoID != nil {
		return *m.LevelInfoID
	}
	return ""
}

func (m *FENGHUOMAIN_DoubleBossData) GetMaxNum() uint32 {
	if m != nil && m.MaxNum != nil {
		return *m.MaxNum
	}
	return Default_FENGHUOMAIN_DoubleBossData_MaxNum
}

func (m *FENGHUOMAIN_DoubleBossData) GetNormalDrop() string {
	if m != nil && m.NormalDrop != nil {
		return *m.NormalDrop
	}
	return ""
}

func (m *FENGHUOMAIN_DoubleBossData) GetSeniorSweepDrop() string {
	if m != nil && m.SeniorSweepDrop != nil {
		return *m.SeniorSweepDrop
	}
	return ""
}

func (m *FENGHUOMAIN_DoubleBossData) GetLevelType() uint32 {
	if m != nil && m.LevelType != nil {
		return *m.LevelType
	}
	return Default_FENGHUOMAIN_DoubleBossData_LevelType
}

type FENGHUOMAIN_BigBossData struct {
	// * 大王关
	LevelInfoID *string `protobuf:"bytes,1,req,def=" json:"LevelInfoID,omitempty"`
	// * 最高出现次数
	MaxNum *uint32 `protobuf:"varint,2,opt,def=0" json:"MaxNum,omitempty"`
	// * 普通掉落（普通扫荡也用此列）
	NormalDrop *string `protobuf:"bytes,3,opt,def=" json:"NormalDrop,omitempty"`
	// * 高级扫荡掉落
	SeniorSweepDrop *string `protobuf:"bytes,4,opt,def=" json:"SeniorSweepDrop,omitempty"`
	// * 关卡类型（1=大boss，2=双子boss，3=小怪
	LevelType        *uint32 `protobuf:"varint,5,opt,def=0" json:"LevelType,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FENGHUOMAIN_BigBossData) Reset()         { *m = FENGHUOMAIN_BigBossData{} }
func (m *FENGHUOMAIN_BigBossData) String() string { return proto.CompactTextString(m) }
func (*FENGHUOMAIN_BigBossData) ProtoMessage()    {}

const Default_FENGHUOMAIN_BigBossData_MaxNum uint32 = 0
const Default_FENGHUOMAIN_BigBossData_LevelType uint32 = 0

func (m *FENGHUOMAIN_BigBossData) GetLevelInfoID() string {
	if m != nil && m.LevelInfoID != nil {
		return *m.LevelInfoID
	}
	return ""
}

func (m *FENGHUOMAIN_BigBossData) GetMaxNum() uint32 {
	if m != nil && m.MaxNum != nil {
		return *m.MaxNum
	}
	return Default_FENGHUOMAIN_BigBossData_MaxNum
}

func (m *FENGHUOMAIN_BigBossData) GetNormalDrop() string {
	if m != nil && m.NormalDrop != nil {
		return *m.NormalDrop
	}
	return ""
}

func (m *FENGHUOMAIN_BigBossData) GetSeniorSweepDrop() string {
	if m != nil && m.SeniorSweepDrop != nil {
		return *m.SeniorSweepDrop
	}
	return ""
}

func (m *FENGHUOMAIN_BigBossData) GetLevelType() uint32 {
	if m != nil && m.LevelType != nil {
		return *m.LevelType
	}
	return Default_FENGHUOMAIN_BigBossData_LevelType
}

type FENGHUOMAIN_LevelInfoData struct {
	// * 关卡A
	LevelInfoID *string `protobuf:"bytes,1,req,def=" json:"LevelInfoID,omitempty"`
	// * 普通掉落（普通扫荡也用此列）
	NormalDrop *string `protobuf:"bytes,5,opt,def=" json:"NormalDrop,omitempty"`
	// * 高级扫荡掉落
	SeniorSweepDrop *string `protobuf:"bytes,7,opt,def=" json:"SeniorSweepDrop,omitempty"`
	// * 关卡类型（1=大boss，2=双子boss，3=小怪）
	LevelType        *uint32 `protobuf:"varint,4,opt,def=0" json:"LevelType,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FENGHUOMAIN_LevelInfoData) Reset()         { *m = FENGHUOMAIN_LevelInfoData{} }
func (m *FENGHUOMAIN_LevelInfoData) String() string { return proto.CompactTextString(m) }
func (*FENGHUOMAIN_LevelInfoData) ProtoMessage()    {}

const Default_FENGHUOMAIN_LevelInfoData_LevelType uint32 = 0

func (m *FENGHUOMAIN_LevelInfoData) GetLevelInfoID() string {
	if m != nil && m.LevelInfoID != nil {
		return *m.LevelInfoID
	}
	return ""
}

func (m *FENGHUOMAIN_LevelInfoData) GetNormalDrop() string {
	if m != nil && m.NormalDrop != nil {
		return *m.NormalDrop
	}
	return ""
}

func (m *FENGHUOMAIN_LevelInfoData) GetSeniorSweepDrop() string {
	if m != nil && m.SeniorSweepDrop != nil {
		return *m.SeniorSweepDrop
	}
	return ""
}

func (m *FENGHUOMAIN_LevelInfoData) GetLevelType() uint32 {
	if m != nil && m.LevelType != nil {
		return *m.LevelType
	}
	return Default_FENGHUOMAIN_LevelInfoData_LevelType
}

type FENGHUOMAIN_ARRAY struct {
	Items            []*FENGHUOMAIN `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *FENGHUOMAIN_ARRAY) Reset()         { *m = FENGHUOMAIN_ARRAY{} }
func (m *FENGHUOMAIN_ARRAY) String() string { return proto.CompactTextString(m) }
func (*FENGHUOMAIN_ARRAY) ProtoMessage()    {}

func (m *FENGHUOMAIN_ARRAY) GetItems() []*FENGHUOMAIN {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
