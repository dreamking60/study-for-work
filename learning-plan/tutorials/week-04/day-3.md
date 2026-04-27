# Week 04 Day 3：做超时回收，不让死连接永久留着

## 今天你要完成什么
今天做一个最关键的稳定性动作：
- 长时间没 heartbeat 的连接要被清掉

## 为什么这一步现在必须做
因为只要有了房间成员表，没有超时回收，房间状态迟早会失真。

## 你今天的最小做法
1. 周期性扫描 session
2. 判断 `now - LastSeenAt` 是否超过阈值
3. 超过就清理 session
4. 从房间成员中移除

## 今天建议新增的方法
- `CleanupExpiredSessions(now time.Time)`

## 今天完成后的标准
- 失活连接不会永久残留
- 连接清理和房间清理是一条完整链路

## 今天会用到的库和最小代码

### `time.NewTicker`
```go
ticker := time.NewTicker(5 * time.Second)
defer ticker.Stop()
for range ticker.C {
    cleanupExpiredSessions(time.Now())
}
```

### 超时判断
```go
expired := now.Sub(sess.LastSeenAt) > 30*time.Second
```
