# Terminal Slack Extension

## 定位
这个 extension 不是改你现有项目的主线，而是在现有实时系统能力基础上，外扩一个“纯 terminal 的团队协作程序”。

目标不是复刻完整 Slack Web，而是做一个：

- 纯 terminal 使用体验。
- 支持类似 Vim 的高效键位操作。
- 面向实时聊天和团队协作。
- 能承载 bot、命令、通知、线程、频道等 Slack 风格能力。
- 明确利用 Go 协程模型处理连接、事件分发、消息同步和 UI 更新。

建议项目代号：`termflow` 或 `slackt`

## 为什么适合接在当前项目后面
你当前仓库里已经有一个 `edge-gateway` 方向，它天然适合承接这类系统的实时接入层能力。

这个 extension 可以这样挂接：

- `edge-gateway` 负责连接管理、鉴权、消息路由、广播。
- 新增 `terminal-slack-extension` 负责产品化方向和技术设计。
- 后续如果要真正实现，可以拆成：
  - `termchat-server`
  - `termchat-tui`
  - `termchat-bot-sdk`

这比“单做一个聊天 CLI”更有项目价值，因为它能同时覆盖：

- Go 并发模型
- 长连接/实时通信
- TUI 交互设计
- bot 和事件系统
- 工程化与可扩展架构

## 产品目标
第一阶段先聚焦下面这类能力：

1. 用户登录和工作区接入
2. 频道列表、私聊列表、未读状态
3. 实时消息收发
4. 线程回复
5. @mention 和系统通知
6. slash command
7. bot 接入
8. 消息搜索
9. 终端内高效率操作

不建议第一阶段就做：

- 富文本完全兼容
- 音视频
- 文件大规模上传预览
- Slack 全量权限模型
- 复杂工作流自动化平台

## 交互体验方向
核心体验应该更像 “Vim + tmux + Slack” 的组合，而不是一个鼠标优先的聊天窗口。

### 建议的界面布局
- 左侧：workspace / channels / DM
- 中间：消息流
- 右侧：thread / members / bot output
- 底部：command line / compose box / mode indicator

### 建议的操作模式
- `Normal mode`
  - `j/k` 上下移动消息
  - `h/l` 切换面板
  - `gg/G` 跳到头尾
  - `/` 搜索
  - `n/N` 搜索结果跳转
  - `r` 回复
  - `t` 打开 thread
  - `i` 进入输入模式
  - `:` 打开命令模式
- `Insert mode`
  - 编辑消息
  - 输入 slash command
- `Command mode`
  - `:channel general`
  - `:dm alice`
  - `:bot deploy status`
  - `:unread`
  - `:quit`

这个交互模型能让程序明显区别于普通聊天客户端，也更适合 terminal power user。

## 推荐整体架构
建议采用三层结构：

1. `TUI Client`
2. `Realtime Gateway`
3. `Chat Core Services`

### 1. TUI Client
职责：

- terminal 渲染
- 键盘事件处理
- 本地状态缓存
- websocket 事件接收
- 本地通知、未读状态、模式切换

推荐库：

- TUI：`bubbletea`
- 组件：`bubbles`
- 样式：`lipgloss`
- Vim 键位可以自己封装一层 input state machine

原因：

- Go 原生生态成熟
- 状态驱动 UI，适合事件流
- 比自己直接操作 ANSI 控制序列稳定很多

### 2. Realtime Gateway
职责：

- websocket 连接管理
- session 鉴权
- 用户订阅关系维护
- fan-out 广播
- ack、重试、心跳
- backpressure 控制

这层可以直接延续你现在 `edge-gateway` 的设计思路。

### 3. Chat Core Services
职责：

- channel / dm / thread 模型
- message store
- mention 解析
- search indexing
- bot event dispatch
- command execution
- permission check

这层和 gateway 分离后，后面扩 bot、审计、消息回放会轻松很多。

## Go 协程应该怎么用
这里不是“为了用协程而用协程”，而是把协程放在最能体现价值的位置。

### 服务端协程模型
建议每个连接、每个房间、每类后台任务都有清晰的 goroutine owner。

#### 连接级
- 一个 goroutine 读 socket
- 一个 goroutine 写 socket
- 一个 goroutine 处理 session 超时/心跳

这样可以避免读写互相阻塞。

#### 频道级 / 房间级
- 每个 channel 可有一个 event loop goroutine
- 负责串行处理该 channel 的广播、线程事件、未读状态更新

这样可以减少共享锁竞争，很多状态在“单 owner goroutine”内完成。

#### 后台任务级
- mention 提醒
- search indexing
- bot webhook dispatch
- audit log append
- offline notification

这些通过 worker pool 或异步任务通道处理。

### 客户端协程模型
- 一个 goroutine 收 websocket 事件
- 一个 goroutine 处理 terminal input
- 一个 goroutine 驱动 UI state update / render tick
- 若有本地缓存落盘，再单独一个 goroutine 做异步 IO

