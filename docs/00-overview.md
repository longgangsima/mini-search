# Go 3 周压缩学习计划 · mini-search 总览

> **推荐**：当前仓库的 **主执行索引** 为 `docs/handbook-v2/README.md`（**Go 23 天 v2**，与《Go 23天 Day-by-Day 执行手册》docx 一致）。下文为旧「3 周叙事」总览，仍可作背景阅读。

面向：Jiulong Lin · 每天 2–3 小时 · Mac 开发。

## 这份课程表怎么用

每天打开对应 Day 卡片，按三个部分走：

- **学什么**：今天要掌握的 Go 概念 + 推荐资源  
- **项目任务**：加到 mini-search 项目里的具体代码  
- **记忆激活**：今天这个概念对应 Walmart.ca 的什么场景

**v23 按天**：`docs/handbook-v2/README.md` + 本地 Day-by-Day docx。  
**旧 20 天线**：`docs/05-day0.md`、`docs/10-week1.md`、`docs/20-week2.md`、`docs/30-week3.md`。

## 核心原则

1. 每天都要写代码，哪怕 20 行——阅读不算学习
2. 遇到卡壳 → 查 Go by Example、gobyexample.com、pkg.go.dev
3. 每天最后 10 分钟，在 Go Self-Brief 文档里写一行激活笔记
4. 3 周结束，你应该有一个可以 push 到 GitHub 的 mini-search service

## 项目目标（3 周后你将拥有）

一个 Go 写的 HTTP search service，包含：

- HTTP API：`GET /search?q=...&store=...&page=...`
- 并发调用 3 个 mock 下游（catalog / inventory / pricing）
- goroutine + channel fan-out/fan-in
- context 贯穿的超时控制
- Handler / Service / Client 三层架构
- 结构化 error handling + graceful degradation
- Unit + integration 测试
- Kafka consumer（消费 search log events）
- Docker 化部署

## 时间线

- **Week 1**：语法快速过 + HTTP 基础（Day 1–5，周末可选巩固）  
- **Week 2**：并发核心（Day 6–10）  
- **Week 3**：工程打磨 + Kafka + 记忆激活（Day 11–20）

## 成功标准（自检）

- 能本地跑通并演示关键 API  
- 能讲清并发模型（goroutine、channel、Mutex 选型）与 context 超时取消  
- 能用测试复现并讲清 store filter bug 故事  
- 能完成中英文 30 秒定位 + 约 5 分钟项目讲述

