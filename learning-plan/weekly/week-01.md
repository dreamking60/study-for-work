# Week 01：Go 基础与工程化（详细执行版）

## 本周目标
- 完成 Go 开发环境与项目初始化。
- 建立标准目录结构：`cmd/app`、`internal`。
- 产出可运行 CLI v0.1 与最小测试闭环。

## 为什么先做这周
- 后续 Week2-Week10 都基于这个骨架扩展。
- 你需要先把“会运行代码”变成“会组织工程”。

## Day 1
- 学习：Go 安装、`go mod`、包与模块关系。
- 编码：初始化项目，创建 `cmd/app/main.go`。
- 验收：`go run ./cmd/app`、`go test ./...`。
- 产出：`README` 初版 + `day1.md`。

## Day 2
- 学习：`struct`、method、interface、slice/map。
- 编码：实现 `internal/config`、`internal/model`。
- 验收：`main` 能调用模块并输出。
- 产出：新增 2 个内部包。

## Day 3
- 学习：`error`、错误包装、JSON/文件读取。
- 编码：配置加载函数（先 JSON，再考虑 YAML）。
- 验收：非法配置时报错信息可定位字段。
- 产出：`docs/error-handling.md`。

## Day 4
- 学习：`flag` 参数解析与日志设计。
- 编码：实现 `-config`、`-env`、`-verbose`。
- 验收：参数错误时有清晰提示。
- 产出：日志规范小节写入 README。

## Day 5
- 学习：单测、表驱动、覆盖率。
- 编码：给 config/model 关键逻辑补测试。
- 验收：`go test ./... -cover` 通过，核心包 >= 60%。
- 产出：`TESTING.md`。

## Day 6
- 学习：CLI 命令组织方式。
- 编码：实现至少 2 个子命令（如 `plan`、`check`）。
- 验收：命令重复执行结果稳定。
- 产出：`v0.1.0` 本地标签。

## Day 7
- 学习：复盘方法与文档结构。
- 编码：清理命名、注释、目录结构。
- 验收：新同学按 README 10 分钟可跑通。
- 产出：`weekly-review/week1.md`。

## 本周完成标准
1. CLI 可运行。
2. 至少 5 个 commit（体现真实过程）。
3. 测试可通过并有覆盖率记录。
4. 文档齐全（README + 周复盘）。

## 常见卡点
1. `missing import path`：检查 `import` 是否用 `()`。
2. `go: command not found`：确认 PATH 包含 `/usr/local/go/bin`。
3. 测试空跑：至少给一个断言，不要只留空函数。
