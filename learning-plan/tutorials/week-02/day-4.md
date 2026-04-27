# Week 02 Day 4：实现 `join`，先把房间成员加进去

## 今天你要完成什么
今天只解决一个问题：
- 用户怎么加入房间

这一天不要同时做广播、鉴权、心跳。

## 今天你要重点改的文件
- `internal/room/service.go`
- `internal/room/room.go`

## 你今天的业务逻辑最小要求
当用户发来 `join` 消息时：
1. 根据 `room_id` 找到房间
2. 如果没有房间，就创建一个
3. 把当前 session 加进去
4. 返回当前成员数或加入成功结果

## 为什么今天要做 `room.Service`
因为你后面会发现，房间相关逻辑如果散在 router 和 handler 里，会越来越难维护。

今天最好有一个统一房间服务，负责：
- 获取房间
- 创建房间
- 注册成员
- 查询成员数

## 今天建议补的测试
- `TestJoinCreatesRoom`
- `TestJoinAddsMember`

## 今天不要做什么
- 不要做广播
- 不要做断开清理

今天只把加入这件事做扎实。

## 今天完成后的标准
- `join` 跑通
- 房间能创建
- 成员能被记录
- 至少两条相关测试通过

## 今天会用到的库和最小代码

### 最小 `Room` 结构
```go
type Room struct {
    ID      string
    Members map[string]*session.Session
}
```

### 最小加入逻辑
```go
func (s *Service) Join(roomID string, sess *session.Session) *Room {
    room, ok := s.rooms[roomID]
    if !ok {
        room = &Room{
            ID:      roomID,
            Members: map[string]*session.Session{},
        }
        s.rooms[roomID] = room
    }
    room.Members[sess.ConnID] = sess
    return room
}
```
