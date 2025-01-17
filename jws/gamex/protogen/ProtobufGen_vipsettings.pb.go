// Code generated by protoc-gen-go.
// source: ProtobufGen_vipsettings.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type VIPSETTINGS struct {
	// * VIP等级
	VIPLevel *uint32 `protobuf:"varint,1,req,def=0" json:"VIPLevel,omitempty"`
	// * 充值额度
	RMBpoints *uint32 `protobuf:"varint,2,opt,def=0" json:"RMBpoints,omitempty"`
	// * 图标颜色
	VIPiconEnum *uint32 `protobuf:"varint,57,opt,def=0" json:"VIPiconEnum,omitempty"`
	// * 体力购买次数限制
	EnergyPurchaseLimit *uint32 `protobuf:"varint,3,opt,def=0" json:"EnergyPurchaseLimit,omitempty"`
	// * 包子购买次数限制
	HighEnergyLimitt *uint32 `protobuf:"varint,63,opt,def=0" json:"HighEnergyLimitt,omitempty"`
	// * 一键吃包权限
	EatBaoziQuik *uint32 `protobuf:"varint,64,opt,def=0" json:"EatBaoziQuik,omitempty"`
	// * SC购买次数限制
	SCPurchaseLimit *uint32 `protobuf:"varint,4,opt,def=0" json:"SCPurchaseLimit,omitempty"`
	// * 令牌购买次数限制
	SprintPurchaseLimit *int32 `protobuf:"varint,5,opt,def=-1" json:"SprintPurchaseLimit,omitempty"`
	// * 竞技场购买次数限制
	PVPTimeLimit *int32 `protobuf:"varint,6,opt,def=-1" json:"PVPTimeLimit,omitempty"`
	// * 比武场购买次数限制
	TPVPTimeLimit *int32 `protobuf:"varint,7,opt,def=-1" json:"TPVPTimeLimit,omitempty"`
	// * 无双竞技场购买次数限制
	WSPVPTimeLimit *int32 `protobuf:"varint,75,opt,def=-1" json:"WSPVPTimeLimit,omitempty"`
	// * VIP商店刷新次数限制
	StoreRefreshLimit *uint32 `protobuf:"varint,8,opt,def=0" json:"StoreRefreshLimit,omitempty"`
	// * 1v1竞技场商店刷新次数限制
	PvpStoreRefreshLimit *uint32 `protobuf:"varint,68,opt,def=0" json:"PvpStoreRefreshLimit,omitempty"`
	// * 3v3竞技场商店刷新次数限制
	TpvpStoreRefreshLimit *uint32 `protobuf:"varint,69,opt,def=0" json:"TpvpStoreRefreshLimit,omitempty"`
	// * 无双竞技场商店刷新次数限制
	WSPVPStoreRefreshLimit *uint32 `protobuf:"varint,76,opt,def=0" json:"WSPVPStoreRefreshLimit,omitempty"`
	// * 名将试炼商店刷新次数限制
	BossStoreRefreshLimit *uint32 `protobuf:"varint,70,opt,def=0" json:"BossStoreRefreshLimit,omitempty"`
	// * 军团商店刷新次数限制
	GuildStoreRefreshLimit *uint32 `protobuf:"varint,71,opt,def=0" json:"GuildStoreRefreshLimit,omitempty"`
	// * 远征商店刷新次数限制
	ExpeditionStoreRefreshLimit *uint32 `protobuf:"varint,72,opt,def=0" json:"ExpeditionStoreRefreshLimit,omitempty"`
	// * 出奇制胜商店刷新次数限制
	HDPStoreRefreshLimit *uint32 `protobuf:"varint,74,opt,def=0" json:"HDPStoreRefreshLimit,omitempty"`
	// * 灵宠商店刷新次数限制
	PetStoreRefreshLimit *uint32 `protobuf:"varint,88,opt,def=0" json:"PetStoreRefreshLimit,omitempty"`
	// * 是否开启名将投资，0=不开启，1=开启
	HeroInvestLimit *uint32 `protobuf:"varint,77,opt,def=0" json:"HeroInvestLimit,omitempty"`
	// * 劫营夺粮运粮次数
	Cropstimes *uint32 `protobuf:"varint,78,opt,def=2" json:"Cropstimes,omitempty"`
	// * 劫营夺粮运粮时长（分钟）
	CropsDuration *uint32 `protobuf:"varint,79,opt,def=0" json:"CropsDuration,omitempty"`
	// * 劫营夺粮掠夺次数
	Robtimes *uint32 `protobuf:"varint,80,opt,def=3" json:"Robtimes,omitempty"`
	// * 劫营夺粮协助次数
	Assisttimes *uint32 `protobuf:"varint,81,opt,def=2" json:"Assisttimes,omitempty"`
	// * 挑战世界boss的次数
	WorldBosstimes *uint32 `protobuf:"varint,83,opt,def=1" json:"WorldBosstimes,omitempty"`
	// * 普通碎片回收宝箱每日次数
	SurplusGacha3 *uint32 `protobuf:"varint,84,opt,def=0" json:"SurplusGacha3,omitempty"`
	// * 精锐碎片回收宝箱每日次数
	SurplusGacha4 *uint32 `protobuf:"varint,85,opt,def=0" json:"SurplusGacha4,omitempty"`
	// * 无双碎片回收宝箱每日次数
	SurplusGacha5 *uint32 `protobuf:"varint,86,opt,def=0" json:"SurplusGacha5,omitempty"`
	// * VIP描述
	VIPDescripition *string `protobuf:"bytes,9,opt,def=" json:"VIPDescripition,omitempty"`
	// * 特权ID
	PrivilegeID *int32 `protobuf:"varint,10,opt,def=0" json:"PrivilegeID,omitempty"`
	// * 开启扫荡功能
	Sweep *uint32 `protobuf:"varint,11,opt,def=0" json:"Sweep,omitempty"`
	// * 开启普通副本非三星扫荡
	NormalSweep *uint32 `protobuf:"varint,12,opt,def=0" json:"NormalSweep,omitempty"`
	// * 名将乱入敌将出现概率提升
	BossRatioUp *uint32 `protobuf:"varint,13,opt,def=0" json:"BossRatioUp,omitempty"`
	// * 开启精英副本非三星扫荡
	EliteSweep *uint32 `protobuf:"varint,14,opt,def=0" json:"EliteSweep,omitempty"`
	// * 开启地狱难度副本非三星扫荡
	HardSweep *uint32 `protobuf:"varint,15,opt,def=0" json:"HardSweep,omitempty"`
	// * 天命关收入加成%
	DCLevelAdd *float32 `protobuf:"fixed32,16,opt,def=0" json:"DCLevelAdd,omitempty"`
	// * 精铁关收入加成%
	IronLevelAdd *float32 `protobuf:"fixed32,17,opt,def=0" json:"IronLevelAdd,omitempty"`
	// * 金币关收入加成%
	GoldLevelAdd *float32 `protobuf:"fixed32,18,opt,def=0" json:"GoldLevelAdd,omitempty"`
	// * 精英关卡重置次数
	ElitePurchase *uint32 `protobuf:"varint,19,opt,def=0" json:"ElitePurchase,omitempty"`
	// * 地狱难度关卡重置次数
	HardPurchase *uint32 `protobuf:"varint,20,opt,def=0" json:"HardPurchase,omitempty"`
	// * 资源关扫荡的权限
	SweepResource *uint32 `protobuf:"varint,21,opt,def=0" json:"SweepResource,omitempty"`
	// * 名将乱入A级BOSS出现比率提升
	ABossRatioUp *uint32 `protobuf:"varint,22,opt,def=0" json:"ABossRatioUp,omitempty"`
	// * 开启名将乱入扫荡
	BOSSFightSweep *uint32 `protobuf:"varint,23,opt,def=0" json:"BOSSFightSweep,omitempty"`
	// * 名将乱入S级敌将出现比率提升
	SBossRatioUp *uint32 `protobuf:"varint,24,opt,def=0" json:"SBossRatioUp,omitempty"`
	// * 每日祈福签到次数
	GuildSignLimit *uint32 `protobuf:"varint,25,opt,def=0" json:"GuildSignLimit,omitempty"`
	// *
	PVPPurchaseLimit *int32 `protobuf:"varint,26,opt,def=-1" json:"PVPPurchaseLimit,omitempty"`
	// * 使用钻石升星的次数限制
	StarupUseHCLimit *uint32 `protobuf:"varint,27,opt,def=0" json:"StarupUseHCLimit,omitempty"`
	// * 是否可以用VIPgacha
	VIPGachaLimit *uint32 `protobuf:"varint,28,opt,def=0" json:"VIPGachaLimit,omitempty"`
	// * 是否可以看到VIPgacha
	VIPGachaVisible *uint32 `protobuf:"varint,67,opt,def=0" json:"VIPGachaVisible,omitempty"`
	// * 军团膜拜次数
	GuildWorshipLimit *uint32 `protobuf:"varint,73,opt,def=2" json:"GuildWorshipLimit,omitempty"`
	// * 可勾选双人必中红箱权限
	TeamBossRedBox *uint32 `protobuf:"varint,87,opt,def=0" json:"TeamBossRedBox,omitempty"`
	// * VIP特权图片1
	PrivilegeImg1 *string `protobuf:"bytes,29,opt,def=" json:"PrivilegeImg1,omitempty"`
	// * 在VIP特权图片1上方的描述
	PrivilegemgUpDes1 *string `protobuf:"bytes,30,opt,def=" json:"PrivilegemgUpDes1,omitempty"`
	// * 在VIP特权图片1下方的描述
	PrivilegemgDownDes1 *string `protobuf:"bytes,31,opt,def=" json:"PrivilegemgDownDes1,omitempty"`
	// * VIP特权图片2
	PrivilegeImg2 *string `protobuf:"bytes,32,opt,def=" json:"PrivilegeImg2,omitempty"`
	// * 在VIP特权图片2上方的描述
	PrivilegemgUpDes2 *string `protobuf:"bytes,33,opt,def=" json:"PrivilegemgUpDes2,omitempty"`
	// * 在VIP特权图片2下方的描述
	PrivilegemgDownDes2 *string `protobuf:"bytes,34,opt,def=" json:"PrivilegemgDownDes2,omitempty"`
	// * VIP特权图片3
	PrivilegeImg3 *string `protobuf:"bytes,35,opt,def=" json:"PrivilegeImg3,omitempty"`
	// * 在VIP特权图片3上方的描述
	PrivilegemgUpDes3 *string `protobuf:"bytes,36,opt,def=" json:"PrivilegemgUpDes3,omitempty"`
	// * 在VIP特权图片3下方的描述
	PrivilegemgDownDes3 *string `protobuf:"bytes,37,opt,def=" json:"PrivilegemgDownDes3,omitempty"`
	// * VIP特权图片4
	PrivilegeImg4 *string `protobuf:"bytes,38,opt,def=" json:"PrivilegeImg4,omitempty"`
	// * 在VIP特权图片4上方的描述
	PrivilegemgUpDes4 *string `protobuf:"bytes,39,opt,def=" json:"PrivilegemgUpDes4,omitempty"`
	// * 在VIP特权图片4下方的描述
	PrivilegemgDownDes4 *string `protobuf:"bytes,40,opt,def=" json:"PrivilegemgDownDes4,omitempty"`
	// * 日常奖励1
	VIPDailyGift1 *string `protobuf:"bytes,41,opt,def=" json:"VIPDailyGift1,omitempty"`
	// * 数量1
	Count1 *uint32 `protobuf:"varint,42,opt,def=0" json:"Count1,omitempty"`
	// * 日常奖励2
	VIPDailyGift2 *string `protobuf:"bytes,43,opt,def=" json:"VIPDailyGift2,omitempty"`
	// * 数量2
	Count2 *uint32 `protobuf:"varint,44,opt,def=0" json:"Count2,omitempty"`
	// * 日常奖励3
	VIPDailyGift3 *string `protobuf:"bytes,45,opt,def=" json:"VIPDailyGift3,omitempty"`
	// * 数量3
	Count3 *uint32 `protobuf:"varint,46,opt,def=0" json:"Count3,omitempty"`
	// * 日常奖励4
	VIPDailyGift4 *string `protobuf:"bytes,47,opt,def=" json:"VIPDailyGift4,omitempty"`
	// * 数量4
	Count4 *uint32 `protobuf:"varint,48,opt,def=0" json:"Count4,omitempty"`
	// * 对应VIP等级弹出的宣传画
	VIPPoster *string `protobuf:"bytes,49,opt,def=" json:"VIPPoster,omitempty"`
	// * 购买成长基金资格
	GrowFund *uint32 `protobuf:"varint,50,opt,def=0" json:"GrowFund,omitempty"`
	// * 开启十连扫荡
	TenSweep *uint32 `protobuf:"varint,51,opt,def=0" json:"TenSweep,omitempty"`
	// * 名将试炼扫荡
	BossFightSweep *uint32 `protobuf:"varint,53,opt,def=0" json:"BossFightSweep,omitempty"`
	// * 金币关扫荡
	GoldLevelSweep *uint32 `protobuf:"varint,54,opt,def=0" json:"GoldLevelSweep,omitempty"`
	// * 天命关扫荡
	DestinyLevelSweep *uint32 `protobuf:"varint,55,opt,def=0" json:"DestinyLevelSweep,omitempty"`
	// * 精铁关扫荡
	IronLevelSweep *uint32 `protobuf:"varint,56,opt,def=0" json:"IronLevelSweep,omitempty"`
	// * 烽火普通扫荡资格
	FengHuoNormalSweep *uint32 `protobuf:"varint,65,opt,def=0" json:"FengHuoNormalSweep,omitempty"`
	// * 烽火燎原高级扫荡资格
	FengHuoSeniorSweep *uint32 `protobuf:"varint,58,opt,def=0" json:"FengHuoSeniorSweep,omitempty"`
	// * 兵临城下鼓舞开启限制
	GEBuffEncourage *uint32 `protobuf:"varint,59,opt,def=0" json:"GEBuffEncourage,omitempty"`
	// * 神兽高级培养免费次数（0代表不可培养）
	DGVIPTrain *uint32 `protobuf:"varint,60,opt,def=0" json:"DGVIPTrain,omitempty"`
	// * 神兽连续培养次数开启
	DGVIPContinuityTrain *uint32 `protobuf:"varint,61,opt,def=0" json:"DGVIPContinuityTrain,omitempty"`
	// * 远征每日重置次数（累积上限）
	ExpeditionTime     *uint32                     `protobuf:"varint,66,opt,def=0" json:"ExpeditionTime,omitempty"`
	VIP_Table          []*VIPSETTINGS_VIP          `protobuf:"bytes,52,rep" json:"VIP_Table,omitempty"`
	StoreRefresh_Table []*VIPSETTINGS_StoreRefresh `protobuf:"bytes,82,rep" json:"StoreRefresh_Table,omitempty"`
	XXX_unrecognized   []byte                      `json:"-"`
}

