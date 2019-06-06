// Code generated by protoc-gen-go.
// source: ProtobufGen_storegroup.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type STOREGROUP struct {
	// * 商品唯一ID
	StoreGoodsID *uint32 `protobuf:"varint,11,req,def=0" json:"StoreGoodsID,omitempty"`
	// * 物品组ID
	GroupID *uint32 `protobuf:"varint,1,req,def=0" json:"GroupID,omitempty"`
	// * 权重
	Weight *uint32 `protobuf:"varint,2,opt,def=0" json:"Weight,omitempty"`
	// * 是否为珍贵道具
	IsTreasure *uint32 `protobuf:"varint,3,opt,def=0" json:"IsTreasure,omitempty"`
	// * VIP限制
	VIPLimit *uint32 `protobuf:"varint,8,opt,def=0" json:"VIPLimit,omitempty"`
	// * 物品ID
	GoodsID *string `protobuf:"bytes,4,opt,def=" json:"GoodsID,omitempty"`
	// * 数量
	GoodsCount *uint32 `protobuf:"varint,5,opt,def=0" json:"GoodsCount,omitempty"`
	// * 购买所需物品ID
	CoinItemID *string `protobuf:"bytes,6,opt,def=VI_HC" json:"CoinItemID,omitempty"`
	// * 数量
	CoinItemCount *uint32 `protobuf:"varint,7,opt,def=0" json:"CoinItemCount,omitempty"`
	// * 原价数量
	OriginalCount *uint32 `protobuf:"varint,9,opt,def=0" json:"OriginalCount,omitempty"`
	// * 折扣
	Discount         *uint32 `protobuf:"varint,10,opt,def=0" json:"Discount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *STOREGROUP) Reset()         { *m = STOREGROUP{} }
func (m *STOREGROUP) String() string { return proto.CompactTextString(m) }
func (*STOREGROUP) ProtoMessage()    {}

const Default_STOREGROUP_StoreGoodsID uint32 = 0
const Default_STOREGROUP_GroupID uint32 = 0
const Default_STOREGROUP_Weight uint32 = 0
const Default_STOREGROUP_IsTreasure uint32 = 0
const Default_STOREGROUP_VIPLimit uint32 = 0
const Default_STOREGROUP_GoodsCount uint32 = 0
const Default_STOREGROUP_CoinItemID string = "VI_HC"
const Default_STOREGROUP_CoinItemCount uint32 = 0
const Default_STOREGROUP_OriginalCount uint32 = 0
const Default_STOREGROUP_Discount uint32 = 0

func (m *STOREGROUP) GetStoreGoodsID() uint32 {
	if m != nil && m.StoreGoodsID != nil {
		return *m.StoreGoodsID
	}
	return Default_STOREGROUP_StoreGoodsID
}

func (m *STOREGROUP) GetGroupID() uint32 {
	if m != nil && m.GroupID != nil {
		return *m.GroupID
	}
	return Default_STOREGROUP_GroupID
}

func (m *STOREGROUP) GetWeight() uint32 {
	if m != nil && m.Weight != nil {
		return *m.Weight
	}
	return Default_STOREGROUP_Weight
}

func (m *STOREGROUP) GetIsTreasure() uint32 {
	if m != nil && m.IsTreasure != nil {
		return *m.IsTreasure
	}
	return Default_STOREGROUP_IsTreasure
}

func (m *STOREGROUP) GetVIPLimit() uint32 {
	if m != nil && m.VIPLimit != nil {
		return *m.VIPLimit
	}
	return Default_STOREGROUP_VIPLimit
}

func (m *STOREGROUP) GetGoodsID() string {
	if m != nil && m.GoodsID != nil {
		return *m.GoodsID
	}
	return ""
}

func (m *STOREGROUP) GetGoodsCount() uint32 {
	if m != nil && m.GoodsCount != nil {
		return *m.GoodsCount
	}
	return Default_STOREGROUP_GoodsCount
}

func (m *STOREGROUP) GetCoinItemID() string {
	if m != nil && m.CoinItemID != nil {
		return *m.CoinItemID
	}
	return Default_STOREGROUP_CoinItemID
}

func (m *STOREGROUP) GetCoinItemCount() uint32 {
	if m != nil && m.CoinItemCount != nil {
		return *m.CoinItemCount
	}
	return Default_STOREGROUP_CoinItemCount
}

func (m *STOREGROUP) GetOriginalCount() uint32 {
	if m != nil && m.OriginalCount != nil {
		return *m.OriginalCount
	}
	return Default_STOREGROUP_OriginalCount
}

func (m *STOREGROUP) GetDiscount() uint32 {
	if m != nil && m.Discount != nil {
		return *m.Discount
	}
	return Default_STOREGROUP_Discount
}

type STOREGROUP_ARRAY struct {
	Items            []*STOREGROUP `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *STOREGROUP_ARRAY) Reset()         { *m = STOREGROUP_ARRAY{} }
func (m *STOREGROUP_ARRAY) String() string { return proto.CompactTextString(m) }
func (*STOREGROUP_ARRAY) ProtoMessage()    {}

func (m *STOREGROUP_ARRAY) GetItems() []*STOREGROUP {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}