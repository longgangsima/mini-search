# Go 后端架构复盘：Handler / Service、「消失的 call」、DI（Day 1 补充）

> 与 [day1.md](day1.md) 第一天的项目实践配套：从「第一直觉」对照 **industry standard** 的工程说法。  
> 和 [questions.md](questions.md) 里关于 `NewSearchHandler`、**composition root** 的问答**互补**；本文件偏「总览 + 一张词汇卡」。

---

## 一、角色分工：为什么不把代码写在一起？

**你的疑问**：Controller (Handler) 和 Service 为什么要分家？

**工程说法**：**Separation of concerns（关注点分离）**。

- **Service（业务）**：核心规则与编排（查库、算价、调下游等）。
- **Handler（HTTP 适配层）**：只处理协议——解析 query、header、写状态码、JSON 等；把「HTTP 里来的东西」变成 `struct`，再交给 Service。

**术语**：**Decoupling（解耦）**。以后从网页换成 gRPC/CLI，往往只改接入层，Service 可大量复用。

---

## 二、「消失的 call」：谁调用了 `ServeHTTP`？

**你的疑问**：`main` 里只看到 `mux.Handle("/search", searchH)`，没有手写 `searchH.ServeHTTP()`，它怎么会跑起来？

**工程说法**：**Inversion of control (IoC，控制反转)**，配合 **standard library 里的 HTTP 服务循环**（常口语叫 engine / server loop）。

1. **契约**：`http.Handler` 要求实现 `ServeHTTP(ResponseWriter, *Request)`。你的 `SearchHandler` 通过 **interface satisfaction** 满足契约（不是 Java 的 `implements` 关键字那种写法）。
2. **登记**：`mux.Handle("/search", searchH)` 把「路径前缀 / 你提供的实例」记进**路由表**（内部是 map + 树等实现，此处略）。
3. **谁调用**：`http.ListenAndServe` 内部接受连接、解析 HTTP、根据 path 在路由表上 **dispatch** 到对应 handler，再调用其 `ServeHTTP`。

**术语**：**Runtime dispatch（运行时分发）**。你写的是「登记与实现」，**调用发生在 net/http 运行时**，所以你在 `main` 里「看不见」那行 `ServeHTTP` 字样。

---

## 三、为什么 `main` 要组装依赖、而不是让 Handler 里 `new Service`？

**你的疑问**：为什么不让 Handler 内部自己 `NewSearchService()`？

**工程说法**：**显式依赖**、**dependency injection (DI，依赖注入)**、**composition root（组合根）**。

- **Hidden dependency（坏味道）**：在 handler 里偷偷 new，**调用方不知道依赖图**，测试也难塞 **mock**。
- **做法**：`main`（或一个专门的 `cmd` 层）**唯一负责** new 出 `Service`，再 `NewSearchHandler(svc)` 塞进去。测试里可以 `NewSearchHandler(fakeSvc)`，不必起真数据库。

**术语**：`main` 常被称为 **composition root**——**总装车间**；依赖从这一层流进各包，**透明、可测**。

---

## 四、import 为什么写一整条 module 路径、而不是 `../../`？

**你的疑问**：明明在隔壁目录，为什么不用相对路径跳文件夹？

**工程说法**：**Module path + 包路径** 是 **location independence（与磁盘布局弱绑定）** 的；只要 `go.mod` 的 `module` 不变，import 在仓库里各处一致。相对 `../` 跳文件在 Go 里**不是**常规 import 方式（`go mod` 体系下以 module 为根）。

**术语**：**Go modules** 下的 **canonical import path**；协作、重命名包、多模块拆分时都更稳。

---

## 核心词汇卡片（和「通电」瞬间对齐）


| English term                   | 可以怎么记                                            |
| ------------------------------ | ------------------------------------------------ |
| **Interface satisfaction**     | 有 `ServeHTTP` 就「像」`http.Handler`，不用写 implements。 |
| **net/http server loop**       | 听端口、接请求、查路由、调你实现的 `ServeHTTP` 的那层运行时。            |
| **Inversion of control (IoC)** | 从「我到处 `foo()`」到「我注册好，**运行时**来 `foo()` 我」。        |
| **Dependency injection (DI)**  | 依赖从外头传进来，不在深处偷偷 `new`。                           |
| **Composition root**           | 通常是 `main`：唯一「总装配」处。                             |
| **Separation of concerns**     | Handler 管 HTTP，Service 管业务。                      |


---

## 写在最后

这些「看起来绕路」的拆分，是为了在系统变大时仍能 **换零件式维护**（改接入、换实现、加测试）。和 mini-search 里 **handler → service →（将来）client** 的目录是对齐的：第一天搭的不是语法，是**能长大的骨架**。