func (m *VIPSETTINGS) Reset()         { *m = VIPSETTINGS{} }
func (m *VIPSETTINGS) String() string { return proto.CompactTextString(m) }
func (*VIPSETTINGS) ProtoMessage()    {}

const Default_VIPSETTINGS_VIPLevel uint32 = 0
const Default_VIPSETTINGS_RMBpoints uint32 = 0
const Default_VIPSETTINGS_VIPiconEnum uint32 = 0
const Default_VIPSETTINGS_EnergyPurchaseLimit uint32 = 0
const Default_VIPSETTINGS_HighEnergyLimitt uint32 = 0
const Default_VIPSETTINGS_EatBaoziQuik uint32 = 0
const Default_VIPSETTINGS_SCPurchaseLimit uint32 = 0
const Default_VIPSETTINGS_SprintPurchaseLimit int32 = -1
const Default_VIPSETTINGS_PVPTimeLimit int32 = -1
const Default_VIPSETTINGS_TPVPTimeLimit int32 = -1
const Default_VIPSETTINGS_WSPVPTimeLimit int32 = -1
const Default_VIPSETTINGS_StoreRefreshLimit uint32 = 0
const Default_VIPSETTINGS_PvpStoreRefreshLimit uint32 = 0
const Default_VIPSETTINGS_TpvpStoreRefreshLimit uint32 = 0
const Default_VIPSETTINGS_WSPVPStoreRefreshLimit uint32 = 0
const Default_VIPSETTINGS_BossStoreRefreshLimit uint32 = 0
const Default_VIPSETTINGS_GuildStoreRefreshLimit uint32 = 0
const Default_VIPSETTINGS_ExpeditionStoreRefreshLimit uint32 = 0
const Default_VIPSETTINGS_HDPStoreRefreshLimit uint32 = 0
const Default_VIPSETTINGS_PetStoreRefreshLimit uint32 = 0
const Default_VIPSETTINGS_HeroInvestLimit uint32 = 0
const Default_VIPSETTINGS_Cropstimes uint32 = 2
const Default_VIPSETTINGS_CropsDuration uint32 = 0
const Default_VIPSETTINGS_Robtimes uint32 = 3
const Default_VIPSETTINGS_Assisttimes uint32 = 2
const Default_VIPSETTINGS_WorldBosstimes uint32 = 1
const Default_VIPSETTINGS_SurplusGacha3 uint32 = 0
const Default_VIPSETTINGS_SurplusGacha4 uint32 = 0
const Default_VIPSETTINGS_SurplusGacha5 uint32 = 0
const Default_VIPSETTINGS_PrivilegeID int32 = 0
const Default_VIPSETTINGS_Sweep uint32 = 0
const Default_VIPSETTINGS_NormalSweep uint32 = 0
const Default_VIPSETTINGS_BossRatioUp uint32 = 0
const Default_VIPSETTINGS_EliteSweep uint32 = 0
const Default_VIPSETTINGS_HardSweep uint32 = 0
const Default_VIPSETTINGS_DCLevelAdd float32 = 0
const Default_VIPSETTINGS_IronLevelAdd float32 = 0
const Default_VIPSETTINGS_GoldLevelAdd float32 = 0
const Default_VIPSETTINGS_ElitePurchase uint32 = 0
const Default_VIPSETTINGS_HardPurchase uint32 = 0
const Default_VIPSETTINGS_SweepResource uint32 = 0
const Default_VIPSETTINGS_ABossRatioUp uint32 = 0
const Default_VIPSETTINGS_BOSSFightSweep uint32 = 0
const Default_VIPSETTINGS_SBossRatioUp uint32 = 0
const Default_VIPSETTINGS_GuildSignLimit uint32 = 0
const Default_VIPSETTINGS_PVPPurchaseLimit int32 = -1
const Default_VIPSETTINGS_StarupUseHCLimit uint32 = 0
const Default_VIPSETTINGS_VIPGachaLimit uint32 = 0
const Default_VIPSETTINGS_VIPGachaVisible uint32 = 0
const Default_VIPSETTINGS_GuildWorshipLimit uint32 = 2
const Default_VIPSETTINGS_TeamBossRedBox uint32 = 0
const Default_VIPSETTINGS_Count1 uint32 = 0
const Default_VIPSETTINGS_Count2 uint32 = 0
const Default_VIPSETTINGS_Count3 uint32 = 0
const Default_VIPSETTINGS_Count4 uint32 = 0
const Default_VIPSETTINGS_GrowFund uint32 = 0
const Default_VIPSETTINGS_TenSweep uint32 = 0
const Default_VIPSETTINGS_BossFightSweep uint32 = 0
const Default_VIPSETTINGS_GoldLevelSweep uint32 = 0
const Default_VIPSETTINGS_DestinyLevelSweep uint32 = 0
const Default_VIPSETTINGS_IronLevelSweep uint32 = 0
const Default_VIPSETTINGS_FengHuoNormalSweep uint32 = 0
const Default_VIPSETTINGS_FengHuoSeniorSweep uint32 = 0
const Default_VIPSETTINGS_GEBuffEncourage uint32 = 0
const Default_VIPSETTINGS_DGVIPTrain uint32 = 0
const Default_VIPSETTINGS_DGVIPContinuityTrain uint32 = 0
const Default_VIPSETTINGS_ExpeditionTime uint32 = 0

