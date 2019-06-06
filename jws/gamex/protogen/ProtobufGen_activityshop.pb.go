// Code generated by protoc-gen-go.
// source: ProtobufGen_activityshop.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ACTIVITYSHOP struct {
	// * 促销活动ID
	PromotionID *uint32 `protobuf:"varint,1,req,def=0" json:"PromotionID,omitempty"`
	// * 活动类型
	PromotionType *uint32 `protobuf:"varint,2,opt,def=0" json:"PromotionType,omitempty"`
	// * 开启条件
	OpeningCondition *uint32 `protobuf:"varint,3,opt,def=0" json:"OpeningCondition,omitempty"`
	// * 开启参数
	OpeningParameters *uint32 `protobuf:"varint,4,opt,def=0" json:"OpeningParameters,omitempty"`
	// * 物品ID
	ItemID *string `protobuf:"bytes,5,opt,def=" json:"ItemID,omitempty"`
	// * 数量
	GoodsCount *uint32 `protobuf:"varint,6,opt,def=0" json:"GoodsCount,omitempty"`
	// * 购买所需物品ID
	CoinItemID *string `protobuf:"bytes,7,opt,def=VI_HC" json:"CoinItemID,omitempty"`
	// * 现价
	CurrentPrice *uint32 `protobuf:"varint,8,opt,def=0" json:"CurrentPrice,omitempty"`
	// * 原价
	OriginalCost *uint32 `protobuf:"varint,9,opt,def=0" json:"OriginalCost,omitempty"`
	// * 购买所需的VIP等级
	VIPLimit *uint32 `protobuf:"varint,10,opt,def=0" json:"VIPLimit,omitempty"`
	// * 限购数量
	CountLimit *uint32 `protobuf:"varint,11,opt,def=0" json:"CountLimit,omitempty"`
	// * 限购类型
	LimitType *string `protobuf:"bytes,12,opt,def=Forever" json:"LimitType,omitempty"`
	// * 服务器限购数量
	ServerCountLimit *uint32 `protobuf:"varint,13,opt,def=0" json:"ServerCountLimit,omitempty"`
	// * 任务获得的活跃值
	ActiveValue      *uint32 `protobuf:"varint,14,opt,def=0" json:"ActiveValue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ACTIVITYSHOP) Reset()         { *m = ACTIVITYSHOP{} }
func (m *ACTIVITYSHOP) String() string { return proto.CompactTextString(m) }
func (*ACTIVITYSHOP) ProtoMessage()    {}

const Default_ACTIVITYSHOP_PromotionID uint32 = 0
const Default_ACTIVITYSHOP_PromotionType uint32 = 0
const Default_ACTIVITYSHOP_OpeningCondition uint32 = 0
const Default_ACTIVITYSHOP_OpeningParameters uint32 = 0
const Default_ACTIVITYSHOP_GoodsCount uint32 = 0
const Default_ACTIVITYSHOP_CoinItemID string = "VI_HC"
const Default_ACTIVITYSHOP_CurrentPrice uint32 = 0
const Default_ACTIVITYSHOP_OriginalCost uint32 = 0
const Default_ACTIVITYSHOP_VIPLimit uint32 = 0
const Default_ACTIVITYSHOP_CountLimit uint32 = 0
const Default_ACTIVITYSHOP_LimitType string = "Forever"
const Default_ACTIVITYSHOP_ServerCountLimit uint32 = 0
const Default_ACTIVITYSHOP_ActiveValue uint32 = 0

func (m *ACTIVITYSHOP) GetPromotionID() uint32 {
	if m != nil && m.PromotionID != nil {
		return *m.PromotionID
	}
	return Default_ACTIVITYSHOP_PromotionID
}

func (m *ACTIVITYSHOP) GetPromotionType() uint32 {
	if m != nil && m.PromotionType != nil {
		return *m.PromotionType
	}
	return Default_ACTIVITYSHOP_PromotionType
}

func (m *ACTIVITYSHOP) GetOpeningCondition() uint32 {
	if m != nil && m.OpeningCondition != nil {
		return *m.OpeningCondition
	}
	return Default_ACTIVITYSHOP_OpeningCondition
}

func (m *ACTIVITYSHOP) GetOpeningParameters() uint32 {
	if m != nil && m.OpeningParameters != nil {
		return *m.OpeningParameters
	}
	return Default_ACTIVITYSHOP_OpeningParameters
}

func (m *ACTIVITYSHOP) GetItemID() string {
	if m != nil && m.ItemID != nil {
		return *m.ItemID
	}
	return ""
}

func (m *ACTIVITYSHOP) GetGoodsCount() uint32 {
	if m != nil && m.GoodsCount != nil {
		return *m.GoodsCount
	}
	return Default_ACTIVITYSHOP_GoodsCount
}

func (m *ACTIVITYSHOP) GetCoinItemID() string {
	if m != nil && m.CoinItemID != nil {
		return *m.CoinItemID
	}
	return Default_ACTIVITYSHOP_CoinItemID
}

func (m *ACTIVITYSHOP) GetCurrentPrice() uint32 {
	if m != nil && m.CurrentPrice != nil {
		return *m.CurrentPrice
	}
	return Default_ACTIVITYSHOP_CurrentPrice
}

func (m *ACTIVITYSHOP) GetOriginalCost() uint32 {
	if m != nil && m.OriginalCost != nil {
		return *m.OriginalCost
	}
	return Default_ACTIVITYSHOP_OriginalCost
}

func (m *ACTIVITYSHOP) GetVIPLimit() uint32 {
	if m != nil && m.VIPLimit != nil {
		return *m.VIPLimit
	}
	return Default_ACTIVITYSHOP_VIPLimit
}

func (m *ACTIVITYSHOP) GetCountLimit() uint32 {
	if m != nil && m.CountLimit != nil {
		return *m.CountLimit
	}
	return Default_ACTIVITYSHOP_CountLimit
}

func (m *ACTIVITYSHOP) GetLimitType() string {
	if m != nil && m.LimitType != nil {
		return *m.LimitType
	}
	return Default_ACTIVITYSHOP_LimitType
}

func (m *ACTIVITYSHOP) GetServerCountLimit() uint32 {
	if m != nil && m.ServerCountLimit != nil {
		return *m.ServerCountLimit
	}
	return Default_ACTIVITYSHOP_ServerCountLimit
}

func (m *ACTIVITYSHOP) GetActiveValue() uint32 {
	if m != nil && m.ActiveValue != nil {
		return *m.ActiveValue
	}
	return Default_ACTIVITYSHOP_ActiveValue
}

type ACTIVITYSHOP_ARRAY struct {
	Items            []*ACTIVITYSHOP `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ACTIVITYSHOP_ARRAY) Reset()         { *m = ACTIVITYSHOP_ARRAY{} }
func (m *ACTIVITYSHOP_ARRAY) String() string { return proto.CompactTextString(m) }
func (*ACTIVITYSHOP_ARRAY) ProtoMessage()    {}

func (m *ACTIVITYSHOP_ARRAY) GetItems() []*ACTIVITYSHOP {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}