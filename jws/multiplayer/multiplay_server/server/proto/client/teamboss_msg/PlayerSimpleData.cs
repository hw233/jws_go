// automatically generated by the FlatBuffers compiler, do not modify

namespace teamboss_msg
{

using System;
using FlatBuffers;

/// [Notify]伤害\损失HP通知
public sealed class PlayerSimpleData : Table {
  public static PlayerSimpleData GetRootAsPlayerSimpleData(ByteBuffer _bb) { return GetRootAsPlayerSimpleData(_bb, new PlayerSimpleData()); }
  public static PlayerSimpleData GetRootAsPlayerSimpleData(ByteBuffer _bb, PlayerSimpleData obj) { return (obj.__init(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public PlayerSimpleData __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; return this; }

  ///玩家账号ID
  public string AccountId { get { int o = __offset(4); return o != 0 ? __string(o + bb_pos) : null; } }
  public ArraySegment<byte>? GetAccountIdBytes() { return __vector_as_arraysegment(4); }
  ///请求进入房间ID
  public string RoomID { get { int o = __offset(6); return o != 0 ? __string(o + bb_pos) : null; } }
  public ArraySegment<byte>? GetRoomIDBytes() { return __vector_as_arraysegment(6); }
  public int State { get { int o = __offset(8); return o != 0 ? bb.GetInt(o + bb_pos) : (int)0; } }
  public teamboss_msg.Vector Position { get { return GetPosition(new teamboss_msg.Vector()); } }
  public teamboss_msg.Vector GetPosition(teamboss_msg.Vector obj) { int o = __offset(10); return o != 0 ? obj.__init(__indirect(o + bb_pos), bb) : null; }
  public teamboss_msg.Vector Direction { get { return GetDirection(new teamboss_msg.Vector()); } }
  public teamboss_msg.Vector GetDirection(teamboss_msg.Vector obj) { int o = __offset(12); return o != 0 ? obj.__init(__indirect(o + bb_pos), bb) : null; }
  public float CurSpeed { get { int o = __offset(14); return o != 0 ? bb.GetFloat(o + bb_pos) : (float)0.0f; } }
  public long NotifyTimeStamp { get { int o = __offset(16); return o != 0 ? bb.GetLong(o + bb_pos) : (long)0; } }
  public int Anim { get { int o = __offset(18); return o != 0 ? bb.GetInt(o + bb_pos) : (int)0; } }

  public static Offset<PlayerSimpleData> CreatePlayerSimpleData(FlatBufferBuilder builder,
      StringOffset accountIdOffset = default(StringOffset),
      StringOffset roomIDOffset = default(StringOffset),
      int state = 0,
      Offset<teamboss_msg.Vector> positionOffset = default(Offset<teamboss_msg.Vector>),
      Offset<teamboss_msg.Vector> directionOffset = default(Offset<teamboss_msg.Vector>),
      float curSpeed = 0.0f,
      long notifyTimeStamp = 0,
      int anim = 0) {
    builder.StartObject(8);
    PlayerSimpleData.AddNotifyTimeStamp(builder, notifyTimeStamp);
    PlayerSimpleData.AddAnim(builder, anim);
    PlayerSimpleData.AddCurSpeed(builder, curSpeed);
    PlayerSimpleData.AddDirection(builder, directionOffset);
    PlayerSimpleData.AddPosition(builder, positionOffset);
    PlayerSimpleData.AddState(builder, state);
    PlayerSimpleData.AddRoomID(builder, roomIDOffset);
    PlayerSimpleData.AddAccountId(builder, accountIdOffset);
    return PlayerSimpleData.EndPlayerSimpleData(builder);
  }

  public static void StartPlayerSimpleData(FlatBufferBuilder builder) { builder.StartObject(8); }
  public static void AddAccountId(FlatBufferBuilder builder, StringOffset accountIdOffset) { builder.AddOffset(0, accountIdOffset.Value, 0); }
  public static void AddRoomID(FlatBufferBuilder builder, StringOffset roomIDOffset) { builder.AddOffset(1, roomIDOffset.Value, 0); }
  public static void AddState(FlatBufferBuilder builder, int state) { builder.AddInt(2, state, 0); }
  public static void AddPosition(FlatBufferBuilder builder, Offset<teamboss_msg.Vector> positionOffset) { builder.AddOffset(3, positionOffset.Value, 0); }
  public static void AddDirection(FlatBufferBuilder builder, Offset<teamboss_msg.Vector> directionOffset) { builder.AddOffset(4, directionOffset.Value, 0); }
  public static void AddCurSpeed(FlatBufferBuilder builder, float curSpeed) { builder.AddFloat(5, curSpeed, 0.0f); }
  public static void AddNotifyTimeStamp(FlatBufferBuilder builder, long notifyTimeStamp) { builder.AddLong(6, notifyTimeStamp, 0); }
  public static void AddAnim(FlatBufferBuilder builder, int anim) { builder.AddInt(7, anim, 0); }
  public static Offset<PlayerSimpleData> EndPlayerSimpleData(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<PlayerSimpleData>(o);
  }
  public static void FinishPlayerSimpleDataBuffer(FlatBufferBuilder builder, Offset<PlayerSimpleData> offset) { builder.Finish(offset.Value); }
};


}