func (m *VIPSETTINGS) GetVIPLevel() uint32 {
	if m != nil && m.VIPLevel != nil {
		return *m.VIPLevel
	}
	return Default_VIPSETTINGS_VIPLevel
}

func (m *VIPSETTINGS) GetRMBpoints() uint32 {
	if m != nil && m.RMBpoints != nil {
		return *m.RMBpoints
	}
	return Default_VIPSETTINGS_RMBpoints
}

func (m *VIPSETTINGS) GetVIPiconEnum() uint32 {
	if m != nil && m.VIPiconEnum != nil {
		return *m.VIPiconEnum
	}
	return Default_VIPSETTINGS_VIPiconEnum
}

func (m *VIPSETTINGS) GetEnergyPurchaseLimit() uint32 {
	if m != nil && m.EnergyPurchaseLimit != nil {
		return *m.EnergyPurchaseLimit
	}
	return Default_VIPSETTINGS_EnergyPurchaseLimit
}

func (m *VIPSETTINGS) GetHighEnergyLimitt() uint32 {
	if m != nil && m.HighEnergyLimitt != nil {
		return *m.HighEnergyLimitt
	}
	return Default_VIPSETTINGS_HighEnergyLimitt
}

func (m *VIPSETTINGS) GetEatBaoziQuik() uint32 {
	if m != nil && m.EatBaoziQuik != nil {
		return *m.EatBaoziQuik
	}
	return Default_VIPSETTINGS_EatBaoziQuik
}

