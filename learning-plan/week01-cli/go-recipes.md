# Go 标准库速查（Week1-2 常用）

> 离线环境下可直接复制示例片段，按需调整命名与错误处理。

## Flag / CLI
```go
fs := flag.NewFlagSet("task-add", flag.ContinueOnError)
fs.SetOutput(io.Discard)
name := fs.String("name", "", "Task name")
var tags multiFlag
fs.Var(&tags, "tag", "Repeat to add multiple tags")
if err := fs.Parse(argv); err != nil { return err }
if *name == "" { return fmt.Errorf("name is required") }
```
- `multiFlag` 可以实现 `Set`/`String` 方法来接受重复参数。

## JSON 读写
```go
f, _ := os.Create(tempPath)
defer f.Close()
enc := json.NewEncoder(f)
enc.SetIndent("", "  ")
if err := enc.Encode(tasks); err != nil { return err }
```
```go
f, _ := os.Open(path)
defer f.Close()
dec := json.NewDecoder(bufio.NewReader(f))
if err := dec.Decode(&tasks); errors.Is(err, io.EOF) { return nil }
```
- 使用 `json.Unmarshal` 读取小文件，`json.Decoder` 更适合流式。

## 文件/路径
```go
home, _ := os.UserHomeDir()
defaultPath := filepath.Join(home, ".week01-cli", "tasks.json")
abs, _ := filepath.Abs(input)
if _, err := os.Stat(abs); errors.Is(err, fs.ErrNotExist) { /* 提示 */ }
```
- 写文件前执行 `os.MkdirAll(filepath.Dir(abs), 0o755)`。

## 时间与持续时间
```go
since := time.Now().Add(-30 * 24 * time.Hour)
cutoff := time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local)
d, err := time.ParseDuration("72h")
```
- 自定义天数解析：`days := flag.Int("days", 7, "..."); duration := time.Duration(*days) * 24 * time.Hour`。

## 排序/过滤
```go
sort.SliceStable(tasks, func(i, j int) bool {
    if tasks[i].Status == tasks[j].Status {
        return tasks[i].ID > tasks[j].ID
    }
    return tasks[i].Status < tasks[j].Status
})
```
```go
filtered := slices.DeleteFunc(tasks, func(t model.Task) bool {
    return !wantedStatuses[t.Status]
})
```
- `slices.DeleteFunc` 需要 Go 1.21+；若版本较低，可手写过滤。

## 文本输出
```go
w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', tabwriter.Debug)
fmt.Fprintf(w, "ID\tName\tStatus\tTags\n")
for _, t := range tasks {
    fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", t.ID, t.Name, strings.Join(t.Tags, ","))
}
w.Flush()
```
- `tabwriter.Debug` 会显示列线，调试完成后可删除。

## 测试片段
```go
func TestFileStoreSaveAndLoad(t *testing.T) {
    dir := t.TempDir()
    path := filepath.Join(dir, "tasks.json")
    store := storage.FileStore{Path: path}
    want := []model.Task{{ID: 1, Name: "demo"}}
    require.NoError(t, store.Save(want))
    got, err := store.Load()
    require.NoError(t, err)
    assert.Equal(t, want, got)
}
```
- 如无法使用 `require/assert`，可用标准库：`if !reflect.DeepEqual(want, got) { t.Fatalf(...) }`。

## 错误包装
```go
if err != nil {
    return fmt.Errorf("load config %s: %w", path, err)
}
```
- 解包：`if errors.Is(err, fs.ErrNotExist) { ... }`。

## 环境变量 / 配置
```go
path := os.Getenv("TASK_DATA")
if path == "" {
    path = defaultPath
}
env := os.Getenv("APP_ENV")
switch env {
case "", "local", "dev", "test", "prod":
default:
    return fmt.Errorf("invalid APP_ENV %q", env)
}
```
- 设置默认值可用 `if v, ok := os.LookupEnv("APP_PORT"); ok { ... }`。

## 临时文件与目录
```go
tmpDir := t.TempDir()              // 测试
tmpFile, err := os.CreateTemp("", "cfg-*.yaml")
if err != nil { return err }
defer os.Remove(tmpFile.Name())
```
- 需要一次性写入可用 `os.WriteFile(tmpFile.Name(), data, 0o600)`。

## 字符串与缓冲
```go
var b strings.Builder
fmt.Fprintf(&b, "Task[%d] %s", task.ID, task.Name)
fmt.Println(b.String())
```
```go
buf := &bytes.Buffer{}
logger := log.New(buf, "task ", log.LstdFlags)
logger.Printf("created id=%d", task.ID)
```
- `strings.Builder` 避免频繁分配；`bytes.Buffer` 适合测试中捕获输出。

## Context 与超时（可用于未来网络/IO）
```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
if err := doWork(ctx); errors.Is(err, context.DeadlineExceeded) {
    // 提示用户重试
}
```
- 在 CLI 中可将 `context.Context` 传递给需要可取消的操作，例如模拟远程同步。

## 使用 bufio.Scanner 逐行读取
```go
f, err := os.Open(path)
if err != nil { return err }
defer f.Close()
scanner := bufio.NewScanner(f)
for scanner.Scan() {
    line := scanner.Text()
    // 处理每一行
}
if err := scanner.Err(); err != nil { return err }
```
- 当需要导入行式数据（例如批量任务）时非常有用。

## HTTP / REST 访问
```go
req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
resp, err := http.DefaultClient.Do(req)
if err != nil { return err }
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("unexpected status %d", resp.StatusCode)
}
var payload struct {
    Tasks []model.Task `json:"tasks"`
}
if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
    return err
}
```
- 若需自定义超时，可创建 `http.Client{Timeout: 5 * time.Second}`。

## OS 信号处理
```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()
go func() {
    <-ctx.Done()
    fmt.Println("graceful shutdown...")
    // flush / save
}()
```
- 在 CLI 中接收 Ctrl+C，配合 `context` 让长任务可取消。

## sync.Mutex / sync.RWMutex
```go
type SafeStore struct {
    mu sync.RWMutex
    tasks []model.Task
}
func (s *SafeStore) List() []model.Task {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return append([]model.Task(nil), s.tasks...)
}
func (s *SafeStore) Add(t model.Task) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.tasks = append(s.tasks, t)
}
```
- `RWMutex` 读多写少时有优势，记得遵守 Lock/Unlock 成对调用。

## Goroutine + channel 模式
```go
jobs := make(chan model.Task)
go func() {
    defer close(jobs)
    for _, t := range tasks {
        jobs <- t
    }
}()
for t := range jobs {
    fmt.Println("processing", t.ID)
}
```
- 若需要 worker pool，可额外开多个 goroutine 消费 `jobs`。

## CLI 输出测试
```go
func TestListCommand(t *testing.T) {
    buf := &bytes.Buffer{}
    cmd := taskListCommand{writer: buf, store: fakeStore}
    require.NoError(t, cmd.Run(context.Background(), args))
    assert.Contains(t, buf.String(), "Task[1]")
}
```
- 通过依赖注入 `io.Writer`，可以在测试中断言输出内容。

## 小贴士
- 可在 `cmd/app` 下放一个 `logWriter io.Writer` 便于测试中注入 `bytes.Buffer`。  
- 需要 mock 当前时间时，可定义 `var now = time.Now` 并在测试里覆盖 `now = func() time.Time { ... }`。  
- 当多个命令共享 store/config，可创建 `type AppContext struct { Cfg config.Config; Store storage.TaskStore }`，函数接收该上下文。
- 组合多个错误可用 `errors.Join(errs...)`，向用户展示时再配合 `errors.Is/As` 做拆解。
