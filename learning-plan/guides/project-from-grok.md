# 候选项目方向整理

这份文档保留 Grok 提供的线索，但改成可直接用于立项的版本。

## 结论先说
不是把这些仓库全部学一遍，而是把它们分成三类用途：
- 主项目参考：决定我们自己要做什么。
- 实现借鉴：遇到具体问题时参考其模块设计。
- 教程素材：给后续 tutorial 生成提供高质量上下文。

## 候选池 A：游戏服务器 / 实时链路

### `lonng/nano`
- 价值：非常适合作为实时接入、组件拆分、Session 管理、集群模式的参考。
- 适合借鉴的部分：
  - Handler 组织方式。
  - Session / route / component 的拆分思路。
  - 从聊天 demo 过渡到房间服务的最小路径。
- 不直接照搬的原因：
  - 你现在更需要“自己实现可讲清楚的核心链路”，而不是把框架用熟。

### `heroiclabs/nakama`
- 价值：适合作为生产级游戏后端的能力清单参考。
- 适合借鉴的部分：
  - 运行时模块扩展方式。
  - 房间、匹配、排行榜、社交能力如何拼成完整产品。
  - Docker 化和服务化交付思路。
- 不直接作为首个项目主仓库的原因：
  - 体量大，容易变成“看懂一点点但讲不清自己做了什么”。

### `grpc-game-example` / `cgo-game-server`
- 价值：补充实时通信和压测思路。
- 用法：当你做到 `edge-gateway` 的协议层和压测层时，再拿来参考，不作为主教程骨架。

## 候选池 B：编排 / 调度 / 控制面

### `Build an Orchestrator in Go`
- 价值：最适合拿来理解“为什么需要 scheduler / worker / manager”。
- 适合借鉴的部分：
  - 控制面与执行面的职责边界。
  - 调度循环、任务状态、资源视角。
- 用法：
  - 用它来指导 `room-orchestrator` 和后续 `cluster-operator` 的设计，而不是全文照抄。

### `Kubebuilder`
- 价值：做云平台方向项目时最重要的官方实践入口。
- 适合借鉴的部分：
  - CRD + Controller 的项目骨架。
  - Reconcile 模型。
  - 本地 kind/minikube 验证流程。
- 用法：
  - 在 Week 08 以后作为 `cluster-operator` 的直接实现基线。

### Go + Kubernetes CI/CD 示例
- 价值：适合作为发布链路、镜像构建、部署脚本、可观测性的补充案例。
- 用法：
  - 放到 Week 09 的交付阶段，用于补齐 GitHub Actions / Docker / K8s 部署流程。

## 本轮最终采用的项目路线

### 项目 0：`opsctl`
- 来源：保留你当前 `week01-cli` 工作。
- 目标：工程 CLI，不再是孤立练习。
- 覆盖能力：配置、命令分发、测试、后续数据导入导出与运维辅助。

### 项目 1：`edge-gateway`
- 参考优先级：`nano` > `grpc-game-example`
- 目标：实现一个可讲清楚的实时接入网关，至少覆盖连接管理、会话校验、消息路由、限流、日志与压测。

### 项目 2：`room-orchestrator`
- 参考优先级：`Build an Orchestrator in Go` > `Nakama` 的房间/匹配思路
- 目标：实现房间创建、状态机、异步任务、重试补偿、事件记录。

### 项目 3：`cluster-operator`
- 参考优先级：`Kubebuilder` > Go + Kubernetes DevOps 示例
- 目标：实现本地集群中的房间服务部署控制、环境配置、发布与观测。

## 这份候选池怎么参与 tutorial 生成
- tutorial 不是“从 Go 基础重新讲起”，而是按当前周目标，选一两个最相关参考仓库来喂上下文。
- 每次只围绕一个问题生成 tutorial：
  - 例如“如何设计 `edge-gateway` 的连接管理”。
  - 例如“如何用 Kubebuilder 写一个最小 Reconcile 循环”。
- 不再用“大而全 prompt”一次性生成两周甚至十周内容。
