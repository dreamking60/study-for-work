# 常用库具体用法速查

这份文档不是概念说明，而是“今天就能抄过去改”的最小用法。

如果你在教程里看到某个库不知道怎么写，先来这里找最接近的例子。

## `flag`

```go
fs := flag.NewFlagSet("run", flag.ContinueOnError)
configPath := fs.String("config", "", "config file path")
dataPath := fs.String("data", "", "task data path")
port := fs.Int("port", 8080, "listen port")
if err := fs.Parse(argv); err != nil {
    return Args{}, err
}
args := Args{
    ConfigPath: *configPath,
    DataPath:   *dataPath,
    Port:       *port,
}
```

## `os`

```go
if v, ok := os.LookupEnv("TASK_DATA"); ok && v != "" {
    dataPath = v
}
```

```go
home, err := os.UserHomeDir()
if err != nil {
    return "", err
}
```

```go
if _, err := os.Stat(path); err != nil {
    if errors.Is(err, os.ErrNotExist) {
        return fmt.Errorf("config file %s does not exist", path)
    }
    return err
}
```

## `filepath`

```go
path := filepath.Join(home, ".opsctl", "tasks.json")
```

```go
abs, err := filepath.Abs(input)
if err != nil {
    return "", err
}
```

## `encoding/json`

```go
f, err := os.Open(path)
if err != nil {
    return Config{}, err
}
defer f.Close()

var cfg Config
if err := json.NewDecoder(f).Decode(&cfg); err != nil {
    return Config{}, err
}
```

```go
var msg Message
if err := json.Unmarshal(raw, &msg); err != nil {
    return err
}
```

```go
enc := json.NewEncoder(f)
enc.SetIndent("", "  ")
if err := enc.Encode(tasks); err != nil {
    return err
}
```

## `testing`

```go
func TestParseArgs_ConfigFlag(t *testing.T) {
    args, err := parseArgsFrom([]string{"--config", "sample.json"})
    if err != nil {
        t.Fatalf("parse args: %v", err)
    }
    if args.ConfigPath != "sample.json" {
        t.Fatalf("unexpected config path: %q", args.ConfigPath)
    }
}
```

```go
dir := t.TempDir()
path := filepath.Join(dir, "config.json")
```

```go
t.Setenv("TASK_DATA", "/tmp/tasks.json")
```

## `time`

```go
sess.LastSeenAt = time.Now()
```

```go
expired := now.Sub(sess.LastSeenAt) > 30*time.Second
```

```go
ticker := time.NewTicker(5 * time.Second)
defer ticker.Stop()
for range ticker.C {
    cleanupExpiredSessions(time.Now())
}
```

## `sync.RWMutex`

```go
type RoomStore struct {
    mu    sync.RWMutex
    rooms map[string]*Room
}

func (s *RoomStore) Get(id string) (*Room, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    room, ok := s.rooms[id]
    return room, ok
}

func (s *RoomStore) Save(room *Room) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.rooms[room.ID] = room
}
```

## `context`

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

select {
case <-ctx.Done():
    return ctx.Err()
case result := <-workCh:
    return result
}
```

## `net/http`

```go
mux := http.NewServeMux()
mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("ok"))
})

srv := &http.Server{
    Addr:    ":8080",
    Handler: mux,
}

if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    return err
}
```

## `errors.Is` / `fmt.Errorf("%w")`

```go
if err != nil {
    return fmt.Errorf("load config %s: %w", path, err)
}
```

```go
if errors.Is(err, os.ErrNotExist) {
    return fmt.Errorf("config file not found")
}
```
