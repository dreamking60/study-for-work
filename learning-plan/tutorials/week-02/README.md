# Week 02 教程：做 `edge-gateway` 的最小实时接入服务

## 本周目标
- 基于 `nano` 的聊天室思路，自己做一个最小接入服务。
- 至少支持 `join` 和 `chat` 两类消息。
- 会话退出时能正确清理。

## 本周新仓库位置
```bash
mkdir -p ../../../go-backend-learning/edge-gateway
cd ../../../go-backend-learning/edge-gateway
go mod init edge-gateway
```

## 本周建议目录
```text
cmd/server/main.go
internal/protocol/message.go
internal/session/session.go
internal/room/room.go
internal/room/service.go
internal/server/router.go
internal/server/server.go
```

## 开始前先读
- [foundation-reference.md](../foundation-reference.md)
- [concepts-and-libraries.md](./concepts-and-libraries.md)

## 本周先不要做什么
- 不要加数据库
- 不要做 JWT
- 不要做 K8s
- 不要把 `edge-gateway` 和 `room-orchestrator` 写在一个仓库里

## 参考资料
- `nano` 入门文档
- `nano` 设计模式文档

## 第 1 天：建骨架
今天把这些文件先建出来，就算内容很少也可以：
- `cmd/server/main.go`
- `internal/protocol/message.go`
- `internal/session/session.go`
- `internal/room/room.go`
- `internal/server/router.go`

完成标准：
- 目录创建完成
- `go test ./...` 至少能跑空包
- `go run ./cmd/server` 至少能启动一个空 server

建议你在 `main.go` 里先只做一件事：
- 创建 server
- 打一行启动日志
- 阻塞运行

## 第 2 天：定义消息协议
先别追求完美，先用 JSON：
```json
{
  "type": "join",
  "room_id": "room-1",
  "user_id": "u-1",
  "content": "hello"
}
```

你要定义：
- `Message`
- `Session`
- `Room`

建议你直接落到这些文件：
- `internal/protocol/message.go`
- `internal/session/session.go`
- `internal/room/room.go`

完成标准：
- 协议结构体落地到代码
- 有基础解析测试

今天建议你最少写 2 个测试：
- `TestDecodeJoinMessage`
- `TestDecodeUnknownMessageType`

## 第 3 天：写统一入口和路由分发
你要做一个统一入口：
- 收到消息
- 解析消息
- 根据 `type` 分发

至少支持：
- `join`
- `chat`

建议你新增：
- `internal/server/router.go`
- `func Route(msg protocol.Message, sess *session.Session) error`

完成标准：
- 路由逻辑存在
- 未知消息类型会返回明确错误

今天的停止点：
- 可以把一条 JSON 消息解到路由器
- `join` 和 `chat` 至少能进入对应 handler

## 第 4 天：实现房间成员管理
今天重点做 `join`。

行为要求：
- 用户加入房间后，session 被记录
- 房间成员数可返回
- 同一个连接退出时可清理

建议增加一个 `room.Service`，职责只有：
- 根据 `room_id` 获取或创建房间
- 向房间注册 session
- 提供成员快照或人数

完成标准：
- `join` 跑通
- 成员集合准确

建议测试：
- `TestJoinCreatesRoom`
- `TestJoinAddsMember`

## 第 5 天：实现消息广播
今天做 `chat`。

行为要求：
- 消息发给当前房间所有在线成员
- 至少有一个内存版广播逻辑

如果你今天还没有真实网络连接层，也没关系。
最小做法：
- 先让房间服务返回“应该广播给哪些 session”
- 后续再把它接到真正连接写回

完成标准：
- `chat` 跑通
- 广播逻辑和房间成员逻辑解耦

建议测试：
- `TestBroadcastTargetsMembersInSameRoom`
- `TestBroadcastDoesNotLeakToOtherRooms`

## 第 6 天：做连接关闭清理和测试
今天专门做生命周期收口。

你要验证：
- 连接退出后是否从房间里删除
- 房间为空后是否要回收

完成标准：
- 至少有 1 个连接生命周期测试

建议你新增一个方法：
- `func (s *Service) Disconnect(connID string)`

这样后面补心跳和超时会简单很多。

## 第 7 天：写一份阶段总结
总结只写：
- 借鉴了 `nano` 的哪些思想
- 自己的代码结构怎么对应 `component/handler/session/group`
- 这一版故意没做什么

## 本周不要做什么
- 不要一开始加 Redis/MySQL
- 不要先做 JWT
- 不要同时做压测和鉴权

先把最小接入链路做通。

## 本周最终自检命令
```bash
go test ./...
go run ./cmd/server
```

## 按天执行
- [day-1.md](./day-1.md)
- [day-2.md](./day-2.md)
- [day-3.md](./day-3.md)
- [day-4.md](./day-4.md)
- [day-5.md](./day-5.md)
- [day-6.md](./day-6.md)
- [day-7.md](./day-7.md)
