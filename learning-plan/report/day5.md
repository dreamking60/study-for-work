# Day5 Report

## Topic
今天的重点是把程序入口从“硬编码演示”推进到“可接收命令行参数、可稳定输出日志”的阶段。

相比 day4 的重点是 config pipeline，day5 更关注 CLI entry 的可运维性：
- 参数从哪里来
- 参数怎么校验
- 错误怎么统一返回
- 日志怎么稳定输出

## Design Standard
- Keep `main()` thin and let it only start the program.
- Use `run() error` as the application entry flow.
- Parse CLI arguments before building runtime config.
- Print stable log fields for both success and failure cases.

## Why I Changed The Parse Design
一开始我把 `parseArgs()` 直接绑定到了全局命令行环境上，也就是直接使用默认的 `flag` 和真实的进程参数。

这样写虽然能跑，但它有一个明显问题：**不够好测**。

Problems with the original design:
- it depended on global flag state
- tests could affect each other
- input was implicit instead of explicit
- parse behavior and process environment were too tightly coupled

所以我把参数解析拆成了两层：
- `parseArgs()`: read real CLI input from `os.Args`
- `parseArgsFrom(argv []string)`: parse an explicit argument list

This split is better because:
- test input becomes explicit
- parsing logic is easier to verify
- the program entry still stays simple
- error handling remains centralized

## Current Structure
Files in `cmd/app`:

- `main.go`: call `run()` and delegate final error handling
- `app.go`: build runtime config and run the demo flow
- `args.go`: parse and validate CLI arguments
- `error_handler.go`: print final error message

## Implementation Notes
### 1. Argument Parsing
我今天没有直接上第三方 CLI 框架，而是先用 Go 标准库 `flag` 做最小可用版本。

Supported flags for now:
- `--env`
- `--port`
- `--app`

Validation rules:
- `env` must be one of `local`, `dev`, `test`, `prod`
- `port` must be in range `1..65535`

### 2. Main Flow
主流程现在不再直接把默认配置写死在 `main()` 里，而是先解析参数，再把参数覆盖到默认配置上。

Current flow:
- parse args
- build config from defaults plus CLI overrides
- create demo task
- print startup log
- print task summary

### 3. Error Handling
错误处理集中在最外层，而不是每个函数内部自己打印。

This follows the same idea as day4:
- lower-level functions return `error`
- top-level entry decides how to print and exit

## Logging Standard
今天的日志先做到“字段稳定”，不追求复杂日志框架。

Current goal:
- success log has stable fields
- error log has stable fields
- output is easy to read and easy to assert later

Example fields:
- `level`
- `msg`
- `app`
- `env`
- `port`

## Test Coverage
今天最重要的测试不是 `main()`，而是参数解析逻辑。

Covered cases:
- default CLI args
- custom CLI args
- invalid `env`
- invalid `port`

What this proves:
- argument parsing is no longer tied only to process state
- CLI behavior can be verified with unit tests
- validation errors are explicit and stable

## Current Status
目前 day5 的基础骨架已经搭起来了：
- `main -> run -> parseArgs -> handleError` 的入口分层已经完成
- 参数解析已经改成可测试设计
- 日志输出已经开始往稳定字段靠拢

## Next Step
如果继续推进，我下一步更适合做两件事：

- add `--config` support and connect it to `LoadConfig`
- improve log format and decide whether CLI args should override file config
