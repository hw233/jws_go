// Code generated by protoc-gen-go.
// source: ProtobufGen_resetcost.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type RESETCOST struct {
	// * 重置次数
	ResetTimes *uint32 `protobuf:"varint,1,req,def=0" json:"ResetTimes,omitempty"`
	// * 重置费用
	ResetCost        *uint32 `protobuf:"varint,2,req,def=0" json:"ResetCost,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RESETCOST) Reset()         { *m = RESETCOST{} }
func (m *RESETCOST) String() string { return proto.CompactTextString(m) }
func (*RESETCOST) ProtoMessage()    {}

const Default_RESETCOST_ResetTimes uint32 = 0
const Default_RESETCOST_ResetCost uint32 = 0

func (m *RESETCOST) GetResetTimes() uint32 {
	if m != nil && m.ResetTimes != nil {
		return *m.ResetTimes
	}
	return Default_RESETCOST_ResetTimes
}

func (m *RESETCOST) GetResetCost() uint32 {
	if m != nil && m.ResetCost != nil {
		return *m.ResetCost
	}
	return Default_RESETCOST_ResetCost
}

type RESETCOST_ARRAY struct {
	Items            []*RESETCOST `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *RESETCOST_ARRAY) Reset()         { *m = RESETCOST_ARRAY{} }
func (m *RESETCOST_ARRAY) String() string { return proto.CompactTextString(m) }
func (*RESETCOST_ARRAY) ProtoMessage()    {}

func (m *RESETCOST_ARRAY) GetItems() []*RESETCOST {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}