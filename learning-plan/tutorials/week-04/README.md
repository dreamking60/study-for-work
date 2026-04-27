# Week 04 教程：做 `edge-gateway` 的稳定性第一轮

## 本周目标
- 补连接生命周期管理
- 加心跳和超时回收
- 修一轮并发安全问题

## 开始前先读
- [foundation-reference.md](../foundation-reference.md)
- [concepts-and-libraries.md](./concepts-and-libraries.md)

## 本周主要改哪些文件
- `internal/server/server.go`
- `internal/session/session.go`
- `internal/room/service.go`
- 如需要新增 `internal/heartbeat/`

## 本周通过线
- 有心跳或保活机制
- 有超时回收
- 有至少一个并发相关测试
- 有一页架构说明

## 第 1 天：画连接状态图
先画清楚：
- 建立连接
- 加入房间
- 保活
- 超时
- 断开
- 清理

完成标准：
- 你能讲清每个状态切换点在哪里

建议你把这张状态图直接写进本周报告，不要只在脑子里想。

## 第 2 天：加心跳
无论你用什么协议，今天至少有一个保活动作。

最小做法：
- 连接收到 `heartbeat`
- 更新 `LastSeenAt`
- 返回 `pong` 或简单 ack

完成标准：
- 服务端能识别活连接和失活连接

建议你把 `LastSeenAt` 明确放到 `Session` 上。

## 第 3 天：做超时回收
你要实现：
- 定时扫描
- 失活连接清理
- 必要时从房间成员中移除

完成标准：
- 失活连接不会永久残留

建议你新增一个定时清理函数，例如：
- `CleanupExpiredSessions(now time.Time)`

## 第 4 天：梳理并发风险
把这些地方全部过一遍：
- 成员 map 读写
- 会话表读写
- 广播时遍历和删除冲突

完成标准：
- 明确哪些地方要上锁或重构

今天先不要急着全修。
先写出风险清单，再决定最重要的一个点。

## 第 5 天：修并发安全问题
今天只修最关键的一轮。

完成标准：
- 并发测试不再随机失败

建议至少写一个测试：
- 并发 join
- 并发 disconnect
- 广播与 disconnect 并发

三选一，但必须有真实并发场景。

## 第 6 天：整理错误模型
给接入层定义固定错误类型：
- 鉴权失败
- 参数错误
- 房间不存在
- 会话失效

完成标准：
- 错误输出有分类，不是散乱字符串

建议你同步统一日志字段，至少保留：
- `conn_id`
- `user_id`
- `room_id`
- `event`

## 第 7 天：产出架构文档
只写 3 件事：
- 模块边界
- 连接生命周期
- 并发安全策略

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
