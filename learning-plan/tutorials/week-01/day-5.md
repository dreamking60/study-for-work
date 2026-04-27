# Week 01 Day 5：把错误提示和日志收干净

## 今天你要完成什么
今天的重点不是加功能，而是让你以后用这个工具时不想骂人。

好的 CLI 最重要的一点是：
- 出错时告诉你哪里错了
- 成功时告诉你程序现在用了什么配置

## 你今天要检查的 3 类错误

### 1. 参数错误
例如：
- `--port` 传错类型
- 缺少必要参数

### 2. 文件错误
例如：
- 配置文件不存在
- 路径不合法

### 3. 配置校验错误
例如：
- 环境值不合法
- 端口无效

## 今天建议跑的命令
```bash
go run ./cmd/app run --port bad
go run ./cmd/app run --config /tmp/not-exist.json
go run ./cmd/app run --env invalid
```

## 你今天要观察什么
不要只看“失败了”，要看失败信息质量：
- 有没有指出具体参数名
- 有没有指出具体文件路径
- 有没有指出具体字段

## 今天日志里至少应该带什么
成功路径里，建议至少带：
- 配置来源
- 环境
- 端口
- 数据路径

这不是为了好看，是为了后面每次运行都能快速确认“程序这次到底吃了什么配置”。

## 今天不要做什么
- 不要加花哨日志框架
- 不要一上来搞 trace id 体系

先把最小可读性做好。

## 今天完成后的标准
- 参数错误可读
- 文件错误可读
- 配置校验错误可读
- 成功日志里能看见最关键的运行信息

## 今天会用到的库和最小代码

### `fmt.Errorf("%w")`
包装错误，保留底层原因：
```go
if err != nil {
    return fmt.Errorf("load config %s: %w", path, err)
}
```

### `errors.Is`
判断文件不存在：
```go
if errors.Is(err, os.ErrNotExist) {
    return fmt.Errorf("config file %s does not exist", path)
}
```

### 成功日志最小示例
```go
fmt.Printf("env=%s port=%d data=%s source=%s\n", cfg.Env, cfg.Port, dataPath, configSource)
```

## 根据你之前报告补充的真实坑
- 你之前已经把 `main()` 收薄、把错误交给外层统一处理了，所以今天不要再回到“每个函数自己打印错误”的老路
