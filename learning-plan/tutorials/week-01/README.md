# Week 01 教程：收口 `opsctl`，把现有 CLI 变成后续项目工具

## 本周目标
- 不推翻你已经做完的 Day1-Day7。
- 把 `go-backend-learning/week01-cli` 收口成后续项目的工程 CLI。
- 补配置文件支持、数据路径约定、使用文档。

## 本周唯一工作目录
```bash
cd ../../../go-backend-learning/week01-cli
```

## 本周主要改哪些文件
- `cmd/app/args.go`
- `cmd/app/app.go`
- `cmd/app/args_test.go`
- `cmd/app/main_test.go`
- `internal/config/*.go`
- 如需要再新增 `internal/storage/`

## 开始前先读
- [foundation-reference.md](../foundation-reference.md)
- [concepts-and-libraries.md](./concepts-and-libraries.md)
- [common-problems-from-your-report.md](./common-problems-from-your-report.md)

## 本周你先不要新建什么
- 不要新建第二个 CLI 仓库
- 不要重命名 module
- 不要先做 `task add/list/update`
- 不要把它改造成真正的业务服务

## 本周结束时你应该得到什么
- 一个能正常跑、能测试、能解释定位的 CLI。
- 一套固定的配置优先级：`flag > env > config > default`
- 一份清晰的 README，说明它如何服务后续项目。

## 开始前先确认
先跑一遍当前状态：
```bash
go test ./...
go run ./cmd/app run
go run ./cmd/app task-summary
```

如果这里跑不通，先修现有问题，不要直接加新功能。

建议把当前输出记到你的本周报告里，作为“改造前基线”。

## Day 1：盘点已完成代码
今天不要写太多代码，先确认现状。

你要读的文件：
- `cmd/app/main.go`
- `cmd/app/app.go`
- `cmd/app/args.go`
- `internal/config/*.go`
- `internal/model/task.go`

你要写下来的结论：
- 现在有哪些命令已经真的存在。
- 哪些测试已经有。
- 哪些文档写了但代码没落地。

完成标准：
- 你能用 5 句话讲清当前 CLI 已有能力。

今天的停止点：
- 你已经能列出“已有能力”和“本周要补的两个缺口：配置文件、数据路径”

## Day 2：明确它的工具定位
今天只做一件事：把 CLI 从“练手项目”改成“后续项目工具”。

你要更新的文档：
- `learning-plan/tutorials/week-01-opsctl.md`
- 如有必要，在仓库根目录补一个 README

要写清楚的职责：
- 配置调试
- 本地数据播种
- 本地回放
- 运维辅助命令

完成标准：
- 文档里能明确说明它不再是独立求职项目，而是 `edge-gateway` 和 `room-orchestrator` 的工具层。

建议你今天顺手补一段“未来命令名单”，但只写名字，不实现：
- `opsctl seed`
- `opsctl replay`
- `opsctl export`

## Day 3：实现 `--config`
今天接你原来 `day8` 的计划，真正把配置文件支持做进代码。

要改的文件：
- `cmd/app/args.go`
- `cmd/app/app.go`
- 必要时新增测试

具体动作：
1. 给 `run` 增加 `--config`
2. 如果传了配置文件，就调用 `internal/config` 的加载逻辑
3. 合并顺序固定为：
   - flag
   - env
   - config file
   - default
4. 启动日志里打印配置来源

你可以直接跑：
```bash
go test ./cmd/app ./internal/config
go run ./cmd/app run --config ./sample.json
```

如果还没有样例配置文件，今天先自己创建一个最小样例，例如：
```json
{
  "app_name": "opsctl-local",
  "environment": "dev",
  "port": 8080
}
```

完成标准：
- 配置文件能生效
- 非法配置会报清晰错误
- 测试覆盖 happy path 和 bad case

今天的停止点：
- `run --config` 跑通
- 至少新增 2 条测试
- 启动日志能看出配置来源

## Day 4：补数据路径约定
今天的重点不是完整存储系统，而是把未来要用的数据路径定下来。

建议你先支持：
- `--data`
- `TASK_DATA`
- 默认路径

优先级建议：
- flag > env > default

默认路径建议：
- `~/.opsctl/tasks.json`
或者保留现有路径，但你必须写清楚

如果今天要顺手落一版代码，最小做法是：
1. 在 `Args` 里加 `DataPath`
2. 支持 `--data`
3. 在 `run` 日志里打印最终路径
4. 不急着真正读写任务数据

完成标准：
- CLI 输出里能看到实际数据路径
- 文档里写清楚默认目录布局

建议目录布局直接写死在教程里：
```text
~/.opsctl/
  tasks.json
  seeds/
  exports/
```

## Day 5：清理日志和错误提示
今天专门做“易用性收口”。

你要检查：
- 参数错误是不是一眼能看懂
- 缺配置文件时是不是会带文件名
- 配置校验错误是不是能定位字段

建议命令：
```bash
go run ./cmd/app run --port bad
go run ./cmd/app run --config /tmp/not-exist.json
```

完成标准：
- 失败信息不是模糊的 panic 或杂乱输出

今天顺手检查：
- 参数缺失时是不是给出 flag 名称
- 配置文件不存在时是不是带绝对路径
- 校验失败时是不是指出具体字段

## Day 6：整理 README 和快速开始
今天只做文档收口。

你要补：
- 如何运行
- 如何测试
- 如何传配置文件
- 数据文件放哪
- 这个工具未来服务哪些项目

完成标准：
- 未来你隔一周回来，不需要重新猜用法

建议你至少把这三段写出来：
1. 现在能做什么
2. 现在还不能做什么
3. 这个工具将如何服务后续三个项目

## Day 7：写本周报告
这份报告只写 4 件事：
- 保留了哪些旧成果
- 本周新增了哪些真实能力
- 它将如何服务 `edge-gateway`
- 它将如何服务 `room-orchestrator`

## 本周不要做什么
- 不要继续无上限扩展 task 子命令
- 不要把它膨胀成新的主项目
- 不要为了“功能多”牺牲结构和可解释性

## 本周通过线
只要你满足下面这些，就可以进入 Week 02：
- `go test ./...` 通过
- `run` 可以加载配置文件
- 有固定数据路径规则
- README 已更新
- 你能讲清这不是练手代码，而是工具层

## 本周最终自检命令
```bash
go test ./...
go run ./cmd/app run
go run ./cmd/app run --config ./sample.json
go run ./cmd/app task-summary
```

## 按天执行
- [day-1.md](./day-1.md)
- [day-2.md](./day-2.md)
- [day-3.md](./day-3.md)
- [day-4.md](./day-4.md)
- [day-5.md](./day-5.md)
- [day-6.md](./day-6.md)
- [day-7.md](./day-7.md)
