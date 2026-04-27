# Week 03 Day 6：做一次最小联调，不要只靠单元测试想象

## 今天你要完成什么
今天你要把前几天的零散功能串起来，做一轮最小联调。

联调的意义不是“跑起来一次就算完”，而是验证这些东西真的接起来了：
- 鉴权
- join
- message
- leave 或 heartbeat

## 今天建议准备的联调清单
至少测这几条：
1. 合法 token + join
2. 非法 token + join
3. join 后 message
4. leave 后再发 message

## 为什么今天必须做联调
因为你现在已经开始跨越多个模块：
- protocol
- auth
- router
- room
- store

单测通过不代表它们连起来就一定对。

## 今天完成后的标准
- 至少一条成功路径跑通
- 至少一条失败路径跑通
- 你能明确指出当前还有哪些边界没做好

## 今天会用到的库和最小代码

### 联调清单建议写成这样
```text
1. valid token + join -> success
2. invalid token + join -> unauthorized
3. join + heartbeat -> last seen updated
4. leave + message -> room/session error
```

### 为什么今天不只靠单测
因为现在你已经跨过多个模块了，联调是验证“模块之间接线”。
