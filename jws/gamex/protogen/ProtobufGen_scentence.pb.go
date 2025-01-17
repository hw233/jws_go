// Code generated by protoc-gen-go.
// source: ProtobufGen_scentence.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SCENTENCE struct {
	// * 句子ID
	ID *string `protobuf:"bytes,1,req,def=" json:"ID,omitempty"`
	// * 句子的stringID
	Content *string `protobuf:"bytes,2,opt,def=" json:"Content,omitempty"`
	// * 讲话人类别
	RoleType *int32 `protobuf:"varint,3,opt,def=0" json:"RoleType,omitempty"`
	// * 具有特殊指示意义的ID
	RoleID *string `protobuf:"bytes,4,opt,def=" json:"RoleID,omitempty"`
	// * 图像显示位置
	RolePic *int32 `protobuf:"varint,5,opt,def=0" json:"RolePic,omitempty"`
	// * 对话表情
	RoleFace *string `protobuf:"bytes,6,opt,def=" json:"RoleFace,omitempty"`
	// * 对话持续时间
	RoleTime         *float32 `protobuf:"fixed32,7,opt,def=3" json:"RoleTime,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SCENTENCE) Reset()         { *m = SCENTENCE{} }
func (m *SCENTENCE) String() string { return proto.CompactTextString(m) }
func (*SCENTENCE) ProtoMessage()    {}

const Default_SCENTENCE_RoleType int32 = 0
const Default_SCENTENCE_RolePic int32 = 0
const Default_SCENTENCE_RoleTime float32 = 3

func (m *SCENTENCE) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *SCENTENCE) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *SCENTENCE) GetRoleType() int32 {
	if m != nil && m.RoleType != nil {
		return *m.RoleType
	}
	return Default_SCENTENCE_RoleType
}

func (m *SCENTENCE) GetRoleID() string {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return ""
}

func (m *SCENTENCE) GetRolePic() int32 {
	if m != nil && m.RolePic != nil {
		return *m.RolePic
	}
	return Default_SCENTENCE_RolePic
}

func (m *SCENTENCE) GetRoleFace() string {
	if m != nil && m.RoleFace != nil {
		return *m.RoleFace
	}
	return ""
}

func (m *SCENTENCE) GetRoleTime() float32 {
	if m != nil && m.RoleTime != nil {
		return *m.RoleTime
	}
	return Default_SCENTENCE_RoleTime
}

type SCENTENCE_ARRAY struct {
	Items            []*SCENTENCE `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SCENTENCE_ARRAY) Reset()         { *m = SCENTENCE_ARRAY{} }
func (m *SCENTENCE_ARRAY) String() string { return proto.CompactTextString(m) }
func (*SCENTENCE_ARRAY) ProtoMessage()    {}

func (m *SCENTENCE_ARRAY) GetItems() []*SCENTENCE {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
