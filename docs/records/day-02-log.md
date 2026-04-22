# Day 2 - 2026-04-22

- Today's plan (30-second version):
  按 v23 Phase 1 Day 2：service 层校验 + 自定义 `ValidationError` + handler 映射 HTTP 状态与 JSON 错误体；巩固 `errors.As`、`defer`、query 解析。

- What I completed:
  - `internal/service/errors.go`：`ValidationError` + `Error() string`（pointer receiver，便于 `errors.As`）
  - `internal/service/search.go`：`Search` 内校验空 `query`、**`utf8.RuneCountInString`** 限制 query 长度、`page < 0`、`store` 默认 `"default"`
  - `internal/handler/search.go`：`r.URL.Query()` + `strconv.Atoi` 读 `q`/`store`/`page`；`errors.As` 识别 `*service.ValidationError` → **400** + JSON；其它 err → **500**；`writeJSONError`；`log` + `defer` 打 start/done
  - 文档：`docs/week-1/errors-panic-defer.md`（`Is` / `As`、`panic`、`defer` 备忘）、`go-basic.md` 预习索引；`day2.md` / `week-1/README.md` 链过去；根目录 `README.md`、`docs/handbook-v2/README.md`、`docs/50-interview-alignment.md` 小更新
  - 已提交并 push：`cursor/mini-search-scaffold-and-docs`（commit `602a63d` 及所含文件）

- Verification command(s):
  - `go build ./...`
  - `go run ./cmd/server`（改代码后记得重启进程）
  - `curl -s "http://127.0.0.1:8080/search?q=milk&store=ny01&page=1"`
  - `curl -s "http://127.0.0.1:8080/search?q="`
  - `curl -s "http://127.0.0.1:8080/search?q=milk&page=-1"`
  - （可选）超长 `q` 验证 **400** + `query must be at most 50 characters`

- Verification result:
  - build 通过
  - 空 `q`、负 `page` → **400**，body 含 JSON `error`
  - 正常请求 → **200**，mock `results` 含 `store`

- Blocker today:
  - 曾混淆 **`errors.Is` vs `errors.As`**；`fmt.Errorf` 与 `Println` 区别；Go 字符串长度用 **`len` / `utf8.RuneCountInString`** 而非 `.length`；**改代码后需重启 `go run`** 才生效。
  - `fmt.Println` 打在 **server 终端**，不会出现在 `curl` 的 response body 里。

- How I resolved it:
  对照 `docs/week-1/errors-panic-defer.md` 与 `day2.md`；用 `errors.As` + `*service.ValidationError`；本地验证后 `git commit` / `push`。

- Tomorrow first step (must be executable):
  打开 `docs/handbook-v2/README.md` 对照 **Phase 1 Day 3**（或手册下一日卡片），只做当天主任务；继续 **`go build` + `curl -i`** 验收。

- One Go concept reinforced today:
  **`errors.As(err, &ptr)`** 用于带 **fields** 的自定义错误；**`errors.Is`** 用于 **sentinel**；业务失败用 **`return ..., err`**，不要用 **`panic`** 当校验。

---

## Deep dive notes（可选）

- `errors` / `panic` / `defer` 场景备忘：`docs/week-1/errors-panic-defer.md`
- Week 1 索引：`docs/week-1/README.md`