func (m *VIPSETTINGS) GetSCPurchaseLimit() uint32 {
	if m != nil && m.SCPurchaseLimit != nil {
		return *m.SCPurchaseLimit
	}
	return Default_VIPSETTINGS_SCPurchaseLimit
}

func (m *VIPSETTINGS) GetSprintPurchaseLimit() int32 {
	if m != nil && m.SprintPurchaseLimit != nil {
		return *m.SprintPurchaseLimit
	}
	return Default_VIPSETTINGS_SprintPurchaseLimit
}

func (m *VIPSETTINGS) GetPVPTimeLimit() int32 {
	if m != nil && m.PVPTimeLimit != nil {
		return *m.PVPTimeLimit
	}
	return Default_VIPSETTINGS_PVPTimeLimit
}

func (m *VIPSETTINGS) GetTPVPTimeLimit() int32 {
	if m != nil && m.TPVPTimeLimit != nil {
		return *m.TPVPTimeLimit
	}
	return Default_VIPSETTINGS_TPVPTimeLimit
}

func (m *VIPSETTINGS) GetWSPVPTimeLimit() int32 {
	if m != nil && m.WSPVPTimeLimit != nil {
		return *m.WSPVPTimeLimit
	}
	return Default_VIPSETTINGS_WSPVPTimeLimit
}

