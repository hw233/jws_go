// Code generated by protoc-gen-go.
// source: ProtobufGen_dialogue_info.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type DIALOGUE_INFO struct {
	// * 对话ID
	ID *string `protobuf:"bytes,1,req,def=" json:"ID,omitempty"`
	// * 对话类型
	Type *string `protobuf:"bytes,2,opt,def=" json:"Type,omitempty"`
	// * 对话内容
	Sentence *string `protobuf:"bytes,3,opt,name=sentence,def=" json:"sentence,omitempty"`
	// * 对话场景
	Scene            *string `protobuf:"bytes,4,opt,name=scene,def=" json:"scene,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DIALOGUE_INFO) Reset()         { *m = DIALOGUE_INFO{} }
func (m *DIALOGUE_INFO) String() string { return proto.CompactTextString(m) }
func (*DIALOGUE_INFO) ProtoMessage()    {}

func (m *DIALOGUE_INFO) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *DIALOGUE_INFO) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *DIALOGUE_INFO) GetSentence() string {
	if m != nil && m.Sentence != nil {
		return *m.Sentence
	}
	return ""
}

func (m *DIALOGUE_INFO) GetScene() string {
	if m != nil && m.Scene != nil {
		return *m.Scene
	}
	return ""
}

type DIALOGUE_INFO_ARRAY struct {
	Items            []*DIALOGUE_INFO `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *DIALOGUE_INFO_ARRAY) Reset()         { *m = DIALOGUE_INFO_ARRAY{} }
func (m *DIALOGUE_INFO_ARRAY) String() string { return proto.CompactTextString(m) }
func (*DIALOGUE_INFO_ARRAY) ProtoMessage()    {}

func (m *DIALOGUE_INFO_ARRAY) GetItems() []*DIALOGUE_INFO {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
