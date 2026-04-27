# 基础参考与常用库说明

这份文档是给“完全不懂的人”准备的总入口。

如果你在某一周教程里看到这些词不理解：
- `flag`
- `JSON`
- `context`
- `goroutine`
- `mutex`
- `state machine`
- `reconcile`

先回来看这份文档，再继续做当天任务。

如果你已经知道概念，但不知道代码到底怎么写，再看：
- [library-usage-cheatsheet.md](./library-usage-cheatsheet.md)

## 一、为什么前几周尽量优先用标准库
因为你现在不是在比“谁会装更多第三方库”，而是在建立这几个最重要的能力：
- 读懂 Go 工程
- 控制数据结构
- 管理配置和错误
- 做清楚的模块边界

标准库的好处是：
- 文档稳定
- 学完以后迁移到别的项目更容易
- 出问题时更容易定位

所以在前几周，我们会优先使用这些标准库：
- `flag`
- `os`
- `filepath`
- `encoding/json`
- `testing`
- `time`
- `context`
- `sync`

## 二、Week 01 最常用的库

### `flag`
- 文档：https://pkg.go.dev/flag
- 用来做什么：解析命令行参数，例如 `--config`、`--data`
- 为什么现在用它：你正在做 CLI，参数解析是核心工作

你现在最该认识的对象：
- `flag.NewFlagSet`
- `String`
- `Int`
- `Parse`

### `os`
- 文档：https://pkg.go.dev/os
- 用来做什么：读环境变量、判断文件是否存在、拿用户目录
- 为什么现在用它：配置路径和数据路径都要依赖它

你现在最常见的用法：
- `os.Getenv`
- `os.LookupEnv`
- `os.UserHomeDir`
- `os.Stat`

### `filepath`
- 文档：https://pkg.go.dev/path/filepath
- 用来做什么：拼路径、清理路径、拿绝对路径
- 为什么现在用它：不要手写字符串去拼文件路径

### `encoding/json`
- 文档：https://pkg.go.dev/encoding/json
- 用来做什么：读取和输出 JSON 配置或数据
- 为什么现在用它：JSON 是最容易上手、最容易观察的配置格式

### `testing`
- 文档：https://pkg.go.dev/testing
- 用来做什么：写单元测试
- 为什么现在用它：教程要求你每一步都能验证，不靠肉眼猜

## 三、Week 02-05 最常用的库

### `encoding/json`
- 为什么还在用：因为你先要把消息协议和数据结构讲清楚
- 为什么暂时不用二进制协议：对初学者来说，先看得见比先追求性能更重要

### `sync`
- 文档：https://pkg.go.dev/sync
- 用来做什么：保护共享状态，例如房间成员表
- 为什么会用到：一旦有多个连接或多个 goroutine，同一个 map 就可能被同时读写

### `time`
- 文档：https://pkg.go.dev/time
- 用来做什么：心跳、超时、定时清理、重试间隔
- 为什么会用到：实时服务离不开“时间”这个维度

### `context`
- 文档：https://pkg.go.dev/context
- 用来做什么：超时、取消、请求生命周期控制
- 为什么会用到：后面做连接管理和服务关闭时会非常重要

### `net/http` 和 `net`
- 文档：
  - https://pkg.go.dev/net/http
  - https://pkg.go.dev/net
- 用来做什么：搭网络入口
- 为什么要先认识它们：即使后面用了别的库，你也要知道 Go 网络编程底层是怎么组织的

## 四、Week 06-07 最常见的概念

### 状态机
- 它是什么：一个对象只允许从某些状态转到某些状态
- 为什么要用：这样房间流程不会乱跳

### 调度器
- 它是什么：定期检查哪些对象该推进到下一步
- 为什么要用：房间不会自己变化，必须有一层控制循环

### 幂等
- 它是什么：同一个请求重复执行，结果仍然可控
- 为什么要用：真实系统里重复消息和重试很常见

### 补偿
- 它是什么：某一步失败后，用另一组动作把系统拉回可接受状态
- 为什么要用：分布式流程不可能永远只走成功路径

## 五、Week 08 以后最常见的库和框架

### Kubebuilder
- 文档：https://book.kubebuilder.io/
- 用来做什么：帮你生成 K8s Operator / Controller 项目骨架
- 为什么要用：你不是从零手写整个控制器框架，而是站在标准脚手架上学习控制器思想

### `controller-runtime`
- 文档：https://pkg.go.dev/sigs.k8s.io/controller-runtime
- 用来做什么：写 Reconcile、管理 K8s 对象
- 为什么会出现：Kubebuilder 生成的项目就是围绕它工作的

## 六、现在先不用着急学的东西
前几周先不要钻进这些坑里：
- 复杂 ORM
- 完整微服务框架
- 很多第三方配置框架
- 一大堆日志库对比

你当前最需要的是：
- 会用标准库做最小闭环
- 知道什么时候该加第三方库，而不是一开始就依赖它
