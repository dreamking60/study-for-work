# Day4 Report

## Topic
今天的重点是把 `config` 从“只有默认值的结构体”推进到“可加载、可解析、可校验”的阶段。

Day4 的核心不是把所有格式一次做完，而是先建立一个清晰的 error flow：
- `load`: 负责读取原始内容
- `parse`: 负责把内容解码成 `Config`
- `validate`: 负责检查字段是否合法

这样做的原因很直接：以后出错时，我可以马上判断问题是在文件读取、格式解析，还是字段校验阶段，而不是把所有错误混在一个函数里。

## Design Standard
- Use `Loader + Parser + Validator` to separate responsibilities.
- Keep `LoadConfig(path string) (Config, error)` as the public entry.
- Wrap errors with stage context such as `load config`, `parse config`, and `validate config`.
- Return field-specific validation errors when possible.

## Current Structure
Files in `internal/config`:

- `config.go`: define `Config` and `Default()`
- `loader.go`: read file bytes and detect extension
- `parser.go`: decode config by extension
- `validator.go`: validate business rules

## Implementation Notes
### 1. Loader
这里我把读取文件和识别扩展名放在一起处理。

Its responsibility:
- read config data from a file path
- detect file extension
- return raw bytes and extension for the parser

Expected behavior:
- empty path -> return `load config: empty path`
- missing file -> return wrapped read error
- missing extension -> return a clear load-stage error

### 2. Parser
这里的目标是让解析逻辑和文件读取解耦。

Supported formats for now:
- `.json`
- `.yaml` / `.yml`
- `.toml`

Expected behavior:
- invalid JSON/YAML/TOML -> return parse-stage error
- unsupported extension -> return `unsupported type`

### 3. Validator
这一层只关心业务合法性，不关心文件怎么读、内容怎么解码。

Validation rules for now:
- `AppName` must not be empty
- `Env` must be one of `local`, `dev`, `test`, `prod`
- `Port` must be in range `1..65535`

## Problems I Hit
### 1. 中文全角符号导致 Go 编译失败
一开始我在 `parser.go` 里误用了中文全角括号 `（ ）`，Go 编译器直接报非法字符。

Error example:
```bash
internal/config/parser.go:10:42: illegal character U+FF08 '（'
internal/config/parser.go:10:58: illegal character U+FF09 '）'
internal/config/parser.go:13:5: expected declaration, found 'switch'
```

What I learned:
- Go source code is very sensitive to punctuation in syntax positions.
- Use ASCII punctuation in code even if the note itself is written in Chinese.

### 2. `path` 变量未定义
在 `loader.go` 里，我一开始函数参数或者局部变量没有对齐好，结果后面多处直接引用了不存在的 `path`。

Error example:
```bash
internal/config/loader.go:28:5: undefined: path
internal/config/loader.go:32:27: undefined: path
internal/config/loader.go:34:52: undefined: path
internal/config/loader.go:37:22: undefined: path
internal/config/loader.go:39:72: undefined: path
```

这类问题说明我在写拆分函数时，参数名和职责边界还不够稳定。修复方式很简单，就是把 `loadBytes(path string)` 的参数和内部引用统一起来。

## Why This Design Is Better
之前如果把逻辑都塞进一个函数里，后面只会越来越难维护。

This split is better because:
- each stage has one clear job
- errors are easier to locate
- future formats can be added without changing validation logic
- tests can target each stage independently

## Acceptance Criteria
- Invalid file path returns a load-stage error.
- Invalid file content returns a parse-stage error.
- Invalid field value returns a validate-stage error with the field name.
- `LoadConfig` stays small and only coordinates the pipeline.

## Test Coverage
今天除了把 `Loader / Parser / Validator` 拆出来，我还把核心测试补上了。

Covered cases:
- `validate`: success, empty `AppName`, invalid `Env`, invalid `Port`
- `parseByExt`: valid JSON/YAML/TOML, invalid JSON, unsupported extension
- `LoadConfig`: success path, invalid path, invalid format, invalid field
- `model.Task`: `Summary()` output check

What this proves:
- the config pipeline works on the happy path
- errors can be located at `load`, `parse`, and `validate` stages
- current model behavior has at least one stable unit test

## Current Status
目前 `Loader / Parser / Validator` 的代码和对应测试都已经补出来了，day4 的主要设计目标算是完成了。

Completed today:
- finish config loading pipeline split
- fix syntax and variable errors during refactor
- add unit tests for `internal/config`
- add one unit test for `internal/model`

## Next Step
明天如果继续推进，我更适合把注意力放到 `main` 接入或者 day5 内容上，因为 day4 的“错误可定位”这一部分已经通过测试固定下来了。

Recommended next actions:
- integrate `LoadConfig` into `main`
- prepare sample config files for manual run
- continue to day5 CLI/logging work
