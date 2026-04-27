# WebSocket Protocol Draft

## 设计目标
- 简单
- 可扩展
- 适合终端 client 的事件驱动模型

## 基础消息格式

```json
{
  "id": "evt_123",
  "type": "message.send",
  "workspace_id": "w1",
  "channel_id": "c1",
  "user_id": "u1",
  "ts": 1710000000,
  "body": {}
}
```

## Client -> Server

### `auth.login`
- token 登录

### `channel.subscribe`
- 订阅频道

### `channel.unsubscribe`
- 取消订阅频道

### `message.send`
- 发送消息

### `message.edit`
- 编辑消息

### `thread.reply`
- 回复线程

### `reaction.add`
- 添加 reaction

### `command.exec`
- 执行 slash command 或 `:` 命令

## Server -> Client

### `auth.ok`
- 登录成功

### `auth.failed`
- 登录失败

### `message.created`
- 新消息

### `message.updated`
- 消息更新

### `thread.updated`
- 线程变化

### `presence.changed`
- 在线状态变化

### `channel.unread`
- 未读计数变化

### `bot.reply`
- bot 执行结果

### `error`
- 协议或业务错误

## ACK 机制
- client 发送带 `id`
- server 返回 `ack` 或具体事件结果
- 超时可重试，但要带幂等 key

## 错误码建议
- `unauthorized`
- `forbidden`
- `invalid_payload`
- `channel_not_found`
- `message_not_found`
- `rate_limited`
- `internal_error`
