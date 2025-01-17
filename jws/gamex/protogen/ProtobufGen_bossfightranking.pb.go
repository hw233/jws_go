// Code generated by protoc-gen-go.
// source: ProtobufGen_bossfightranking.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type BOSSFIGHTRANKING struct {
	// * 玩家每日22点结算功勋时的排名，排名只记录到300,301开始属于未进入排名，不发放排名奖励
	Ranking *uint32 `protobuf:"varint,1,req,def=0" json:"Ranking,omitempty"`
	// * 每日累计功勋排行-难度1-代币奖励
	AccfeatsLv1_BOSSCurrency *uint32 `protobuf:"varint,2,opt,name=Accfeats_lv1_BOSSCurrency,def=0" json:"Accfeats_lv1_BOSSCurrency,omitempty"`
	// * 每日累计功勋排行-难度1-硬通奖励
	AccfeatsLv1_HC *uint32 `protobuf:"varint,3,opt,name=Accfeats_lv1_HC,def=0" json:"Accfeats_lv1_HC,omitempty"`
	// * 每日累计功勋排行-难度2-代币奖励
	AccfeatsLv2_BOSSCurrency *uint32 `protobuf:"varint,4,opt,name=Accfeats_lv2_BOSSCurrency,def=0" json:"Accfeats_lv2_BOSSCurrency,omitempty"`
	// * 每日累计功勋排行-难度2-硬通奖励
	AccfeatsLv2_HC *uint32 `protobuf:"varint,5,opt,name=Accfeats_lv2_HC,def=0" json:"Accfeats_lv2_HC,omitempty"`
	// * 每日单场最高功勋排行-难度1-代币奖励
	HfeatsLv1_BOSSCurrency *uint32 `protobuf:"varint,6,opt,name=Hfeats_lv1_BOSSCurrency,def=0" json:"Hfeats_lv1_BOSSCurrency,omitempty"`
	// * 每日单场最高功勋排行-难度1-硬通奖励
	HfeatsLv1_HC *uint32 `protobuf:"varint,7,opt,name=Hfeats_lv1_HC,def=0" json:"Hfeats_lv1_HC,omitempty"`
	// * 每日单场最高功勋排行-难度2-代币奖励
	HfeatsLv2_BOSSCurrency *uint32 `protobuf:"varint,8,opt,name=Hfeats_lv2_BOSSCurrency,def=0" json:"Hfeats_lv2_BOSSCurrency,omitempty"`
	// * 每日单场最高功勋排行-难度2-硬通奖励
	HfeatsLv2_HC *uint32 `protobuf:"varint,9,opt,name=Hfeats_lv2_HC,def=0" json:"Hfeats_lv2_HC,omitempty"`
	// * 累计积分排行服务器推送IDS
	AccfeatsPushIDS *string `protobuf:"bytes,10,opt,def=" json:"AccfeatsPushIDS,omitempty"`
	// * 单场排行服务器推送IDS
	HfeatsPushIDS    *string `protobuf:"bytes,11,opt,def=" json:"HfeatsPushIDS,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BOSSFIGHTRANKING) Reset()         { *m = BOSSFIGHTRANKING{} }
func (m *BOSSFIGHTRANKING) String() string { return proto.CompactTextString(m) }
func (*BOSSFIGHTRANKING) ProtoMessage()    {}

const Default_BOSSFIGHTRANKING_Ranking uint32 = 0
const Default_BOSSFIGHTRANKING_AccfeatsLv1_BOSSCurrency uint32 = 0
const Default_BOSSFIGHTRANKING_AccfeatsLv1_HC uint32 = 0
const Default_BOSSFIGHTRANKING_AccfeatsLv2_BOSSCurrency uint32 = 0
const Default_BOSSFIGHTRANKING_AccfeatsLv2_HC uint32 = 0
const Default_BOSSFIGHTRANKING_HfeatsLv1_BOSSCurrency uint32 = 0
const Default_BOSSFIGHTRANKING_HfeatsLv1_HC uint32 = 0
const Default_BOSSFIGHTRANKING_HfeatsLv2_BOSSCurrency uint32 = 0
const Default_BOSSFIGHTRANKING_HfeatsLv2_HC uint32 = 0

func (m *BOSSFIGHTRANKING) GetRanking() uint32 {
	if m != nil && m.Ranking != nil {
		return *m.Ranking
	}
	return Default_BOSSFIGHTRANKING_Ranking
}

func (m *BOSSFIGHTRANKING) GetAccfeatsLv1_BOSSCurrency() uint32 {
	if m != nil && m.AccfeatsLv1_BOSSCurrency != nil {
		return *m.AccfeatsLv1_BOSSCurrency
	}
	return Default_BOSSFIGHTRANKING_AccfeatsLv1_BOSSCurrency
}

func (m *BOSSFIGHTRANKING) GetAccfeatsLv1_HC() uint32 {
	if m != nil && m.AccfeatsLv1_HC != nil {
		return *m.AccfeatsLv1_HC
	}
	return Default_BOSSFIGHTRANKING_AccfeatsLv1_HC
}

func (m *BOSSFIGHTRANKING) GetAccfeatsLv2_BOSSCurrency() uint32 {
	if m != nil && m.AccfeatsLv2_BOSSCurrency != nil {
		return *m.AccfeatsLv2_BOSSCurrency
	}
	return Default_BOSSFIGHTRANKING_AccfeatsLv2_BOSSCurrency
}

func (m *BOSSFIGHTRANKING) GetAccfeatsLv2_HC() uint32 {
	if m != nil && m.AccfeatsLv2_HC != nil {
		return *m.AccfeatsLv2_HC
	}
	return Default_BOSSFIGHTRANKING_AccfeatsLv2_HC
}

func (m *BOSSFIGHTRANKING) GetHfeatsLv1_BOSSCurrency() uint32 {
	if m != nil && m.HfeatsLv1_BOSSCurrency != nil {
		return *m.HfeatsLv1_BOSSCurrency
	}
	return Default_BOSSFIGHTRANKING_HfeatsLv1_BOSSCurrency
}

func (m *BOSSFIGHTRANKING) GetHfeatsLv1_HC() uint32 {
	if m != nil && m.HfeatsLv1_HC != nil {
		return *m.HfeatsLv1_HC
	}
	return Default_BOSSFIGHTRANKING_HfeatsLv1_HC
}

func (m *BOSSFIGHTRANKING) GetHfeatsLv2_BOSSCurrency() uint32 {
	if m != nil && m.HfeatsLv2_BOSSCurrency != nil {
		return *m.HfeatsLv2_BOSSCurrency
	}
	return Default_BOSSFIGHTRANKING_HfeatsLv2_BOSSCurrency
}

func (m *BOSSFIGHTRANKING) GetHfeatsLv2_HC() uint32 {
	if m != nil && m.HfeatsLv2_HC != nil {
		return *m.HfeatsLv2_HC
	}
	return Default_BOSSFIGHTRANKING_HfeatsLv2_HC
}

func (m *BOSSFIGHTRANKING) GetAccfeatsPushIDS() string {
	if m != nil && m.AccfeatsPushIDS != nil {
		return *m.AccfeatsPushIDS
	}
	return ""
}

func (m *BOSSFIGHTRANKING) GetHfeatsPushIDS() string {
	if m != nil && m.HfeatsPushIDS != nil {
		return *m.HfeatsPushIDS
	}
	return ""
}

type BOSSFIGHTRANKING_ARRAY struct {
	Items            []*BOSSFIGHTRANKING `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *BOSSFIGHTRANKING_ARRAY) Reset()         { *m = BOSSFIGHTRANKING_ARRAY{} }
func (m *BOSSFIGHTRANKING_ARRAY) String() string { return proto.CompactTextString(m) }
func (*BOSSFIGHTRANKING_ARRAY) ProtoMessage()    {}

func (m *BOSSFIGHTRANKING_ARRAY) GetItems() []*BOSSFIGHTRANKING {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
