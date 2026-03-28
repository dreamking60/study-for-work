# Week2 Day6 — `task prune` 与备份

## 今日目标
- 实现 `task prune` 命令，支持删除已完成/过期任务。
- 在删除前自动备份数据，并提供恢复提示。
- 提升文件操作的安全性，避免并发写入冲突。

## 产出物
- `cmd/app/task_prune.go`。
- `internal/storage/backup.go`（或合并到 file_store）。
- 针对备份/删除的端到端测试。

## 实施步骤
1. **备份策略**  
   - `func Backup(path string) (string, error)`：复制到 `path + ".bak-20060102-150405"`。  
   - 可使用 `io.Copy`，或 `os.ReadFile` + `os.WriteFile`（注意权限）。
2. **命令参数**  
   - `--status success`（可多选），`--before 2024-12-31`（ISO 日期）。  
   - 解析日期：`time.Parse("2006-01-02", value)`。  
   - 支持 `--dry-run` 输出将被删除的 ID 而不落盘。
3. **删除流程**  
   - 先 `Backup`，成功后才继续。  
   - 遍历任务，保留不满足条件的项，统计被删除数。  
   - 删除后写回文件并打印 `pruned=3 backup=/path/to/file`。  
   - 若没有任务符合条件，打印 `nothing to prune` 并退出 0。
4. **并发防护（可选加分）**  
   - 使用 `flock`? 在纯 Go 中可利用 `syscall.Flock`（仅 Unix）。或写一个简易锁文件 `path + ".lock"`，在命令执行期间持有。  
   - 确保即使命令失败也释放锁。
5. **测试**  
   - `TestPruneByStatus`：准备 3 个任务（success/pending），确认只删除 success。  
   - `TestPruneBeforeDate`：创建不同 CreatedAt，断言结果。  
   - `TestDryRunNoWrite`：dry-run 时文件不变，可通过比较 `os.ReadFile` 内容实现。

## 调试清单
- `go run ./cmd/app task prune --status success --before 2024-01-01`。  
- `go run ./cmd/app task prune --dry-run --status failed`。

## Hints
- 备份文件命名可使用 `time.Now().Format("20060102-150405")`。  
- 删除时尽量使用纯函数 `filterTasks(tasks []Task, predicate func(Task) bool)`，测试更简单。  
- 若实现锁文件，建议写到与数据相同的目录，命名 `tasks.json.lock`。  
- 输出中提示备份路径和恢复方法（`cp backup tasks.json`），方便离线排障。
