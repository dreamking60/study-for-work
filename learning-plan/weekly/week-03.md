# Week 03：服务化与存储（详细执行版）

## 本周目标
- 建立登录服务 v1（HTTP API）。
- 接入 MySQL + Redis，Mongo 先完成建模与示例。
- 完成认证链路（JWT + 会话缓存）。

## Day 1
- 学习：MySQL 索引、事务、慢查询基础。
- 编码：建用户表与索引。
- 验收：常用查询命中索引。
- 产出：`sql/schema.sql`。

## Day 2
- 学习：Redis string/hash/ttl 及缓存穿透问题。
- 编码：实现 session 存储到 Redis。
- 验收：登录后可从 Redis 读取会话。
- 产出：`internal/store/redis`。

## Day 3
- 学习：Mongo 文档模型场景与边界。
- 编码：设计 1 个半结构化集合（审计日志）。
- 验收：可写入并按条件查询。
- 产出：`docs/mongo-model.md`。

## Day 4
- 学习：Gin 路由、中间件、参数绑定。
- 编码：实现 `/register`、`/login`。
- 验收：接口返回码与错误结构统一。
- 产出：`cmd/auth-service`。

## Day 5
- 学习：JWT 签发、刷新、过期策略。
- 编码：完成登录鉴权中间件。
- 验收：未登录请求被拒绝，已登录可访问。
- 产出：`internal/auth`。

## Day 6
- 学习：测试分层（单测/集成测试）。
- 编码：为 auth/store 加测试。
- 验收：`go test ./...` 通过。
- 产出：`reports/week3-test.md`。

## Day 7
- 学习：API 文档最小规范。
- 编码：补 README 与接口说明。
- 验收：他人可按文档调用接口。
- 产出：`docs/api.md`。

## 本周完成标准
1. 登录流程跑通。
2. MySQL + Redis 实际参与链路。
3. 有接口文档和测试记录。