关键点是不要让网络事件直接阻塞 UI。

## 推荐的并发设计原则
这个项目如果要稳，必须明确下面几个原则：

1. 尽量单写者
- 某类状态只由一个 goroutine 持有和修改。

2. channel 用于事件传递，不要到处共享 map + mutex
- mutex 还是要有，但不要把整个系统写成锁表。

3. 网络连接与业务处理解耦
- socket goroutine 只负责 IO，不负责复杂业务。

4. 慢操作异步化
- bot 执行、索引写入、通知投递都不要阻塞主消息链路。

5. 明确 backpressure 策略
- 用户终端卡住时，队列如何限长、丢什么、怎么提示。

## Slack 风格功能怎么落地

### 频道和私聊
最小数据模型：

- `Workspace`
- `User`
- `Channel`
- `Membership`
- `Conversation`
- `Message`
- `Thread`
- `Reaction`
- `Bot`

### 消息能力
第一阶段建议支持：

- 普通文本消息
- 回复某条消息
- thread reply
- edit/delete
- pin
- reaction
- mention
- code block

terminal 中没必要追求 HTML 富文本，重点是：

- 文本可读性
- 快捷操作
- 结构化展示

### bot 能力
这是很值得做的亮点。

建议分三类：

1. `slash command bot`
- 例如 `/deploy status`

2. `event bot`
- 监听 message posted、mention、channel joined

3. `local helper bot`
- 直接在 terminal 本地执行某些效率命令，例如日志查询、Git 状态、CI 查询

如果你把 bot 做成统一事件接口，这个项目会非常像“终端协作工作台”，而不只是“聊天工具”。

## 推荐协议和通信方式
优先级建议：

1. 客户端到服务端实时链路：`WebSocket`
2. 内部服务通信：先用内存接口 / HTTP，后面再考虑 gRPC
3. bot 回调：HTTP webhook 或内部事件总线

为什么先不用 gRPC 做前端链路：

- terminal client 和服务端最核心是事件推送
- WebSocket 对双向实时消息更直接
- 调试成本低

## 存储建议
第一版不要上太重。

### 可行组合
- 元数据：`PostgreSQL`
- 热状态 / presence / 短期队列：`Redis`
- 搜索：第一版先数据库 LIKE / trigram，后续再接 Meilisearch 或 OpenSearch

### 这样拆的原因
- 聊天系统需要可靠消息落库
- presence 和未读计数更适合内存态/缓存态
- 搜索通常后补，不要一开始引入太重的依赖

## 模块拆分建议
如果后面真正实现，我建议目录这么拆：

```text
terminal-slack-extension/
  docs/
    architecture.md
    roadmap.md
    protocol.md
  termchat-server/
  termchat-tui/
  termchat-bot-sdk/
```

服务端内部可以拆：

```text
internal/
  gateway/
  session/
  channel/
  message/
  thread/
  presence/
  bot/
  search/
  notify/
```

## MVP 路线
第一版不要做“大而全”，建议按下面顺序推进。

### MVP-1：最小可聊天
- 登录
- channel 列表
- 实时消息收发
- TUI 基础布局
- Vim 风格浏览和输入切换

### MVP-2：更像 Slack
- thread
- mention
- unread
- DM
- reaction
- 搜索

### MVP-3：差异化亮点
- slash command
- bot framework
- terminal 内执行运维动作
- 可配置快捷键
- 消息过滤器 / 订阅器

## 这个项目真正的难点
难点不在“terminal 能不能画出来”，而在下面这些点：

1. UI 状态和实时事件同步
2. 高并发下的 fan-out 和未读一致性
3. thread / channel / dm 三套会话模型的统一抽象
4. bot 执行不能拖慢主链路
5. terminal 交互既快又不乱

## 我给你的技术判断
这个方向是成立的，而且有亮点，但前提是别把它做成：

- 只有 TUI，没有后端设计
- 只有 goroutine 口号，没有明确 owner model
- 只会发消息，没有 thread / bot / command 这些高价值功能

更好的讲法应该是：

“基于 Go 构建一个纯终端的实时团队协作系统，采用 goroutine + event loop 模型支撑高并发消息分发，并通过 Vim 风格交互、bot 扩展和 terminal-native workflow，提供区别于传统 GUI IM 的高效率协作体验。”

这句以后可以直接发展成项目介绍。

## 下一步建议
如果你认可这个方向，下一步我建议不是立刻写代码，而是继续补三份文档：

1. `architecture.md`
- 系统组件、事件流、协程模型、模块边界

2. `roadmap.md`
- MVP 分阶段、哪些先不做、每阶段验收标准

3. `protocol.md`
- websocket 事件类型、消息结构、ack、错误码

等这三份稳定后，再决定是先做 `termchat-tui`，还是先补 `edge-gateway` 的实时基础能力。
