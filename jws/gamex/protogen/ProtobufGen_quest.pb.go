// Code generated by protoc-gen-go.
// source: ProtobufGen_quest.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Quest struct {
	// * 任务ID
	QuestID *uint32 `protobuf:"varint,1,req,def=0" json:"QuestID,omitempty"`
	// * 是否自动接取
	AutoAccept *uint32 `protobuf:"varint,33,opt,def=0" json:"AutoAccept,omitempty"`
	// * 任务类型
	QuestType *uint32 `protobuf:"varint,2,opt,def=0" json:"QuestType,omitempty"`
	// * 领取NPC
	QuestNpc *string `protobuf:"bytes,3,opt,def=" json:"QuestNpc,omitempty"`
	// * 完成NPC
	ReportNpc *string `protobuf:"bytes,4,opt,def=" json:"ReportNpc,omitempty"`
	// * 任务对话
	ReportTalk   *string                  `protobuf:"bytes,5,opt,def=" json:"ReportTalk,omitempty"`
	AccCon_Table []*Quest_AcceptCondition `protobuf:"bytes,6,rep" json:"AccCon_Table,omitempty"`
	// * 任务名称
	QuestName *string `protobuf:"bytes,7,opt,def=" json:"QuestName,omitempty"`
	// * 任务描述
	QuestDes *string `protobuf:"bytes,8,opt,def=" json:"QuestDes,omitempty"`
	// * 完成条件类型
	FCType *uint32 `protobuf:"varint,9,opt,def=0" json:"FCType,omitempty"`
	// * 完成条件数值参数1
	FCValueIP1 *uint32 `protobuf:"varint,10,opt,def=1" json:"FCValueIP1,omitempty"`
	// * 完成条件数值参数2
	FCValueIP2 *uint32 `protobuf:"varint,11,opt,def=0" json:"FCValueIP2,omitempty"`
	// * 完成条件字符串参数1
	FCValueSP1 *string `protobuf:"bytes,12,opt,def=" json:"FCValueSP1,omitempty"`
	// * 完成条件字符串参数2
	FCValueSP2 *string `protobuf:"bytes,13,opt,def=" json:"FCValueSP2,omitempty"`
	// * 后置任务（自动，无视条件）
	PostQuest *string `protobuf:"bytes,14,opt,def=" json:"PostQuest,omitempty"`
	// * 战队经验奖励
	CorpXP *uint32 `protobuf:"varint,15,opt,def=0" json:"CorpXP,omitempty"`
	// * 角色经验奖励
	CharaXP *uint32 `protobuf:"varint,16,opt,def=0" json:"CharaXP,omitempty"`
	// * 软通奖励
	SC *uint32 `protobuf:"varint,17,opt,def=0" json:"SC,omitempty"`
	// * 硬通奖励
	HC *uint32 `protobuf:"varint,18,opt,def=0" json:"HC,omitempty"`
	// * 物品ID
	Item1 *string `protobuf:"bytes,19,opt,def=" json:"Item1,omitempty"`
	// * 物品数量
	Count1 *uint32 `protobuf:"varint,20,opt,def=1" json:"Count1,omitempty"`
	// * 物品ID
	Item2 *string `protobuf:"bytes,21,opt,def=" json:"Item2,omitempty"`
	// * 物品数量
	Count2 *uint32 `protobuf:"varint,22,opt,def=1" json:"Count2,omitempty"`
	// * 物品ID
	Item3 *string `protobuf:"bytes,23,opt,def=" json:"Item3,omitempty"`
	// * 物品数量
	Count3 *uint32 `protobuf:"varint,24,opt,def=1" json:"Count3,omitempty"`
	// * 物品ID
	Item4 *string `protobuf:"bytes,25,opt,def=" json:"Item4,omitempty"`
	// * 物品数量
	Count4 *uint32 `protobuf:"varint,26,opt,def=1" json:"Count4,omitempty"`
	// * 跳转类型
	TeleType *uint32 `protobuf:"varint,27,opt,def=0" json:"TeleType,omitempty"`
	// * 跳转ID
	TeleID *string `protobuf:"bytes,28,opt,def=" json:"TeleID,omitempty"`
	// * 任务获得的活跃值
	ActiveValue *uint32 `protobuf:"varint,29,opt,def=0" json:"ActiveValue,omitempty"`
	// * 开服任务活跃值
	SevenDaysActiveValue *uint32 `protobuf:"varint,30,opt,def=0" json:"SevenDaysActiveValue,omitempty"`
	// * 回复任务时是否强制跑路
	RunForce *uint32 `protobuf:"varint,31,opt,def=0" json:"RunForce,omitempty"`
	// * 对应跳转界面解锁的枚举
	ConditionID      *uint32 `protobuf:"varint,32,opt,def=0" json:"ConditionID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Quest) Reset()         { *m = Quest{} }
func (m *Quest) String() string { return proto.CompactTextString(m) }
func (*Quest) ProtoMessage()    {}

const Default_Quest_QuestID uint32 = 0
const Default_Quest_AutoAccept uint32 = 0
const Default_Quest_QuestType uint32 = 0
const Default_Quest_FCType uint32 = 0
const Default_Quest_FCValueIP1 uint32 = 1
const Default_Quest_FCValueIP2 uint32 = 0
const Default_Quest_CorpXP uint32 = 0
const Default_Quest_CharaXP uint32 = 0
const Default_Quest_SC uint32 = 0
const Default_Quest_HC uint32 = 0
const Default_Quest_Count1 uint32 = 1
const Default_Quest_Count2 uint32 = 1
const Default_Quest_Count3 uint32 = 1
const Default_Quest_Count4 uint32 = 1
const Default_Quest_TeleType uint32 = 0
const Default_Quest_ActiveValue uint32 = 0
const Default_Quest_SevenDaysActiveValue uint32 = 0
const Default_Quest_RunForce uint32 = 0
const Default_Quest_ConditionID uint32 = 0

func (m *Quest) GetQuestID() uint32 {
	if m != nil && m.QuestID != nil {
		return *m.QuestID
	}
	return Default_Quest_QuestID
}

func (m *Quest) GetAutoAccept() uint32 {
	if m != nil && m.AutoAccept != nil {
		return *m.AutoAccept
	}
	return Default_Quest_AutoAccept
}

func (m *Quest) GetQuestType() uint32 {
	if m != nil && m.QuestType != nil {
		return *m.QuestType
	}
	return Default_Quest_QuestType
}

func (m *Quest) GetQuestNpc() string {
	if m != nil && m.QuestNpc != nil {
		return *m.QuestNpc
	}
	return ""
}

func (m *Quest) GetReportNpc() string {
	if m != nil && m.ReportNpc != nil {
		return *m.ReportNpc
	}
	return ""
}

func (m *Quest) GetReportTalk() string {
	if m != nil && m.ReportTalk != nil {
		return *m.ReportTalk
	}
	return ""
}

func (m *Quest) GetAccCon_Table() []*Quest_AcceptCondition {
	if m != nil {
		return m.AccCon_Table
	}
	return nil
}

func (m *Quest) GetQuestName() string {
	if m != nil && m.QuestName != nil {
		return *m.QuestName
	}
	return ""
}

func (m *Quest) GetQuestDes() string {
	if m != nil && m.QuestDes != nil {
		return *m.QuestDes
	}
	return ""
}

func (m *Quest) GetFCType() uint32 {
	if m != nil && m.FCType != nil {
		return *m.FCType
	}
	return Default_Quest_FCType
}

func (m *Quest) GetFCValueIP1() uint32 {
	if m != nil && m.FCValueIP1 != nil {
		return *m.FCValueIP1
	}
	return Default_Quest_FCValueIP1
}

func (m *Quest) GetFCValueIP2() uint32 {
	if m != nil && m.FCValueIP2 != nil {
		return *m.FCValueIP2
	}
	return Default_Quest_FCValueIP2
}

func (m *Quest) GetFCValueSP1() string {
	if m != nil && m.FCValueSP1 != nil {
		return *m.FCValueSP1
	}
	return ""
}

func (m *Quest) GetFCValueSP2() string {
	if m != nil && m.FCValueSP2 != nil {
		return *m.FCValueSP2
	}
	return ""
}

func (m *Quest) GetPostQuest() string {
	if m != nil && m.PostQuest != nil {
		return *m.PostQuest
	}
	return ""
}

func (m *Quest) GetCorpXP() uint32 {
	if m != nil && m.CorpXP != nil {
		return *m.CorpXP
	}
	return Default_Quest_CorpXP
}

func (m *Quest) GetCharaXP() uint32 {
	if m != nil && m.CharaXP != nil {
		return *m.CharaXP
	}
	return Default_Quest_CharaXP
}

func (m *Quest) GetSC() uint32 {
	if m != nil && m.SC != nil {
		return *m.SC
	}
	return Default_Quest_SC
}

func (m *Quest) GetHC() uint32 {
	if m != nil && m.HC != nil {
		return *m.HC
	}
	return Default_Quest_HC
}

func (m *Quest) GetItem1() string {
	if m != nil && m.Item1 != nil {
		return *m.Item1
	}
	return ""
}

func (m *Quest) GetCount1() uint32 {
	if m != nil && m.Count1 != nil {
		return *m.Count1
	}
	return Default_Quest_Count1
}

func (m *Quest) GetItem2() string {
	if m != nil && m.Item2 != nil {
		return *m.Item2
	}
	return ""
}

func (m *Quest) GetCount2() uint32 {
	if m != nil && m.Count2 != nil {
		return *m.Count2
	}
	return Default_Quest_Count2
}

func (m *Quest) GetItem3() string {
	if m != nil && m.Item3 != nil {
		return *m.Item3
	}
	return ""
}

func (m *Quest) GetCount3() uint32 {
	if m != nil && m.Count3 != nil {
		return *m.Count3
	}
	return Default_Quest_Count3
}

func (m *Quest) GetItem4() string {
	if m != nil && m.Item4 != nil {
		return *m.Item4
	}
	return ""
}

func (m *Quest) GetCount4() uint32 {
	if m != nil && m.Count4 != nil {
		return *m.Count4
	}
	return Default_Quest_Count4
}

func (m *Quest) GetTeleType() uint32 {
	if m != nil && m.TeleType != nil {
		return *m.TeleType
	}
	return Default_Quest_TeleType
}

func (m *Quest) GetTeleID() string {
	if m != nil && m.TeleID != nil {
		return *m.TeleID
	}
	return ""
}

func (m *Quest) GetActiveValue() uint32 {
	if m != nil && m.ActiveValue != nil {
		return *m.ActiveValue
	}
	return Default_Quest_ActiveValue
}

func (m *Quest) GetSevenDaysActiveValue() uint32 {
	if m != nil && m.SevenDaysActiveValue != nil {
		return *m.SevenDaysActiveValue
	}
	return Default_Quest_SevenDaysActiveValue
}

func (m *Quest) GetRunForce() uint32 {
	if m != nil && m.RunForce != nil {
		return *m.RunForce
	}
	return Default_Quest_RunForce
}

func (m *Quest) GetConditionID() uint32 {
	if m != nil && m.ConditionID != nil {
		return *m.ConditionID
	}
	return Default_Quest_ConditionID
}

type Quest_AcceptCondition struct {
	// * 接取条件类型
	ACType *uint32 `protobuf:"varint,1,opt,def=0" json:"ACType,omitempty"`
	// * 接取数值条件1
	ACValueIP1 *uint32 `protobuf:"varint,2,opt,def=0" json:"ACValueIP1,omitempty"`
	// * 接取数值条件2
	ACValueIP2 *uint32 `protobuf:"varint,3,opt,def=0" json:"ACValueIP2,omitempty"`
	// * 接取字符串条件
	ACValueSP        *string `protobuf:"bytes,4,opt,def=" json:"ACValueSP,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Quest_AcceptCondition) Reset()         { *m = Quest_AcceptCondition{} }
