// automatically generated by the FlatBuffers compiler, do not modify

namespace fenghuomsg
{

public enum FenghuoGameStatus : sbyte
{
 WaitingInit = 0,
  /// 可以开始战斗,选择小关卡后,收到了StartFight消息
 SubLevelStart = 1,
  /// 小关卡战斗胜利后
 SubLevelDone = 2,
 GameOver = 3,
  /// 强制结束
 ForceOver = 4,
};


}
