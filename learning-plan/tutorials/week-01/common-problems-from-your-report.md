# Week 01 常见问题：根据你已写报告反向补充

这份文档不是泛泛而谈，而是根据你已经写过的 Day1-Day7 报告整理出来的真实坑点。

## 1. Go 装好了，但命令还是不能用

### 你之前遇到的问题
- 安装完 Go 之后，命令不能直接使用

### 根因
- `PATH` 没配好

### 你应该检查什么
```bash
go version
echo $PATH
```

### 典型修法
```bash
echo 'export PATH=/usr/local/go/bin:$HOME/go/bin:$PATH' >> ~/.bashrc
source ~/.bashrc
```

## 2. 多包导入时写成了 `{}`，不是 `()`

### 你之前遇到的问题
```bash
cmd/app/main.go:2:8: missing import path
```

### 根因
- Go 的多包导入必须用小括号 `()`
- 不能用花括号 `{}`

### 正确写法
```go
import (
    "fmt"
    "time"
)
```

## 3. 该返回字符串时用了 `fmt.Print`

### 你之前遇到的问题
- `Summary()` 这种函数应该返回字符串，但误用了直接打印函数

### 根因
- `fmt.Print` / `fmt.Printf` 是直接输出
- `fmt.Sprintf` 才是“生成字符串并返回”

### 正确写法
```go
func (t Task) Summary() string {
    return fmt.Sprintf("Task[%d] %s status=%s", t.ID, t.Name, t.Status)
}
```

## 4. 结构体字面量忘记写逗号

### 你之前遇到的问题
```bash
syntax error: unexpected newline in composite literal; possibly missing comma or }
```

### 根因
- 多行结构体字面量最后一个字段后面也要保留逗号

### 正确写法
```go
task := model.Task{
    ID:     1,
    Name:   "demo",
    Status: "running",
}
```

## 5. 代码里混进了中文全角符号

### 你之前遇到的问题
```bash
illegal character U+FF08 '（'
```

### 根因
- Go 语法位置必须使用 ASCII 标点
- 中文输入法会把括号、逗号、冒号变成全角字符

### 处理方式
- 代码区全部切英文输入法
- 注释可以是中文，但语法符号必须是英文半角

## 6. 拆函数时参数名没对齐，出现 `undefined: path`

### 你之前遇到的问题
```bash
undefined: path
```

### 根因
- 函数参数、局部变量和调用处没统一

### 处理方式
- 一旦把逻辑拆成 `load(path string)` 这种函数，就要先检查：
  - 参数名是否一致
  - 内部引用是否还在用旧变量名
  - 调用处是否传了正确参数

## 7. `flag` 直接绑全局状态，后面测试会很难写

### 你之前已经意识到的改进
- `parseArgs()` 读取真实 CLI
- `parseArgsFrom(argv []string)` 负责真正解析

### 为什么这是对的
- 测试输入显式
- 不会被全局 `flag` 状态污染
- 更容易覆盖 bad case

## 8. coverage 低不一定代表代码差

### 你之前已经观察到的现象
- `internal/config` 和 `internal/model` coverage 高
- `cmd/app` coverage 低一些

### 为什么正常
- 纯函数和配置管线更容易测
- 入口编排、输出和进程行为天然更难测

### 你现在应该怎么理解 coverage
- 用 coverage 找漏掉的分支
- 不要为了数字好看去补低价值测试

## 9. 命令分发应该和参数解析分开

### 你之前已经做对的事情
- `run()` 负责命令分发
- `parseArgsFrom()` 负责参数解析

### 为什么这很重要
- 后面加新子命令更轻松
- 测试边界更清楚
- 错误也更容易定位

## 10. 当前 Week 01 最容易重复犯的错
- 还没看清现状就开始大改
- 一上来加太多命令
- 想同时做配置、存储、日志、子命令、README
- 文档和代码行为不一致

## 你现在该怎么用这份文档
- 在开始 Week 01 前先读一遍
- 每天卡住时先来这里对照
- 如果后面又踩到新坑，继续把它补进来
