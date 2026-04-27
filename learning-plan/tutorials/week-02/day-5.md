# Week 02 Day 5：实现 `chat`，但先做内存版广播

## 今天你要完成什么
今天做第二个核心动作：
- `chat`

也就是：一个房间里的消息如何发给房间成员。

## 先说一个很重要的原则
今天你不一定要实现真实网络写回。

如果你现在网络层还没完全搭好，今天也可以先做：
- 计算应该广播给哪些成员
- 让房间服务返回广播目标列表

这依然是正确的进展。

## 今天重点改哪些地方
- `internal/room/service.go`
- `internal/server/router.go`

## 你今天的最小广播逻辑
当收到 `chat` 时：
1. 找到当前房间
2. 找到房间里的成员
3. 生成广播目标
4. 把消息分发给这些目标

## 今天为什么不追求真实 socket 写回
因为今天的目标是“把房间内消息流向讲清楚”。

如果你把网络细节和业务逻辑绑死，后面你补：
- 鉴权
- 心跳
- 超时回收

都会更难。

## 今天建议补的测试
- `TestBroadcastTargetsMembersInSameRoom`
- `TestBroadcastDoesNotLeakToOtherRooms`

## 今天完成后的标准
- `chat` 路由已存在
- 广播目标生成逻辑已存在
- 不同房间之间不会串消息

## 今天会用到的库和最小代码

### 最小广播目标生成
```go
func (s *Service) BroadcastTargets(roomID string) []*session.Session {
    room, ok := s.rooms[roomID]
    if !ok {
        return nil
    }
    out := make([]*session.Session, 0, len(room.Members))
    for _, sess := range room.Members {
        out = append(out, sess)
    }
    return out
}
```

### 为什么先返回目标列表
因为这样你可以把“房间逻辑”和“网络写回逻辑”先拆开。
