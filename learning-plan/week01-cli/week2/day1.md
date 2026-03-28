# Week2 Day1 — 本地文件存储层

## 今日目标
- 设计并实现面向任务数据的文件存储（JSON 数组）。
- 支持 flag/env 默认的三层路径优先级。
- 为读写逻辑提供单元测试，确保离线环境也能运行。

## 产出物
- `internal/storage/file_store.go` 与对应测试文件。
- `cmd/app/args.go` 增加 `--data` flag（默认 `""`）。
- `cmd/app/app.go` 引入 `FileStore` 实例创建（后续命令复用）。

## 实施步骤
1. **确定数据文件路径**  
   - Flag `--data` > 环境变量 `TASK_DATA` > 默认 `~/.week01-cli/tasks.json`。  
   - 可使用 `os.UserHomeDir()` 与 `filepath.Join` 创建默认路径。  
   - 在 `Args` 结构中新增 `DataPath string`，并在 `parseArgsFrom` 返回。
2. **实现 FileStore**  
   - 结构体：`type FileStore struct { Path string }`，可额外挂 `mu sync.Mutex`。  
   - `func (s FileStore) Load() ([]model.Task, error)`：
     - 若文件不存在返回空 slice；使用 `os.IsNotExist` 检测。  
     - `json.NewDecoder` 解码，注意 `io.EOF` 表示空文件。  
   - `func (s FileStore) Save(tasks []model.Task) error`：
     - 将目录 `os.MkdirAll(filepath.Dir(s.Path), 0o755)`。  
     - 写到 `tmp, err := os.CreateTemp(dir, "tasks-*.json")`，完成后 `os.Rename` 原子替换。
3. **集成到 `runCommand`**  
   - 根据 `args.DataPath` 创建 `store := storage.NewFileStore(path)`，暂时只做初始化和日志输出。  
   - 返回结构中带上数据路径，方便后续命令使用。
4. **测试**  
   - `TestFileStoreLoadMissing`：使用 `t.TempDir()`，确认返回空 slice 且不报错。  
   - `TestFileStoreSaveAndLoad`：写入两条任务，重新读取并校验顺序/字段。  
   - 记得在测试中注入 `model.Task{}`，无需真实 CLI。

## 调试清单
- `go test ./internal/storage -v`。  
- `go run ./cmd/app run --data /tmp/tasks.json`，确认日志中显示实际路径。

## Hints
- `encoding/json` 在写出时可以使用 `json.NewEncoder(f).Encode(tasks)`，记得 `SetIndent` 提升可读性。  
- 若需要 ID 生成器，可在 `model` 层补充辅助函数，但今天可先留空。  
- 写文件时使用 `defer f.Close()`，确保 `os.Rename` 在 close 之后执行。  
- 离线环境下可把 sample JSON 预先放到 `docs/samples/`，方便未来 copy。
