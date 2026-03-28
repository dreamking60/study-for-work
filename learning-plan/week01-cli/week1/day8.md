# Week1 Day8 — CLI 配置文件与参数融合

## 今日目标
- 在 `run` 子命令中读取 `--config` 文件（支持 JSON/YAML/TOML）。
- 使用 `internal/config` 包完成加载、校验，并与命令行 flag 合并。
- 确保命令行参数始终优先于配置文件，提供明确的日志输出。

## 产出物
- 更新后的 `cmd/app/args.go`（新增 `--config`）。
- 在 `cmd/app/app.go` 中加载配置、合并并输出最终值。
- 至少 2 个针对 happy/edge case 的单元测试（`args_test.go` 或新增文件）。

## 实施步骤
1. **扩展 flag 解析**  
   - 在 `parseArgsFrom` 中注册 `config := fs.String("config", "", "config file path")`。  
   - 解析完成后使用 `filepath.Clean` 与 `filepath.Abs` 标准化路径，允许空字符串（表示不使用文件）。
2. **分离配置加载逻辑**  
   - 在 `runCommand` 中检测 `args.ConfigPath`，若非空则调用 `config.LoadConfig`。  
   - 记录加载阶段的错误并返回；提示语包含文件名便于排障。
3. **合并优先级**  
   - 先以 `cfg := config.Default()` 初始化，再合并文件值，最后覆盖命令行 flag。  
   - 可以实现一个 `mergeConfig(base Config, incoming Config) Config` 辅助函数，或直接在 `runCommand` 内写出覆盖逻辑。  
   - 约定：非零值/非空字符串才覆盖，`Tags` 等复杂字段后续再扩展。
4. **输出可观察信息**  
   - 在 `fmt.Printf` 中补充 `configSource` 字段（`flags`, `file`, `flags+file`）。  
   - 若未提供文件，日志中明确提示 "using defaults"。
5. **测试**  
   - 在 `cmd/app/args_test.go` 中新增 `TestParseArgs_ConfigFlag` 等用例。  
   - 为 `runCommand` 单独写测试：使用临时文件创建一个 config，并验证最终日志或返回值。  
   - 使用 `t.Setenv("APP_CONFIG", tmpFile)`（若你决定支持 env fallback）。

## 调试清单
- `go test ./cmd/app -run Config -v` — 确保 flag/合并逻辑覆盖。  
- `go run ./cmd/app run --config ./local.yaml --env prod --port 8080` — 应优先使用 flag 中的 env/port。

## Hints
- `os.Stat(path)` 可区分 "文件不存在" 与其他错误。  
- `config.LoadConfig` 内部已做格式判断与校验，不需要重复验证。  
- 使用 `errors.Is(err, fs.ErrNotExist)` 提示用户文件缺失。  
- 若需要在日志中隐藏敏感信息，可在 `Config` 上实现 `String()` 并自行格式化。
