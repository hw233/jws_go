// automatically generated by the FlatBuffers compiler, do not modify

namespace teamboss_msg
{

using System;
using FlatBuffers;

/// [Notify]主动离开战斗服务器(状态算退出)
public sealed class LeaveMultiplayGameNotify : Table {
  public static LeaveMultiplayGameNotify GetRootAsLeaveMultiplayGameNotify(ByteBuffer _bb) { return GetRootAsLeaveMultiplayGameNotify(_bb, new LeaveMultiplayGameNotify()); }
  public static LeaveMultiplayGameNotify GetRootAsLeaveMultiplayGameNotify(ByteBuffer _bb, LeaveMultiplayGameNotify obj) { return (obj.__init(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public LeaveMultiplayGameNotify __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; return this; }

  ///玩家账号ID
  public string AccountId { get { int o = __offset(4); return o != 0 ? __string(o + bb_pos) : null; } }
  public ArraySegment<byte>? GetAccountIdBytes() { return __vector_as_arraysegment(4); }
  ///请求进入房间ID
  public string RoomID { get { int o = __offset(6); return o != 0 ? __string(o + bb_pos) : null; } }
  public ArraySegment<byte>? GetRoomIDBytes() { return __vector_as_arraysegment(6); }

  public static Offset<LeaveMultiplayGameNotify> CreateLeaveMultiplayGameNotify(FlatBufferBuilder builder,
      StringOffset accountIdOffset = default(StringOffset),
      StringOffset roomIDOffset = default(StringOffset)) {
    builder.StartObject(2);
    LeaveMultiplayGameNotify.AddRoomID(builder, roomIDOffset);
    LeaveMultiplayGameNotify.AddAccountId(builder, accountIdOffset);
    return LeaveMultiplayGameNotify.EndLeaveMultiplayGameNotify(builder);
  }

  public static void StartLeaveMultiplayGameNotify(FlatBufferBuilder builder) { builder.StartObject(2); }
  public static void AddAccountId(FlatBufferBuilder builder, StringOffset accountIdOffset) { builder.AddOffset(0, accountIdOffset.Value, 0); }
  public static void AddRoomID(FlatBufferBuilder builder, StringOffset roomIDOffset) { builder.AddOffset(1, roomIDOffset.Value, 0); }
  public static Offset<LeaveMultiplayGameNotify> EndLeaveMultiplayGameNotify(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<LeaveMultiplayGameNotify>(o);
  }
  public static void FinishLeaveMultiplayGameNotifyBuffer(FlatBufferBuilder builder, Offset<LeaveMultiplayGameNotify> offset) { builder.Finish(offset.Value); }
};


}
