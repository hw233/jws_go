// Code generated by protoc-gen-go.
// source: ProtobufGen_subquest.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SUBQUEST struct {
	// * 任务ID
	QuestID *uint32 `protobuf:"varint,1,req,def=0" json:"QuestID,omitempty"`
	// * 任务类型
	QuestType *uint32 `protobuf:"varint,2,opt,def=0" json:"QuestType,omitempty"`
	// * 任务名称
	QuestName *string `protobuf:"bytes,3,opt,def=" json:"QuestName,omitempty"`
	// * 任务描述
	QuestDes *string `protobuf:"bytes,4,opt,def=" json:"QuestDes,omitempty"`
	// * 完成条件类型
	FCType *uint32 `protobuf:"varint,5,opt,def=0" json:"FCType,omitempty"`
	// * 完成条件数值参数1
	FCValueIP1 *uint32 `protobuf:"varint,6,opt,def=1" json:"FCValueIP1,omitempty"`
	// * 完成条件数值参数2
	FCValueIP2 *uint32 `protobuf:"varint,7,opt,def=0" json:"FCValueIP2,omitempty"`
	// * 完成条件字符串参数1
	FCValueSP1 *string `protobuf:"bytes,8,opt,def=" json:"FCValueSP1,omitempty"`
	// * 跳转类型
	TeleType *uint32 `protobuf:"varint,9,opt,def=0" json:"TeleType,omitempty"`
	// * 跳转ID
	TeleID *string `protobuf:"bytes,10,opt,def=" json:"TeleID,omitempty"`
	// * 任务获得的活跃值
	ActiveValue *uint32 `protobuf:"varint,11,opt,def=0" json:"ActiveValue,omitempty"`
	// * 对应跳转界面解锁的枚举
	ConditionID      *uint32                   `protobuf:"varint,12,opt,def=999" json:"ConditionID,omitempty"`
	Item_Table       []*SUBQUEST_ItemCondition `protobuf:"bytes,13,rep" json:"Item_Table,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *SUBQUEST) Reset()         { *m = SUBQUEST{} }
func (m *SUBQUEST) String() string { return proto.CompactTextString(m) }
func (*SUBQUEST) ProtoMessage()    {}

const Default_SUBQUEST_QuestID uint32 = 0
const Default_SUBQUEST_QuestType uint32 = 0
const Default_SUBQUEST_FCType uint32 = 0
const Default_SUBQUEST_FCValueIP1 uint32 = 1
const Default_SUBQUEST_FCValueIP2 uint32 = 0
const Default_SUBQUEST_TeleType uint32 = 0
const Default_SUBQUEST_ActiveValue uint32 = 0
const Default_SUBQUEST_ConditionID uint32 = 999

func (m *SUBQUEST) GetQuestID() uint32 {
	if m != nil && m.QuestID != nil {
		return *m.QuestID
	}
	return Default_SUBQUEST_QuestID
}

func (m *SUBQUEST) GetQuestType() uint32 {
	if m != nil && m.QuestType != nil {
		return *m.QuestType
	}
	return Default_SUBQUEST_QuestType
}

func (m *SUBQUEST) GetQuestName() string {
	if m != nil && m.QuestName != nil {
		return *m.QuestName
	}
	return ""
}

func (m *SUBQUEST) GetQuestDes() string {
	if m != nil && m.QuestDes != nil {
		return *m.QuestDes
	}
	return ""
}

func (m *SUBQUEST) GetFCType() uint32 {
	if m != nil && m.FCType != nil {
		return *m.FCType
	}
	return Default_SUBQUEST_FCType
}

func (m *SUBQUEST) GetFCValueIP1() uint32 {
	if m != nil && m.FCValueIP1 != nil {
		return *m.FCValueIP1
	}
	return Default_SUBQUEST_FCValueIP1
}

func (m *SUBQUEST) GetFCValueIP2() uint32 {
	if m != nil && m.FCValueIP2 != nil {
		return *m.FCValueIP2
	}
	return Default_SUBQUEST_FCValueIP2
}

func (m *SUBQUEST) GetFCValueSP1() string {
	if m != nil && m.FCValueSP1 != nil {
		return *m.FCValueSP1
	}
	return ""
}

func (m *SUBQUEST) GetTeleType() uint32 {
	if m != nil && m.TeleType != nil {
		return *m.TeleType
	}
	return Default_SUBQUEST_TeleType
}

func (m *SUBQUEST) GetTeleID() string {
	if m != nil && m.TeleID != nil {
		return *m.TeleID
	}
	return ""
}

func (m *SUBQUEST) GetActiveValue() uint32 {
	if m != nil && m.ActiveValue != nil {
		return *m.ActiveValue
	}
	return Default_SUBQUEST_ActiveValue
}

func (m *SUBQUEST) GetConditionID() uint32 {
	if m != nil && m.ConditionID != nil {
		return *m.ConditionID
	}
	return Default_SUBQUEST_ConditionID
}

func (m *SUBQUEST) GetItem_Table() []*SUBQUEST_ItemCondition {
	if m != nil {
		return m.Item_Table
	}
	return nil
}

type SUBQUEST_ItemCondition struct {
	// * 奖励物品ID
	ItemID *string `protobuf:"bytes,1,opt,def=" json:"ItemID,omitempty"`
	// * 奖励物品数量
	ItemCount        *uint32 `protobuf:"varint,2,opt,def=0" json:"ItemCount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SUBQUEST_ItemCondition) Reset()         { *m = SUBQUEST_ItemCondition{} }
func (m *SUBQUEST_ItemCondition) String() string { return proto.CompactTextString(m) }
func (*SUBQUEST_ItemCondition) ProtoMessage()    {}

const Default_SUBQUEST_ItemCondition_ItemCount uint32 = 0

func (m *SUBQUEST_ItemCondition) GetItemID() string {
	if m != nil && m.ItemID != nil {
		return *m.ItemID
	}
	return ""
}

func (m *SUBQUEST_ItemCondition) GetItemCount() uint32 {
	if m != nil && m.ItemCount != nil {
		return *m.ItemCount
	}
	return Default_SUBQUEST_ItemCondition_ItemCount
}

type SUBQUEST_ARRAY struct {
	Items            []*SUBQUEST `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SUBQUEST_ARRAY) Reset()         { *m = SUBQUEST_ARRAY{} }
func (m *SUBQUEST_ARRAY) String() string { return proto.CompactTextString(m) }
func (*SUBQUEST_ARRAY) ProtoMessage()    {}

func (m *SUBQUEST_ARRAY) GetItems() []*SUBQUEST {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
