// automatically generated by the FlatBuffers compiler, do not modify

package multiplayMsg

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///[RPC]进入同步战斗服务器 返回
type EnterMultiplayGameRsp struct {
	_tab flatbuffers.Table
}

func GetRootAsEnterMultiplayGameRsp(buf []byte, offset flatbuffers.UOffsetT) *EnterMultiplayGameRsp {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EnterMultiplayGameRsp{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *EnterMultiplayGameRsp) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

///当前战斗状态(等待开始\已开始\已结束)
func (rcv *EnterMultiplayGameRsp) Stat() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

///当前战斗状态(等待开始\已开始\已结束)
func (rcv *EnterMultiplayGameRsp) MutateStat(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

///当前状态开始时间
func (rcv *EnterMultiplayGameRsp) StartTime() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

///当前状态开始时间
func (rcv *EnterMultiplayGameRsp) MutateStartTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

///当前状态结束时间
func (rcv *EnterMultiplayGameRsp) EndTime() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

///当前状态结束时间
func (rcv *EnterMultiplayGameRsp) MutateEndTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

///战斗难度
func (rcv *EnterMultiplayGameRsp) GameClass() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

///战斗难度
func (rcv *EnterMultiplayGameRsp) MutateGameClass(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

///战斗场景信息
func (rcv *EnterMultiplayGameRsp) GameScene() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///战斗场景信息
///当前房间中玩家的数据
func (rcv *EnterMultiplayGameRsp) AccDatas(obj *AccountInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
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

func (rcv *EnterMultiplayGameRsp) AccDatasLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

///当前房间中玩家的数据
///当前房间中Boss的数据
func (rcv *EnterMultiplayGameRsp) AcDatas(obj *AcDataInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(AcDataInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *EnterMultiplayGameRsp) AcDatasLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

///当前房间中Boss的数据
///当前房间中玩家的状态
func (rcv *EnterMultiplayGameRsp) PlayerStat(obj *PlayerState, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
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

func (rcv *EnterMultiplayGameRsp) PlayerStatLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

///当前房间中玩家的状态
///当前房间中Boss的状态
func (rcv *EnterMultiplayGameRsp) BossStat(obj *BossState, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
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

func (rcv *EnterMultiplayGameRsp) BossStatLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

///当前房间中Boss的状态
func (rcv *EnterMultiplayGameRsp) TimeStampS() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *EnterMultiplayGameRsp) MutateTimeStampS(n int64) bool {
	return rcv._tab.MutateInt64Slot(22, n)
}

func (rcv *EnterMultiplayGameRsp) TimeStampNS() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *EnterMultiplayGameRsp) MutateTimeStampNS(n int32) bool {
	return rcv._tab.MutateInt32Slot(24, n)
}

func (rcv *EnterMultiplayGameRsp) Pos() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *EnterMultiplayGameRsp) MutatePos(n int32) bool {
	return rcv._tab.MutateInt32Slot(26, n)
}

func EnterMultiplayGameRspStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func EnterMultiplayGameRspAddStat(builder *flatbuffers.Builder, stat int32) {
	builder.PrependInt32Slot(0, stat, 0)
}
func EnterMultiplayGameRspAddStartTime(builder *flatbuffers.Builder, startTime int64) {
	builder.PrependInt64Slot(1, startTime, 0)
}
func EnterMultiplayGameRspAddEndTime(builder *flatbuffers.Builder, endTime int64) {
	builder.PrependInt64Slot(2, endTime, 0)
}
func EnterMultiplayGameRspAddGameClass(builder *flatbuffers.Builder, GameClass int32) {
	builder.PrependInt32Slot(3, GameClass, 0)
}
func EnterMultiplayGameRspAddGameScene(builder *flatbuffers.Builder, GameScene flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(GameScene), 0)
}
func EnterMultiplayGameRspAddAccDatas(builder *flatbuffers.Builder, accDatas flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(accDatas), 0)
}
func EnterMultiplayGameRspStartAccDatasVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EnterMultiplayGameRspAddAcDatas(builder *flatbuffers.Builder, acDatas flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(acDatas), 0)
}
func EnterMultiplayGameRspStartAcDatasVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EnterMultiplayGameRspAddPlayerStat(builder *flatbuffers.Builder, playerStat flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(playerStat), 0)
}
func EnterMultiplayGameRspStartPlayerStatVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EnterMultiplayGameRspAddBossStat(builder *flatbuffers.Builder, bossStat flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(bossStat), 0)
}
func EnterMultiplayGameRspStartBossStatVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EnterMultiplayGameRspAddTimeStampS(builder *flatbuffers.Builder, timeStampS int64) {
	builder.PrependInt64Slot(9, timeStampS, 0)
}
func EnterMultiplayGameRspAddTimeStampNS(builder *flatbuffers.Builder, timeStampNS int32) {
	builder.PrependInt32Slot(10, timeStampNS, 0)
}
func EnterMultiplayGameRspAddPos(builder *flatbuffers.Builder, Pos int32) {
	builder.PrependInt32Slot(11, Pos, 0)
}
func EnterMultiplayGameRspEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}