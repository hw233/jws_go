// Code generated by protoc-gen-go.
// source: ProtobufGen_subtitlegroup.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SUBTITLEGROUP struct {
	// * 字幕组ID
	SubtitleGroupID *string `protobuf:"bytes,1,req,def=" json:"SubtitleGroupID,omitempty"`
	// * 开始时间(秒）
	StartTime1 *float32 `protobuf:"fixed32,2,opt,def=0" json:"StartTime1,omitempty"`
	// * 语言2_开始时间(秒）
	Language1_StartTime1 *float32 `protobuf:"fixed32,8,opt,def=0" json:"Language1_StartTime1,omitempty"`
	// * 语言3_开始时间(秒）
	Language2_StartTime1 *float32 `protobuf:"fixed32,9,opt,def=0" json:"Language2_StartTime1,omitempty"`
	// * 对应的IDS
	IDS1 *string `protobuf:"bytes,3,opt,def=" json:"IDS1,omitempty"`
	// * 开始时间(秒）
	StartTime2 *float32 `protobuf:"fixed32,4,opt,def=0" json:"StartTime2,omitempty"`
	// * 语言2_开始时间(秒）
	Language1_StartTime2 *float32 `protobuf:"fixed32,10,opt,def=0" json:"Language1_StartTime2,omitempty"`
	// * 语言3_开始时间(秒）
	Language2_StartTime2 *float32 `protobuf:"fixed32,11,opt,def=0" json:"Language2_StartTime2,omitempty"`
	// * 对应的IDS
	IDS2 *string `protobuf:"bytes,5,opt,def=" json:"IDS2,omitempty"`
	// * 结束时间（秒）
	EndTime *float32 `protobuf:"fixed32,6,opt,def=0" json:"EndTime,omitempty"`
	// * 语言2_结束时间（秒）
	Language1_EndTime *float32 `protobuf:"fixed32,12,opt,def=0" json:"Language1_EndTime,omitempty"`
	// * 语言3_结束时间（秒）
	Language2_EndTime *float32 `protobuf:"fixed32,13,opt,def=0" json:"Language2_EndTime,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *SUBTITLEGROUP) Reset()         { *m = SUBTITLEGROUP{} }
func (m *SUBTITLEGROUP) String() string { return proto.CompactTextString(m) }
func (*SUBTITLEGROUP) ProtoMessage()    {}

const Default_SUBTITLEGROUP_StartTime1 float32 = 0
const Default_SUBTITLEGROUP_Language1_StartTime1 float32 = 0
const Default_SUBTITLEGROUP_Language2_StartTime1 float32 = 0
const Default_SUBTITLEGROUP_StartTime2 float32 = 0
const Default_SUBTITLEGROUP_Language1_StartTime2 float32 = 0
const Default_SUBTITLEGROUP_Language2_StartTime2 float32 = 0
const Default_SUBTITLEGROUP_EndTime float32 = 0
const Default_SUBTITLEGROUP_Language1_EndTime float32 = 0
const Default_SUBTITLEGROUP_Language2_EndTime float32 = 0

func (m *SUBTITLEGROUP) GetSubtitleGroupID() string {
	if m != nil && m.SubtitleGroupID != nil {
		return *m.SubtitleGroupID
	}
	return ""
}

func (m *SUBTITLEGROUP) GetStartTime1() float32 {
	if m != nil && m.StartTime1 != nil {
		return *m.StartTime1
	}
	return Default_SUBTITLEGROUP_StartTime1
}

func (m *SUBTITLEGROUP) GetLanguage1_StartTime1() float32 {
	if m != nil && m.Language1_StartTime1 != nil {
		return *m.Language1_StartTime1
	}
	return Default_SUBTITLEGROUP_Language1_StartTime1
}

func (m *SUBTITLEGROUP) GetLanguage2_StartTime1() float32 {
	if m != nil && m.Language2_StartTime1 != nil {
		return *m.Language2_StartTime1
	}
	return Default_SUBTITLEGROUP_Language2_StartTime1
}

func (m *SUBTITLEGROUP) GetIDS1() string {
	if m != nil && m.IDS1 != nil {
		return *m.IDS1
	}
	return ""
}

func (m *SUBTITLEGROUP) GetStartTime2() float32 {
	if m != nil && m.StartTime2 != nil {
		return *m.StartTime2
	}
	return Default_SUBTITLEGROUP_StartTime2
}

func (m *SUBTITLEGROUP) GetLanguage1_StartTime2() float32 {
	if m != nil && m.Language1_StartTime2 != nil {
		return *m.Language1_StartTime2
	}
	return Default_SUBTITLEGROUP_Language1_StartTime2
}

func (m *SUBTITLEGROUP) GetLanguage2_StartTime2() float32 {
	if m != nil && m.Language2_StartTime2 != nil {
		return *m.Language2_StartTime2
	}
	return Default_SUBTITLEGROUP_Language2_StartTime2
}

func (m *SUBTITLEGROUP) GetIDS2() string {
	if m != nil && m.IDS2 != nil {
		return *m.IDS2
	}
	return ""
}

func (m *SUBTITLEGROUP) GetEndTime() float32 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return Default_SUBTITLEGROUP_EndTime
}

func (m *SUBTITLEGROUP) GetLanguage1_EndTime() float32 {
	if m != nil && m.Language1_EndTime != nil {
		return *m.Language1_EndTime
	}
	return Default_SUBTITLEGROUP_Language1_EndTime
}

func (m *SUBTITLEGROUP) GetLanguage2_EndTime() float32 {
	if m != nil && m.Language2_EndTime != nil {
		return *m.Language2_EndTime
	}
	return Default_SUBTITLEGROUP_Language2_EndTime
}

type SUBTITLEGROUP_ARRAY struct {
	Items            []*SUBTITLEGROUP `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *SUBTITLEGROUP_ARRAY) Reset()         { *m = SUBTITLEGROUP_ARRAY{} }
func (m *SUBTITLEGROUP_ARRAY) String() string { return proto.CompactTextString(m) }
func (*SUBTITLEGROUP_ARRAY) ProtoMessage()    {}

func (m *SUBTITLEGROUP_ARRAY) GetItems() []*SUBTITLEGROUP {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
