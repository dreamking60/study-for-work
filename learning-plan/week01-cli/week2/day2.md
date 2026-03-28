# Week2 Day2 — `task add` 子命令

## 今日目标
- 新增 `task add` 子命令用于创建任务并写入文件存储。
- 支持标签、元数据与默认状态，自动生成任务 ID。
- 完成 CLI 层与存储层之间的串联测试。

## 产出物
- `cmd/app/task_add.go`（或类似文件）封装命令逻辑。
- 在 `main.go`/`app.go` 中增加命令路由：`task add`。
- 针对 `task add` 的测试（可用 fake store）。

## 实施步骤
1. **命令路由设计**  
   - 顶层结构建议：`app run`, `app task-summary`, `app task add`。  
   - 在 `run()` 中检测 `os.Args[1] == "task"`，再根据第二个参数分派。  
   - 若缺少 action，返回 `fmt.Errorf("missing task action: add/update/...")`。
2. **解析输入**  
   - 使用 `flag.NewFlagSet("task-add", ...)` 支持：
     - `--name` (必填)  
     - `--status` 默认 `pending`  
     - `--tags` 逗号分隔，使用 `strings.Split` 并 `slices.DeleteFunc` 清理空值  
     - `--meta` 形式如 `owner=dreamking,priority=high`，可用 `strings.SplitN` 处理。  
   - 校验：`name` 不能为空，`status` 限制在 `pending|running|success|failed`。
3. **生成任务并写入**  
   - ID 可以采用 `time.Now().UnixNano()` 或维护递增计数（前者更简单）。  
   - 读取现有任务：`tasks, err := store.Load()`，将新任务 append 后保存。  
   - 写入成功后打印 `Task[%d] created name=%s status=%s`。
4. **测试**  
   - fake store：实现 `type fakeStore struct { loadErr error; tasks []model.Task }`，便于注入。  
   - `TestTaskAddCreatesFile`：使用临时目录执行真实存储写入，随后 `Load` 验证数据。  
   - `TestTaskAddValidation`：传空 name，断言报错信息包含 `name is required`。

## 调试清单
- `go run ./cmd/app task add --name "Write docs" --tags go,cli --meta owner=dreamking`。  
- `jq '.' ~/.week01-cli/tasks.json` 验证写入格式。

## Hints
- 为了重用 flag 解析逻辑，可把 `parseTaskAddArgs(argv []string)` 抽成独立函数并在测试中直接调用。  
- 将 store 接口声明为 `type TaskStore interface { Load() ([]model.Task, error); Save([]model.Task) error }`，让未来命令可以依赖接口。  
- 在追加任务前可检查名称重复：`slices.ContainsFunc(tasks, func(t model.Task) bool { return t.Name == name })`。  
- 离线环境下若需要种子数据，可预置一个 JSON 文件供手工拷贝。
