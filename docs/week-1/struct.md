# Struct

> 与 [memory-pointers-escape-and-map.md](memory-pointers-escape-and-map.md) 有重叠：本文件保留 **例题与 struct 细节**；总览与术语表以 memory 文件为辅。

## Go 语言中 **Struct**（结构体）的定义、初始化以及如何通过 **Pointer**（指针）来操作它们。

---

### 1. `type person struct`：定义结构
```go
type person struct {
    name string
    age  int
}
```
这里定义了一个名为 `person` 的 **Named Struct**。
* **Fields**: 它包含两个 **Fields**（字段）：`name` 和 `age`。这就像是自定义了一个 **Type Identity**。

---

### 2. 多种 Initialization（初始化）方式
代码展示了在 **Main Function** 中创建 **Struct Instance** 的不同方法：

* **Positional Syntax**: `person{"Bob", 20}` 必须按照定义的顺序传入 **Values**。
* **Named Fields**: `person{name: "Alice", age: 30}` 推荐做法，更清晰且顺序可变。
* **Omitted Fields**: `person{name: "Fred"}` 没提到的字段（如 `age`）会自动赋为该类型的 **Zero Value**（这里是 `0`）。
* **Pointer Literal**: `&person{name: "Ann", age: 40}` 直接创建一个指向该实例的 **Pointer**。



---

### 3. `newPerson`：构造函数模式
```go
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p // 返回 Pointer
}
```
* **Encapsulation**: 这是一个典型的 **Constructor Pattern**。
* **Escape Analysis**: 这里的逻辑非常有趣——在 C++ 中，返回局部变量的地址是危险的；但在 Go 中，**Compiler** 会执行 **Escape Analysis**（逃逸分析），自动把 `p` 分配到 **Heap**（堆）上，确保函数结束后它依然有效。

---

### 4. Accessing and Modifying Fields（访问与修改）
```go
sp := &s
fmt.Println(sp.age)
sp.age = 51
```
* **Dot Notation**: 无论 `s` 是 **Struct Value** 还是 **Struct Pointer**，你都直接用 `.` 来访问字段。
* **Automatic Dereferencing**: Go 非常聪明，当你用 `sp.age` 时，它会自动执行 **Dereferencing**（解引用），你不需要像 C 语言那样写 `(*sp).age` 或者 `sp->age`。

---

### 5. `dog := struct { ... }`：匿名结构体
```go
dog := struct {
    name   string
    isGood bool
}{ "Rex", true }
```
* **Anonymous Struct**: 这定义了一个没有名字的 **Struct**。
* **Use Case**: 当你只需要临时用一次（比如解析特定的 JSON 或者写个小测试）而不想正式定义一个 **Type** 时，这种写法非常方便。

---

### 总结 (Terminology Summary)

| 术语 (Terminology) | 作用 (Role) |
| :--- | :--- |
| **Struct** | 将不同 **Types** 的数据组合在一起的集合。 |
| **Field** | **Struct** 内部的一个命名数据项。 |
| **Pointer (`*`)** | 存储 **Struct Instance** 在内存中的地址。 |
| **Anonymous Struct** | 没有正式 **Type Name** 的临时结构。 |
| **Zero Value** | 字段未赋值时的默认值（如 `int` 为 `0`, `string` 为 `""`）。 |

**一句话总结：**
这段代码在演示如何像“搭积木”一样把数据封装进 **Struct**，以及如何利用 **Pointer** 高效地引用和修改这些积木块，而不需要复制整个数据块。

你觉得这种直接用 `.` 就能同时操作 **Value** 和 **Pointer** 的设计，是不是比 C 语言那种区分 `.` 和 `->` 的做法省心多了？

##Questions:

###An & prefix yields a pointer to the struct. 这样写的作用是什么？跟别的区别是什么？

###Escape Analysis: 这里的逻辑非常有趣——在 C++ 中，返回局部变量的地址是危险的；但在 Go 中，Compiler 会执行 Escape Analysis（逃逸分析），自动把 p 分配到 Heap（堆）上，确保函数结束后它依然有效。没太懂你的意思。

###sp := &s 这是什么？

###Struct Value 还是 Struct Pointer。这2者是什么？

###Pointer (*)存储 Struct Instance 在内存中的地址。 这是啥意思啊？