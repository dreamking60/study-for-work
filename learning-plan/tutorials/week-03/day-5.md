# Week 03 Day 5：统一错误和日志字段，让后面排障不崩

## 今天你要完成什么
今天是工程质量日。

你现在已经有了：
- 路由
- 房间
- 消息
- 初步鉴权

如果现在不统一错误和日志，后面会非常难排查。

## 今天你要统一两样东西

### 1. 错误分类
建议至少先有：
- `unauthorized`
- `invalid_message`
- `room_not_found`
- `session_not_found`

### 2. 日志字段
建议至少带：
- `conn_id`
- `user_id`
- `room_id`
- `message_type`

## 为什么这一步这么重要
因为后面你开始做：
- 心跳
- 超时
- 并发
- 压测

没有统一日志和错误分类，所有问题都会变成“看起来哪里都像有问题”。

## 今天建议改哪些地方
- `internal/server/router.go`
- `internal/server/server.go`
- 相关错误定义文件

## 今天完成后的标准
- 错误不是散乱字符串
- 日志不是只打印一句“failed”
- 一次消息处理过程至少能看见关键信息

## 今天会用到的库和最小代码

### 错误包装
```go
if err != nil {
    return fmt.Errorf("route %s: %w", msg.Type, err)
}
```

### 最小日志
```go
fmt.Printf("event=%s conn_id=%s user_id=%s room_id=%s\n", msg.Type, sess.ConnID, sess.UserID, sess.RoomID)
```
