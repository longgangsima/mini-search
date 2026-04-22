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
go mod init github.com/jiulonglin/mini-search
```

本仓库已用模块路径：`github.com/jiulonglin/mini-search`。

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

下面 import 已按本仓库 `go.mod` 写成 `github.com/jiulonglin/mini-search/...`。

`**internal/config/config.go**`

- 小写 `**getEnv**` 是你**本包里的辅助函数**，内部应调用标准库的 `**os.Getenv`**（**G 大写**）。  
- **不要**写 `os.getEnv`，标准库里不存在，会报 `undefined`。

```go
package config

import (
	"os"
	"time"
)

type Config struct {
	Port            string
	UpstreamTimeout time.Duration
}

func Load() Config {
	return Config{
		Port:            getEnv("PORT", ":8080"),
		UpstreamTimeout: 500 * time.Millisecond,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
```

`**internal/client/client.go**`（interface 占位）

```go
package client

import "context"

// UpstreamClient 下游通用接口；Day 4 再写真实实现。
type UpstreamClient interface {
	Fetch(ctx context.Context, query string) (string, error)
}
```

`**internal/service/search.go**`

```go
package service

import "context"

type SearchRequest struct {
	Query string `json:"query"`
	Store string `json:"store,omitempty"`
	Page  int    `json:"page,omitempty"`
}

type SearchResponse struct {
	Results []string `json:"results"`
	Total   int      `json:"total"`
}

type SearchService struct{}

func NewSearchService() *SearchService {
	return &SearchService{}
}

func (s *SearchService) Search(ctx context.Context, req SearchRequest) (SearchResponse, error) {
	_ = ctx // Day 1 占位；Day 7 起贯穿取消与超时
	return SearchResponse{
		Results: []string{"result for " + req.Query},
		Total:   1,
	}, nil
}
```

`**internal/handler/search.go**`

```go
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jiulonglin/mini-search/internal/service"
)

type SearchHandler struct {
	svc *service.SearchService
}

func NewSearchHandler(svc *service.SearchService) *SearchHandler {
	return &SearchHandler{svc: svc}
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := service.SearchRequest{
		Query: r.URL.Query().Get("q"),
	}

	resp, err := h.svc.Search(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

type HealthHandler struct{}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
```

`**cmd/server/main.go**`

```go
package main

import (
	"log"
	"net/http"

	"github.com/jiulonglin/mini-search/internal/config"
	"github.com/jiulonglin/mini-search/internal/handler"
	"github.com/jiulonglin/mini-search/internal/service"
)

func main() {
	cfg := config.Load()

	svc := service.NewSearchService()
	searchH := handler.NewSearchHandler(svc)
	healthH := &handler.HealthHandler{}

	mux := http.NewServeMux()
	mux.Handle("/search", searchH)
	mux.Handle("/health", healthH)

	log.Printf("listening on %s", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}
```

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