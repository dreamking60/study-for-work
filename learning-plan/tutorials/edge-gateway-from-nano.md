# 教程 02：基于 `nano` 思路做 `edge-gateway` 最小版本

## 参考资料
- `lonng/nano` 入门文档
  - https://github.com/lonng/nano/blob/master/docs/get_started.md
- `lonng/nano` 设计模式文档
  - https://github.com/lonng/nano/blob/master/docs/design_patterns.md

## 为什么这份教程适合你
`nano` 的入门文档本质上就是一个最小聊天室服务，它把几个最关键的东西讲清楚了：
- `component`
- `handler`
- `route`
- `session`
- `group`

这些概念不是拿来背的，你要把它们翻译成你自己的 `edge-gateway` 骨架。

文档里的核心样例就是一个 `Room` 组件：
- `AfterInit()` 里注册会话关闭回调
- `Join()` 里绑定用户、广播新用户、加入组
- `Message()` 里广播消息

你的任务不是照搬 `nano` 框架，而是照着这个职责拆分思路，自己做一个最小接入服务。

## 目标产物
- 一个独立目录，例如：
  - `go-backend-learning/edge-gateway`
- 支持一个最小连接入口
- 支持两个消息动作：
  - `join`
  - `chat`
- 支持会话退出清理
- 至少有 1 个连接生命周期测试

## 第一天怎么做

### 第 1 步：只搭骨架，不做复杂功能
先建目录：
```bash
mkdir -p ../../go-backend-learning/edge-gateway
cd ../../go-backend-learning/edge-gateway
go mod init edge-gateway
```

建议目录：
```text
cmd/server
internal/session
internal/room
internal/protocol
internal/server
```

## 第 2 步：先把 `nano` 的 5 个概念翻译成你自己的模块

按这个映射做：
- `component` -> 你的业务模块，例如 `room.Service`
- `handler` -> 你的消息处理函数，例如 `HandleJoin` / `HandleChat`
- `route` -> 你的消息类型字段，例如 `type=join` / `type=chat`
- `session` -> 你的连接上下文，保存连接 ID、用户 ID、房间 ID
- `group` -> 你的房间成员集合

今天先不要上来写鉴权、限流、存储。

## 第 3 步：定义最小消息协议
建议先用 JSON，消息结构固定成这样：
```json
{
  "type": "join",
  "room_id": "room-1",
  "user_id": "u-1001",
  "content": "hello"
}
```

你至少要定义：
- `join`
- `chat`

你需要有一个统一入口：
- 先解析消息
- 根据 `type` 分发到不同 handler

## 第 4 步：先做最小房间服务
参考 `nano` 样例里 `Join()` 和 `Message()` 的职责：

你要实现的行为：
- 用户发送 `join`
  - 把 session 放进房间成员集合
  - 返回当前房间成员数
- 用户发送 `chat`
  - 广播给当前房间的所有在线连接
- 连接关闭
  - 从房间成员集合中删除

## 建议你直接写的几个类型
```go
type Message struct {
    Type    string `json:"type"`
    RoomID  string `json:"room_id"`
    UserID  string `json:"user_id"`
    Content string `json:"content"`
}

type Session struct {
    ConnID string
    UserID string
    RoomID string
}
```

```go
type Room struct {
    ID      string
    Members map[string]*Session
}
```

## 今天的停止点
只要你完成下面 3 件事就停：
- 能启动一个最小服务
- 能处理 `join`
- 能处理 `chat`

不要在第一天加这些：
- Redis
- MySQL
- JWT
- 指标系统
- 多房间调度

## 今天的验证方式
你至少要有：
```bash
go test ./...
go run ./cmd/server
```

如果你先做的是内存版假连接，也可以接受，前提是：
- 路由逻辑真实存在
- 房间成员管理真实存在
- 连接关闭清理真实存在

## 第二天继续做什么
等第一天的骨架跑通，再继续补：
- 心跳
- 超时清理
- 结构化日志
- 简化版鉴权

这部分对应 `nano` 文档里对 session 和 group 的使用方式，但实现要保持你自己的代码结构。
