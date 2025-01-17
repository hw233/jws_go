// Code generated by protoc-gen-go.
// source: ProtobufGen_dubbleraward.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type DUBBLERAWARD struct {
	// * 活动ID
	ActivityID *uint32 `protobuf:"varint,1,req,def=0" json:"ActivityID,omitempty"`
	// * 激活条件
	GetCondition *uint32 `protobuf:"varint,2,opt,def=0" json:"GetCondition,omitempty"`
	// * 条件参数1（IOS)
	ConditionValue1ForIOS *uint32 `protobuf:"varint,3,opt,name=ConditionValue1forIOS,def=0" json:"ConditionValue1forIOS,omitempty"`
	// * 条件参数2（IOS)
	ConditionValue2ForIOS *uint32 `protobuf:"varint,4,opt,name=ConditionValue2forIOS,def=0" json:"ConditionValue2forIOS,omitempty"`
	// * 条件参数1（Android)
	ConditionValue1ForAndroid *uint32 `protobuf:"varint,5,opt,name=ConditionValue1forAndroid,def=0" json:"ConditionValue1forAndroid,omitempty"`
	// * 条件参数2(Android)
	ConditionValue2ForAndroid *uint32 `protobuf:"varint,6,opt,name=ConditionValue2forAndroid,def=0" json:"ConditionValue2forAndroid,omitempty"`
	// * 激活奖励类型
	GetTimeType *uint32 `protobuf:"varint,7,opt,def=0" json:"GetTimeType,omitempty"`
	// *
	GetDay *uint32 `protobuf:"varint,8,opt,def=0" json:"GetDay,omitempty"`
	// * 领取奖励类型
	GiftAcceptType   *uint32                     `protobuf:"varint,9,opt,def=0" json:"GiftAcceptType,omitempty"`
	Loots            []*DUBBLERAWARD_RawardLoots `protobuf:"bytes,10,rep" json:"Loots,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *DUBBLERAWARD) Reset()         { *m = DUBBLERAWARD{} }
func (m *DUBBLERAWARD) String() string { return proto.CompactTextString(m) }
func (*DUBBLERAWARD) ProtoMessage()    {}

const Default_DUBBLERAWARD_ActivityID uint32 = 0
const Default_DUBBLERAWARD_GetCondition uint32 = 0
const Default_DUBBLERAWARD_ConditionValue1ForIOS uint32 = 0
const Default_DUBBLERAWARD_ConditionValue2ForIOS uint32 = 0
const Default_DUBBLERAWARD_ConditionValue1ForAndroid uint32 = 0
const Default_DUBBLERAWARD_ConditionValue2ForAndroid uint32 = 0
const Default_DUBBLERAWARD_GetTimeType uint32 = 0
const Default_DUBBLERAWARD_GetDay uint32 = 0
const Default_DUBBLERAWARD_GiftAcceptType uint32 = 0

func (m *DUBBLERAWARD) GetActivityID() uint32 {
	if m != nil && m.ActivityID != nil {
		return *m.ActivityID
	}
	return Default_DUBBLERAWARD_ActivityID
}

func (m *DUBBLERAWARD) GetGetCondition() uint32 {
	if m != nil && m.GetCondition != nil {
		return *m.GetCondition
	}
	return Default_DUBBLERAWARD_GetCondition
}

func (m *DUBBLERAWARD) GetConditionValue1ForIOS() uint32 {
	if m != nil && m.ConditionValue1ForIOS != nil {
		return *m.ConditionValue1ForIOS
	}
	return Default_DUBBLERAWARD_ConditionValue1ForIOS
}

func (m *DUBBLERAWARD) GetConditionValue2ForIOS() uint32 {
	if m != nil && m.ConditionValue2ForIOS != nil {
		return *m.ConditionValue2ForIOS
	}
	return Default_DUBBLERAWARD_ConditionValue2ForIOS
}

func (m *DUBBLERAWARD) GetConditionValue1ForAndroid() uint32 {
	if m != nil && m.ConditionValue1ForAndroid != nil {
		return *m.ConditionValue1ForAndroid
	}
	return Default_DUBBLERAWARD_ConditionValue1ForAndroid
}

func (m *DUBBLERAWARD) GetConditionValue2ForAndroid() uint32 {
	if m != nil && m.ConditionValue2ForAndroid != nil {
		return *m.ConditionValue2ForAndroid
	}
	return Default_DUBBLERAWARD_ConditionValue2ForAndroid
}

func (m *DUBBLERAWARD) GetGetTimeType() uint32 {
	if m != nil && m.GetTimeType != nil {
		return *m.GetTimeType
	}
	return Default_DUBBLERAWARD_GetTimeType
}

func (m *DUBBLERAWARD) GetGetDay() uint32 {
	if m != nil && m.GetDay != nil {
		return *m.GetDay
	}
	return Default_DUBBLERAWARD_GetDay
}

func (m *DUBBLERAWARD) GetGiftAcceptType() uint32 {
	if m != nil && m.GiftAcceptType != nil {
		return *m.GiftAcceptType
	}
	return Default_DUBBLERAWARD_GiftAcceptType
}

func (m *DUBBLERAWARD) GetLoots() []*DUBBLERAWARD_RawardLoots {
	if m != nil {
		return m.Loots
	}
	return nil
}

type DUBBLERAWARD_RawardLoots struct {
	// * 可获得奖励物品ID
	CoinItemID *string `protobuf:"bytes,1,opt,def=VI_HC" json:"CoinItemID,omitempty"`
	// * 数量
	Count            *uint32 `protobuf:"varint,2,opt,def=0" json:"Count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DUBBLERAWARD_RawardLoots) Reset()         { *m = DUBBLERAWARD_RawardLoots{} }
func (m *DUBBLERAWARD_RawardLoots) String() string { return proto.CompactTextString(m) }
func (*DUBBLERAWARD_RawardLoots) ProtoMessage()    {}

const Default_DUBBLERAWARD_RawardLoots_CoinItemID string = "VI_HC"
const Default_DUBBLERAWARD_RawardLoots_Count uint32 = 0

func (m *DUBBLERAWARD_RawardLoots) GetCoinItemID() string {
	if m != nil && m.CoinItemID != nil {
		return *m.CoinItemID
	}
	return Default_DUBBLERAWARD_RawardLoots_CoinItemID
}

func (m *DUBBLERAWARD_RawardLoots) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return Default_DUBBLERAWARD_RawardLoots_Count
}

type DUBBLERAWARD_ARRAY struct {
	Items            []*DUBBLERAWARD `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *DUBBLERAWARD_ARRAY) Reset()         { *m = DUBBLERAWARD_ARRAY{} }
func (m *DUBBLERAWARD_ARRAY) String() string { return proto.CompactTextString(m) }
func (*DUBBLERAWARD_ARRAY) ProtoMessage()    {}

func (m *DUBBLERAWARD_ARRAY) GetItems() []*DUBBLERAWARD {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
