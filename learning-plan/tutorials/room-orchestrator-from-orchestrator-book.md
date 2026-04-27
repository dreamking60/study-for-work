# 教程 03：基于 Orchestrator 书的思路做 `room-orchestrator`

## 参考资料
- 仓库：
  - https://github.com/buildorchestratoringo/code
- 书籍主页：
  - https://www.manning.com/books/build-an-orchestrator-in-go-from-scratch
- 辅助参考：
  - https://heroiclabs.com/docs/nakama/server-framework/go-runtime/code-samples/

## 这份教程的目标
不是让你去复刻一个 Kubernetes。

你现在真正要学的是：
- 为什么要有 `manager`
- 为什么要有 `scheduler`
- 为什么要有 `worker`
- 为什么状态机和调度循环是同一类问题

然后把这些能力翻译成你的游戏业务场景：
- 房间创建
- 玩家加入
- 房间开始
- 房间结束
- 超时回收

## 你要做出来的最小版本
- 一个新目录，例如：
  - `go-backend-learning/room-orchestrator`
- 一个房间状态机
- 一个简单调度循环
- 一个事件记录或任务日志
- 至少 2 个异常场景测试

## 第一天只做状态机

### 第 1 步：先定一个最小业务对象
今天不要做“所有玩法”。

只选一个对象：
- `room`

房间状态先限制成：
- `created`
- `waiting`
- `running`
- `finished`
- `failed`

### 第 2 步：写出合法转移
你先别写代码，先把这张表写出来：

```text
created -> waiting
waiting -> running
running -> finished
running -> failed
waiting -> failed
```

明确禁止的转移：
- `finished -> running`
- `failed -> running`

### 第 3 步：把状态机落成代码
建议你至少有：
```go
type RoomStatus string

const (
    RoomCreated  RoomStatus = "created"
    RoomWaiting  RoomStatus = "waiting"
    RoomRunning  RoomStatus = "running"
    RoomFinished RoomStatus = "finished"
    RoomFailed   RoomStatus = "failed"
)
```

再写：
```go
func CanTransit(from, to RoomStatus) bool
func ApplyTransition(room *Room, to RoomStatus) error
```

### 第 4 步：把状态变化记录下来
不要只改一个字段。

至少要记一条事件：
```go
type RoomEvent struct {
    RoomID string
    From   RoomStatus
    To     RoomStatus
    Reason string
}
```

## 第二天再做调度循环
当你把状态机写稳以后，再继续：
- 一个简单 manager
- 一个简单 scheduler
- 一个超时回收逻辑

最小职责可以这样拆：
- `manager`
  - 持有房间集合
  - 接收创建请求
- `scheduler`
  - 周期性扫描房间
  - 推进满足条件的房间状态

## 你应该从参考资料里学什么
从 Orchestrator 资料里学：
- 控制循环
- 状态与资源视角
- manager / worker / scheduler 边界

从 Nakama 代码样例里学：
- 房间或 match 生命周期其实就是一组有节奏的状态变化

## 今天的完成标准
- 房间状态机代码已存在
- 非法状态转移会报错
- 有事件记录
- 有测试覆盖正常和非法转移

## 今天不要做什么
- 不要接数据库
- 不要做分布式
- 不要写复杂任务队列
- 不要急着上 K8s

先把“房间生命周期就是一个可验证状态机”这件事做扎实。
