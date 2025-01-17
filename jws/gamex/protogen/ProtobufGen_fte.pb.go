// Code generated by protoc-gen-go.
// source: ProtobufGen_fte.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type FTE struct {
	// * 引导的ID
	ID *string `protobuf:"bytes,1,req,def=" json:"ID,omitempty"`
	// * 引导的条件
	Conditions *string `protobuf:"bytes,2,opt,def=" json:"Conditions,omitempty"`
	// * 引导条件的参数
	ConditionParameters *string `protobuf:"bytes,3,opt,def=" json:"ConditionParameters,omitempty"`
	// * 引导的按钮
	PushButton *string `protobuf:"bytes,4,opt,def=" json:"PushButton,omitempty"`
	// * 按钮的参数
	ButtonParameters *string `protobuf:"bytes,5,opt,def=" json:"ButtonParameters,omitempty"`
	// * 引导按钮所在界面
	ButtonContainer *string `protobuf:"bytes,13,opt,def=" json:"ButtonContainer,omitempty"`
	// * 解释的文字
	Explain *string `protobuf:"bytes,6,opt,def=" json:"Explain,omitempty"`
	// * 播放语音
	Voice *string `protobuf:"bytes,15,opt,def=" json:"Voice,omitempty"`
	// * 解释文字框的位置
	Position *string `protobuf:"bytes,7,opt,def=" json:"Position,omitempty"`
	// * 引导后的行为
	Actions *string `protobuf:"bytes,8,opt,def=" json:"Actions,omitempty"`
	// * 引导后行为的参数
	ActionParameters *string `protobuf:"bytes,9,opt,def=" json:"ActionParameters,omitempty"`
	// * 引导允许弹出的界面类型
	AllowedNoticeItem *string `protobuf:"bytes,10,opt,def=" json:"AllowedNoticeItem,omitempty"`
	// * Skip按钮延迟出现的时间
	CanSkipDelayTime *string `protobuf:"bytes,11,opt,def=" json:"CanSkipDelayTime,omitempty"`
	// * Skip按钮延迟出现次数
	CanSkipClickTime *uint32 `protobuf:"varint,14,opt,def=0" json:"CanSkipClickTime,omitempty"`
	// * 引导的名字
	Name             *string `protobuf:"bytes,12,opt,def=" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FTE) Reset()         { *m = FTE{} }
func (m *FTE) String() string { return proto.CompactTextString(m) }
func (*FTE) ProtoMessage()    {}

const Default_FTE_CanSkipClickTime uint32 = 0

func (m *FTE) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *FTE) GetConditions() string {
	if m != nil && m.Conditions != nil {
		return *m.Conditions
	}
	return ""
}

func (m *FTE) GetConditionParameters() string {
	if m != nil && m.ConditionParameters != nil {
		return *m.ConditionParameters
	}
	return ""
}

func (m *FTE) GetPushButton() string {
	if m != nil && m.PushButton != nil {
		return *m.PushButton
	}
	return ""
}

func (m *FTE) GetButtonParameters() string {
	if m != nil && m.ButtonParameters != nil {
		return *m.ButtonParameters
	}
	return ""
}

func (m *FTE) GetButtonContainer() string {
	if m != nil && m.ButtonContainer != nil {
		return *m.ButtonContainer
	}
	return ""
}

func (m *FTE) GetExplain() string {
	if m != nil && m.Explain != nil {
		return *m.Explain
	}
	return ""
}

func (m *FTE) GetVoice() string {
	if m != nil && m.Voice != nil {
		return *m.Voice
	}
	return ""
}

func (m *FTE) GetPosition() string {
	if m != nil && m.Position != nil {
		return *m.Position
	}
	return ""
}

func (m *FTE) GetActions() string {
	if m != nil && m.Actions != nil {
		return *m.Actions
	}
	return ""
}

func (m *FTE) GetActionParameters() string {
	if m != nil && m.ActionParameters != nil {
		return *m.ActionParameters
	}
	return ""
}

func (m *FTE) GetAllowedNoticeItem() string {
	if m != nil && m.AllowedNoticeItem != nil {
		return *m.AllowedNoticeItem
	}
	return ""
}

func (m *FTE) GetCanSkipDelayTime() string {
	if m != nil && m.CanSkipDelayTime != nil {
		return *m.CanSkipDelayTime
	}
	return ""
}

func (m *FTE) GetCanSkipClickTime() uint32 {
	if m != nil && m.CanSkipClickTime != nil {
		return *m.CanSkipClickTime
	}
	return Default_FTE_CanSkipClickTime
}

func (m *FTE) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type FTE_ARRAY struct {
	Items            []*FTE `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FTE_ARRAY) Reset()         { *m = FTE_ARRAY{} }
func (m *FTE_ARRAY) String() string { return proto.CompactTextString(m) }
func (*FTE_ARRAY) ProtoMessage()    {}

func (m *FTE_ARRAY) GetItems() []*FTE {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
