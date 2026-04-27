# Week 03 教程：给 `edge-gateway` 补鉴权、路由和状态存储

## 本周目标
- 从“能连接”升级到“能处理业务消息”
- 补最小鉴权或会话校验
- 做至少两类业务消息路由

## 开始前先读
- [foundation-reference.md](../foundation-reference.md)
- [concepts-and-libraries.md](./concepts-and-libraries.md)

## 本周主要改哪些文件
- `internal/server/router.go`
- `internal/server/server.go`
- `internal/session/session.go`
- 新增 `internal/auth/`
- 如需要新增 `internal/store/`

## 第 1 天：设计最小会话校验方案
今天先定方案，不要先乱写。

你要回答：
- 用户怎么拿到身份
- 服务端如何校验
- 会话过期如何表示

这一版可以很简单：
- 静态 token
- 内存 session

建议你今天明确这 3 个字段：
- `user_id`
- `token`
- `expires_at`

完成标准：
- 方案写到文档或注释里

今天的停止点：
- 你能说清 token 在哪里校验
- 你能说清 session 在哪里保存

## 第 2 天：接入鉴权前置校验
今天把校验接到消息入口前。

你要保证：
- 未鉴权消息不会进入业务 handler
- 错误返回稳定可识别

建议新增：
- `internal/auth/auth.go`
- `func ValidateToken(token string) error`

完成标准：
- 至少一条未鉴权失败测试

建议命令：
```bash
go test ./...
```

## 第 3 天：增加业务消息类型
建议补这两类之一：
- `heartbeat`
- `leave`
- `room-info`

建议优先级：
1. `heartbeat`
2. `leave`
3. `room-info`

完成标准：
- 至少有 2 到 3 类消息分发

如果你时间不够，本周至少做到：
- `join`
- `message`
- `heartbeat`

## 第 4 天：抽出状态存储层
你现在至少要有一层抽象：
- 内存 session store
- 内存 room store

如果要更进一步，再考虑文件版或 Redis 版。

建议文件：
- `internal/store/session_store.go`
- `internal/store/room_store.go`

完成标准：
- handler 不直接操作裸 map

今天的停止点：
- handler 只依赖 store 或 service
- 业务逻辑不直接改全局变量

## 第 5 天：整理错误和日志字段
补这些字段：
- `conn_id`
- `user_id`
- `room_id`
- `message_type`

完成标准：
- 出错时能快速定位是哪个连接、哪个房间

建议同时统一错误字符串或错误码：
- `unauthorized`
- `invalid_message`
- `room_not_found`
- `session_not_found`

## 第 6 天：做联调验证
准备几个输入样例：
- 正常 join
- 未鉴权 join
- 正常 message
- 离开后 message

完成标准：
- 关键路径能回归

今天建议准备一个最小联调清单：
```text
1. join with valid token
2. join with invalid token
3. heartbeat after join
4. message after leave
```

## 第 7 天：输出 v1 小结
写清：
- 现在 `edge-gateway` 已覆盖哪些岗位能力
- 下一周要开始补哪些稳定性能力

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
