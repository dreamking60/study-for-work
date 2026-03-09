# Week 02：并发与网络编程（详细执行版）

## 本周目标
- 掌握 goroutine/channel/context 的实际使用。
- 实现一个 WebSocket 聊天室服务 v1。
- 完成首次压测并输出指标。

## Day 1
- 学习：goroutine 生命周期、channel 阻塞语义。
- 编码：写 worker pool 示例。
- 验收：并发任务正确收敛，无死锁。
- 产出：`examples/worker_pool.go`。

## Day 2
- 学习：`context` 取消、超时、传值边界。
- 编码：把 Day1 示例改成支持超时取消。
- 验收：超时后 goroutine 能退出。
- 产出：`examples/context_cancel.go`。

## Day 3
- 学习：TCP 拆包粘包与协议边界。
- 编码：实现长度头协议解码器。
- 验收：连续发送多条消息不串包。
- 产出：`internal/protocol`。

## Day 4
- 学习：HTTP 中间件、请求生命周期。
- 编码：实现 healthz、request logging。
- 验收：中间件顺序正确，日志有请求耗时。
- 产出：`cmd/http-demo`。

## Day 5
- 学习：WebSocket 连接管理、心跳。
- 编码：实现连接注册/注销和广播。
- 验收：多客户端可收发，离线自动清理。
- 产出：聊天室核心逻辑初版。

## Day 6
- 学习：并发安全（锁/原子/chan）选型。
- 编码：聊天室 v1 收口，补单测。
- 验收：`go test ./...`，关键逻辑可测。
- 产出：`README` 更新聊天室架构图。

## Day 7
- 学习：压测工具使用（vegeta/wrk 二选一）。
- 编码：写压测脚本并跑 2 组数据。
- 验收：输出 QPS、P95/P99、错误率。
- 产出：`reports/week2-loadtest.md`。

## 本周完成标准
1. 聊天室支持连接、广播、心跳。
2. 有一份压测报告。
3. 有至少 2 个并发相关测试用例。

## 常见卡点
1. goroutine 泄漏：务必使用 context + defer close。
2. map 并发写 panic：加锁或使用单线程事件循环。
