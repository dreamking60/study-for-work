# 项目参考资料怎么使用

这份文档的重点不是“列链接”，而是规定每类资料在我们的 tutorial 和实现里扮演什么角色。

## 使用原则
- 参考仓库是“实现和教程的依据”，不是“要完整复刻的目标”。
- 每次 tutorial 只绑定一个当前问题和一到两个参考资料。
- 优先使用官方文档、项目内 README、设计文档，再看社区博客。
- 生成教程时必须说明：
  - 当前目标是什么。
  - 参考哪一份资料。
  - 只借鉴哪些模块，不照抄哪些部分。

## 基础资料与标准库入口
如果教程里提到某个库或概念，而你还不熟，先从这些稳定资料开始：

- Go 官方教程
  - https://go.dev/doc/tutorial/getting-started
- Go Tour
  - https://go.dev/tour/welcome/1
- `flag`
  - https://pkg.go.dev/flag
- `encoding/json`
  - https://pkg.go.dev/encoding/json
- `context`
  - https://pkg.go.dev/context
- `sync`
  - https://pkg.go.dev/sync
- `testing`
  - https://pkg.go.dev/testing
- `time`
  - https://pkg.go.dev/time
- `net/http`
  - https://pkg.go.dev/net/http

## 新增的写作与学习参考
下面这些链接不一定都直接指导某一周的代码实现，但它们对“教程怎么写得更完整、学习内容怎么铺开、面试和简历材料怎么组织”有参考价值：

- `tmnhs/go-interview-resume`
  - https://github.com/tmnhs/go-interview-resume
  - 用途：参考 Go 后端求职资料的组织方式，尤其是面试、简历和系统设计材料如何分区。
- `leezchuan` 的 Go 学习文章
  - https://leezchuan.github.io/blog/computer-technology/go/learning_go/
  - 用途：参考长篇学习型文章的讲解节奏和章节铺排方式。
- `datasea` 的 Go 教程文章
  - https://datasea.cn/go0225510828.html
  - 用途：参考偏入门、偏教学式写法，尤其是如何把概念拆成适合新手跟做的段落。

## 项目 1：`edge-gateway` 的参考资料

### 一级参考
- `lonng/nano`
  - `docs/get_started.md`
  - `docs/get_started_zh_CN.md`
  - `docs/design_patterns.md`

### 二级参考
- `grpc-game-example`
- `luk4z7/go-concurrency-guide`

### 适合拿来生成 tutorial 的主题
- 如何组织连接管理和消息路由。
- 如何从最小聊天示例抽象出房间服务。
- 如何设计 Session、Handler、广播模型。
- 如何做最小压测和并发问题排查。

## 项目 2：`room-orchestrator` 的参考资料

### 一级参考
- `Build an Orchestrator in Go`
  - GitHub 代码仓库
  - 书籍章节预览

### 二级参考
- `heroiclabs/nakama` 的房间、匹配、运行时模块思路

### 适合拿来生成 tutorial 的主题
- 如何拆分 manager / scheduler / worker。
- 如何定义任务状态机、重试和补偿。
- 如何把游戏房间生命周期讲成一个编排问题。
- 如何把本地编排设计映射到未来的 K8s 控制面。

## 项目 3：`cluster-operator` 的参考资料

### 一级参考
- Kubebuilder 官方书
  - `quick-start`
  - `getting-started`

### 二级参考
- Go + Kubernetes CI/CD 示例

### 适合拿来生成 tutorial 的主题
- 如何创建最小 CRD 和 Controller。
- 如何写 Reconcile 循环并处理状态同步。
- 如何把 `room-orchestrator` 的部署动作搬进 K8s。
- 如何补齐 Docker、镜像构建、GitHub Actions、kind 验证。

## 给后续 tutorial 的统一输入模板
每次让 AI 生成教程时，输入至少包含这四部分：

1. 当前阶段
- 例如：`Week 04，正在做 edge-gateway 的会话鉴权和消息路由`

2. 当前目标
- 例如：`实现连接注册、鉴权校验、消息类型分发，并保留现有日志结构`

3. 参考资料
- 明确写出文档链接或仓库内具体文件。

4. 输出要求
- 步骤编号。
- 对应代码落点。
- 为什么这样设计。
- 如何运行和验证。
- 这一版先不做什么。

## 反例
- 不要一次要求生成“完整十周教程”。
- 不要把多个参考仓库混成一个 prompt 又不说明优先级。
- 不要让教程脱离你当前仓库现状，从零另起一个不相干的 demo。
