# Week 04 Day 5：修一轮最关键的并发安全问题

## 今天你要完成什么
今天不追求把所有并发问题一次解决，而是修最危险、最容易出现随机错误的一个点。

## 今天你可以优先选哪个点
三选一即可：
- 并发 join
- 并发 disconnect
- 广播和 disconnect 并发

## 为什么今天只修一个点
因为这周的重点是建立方法，不是堆修改量。

你要学会的是：
- 找风险
- 选优先级
- 改一轮
- 用测试证明它比之前稳

## 今天必须补的内容
- 至少一条真实并发测试

如果你不知道 `sync.Mutex` / `sync.RWMutex` 该怎么包，先看：
- [library-usage-cheatsheet.md](../library-usage-cheatsheet.md)
  重点看 `sync.Mutex` / `sync.RWMutex`

## 今天完成后的标准
- 并发测试存在
- 最关键的共享状态读写不再裸奔

## 今天会用到的库和最小代码

### `sync.RWMutex`
```go
type RoomStore struct {
    mu    sync.RWMutex
    rooms map[string]*Room
}

func (s *RoomStore) Get(id string) (*Room, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    room, ok := s.rooms[id]
    return room, ok
}

func (s *RoomStore) Save(room *Room) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.rooms[room.ID] = room
}
```

### 并发测试最小样子
```go
func TestConcurrentJoin(t *testing.T) {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // 调用 join
        }()
    }
    wg.Wait()
}
```
