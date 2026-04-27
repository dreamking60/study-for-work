# Week 01 Day 4：定下数据路径规则，但今天不做大存储层

## 今天你要完成什么
今天你不是做完整存储系统，而是先把“以后数据放哪里”这件事定下来。

这一步很关键，因为后面不管是：
- 本地任务
- 数据播种
- 回放输入
- 导出结果

都离不开统一的数据目录。

## 你今天要做的最小版本
先只支持路径规则，不急着写完整读写逻辑。

建议顺序：
1. 在 `Args` 里增加 `DataPath`
2. 支持 `--data`
3. 支持 `TASK_DATA`
4. 定一个默认路径

## 推荐优先级
- `flag > env > default`

今天你可以先不做配置文件里的数据路径，因为那会把问题同时拉大两层。

## 推荐默认目录
```text
~/.opsctl/
  tasks.json
  seeds/
  exports/
```

你可以保留旧路径，但要满足一个条件：
- 文档里必须明确写出来

## 今天要改的文件
- `cmd/app/args.go`
- `cmd/app/app.go`
- 如有需要，补一个专门的路径解析测试

## 今天运行逻辑只要做到什么程度
只要做到下面这个程度就够了：
- 程序能解析最终数据路径
- 日志里能打印它
- 你能验证优先级是对的

今天先不要急着做：
- `task add`
- `task list`
- JSON 文件读写

这些是下一阶段再做的，不是今天的重点。

## 今天建议的验证方式
你可以这样测：
```bash
go run ./cmd/app run --data /tmp/opsctl-tasks.json
```

然后再试环境变量：
```bash
TASK_DATA=/tmp/from-env.json go run ./cmd/app run
```

你要看的是：
- 最终路径是不是按你预期生效
- flag 有没有覆盖 env

## 今天完成后的标准
- `DataPath` 已进参数结构
- CLI 已支持 `--data`
- 文档里已经写清默认目录规则
- 日志里能看见最终路径

## 今天会用到的库和最小代码

### `os.LookupEnv`
读环境变量：
```go
if v, ok := os.LookupEnv("TASK_DATA"); ok && v != "" {
    dataPath = v
}
```

### `os.UserHomeDir` + `filepath.Join`
拼默认路径：
```go
home, err := os.UserHomeDir()
if err != nil {
    return "", err
}
defaultPath := filepath.Join(home, ".opsctl", "tasks.json")
```

### 优先级合并的最小写法
```go
path := defaultPath
if envPath != "" {
    path = envPath
}
if flagPath != "" {
    path = flagPath
}
```

## 根据你之前报告补充的真实坑
- 你之前已经在 config 管线里踩过“参数名不统一”的坑，所以今天加 `DataPath` 时要特别小心结构体字段名、flag 变量名和日志输出名是否一致
