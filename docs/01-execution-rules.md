# 执行规则（与《Go 3周学习计划 - mini-search 项目》docx 一致）

## 每日节奏

- 工作日：每天约 **2–3 小时**（与 docx 一致）  
- 周末 Week 1：可选 **2–3 小时** 巩固任务（见 Week 1 文末）

## 每天怎么安排（与 docx 卡片一致）

每天按 **三块** 走（顺序随意，但三块都要覆盖）：

1. **学什么**
2. **项目任务**（硬性要求：必须写代码）
3. **记忆激活**

每天最后 **10 分钟**：在 **Go Self-Brief** 文档里写一行激活笔记。

## 核心原则（与 docx 相同）

1. 每天都要写代码，哪怕 20 行——阅读不算学习
2. 遇到卡壳 → 查 Go by Example、gobyexample.com、pkg.go.dev
3. 每天最后 10 分钟，在 Go Self-Brief 文档里写一行激活笔记
4. 3 周结束，你应该有一个可以 push 到 GitHub 的 mini-search service

## 每日完成标准（便于你打卡）

一天算「完成」建议同时满足：

- 当天 **项目任务** 有代码落库  
- 能用命令验证：`go run`、`curl` 或 `go test` 至少一种  
- Self-Brief 或在 `docs/worklog/HISTORY.local.md` 里有一条极简会话记录（二选一或都写）

## 编码约定（不改动 docx 日程，仅避免走偏）

- 业务逻辑不要堆在 `cmd/server/main.go` 里；分层按各 Week 的 Day 推进。  
- **context 贯穿**：按 docx **Week 2 Day 10** 的任务卡执行（handler → service → client）。Day 10 之前不要求提前强行改全链路签名。  
- 错误处理：显式处理 `error`，进阶部分按 Week 3 Day 12。

## 卡住怎么办

- 某一天卡住 → docx 原话：**不要硬撑，第二天回头看，往往就通了**  
- Week 2 结束自检做不到 → **Week 2 某天重来，不要带着漏洞进 Week 3**

