# Go 23 天 Day-by-Day 执行手册 · 仓库骨架（做法 B）

本目录是 **《Go 23天 Day-by-Day 执行手册》docx** 的「结构版」：只保留 **Phase / Day 标题、时长、完成标志、产出路径**，**不**粘贴大段代码。详细步骤、代码片段、Senior 视角与常见坑 **以你本机 docx 为准**。

## 真源文档（权威）


| 文档                             | 用途                                          | 建议路径（本机）                                  |
| ------------------------------ | ------------------------------------------- | ----------------------------------------- |
| 《Go 23 天学习计划 v2 - mini-search》 | 总览：决策、参数、概念依赖、JD 对齐，**读 1 次**               | 与你的 v2 总览 docx 同目录                        |
| 《Go 23天 Day-by-Day 执行手册》       | **每天照做**：学什么 / 主项目任务 / Senior 视角 / 坑 / 完成标志 | `~/Downloads/Go 23天 Day-by-Day 执行手册.docx` |


Day 0 环境准备以 **总览 docx** 里的 Day 0 为准；本仓库另有 `docs/05-day0.md`（旧 20 天对齐版），若你以 **v23** 为主，请优先跟总览 docx。

## 使用方式（与手册一致）

- **Day 1 起**：每天只打开 **Day-by-Day docx**，找到对应 Day，按卡片执行。  
- **卡壳超过 30 分钟**：记下来，跳过或改天补；**某天没做完 → 整体顺延 1 天**，不必自责。  
- **只有 1 小时**：手册建议只做 **Step 1 + Step 2**，后面跳过。

时间预算：**每天 2–3 小时**；**Day 1 为长日 3–4 小时**。

---

## Phase 1 · 骨架搭建（Day 1–3）

**Phase 目标**：项目骨架一次搭好，之后每天填肉、不大重构。  
**Phase 结束时你有**：完整目录（4 层 `internal`）、`/health` + `/search`、单测 + httptest、简陋可跑的 **React 前端 dev tool**（Vite）。


| Day   | 主题                    | 预计时长           | 完成标志（摘要）                                                                                                 |
| ----- | --------------------- | -------------- | -------------------------------------------------------------------------------------------------------- |
| **1** | 长日 · 建完整框架            | 3–4h           | `curl /health`、`curl /search?q=test` OK；`internal/{config,handler,service,client}` 齐；`go build ./...` 无错 |
| **2** | Error + Validation 巩固 | 2–3h           | 空 `q` / 负 `page` → 400；正常 200；handler 入口 log + `defer log` 可见                                            |
| **3** | 测试入门 + 前端初版（下午）       | 2–3h + 下午 2–3h | `go test ./...` 全过；覆盖率 **>70%**；前端 Search Playground / Health Dashboard 能打通（含 CORS）                      |


---

## Phase 2 · 并发引擎（Day 4–8）

**Phase 目标**：goroutine、channel、context 贯通；graceful degradation；graceful shutdown。  
**Phase 结束时你有**：3 mock client 并发 fan-out/fan-in；`ctx` 全链路；慢下游不拖垮；**Ctrl+C 优雅退出**。


| Day   | 主题                                      | 预计时长   | 完成标志（摘要）                                                               |
| ----- | --------------------------------------- | ------ | ---------------------------------------------------------------------- |
| **4** | Goroutine + WaitGroup + Mutex           | 2–3h   | 3 client 并发；总延迟 ≈ 最慢单路；3 段结果；测试仍绿                                      |
| **5** | Channel + race detector                 | 2–3h   | `-race` 干净；Channel 与 Mutex 行为一致；能简述 Mutex vs Channel                   |
| **6** | Select + timeout + graceful degradation | 2–3h   | 正常 3 结果；超时部分结果 + `degraded`；总延迟 ≤ 约 500ms                              |
| **7** | Context + Graceful shutdown ⭐           | 2.5–3h | `ctx` 贯穿 handler→service→client；慢下游取消正确；**Ctrl+C 优雅退出**；能简述 context    |
| **8** | 巩固日 + 前端对接                              | 2–3h   | `docs/adr/001-concurrency-model.md` 写好；前端展示 `degraded`；能约 30s 讲清项目在做什么 |


---

## Phase 3 · 工程化（Day 9–14）

**Phase 目标**：从「能跑」到 **production-grade**（错误链、结构化日志、DI、PG、API 与 ADR）。  
**Phase 结束时你有**：`%w` + slog；interface + mock 测试；PostgreSQL + repository；`**docs/api.md` + 3 份 ADR**。


