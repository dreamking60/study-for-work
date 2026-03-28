# Week2 Day3 — `task list` 与过滤/格式化

## 今日目标
- 提供 `task list` 子命令，支持多种过滤条件与排序。
- 提供表格模式和 JSON 模式两种输出格式。
- 构建覆盖主要过滤路径的测试。

## 产出物
- `cmd/app/task_list.go`（命令实现）。
- `internal/present/table.go`（可选）集中处理输出格式。
- 对过滤函数的单元测试。

## 实施步骤
1. **解析过滤参数**  
   - flag 列表：`--status`, `--tag`, `--owner`, `--limit`, `--json`.  
   - 允许重复 flag，例如 `--tag go --tag cli`，可用自定义 `type multiFlag []string`。  
   - 对于 limit，若 <=0 则忽略或视为无限。
2. **加载并过滤任务**  
   - `tasks := store.MustLoad()`（需要处理错误）。  
   - 使用 `slices.DeleteFunc` 或自定义过滤器链：
     ```go
     filtered := tasks
     filtered = filterByStatus(filtered, statuses)
     filtered = filterByTags(filtered, tags)
     filtered = filterByOwner(filtered, owner)
     ```
   - `filterByTags` 可要求任务包含所有指定标签。
3. **排序与限制**  
   - 默认按 `CreatedAt`（若无，则按 `ID` 降序）排序。若尚未记录时间，可在 Day2 生成任务时补上时间戳字段。  
   - 使用 `sort.SliceStable`。应用 limit 前先排序。
4. **渲染输出**  
   - 表格：`text/tabwriter`，列为 `ID | Name | Status | Tags | Owner`。  
   - JSON：`json.MarshalIndent(filtered, "", "  ")`。  
   - 若过滤结果为空，打印 `no tasks matched`，返回 `nil`。
5. **测试**  
   - 针对过滤函数写纯 Go 测试，输入固定数组输出期望数组。  
   - Command 层测试可以 fake store，断言写入的 `bytes.Buffer` 内容包含特定字符串。

## 调试清单
- `go run ./cmd/app task list --status running --tag go`。  
- `go run ./cmd/app task list --json | jq '.'`。  
- 若 tasks 为空，命令应退出码 0 且提示信息友好。

## Hints
- 可以给 `Task` 增加 `Owner()` 辅助方法，默认读取 `Meta["owner"]`，避免在多处硬编码 key。  
- 过滤逻辑越纯越好，保持与 IO/打印解耦，方便测试。  
- `multiFlag` 的实现示例：
  ```go
  type multiFlag []string
  func (m *multiFlag) String() string { return strings.Join(*m, ",") }
  func (m *multiFlag) Set(v string) error { *m = append(*m, v); return nil }
  ```
- 离线模式下若需要示例数据，可编写 `docs/samples/tasks_seed.json` 并在 README 中说明。
