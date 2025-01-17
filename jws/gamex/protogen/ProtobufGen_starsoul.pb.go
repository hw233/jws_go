// Code generated by protoc-gen-go.
// source: ProtobufGen_starsoul.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type STARSOUL struct {
	// * 星魂品质
	RareLevel          *uint32               `protobuf:"varint,1,req,def=0" json:"RareLevel,omitempty"`
	SSResolve_Template []*STARSOUL_SSResolve `protobuf:"bytes,3,rep" json:"SSResolve_Template,omitempty"`
	XXX_unrecognized   []byte                `json:"-"`
}

func (m *STARSOUL) Reset()         { *m = STARSOUL{} }
func (m *STARSOUL) String() string { return proto.CompactTextString(m) }
func (*STARSOUL) ProtoMessage()    {}

const Default_STARSOUL_RareLevel uint32 = 0

func (m *STARSOUL) GetRareLevel() uint32 {
	if m != nil && m.RareLevel != nil {
		return *m.RareLevel
	}
	return Default_STARSOUL_RareLevel
}

func (m *STARSOUL) GetSSResolve_Template() []*STARSOUL_SSResolve {
	if m != nil {
		return m.SSResolve_Template
	}
	return nil
}

type STARSOUL_SSResolve struct {
	// * 星魂分解材料
	SSRResolveMaterial *string `protobuf:"bytes,1,opt,def=" json:"SSRResolveMaterial,omitempty"`
	// * 星魂分解材料数量
	SSRResolveMaterialCount *uint32 `protobuf:"varint,2,opt,def=1" json:"SSRResolveMaterialCount,omitempty"`
	XXX_unrecognized        []byte  `json:"-"`
}

func (m *STARSOUL_SSResolve) Reset()         { *m = STARSOUL_SSResolve{} }
func (m *STARSOUL_SSResolve) String() string { return proto.CompactTextString(m) }
func (*STARSOUL_SSResolve) ProtoMessage()    {}

const Default_STARSOUL_SSResolve_SSRResolveMaterialCount uint32 = 1

func (m *STARSOUL_SSResolve) GetSSRResolveMaterial() string {
	if m != nil && m.SSRResolveMaterial != nil {
		return *m.SSRResolveMaterial
	}
	return ""
}

func (m *STARSOUL_SSResolve) GetSSRResolveMaterialCount() uint32 {
	if m != nil && m.SSRResolveMaterialCount != nil {
		return *m.SSRResolveMaterialCount
	}
	return Default_STARSOUL_SSResolve_SSRResolveMaterialCount
}

type STARSOUL_ARRAY struct {
	Items            []*STARSOUL `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *STARSOUL_ARRAY) Reset()         { *m = STARSOUL_ARRAY{} }
func (m *STARSOUL_ARRAY) String() string { return proto.CompactTextString(m) }
func (*STARSOUL_ARRAY) ProtoMessage()    {}

func (m *STARSOUL_ARRAY) GetItems() []*STARSOUL {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