func (m *VIPSETTINGS) GetStoreRefreshLimit() uint32 {
	if m != nil && m.StoreRefreshLimit != nil {
		return *m.StoreRefreshLimit
	}
	return Default_VIPSETTINGS_StoreRefreshLimit
}

func (m *VIPSETTINGS) GetPvpStoreRefreshLimit() uint32 {
	if m != nil && m.PvpStoreRefreshLimit != nil {
		return *m.PvpStoreRefreshLimit
	}
	return Default_VIPSETTINGS_PvpStoreRefreshLimit
}

func (m *VIPSETTINGS) GetTpvpStoreRefreshLimit() uint32 {
	if m != nil && m.TpvpStoreRefreshLimit != nil {
		return *m.TpvpStoreRefreshLimit
	}
	return Default_VIPSETTINGS_TpvpStoreRefreshLimit
}

func (m *VIPSETTINGS) GetWSPVPStoreRefreshLimit() uint32 {
	if m != nil && m.WSPVPStoreRefreshLimit != nil {
		return *m.WSPVPStoreRefreshLimit
	}
	return Default_VIPSETTINGS_WSPVPStoreRefreshLimit
}

func (m *VIPSETTINGS) GetBossStoreRefreshLimit() uint32 {
	if m != nil && m.BossStoreRefreshLimit != nil {
		return *m.BossStoreRefreshLimit
	}
	return Default_VIPSETTINGS_BossStoreRefreshLimit
}

func (m *VIPSETTINGS) GetGuildStoreRefreshLimit() uint32 {
	if m != nil && m.GuildStoreRefreshLimit != nil {
		return *m.GuildStoreRefreshLimit
	}
	return Default_VIPSETTINGS_GuildStoreRefreshLimit
}

func (m *VIPSETTINGS) GetExpeditionStoreRefreshLimit() uint32 {
	if m != nil && m.ExpeditionStoreRefreshLimit != nil {
		return *m.ExpeditionStoreRefreshLimit
	}
	return Default_VIPSETTINGS_ExpeditionStoreRefreshLimit
}

func (m *VIPSETTINGS) GetHDPStoreRefreshLimit() uint32 {
	if m != nil && m.HDPStoreRefreshLimit != nil {
		return *m.HDPStoreRefreshLimit
	}
	return Default_VIPSETTINGS_HDPStoreRefreshLimit
}

func (m *VIPSETTINGS) GetPetStoreRefreshLimit() uint32 {
	if m != nil && m.PetStoreRefreshLimit != nil {
		return *m.PetStoreRefreshLimit
	}
	return Default_VIPSETTINGS_PetStoreRefreshLimit
}

func (m *VIPSETTINGS) GetHeroInvestLimit() uint32 {
	if m != nil && m.HeroInvestLimit != nil {
		return *m.HeroInvestLimit
	}
	return Default_VIPSETTINGS_HeroInvestLimit
}

func (m *VIPSETTINGS) GetCropstimes() uint32 {
	if m != nil && m.Cropstimes != nil {
		return *m.Cropstimes
	}
	return Default_VIPSETTINGS_Cropstimes
}

func (m *VIPSETTINGS) GetCropsDuration() uint32 {
	if m != nil && m.CropsDuration != nil {
		return *m.CropsDuration
	}
	return Default_VIPSETTINGS_CropsDuration
}

