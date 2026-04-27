# Week 04 Day 6：统一错误模型和日志字段，为后面压测做准备

## 今天你要完成什么
今天你要把稳定性相关的可观测性基础先垫起来。

## 为什么现在就要做这件事
因为你下一周就要开始做：
- 限流
- 压测
- 热点定位

如果日志和错误模型现在不统一，下一周基本上什么都不好看。

## 今天至少统一这两层

### 错误
- unauthorized
- invalid_message
- room_not_found
- session_expired

### 日志字段
- conn_id
- user_id
- room_id
- event

## 今天完成后的标准
- 稳定性相关行为不是“默默发生”
- 你能从日志里看见连接状态变化

## 今天会用到的库和最小代码

### 统一事件日志
```go
fmt.Printf("event=%s conn_id=%s user_id=%s room_id=%s\n", event, sess.ConnID, sess.UserID, sess.RoomID)
```

### 统一错误返回
```go
return fmt.Errorf("session_expired")
```
