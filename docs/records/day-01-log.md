# Day 1 - 2026-04-21

- Today's plan (30-second version):
搭好 mini-search 骨架，跑通 /health 和 /search（mock）。
- What I completed:
  - 配置层 `internal/config/config.go`（Load + getEnv）
  - handler/service/client 基础骨架
  - `cmd/server/main.go` 路由注册（/health, /search）
  - 修复并确认 `os.Getenv` 用法
  - 给主要 Go 文件补了简短注释
- Verification command(s):
  - go build ./...
  - go run ./cmd/server
  - curl -s "[http://127.0.0.1:8080/health](http://127.0.0.1:8080/health)"
  - curl -s "[http://127.0.0.1:8080/search?q=milk](http://127.0.0.1:8080/search?q=milk)"
- Verification result:
  - build 通过
  - /health 返回 {"status":"ok"}
  - /search 返回 {"results":["result for milk"],"total":1}
- Blocker today:
一度不知道 8080 被谁占用；以及 `os.getEnv`/`os.Getenv` 大小写混淆。
- How I resolved it:
查监听进程并结束占用；改为 `os.Getenv`。
- Tomorrow first step (must be executable):
在 handler 里解析 `store` 和 `page`，并扩展 `SearchRequest` 参数传递。
- One Go concept reinforced today:
Go 导出符号首字母大写；环境变量读取用标准库 `os.Getenv`。

---

## Deep dive notes（可选，长文放 week-1）

- 指针、`&` / `*`、struct value vs pointer、**escape analysis**、map **comma ok**：已整理到  
`docs/week-1/memory-pointers-escape-and-map.md`  
- Week 1 全部笔记索引与去重说明：`docs/week-1/README.md`  
- 以后同类「概念长文」默认放 `docs/week-1/`，`records` 里只留链接 + 一句话，避免 records 变成第二份讲义。