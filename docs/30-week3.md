# Week 3：工程打磨 + Kafka + 记忆激活

> **若你主跟 v23 手册**：见 `docs/handbook-v2/README.md`（v23 为 Day 1–23，与本周文件天数不一致）。本文件为旧 20 天线。

---


## Week 3 双重目标（与 docx 一致）

1. 把 mini-search 打磨成可展示的项目（interface + error + 测试完善 + Docker）  
2. 穿插 Kafka 学习（**能用 + 让 Compass 故事成立**的程度）  
3. **关键：Day 15** 做 store filter bug 重现——未来面试的杀手级故事来源  

---

## Day 11 · Interface + 依赖注入（2–3h）

### 学什么

- Go by Example：Interfaces  
- 重点：Go interface **隐式实现**；*"accept interfaces, return structs"*  
- 依赖注入：service 接收 **interface** 而不是具体 struct  

### 项目任务

- 定义 `CatalogClient`, `InventoryClient`, `PricingClient` 三个 interface  
- 具体 struct 实现这些 interface  
- `SearchService` 接收 3 个 interface 参数  
- 好处：测试时可注入 mock，不必真 HTTP  

### 记忆激活

- Walmart.ca 代码里一定有 interface——Go 项目标配结构  

---

## Day 12 · Error Handling 进阶（2–3h）

### 学什么

- Go by Example：Error Wrapping  
- `fmt.Errorf("... %w", err)`；`errors.Is` / `errors.As`  
- 自定义 error：实现 `Error() string`  

### 项目任务

- 自定义 error：`ErrUpstreamTimeout`, `ErrInvalidQuery`, `ErrUpstreamUnavailable`  
- client 层返回具体 error，service 层 `%w` wrap 加上下文  
- handler 层 `errors.Is` 判断，映射 HTTP 状态码（400 / 504 / 503）  
- 写测试验证 error 链可识别  

### 记忆激活

- 面试常问 Go 的 error handling——能讲 **wrap + Is + As** 三件套  

---

## Day 13 · 测试全面覆盖 + Table-driven（2–3h）

### 学什么

- Table-driven tests 是 Go 测试惯用法  
- mock interface 做 service 层隔离测试  
- `httptest` 做 handler 端到端测试  

### 项目任务

- 用 Day 11 的 interface 写 mock client：**成功 / 超时 / 错误** 三种  
- table-driven 写 service 测试：每行一个 case（输入 + 期望输出）  
- handler 集成测试：`httptest.Server` 模拟真实 HTTP  
- **目标：覆盖率 > 70%**（`go test -cover` 验证）  

### 记忆激活

- Walmart.ca 写过测试——这里看到完整 Go 测试生态  

---

## Day 14 · Kafka 基础 + Docker 启动（2.5–3h）

### 学什么

- Kafka：Topic, Partition, Consumer Group, Offset  
- Confluent「Apache Kafka 101」YouTube 前 3 节（约 30min）  
- Go Kafka 库：`kafka-go`（segment.io）或 `confluent-kafka-go`，**推荐 kafka-go（纯 Go）**  

### 项目任务

- 写 `docker-compose.yml` 启动 **Kafka + Zookeeper**（网上抄一份标准配置即可）  
- `docker-compose up -d`，验证 Kafka 运行  
- CLI 创建 topic：`search-log`  
- CLI **生产一条、消费一条**（熟悉工具）  
- `go get github.com/segmentio/kafka-go`  

### 记忆激活

- Compass Map Refresh Pipeline 也是 Kafka——开始有 hands-on  

---

## Day 15 · Store Filter Bug 重现（杀手级故事日）（3h）

### 学什么

- 今天没有新概念——用之前所有东西做刻意练习  
- 目的：重现你记得的真实 bug，变成面试真实素材  

### 项目任务

1. **约 30min**：故意引入 bug——某路径上 **store 参数不透传**到 client  
2. **约 30min**：写测试——模拟「选了 store NY 却返回全国数据」，让测试失败  
3. **约 60min**：debug——log / 断点定位  
4. **约 30min**：修 bug，测试通过  
5. **约 30min**：bugfix 总结（README 里加一节）  

### 记忆激活

- 面试：「我在 mini-search 里重现过当年类似 Walmart.ca 的 store filter 不透传……」真实 + 代码 + 故事  

---

## Day 16 · Kafka Consumer 集成（2.5–3h）

### 学什么

- `kafka-go` 的 Reader API  
- Consumer group + auto commit offset  
- 上下文取消让 consumer **优雅退出**  

