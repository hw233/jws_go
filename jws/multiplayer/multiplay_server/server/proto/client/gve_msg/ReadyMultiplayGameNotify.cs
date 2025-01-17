// automatically generated by the FlatBuffers compiler, do not modify

namespace gve_msg
{

using System;
using FlatBuffers;

///[Notify]准备开始战斗
public sealed class ReadyMultiplayGameNotify : Table {
  public static ReadyMultiplayGameNotify GetRootAsReadyMultiplayGameNotify(ByteBuffer _bb) { return GetRootAsReadyMultiplayGameNotify(_bb, new ReadyMultiplayGameNotify()); }
  public static ReadyMultiplayGameNotify GetRootAsReadyMultiplayGameNotify(ByteBuffer _bb, ReadyMultiplayGameNotify obj) { return (obj.__init(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public ReadyMultiplayGameNotify __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; return this; }

  ///玩家账号ID
  public string AccountId { get { int o = __offset(4); return o != 0 ? __string(o + bb_pos) : null; } }
  public ArraySegment<byte>? GetAccountIdBytes() { return __vector_as_arraysegment(4); }
  ///请求进入房间ID
  public string RoomID { get { int o = __offset(6); return o != 0 ? __string(o + bb_pos) : null; } }
  public ArraySegment<byte>? GetRoomIDBytes() { return __vector_as_arraysegment(6); }

  public static Offset<ReadyMultiplayGameNotify> CreateReadyMultiplayGameNotify(FlatBufferBuilder builder,
      StringOffset accountIdOffset = default(StringOffset),
      StringOffset roomIDOffset = default(StringOffset)) {
    builder.StartObject(2);
    ReadyMultiplayGameNotify.AddRoomID(builder, roomIDOffset);
    ReadyMultiplayGameNotify.AddAccountId(builder, accountIdOffset);
    return ReadyMultiplayGameNotify.EndReadyMultiplayGameNotify(builder);
  }

  public static void StartReadyMultiplayGameNotify(FlatBufferBuilder builder) { builder.StartObject(2); }
  public static void AddAccountId(FlatBufferBuilder builder, StringOffset accountIdOffset) { builder.AddOffset(0, accountIdOffset.Value, 0); }
  public static void AddRoomID(FlatBufferBuilder builder, StringOffset roomIDOffset) { builder.AddOffset(1, roomIDOffset.Value, 0); }
  public static Offset<ReadyMultiplayGameNotify> EndReadyMultiplayGameNotify(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<ReadyMultiplayGameNotify>(o);
  }
  public static void FinishReadyMultiplayGameNotifyBuffer(FlatBufferBuilder builder, Offset<ReadyMultiplayGameNotify> offset) { builder.Finish(offset.Value); }
};


}
