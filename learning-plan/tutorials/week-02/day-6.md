# Week 02 Day 6：做连接退出清理，不要让房间里留僵尸成员

## 今天你要完成什么
今天专门处理一个很容易被忽略的问题：
- 连接断开之后怎么办

如果你不做这一步，房间里的成员表很快就会出现大量无效 session。

## 今天你至少要解决两个问题
1. 连接断开后，session 从房间移除
2. 房间为空后，是否回收

## 今天建议新增的方法
例如：
- `func (s *Service) Disconnect(connID string)`

这个方法的好处是，后面不管是：
- 主动断开
- 心跳超时
- 网络错误

都可以复用同一套清理逻辑。

## 今天建议你重点看什么
你要确保清理逻辑不会留下这几类垃圾状态：
- session 还在，房间已无效
- 房间成员还在，但连接早没了
- 房间没人了，但房间对象永远不删

## 今天建议补的测试
- `TestDisconnectRemovesSessionFromRoom`
- `TestEmptyRoomCleanup`

## 今天完成后的标准
- 断开逻辑已存在
- 房间成员可被清理
- 至少 1 条生命周期测试通过

## 今天会用到的库和最小代码

### 最小断开清理
```go
func (s *Service) Disconnect(connID string) {
    for _, room := range s.rooms {
        delete(room.Members, connID)
        if len(room.Members) == 0 {
            delete(s.rooms, room.ID)
        }
    }
}
```

### 为什么这里后面还要重构
因为这段代码在并发场景下还不安全，但今天先把生命周期逻辑讲清楚。
