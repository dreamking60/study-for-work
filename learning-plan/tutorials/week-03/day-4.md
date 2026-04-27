# Week 03 Day 4：把共享状态从 handler 里抽出来

## 今天你要完成什么
今天最重要的不是加新业务，而是开始整理结构。

到现在为止，你已经很可能在这些地方保存状态：
- 房间成员
- session 列表
- 用户和连接的对应关系

如果这些状态还散在 handler 或 router 里，今天就该收了。

## 今天你要抽的两层
- `session store`
- `room store`

## 为什么今天必须抽 store
因为从今天开始，功能会越来越多：
- 鉴权
- 心跳
- leave
- 断开清理
- 超时回收

如果所有逻辑都直接改 map，Week 04 会非常难收口。

## 今天建议新增哪些文件
- `internal/store/session_store.go`
- `internal/store/room_store.go`

如果你不想新建目录，也至少要把 store 概念抽成单独类型。

## 今天完成后的标准
- handler 不再直接改裸 map
- 至少有一层统一的状态访问入口
- 后面做清理和超时时不需要到处找变量

## 今天会用到的库和最小代码

### 最小 `SessionStore`
```go
type SessionStore struct {
    sessions map[string]*session.Session
}

func (s *SessionStore) Save(sess *session.Session) {
    s.sessions[sess.ConnID] = sess
}

func (s *SessionStore) Get(connID string) (*session.Session, bool) {
    sess, ok := s.sessions[connID]
    return sess, ok
}
```

### 为什么先接受这个版本
因为你今天的重点是“抽象边界”，并发安全放到 Week 04 再补。