func (m *VIPSETTINGS) GetRobtimes() uint32 {
	if m != nil && m.Robtimes != nil {
		return *m.Robtimes
	}
	return Default_VIPSETTINGS_Robtimes
}

func (m *VIPSETTINGS) GetAssisttimes() uint32 {
	if m != nil && m.Assisttimes != nil {
		return *m.Assisttimes
	}
	return Default_VIPSETTINGS_Assisttimes
}

func (m *VIPSETTINGS) GetWorldBosstimes() uint32 {
	if m != nil && m.WorldBosstimes != nil {
		return *m.WorldBosstimes
	}
	return Default_VIPSETTINGS_WorldBosstimes
}

func (m *VIPSETTINGS) GetSurplusGacha3() uint32 {
	if m != nil && m.SurplusGacha3 != nil {
		return *m.SurplusGacha3
	}
	return Default_VIPSETTINGS_SurplusGacha3
}

func (m *VIPSETTINGS) GetSurplusGacha4() uint32 {
	if m != nil && m.SurplusGacha4 != nil {
		return *m.SurplusGacha4
	}
	return Default_VIPSETTINGS_SurplusGacha4
}

func (m *VIPSETTINGS) GetSurplusGacha5() uint32 {
	if m != nil && m.SurplusGacha5 != nil {
		return *m.SurplusGacha5
	}
	return Default_VIPSETTINGS_SurplusGacha5
}

