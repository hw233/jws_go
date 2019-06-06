// Code generated by protoc-gen-go.
// source: ProtobufGen_gvgwinspoint.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GVGWINSPOINT struct {
	// * 连胜场数
	WinNum *uint32 `protobuf:"varint,1,req,def=0" json:"WinNum,omitempty"`
	// * 额外增加积分
	GVGPoint         *uint32 `protobuf:"varint,2,opt,def=0" json:"GVGPoint,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GVGWINSPOINT) Reset()         { *m = GVGWINSPOINT{} }
func (m *GVGWINSPOINT) String() string { return proto.CompactTextString(m) }
func (*GVGWINSPOINT) ProtoMessage()    {}

const Default_GVGWINSPOINT_WinNum uint32 = 0
const Default_GVGWINSPOINT_GVGPoint uint32 = 0

func (m *GVGWINSPOINT) GetWinNum() uint32 {
	if m != nil && m.WinNum != nil {
		return *m.WinNum
	}
	return Default_GVGWINSPOINT_WinNum
}

func (m *GVGWINSPOINT) GetGVGPoint() uint32 {
	if m != nil && m.GVGPoint != nil {
		return *m.GVGPoint
	}
	return Default_GVGWINSPOINT_GVGPoint
}

type GVGWINSPOINT_ARRAY struct {
	Items            []*GVGWINSPOINT `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GVGWINSPOINT_ARRAY) Reset()         { *m = GVGWINSPOINT_ARRAY{} }
func (m *GVGWINSPOINT_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GVGWINSPOINT_ARRAY) ProtoMessage()    {}

func (m *GVGWINSPOINT_ARRAY) GetItems() []*GVGWINSPOINT {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}