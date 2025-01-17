// Code generated by protoc-gen-go.
// source: ProtobufGen_gankids.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GANKIDS struct {
	// * 信息类型，0为胜利，1为失败
	Type *uint32 `protobuf:"varint,1,req,def=0" json:"Type,omitempty"`
	// * 切磋信息的ids
	InfoIds          *string `protobuf:"bytes,2,req,def=" json:"InfoIds,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GANKIDS) Reset()         { *m = GANKIDS{} }
func (m *GANKIDS) String() string { return proto.CompactTextString(m) }
func (*GANKIDS) ProtoMessage()    {}

const Default_GANKIDS_Type uint32 = 0

func (m *GANKIDS) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_GANKIDS_Type
}

func (m *GANKIDS) GetInfoIds() string {
	if m != nil && m.InfoIds != nil {
		return *m.InfoIds
	}
	return ""
}

type GANKIDS_ARRAY struct {
	Items            []*GANKIDS `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *GANKIDS_ARRAY) Reset()         { *m = GANKIDS_ARRAY{} }
func (m *GANKIDS_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GANKIDS_ARRAY) ProtoMessage()    {}

func (m *GANKIDS_ARRAY) GetItems() []*GANKIDS {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
