// Code generated by protoc-gen-go.
// source: ProtobufGen_hotgachaspecial.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type HOTGACHASPECIAL struct {
	// * 活动ID
	ActivityID *uint32 `protobuf:"varint,1,opt,def=0" json:"ActivityID,omitempty"`
	// * 活动子ID
	ActivitySubID *uint32 `protobuf:"varint,5,opt,def=0" json:"ActivitySubID,omitempty"`
	// * 暗控次数
	SpecialTimes *uint32 `protobuf:"varint,2,opt,def=0" json:"SpecialTimes,omitempty"`
	// * 抽取ID
	GachaID          *uint32 `protobuf:"varint,3,opt,def=0" json:"GachaID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HOTGACHASPECIAL) Reset()         { *m = HOTGACHASPECIAL{} }
func (m *HOTGACHASPECIAL) String() string { return proto.CompactTextString(m) }
func (*HOTGACHASPECIAL) ProtoMessage()    {}

const Default_HOTGACHASPECIAL_ActivityID uint32 = 0
const Default_HOTGACHASPECIAL_ActivitySubID uint32 = 0
const Default_HOTGACHASPECIAL_SpecialTimes uint32 = 0
const Default_HOTGACHASPECIAL_GachaID uint32 = 0

func (m *HOTGACHASPECIAL) GetActivityID() uint32 {
	if m != nil && m.ActivityID != nil {
		return *m.ActivityID
	}
	return Default_HOTGACHASPECIAL_ActivityID
}

func (m *HOTGACHASPECIAL) GetActivitySubID() uint32 {
	if m != nil && m.ActivitySubID != nil {
		return *m.ActivitySubID
	}
	return Default_HOTGACHASPECIAL_ActivitySubID
}

func (m *HOTGACHASPECIAL) GetSpecialTimes() uint32 {
	if m != nil && m.SpecialTimes != nil {
		return *m.SpecialTimes
	}
	return Default_HOTGACHASPECIAL_SpecialTimes
}

func (m *HOTGACHASPECIAL) GetGachaID() uint32 {
	if m != nil && m.GachaID != nil {
		return *m.GachaID
	}
	return Default_HOTGACHASPECIAL_GachaID
}

type HOTGACHASPECIAL_ARRAY struct {
	Items            []*HOTGACHASPECIAL `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HOTGACHASPECIAL_ARRAY) Reset()         { *m = HOTGACHASPECIAL_ARRAY{} }
func (m *HOTGACHASPECIAL_ARRAY) String() string { return proto.CompactTextString(m) }
func (*HOTGACHASPECIAL_ARRAY) ProtoMessage()    {}

func (m *HOTGACHASPECIAL_ARRAY) GetItems() []*HOTGACHASPECIAL {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