func (m *VIPSETTINGS) GetVIPDescripition() string {
	if m != nil && m.VIPDescripition != nil {
		return *m.VIPDescripition
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegeID() int32 {
	if m != nil && m.PrivilegeID != nil {
		return *m.PrivilegeID
	}
	return Default_VIPSETTINGS_PrivilegeID
}

func (m *VIPSETTINGS) GetSweep() uint32 {
	if m != nil && m.Sweep != nil {
		return *m.Sweep
	}
	return Default_VIPSETTINGS_Sweep
}

func (m *VIPSETTINGS) GetNormalSweep() uint32 {
	if m != nil && m.NormalSweep != nil {
		return *m.NormalSweep
	}
	return Default_VIPSETTINGS_NormalSweep
}

func (m *VIPSETTINGS) GetBossRatioUp() uint32 {
	if m != nil && m.BossRatioUp != nil {
		return *m.BossRatioUp
	}
	return Default_VIPSETTINGS_BossRatioUp
}

func (m *VIPSETTINGS) GetEliteSweep() uint32 {
	if m != nil && m.EliteSweep != nil {
		return *m.EliteSweep
	}
	return Default_VIPSETTINGS_EliteSweep
}

func (m *VIPSETTINGS) GetHardSweep() uint32 {
	if m != nil && m.HardSweep != nil {
		return *m.HardSweep
	}
	return Default_VIPSETTINGS_HardSweep
}

func (m *VIPSETTINGS) GetDCLevelAdd() float32 {
	if m != nil && m.DCLevelAdd != nil {
		return *m.DCLevelAdd
	}
	return Default_VIPSETTINGS_DCLevelAdd
}

func (m *VIPSETTINGS) GetIronLevelAdd() float32 {
	if m != nil && m.IronLevelAdd != nil {
		return *m.IronLevelAdd
	}
	return Default_VIPSETTINGS_IronLevelAdd
}

func (m *VIPSETTINGS) GetGoldLevelAdd() float32 {
	if m != nil && m.GoldLevelAdd != nil {
		return *m.GoldLevelAdd
	}
	return Default_VIPSETTINGS_GoldLevelAdd
}

func (m *VIPSETTINGS) GetElitePurchase() uint32 {
	if m != nil && m.ElitePurchase != nil {
		return *m.ElitePurchase
	}
	return Default_VIPSETTINGS_ElitePurchase
}

func (m *VIPSETTINGS) GetHardPurchase() uint32 {
	if m != nil && m.HardPurchase != nil {
		return *m.HardPurchase
	}
	return Default_VIPSETTINGS_HardPurchase
}

func (m *VIPSETTINGS) GetSweepResource() uint32 {
	if m != nil && m.SweepResource != nil {
		return *m.SweepResource
	}
	return Default_VIPSETTINGS_SweepResource
}

func (m *VIPSETTINGS) GetABossRatioUp() uint32 {
	if m != nil && m.ABossRatioUp != nil {
		return *m.ABossRatioUp
	}
	return Default_VIPSETTINGS_ABossRatioUp
}

func (m *VIPSETTINGS) GetBOSSFightSweep() uint32 {
	if m != nil && m.BOSSFightSweep != nil {
		return *m.BOSSFightSweep
	}
	return Default_VIPSETTINGS_BOSSFightSweep
}

func (m *VIPSETTINGS) GetSBossRatioUp() uint32 {
	if m != nil && m.SBossRatioUp != nil {
		return *m.SBossRatioUp
	}
	return Default_VIPSETTINGS_SBossRatioUp
}

func (m *VIPSETTINGS) GetGuildSignLimit() uint32 {
	if m != nil && m.GuildSignLimit != nil {
		return *m.GuildSignLimit
	}
	return Default_VIPSETTINGS_GuildSignLimit
}

func (m *VIPSETTINGS) GetPVPPurchaseLimit() int32 {
	if m != nil && m.PVPPurchaseLimit != nil {
		return *m.PVPPurchaseLimit
	}
	return Default_VIPSETTINGS_PVPPurchaseLimit
}

func (m *VIPSETTINGS) GetStarupUseHCLimit() uint32 {
	if m != nil && m.StarupUseHCLimit != nil {
		return *m.StarupUseHCLimit
	}
	return Default_VIPSETTINGS_StarupUseHCLimit
}

func (m *VIPSETTINGS) GetVIPGachaLimit() uint32 {
	if m != nil && m.VIPGachaLimit != nil {
		return *m.VIPGachaLimit
	}
	return Default_VIPSETTINGS_VIPGachaLimit
}

func (m *VIPSETTINGS) GetVIPGachaVisible() uint32 {
	if m != nil && m.VIPGachaVisible != nil {
		return *m.VIPGachaVisible
	}
	return Default_VIPSETTINGS_VIPGachaVisible
}

func (m *VIPSETTINGS) GetGuildWorshipLimit() uint32 {
	if m != nil && m.GuildWorshipLimit != nil {
		return *m.GuildWorshipLimit
	}
	return Default_VIPSETTINGS_GuildWorshipLimit
}

func (m *VIPSETTINGS) GetTeamBossRedBox() uint32 {
	if m != nil && m.TeamBossRedBox != nil {
		return *m.TeamBossRedBox
	}
	return Default_VIPSETTINGS_TeamBossRedBox
}

func (m *VIPSETTINGS) GetPrivilegeImg1() string {
	if m != nil && m.PrivilegeImg1 != nil {
		return *m.PrivilegeImg1
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegemgUpDes1() string {
	if m != nil && m.PrivilegemgUpDes1 != nil {
		return *m.PrivilegemgUpDes1
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegemgDownDes1() string {
	if m != nil && m.PrivilegemgDownDes1 != nil {
		return *m.PrivilegemgDownDes1
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegeImg2() string {
	if m != nil && m.PrivilegeImg2 != nil {
		return *m.PrivilegeImg2
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegemgUpDes2() string {
	if m != nil && m.PrivilegemgUpDes2 != nil {
		return *m.PrivilegemgUpDes2
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegemgDownDes2() string {
	if m != nil && m.PrivilegemgDownDes2 != nil {
		return *m.PrivilegemgDownDes2
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegeImg3() string {
	if m != nil && m.PrivilegeImg3 != nil {
		return *m.PrivilegeImg3
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegemgUpDes3() string {
	if m != nil && m.PrivilegemgUpDes3 != nil {
		return *m.PrivilegemgUpDes3
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegemgDownDes3() string {
	if m != nil && m.PrivilegemgDownDes3 != nil {
		return *m.PrivilegemgDownDes3
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegeImg4() string {
	if m != nil && m.PrivilegeImg4 != nil {
		return *m.PrivilegeImg4
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegemgUpDes4() string {
	if m != nil && m.PrivilegemgUpDes4 != nil {
		return *m.PrivilegemgUpDes4
	}
	return ""
}

func (m *VIPSETTINGS) GetPrivilegemgDownDes4() string {
	if m != nil && m.PrivilegemgDownDes4 != nil {
		return *m.PrivilegemgDownDes4
	}
	return ""
}

func (m *VIPSETTINGS) GetVIPDailyGift1() string {
	if m != nil && m.VIPDailyGift1 != nil {
		return *m.VIPDailyGift1
	}
	return ""
}

func (m *VIPSETTINGS) GetCount1() uint32 {
	if m != nil && m.Count1 != nil {
		return *m.Count1
	}
	return Default_VIPSETTINGS_Count1
}

func (m *VIPSETTINGS) GetVIPDailyGift2() string {
	if m != nil && m.VIPDailyGift2 != nil {
		return *m.VIPDailyGift2
	}
	return ""
}

func (m *VIPSETTINGS) GetCount2() uint32 {
	if m != nil && m.Count2 != nil {
		return *m.Count2
	}
	return Default_VIPSETTINGS_Count2
}

func (m *VIPSETTINGS) GetVIPDailyGift3() string {
	if m != nil && m.VIPDailyGift3 != nil {
		return *m.VIPDailyGift3
	}
	return ""
}

func (m *VIPSETTINGS) GetCount3() uint32 {
	if m != nil && m.Count3 != nil {
		return *m.Count3
	}
	return Default_VIPSETTINGS_Count3
}

func (m *VIPSETTINGS) GetVIPDailyGift4() string {
	if m != nil && m.VIPDailyGift4 != nil {
		return *m.VIPDailyGift4
	}
	return ""
}

func (m *VIPSETTINGS) GetCount4() uint32 {
	if m != nil && m.Count4 != nil {
		return *m.Count4
	}
	return Default_VIPSETTINGS_Count4
}

func (m *VIPSETTINGS) GetVIPPoster() string {
	if m != nil && m.VIPPoster != nil {
		return *m.VIPPoster
	}
	return ""
}

func (m *VIPSETTINGS) GetGrowFund() uint32 {
	if m != nil && m.GrowFund != nil {
		return *m.GrowFund
	}
	return Default_VIPSETTINGS_GrowFund
}

func (m *VIPSETTINGS) GetTenSweep() uint32 {
	if m != nil && m.TenSweep != nil {
		return *m.TenSweep
	}
	return Default_VIPSETTINGS_TenSweep
}

func (m *VIPSETTINGS) GetBossFightSweep() uint32 {
	if m != nil && m.BossFightSweep != nil {
		return *m.BossFightSweep
	}
	return Default_VIPSETTINGS_BossFightSweep
}

func (m *VIPSETTINGS) GetGoldLevelSweep() uint32 {
	if m != nil && m.GoldLevelSweep != nil {
		return *m.GoldLevelSweep
	}
	return Default_VIPSETTINGS_GoldLevelSweep
}

func (m *VIPSETTINGS) GetDestinyLevelSweep() uint32 {
	if m != nil && m.DestinyLevelSweep != nil {
		return *m.DestinyLevelSweep
	}
	return Default_VIPSETTINGS_DestinyLevelSweep
}

func (m *VIPSETTINGS) GetIronLevelSweep() uint32 {
	if m != nil && m.IronLevelSweep != nil {
		return *m.IronLevelSweep
	}
	return Default_VIPSETTINGS_IronLevelSweep
}

func (m *VIPSETTINGS) GetFengHuoNormalSweep() uint32 {
	if m != nil && m.FengHuoNormalSweep != nil {
		return *m.FengHuoNormalSweep
	}
	return Default_VIPSETTINGS_FengHuoNormalSweep
}

func (m *VIPSETTINGS) GetFengHuoSeniorSweep() uint32 {
	if m != nil && m.FengHuoSeniorSweep != nil {
		return *m.FengHuoSeniorSweep
	}
	return Default_VIPSETTINGS_FengHuoSeniorSweep
}

func (m *VIPSETTINGS) GetGEBuffEncourage() uint32 {
	if m != nil && m.GEBuffEncourage != nil {
		return *m.GEBuffEncourage
	}
	return Default_VIPSETTINGS_GEBuffEncourage
}

func (m *VIPSETTINGS) GetDGVIPTrain() uint32 {
	if m != nil && m.DGVIPTrain != nil {
		return *m.DGVIPTrain
	}
	return Default_VIPSETTINGS_DGVIPTrain
}

func (m *VIPSETTINGS) GetDGVIPContinuityTrain() uint32 {
	if m != nil && m.DGVIPContinuityTrain != nil {
		return *m.DGVIPContinuityTrain
	}
	return Default_VIPSETTINGS_DGVIPContinuityTrain
}

func (m *VIPSETTINGS) GetExpeditionTime() uint32 {
	if m != nil && m.ExpeditionTime != nil {
		return *m.ExpeditionTime
	}
	return Default_VIPSETTINGS_ExpeditionTime
}

func (m *VIPSETTINGS) GetVIP_Table() []*VIPSETTINGS_VIP {
	if m != nil {
		return m.VIP_Table
	}
	return nil
}

func (m *VIPSETTINGS) GetStoreRefresh_Table() []*VIPSETTINGS_StoreRefresh {
	if m != nil {
		return m.StoreRefresh_Table
	}
	return nil
}

type VIPSETTINGS_VIP struct {
	// * VIP特权的描述
	VIPPrivilegeIDS  *string `protobuf:"bytes,1,opt,def=" json:"VIPPrivilegeIDS,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VIPSETTINGS_VIP) Reset()         { *m = VIPSETTINGS_VIP{} }
func (m *VIPSETTINGS_VIP) String() string { return proto.CompactTextString(m) }
func (*VIPSETTINGS_VIP) ProtoMessage()    {}

func (m *VIPSETTINGS_VIP) GetVIPPrivilegeIDS() string {
	if m != nil && m.VIPPrivilegeIDS != nil {
		return *m.VIPPrivilegeIDS
	}
	return ""
}

type VIPSETTINGS_StoreRefresh struct {
	// * 军团商店
	StoreID *uint32 `protobuf:"varint,1,opt,def=0" json:"StoreID,omitempty"`
	// * 刷新次数
	RefreshTime      *uint32 `protobuf:"varint,2,opt,def=0" json:"RefreshTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VIPSETTINGS_StoreRefresh) Reset()         { *m = VIPSETTINGS_StoreRefresh{} }
func (m *VIPSETTINGS_StoreRefresh) String() string { return proto.CompactTextString(m) }
func (*VIPSETTINGS_StoreRefresh) ProtoMessage()    {}

const Default_VIPSETTINGS_StoreRefresh_StoreID uint32 = 0
const Default_VIPSETTINGS_StoreRefresh_RefreshTime uint32 = 0

func (m *VIPSETTINGS_StoreRefresh) GetStoreID() uint32 {
	if m != nil && m.StoreID != nil {
		return *m.StoreID
	}
	return Default_VIPSETTINGS_StoreRefresh_StoreID
}

func (m *VIPSETTINGS_StoreRefresh) GetRefreshTime() uint32 {
	if m != nil && m.RefreshTime != nil {
		return *m.RefreshTime
	}
	return Default_VIPSETTINGS_StoreRefresh_RefreshTime
}

type VIPSETTINGS_ARRAY struct {
	Items            []*VIPSETTINGS `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *VIPSETTINGS_ARRAY) Reset()         { *m = VIPSETTINGS_ARRAY{} }
func (m *VIPSETTINGS_ARRAY) String() string { return proto.CompactTextString(m) }
func (*VIPSETTINGS_ARRAY) ProtoMessage()    {}

func (m *VIPSETTINGS_ARRAY) GetItems() []*VIPSETTINGS {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
