# Agent Infra 工程师 — 岗位要求分析

## 岗位本质
为成千上万个智能体提供稳定、高效、安全的运行底座——"智能体时代的操作系统与云计算基础设施"。

## 典型工作内容
- 设计 Agent 运行时（循环推理、工具调用、记忆读写）
- 多 Agent 安全协同与工作流编排（DAG、状态机）
- 工具/插件生态（注册、发现、鉴权、沙箱执行）
- 大规模 Agent 集群管理（弹性扩缩、负载均衡、多租户隔离）
- 可观测性（全链路追踪、Token 成本审计、延迟监控）
- 行业协议兼容（MCP、A2A 等）

## 需要证明的能力

### 编程语言
- **Python**（主力，所有 AI 框架的基础）
- **Go 或 Rust** 之一（高性能 Infra 组件：网关、沙箱、消息中间件）

### 计算机基础
- 操作系统：进程/容器、资源隔离
- 计算机网络：HTTP/gRPC、WebSocket、NAT 穿透
- 并发编程：asyncio、线程模型、GIL 理解
- 数据结构与算法

### 分布式系统
- 容器与 K8s
- 消息队列（Kafka/NATS/RabbitMQ）
- RPC/gRPC、分布式锁、一致性、服务发现、微服务设计

### 数据库与存储
- 关系型/文档型：PostgreSQL、Redis、MongoDB
- 向量数据库：Milvus、Qdrant、Chroma
- 可选：图数据库（知识图谱记忆）

### 大模型基础
- Transformer 基本原理、Prompt Engineering、RAG
- Function Calling 机制
- 推理优化概念：vLLM、Continuous Batching

### Agent 理论
- ReAct、Plan-and-Execute、Reflexion、Tree of Thoughts
- 多 Agent 协作模式（辩论、角色分工、分层监督）

### 行业协议
- OpenAI Function Calling / Tool Use
- Anthropic MCP
- Google A2A

## 面试常考
- 系统设计："设计一个可扩展的 Agent 平台"
- 分布式理论：一致性与可用性权衡
- 并发编程：Python/Go 并发模型
- Agent 理论深度理解
