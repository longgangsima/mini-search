# Memory, pointers, escape analysis, and map lookup (study notes)

> 学习过程中「不懂 → 弄懂」的整理；术语用英文，解释中英混排。  
> 本目录索引与去重规则见 [README.md](README.md)。map / struct **详细版**分别以 [maps.md](maps.md)、[struct.md](struct.md) 为准。

---

## 1. `&` prefix and struct pointer (`*struct`)

**`&`** in Go is the **address-of operator**（取地址符）。

- **作用**：不新建一份「物体」，而是拿到已有值在内存里的地址。
- **Struct value** (`p := person{...}`)：你手里是**实物**；把 `p` 传给函数时，Go 会按值**复制**一整份 struct（大 struct 时成本高）。
- **Struct pointer** (`pp := &person{...}`)：你手里是**纸条**（地址），传参只传指针大小（通常 8 字节），快。

**为什么要用 `&`？**  
大 struct 每次传值都复制会伤 **performance**；传 **pointer** 让多处共享同一份底层数据，并能在 callee 里改到「原件」（通过指针间接写回）。

---

## 2. Struct value vs struct pointer（和你代码最相关）

- **Struct value**  
  `s := person{name: "Sean"}`：`s` 就是那块内存里的数据本身。`s2 := s` 会得到第二份拷贝。

- **Struct pointer**  
  `sp := &s`：`sp` 的类型是 `*person`，存的是 `s` 的地址。通过 `sp` 改字段，改的是同一份 `s`（Go 对字段访问会做合理的解引用，写法上常直接 `sp.Field`）。

---

## 3. Pointer (`*`) and “存的是地址”是什么意思？

把内存想成酒店房间：

- **Struct instance**：房间里的客人（数据）。
- **Pointer**：房卡 / 房间号（地址）。
- **Type `*person`**：一张「只能指向 `person` 这种房间」的房卡类型。

拿到地址后，可以通过它去读写那个房间里的字段。

---

## 4. Escape analysis（逃逸分析）

Go **compiler** 会做 **escape analysis**：

- 很多局部变量在 **stack** 上分配很快，函数结束栈帧回收。
- 若你把某个局部变量的 **address** 泄露到函数外（例如 `return &p`），compiler 发现「这个地址在函数返回后仍可能被用」，就会把 `p` **搬到 heap**，由 GC 管理生命周期，避免悬垂指针。

对比 **C++**：返回局部变量地址是经典 UB；Go 用逃逸分析 + heap 分配来避免这类崩溃（代价是可能多一点堆分配）。

---

## 5. `sp := &s` 一行里发生了什么？

1. **Address-of**：取 `s` 的地址。  
2. **Short variable declaration**：声明 `sp`。  
3. **Type inference**：compiler 推断 `sp` 的类型为 `*person`。

---

## 6. Map 与 **comma ok** idiom（摘要）

- **问题**：`m["k"]` 在 key 不存在时，会返回该 value 类型的 **zero value**（`0`、`""` 等），你无法区分「真的存了 0」还是「没有这个 key」。
- **解决**：`v, ok := m["k"]` 第二个返回值 **bool** 表示是否存在。  
- **`_` blank identifier**：`_, ok := m["k"]` 只关心存在性时，丢弃 value。

这叫 **comma ok idiom**（两返回值习惯写法）。**例题级展开、术语表与代码片段以 [maps.md](maps.md) 为准**，本节不重复粘贴。

---

## Terminology summary（Global setting）

| Terminology | 形象理解 | 核心价值 |
| :--- | :--- | :--- |
| **Address (`&`)** | 查房号 | 找到数据在内存的源头 |
| **Pointer (`*`)** | 房卡 / 地址纸条 | 少复制、共享、间接修改 |
| **Stack** | 临时小仓库 | 快，函数结束随栈帧回收 |
| **Heap** | 长期存放区 | 可跨函数生命周期，由 GC 回收 |
| **Escape analysis** | 智能搬家 | Compiler 决定值放栈还是堆，避免悬垂引用 |
| **Composite literal** | `T{...}` 一眼初始化 | 可读、可维护 |
| **Anonymous struct** | 无名 struct 类型 | 临时结构，少起类型名 |

---

## 今日自检清单（Checklist）

能立刻在脑子里「对应到内存样子」就算过关：

1. **Type identity**：为何 `[3]int` 与 `[4]int` 不同类型、不能乱赋值。  
2. **Bound check elimination (BCE)**（概念级）：编译器在可证明安全时减少边界检查，数组/切片热点路径更快（细节以后写性能时再深挖）。  
3. **Pointer & address (`*` & `&`)**：值 vs 指针，何时复制、何时共享。  
4. **Escape analysis**：`return &local` 为什么通常仍然安全。  
5. **Composite literal**：`person{name: "Alice"}` 这类写法。  
6. **Anonymous struct**：临时凑合用的无名 struct。  
7. **Map comma ok**：`v, ok := m[k]` 区分「零值」与「不存在」。

---

## 和 mini-search 项目的关系（落地一句）

- `NewSearchService() *SearchService`：返回指针，让 handler 持有**同一份** service 实例，避免无意义复制，也方便以后给 struct 加字段（clients、repo）。  
- `r.Context()`：请求级上下文，后面接超时/取消；和指针/逃逸是不同维度，但都属于「写生产级 Go」必备拼图。

---

## 备注（诚实边界）

- **Escape analysis** 的具体规则以官方文档与 `go build -gcflags=-m=2` 输出为准；上面是面试友好的直觉版。  
- **BCE** 是否发生取决于编译器版本与代码形状，不必 Day 1 背细节。
