// Code generated by protoc-gen-go.
// source: ProtobufGen_conditionpage.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type CONDITIONPAGE struct {
	// * 引导掉落的界面名
	ConditionPage *string `protobuf:"bytes,1,req,def=" json:"ConditionPage,omitempty"`
	// * 开放的功能名称ID
	ConditionID *uint32 `protobuf:"varint,2,req,def=0" json:"ConditionID,omitempty"`
	// * 引导掉落途径对应的文字
	ConditionIDS     *string `protobuf:"bytes,3,req,def=" json:"ConditionIDS,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CONDITIONPAGE) Reset()         { *m = CONDITIONPAGE{} }
func (m *CONDITIONPAGE) String() string { return proto.CompactTextString(m) }
func (*CONDITIONPAGE) ProtoMessage()    {}

const Default_CONDITIONPAGE_ConditionID uint32 = 0

func (m *CONDITIONPAGE) GetConditionPage() string {
	if m != nil && m.ConditionPage != nil {
		return *m.ConditionPage
	}
	return ""
}

func (m *CONDITIONPAGE) GetConditionID() uint32 {
	if m != nil && m.ConditionID != nil {
		return *m.ConditionID
	}
	return Default_CONDITIONPAGE_ConditionID
}

func (m *CONDITIONPAGE) GetConditionIDS() string {
	if m != nil && m.ConditionIDS != nil {
		return *m.ConditionIDS
	}
	return ""
}

type CONDITIONPAGE_ARRAY struct {
	Items            []*CONDITIONPAGE `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CONDITIONPAGE_ARRAY) Reset()         { *m = CONDITIONPAGE_ARRAY{} }
func (m *CONDITIONPAGE_ARRAY) String() string { return proto.CompactTextString(m) }
func (*CONDITIONPAGE_ARRAY) ProtoMessage()    {}

func (m *CONDITIONPAGE_ARRAY) GetItems() []*CONDITIONPAGE {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}