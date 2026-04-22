# Day 1 · 长日 · 建完整框架（v23，约 3–4h）

与 `docs/handbook-v2/README.md` 中 Phase 1 Day 1 一致：第一天把 **handler / service / client / config** 四层骨架搭好，后面在骨架上填肉。

---

## 学什么

- Go by Example 前约 15 节（到 Methods / Interfaces 之前可停）：变量、控制流、多返回值、slice/map、struct、JSON  
- `internal/`：其他 module 无法 import，适合放实现细节  
- `net/http`：`http.ServeMux`、`ListenAndServe`、实现 `http.Handler` 的 `ServeHTTP`

---

## 主项目任务

### Step 1：环境确认（约 15 min）

```bash
go version   # go1.22+
```

若从零建目录（你已有 `go-pro` 可跳过 `mkdir`）：

```bash
cd ~/Projects
mkdir mini-search && cd mini-search
go mod init github.com/longgangsima/mini-search
```

本仓库已用模块路径：以根目录 `go.mod` 为准（当前为 `github.com/longgangsima/mini-search`）。

---

### Step 2：目录结构（约 10 min）

目标树：

```text
mini-search/
├── go.mod
├── cmd/server/main.go
├── internal/config/config.go
├── internal/handler/search.go
├── internal/service/search.go
├── internal/client/client.go
├── .env.example
└── README.md
```

```bash
mkdir -p cmd/server internal/config internal/handler internal/service internal/client
touch cmd/server/main.go internal/config/config.go internal/handler/search.go \
  internal/service/search.go internal/client/client.go
touch .env.example README.md
```

---

### Step 3：五个文件（约 150 min）

下面 import 路径请与根目录 `**go.mod**` 中 `module` 一行保持一致。

**Step 3 代码以仓库真源为准**（避免与文档双份维护、产生漂移）：


| 路径                           | 作用                                                                         |
| ---------------------------- | -------------------------------------------------------------------------- |
| `internal/config/config.go`  | 读 `PORT` 等；包内 `getEnv` 内部必须用标准库 `**os.Getenv`**（**G 大写**），不要写 `os.getEnv`。 |
| `internal/client/client.go`  | `UpstreamClient` 接口占位。                                                     |
| `internal/service/search.go` | `Search` mock 与类型定义。                                                       |
| `internal/handler/search.go` | `/search` 与 `/health` 的 HTTP 处理。                                           |
| `cmd/server/main.go`         | `Load`、组装 handler、注册路由、`ListenAndServe`。                                   |


若需要「可打印的讲义版」单文件，仍以本仓库 `internal/`、`cmd/` 为准复制，不要从旧版 docx 表格里粘半行代码。

---

### Step 4：跑通与验证（约 15 min）

```bash
go mod tidy
go run ./cmd/server
```

另一终端：

```bash
curl -s "http://127.0.0.1:8080/health"
curl -s "http://127.0.0.1:8080/search?q=milk"
```

期望：`health` 返回带 `"ok"` 的 JSON；`search` 返回 `result for milk` 类 mock。

---

## 可选（30–45 min）

- `.env.example` 列出 `PORT` 等  
- `Makefile`：`make run` / `make build`  
- README：项目名 + 一句话 + 如何 `go run`

---

## Senior 视角（话术提示）

- Day 1 就四层分包，减少后面大重构  
- 业务入口函数带 `context.Context`，符合常见 Go 服务约定  
- Config 从环境变量读，贴近 12-factor

---

## 常见坑

- `**os.Getenv**`：标准库只有大写 **G**，没有 `os.getEnv`  
- `go mod init` 路径尽量与将来 GitHub 一致  
- JSON 序列化到对外的字段需 **大写开头**（exported）  
- `SearchService` 方法接收者一般用 `*SearchService`（单例式服务）

---

## 完成标志

- `curl /health` 与 `curl '/search?q=test'` 行为正确  
- `internal` 下 **config、handler、service、client** 四个子包齐全  
- `go build ./...` 无错误

---

## 记忆激活

Express 的 `app.get('/health')` 与 Go 里注册 `Handler` 到 `ServeMux` 一样，都是「路由 → 处理函数 → 写 response」。

## 延伸阅读（架构总览）

谁调 `ServeHTTP`、为什么要 **DI**、**composition root**、长 `import` 路径：见 [architecture-web-handler-ioc.md](architecture-web-handler-ioc.md)。

**下一日（v23 Day 2）**：[day2.md](day2.md)（Error + Validation）。

## Summary（与 architecture 合并维护）

原先写在本节的长问答（`NewSearchHandler` 为何在 `handler` 包、`main` 为何仍 `import service`、DI、封装与对比表）与 [architecture-web-handler-ioc.md](architecture-web-handler-ioc.md) **重复**，已并入该文件的 **§三～§五** 统一维护，避免两处改两份。

若你关心的是「谁调 `ServeHTTP`、IoC、词汇卡」，仍从该文件开头读即可；本 `day1.md` 保留课表与验收，原理以专题文档为真源。