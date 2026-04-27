# Week 02 Day 2：定义消息协议和最小数据结构

## 今天你要完成什么
今天是“先定义语言，再写动作”。

如果消息协议不先定下来，后面所有路由、房间管理、广播逻辑都会乱。

## 今天先定一个最小 JSON 协议
今天先用 JSON，不要先做二进制协议。

最小消息长这样：
```json
{
  "type": "chat",
  "room_id": "room-1",
  "user_id": "u-1",
  "content": "hello"
}
```

## 你今天要写的 3 个核心结构

### 1. Message
放在：
- `internal/protocol/message.go`

### 2. Session
放在：
- `internal/session/session.go`

### 3. Room
放在：
- `internal/room/room.go`

如果你不知道 JSON 结构体标签怎么写，先看：
- [library-usage-cheatsheet.md](../library-usage-cheatsheet.md)
  重点看 `encoding/json`

## 为什么今天不写太多字段
因为你现在要的是“能跑通一个最小闭环”，不是做一个终极协议。

这一版字段只要够支撑：
- 加入房间
- w

就够了。

## 今天至少写两个测试
- `TestDecodeJoinMessage`
- `TestDecodeUnknownMessageType`

这样做的意义是：
- 你能验证协议输入是稳定的
- 你能提前定义异常输入的行为

## 今天完成后的标准
- 协议结构体已存在
- 会话和房间结构体已存在
- 至少 2 条协议解析测试已存在

## 今天会用到的库和最小代码

### `encoding/json` 结构体标签
```go
type Message struct {
    Type    string `json:"type"`
    RoomID  string `json:"room_id"`
    UserID  string `json:"user_id"`
    Content string `json:"content"`
}
```

### `json.Unmarshal`
把一条 JSON 消息解出来：
```go
var msg Message
if err := json.Unmarshal(raw, &msg); err != nil {
    return err
}
```

### `testing`
最小解析测试：
```go
func TestDecodeJoinMessage(t *testing.T) {
    raw := []byte(`{"type":"join","room_id":"room-1","user_id":"u-1"}`)
    var msg Message
    if err := json.Unmarshal(raw, &msg); err != nil {
        t.Fatalf("decode message: %v", err)
    }
    if msg.Type != "join" {
        t.Fatalf("unexpected type: %q", msg.Type)
    }
}
```
