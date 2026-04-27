# Terminal Slack Extension Roadmap

## Phase 0: 方案冻结
验收标准：
- 明确系统边界
- 明确客户端和服务端职责
- 明确 websocket 事件协议

## Phase 1: MVP Chat
范围：
- 用户登录
- channel 列表
- 实时收发消息
- terminal 基础布局
- Normal / Insert / Command 三种模式

验收标准：
- 两个终端用户可稳定聊天
- UI 不因高频消息卡死
- 基础快捷键流畅可用

## Phase 2: Slack-like Core
范围：
- DM
- thread
- mention
- unread
- reaction
- 搜索

验收标准：
- 支持多会话切换
- thread 视图可独立使用
- mention 和未读状态可追踪

## Phase 3: Bot and Workflow
范围：
- slash command
- event bot
- terminal helper bot
- webhook 接入

验收标准：
- bot 执行不阻塞主消息链路
- 命令结果可回显到消息流或 side panel

## Phase 4: 工程化
范围：
- 压测
- 可观测性
- 重连恢复
- 配置化快捷键
- 权限基础模型

验收标准：
- 有最小压测数据
- 异常重连可恢复
- 有基础日志和指标
