// Code generated by protoc-gen-go.
// source: ProtobufGen_fbconfig.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type FBCONFIG struct {
	// * 节日类型ID
	ActivityID *uint32 `protobuf:"varint,1,opt,def=0" json:"ActivityID,omitempty"`
	// * 距离结束X分钟不能挑战
	OnlyShop *uint32 `protobuf:"varint,2,opt,def=0" json:"OnlyShop,omitempty"`
	// * 采用的Boss画像
	BossPic *string `protobuf:"bytes,3,opt,def=" json:"BossPic,omitempty"`
	// * 采用的活动图标
	TownIcon *string `protobuf:"bytes,20,opt,def=" json:"TownIcon,omitempty"`
	// * 采用的标题IDS
	BossTittle *string `protobuf:"bytes,4,opt,def=" json:"BossTittle,omitempty"`
	// * 采用的击杀文案
	BossKillCountIDS *string `protobuf:"bytes,5,opt,def=" json:"BossKillCountIDS,omitempty"`
	// * 采用的商店名称
	ShopNameIDS *string `protobuf:"bytes,6,opt,def=" json:"ShopNameIDS,omitempty"`
	// * BOSS展示名称
	BossNameIDS *string `protobuf:"bytes,7,opt,def=" json:"BossNameIDS,omitempty"`
	// * 广告语IDS（在商店按钮上面）
	JieShaoIDS *string `protobuf:"bytes,8,opt,def=" json:"JieShaoIDS,omitempty"`
	// * 击杀了什么的IDS
	SkillWhatIDS *string `protobuf:"bytes,9,opt,def=" json:"SkillWhatIDS,omitempty"`
	// * 挑战消耗的道具
	ChallengeCost *string `protobuf:"bytes,10,opt,def=" json:"ChallengeCost,omitempty"`
	// * 挑战消耗的道具个数
	ChallengeCostNum *uint32 `protobuf:"varint,11,opt,def=0" json:"ChallengeCostNum,omitempty"`
	// * 翻倍策略1（翻几次）
	MoreRewardOne *uint32 `protobuf:"varint,12,opt,def=0" json:"MoreRewardOne,omitempty"`
	// * 翻倍消耗的道具
	MoreRewardOneCost *string `protobuf:"bytes,13,opt,def=" json:"MoreRewardOneCost,omitempty"`
	// * 数量
	MoreRewardOneCostNum *uint32 `protobuf:"varint,14,opt,def=0" json:"MoreRewardOneCostNum,omitempty"`
	// * 显示的翻倍IDS
	MoreRewardOneIDS *string `protobuf:"bytes,15,opt,def=" json:"MoreRewardOneIDS,omitempty"`
	// * 翻倍策略2（翻几次）
	MoreRewardTwo *uint32 `protobuf:"varint,16,opt,def=0" json:"MoreRewardTwo,omitempty"`
	// * 翻倍消耗的道具
	MoreRewardTwoCost *string `protobuf:"bytes,17,opt,def=" json:"MoreRewardTwoCost,omitempty"`
	// * 数量
	MoreRewardTwoCostNum *uint32 `protobuf:"varint,18,opt,def=0" json:"MoreRewardTwoCostNum,omitempty"`
	// * 显示的翻倍IDS
	MoreRewardTwoIDS *string `protobuf:"bytes,19,opt,def=" json:"MoreRewardTwoIDS,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FBCONFIG) Reset()         { *m = FBCONFIG{} }
func (m *FBCONFIG) String() string { return proto.CompactTextString(m) }
func (*FBCONFIG) ProtoMessage()    {}

const Default_FBCONFIG_ActivityID uint32 = 0
const Default_FBCONFIG_OnlyShop uint32 = 0
const Default_FBCONFIG_ChallengeCostNum uint32 = 0
const Default_FBCONFIG_MoreRewardOne uint32 = 0
const Default_FBCONFIG_MoreRewardOneCostNum uint32 = 0
const Default_FBCONFIG_MoreRewardTwo uint32 = 0
const Default_FBCONFIG_MoreRewardTwoCostNum uint32 = 0

func (m *FBCONFIG) GetActivityID() uint32 {
	if m != nil && m.ActivityID != nil {
		return *m.ActivityID
	}
	return Default_FBCONFIG_ActivityID
}

func (m *FBCONFIG) GetOnlyShop() uint32 {
	if m != nil && m.OnlyShop != nil {
		return *m.OnlyShop
	}
	return Default_FBCONFIG_OnlyShop
}

func (m *FBCONFIG) GetBossPic() string {
	if m != nil && m.BossPic != nil {
		return *m.BossPic
	}
	return ""
}

func (m *FBCONFIG) GetTownIcon() string {
	if m != nil && m.TownIcon != nil {
		return *m.TownIcon
	}
	return ""
}

func (m *FBCONFIG) GetBossTittle() string {
	if m != nil && m.BossTittle != nil {
		return *m.BossTittle
	}
	return ""
}

func (m *FBCONFIG) GetBossKillCountIDS() string {
	if m != nil && m.BossKillCountIDS != nil {
		return *m.BossKillCountIDS
	}
	return ""
}

func (m *FBCONFIG) GetShopNameIDS() string {
	if m != nil && m.ShopNameIDS != nil {
		return *m.ShopNameIDS
	}
	return ""
}

func (m *FBCONFIG) GetBossNameIDS() string {
	if m != nil && m.BossNameIDS != nil {
		return *m.BossNameIDS
	}
	return ""
}

func (m *FBCONFIG) GetJieShaoIDS() string {
	if m != nil && m.JieShaoIDS != nil {
		return *m.JieShaoIDS
	}
	return ""
}

func (m *FBCONFIG) GetSkillWhatIDS() string {
	if m != nil && m.SkillWhatIDS != nil {
		return *m.SkillWhatIDS
	}
	return ""
}

func (m *FBCONFIG) GetChallengeCost() string {
	if m != nil && m.ChallengeCost != nil {
		return *m.ChallengeCost
	}
	return ""
}

func (m *FBCONFIG) GetChallengeCostNum() uint32 {
	if m != nil && m.ChallengeCostNum != nil {
		return *m.ChallengeCostNum
	}
	return Default_FBCONFIG_ChallengeCostNum
}

func (m *FBCONFIG) GetMoreRewardOne() uint32 {
	if m != nil && m.MoreRewardOne != nil {
		return *m.MoreRewardOne
	}
	return Default_FBCONFIG_MoreRewardOne
}

func (m *FBCONFIG) GetMoreRewardOneCost() string {
	if m != nil && m.MoreRewardOneCost != nil {
		return *m.MoreRewardOneCost
	}
	return ""
}

func (m *FBCONFIG) GetMoreRewardOneCostNum() uint32 {
	if m != nil && m.MoreRewardOneCostNum != nil {
		return *m.MoreRewardOneCostNum
	}
	return Default_FBCONFIG_MoreRewardOneCostNum
}

func (m *FBCONFIG) GetMoreRewardOneIDS() string {
	if m != nil && m.MoreRewardOneIDS != nil {
		return *m.MoreRewardOneIDS
	}
	return ""
}

func (m *FBCONFIG) GetMoreRewardTwo() uint32 {
	if m != nil && m.MoreRewardTwo != nil {
		return *m.MoreRewardTwo
	}
	return Default_FBCONFIG_MoreRewardTwo
}

func (m *FBCONFIG) GetMoreRewardTwoCost() string {
	if m != nil && m.MoreRewardTwoCost != nil {
		return *m.MoreRewardTwoCost
	}
	return ""
}

func (m *FBCONFIG) GetMoreRewardTwoCostNum() uint32 {
	if m != nil && m.MoreRewardTwoCostNum != nil {
		return *m.MoreRewardTwoCostNum
	}
	return Default_FBCONFIG_MoreRewardTwoCostNum
}

func (m *FBCONFIG) GetMoreRewardTwoIDS() string {
	if m != nil && m.MoreRewardTwoIDS != nil {
		return *m.MoreRewardTwoIDS
	}
	return ""
}

type FBCONFIG_ARRAY struct {
	Items            []*FBCONFIG `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *FBCONFIG_ARRAY) Reset()         { *m = FBCONFIG_ARRAY{} }
func (m *FBCONFIG_ARRAY) String() string { return proto.CompactTextString(m) }
func (*FBCONFIG_ARRAY) ProtoMessage()    {}

func (m *FBCONFIG_ARRAY) GetItems() []*FBCONFIG {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
