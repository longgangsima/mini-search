# Week 1:

### Slice

详见 [slice.md](slice.md)。

### Map

详见 [maps.md](maps.md)。

### struct

详见 [struct.md](struct.md)。

### 环境变量

详见 [环境变量.md](环境变量.md)。

### 指针 / 逃逸 / map comma ok 总览

详见 [memory-pointers-escape-and-map.md](memory-pointers-escape-and-map.md)（详细 map 仍以 `maps.md` 为准）。

## Questions

### Q: 什么是 gRPC？ (The Communication Protocol)

**gRPC** 的全称是 **Google Remote Procedure Call**。

在传统的网页开发中，我们常用 **REST API**（就像你发短信给别人，对方回你短信）。而 **gRPC** 是一种更硬核、更高效的通信方式，它让调用“远程服务器上的函数”变得像调用“自己本地的函数”一样简单。

#### 核心术语拆解：

1. **Protocol Buffers (Protobuf)**: 这是 gRPC 的 **Serialization（序列化）** 机制。它不像 JSON 那样是人类能读懂的文本，它是**二进制数据**。就像把一封长信压缩成了一个极小的快递包，传输速度极快。
2. **HTTP/2**: 这是它的底座。它支持 **Multiplexing（多路复用）**，意味着一条连接可以同时传一堆数据，不用排队。
3. **Language Agnostic（语言无关性）**: 你可以用 Go 写服务端，别人用 Java 或 Python 写客户端，gRPC 像翻译官一样让它们无缝对接。

### Q: 为什么接口定义（Interface）让你听不懂？

听不懂很正常，因为 **Interface（接口）** 并不是一段“执行逻辑”，它是一份 **Behavioral Contract（行为契约）**。

我们用**“电脑维修”**这个工程场景来理解你代码里的 `UpstreamClient`：

#### 1. 场景模拟

假设你在设计一台电脑。你想给电脑加一个“散热器”。

- **如果不使用 Interface（低级写法）**：你在主板上焊死了一个“小明牌散热器”。结果 Day 4 小明倒闭了，你想换个“猫头鹰牌”，对不起，主板废了，得重焊。
- **如果使用 Interface（工程思维）**：你在主板上只定义了一个 **“散热器接口”**（就像 CPU 插座）。

**Interface** 是为了让你的代码**“可替换”**；而 gRPC 是为了让你的数据传输**“极速且标准”**。

### Q:  为什么要return 回来一个service 的地址呢？`*SearchService`)？

在 Go 中，当你传递一个变量时，默认是 **Pass-by-value（值传递/副本传递）**。如果你返回的是对象本身，Go 会在内存中把整个对象复制一遍。

#### A. Performance & Memory Efficiency（性能与内存效率）

- **The Problem**: 如果你的 `SearchService` 结构体很大（比如里面包含了很多配置信息或大型缓存），每次复制都会消耗大量的 **CPU** 和 **Memory Resource**。
- **The Solution**: 返回 `&`（地址），本质上只是传递了一个 **64-bit Address**（一个极小的数字）。无论原来的结构体有多大，传地址永远是极速的。

#### B. Singleton Pattern（单例模式）与 Shared State（状态共享）

这是最重要的工程原因。`SearchService` 通常是一个 **Service Object（服务对象）**，它内部可能持有数据库连接池或配置。

- **Shared State**: 如果你返回的是指针（地址），那么程序的其他部分拿到的是指向 **同一个内存空间** 的钥匙。这意味着大家都在操作同一个实例。
- **Data Consistency（数据一致性）**: 如果返回的是值（副本），每个模块拿到的都是一份“快照”。你在模块 A 里修改了 Service 的状态，模块 B 根本看不见。

### Q: func (s *SearchService) Search(ctx context.Context, req SearchRequest)(SearchResponse, error) 是什么意思？

---

### 1. 它是如何构成的？ (The Anatomy of a Method)

在 Go 中，这不叫普通函数（Function），这叫 **Method（方法）**。它和普通函数的区别在于最前面那个小括号。

```go
//    (1) Receiver      (2) Name   (3) Input Parameters      (4) Return Values
func (s *SearchService)  Search  (ctx context.Context, ...) (SearchResponse, error)
```

#### (1) Receiver（接收者）：`(s *SearchService)`

- **什么是 Receiver?**：它就像是给这个函数找了个“亲爹”。它告诉编译器：这个 `Search` 方法不是路边随便谁都能用的，它是**专属于** `SearchService` 这个结构体的。

