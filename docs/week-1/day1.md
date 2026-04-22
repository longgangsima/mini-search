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

下面 import 路径请与根目录 **`go.mod`** 中 `module` 一行保持一致。

**Step 3 代码以仓库真源为准**（避免与文档双份维护、产生漂移）：

| 路径 | 作用 |
|------|------|
| `internal/config/config.go` | 读 `PORT` 等；包内 `getEnv` 内部必须用标准库 **`os.Getenv`**（**G 大写**），不要写 `os.getEnv`。 |
| `internal/client/client.go` | `UpstreamClient` 接口占位。 |
| `internal/service/search.go` | `Search` mock 与类型定义。 |
| `internal/handler/search.go` | `/search` 与 `/health` 的 HTTP 处理。 |
| `cmd/server/main.go` | `Load`、组装 handler、注册路由、`ListenAndServe`。 |

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

## Summary:

### Q: NewSearchHandler 这玩意为什么要写在handler 里而不是在main 里写？我感觉他跟handler 没半毛钱关系

#### **1. 什么是 Encapsulation（封装）？**

在工程上，我们遵循 **"Tell, Don't Ask"** 原则。

- **如果写在** `main` **里**：`main` 函数必须知道 `SearchHandler` 内部所有的细节。比如它需要哪个 `service`？它需要什么初始参数？如果以后 `SearchHandler` 增加了一个新的依赖项（比如需要一个日志记录器 `Logger`），你得去 `main.go` 里改代码。
- **写在** `handler` **里**：`NewSearchHandler` 就像是这个组件的**安装说明书**。`main` 只需要说“我要一个 SearchHandler”，至于这个 Handler 内部是怎么组装的，`main` 完全不需要关心。

**Engineering Term: Information Hiding（信息隐藏）**。我们希望 `main` 越简单越好，它只负责启动，不负责研究每个零件怎么造。

---

#### **2. 为什么说它跟 handler “有半毛钱关系”？**

其实，`NewSearchHandler` 是 `SearchHandler` 这个结构体的 **Constructor（构造函数）**。

在 Go 的工程习惯中，数据结构（Struct）和它的初始化函数（New 函数）必须**绑定**在一起。

- **Access Control（访问控制）**：  
如果你看 `SearchHandler` 的定义：
**Go**
  ```
  type SearchHandler struct {
      svc *service.SearchService // 注意：这里的 svc 是小写的
  }

  ```
  在 Go 中，**小写开头的变量是私有的（Private）**。这意味着 `main.go` 根本无法直接给 `svc` 赋值。只有同一个包（`handler` 包）里的 `NewSearchHandler` 才有权限操作这个私有字段。

**Engineering Logic**: 这样做是为了防止 `main` 这种外部模块乱动 `handler` 内部的零件，保证了 **System Integrity（系统完整性）**。

---

#### **3. Dependency Injection（依赖注入）的入口**

`NewSearchHandler` 的参数 `(svc *service.SearchService)` 是实现 **Decoupling（解耦）** 的关键。

- **The Pattern**: 这叫 **Constructor Injection**。
- **Logic**: `handler` 自己不产生 `service`，而是等着 `main` 把 `service` 传进来。`NewSearchHandler` 就是那个“插口”。
- **Benefit**: 如果以后你想测试这个 `handler`，你可以在测试代码里调用 `NewSearchHandler` 并传一个“假的 Service”进去。如果初始化逻辑写死在 `main` 里，你的测试代码就很难模拟这个过程。

---

#### **4. 总结对比：工程思维 vs. 简单思维**


| **维度**  | **写在** `main` **里 (Quick & Dirty)** | **写在** `handler` **里 (Engineering Standard)** |
| ------- | ----------------------------------- | --------------------------------------------- |
| **归属感** | 逻辑散乱，`main` 变得臃肿。                   | **High Cohesion**（高内聚），逻辑都在家门口。               |
| **安全性** | 必须把字段设为公开（大写），容易被乱改。                | 可以使用私有字段，实现完美的 **Encapsulation**。             |
| **维护性** | 修改 Handler 依赖时，必须改动 `main`。         | 只需要改动 `handler` 包内部，`main` 无感知。               |
| **术语**  | Hard-coded Initialization           | **Factory Pattern / Constructor**             |


**一句话总结：**  
`NewSearchHandler` 虽然不干业务活，但它是 `SearchHandler` 的**出生证明**和**唯一合法的造币厂**。把它放在 `handler` 里是为了保护内部数据不被乱动，并让 `main` 保持简洁。

