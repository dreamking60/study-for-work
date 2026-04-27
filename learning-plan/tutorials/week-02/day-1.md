# Week 02 Day 1：创建 `edge-gateway` 仓库骨架

## 今天你要完成什么
今天的任务不是“写完一个服务器”，而是先把项目骨架建对。

如果骨架没立住，后面你每加一个功能，目录都会越来越乱。

## 今天工作目录
```bash
mkdir -p ../../../go-backend-learning/edge-gateway
cd ../../../go-backend-learning/edge-gateway
go mod init edge-gateway
```

## 今天建议的目录结构
```text
cmd/server/main.go
internal/protocol/message.go
internal/session/session.go
internal/room/room.go
internal/room/service.go
internal/server/router.go
internal/server/server.go
```

## 为什么这样拆
- `cmd/server` 只放启动入口
- `internal/protocol` 只放消息结构
- `internal/session` 只放连接上下文
- `internal/room` 只放房间和房间服务
- `internal/server` 放接入和路由

这是为了让你后面一看目录就知道职责边界，而不是所有东西塞进 `main.go`。

## 今天你至少要创建哪些文件
- `cmd/server/main.go`
- `internal/protocol/message.go`
- `internal/session/session.go`
- `internal/room/room.go`
- `internal/server/server.go`

## 今天最小目标
即使今天每个文件只有很少的内容，也没关系。

你至少做到：
- `go test ./...` 不报目录错误
- `go run ./cmd/server` 能启动一个最小 server

## 今天不要做什么
- 不要接数据库
- 不要做鉴权
- 不要做广播

今天只是立骨架。

## 今天完成后的标准
- 仓库建立完成
- 目录职责明确
- 最小入口可运行

## 今天会用到的库和最小代码

### `net/http`
如果你今天想先起一个最小服务，可以这样写：
```go
mux := http.NewServeMux()
srv := &http.Server{
    Addr:    ":8080",
    Handler: mux,
}
if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    panic(err)
}
```

### `fmt`
最小启动日志：
```go
fmt.Println("edge-gateway starting on :8080")
```
