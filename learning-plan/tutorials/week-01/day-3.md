# Week 01 Day 3：把 `--config` 真的做进去

## 今天你要完成什么
今天是本周第一个真正的代码日。

你要把“配置文件支持”从想法变成代码。目标不是支持所有格式，而是先让当前项目真正能通过 `--config` 读取配置。

## 今天主要改哪些文件
- `cmd/app/args.go`
- `cmd/app/app.go`
- `cmd/app/args_test.go`
- 必要时补 `cmd/app` 下的新测试文件

## 第一步：先看清现在参数是怎么解析的
打开 `cmd/app/args.go`，重点看：
- `Args` 结构里现在有哪些字段
- `parseArgsFrom` 是怎么注册 flag 的
- `run` 命令现在把哪些参数传到运行逻辑里

你今天要做的第一件实际改动，就是给 `Args` 增加一个字段：
- `ConfigPath string`

如果你不知道 `flag` 具体怎么写，先看：
- [library-usage-cheatsheet.md](../library-usage-cheatsheet.md)
  重点看 `flag`、`encoding/json`、`testing`

## 第二步：给 `run` 注册 `--config`
你要在参数解析里新增一个 flag：
- `--config`

这个 flag 目前的作用只有一个：
- 告诉程序配置文件路径在哪里

今天不要做的事情：
- 不要同时支持 4 种配置来源
- 不要一上来把 env fallback 也一起堆进去

先把 `--config` 路径拿到手。

## 第三步：在运行逻辑里真正读取配置文件
打开 `cmd/app/app.go`，找到当前 `run` 或 `runCommand` 逻辑。

你今天要做的是：
1. 如果没有传 `--config`，继续走当前默认逻辑
2. 如果传了 `--config`，调用 `internal/config` 里的加载逻辑
3. 把加载结果和默认配置合并

这里你会直接用到的 API 是：
- `os.Open`
- `json.NewDecoder`
- `fmt.Errorf("...: %w", err)`

## 第四步：明确合并顺序
今天你就固定成这一条：
- `flag > env > config file > default`

如果你今天还没支持 env，也没关系，但顺序必须先写清楚。

最重要的是：
- 代码行为和文档说法要一致

## 第五步：补测试
今天至少写 2 条测试：
- `TestParseArgs_ConfigFlag`
- `TestRunCommand_LoadConfigFile`

如果还有余力，再补一条错误测试：
- 配置文件不存在时，错误信息要带路径

## 今天建议你自己建一个最小样例配置
例如放到临时位置，内容先极简：
```json
{
  "app_name": "opsctl-local",
  "environment": "dev",
  "port": 8080
}
```

你今天只需要一个能触发加载流程的最小文件，不需要追求配置设计完整。

## 今天的验证命令
```bash
go test ./cmd/app ./internal/config
go run ./cmd/app run --config ./sample.json
```

## 今天完成后的标准
- `--config` 能被解析
- 传入配置文件后程序行为发生变化
- 至少有 2 条相关测试
- 错误时输出可读

## 今天会用到的库和最小代码

### `flag`
给 `run` 增加 `--config` 的最小写法：
```go
fs := flag.NewFlagSet("run", flag.ContinueOnError)
configPath := fs.String("config", "", "config file path")
if err := fs.Parse(argv); err != nil {
    return Args{}, err
}
args.ConfigPath = *configPath
```

### `os` + `encoding/json`
从文件读取配置：
```go
f, err := os.Open(path)
if err != nil {
    return Config{}, err
}
defer f.Close()

var cfg Config
if err := json.NewDecoder(f).Decode(&cfg); err != nil {
    return Config{}, err
}
```

### `testing`
最小参数测试：
```go
func TestParseArgs_ConfigFlag(t *testing.T) {
    args, err := parseArgsFrom([]string{"--config", "sample.json"})
    if err != nil {
        t.Fatalf("parse args: %v", err)
    }
    if args.ConfigPath != "sample.json" {
        t.Fatalf("unexpected config path: %q", args.ConfigPath)
    }
}
```

## 根据你之前报告补充的真实坑
- 如果报 `illegal character U+FF08` 这类错误，先检查是不是输入法把括号打成了中文全角
- 如果报 `undefined: path`，先回头看你拆函数时参数名是不是没统一
