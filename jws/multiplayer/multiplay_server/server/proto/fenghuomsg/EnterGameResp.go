// automatically generated by the FlatBuffers compiler, do not modify

package fenghuomsg

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// [RPC]进入同步战斗服务器 resp
type EnterGameResp struct {
	_tab flatbuffers.Table
}

func GetRootAsEnterGameResp(buf []byte, offset flatbuffers.UOffsetT) *EnterGameResp {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EnterGameResp{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *EnterGameResp) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

///当前玩家索引
func (rcv *EnterGameResp) Myidx() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

///当前玩家索引
func (rcv *EnterGameResp) MutateMyidx(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

///当前服务器状态
func (rcv *EnterGameResp) Gamestatus() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

///当前服务器状态
func (rcv *EnterGameResp) MutateGamestatus(n int8) bool {
	return rcv._tab.MutateInt8Slot(6, n)
}

///敌兵所有血量
func (rcv *EnterGameResp) EnemiesHp(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *EnterGameResp) EnemiesHpLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

///敌兵所有血量
///玩家血量
func (rcv *EnterGameResp) Hps(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *EnterGameResp) HpsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

///玩家血量
///玩家Avatar属性
func (rcv *EnterGameResp) Avatars(obj *AccountInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(AccountInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *EnterGameResp) AvatarsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

///玩家Avatar属性
func EnterGameRespStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func EnterGameRespAddMyidx(builder *flatbuffers.Builder, myidx int32) {
	builder.PrependInt32Slot(0, myidx, 0)
}
func EnterGameRespAddGamestatus(builder *flatbuffers.Builder, gamestatus int8) {
	builder.PrependInt8Slot(1, gamestatus, 0)
}
func EnterGameRespAddEnemiesHp(builder *flatbuffers.Builder, enemiesHp flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(enemiesHp), 0)
}
func EnterGameRespStartEnemiesHpVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EnterGameRespAddHps(builder *flatbuffers.Builder, Hps flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(Hps), 0)
}
func EnterGameRespStartHpsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EnterGameRespAddAvatars(builder *flatbuffers.Builder, avatars flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(avatars), 0)
}
func EnterGameRespStartAvatarsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EnterGameRespEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
