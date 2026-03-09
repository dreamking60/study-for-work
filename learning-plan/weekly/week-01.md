# Week 01：Go 基础与工程化（概念优先版）

## 本周核心概念
- `module`：项目依赖边界与导入路径前缀。
- `package`：代码组织单元，按职责分目录。
- `config`：程序如何运行（环境、端口、超时）。
- `model`：业务对象是什么（如 `Task`）。

## 为什么重要
- 你后续所有服务都依赖这套工程骨架。
- 先区分“运行配置”和“业务模型”，后面才不会把逻辑堆在 `main`。

## 参考资料
- https://go.dev/doc/tutorial/getting-started
- https://go.dev/tour/basics/1
- https://go.dev/tour/moretypes/2
- https://go.dev/tour/methods/1
- https://go.dev/tour/methods/9
- https://12factor.net/config

## 本周实现目标
- 项目可运行、可测试、可扩展。
- 完成 `internal/config` 与 `internal/model` 的最小抽象。

## Day 1
- 概念：Go 项目最小闭环（写代码 -> 运行 -> 测试）。
- 实践目标：完成 `main` 入口与 smoke test。
- 验收标准：`go run` 与 `go test` 都通过。

## Day 2
- 概念：`config` 是运行参数，`model` 是业务对象。
- 实践目标：实现 `Config` 和 `Task`，在 `main` 调用。
- 验收标准：输出中同时出现配置信息和任务摘要。

## Day 3
- 概念：错误不是打印日志，而是明确返回与传播。
- 实践目标：配置加载与错误定位（字段级）。
- 验收标准：非法配置能返回可读错误。

## Day 4
- 概念：日志和参数解析是可运维性的基础。
- 实践目标：支持运行参数，输出结构化日志。
- 验收标准：参数错误提示清晰，日志字段稳定。

## Day 5
- 概念：测试是行为约束，不是收尾动作。
- 实践目标：为 config/model 添加单测。
- 验收标准：测试通过且有覆盖率记录。

## Day 6
- 概念：CLI 命令是业务能力的外部接口。
- 实践目标：实现 2 个子命令并保持可重复执行。
- 验收标准：同输入同输出，异常输入可控失败。

## Day 7
- 概念：工程成果必须可复现、可交接。
- 实践目标：整理 README 与周复盘。
- 验收标准：新人 10 分钟内跑通。
