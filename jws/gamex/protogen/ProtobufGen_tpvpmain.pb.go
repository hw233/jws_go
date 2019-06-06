// Code generated by protoc-gen-go.
// source: ProtobufGen_tpvpmain.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type TPVPMAIN struct {
	// * 最大战力
	YourGSMax *uint32 `protobuf:"varint,1,req,def=0" json:"YourGSMax,omitempty"`
	// * 最小战力
	YourGSMin *uint32 `protobuf:"varint,2,opt,def=0" json:"YourGSMin,omitempty"`
	// * 最大判断区间
	HopeLimitMax *float32 `protobuf:"fixed32,3,opt,def=0" json:"HopeLimitMax,omitempty"`
	// * 最小判断区间
	HopeLimitMin *float32 `protobuf:"fixed32,4,opt,def=0" json:"HopeLimitMin,omitempty"`
	// * 持续时间
	TeamPvPTime *int32 `protobuf:"varint,5,opt,def=0" json:"TeamPvPTime,omitempty"`
	// * 检测间隔
	TeamPvPTimeUpdate *int32 `protobuf:"varint,6,opt,def=0" json:"TeamPvPTimeUpdate,omitempty"`
	// * 自动刷新间隔
	TeamPvPRefreshTime *int32 `protobuf:"varint,7,opt,def=0" json:"TeamPvPRefreshTime,omitempty"`
	// * 服务器忍耐时间秒
	ServerServeTime  *uint32 `protobuf:"varint,8,opt,def=0" json:"ServerServeTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TPVPMAIN) Reset()         { *m = TPVPMAIN{} }
func (m *TPVPMAIN) String() string { return proto.CompactTextString(m) }
func (*TPVPMAIN) ProtoMessage()    {}

const Default_TPVPMAIN_YourGSMax uint32 = 0
const Default_TPVPMAIN_YourGSMin uint32 = 0
const Default_TPVPMAIN_HopeLimitMax float32 = 0
const Default_TPVPMAIN_HopeLimitMin float32 = 0
const Default_TPVPMAIN_TeamPvPTime int32 = 0
const Default_TPVPMAIN_TeamPvPTimeUpdate int32 = 0
const Default_TPVPMAIN_TeamPvPRefreshTime int32 = 0
const Default_TPVPMAIN_ServerServeTime uint32 = 0

func (m *TPVPMAIN) GetYourGSMax() uint32 {
	if m != nil && m.YourGSMax != nil {
		return *m.YourGSMax
	}
	return Default_TPVPMAIN_YourGSMax
}

func (m *TPVPMAIN) GetYourGSMin() uint32 {
	if m != nil && m.YourGSMin != nil {
		return *m.YourGSMin
	}
	return Default_TPVPMAIN_YourGSMin
}

func (m *TPVPMAIN) GetHopeLimitMax() float32 {
	if m != nil && m.HopeLimitMax != nil {
		return *m.HopeLimitMax
	}
	return Default_TPVPMAIN_HopeLimitMax
}

func (m *TPVPMAIN) GetHopeLimitMin() float32 {
	if m != nil && m.HopeLimitMin != nil {
		return *m.HopeLimitMin
	}
	return Default_TPVPMAIN_HopeLimitMin
}

func (m *TPVPMAIN) GetTeamPvPTime() int32 {
	if m != nil && m.TeamPvPTime != nil {
		return *m.TeamPvPTime
	}
	return Default_TPVPMAIN_TeamPvPTime
}

func (m *TPVPMAIN) GetTeamPvPTimeUpdate() int32 {
	if m != nil && m.TeamPvPTimeUpdate != nil {
		return *m.TeamPvPTimeUpdate
	}
	return Default_TPVPMAIN_TeamPvPTimeUpdate
}

func (m *TPVPMAIN) GetTeamPvPRefreshTime() int32 {
	if m != nil && m.TeamPvPRefreshTime != nil {
		return *m.TeamPvPRefreshTime
	}
	return Default_TPVPMAIN_TeamPvPRefreshTime
}

func (m *TPVPMAIN) GetServerServeTime() uint32 {
	if m != nil && m.ServerServeTime != nil {
		return *m.ServerServeTime
	}
	return Default_TPVPMAIN_ServerServeTime
}

type TPVPMAIN_ARRAY struct {
	Items            []*TPVPMAIN `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TPVPMAIN_ARRAY) Reset()         { *m = TPVPMAIN_ARRAY{} }
func (m *TPVPMAIN_ARRAY) String() string { return proto.CompactTextString(m) }
func (*TPVPMAIN_ARRAY) ProtoMessage()    {}

func (m *TPVPMAIN_ARRAY) GetItems() []*TPVPMAIN {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}