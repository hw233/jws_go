// automatically generated by the FlatBuffers compiler, do not modify

namespace teamboss_msg
{

using System;
using FlatBuffers;

/// [EnemyHP]伤害\损失HP通知
public sealed class EnenyHP : Table {
  public static EnenyHP GetRootAsEnenyHP(ByteBuffer _bb) { return GetRootAsEnenyHP(_bb, new EnenyHP()); }
  public static EnenyHP GetRootAsEnenyHP(ByteBuffer _bb, EnenyHP obj) { return (obj.__init(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public EnenyHP __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; return this; }

  public int Waves { get { int o = __offset(4); return o != 0 ? bb.GetInt(o + bb_pos) : (int)0; } }
  public long GetHp(int j) { int o = __offset(6); return o != 0 ? bb.GetLong(__vector(o) + j * 8) : (long)0; }
  public int HpLength { get { int o = __offset(6); return o != 0 ? __vector_len(o) : 0; } }
  public ArraySegment<byte>? GetHpBytes() { return __vector_as_arraysegment(6); }

  public static Offset<EnenyHP> CreateEnenyHP(FlatBufferBuilder builder,
      int waves = 0,
      VectorOffset hpOffset = default(VectorOffset)) {
    builder.StartObject(2);
    EnenyHP.AddHp(builder, hpOffset);
    EnenyHP.AddWaves(builder, waves);
    return EnenyHP.EndEnenyHP(builder);
  }

  public static void StartEnenyHP(FlatBufferBuilder builder) { builder.StartObject(2); }
  public static void AddWaves(FlatBufferBuilder builder, int waves) { builder.AddInt(0, waves, 0); }
  public static void AddHp(FlatBufferBuilder builder, VectorOffset hpOffset) { builder.AddOffset(1, hpOffset.Value, 0); }
  public static VectorOffset CreateHpVector(FlatBufferBuilder builder, long[] data) { builder.StartVector(8, data.Length, 8); for (int i = data.Length - 1; i >= 0; i--) builder.AddLong(data[i]); return builder.EndVector(); }
  public static void StartHpVector(FlatBufferBuilder builder, int numElems) { builder.StartVector(8, numElems, 8); }
  public static Offset<EnenyHP> EndEnenyHP(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<EnenyHP>(o);
  }
  public static void FinishEnenyHPBuffer(FlatBufferBuilder builder, Offset<EnenyHP> offset) { builder.Finish(offset.Value); }
};


}
