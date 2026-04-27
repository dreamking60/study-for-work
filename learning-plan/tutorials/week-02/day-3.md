# Week 02 Day 3：写统一入口和最小路由器

## 今天你要完成什么
今天把“消息进来以后去哪”这件事定下来。

这一步的目的不是做完整网络服务，而是先把路由逻辑做出来。

## 今天你要写哪个文件
- `internal/server/router.go`

## 你今天要实现的思路
统一入口做三件事：
1. 收到消息
2. 解析消息
3. 根据 `type` 分发到不同 handler

## 今天你至少支持两个消息类型
- `join`
- `chat`

## 为什么先做路由，再做业务
因为如果没有一层统一分发，后面你很容易把：
- 协议解析
- 业务逻辑
- 房间管理

全部写在一起。

那样你后面再补鉴权和心跳会非常痛苦。

## 今天建议的停止点
你今天做到下面这些就可以停：
- 一条 JSON 消息能进入 router
- `join` 会走到 join handler
- `chat` 会走到 chat handler
- 未知 `type` 会返回明确错误

## 今天建议补的测试
- `TestRouteJoinMessage`
- `TestRouteUnknownMessage`

## 今天完成后的标准
- 路由器存在
- 两类消息可分发
- 未知类型有稳定错误行为

## 今天会用到的库和最小代码

### `switch`
最直接的消息分发写法：
```go
switch msg.Type {
case "join":
    return handleJoin(msg, sess)
case "chat":
    return handleChat(msg, sess)
default:
    return fmt.Errorf("unknown message type %q", msg.Type)
}
```

### 为什么这里先用 `switch`
因为你现在先要把分发边界讲清楚，而不是先设计一套复杂 handler 注册框架。
