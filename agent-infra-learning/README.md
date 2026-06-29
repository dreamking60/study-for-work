# Agent Infra 学习计划（项目驱动版）

## 先看什么
1. 先读 `guides/job-description.md`，确认目标岗位真正要你证明的能力。
2. 再读 `projects/project-overview.md`，确认本轮采用的项目主线。
3. 直接从 `tutorials/README.md` 开始，`tutorials/` 是正文。
4. 每周结束按 `guides/weekly-checklist.md` 和 `guides/execution-standard.md` 产出交付物。

## 目录说明
- `guides/`: 岗位要求整理、执行规范、每周检查清单。
- `projects/`: 最终立项结果、项目拆分、简历话术。
- `tutorials/`: 直接照着做的实操教程，按周组织。
- `report/`: 已完成的日更记录，作为阶段成果存档。
- `templates/`: 架构图模板、设计文档模板、压测报告模板。

## 项目主线
- **项目 1：Agent 运行时** — ReAct 循环、工具调用、记忆系统、流式输出
- **项目 2：多 Agent 协作运行时** — 消息队列通信、监督者-执行者模式、状态检查点
- **项目 3：安全工具网关与沙箱** — 插件注册中心、Docker 隔离执行、鉴权限流
- **项目 4：可观测性与成本控制** — OpenTelemetry 埋点、Grafana 面板、成本熔断

## 核心目标
- 从零构建一个可演示的自研 Agent 基础设施
- 覆盖 Agent Infra 工程师面试中的系统设计与项目深挖
- 每个阶段必须落到代码、测试、设计文档、压测或演练记录
