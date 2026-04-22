# Week 2：并发核心（主战场）

> **若你主跟 v23 手册**：见 `docs/handbook-v2/README.md`。本文件为旧 20 天线。

---


## 心理准备（与 docx 一致）

这是你整个 3 周里**最难的一周**。goroutine + channel 的并发模型和 async/await 是两回事。第一次写出 **deadlock** 和 **panic** 是正常的——调 bug 的过程才是真正学到东西的时候。

**如果某一天卡住了，不要硬撑——第二天回头看，往往就通了。**

这一周的内容 ≈ 面试官 **90%** 会问的 Go 内容，值得多花时间。

---

## Day 6 · Goroutine 基础（2–3h）

### 学什么

- Go by Example：Goroutines, Channels  
- 重点：`go func() {}()`；main goroutine 退出会结束其他 goroutine  
- `sync.WaitGroup`  
- 陷阱：for 循环里启动 goroutine 的闭包问题  

### 项目任务

- `internal/client/`：`catalog.go`, `inventory.go`, `pricing.go`  
- 每个 client：`Fetch(query)`，sleep 100–300ms 后返回 mock 数据  
- `service.Search` 里**串行**调用 3 个 client，看总耗时  
- 测量：总时长 ≈ 三个 client 延迟之和  

### 记忆激活

- Walmart.ca search service 骨架：一个入口调多个下游  

---

## Day 7 · Goroutine + WaitGroup 改并发（2–3h）

### 学什么

- 串行改 goroutine 并发，`WaitGroup` 同步  
- 数据竞争（race）+ `go run -race`  
- 为什么不能无保护地共享变量收集结果  

### 项目任务

- `service.Search` 并发版：3 个 client 同时调  
- `WaitGroup` 等待全部结束  
- `sync.Mutex` 保护共享结果（先简单方式）  
- 测量：总时长 ≈ **最慢**的那个 client  
- `go run -race` 确保无数据竞争  

### 记忆激活

- fan-out/fan-in 模式——Self-Brief 里写过，今天要真写出来  

---

## Day 8 · Channel 取代 Mutex（2–3h）

### 学什么

- Go by Example：Channels, Channel Buffering, Channel Directions  
- 重点：*"don't communicate by sharing memory, share memory by communicating"*  
- buffered vs unbuffered；`close(ch)` 与 `for range ch`  

### 项目任务

- 把 Day 7 的 Mutex 版改写成 **channel** 版  
- 每个 goroutine 把结果塞进 channel，主 goroutine 读齐  
- 对比两种写法可读性  
- 思考题：**何时 Mutex、何时 channel？**（写下答案）  

### 记忆激活

- 面试常问：Mutex 和 channel 怎么选——今天要有体感答案  

---

## Day 9 · Select + 超时模式（2–3h）

### 学什么

- Go by Example：Select, Timeouts, Non-blocking Channel Ops  
- `select`；`time.After` 返回 channel；`select + time.After` = 超时模式  

### 项目任务

- 3 个 mock client 加**慢速模式**（如 20% 概率 sleep 2 秒模拟超时）  
- service 层用 `select`：**500ms** 内没返回则跳过该 client  
- 返回部分结果 + 标记哪些 client 超时（graceful degradation）  

### 记忆激活

- 前端 SLA 500ms、某个下游挂了不能拖死整请求——超时与降级  

---

## Day 10 · Context — Go 的灵魂（2.5–3h）

### 学什么

- Go by Example：Context  
- 重点：`Done`, `Err`, `Deadline`, `Value`；`WithTimeout` / `WithCancel`  
- 为什么 `ctx` 常作为第一个参数；HTTP 里 `r.Context()`  

### 项目任务

- `service.Search` 第一个参数改为 `ctx context.Context`  
- mock client：`Fetch(ctx, query)`  
- handler 入口：`ctx, cancel := context.WithTimeout(r.Context(), 500*time.Millisecond); defer cancel()`  
- client 里：`select { case <-ctx.Done(): return ctx.Err() }`  
- 测试：手动让 mock client sleep 1s，看整条链路是否按 ctx 取消  

### 记忆激活

- Walmart.ca Go 代码里极常见的模式；面试问 context 要能自信答  

---

## Week 2 结束自检（与 docx 一致）

不看文档应能做到：

- 写一个 fan-out/fan-in 的 goroutine 并发调用  
- 用 channel + select 实现超时  
- 让 context 贯穿 handler → service → client  
- 解释 Mutex 和 channel 各自用在什么场景  

**若做不到 → Week 2 的某天重来。不要带着漏洞进 Week 3。**
