# Week 03 Day 3：补消息类型，把路由真正变成“入口”

## 今天你要完成什么
Week 02 里你至少已经有：
- `join`
- `message`

今天你要让路由更像一个真实入口，而不是只处理两个示例动作。

## 今天建议补哪两个消息
优先级建议：
1. `heartbeat`
2. `leave`

如果你还有余力，再做：
3. `room-info`

## 为什么优先补 `heartbeat`
因为它和“连接是否还活着”直接相关，是进入 Week 04 稳定性之前最自然的一步。

## 为什么 `leave` 比 `room-info` 更重要
因为 `leave` 会直接影响房间成员表的正确性，而 `room-info` 更多是查询能力。

## 今天要改哪些地方
- `internal/protocol/message.go`
- `internal/server/router.go`
- `internal/room/service.go`

## 今天建议补的测试
- `TestRouteHeartbeatMessage`
- `TestRouteLeaveMessage`

## 今天完成后的标准
- 路由器不再只处理两个消息
- 你已经为 Week 04 的心跳和清理打好了接口基础

## 今天会用到的库和最小代码

### 扩展消息类型
```go
type Message struct {
    Type    string `json:"type"`
    RoomID  string `json:"room_id"`
    UserID  string `json:"user_id"`
    Token   string `json:"token"`
    Content string `json:"content"`
}
```

### 路由增加 `heartbeat`
```go
case "heartbeat":
    return handleHeartbeat(msg, sess)
```
