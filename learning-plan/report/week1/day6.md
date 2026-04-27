# Day6 Report

## Topic
今天的重点是把“测试通过”进一步推进到“知道自己测到了什么、为什么还没到 100%”。

day6 的核心不是盲目追求高 coverage，而是理解：
- 哪些逻辑最值得测
- 哪些路径还没测到
- coverage 数字为什么会长这样

## Design Standard
- Treat tests as behavior constraints, not as a final checklist item.
- Prefer high-value unit tests over low-value output-only tests.
- Use coverage to find missing branches, not to chase 100% blindly.

## Coverage Result
你本地跑出来的结果是：

```bash
go test ./... --cover
ok      week01-cli/cmd/app           coverage: 37.9% of statements
ok      week01-cli/internal/config   coverage: 91.7% of statements
ok      week01-cli/internal/model    coverage: 100.0% of statements
```

## Why Coverage Looks Like This
### 1. `internal/model` is 100%
这一层只有一个核心行为：`Task.Summary()`。

Because the package is small and focused:
- there is only one main behavior to verify
- one direct unit test can cover almost all executable statements

### 2. `internal/config` is high
这一层 coverage 高，是因为它本身很适合做单元测试：
- input is explicit
- output is explicit
- error branches are clear

Covered areas:
- `Default()`
- `loadBytes()`
- `parseByExt()`
- `validate()`
- `LoadConfig()`

我今天又额外补了几条更深入的测试：
- invalid YAML
- invalid TOML
- default config output

These tests matter because they hit branches that were previously easy to miss.

### 3. `cmd/app` is lower
这里低不是因为代码一定差，而是因为这个 package 里有更多“入口编排逻辑”，天然就比纯函数难测。

Why it is lower:
- `main()` itself is thin and usually not tested directly
- `run()` coordinates several steps and prints output
- `handleError()` writes to stderr and exits the process

这些函数不是完全不能测，而是测试成本更高、收益没 `parseArgsFrom()` 那么直接。

So for day6, I intentionally prioritized:
- `parseArgsFrom()` behavior
- invalid CLI inputs
- explicit argument parsing paths

instead of forcing tests for:
- `main()`
- `handleError()`
- output-only integration behavior

## What I Added Today
Compared with the previous test set, I added deeper coverage in two places.

### 1. `cmd/app`
New covered paths:
- invalid flag parse, for example `--port abc`
- `parseArgs()` wrapper using real `os.Args`

Why this matters:
- it proves the parser handles both explicit argv and real CLI input
- it covers the parse-error path, not only validation errors

### 2. `internal/config`
New covered paths:
- `Default()` baseline config
- invalid YAML parse branch
- invalid TOML parse branch

Why this matters:
- it covers real error branches instead of only happy path
- it explains why previous coverage stopped below full branch coverage

## What I Learned
coverage 不是“分数越高越好”这么简单。

What matters more is:
- whether core behaviors are protected
- whether important error paths are covered
- whether the remaining uncovered code is actually worth testing

For this project:
- `config` and `model` should stay highly covered
- `cmd/app` can stay lower for now, because its remaining code is mostly orchestration and process-exit behavior

## Current Status
目前 day6 的目标可以认为已经完成了：
- config/model tests are in place
- CLI parsing tests were also added
- coverage result is recorded
- missing coverage is now explainable, not mysterious

## Next Step
如果继续推进，我更适合把注意力放到 day7 的 CLI command 设计，而不是继续为了 coverage 数字强行补低价值测试。