你觉得这种“自己的事情自己管，不麻烦主程序”的思维，是不是让代码模块看起来更像一个独立的“黑盒子”？

## 如果按你怎么说那main 不应该import service 啊？

### **1. Main 的真实角色：The Assembler（组装车间）**

在 **Professional Software Architecture**（专业软件架构）中，`main.go` 有一个特殊的身份，叫做 **Object Graph Composer**（对象图合成器）。

- **Engineering Logic**: 想象你在组装一台电脑。
  - `handler` 是显示器，`service` 是主机。
  - **显示器（handler）** 确实不应该自己去生产一台 **主机（service）**，否则它俩就死死地焊在一起了（Tight Coupling）。
  - 但是，总得有一个人把显示器的线插到主机上吧？
- **Main 的职责**: `main` 就是那个**拿线的人**。它不关心主机内部怎么运行，也不关心显示器怎么显示。它的唯一任务是：**实例化（New）所有零件，并将它们“接”在一起。**

所以，`main` 必须 `import service`，是因为它是**组装现场**。

---

### **2. 为什么不在 Handler 内部直接 New？ (The "Anti-Pattern")**

如果你在 `handler` 内部直接调用 `service.NewSearchService()`，会发生什么？

- **Violating Dependency Inversion（违反依赖倒置）**: `handler` 此时不仅依赖 `service` 的接口，还依赖了它的具体创建过程。
- **Testing Blockade（测试阻碍）**: 如果你想单独测试 `handler`，由于它内部写死了 `NewSearchService`，你无法在测试时偷偷换成一个“假 Service”。
- **Engineering Logic**: 通过在 `main` 里创建 `svc` 再传给 `handler`（这就是 **Dependency Injection / 依赖注入**），`handler` 变得非常“纯粹”。它只需要知道：“我手里有个能干活的 Service 就行，至于它是谁给我的，我不关心。”

---

### **3. 隐藏了什么“细节”？ (What is hidden?)**

虽然 `main` 知道了 `service` 的存在，但 `NewSearchHandler` 依然隐藏了**组装的工艺细节**：

假设未来的逻辑变复杂了：

**Go**

```
// 现在的 handler 内部
func NewSearchHandler(svc *service.SearchService) *SearchHandler {
    return &SearchHandler{
        svc: svc,
        cache: cache.NewLocalCache(), // 细节 A：增加了缓存
        stats: metrics.NewStats(),     // 细节 B：增加了埋点统计
    }
}

```

- **细节隐藏**: 在这种情况下，`main` 依然只需要传一个 `svc` 进去。
- **Maintenance Efficiency（维护效率）**: 至于 `handler` 内部是不是偷偷加了“缓存”或“统计”，`main` 完全不需要知道，也不需要为此修改代码。

---

### **4. 终极目标：Interface-Based Injection（基于接口的注入）**

你现在的代码里，`main` 传的是具体的 `*service.SearchService`。在更高级的工程实践中，我们会更进一步：

1. `handler` 定义一个接口：`type Searcher interface { ... }`。
2. `NewSearchHandler` 接收这个接口：`func NewSearchHandler(s Searcher)`。
3. `main` 依然负责组装。

**这样做 main 依然要 import 吗？**  
是的。因为 `main` 是整个系统的 **Constructor（构造者）**。它必须认识所有的零件，才能把零件发给各个部门。

---

### **总结术语 (Engineering Summary)**


| **Term**                       | **含义** | **工程目的**                                   |
| ------------------------------ | ------ | ------------------------------------------ |
| **Object Composition**         | 对象组合   | 在程序启动时，将各个独立的模块拼装成一个完整的系统。                 |
| **Dependency Injection (DI)**  | 依赖注入   | 零件不由使用者生产，而由第三方（Main）提供。                   |
| **Entry Point Responsibility** | 入口职责   | `main` 作为唯一合法的全局组装中心。                      |
| **Loose Coupling**             | 松耦合    | `handler` 和 `service` 互不生产，只通过“接口”或“注入”通信。 |


**一句话总结：**  
`main` 就像是**总司令部**，它必须知道所有的**兵种（import service/handler）**，然后把**坦克（service）分配给坦克兵（handler）**。`handler` 不需要知道坦克是怎么造出来的，但总司令必须知道，否则没人能把坦克交到士兵手里。

这种“在最顶层（Main）进行组装，在底层（Handler/Service）保持纯粹”的思维，正是构建 **Scalable Systems**（可扩展系统）的基石。你觉得这种“总指挥”的逻辑好理解了吗？