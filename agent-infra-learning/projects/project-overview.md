# 项目概览

四个项目递进式实现，覆盖 Agent Infra 工程师的核心能力。

---

## 项目 1：Agent 运行时（Phase 1：第 9 周）

**目标**：从零实现一个单 Agent 运行时，支持 ReAct 循环、工具调用、记忆与流式输出。

**技术栈**：Python asyncio + FastAPI + SSE + Chroma

**覆盖能力**：
- Function Calling 协议理解
- ReAct 循环自实现
- 记忆系统抽象
- 流式输出

**可交付物**：
- 可运行的 Agent 服务（HTTP API）
- 支持注册自定义工具
- 滑动窗口 + 向量检索记忆
- SSE 实时输出 Agent 思考过程

---

## 项目 2：多 Agent 协作运行时（Phase 2：第 10 周）

**目标**：基于消息队列实现多 Agent 通信，支持监督者-执行者模式。

**技术栈**：Redis Pub/Sub 或 NATS + gRPC + 状态检查点

**覆盖能力**：
- 分布式通信
- 状态持久化与恢复
- 任务编排

**可交付物**：
- 多 Agent 协作 Demo
- 暂停/恢复流程

---

## 项目 3：安全工具网关与沙箱（Phase 2：第 11 周）

**目标**：设计插件化工具注册中心 + Docker 沙箱执行环境。

**技术栈**：Docker SDK for Python + 鉴权中间件

**覆盖能力**：
- 沙箱安全
- 鉴权与配额
- 插件架构设计

**可交付物**：
- 工具注册中心 API
- Docker 隔离执行 Demo
- 权限控制

---

## 项目 4：可观测性与成本控制（Phase 2：第 12 周）

**目标**：全链路可观测 + Token 成本熔断。

**技术栈**：OpenTelemetry + Prometheus + Grafana

**覆盖能力**：
- 链路追踪
- 指标监控
- 成本控制策略

**可交付物**：
- Agent 全链路 Trace
- Token 用量 Dashboard
- 成本熔断 Demo

---

## 简历话术要点
- "从零设计并实现了 Agent 运行时，支持 ReAct 循环、工具注册、向量记忆与流式输出"
- "基于消息队列构建多 Agent 协作框架，实现监督者-执行者模式与状态检查点恢复"
- "设计安全沙箱工具网关，支持 Docker 隔离执行、鉴权与配额管理"
- "为 Agent 系统搭建 OpenTelemetry 全链路观测 + Prometheus/Grafana 监控面板"
