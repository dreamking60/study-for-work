# Week 04 Day 2：加最小 heartbeat，让连接状态开始可信

## 今天你要完成什么
今天先做最小 heartbeat，不求复杂，但要成立。

## 你今天的最小实现
1. 客户端或测试输入发来 `heartbeat`
2. 服务端更新 session 的 `LastSeenAt`
3. 返回 `pong` 或简单 ack

## 为什么要把 `LastSeenAt` 放到 Session
因为“连接最近一次活跃时间”本质上属于连接上下文，而不是房间本身。

## 今天建议改哪些地方
- `internal/session/session.go`
- `internal/server/router.go`

## 今天完成后的标准
- heartbeat 已有路由
- Session 上有最近活跃时间
- 后面做超时回收有了基础

## 今天会用到的库和最小代码

### 更新心跳时间
```go
func handleHeartbeat(sess *session.Session) {
    sess.LastSeenAt = time.Now()
}
```

### 最小 ack
```go
fmt.Println("pong")
```
