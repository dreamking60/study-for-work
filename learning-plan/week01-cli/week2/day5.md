# Week2 Day5 — `task report` 统计与可视化

## 今日目标
- 基于任务数据生成报表：状态分布、标签 TopN、完成耗时。
- 支持控制时间窗口（`--since 72h`）与输出格式（表格/JSON）。
- 通过独立的聚合层保持 CLI 与统计逻辑解耦。

## 产出物
- `internal/report/report.go`（聚合逻辑 + 结构体）。
- `cmd/app/task_report.go`（命令入口）。
- 聚合函数的测试覆盖 >80%。

## 实施步骤
1. **定义数据结构**  
   - `type Report struct { StatusCount map[string]int; TagTop []TagStat; AvgCycle time.Duration }`。  
   - `TagStat` 包含 `Name string`, `Count int`，可附带示例任务 ID。  
   - `func BuildReport(tasks []model.Task, since time.Time) Report`：过滤出最近任务再聚合。
2. **统计逻辑**  
   - 状态分布：遍历自增。  
   - Tag TopN：使用 `map[string]int` 累加后转成 slice 排序，输出前 N（默认 5，可用 flag 控制）。  
   - 平均耗时：若任务含 `Meta["duration"]` 或可用 `UpdatedAt.Sub(CreatedAt)`，计算平均完成耗时；缺失则跳过。
3. **CLI 命令**  
   - flag：`--since`（`time.ParseDuration`），`--top`（int），`--json`。  
   - 若 `--since` 为空，默认 30 天：`time.Now().Add(-30*24*time.Hour)`。  
   - 输出：表格可显示三块区域；JSON 直接 `json.MarshalIndent(report)`。
4. **测试**  
   - 使用固定的任务 slice，断言 `BuildReport` 返回的 map/slice 与期望一致。  
   - 对 `--since` 做 2 个边界测试：正好等于阈值、早于阈值。

## 调试清单
- `go run ./cmd/app task report --since 720h --top 3`。  
- `go run ./cmd/app task report --json | jq '.'`。

## Hints
- `time.ParseDuration` 不支持 "d"，需要你自己把 `--since-days` 转换为小时，或提供两个 flag。  
- 排序 Tag 时可用 `sort.Slice(tagStats, func(i, j int) bool { ... })`，当数量相同时按字母序。  
- 若任务不含 `UpdatedAt`，可把统计写成安全模式：缺失时不计入平均耗时，但在输出中提示 `avgDuration=N/A`。  
- 报表逻辑纯函数化，有助于未来做 HTTP API 复用。
