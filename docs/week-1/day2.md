# Day 2 · Error + Validation 巩固日（v23 Phase 1）

与 `docs/handbook-v2/README.md` **Phase 1 · Day 2** 一致。

> **说明**：本文件是 **v23 的学习步骤与验收清单**；**代码由你按步骤自己写**。若未单独说「请帮我改仓库」，不应由他人直接替你落库改代码。  
> Day 1 扎实了就有 buffer。今天把 **error handling** 铺好，为 **Day 4 并发**做准备。

---

## 学什么（约 45 min）

- **Error 返回值模式**：为什么 Go 用 `(result, error)`，而不是 `try/catch`  
- **自定义 error 类型**：实现 `Error() string`  
- `**defer` 的本质**：在函数返回前按 **LIFO** 执行延迟调用（像压栈）  
- `**defer` 常见用法**：关资源、`Unlock`、入口/出口日志  
- `**strconv` 转换**：`strconv.Atoi` 返回 `(int, error)`，调用方必须处理

资源：Go by Example — Errors, Panic, Defer；需要时查 `pkg.go.dev/strconv`。

**备忘**：[errors-panic-defer.md](errors-panic-defer.md)（`Is` / `As`、`panic`、`defer`）

---

## 主项目任务（约 120 min）

### Step 1：自定义校验错误（约 20 min）

新建 `internal/service/errors.go`：

- 定义 `ValidationError`，字段 `Field`、`Message`  
- 实现 `func (e *ValidationError) Error() string`（**pointer receiver**，便于 `errors.As`）

### Step 2：在 `Search` 里做参数校验（约 40 min）

修改 `internal/service/search.go` 的 `Search`：

- `Query` 为空 → `&ValidationError{Field: "query", Message: "cannot be empty"}`  
- `Page < 0` → `&ValidationError{Field: "page", Message: "must be non-negative"}`  
- `Store == ""` → 设为 `"default"`  
- 成功时 mock：`Results` 里带上 `query` 与 `store`（例如 `result for … at …`）

### Step 3：Handler 解析 query + 映射 HTTP 状态（约 40 min）

修改 `internal/handler/search.go`：

- 用 `r.URL.Query()` 读 `q`、`store`；`page` 用 `strconv.Atoi`（v23 示例：`page, _ := strconv.Atoi(...)`，解析失败当 `0` 即可）  
- 调 `svc.Search`；若 `errors.As(err, &verr)` 且为 `*service.ValidationError` → **400** + JSON `{"error":"..."}`  
- 其它错误 → **500** + JSON（不要泄露内部栈给客户端）  
- 抽 `writeJSONError(w, code, msg)`，统一 `Content-Type: application/json`  
- 在 `ServeHTTP` 开头 `log` + `**defer log`**（练习 `defer`）

### Step 4：手动验收（约 10 min）

```bash
go run ./cmd/server
```

```bash
# 正常
curl -s "http://127.0.0.1:8080/search?q=milk&store=ny01&page=1"

# 校验失败
curl -s "http://127.0.0.1:8080/search?q="
curl -s "http://127.0.0.1:8080/search?q=milk&page=-1"
```

期望：空 `q`、负 `page` 返回 **400** 且 body 为 JSON `error` 字段；正常请求 **200**。

---

## 可选（时间够再做）

- 多加规则（如 query 最大长度）  
- 对比 `errors.Is` 与 `errors.As` 的使用场景（Day 12 会再用 `%w`）

---

## Senior 视角（话术）

- 「Go 的 error 是值，不是异常；`errors.As` 用来在 error 链里取具体类型。」  
- 「`defer` 保证出口日志或清理一定会跑，比到处复制 `finally` 更直。」

---

## 常见坑

- 自定义 `Error()` 用 **pointer receiver**，否则 `errors.As` 可能对不上类型  
- `defer` 的参数在 `**defer` 语句执行时**求值，不是函数返回时  
- `err == nil` 判断顺序：先处理无错误路径  
- `strconv.Atoi`：非法字符串时返回 `(0, err)`；若你选择忽略 `err`，要在注释里说明设计取舍

---

## 完成标志

- 空 `q` → **400** + JSON error  
- 负 `page` → **400** + JSON error  
- 正常查询 → **200** + 带 `store` 的 mock 结果  
- `handler` 里能看到 **start / done** 两条日志（`defer`）  
- `go build ./...` 通过

---

## 记忆激活

当年改过的 bug 很多是 **参数校验** 与 **错误码**；今天把「校验在 service、映射在 handler」写进肌肉记忆。

---

## 与旧 20 天 Day 2 的关系

旧课表「Struct + JSON + Query」已迁到 [day02-legacy-20day-struct-json-query.md](day02-legacy-20day-struct-json-query.md)；**struct / JSON** 深挖仍看 [struct.md](struct.md)。