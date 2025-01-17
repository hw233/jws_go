// Code generated by protoc-gen-go.
// source: ProtobufGen_lateraccessloading.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type LATERACCESSLOADING struct {
	// * 玩家等级
	Level               *int32                            `protobuf:"varint,1,req,name=level,def=0" json:"level,omitempty"`
	LoadingImg_Template []*LATERACCESSLOADING_LoadingImgs `protobuf:"bytes,2,rep" json:"LoadingImg_Template,omitempty"`
	XXX_unrecognized    []byte                            `json:"-"`
}

func (m *LATERACCESSLOADING) Reset()         { *m = LATERACCESSLOADING{} }
func (m *LATERACCESSLOADING) String() string { return proto.CompactTextString(m) }
func (*LATERACCESSLOADING) ProtoMessage()    {}

const Default_LATERACCESSLOADING_Level int32 = 0

func (m *LATERACCESSLOADING) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return Default_LATERACCESSLOADING_Level
}

func (m *LATERACCESSLOADING) GetLoadingImg_Template() []*LATERACCESSLOADING_LoadingImgs {
	if m != nil {
		return m.LoadingImg_Template
	}
	return nil
}

type LATERACCESSLOADING_LoadingImgs struct {
	// * loading图片
	LoadingImg *string `protobuf:"bytes,1,opt,def=" json:"LoadingImg,omitempty"`
	// * 随中权值
	Weight           *int32 `protobuf:"varint,2,opt,def=0" json:"Weight,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LATERACCESSLOADING_LoadingImgs) Reset()         { *m = LATERACCESSLOADING_LoadingImgs{} }
func (m *LATERACCESSLOADING_LoadingImgs) String() string { return proto.CompactTextString(m) }
func (*LATERACCESSLOADING_LoadingImgs) ProtoMessage()    {}

const Default_LATERACCESSLOADING_LoadingImgs_Weight int32 = 0

func (m *LATERACCESSLOADING_LoadingImgs) GetLoadingImg() string {
	if m != nil && m.LoadingImg != nil {
		return *m.LoadingImg
	}
	return ""
}

func (m *LATERACCESSLOADING_LoadingImgs) GetWeight() int32 {
	if m != nil && m.Weight != nil {
		return *m.Weight
	}
	return Default_LATERACCESSLOADING_LoadingImgs_Weight
}

type LATERACCESSLOADING_ARRAY struct {
	Items            []*LATERACCESSLOADING `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *LATERACCESSLOADING_ARRAY) Reset()         { *m = LATERACCESSLOADING_ARRAY{} }
func (m *LATERACCESSLOADING_ARRAY) String() string { return proto.CompactTextString(m) }
func (*LATERACCESSLOADING_ARRAY) ProtoMessage()    {}

func (m *LATERACCESSLOADING_ARRAY) GetItems() []*LATERACCESSLOADING {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