### 项目任务

- 新建 `cmd/consumer/main.go`  
- `kafka-go.Reader` 订阅 `search-log` topic  
- 打印每条消息（先不做复杂处理）  
- mini-search 的 handler 里，每次查询后用 `kafka-go.Writer` 发一条到 `search-log`  
- **端到端**：访问 `/search` → consumer 打印出消息  

### 记忆激活

- Compass validator 也是 Kafka consumer——结构类似  

---

## Day 17 · Kafka 面试题 + Offset 策略实验（2–3h）

### 学什么

- at-least-once vs at-most-once（**offset commit 时机**决定）  
- Consumer group 和 partition 的关系  
- Rebalance 何时触发  
- 为什么 **partition 数 = 并行度上限**  

### 项目任务

- **实验 1**：关掉 auto-commit，改手动 commit，观察区别  
- **实验 2**：同一 group 起 **2 个 consumer**，看 partition 怎么分  
- **实验 3**：杀掉一个 consumer，观察 rebalance  
- 把 Kafka 面试必会 **5 个问题**（见《Full-Stack 面试准备提纲》）自己答一遍  

### 记忆激活

- 基本原理 + hands-on 代码 = Compass 故事成立  

---

## Day 18 · Docker 化 + README 收尾（2–3h）

### 学什么

- Go 的 **Dockerfile**（multi-stage build）  
- 补全 `docker-compose.yml`：**mini-search + consumer + kafka** 一起跑  
- Go 项目常规 README 结构  

### 项目任务

- Dockerfile：多阶段构建（builder 编译，最终镜像小）  
- `docker-compose.yml` **一键**起整个系统  
- README 完整版：介绍、架构图、API、本地怎么跑、设计决策  
- 画一张简单架构图（draw.io / Excalidraw 等）  

### 记忆激活

- 「最近我在做一个 Go 项目」——今天能真正拿出来给人看  

---

## Day 19 · Self-Brief 填空 + 面试话术打磨（2–3h）

### 学什么

- 今天不写新代码——回头看 Go Self-Brief 文档  
- 用这 19 天学到的东西填完之前空着的地方  

### 项目任务

- 打开 Go Self-Brief 填空文档，模块 **A–F** 填完  
- 一版**中文** 30 秒自我定位 + 一版**英文**  
- 对着 mini-search 练习「我最近在做一个 Go 项目」**约 5 分钟**讲述  
- **录音一遍**，自己听，找卡壳处  

### 记忆激活

- 真实代码 + 真实学习 + 真实故事收口  

---

## Day 20 · 模拟面试 + 复盘（2–3h）

### 学什么

- 测试学得怎么样  
- 用 Go Self-Brief 里 **10 个面试问题**自测  

### 项目任务

- 自问自答 10 题（来自 Self-Brief 模块 **F**）  
- 录音 **Q1、Q3、Q7、Q8**（关键 4 题）  
- 听录音：停顿多、语无伦次、答不对的地方 → 回去补知识点  
- **20 天回顾**：最骄傲的 **3** 个代码片段、最难的 **2** 个坑、最意外的 **1** 个收获  

### 记忆激活

- 从「Go 跟 JS 差不多」到「能讲清 Go 并发 + 有可展示项目」  

---

## Week 3 结束验收（与 docx「最终检查」一致）

20 天后你应该有：

- **代码**：可跑的 mini-search，在 GitHub 上  
- **技能**：Go 并发 + context + interface + error + testing + Kafka consumer 全部 hands-on  
- **故事**：store filter bug 重现 = 杀手级面试素材  
- **话术**：中英文 30 秒定位 + 5 分钟项目讲述  

## 长期维护（与 docx 一致）

**面试前保持手感**

- 每周 1–2 小时给 mini-search 加小 feature（如 rate limiter、新下游）  
- 每天约 15 分钟读一篇 Go 文章（如 awesome-go、Go Weekly）  
- LeetCode 用 Go 刷，保持语法手感  

**真面试来临时**

- 面试前 **2 天**：重看 Self-Brief 模块 **D、E、F**  
- 面试前 **1 天**：看一遍 mini-search README 和架构图  
- 面试前 **1 小时**：把 **store filter bug** 故事口述一遍  

**最后一条（docx 原意）**

这份表不是「学完 Go」——Go 永远学不完。它是 **够用、真实、能用于面试** 的底牌：  
**「我有一个自己的 Go 项目，这是代码，这是设计决策，这是我踩过的坑。」**  
这句话胜过背八股。
