// Code generated by protoc-gen-go.
// source: ProtobufGen_tpvptime.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type TPVPTIME struct {
	// * 挑战次数购买数量
	TPVPTime *uint32 `protobuf:"varint,1,req,def=0" json:"TPVPTime,omitempty"`
	// * HC消耗
	TPVPPrice        *uint32 `protobuf:"varint,2,req,def=0" json:"TPVPPrice,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TPVPTIME) Reset()         { *m = TPVPTIME{} }
func (m *TPVPTIME) String() string { return proto.CompactTextString(m) }
func (*TPVPTIME) ProtoMessage()    {}

const Default_TPVPTIME_TPVPTime uint32 = 0
const Default_TPVPTIME_TPVPPrice uint32 = 0

func (m *TPVPTIME) GetTPVPTime() uint32 {
	if m != nil && m.TPVPTime != nil {
		return *m.TPVPTime
	}
	return Default_TPVPTIME_TPVPTime
}

func (m *TPVPTIME) GetTPVPPrice() uint32 {
	if m != nil && m.TPVPPrice != nil {
		return *m.TPVPPrice
	}
	return Default_TPVPTIME_TPVPPrice
}

type TPVPTIME_ARRAY struct {
	Items            []*TPVPTIME `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TPVPTIME_ARRAY) Reset()         { *m = TPVPTIME_ARRAY{} }
func (m *TPVPTIME_ARRAY) String() string { return proto.CompactTextString(m) }
func (*TPVPTIME_ARRAY) ProtoMessage()    {}

func (m *TPVPTIME_ARRAY) GetItems() []*TPVPTIME {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
