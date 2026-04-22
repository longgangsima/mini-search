# mini-search

Project-driven Go learning project。

**主执行线（推荐）**：**Go 23 天 v2** —— 见 `docs/handbook-v2/README.md`（做法 B：骨架索引，细节在本地《Go 23天 Day-by-Day 执行手册》docx + 《Go 23 天学习计划 v2》总览 docx）。

**旧版（20 天，对齐《Go 3周学习计划 - mini-search》docx）**：`docs/05-day0.md`、`docs/10-week1.md`、`docs/20-week2.md`、`docs/30-week3.md` 仍可作对照；若你只跟 v23，以 `handbook-v2` 为准即可。

## Quick Start

```bash
go run ./cmd/server
```

In another terminal:

```bash
curl "http://127.0.0.1:8080/health"
```

Expected response:

```json
{"status":"ok"}
```

## Learning Docs

**v23（当前推荐）**

- `docs/handbook-v2/README.md`：**23 天 Phase / Day 骨架 + 完成标志**（与 Day-by-Day docx 一致，无大段代码）

**总览与旧版 20 天**

- `docs/00-overview.md`：项目总览（偏旧 3 周叙事，可选读）  
- `docs/01-execution-rules.md`：执行习惯与完成标准  
- `docs/05-day0.md`：Day 0（旧 20 天线）  
- `docs/10-week1.md` … `docs/30-week3.md`：旧 20 天线逐周  
- `docs/50-interview-alignment.md`：面试映射（补充）

## Records Docs

- `docs/records/00-index.md`: records documentation index
- `docs/records/daily-copy-template.md`: daily copy template (recommended + minimal)
- `docs/records/weekly-review-template.md`: weekly review template
- `docs/records/kafka-reality-checklist.md`: verify real Kafka implementation details before interviews

## Project Structure

- `cmd/server`: HTTP server entrypoint
- `cmd/consumer`: Kafka consumer entrypoint (stub for Week 3)
- `internal/handler`: HTTP handlers
- `internal/service`: business logic
- `internal/client`: downstream client mocks
- `internal/model`: request and response models
- `internal/errors`: custom error definitions
- `test`: integration tests

## Suggested Workflow

1. 打开 `docs/handbook-v2/README.md` 确认当天 **Phase / Day**，再打开 **《Go 23天 Day-by-Day 执行手册》docx** 做当天卡片。  
2. 只完成当天的「主项目任务」；可选任务视时间。  
3. 每天最后 10 分钟写 **Go Self-Brief**；可选 `docs/records/daily-copy-template.md` 打卡。  
4. 验证通过后再 commit（`go test`、`curl`、或手册当日「完成标志」）。

# mini-search
