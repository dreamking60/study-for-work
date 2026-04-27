# day3
## Module
Config Module
```go
package config

// Config stores runtime settings
type Config struct {
	AppName string
	Env     string
	Port    int
}

// Default returns a baseline config for local development.
func Default() Config {
	return Config{
		AppName: "demo",
		Env:     "local",
		Port:    8765,
	}
}
```


model/task module
```go
package model

import "fmt"

// Task is a business object in this small demo.
type Task struct {
	ID int // Unique identifier of the task.
	// Human-readable task name shown in logs/CLI.
	Name string
	// Current lifecycle state, e.g. pending/running/success/failed.
	Status string
	// Category labels used for filtering and grouping.
	Tags []string
	// Extensible key-value metadata, e.g. owner/source/priority.
	Meta map[string]string
}

// Summary returns a short printable description for CLI output.
func (t Task) Summary() string {
	return fmt.Sprintf("Task[%d] %s status=%s tags=%d",
		t.ID, t.Name, t.Status, len(t.Tags))
}
```

## Result
```bash
:!go run ./cmd/app
App=demo, Env=local, Port=8765
Task[1] learning-go-day2 status=running tags=1
```

## Error&Warning
### 1
Use Sprintf instead of Print.
We wonder the task.Summary() function should return a string 
describe the detail of a task info.

  1. fmt.Print(...)

  - 作用：直接打印到标准输出（终端）
  - 返回：(n int, err error)，不是字符串
  - 适合：在 main 里直接打印日志/信息

  2. fmt.Printf(format, ...)

  - 作用：按格式打印到标准输出
  - 返回：(n int, err error)
  - 适合：你想格式化并直接输出

  3. fmt.Sprintf(format, ...)

  - 作用：按格式生成字符串并返回
  - 返回：string
  - 适合：函数需要返回文本（比如 Summary()）

### 2
忘记加逗号，修复即可
```bash
cmd/app/main.go:17:31: syntax error: unexpected newline in composite literal; possibly missing comma or }
cmd/app/main.go:18:27: syntax error: unexpected ] at end of statement
cmd/app/main.go:19:5: syntax error: non-declaration statement outside function body
```

