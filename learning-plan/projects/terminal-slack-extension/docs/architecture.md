# Terminal Slack Extension Architecture Draft

## 系统目标
- 在 terminal 中完成团队沟通和协作
- 复用 Go 的并发模型处理实时事件
- 支持 Slack 风格的频道、私聊、线程、bot 和命令系统

## 组件划分

### TUI Client
- 输入状态机
- 面板布局管理
- 本地缓存
- websocket client
- 通知和未读状态

### Gateway
- 长连接管理
- 鉴权
- 用户订阅
- 事件转发
- 限流和背压

### Chat Service
- 消息写入
- thread 聚合
- mention 解析
- 历史记录查询
- 搜索

### Bot Runtime
- slash command
- 事件订阅
- 内部任务触发

## 关键事件流

### 发消息
1. TUI 输入消息
2. client 发送 `message.send`
3. gateway 做鉴权和路由
4. chat service 落库
5. channel event loop 广播
6. 各 client 更新消息流和未读

### thread 回复
1. 用户打开 thread 面板
2. client 拉取 thread 历史
3. 用户回复
4. thread 事件写入主消息关联链
5. 订阅者收到 thread update

### bot 触发
1. 用户输入 slash command 或 @bot
2. gateway 转为标准事件
3. bot runtime 异步消费
4. bot 结果以消息或 ephemeral event 返回

## 并发模型

### 服务端
- `connReader`：每连接一个
- `connWriter`：每连接一个
- `heartbeatLoop`：每连接一个
- `channelLoop`：每频道一个
- `workerPool`：bot、搜索、通知等后台任务共享

### 客户端
- `inputLoop`
- `networkLoop`
- `uiUpdateLoop`
- `cacheFlushLoop`

## 关键设计原则
- 消息主链路短而稳定
- 慢任务全部异步
- channel 是事件边界，不是共享状态泥团
- UI 和网络状态解耦
- 所有事件都要可观测

## 风险
- TUI 状态复杂度容易快速失控
- terminal 渲染与异步事件竞争会造成闪烁或错位
- thread / unread / mention 一致性不好做
- bot 如果无隔离会污染主服务稳定性

## 建议的第一实现顺序
1. gateway 事件协议
2. TUI 基础布局和模式切换
3. 单频道实时聊天
4. 多频道和未读
5. thread
6. bot framework