func (m *Quest_AcceptCondition) String() string { return proto.CompactTextString(m) }
func (*Quest_AcceptCondition) ProtoMessage()    {}

const Default_Quest_AcceptCondition_ACType uint32 = 0
const Default_Quest_AcceptCondition_ACValueIP1 uint32 = 0
const Default_Quest_AcceptCondition_ACValueIP2 uint32 = 0

func (m *Quest_AcceptCondition) GetACType() uint32 {
	if m != nil && m.ACType != nil {
		return *m.ACType
	}
	return Default_Quest_AcceptCondition_ACType
}

func (m *Quest_AcceptCondition) GetACValueIP1() uint32 {
	if m != nil && m.ACValueIP1 != nil {
		return *m.ACValueIP1
	}
	return Default_Quest_AcceptCondition_ACValueIP1
}

func (m *Quest_AcceptCondition) GetACValueIP2() uint32 {
	if m != nil && m.ACValueIP2 != nil {
		return *m.ACValueIP2
	}
	return Default_Quest_AcceptCondition_ACValueIP2
}

func (m *Quest_AcceptCondition) GetACValueSP() string {
	if m != nil && m.ACValueSP != nil {
		return *m.ACValueSP
	}
	return ""
}

type Quest_ARRAY struct {
	Items            []*Quest `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Quest_ARRAY) Reset()         { *m = Quest_ARRAY{} }
func (m *Quest_ARRAY) String() string { return proto.CompactTextString(m) }
func (*Quest_ARRAY) ProtoMessage()    {}

func (m *Quest_ARRAY) GetItems() []*Quest {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
