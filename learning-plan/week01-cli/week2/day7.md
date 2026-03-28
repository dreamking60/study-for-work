# Week2 Day7 — 发布、打包与质量保障

## 今日目标
- 为 CLI 项目准备发布流程：构建、打包、校验。
- 提升测试覆盖率并生成报告。
- 产出操作手册，方便在离线环境重复执行。

## 产出物
- `Makefile` 或 `Taskfile.yml`，包含 `build/test/format/release` 目标。
- `scripts/release.sh`（可选）将二进制与示例配置/数据一同打包。
- 文档：`docs/release-checklist.md`，记录手动步骤与常见问题。

## 实施步骤
1. **标准化 CLI 构建**  
   - `make build`：运行 `go build -o bin/week01-cli ./cmd/app`。  
   - `make test`：执行 `go test ./...`.  
   - `make lint`：可用 `go vet ./...` + `staticcheck`（若离线不可用，可暂记 TODO）。
2. **发布产物**  
   - `make release VERSION=0.1.0`：
     1. 清理 `bin/`。  
     2. 运行 `GOOS=darwin/linux`、`GOARCH=amd64/arm64` 的交叉编译。  
     3. 将 `bin/week01-cli-darwin-amd64` 等文件连同 `docs/samples` 打包 zip/tar。  
   - 在 `docs/release-checklist.md` 中列出需要预先下载的依赖（如 `zip` 工具）。
3. **测试覆盖率**  
   - `make cover`：`go test ./... -coverprofile=cover.out`，随后 `go tool cover -html=cover.out -o cover.html`。  
   - 在文档中说明如何在无 GUI 环境查看：`go tool cover -func cover.out`。
4. **烟囱测试脚本**  
   - `scripts/smoke.sh`：
     ```bash
     set -euo pipefail
     BIN="${BIN:-./bin/week01-cli}"
     $BIN run
     $BIN task add --name smoke --status pending --tags ci
     $BIN task list --limit 1
     ```
   - 在 release 步骤中自动运行。
5. **文档化**  
   - `docs/release-checklist.md` 提示：版本号规范、如何更新 CHANGELOG、如何验证签名（若有）。  
   - 记录常见故障：Go 版本不一致、缺少执行权限、`GOOS` 不支持等。

## 调试清单
- `make build && ./bin/week01-cli --help`（若你实现了 help）。  
- `make release VERSION=0.1.0` 应生成 `dist/week01-cli-0.1.0-darwin-arm64.zip` 等文件。

## Hints
- `make` 目标可使用 `.PHONY` 声明，防止与同名文件冲突。  
- 交叉编译前确保禁用 CGO：`CGO_ENABLED=0`.  
- 文档中附上 `sha256sum dist/*.zip` 命令，方便校验。  
- 若未来要上传 GitHub Release，可在脚本中预留 `gh release upload` 的命令（离线环境下可注释）。
