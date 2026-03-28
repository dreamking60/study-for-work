# Week2 Day4 — `task update` 与状态流转

## 今日目标
- 为任务提供更新能力：修改状态、名称、标签、元数据。
- 在更新时记录 `UpdatedAt` 字段，并保持幂等。
- 将 `task-summary` 命令改为基于真实数据聚合。

## 产出物
- `cmd/app/task_update.go`。  
- `internal/model/task.go` 扩充 `CreatedAt`, `UpdatedAt`（`time.Time`）。  
- `task-summary` 逻辑更新，读取存储数据。

## 实施步骤
1. **扩展 Task 结构**  
   - 为 `model.Task` 添加 `CreatedAt`, `UpdatedAt` 字段，`Task.Summary()` 中打印时间（可格式化为 `2006-01-02`）。  
   - Day2 中创建任务时确保填写 `CreatedAt: time.Now()`, `UpdatedAt: time.Now()`。
2. **实现 `task update`**  
   - Flag 示例：`--id 123 --status success --name "ship feature" --set-tag go --del-tag legacy --set-meta owner=alice --del-meta owner`。  
   - 更新流程：
     1. 读取任务列表，定位 ID。  
     2. 根据 flag 修改字段；标签/元数据建议写成辅助函数。  
     3. 设置 `UpdatedAt = time.Now()` 并保存。  
   - 若找不到 ID，返回 `fmt.Errorf("task %d not found", id)`。
3. **刷新 summary 命令**  
   - 读取所有任务，统计 `status -> count`、`owner -> count`，并输出 JSON 或表格。  
   - 可新增 `internal/summary` 包来聚合数据。  
   - 输出示例：`status=success count=3 latest=Task[12] ship feature`。
4. **测试**  
   - 针对标签/元数据变更写纯函数测试。  
   - `task update` 命令测试可使用临时文件 + 真实 JSON 进行端到端校验。

## 调试清单
- `go run ./cmd/app task update --id 1 --status success --set-tag done`。  
- `go run ./cmd/app task-summary`（应展示真实统计）。

## Hints
- `map[string]string` 在序列化时键顺序不稳定，如需稳定输出可在打印前复制到 `[]kv` 并排序。  
- 在更新标签时可利用 `slices.Index` 判断是否存在，避免重复。  
- 对 `--set-meta` 这类 kv 输入，可沿用 Day2 中的解析工具函数。  
- 若后续要支持批量更新，考虑把更新逻辑抽象成 `func ApplyPatch(t *Task, patch UpdatePatch)`。
