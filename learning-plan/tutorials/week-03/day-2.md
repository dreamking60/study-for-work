# Week 03 Day 2：把最小鉴权接到消息入口前

## 今天你要完成什么
今天开始真正落鉴权，但只做最小版。

今天不是做完整安全体系，而是做一条很明确的规则：
- 未鉴权消息不允许进入业务 handler

## 今天你主要改哪些文件
- `internal/server/router.go`
- 新增 `internal/auth/auth.go`
- `internal/session/session.go`

## 你今天的最小实现思路
1. 给消息或连接上下文提供 token
2. 在路由分发前校验 token
3. 如果失败，直接返回错误
4. 如果成功，把身份写进 session

## 为什么鉴权最好放在入口
因为你现在的项目正在长大。

如果你不在统一入口挡住非法消息，后面每个 handler 都得重复做这件事：
- `join`
- `message`
- `heartbeat`
- `leave`

这会让结构很快变乱。

## 今天建议定义的最小错误
- `unauthorized`
- `invalid_token`

今天不用做一堆错误类型，只先让非法消息和业务错误分开。

## 今天建议补的测试
- `TestRouteRejectsUnauthorizedMessage`
- `TestRouteAllowsAuthorizedMessage`

## 今天完成后的标准
- 未鉴权消息会被挡在入口
- 合法消息可以继续往下走
- 至少有 2 条相关测试

## 今天会用到的库和最小代码

### 最小 token 校验函数
```go
func ValidateToken(token string) error {
    if token == "" {
        return fmt.Errorf("unauthorized")
    }
    if token != "demo-token" {
        return fmt.Errorf("invalid_token")
    }
    return nil
}
```

### 在入口拦截
```go
if err := auth.ValidateToken(msg.Token); err != nil {
    return err
}
```
