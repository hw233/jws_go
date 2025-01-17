// automatically generated by the FlatBuffers compiler, do not modify

package multiplayMsg

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Ping struct {
	_tab flatbuffers.Table
}

func GetRootAsPing(buf []byte, offset flatbuffers.UOffsetT) *Ping {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Ping{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Ping) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Ping) Time() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Ping) MutateTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *Ping) Acid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func PingStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PingAddTime(builder *flatbuffers.Builder, time int64) {
	builder.PrependInt64Slot(0, time, 0)
}
func PingAddAcid(builder *flatbuffers.Builder, acid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(acid), 0)
}
func PingEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
