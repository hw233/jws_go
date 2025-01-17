// Code generated by protoc-gen-go.
// source: ProtobufGen_npcui.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type NPCUI struct {
	// * UI的id
	UiId *string `protobuf:"bytes,1,req,name=uiId,def=" json:"uiId,omitempty"`
	// * 对应的NPC的id
	NpcId            *string `protobuf:"bytes,2,opt,name=npcId,def=" json:"npcId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NPCUI) Reset()         { *m = NPCUI{} }
func (m *NPCUI) String() string { return proto.CompactTextString(m) }
func (*NPCUI) ProtoMessage()    {}

func (m *NPCUI) GetUiId() string {
	if m != nil && m.UiId != nil {
		return *m.UiId
	}
	return ""
}

func (m *NPCUI) GetNpcId() string {
	if m != nil && m.NpcId != nil {
		return *m.NpcId
	}
	return ""
}

type NPCUI_ARRAY struct {
	Items            []*NPCUI `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NPCUI_ARRAY) Reset()         { *m = NPCUI_ARRAY{} }
func (m *NPCUI_ARRAY) String() string { return proto.CompactTextString(m) }
func (*NPCUI_ARRAY) ProtoMessage()    {}

func (m *NPCUI_ARRAY) GetItems() []*NPCUI {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
