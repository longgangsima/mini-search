# Week 1：语法快速过 + HTTP 基础

> **若你主跟《Go 23天 Day-by-Day 执行手册》v23**：请以 `docs/handbook-v2/README.md` + 本地 docx 为准；本文件为旧版「20 天 / 3 周 docx」对齐稿，日程与 v23 不同。

---

你 JS 基础好，Week 1 压缩。目标：用 **5 天**把 Go 基础语法、HTTP server、struct、分层走一遍。

## Week 1 学习策略

- 不要停下来背语法——遇到不会的随时查  
- 每天任务卡里 **「项目任务」** 是硬性要求（必须写代码）  
- **「学什么」** 是引导资源，看不完不强求  
- 周末可做小 side task：用 Go 重写一个以前写过的 JS 小脚本

---

## Day 1 · Hello Go + 第一个 HTTP endpoint（2–3h）

### 学什么

- Go by Example 前 10 节：Hello World, Values, Variables, Constants, For, If/Else, Switch, Arrays, Slices, Maps  
- 重点：struct 声明 vs JS object、`:=` 和 `var` 的区别  
- `net/http`：`http.ListenAndServe` + `http.HandleFunc`

### 项目任务

- 在 `cmd/server/main.go` 写一个最简 HTTP server  
- 添加 `GET /health`，返回 JSON `{"status":"ok"}`  
- 运行：`go run cmd/server/main.go`  
- 用浏览器或 `curl localhost:8080/health` 验证

### 记忆激活

- JS 里 Express 的 `app.get('/health')`，Go 里就是 `http.HandleFunc`，本质一样

---

## Day 2

> **主跟 v23**：执行稿与代码任务以 `docs/week-1/day2.md` 为准（**Error + Validation**）。  
> 下面保留的是旧「20 天 / 3 周」**Struct + JSON + Query** 原文，便于对照；同一主题深挖见 `docs/week-1/struct.md`。

### 学什么（旧 20 天线原文）

- Go by Example：Structs, Methods, JSON  
- 重点：`json:"name"` struct tag  
- `encoding/json`：Marshal / Unmarshal  
- `net/url` 解析 query params

### 项目任务（旧 20 天线原文）

- 定义 `SearchRequest`：`Q`（query）, `Store`, `Page`  
- 定义 `SearchResponse`：`Results []Product`, `Total int`  
- 添加 `GET /search`，返回 mock 数据  
- 从 `r.URL.Query()` 读取 `q` / `store` / `page` 并打印

### 记忆激活

- Walmart.ca 的 search endpoint 大致也是：handler 接请求、解析 params、返回 JSON

---

## Day 3 · Error Handling 初步 + Validation（2–3h）

### 学什么

- Go by Example：Errors, Panic, Defer  
- 重点：`error` 是 interface；`(result, err)` 多返回值  
- `http.Error()`；`strconv.Atoi` + 错误处理

### 项目任务

- `/search` 参数校验：`q` 不能为空；`page` 必须是正整数  
- 参数错误返回 `400 Bad Request` + error message JSON  
- `store` 为空字符串时默认一个值（如 `"default"`）

### 记忆激活

- 参数校验类 bug：`store` 没传时默认怎么处理

---

## Day 4 · 分层架构：Handler → Service（2–3h）

### 学什么

- Go by Example：Functions, Multiple Return Values  
- Package 基础：`internal/` 的意义  
- 从 `main.go` 拆逻辑出去的模式

### 项目任务

- `internal/service/search.go`：`Search(req SearchRequest) (SearchResponse, error)`  
- `internal/handler/search.go`：`SearchHandler` 调 `service.Search`  
- `main.go` 只负责：创建 service → 注入 handler → 启动 server  
- 重构后整体行为不变，只是结构更整洁

### 记忆激活

- Walmart.ca 的 Go service 一定是分层的；你改过的代码可能在 service 层

---

## Day 5 · Week 1 巩固 + 测试入门（2–3h）

### 学什么

- Go by Example：Testing, Table-driven Tests  
- `testing`：`TestXxx(t *testing.T)`  
- `net/http/httptest`：`NewRequest`, `NewRecorder`

### 项目任务

- 给 `service.Search` 写 2–3 个单测：正常查询、空 query、负数 page  
- 给 `SearchHandler` 写一个 httptest：构造 request → assert status 和 body  
- `go test ./... -v`  
- `go test ./... -cover` 看 coverage

### 记忆激活

- 你写过后端测试——这就是 Go 里对应的一套，标准库比想象简单

---

## Week 1 周末任务（2–3h，可选）

- README 初稿：项目介绍、怎么跑、已有哪些 endpoint  
- 项目 push 到 GitHub（private 也行）  
- 花约 30 分钟看 Go Tour 的 concurrency 部分预习  
- 若轻松：用 Go 重写一个以前用 JS/TS 写的小脚本（如读 CSV 统计）