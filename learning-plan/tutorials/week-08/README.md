# Week 08 教程：用 Kubebuilder 启动 `cluster-operator`

## 本周目标
- 准备本地 K8s 环境
- 初始化 Kubebuilder 项目
- 生成一个最小资源和 Controller

## 本周工作目录
```bash
mkdir -p ../../../go-backend-learning/cluster-operator
cd ../../../go-backend-learning/cluster-operator
```

## 本周建议只碰这些对象
- `RoomService` CRD
- 对应 Controller
- 一个 Deployment 模板

## 开始前先读
- [foundation-reference.md](../foundation-reference.md)
- [concepts-and-libraries.md](./concepts-and-libraries.md)

## 第 1 天：准备环境
至少检查：
- `go version`
- `docker version`
- `kubectl version --client`
- `kind version`
- `kubebuilder version`

缺哪个先补哪个，不要硬跳。

今天的停止点：
- 你已经知道本机缺什么，不会盲目继续

## 第 2 天：初始化项目
```bash
kubebuilder init --domain study.local --repo cluster-operator
```

## 第 3 天：创建 API 和 Controller
```bash
kubebuilder create api --group game --version v1alpha1 --kind RoomService
```

## 第 4 天：定义最小 Spec
这一版先只放：
- 镜像
- 副本
- 端口

完成标准：
- 字段已经写进 API 定义，而不是只在脑子里想

## 第 5 天：写第一版 Reconcile
第一版只做一件事：
- 根据 `RoomService` 创建或更新 Deployment

今天不要同时做 Service、Ingress、HPA。

## 第 6 天：本地 kind 跑起来
你要完成：
- 创建 kind 集群
- 安装 CRD
- 应用 sample
- 本地跑 controller

完成标准：
- controller 至少收到一次 reconcile

## 第 7 天：写阶段总结
说明：
- 你理解的 Reconcile 是什么
- Spec 和 Status 的区别
- 当前这一版故意没做什么

## 本周最终自检命令
```bash
go test ./...
```
