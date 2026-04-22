# Maps

> 与 [memory-pointers-escape-and-map.md](memory-pointers-escape-and-map.md) §6 同一主题：本文件为 **详细版**；总览只看 memory 文件即可。

## `value, ok := m["key"]`

这段话描述的是 Go 语言中处理 **Map** 时一个非常经典且重要的模式，通常被称为 **"Comma ok" idiom**。

在 Go 中，当你从 **Map** 中取值时，它其实可以返回两个 **Values**。我们按照 **Global Setting** 来拆解这个逻辑：

### 1. 为什么需要两个返回值？（The Ambiguity Problem）

在 JS 或 Python 中，如果 **Key** 不存在，可能会返回 `undefined` 或报错。但在 Go 中，如果 **Key** 不存在，它会返回该 **Type** 的 **Zero Value**（零值）。

- **问题来了**：如果你从 `map[string]int` 里取出的值是 `0`，你无法确定是因为 **Key** 真的存在且值就是 `0`，还是因为 **Key** 根本不存在。

---

### 2. "Comma ok" 语法拆解

为了解决这种模糊性（**Disambiguate**），Go 提供了这种写法：

`value, ok := m["key"]`

- `value`: 实际存储的 **Value**。如果 **Key** 不存在，它就是 **Zero Value**。
- `ok` (在你的例子里叫 `prs`，即 present): 一个 **bool** 类型。
  - `true`: 表示 **Key** 确实在 **Map** 里。
  - `false`: 表示 **Key** 不存在。

---

### 3. 你的代码示例分析

Go

```
_, prs := m["k2"]

```

- `_` **(Blank Identifier)**:
由于你现在只关心“这个 **Key** 在不在”，而不关心它具体存的是什么，所以你使用 **Blank Identifier** 丢弃了第一个返回值。在 Go 中，声明了变量但不使用会触发 **Compile Error**，所以 `_` 是非常有用的“垃圾桶”。
- `prs` **(Presence)**:
这是一个 **bool** 变量。
- `fmt.Println("prs:", prs)`:
如果 `k2` 存在，打印 `prs: true`；否则打印 `prs: false`。

---

### 4. 总结 (Terminology Summary)


|                                |                                             |
| ------------------------------ | ------------------------------------------- |
| **术语 (Terminology)**           | **作用 (Role)**                               |
| **Map**                        | 一种存储 **Key-Value pairs** 的数据结构。             |
| **Zero Value**                 | 类型的默认初始值（如 `int` 是 `0`，`string` 是 `""`）。    |
| **Blank Identifier (**`_`**)** | 用来忽略不需要的 **Return Value**。                  |
| **Comma ok idiom**             | 这种 `value, ok :=` 的模式，用于安全地检查 **Key** 是否存在。 |


**一句话总结：**

这行代码是在**“侦测”**：它不想要具体的数据，它只想确认 `"k2"` 这个名字到底有没有在 **Map** 的名单（索引）里。

这种写法在 Go 的逻辑判断中非常常见，比如：

Go

```
if _, ok := m["k2"]; ok {
    // 只有当 k2 存在时，才执行这里的逻辑
}

```

你觉得这种通过返回两个 **Values** 来确认状态的设计，比起其它语言返回 `null` 或抛出异常，哪种更符合你的直觉？

## n := map[string]int{"foo": 1, "bar": 2} 这是什么？

- `map[string]int`: 这是 **Type Identity**。
  - `map`: 关键字，代表这是一个 **Hash Table** 结构。
  - `string`: **Key Type**，表示键必须是字符串。
  - `int`: **Value Type**，表示值必须是整数。
- `{"foo": 1, "bar": 2}`: 这叫 **Composite Literal**（复合字面量）。
  - 它在内存中直接创建了包含两个 **Key-Value Pairs**（键值对）的 **Map**。

