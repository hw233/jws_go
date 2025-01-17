// automatically generated by the FlatBuffers compiler, do not modify

namespace teamboss_msg
{

using System;
using FlatBuffers;

///[RPC]获取当前战斗状态,resp
public sealed class GetGameStateRsp : Table {
  public static GetGameStateRsp GetRootAsGetGameStateRsp(ByteBuffer _bb) { return GetRootAsGetGameStateRsp(_bb, new GetGameStateRsp()); }
  public static GetGameStateRsp GetRootAsGetGameStateRsp(ByteBuffer _bb, GetGameStateRsp obj) { return (obj.__init(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public GetGameStateRsp __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; return this; }

  ///当前战斗状态(等待开始\已开始\已结束)
  public int Stat { get { int o = __offset(4); return o != 0 ? bb.GetInt(o + bb_pos) : (int)0; } }
  ///当前状态开始时间
  public long StartTime { get { int o = __offset(6); return o != 0 ? bb.GetLong(o + bb_pos) : (long)0; } }
  ///当前状态结束时间
  public long EndTime { get { int o = __offset(8); return o != 0 ? bb.GetLong(o + bb_pos) : (long)0; } }
  ///战斗难度
  public int GameClass { get { int o = __offset(10); return o != 0 ? bb.GetInt(o + bb_pos) : (int)0; } }
  ///战斗场景信息
  public string GameScene { get { int o = __offset(12); return o != 0 ? __string(o + bb_pos) : null; } }
  public ArraySegment<byte>? GetGameSceneBytes() { return __vector_as_arraysegment(12); }
  public teamboss_msg.PlayerState GetPlayerStat(int j) { return GetPlayerStat(new teamboss_msg.PlayerState(), j); }
  public teamboss_msg.PlayerState GetPlayerStat(teamboss_msg.PlayerState obj, int j) { int o = __offset(14); return o != 0 ? obj.__init(__indirect(__vector(o) + j * 4), bb) : null; }
  public int PlayerStatLength { get { int o = __offset(14); return o != 0 ? __vector_len(o) : 0; } }
  public teamboss_msg.BossState GetBossStat(int j) { return GetBossStat(new teamboss_msg.BossState(), j); }
  public teamboss_msg.BossState GetBossStat(teamboss_msg.BossState obj, int j) { int o = __offset(16); return o != 0 ? obj.__init(__indirect(__vector(o) + j * 4), bb) : null; }
  public int BossStatLength { get { int o = __offset(16); return o != 0 ? __vector_len(o) : 0; } }

  public static Offset<GetGameStateRsp> CreateGetGameStateRsp(FlatBufferBuilder builder,
      int stat = 0,
      long startTime = 0,
      long endTime = 0,
      int GameClass = 0,
      StringOffset GameSceneOffset = default(StringOffset),
      VectorOffset playerStatOffset = default(VectorOffset),
      VectorOffset bossStatOffset = default(VectorOffset)) {
    builder.StartObject(7);
    GetGameStateRsp.AddEndTime(builder, endTime);
    GetGameStateRsp.AddStartTime(builder, startTime);
    GetGameStateRsp.AddBossStat(builder, bossStatOffset);
    GetGameStateRsp.AddPlayerStat(builder, playerStatOffset);
    GetGameStateRsp.AddGameScene(builder, GameSceneOffset);
    GetGameStateRsp.AddGameClass(builder, GameClass);
    GetGameStateRsp.AddStat(builder, stat);
    return GetGameStateRsp.EndGetGameStateRsp(builder);
  }

  public static void StartGetGameStateRsp(FlatBufferBuilder builder) { builder.StartObject(7); }
  public static void AddStat(FlatBufferBuilder builder, int stat) { builder.AddInt(0, stat, 0); }
  public static void AddStartTime(FlatBufferBuilder builder, long startTime) { builder.AddLong(1, startTime, 0); }
  public static void AddEndTime(FlatBufferBuilder builder, long endTime) { builder.AddLong(2, endTime, 0); }
  public static void AddGameClass(FlatBufferBuilder builder, int GameClass) { builder.AddInt(3, GameClass, 0); }
  public static void AddGameScene(FlatBufferBuilder builder, StringOffset GameSceneOffset) { builder.AddOffset(4, GameSceneOffset.Value, 0); }
  public static void AddPlayerStat(FlatBufferBuilder builder, VectorOffset playerStatOffset) { builder.AddOffset(5, playerStatOffset.Value, 0); }
  public static VectorOffset CreatePlayerStatVector(FlatBufferBuilder builder, Offset<teamboss_msg.PlayerState>[] data) { builder.StartVector(4, data.Length, 4); for (int i = data.Length - 1; i >= 0; i--) builder.AddOffset(data[i].Value); return builder.EndVector(); }
  public static void StartPlayerStatVector(FlatBufferBuilder builder, int numElems) { builder.StartVector(4, numElems, 4); }
  public static void AddBossStat(FlatBufferBuilder builder, VectorOffset bossStatOffset) { builder.AddOffset(6, bossStatOffset.Value, 0); }
  public static VectorOffset CreateBossStatVector(FlatBufferBuilder builder, Offset<teamboss_msg.BossState>[] data) { builder.StartVector(4, data.Length, 4); for (int i = data.Length - 1; i >= 0; i--) builder.AddOffset(data[i].Value); return builder.EndVector(); }
  public static void StartBossStatVector(FlatBufferBuilder builder, int numElems) { builder.StartVector(4, numElems, 4); }
  public static Offset<GetGameStateRsp> EndGetGameStateRsp(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<GetGameStateRsp>(o);
  }
  public static void FinishGetGameStateRspBuffer(FlatBufferBuilder builder, Offset<GetGameStateRsp> offset) { builder.Finish(offset.Value); }
};


}
