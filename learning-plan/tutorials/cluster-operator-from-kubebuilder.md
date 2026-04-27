# 教程 04：基于 Kubebuilder 做 `cluster-operator`

## 参考资料
- Kubebuilder Quick Start
  - https://book.kubebuilder.io/quick-start
- Kubebuilder Getting Started
  - https://book.kubebuilder.io/getting-started

## 为什么现在做这个
你已经有了：
- 接入服务方向
- 编排服务方向

接下来要补的是平台岗会问的东西：
- 你会不会用 Go 做控制器
- 你理解不理解 Reconcile
- 你能不能把服务带进本地 K8s

## 这份教程只要求你做最小闭环
- 初始化 Kubebuilder 项目
- 创建一个最小资源对象
- 写一版 Reconcile
- 在本地 kind 集群中跑起来

## 第 1 步：先准备环境
你至少需要：
- Go
- Docker
- kubectl
- kind
- kubebuilder

如果本地没装全，不要硬跳。
先把缺的东西记下来，然后分开解决。

## 第 2 步：初始化项目
建议新目录：
```bash
mkdir -p ../../go-backend-learning/cluster-operator
cd ../../go-backend-learning/cluster-operator
kubebuilder init --domain study.local --repo cluster-operator
```

这里先不要想“大而全平台”，只做一个最小资源，比如：
- `RoomService`

## 第 3 步：创建 API 和 Controller
按 Kubebuilder 官方流程：
```bash
kubebuilder create api --group game --version v1alpha1 --kind RoomService
```

生成后你先看这几类文件：
- `api/...`
- `internal/controller/...`
- `config/samples/...`

今天的目标不是把所有生成代码看懂，而是抓住两件事：
- Spec 是你期望的状态
- Status 是系统当前的状态

## 第 4 步：定义一个最小 Spec
今天只放最少字段，例如：
- 镜像
- 副本数
- 端口

不要一开始就塞资源限制、租户、节点选择、监控开关。

## 第 5 步：写第一版 Reconcile
第一版只做一件事：
- 当 `RoomService` 被创建时，创建或更新一个 Deployment

你不需要一开始就做复杂控制逻辑。

今天先确认这条链路成立：
- apply 一个 CR
- controller 收到变更
- reconcile 被触发
- 对应对象被创建

## 第 6 步：在本地集群跑起来
Kubebuilder 文档里的关键动作是：
- 安装 CRD
- 应用 sample
- 本地运行 controller

如果你用 kind，本地镜像可以直接加载，不需要先推远端仓库。
官方文档也明确推荐在本地开发和 CI 中优先使用 kind。

## 今天的完成标准
- Kubebuilder 项目初始化完成
- 有一个 `RoomService` 资源定义
- Reconcile 跑起来
- 在本地集群看到对应对象被创建

## 今天不要做什么
- 不要一开始就做多租户
- 不要先接完整 CI/CD
- 不要写复杂平台 UI

先把“一个 Go controller 管一个业务资源对象”这件事走通。
