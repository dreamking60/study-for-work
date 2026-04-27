# Week 06 教程：启动 `room-orchestrator`，先做状态机

## 本周目标
- 把房间生命周期压缩成可验证状态机
- 实现最小 manager / scheduler 骨架
- 明确和 `edge-gateway` 的边界

## 本周新仓库位置
```bash
mkdir -p ../../../go-backend-learning/room-orchestrator
cd ../../../go-backend-learning/room-orchestrator
go mod init room-orchestrator
```

## 本周建议目录
```text
cmd/orchestrator/main.go
internal/room/model.go
internal/room/state_machine.go
internal/event/event.go
internal/manager/manager.go
internal/scheduler/scheduler.go
```

## 开始前先读
- [foundation-reference.md](../foundation-reference.md)
- [concepts-and-libraries.md](./concepts-and-libraries.md)

## 第 1 天：只选一个业务对象
今天只做 `room`，不要做所有玩法。

建议状态：
- `created`
- `waiting`
- `running`
- `finished`
- `failed`

完成标准：
- `Room` 结构和状态常量已经落地

## 第 2 天：写合法转移表
先写规则，再写代码。

你必须明确：
- 哪些状态可以转
- 哪些绝对不能转

建议你今天直接写成测试表，而不是只写文字。

## 第 3 天：状态机落代码
至少实现：
- `CanTransit`
- `ApplyTransition`

完成标准：
- 非法转移会报错

建议至少有：
- `TestCanTransit`
- `TestApplyTransitionRejectsInvalid`

## 第 4 天：记录事件
每次状态变化都要记事件，不要只改字段。

至少有：
- `room_id`
- `from`
- `to`
- `reason`

建议事件先保存在内存 slice，后面再谈持久化。

## 第 5 天：搭 manager / scheduler 骨架
最小职责：
- manager 管房间集合
- scheduler 周期扫描并推进状态

今天的停止点：
- manager 能创建房间
- scheduler 至少能扫描一遍房间列表

## 第 6 天：定义和 `edge-gateway` 的边界
写清：
- `edge-gateway` 负责接入
- `room-orchestrator` 负责状态流转

## 第 7 天：补测试和文档
至少要有：
- 正常转移测试
- 非法转移测试
- 状态图说明

## 本周最终自检命令
```bash
go test ./...
go run ./cmd/orchestrator
```
