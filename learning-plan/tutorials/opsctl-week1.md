# 教程 01：把 `week01-cli` 收成 `opsctl` 第一版

## 这份教程解决什么问题
你当前已经在 `go-backend-learning/week01-cli` 里做了配置、模型、参数解析、命令分发和测试，但这些成果还像练习，不像后续项目的工具入口。

这份教程的目标是：
- 保留现有代码。
- 把它收成一个明确的工程 CLI。
- 让它成为后续 `edge-gateway` 和 `room-orchestrator` 的辅助工具。

## 做之前先确认
- 代码目录：`go-backend-learning/week01-cli`
- 已有能力：
  - `run`
  - `task-summary`
  - `internal/config`
  - `internal/model`

## 今天只做 4 步

### 第 1 步：给项目补清晰定位
打开并更新这些文件：
- `go-backend-learning/week01-cli/cmd/app/main.go`
- `go-backend-learning/week01-cli/cmd/app/app.go`
- `learning-plan/tutorials/week-01-opsctl.md`

要完成的动作：
- 把 CLI 名称从“练习项目”视角改成“工程工具”视角。
- 在文档里写清楚这个工具未来要承担什么：
  - 配置调试
  - 数据播种
  - 本地回放
  - 运维辅助

完成标准：
- 你能用一句话解释它不是最终求职项目，而是后续项目的共用工具层。

### 第 2 步：实现配置文件支持
参考现有实现：
- `go-backend-learning/week01-cli/internal/config/*.go`
- `go-backend-learning/week01-cli/cmd/app/args.go`

你要改的代码位置：
- `go-backend-learning/week01-cli/cmd/app/args.go`
- `go-backend-learning/week01-cli/cmd/app/app.go`
- 必要时补测试文件

你要做的事情：
1. 给 `run` 命令增加 `--config`
2. 读取配置文件
3. 合并优先级固定为：
   - `flag`
   - `env`
   - `config file`
   - `default`
4. 运行时打印当前配置来源

完成标准：
- `go test ./...` 通过
- `go run ./cmd/app run --config ./example.json` 能看到最终生效配置

### 第 3 步：补数据目录约定
你要决定一件事：
- 后续任务数据、本地状态、回放输入放在哪里

建议：
- 默认目录：`~/.opsctl/` 或保留现有 `~/.week01-cli/`，但要在文档里说清楚
- 文件命名至少预留：
  - `tasks.json`
  - `seeds/`
  - `exports/`

如果今天不想一口气做完存储层，至少先把路径约定和命令入口留好。

完成标准：
- 文档里有固定路径约定
- CLI 输出里能看到实际使用的路径

### 第 4 步：收口并写一份阶段报告
把本周报告按下面格式写：
- 已有代码真实完成了什么
- 这次新增了什么
- 它将如何服务 `edge-gateway`
- 它将如何服务 `room-orchestrator`

## 今日命令清单
```bash
cd ../../go-backend-learning/week01-cli
go test ./...
go run ./cmd/app run
go run ./cmd/app task-summary
```

如果你加了配置文件支持，再跑：
```bash
go run ./cmd/app run --config ./sample.json
```

## 今天先不要做什么
- 不要继续往 CLI 里塞太多业务命令。
- 不要为了“看起来完整”去写一堆还服务不到主项目的功能。
- 不要重写已有代码，只做承接和收口。

## 做完以后你应该得到什么
- 一个能继续演进的工程 CLI。
- 一套明确的配置加载规则。
- 一份能解释当前成果价值的文档。