- **指针符号 `*` 的意义**：记得你之前问为什么返回 `&`（地址）吗？这里用 `*` 代表 **Pointer Receiver**。
  - **工程意义**：这意味着当你调用 `Search` 时，你是在操作内存中那个**真实的、唯一的** `SearchService` 对象，而不是它的一个临时副本。这保证了 **Data Consistency**（数据一致性）。

#### (2) Method Name（方法名）：`Search`

- 这是它的身份标识。

#### (3) Input Parameters（输入参数）

- `**ctx context.Context`**：这是 Go 的 **Standard Practice**（标准实践）。
  - **术语：Context Propagation**（上下文传播）。
  - **工程逻辑**：如果这个搜索跑了 10 秒还没出结果，用户关掉网页了，这个 `ctx` 会发出一个 **Cancellation Signal**（取消信号），告诉程序：“别找了，省点 CPU 吧。” 它控制着 **Request Lifecycle**（请求生命周期）。
- `**req SearchRequest`**：这是 **Payload**（有效负载），也就是用户传进来的搜索关键词。

#### (4) Return Values（返回值）：`(SearchResponse, error)`

- Go 支持 **Multiple Return Values**（多返回值）。
- **工程逻辑**：第一个返回搜到的**结果**，第二个返回是否有**报错**。这让你能一眼看出这次调用是 **Success** 还是 **Failure**。

---

### 2. 为什么要这么写？ (The "Why")

如果你觉得这种写法很怪，我们可以对比一下。

**如果没有这种写法（像 C 语言）：**
你可能得写成 `func Search(s *SearchService, ctx ..., req ...)`。

- **问题**：这只是个普通函数。在大型项目中，你很难一眼看出哪些函数是属于“搜索模块”的。

**有了 Method Receiver：**
你现在可以像这样调用代码：

```go
service := NewSearchService() // 拿到地址
service.Search(ctx, req)      // 就像给 service 发送了一条指令
```

这种写法实现了 **Encapsulation（封装）**。它把数据（`SearchService` 里的配置）和行为（`Search` 这个动作）紧紧地捆在了一起。

---

### 3. 形象比喻：自动售货机

我们可以把这行代码想象成一台**自动售货机**：

- `***SearchService` (Receiver)**：这就是那台**具体的售货机**本体（而不是它的照片）。
- `**Search` (Method)**：售货机上的一个按钮——“搜索”。
- `**req SearchRequest` (Input)**：你塞进去的**硬币和选择编号**。
- `**ctx` (Context)**：**投币限时**。如果你 1 分钟没操作，它就会自动退币取消任务。
- `**(SearchResponse, error)` (Return)**：机器吐出来的**可乐（结果）**，或者**红灯报警（错误）**。

### 总结术语表 (Engineering Reference)


| 术语 (Term)             | 含义     | 工程作用                       |
| --------------------- | ------ | -------------------------- |
| **Method Receiver**   | 方法接收者  | 定义函数的“归属权”，实现封装。           |
| **Pointer Receiver**  | 指针接收者  | 避免内存拷贝，确保操作的是同一个实例。        |
| **Request Lifecycle** | 请求生命周期 | 通过 `context` 确保资源不会被无限期占用。 |
| **Error Handling**    | 错误处理   | 通过返回 `error` 强制开发者处理潜在的失败。 |


##### 什么是“归属权”？ (Method Ownership)

想象一个大型超市的后台管理系统：

- 我们有一个**“收银机模块”** (Cashier)
- 我们有一个**“库存模块”** (Warehouse)。

这两个模块可能都有一个动作叫 Open()（打开）。

- 收银机需要 Open()（打开钱箱）。
- 仓库需要 Open()（打开大门）。

如果没有 Receiver（普通函数写法）：

- 如果大家都写成普通的 func Open()，代码就会打架。
- 你调用 Open() 的时候，系统不知道你想开钱箱还是开大门。
- 这在工程上叫 Naming Collision（命名冲突）。

#### 有了 Receiver (`SearchService`)：

```
func (s *SearchService) Search(...) // 搜索服务的 Search
func (a *AdminService) Search(...)  // 管理员服务的 Search
```

这里的 `(s *SearchService)` 就是在申明：**“这个 Search 方法是长在 SearchService 身上的零件。”** **这就叫“专属”：** 你不能拿着 `AdminService` 去调用这个搜索方法，编译器会报错。它确保了代码的 **Type Safety（类型安全）**。