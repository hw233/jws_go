// automatically generated by the FlatBuffers compiler, do not modify

namespace multiplayMsg
{

using System;
using FlatBuffers;

public sealed class Ping : Table {
  public static Ping GetRootAsPing(ByteBuffer _bb) { return GetRootAsPing(_bb, new Ping()); }
  public static Ping GetRootAsPing(ByteBuffer _bb, Ping obj) { return (obj.__init(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public Ping __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; return this; }

  public long Time { get { int o = __offset(4); return o != 0 ? bb.GetLong(o + bb_pos) : (long)0; } }
  public string Acid { get { int o = __offset(6); return o != 0 ? __string(o + bb_pos) : null; } }
  public ArraySegment<byte>? GetAcidBytes() { return __vector_as_arraysegment(6); }

  public static Offset<Ping> CreatePing(FlatBufferBuilder builder,
      long time = 0,
      StringOffset acidOffset = default(StringOffset)) {
    builder.StartObject(2);
    Ping.AddTime(builder, time);
    Ping.AddAcid(builder, acidOffset);
    return Ping.EndPing(builder);
  }

  public static void StartPing(FlatBufferBuilder builder) { builder.StartObject(2); }
  public static void AddTime(FlatBufferBuilder builder, long time) { builder.AddLong(0, time, 0); }
  public static void AddAcid(FlatBufferBuilder builder, StringOffset acidOffset) { builder.AddOffset(1, acidOffset.Value, 0); }
  public static Offset<Ping> EndPing(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<Ping>(o);
  }
  public static void FinishPingBuffer(FlatBufferBuilder builder, Offset<Ping> offset) { builder.Finish(offset.Value); }
};


}
