// automatically generated by the FlatBuffers compiler, do not modify

package multiplayMsg

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///[RPC]获取当前战斗状态,resp
type GetGameStateRsp struct {
	_tab flatbuffers.Table
}

func GetRootAsGetGameStateRsp(buf []byte, offset flatbuffers.UOffsetT) *GetGameStateRsp {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GetGameStateRsp{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *GetGameStateRsp) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

///当前战斗状态(等待开始\已开始\已结束)
func (rcv *GetGameStateRsp) Stat() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

///当前战斗状态(等待开始\已开始\已结束)
func (rcv *GetGameStateRsp) MutateStat(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

///当前状态开始时间
func (rcv *GetGameStateRsp) StartTime() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

///当前状态开始时间
func (rcv *GetGameStateRsp) MutateStartTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

///当前状态结束时间
func (rcv *GetGameStateRsp) EndTime() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

///当前状态结束时间
func (rcv *GetGameStateRsp) MutateEndTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

///战斗难度
func (rcv *GetGameStateRsp) GameClass() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

///战斗难度
func (rcv *GetGameStateRsp) MutateGameClass(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

///战斗场景信息
func (rcv *GetGameStateRsp) GameScene() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///战斗场景信息
func (rcv *GetGameStateRsp) PlayerStat(obj *PlayerState, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(PlayerState)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *GetGameStateRsp) PlayerStatLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *GetGameStateRsp) BossStat(obj *BossState, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(BossState)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *GetGameStateRsp) BossStatLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func GetGameStateRspStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func GetGameStateRspAddStat(builder *flatbuffers.Builder, stat int32) {
	builder.PrependInt32Slot(0, stat, 0)
}
func GetGameStateRspAddStartTime(builder *flatbuffers.Builder, startTime int64) {
	builder.PrependInt64Slot(1, startTime, 0)
}
func GetGameStateRspAddEndTime(builder *flatbuffers.Builder, endTime int64) {
	builder.PrependInt64Slot(2, endTime, 0)
}
func GetGameStateRspAddGameClass(builder *flatbuffers.Builder, GameClass int32) {
	builder.PrependInt32Slot(3, GameClass, 0)
}
func GetGameStateRspAddGameScene(builder *flatbuffers.Builder, GameScene flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(GameScene), 0)
}
func GetGameStateRspAddPlayerStat(builder *flatbuffers.Builder, playerStat flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(playerStat), 0)
}
func GetGameStateRspStartPlayerStatVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func GetGameStateRspAddBossStat(builder *flatbuffers.Builder, bossStat flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(bossStat), 0)
}
func GetGameStateRspStartBossStatVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func GetGameStateRspEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