| Day    | 主题                                 | 预计时长 | 完成标志（摘要）                                                                     |
| ------ | ---------------------------------- | ---- | ---------------------------------------------------------------------------- |
| **9**  | Error wrapping + slog              | 2–3h | 不同 error → 不同 HTTP 码；日志 JSON；能简述 `errors.Is` / `As`                          |
| **10** | Interface + DI + mock 测试           | 2–3h | Service 吃 interface；mock 覆盖成功/超时/错误；`go test ./...` 全过                       |
| **11** | PostgreSQL + Repository            | 2–3h | `search_logs` 有数据；`internal/repository` 独立；`go build` OK                     |
| **12** | PG 测试 + N+1 + 索引                   | 2–3h | 集成测试过；能讲 Seq Scan vs Index Scan；能讲 N+1 与解法                                   |
| **13** | System Design Day（ADR + API Spec）⭐ | 2–3h | `docs/api.md`；ADR 001–003 齐；README 链到文档；能约 1min 讲一则 ADR                      |
| **14** | 巩固日                                | 2–3h | 覆盖率达标（手册给 service/handler/repo 分档）；`go vet`、`-race`、`-cover`、gofmt 过关；代码自审顺眼 |


---

## Phase 4 · 事件驱动（Day 15–19）

**Phase 目标**：Kafka 深度 + 可讲的 debug 故事。  
**Phase 结束时你有**：compose 起 Kafka（及 PG）；consumer + producer；schema；worker pool；retry/DLQ + 实验；`**docs/debug-story-001.md`**。


| Day    | 主题                                         | 预计时长          | 完成标志（摘要）                                                                  |
| ------ | ------------------------------------------ | ------------- | ------------------------------------------------------------------------- |
| **15** | Kafka 概念 + compose + Consumer              | 2.5–3h + 下午前端 | compose 起 Kafka+PG；consumer 能收 console 消息；前端 Event Log 能拉 DB（如 `/events`） |
| **16** | Message schema + Producer 集成               | 2–3h          | `/search` → DB + Kafka；consumer 能解析结构化 event                              |
| **17** | Consumer worker pool ⭐                     | 2–3h          | 手册要求 worker pool 与手册内验收（见 docx）                                           |
| **18** | Failure handling + DLQ + consumer group 实验 | 2–3h          | Retry + DLQ 走通；观察到 rebalance；能约 2min 讲 Kafka 高频 5 题                       |
| **19** | Bug 练习日（参数透传）⭐                             | 2–3h          | 测试防回归；`docs/debug-story-001.md` 完整；能讲 Senior 调试五步                         |


---

## Phase 5 · 打磨 + 展示（Day 20–23）

**Phase 目标**：能见人 + **约 15 分钟讲清楚**。  
**Phase 结束时你有**：全栈 compose；README + 架构图 + `/metrics`；中英文讲述录音；10 题自答；`**docs/learning-journal.md`**。


| Day    | 主题                     | 预计时长        | 完成标志（摘要）                                                |
| ------ | ---------------------- | ----------- | ------------------------------------------------------- |
| **20** | Docker 化 + compose 全栈  | 2–3h + 下午前端 | `docker-compose up -d --build` 一键起；前端截图 + 约 60s demo 视频 |
| **21** | README + 架构图 + metrics | 2–3h        | `/metrics` 可见计数；README 完整；`docs/architecture.png` 存在    |
| **22** | Self-Brief + 5 分钟讲述录音  | 2–3h        | Self-Brief A–F 填完；中英讲稿 + 录音满意；GitHub README 可展示         |
| **23** | 模拟面试 + 复盘              | 2–3h        | 10 题录音与自评；`docs/learning-journal.md` 写完；明确下月 1–2 件事     |


**Day 23 最终验收（手册原文摘要）**：10 题录音、`learning-journal`、public repo + README、知道面试前怎么复习。

---

## 与仓库其他文档的关系


| 路径                                     | 说明                                                                             |
| -------------------------------------- | ------------------------------------------------------------------------------ |
| `docs/handbook-v2/README.md`           | **本文件**：v23 骨架索引（做法 B）                                                         |
| `docs/05-day0.md`、`docs/10-week1.md` 等 | **旧版**：与《Go 3周…20 天》docx 对齐；若你主跟 **v23**，以本目录 + Day-by-Day docx 为准，旧周计划可作对照或忽略 |
| `docs/worklog/`                        | 滚动交接：`HISTORY.local.md` / `CURRENT.local.md`；模板与 **Kafka 核对表** 在同目录 `*.md`；**Day 18** 实验笔记可写进 `kafka-reality-checklist.md` |


---

## 手册里会创建的重要路径（备查）

执行过程中 docx 会点名这些产物（不必一天全有）：

- `docs/adr/001-concurrency-model.md`（及 002、003）  
- `docs/api.md`  
- `docs/debug-story-001.md`  
- `docs/learning-journal.md`  
- `docs/architecture.png`  
- `docs/screenshots/`、`frontend/`（AI dev tool）

具体命令与代码 **只打开 docx 跟做** 即可。