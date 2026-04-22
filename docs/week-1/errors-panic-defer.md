# `errors.Is` / `errors.As`、`panic`、`defer` — 场景备忘

与 [day2.md](day2.md) 的 Error + Validation、`defer` 练习配套；代码以仓库 `internal/` 为准。

---

## 一、什么时候用 `errors.Is`？

**用途**：在一条 **error chain（错误链）**（含 `Unwrap`）里判断：是不是某个固定的 **sentinel error（哨兵错误）**。

`errors.Is` 就像一台 **X 光机**。它不仅仅看最外层，它会一层一层往里剥，看看这个错误的“祖宗十八代”里，有没有你指定的那个错误。

**典型场景**：

- **standard library**：`io.EOF`、`context.Canceled`、`fs.ErrNotExist`
- 自己包：`var ErrNotFound = errors.New("not found")`，各处 `return fmt.Errorf("load: %w", ErrNotFound)`，外层 `errors.Is(err, ErrNotFound)`
- 当你只想知道这个错误**“本质上”**是不是某个特定的错（比如是不是“超时”了、是不是“没找到文件”），而不管它外面套了多少层说明文字时，就用 `errors.Is`。
- **什么是“哨兵”（Sentinel）？** 就是你在包的最开头，用 `var` 定义的那些全局错误变量。它们就像哨兵一样站在那里，等着你去认领。
- **为什么要避开字符串比较？** 如果你用 `err.Error() == "包裹丢了"`，万一别人把文字改成 `"包裹丢失"`，你的代码就崩了。用 `errors.Is` 比的是**“身份”**，而不是文字。

**要点**：

- `errors.New("x")` **每调一次都是新实例**；两个 `New("x")` **不能用 `Is` 互认**，要和 **package-level variable（包级变量）** 比：`errors.Is(err, ErrXxx)`。
- `Is` 比的是 **identity（身份）** / 链上是否出现该哨兵，**不是**比 `Error()` 字符串是否相同。

```go
var ErrEmpty = errors.New("empty")

func f() error {
	return fmt.Errorf("parse: %w", ErrEmpty)
}

if errors.Is(f(), ErrEmpty) { /* true：链里能 unwrap 到 ErrEmpty */ }

var ErrLost = errors.New("包裹丢了") // 这是我们的“哨兵”
func checkStatus() error {
    // 我们把原始错误包装了一层，加了点上下文信息
    return fmt.Errorf("物流部报告: %w", ErrLost) 
}
func main() {
    err := checkStatus()
    // 1. 用 == 比 (会失败)
    fmt.Println(err == ErrLost) // 输出: false
    // 2. 用 errors.Is 比 (会成功)
    // 它会问：这个 err 里面套着 ErrLost 吗？
    fmt.Println(errors.Is(err, ErrLost)) // 输出: true
}
```

---

## 二、什么时候用 `errors.As`？

**用途**：在 **error chain** 里 **取出 concrete type（具体类型）**（通常带 **fields**），做分支或 **map 到 HTTP status code**。

**典型场景**：

- 自定义类型：`*`LoginError（有 Attempts、Timeout）
- 把 `error` 转成具体 **struct** / **interface** 实现，再读 **fields**

**要点**：

- 第二个参数必须是 **pointer to pointer（指向指针变量的指针）**：var le *LoginError errors.As(err, &le) 。
- 适合 **「一种错误一类数据」**；哨兵 `errors.New` 没有 **fields**，一般用 `Is` 而不是 `As`。

```go
type LoginError struct {
    Attempts int
    Timeout  int
}

func (e *LoginError) Error() string {
    return "登录失败过多"
}

func main() {
    err := login() // 假设返回了被包装过的 LoginError

    // ❌ 错误做法：直接拿
    // fmt.Println(err.Attempts) // 报错！error 接口没这个字段

    // ✅ 正确做法：用 errors.As
    var le *LoginError // 1. 先声明一个目标类型的指针变量（空口袋）
    
    if errors.As(err, &le) { // 2. 问：err 里面有没有 LoginError？有的话请装进 &le 这个口袋
        // 3. 成功拿到具体数据！
        fmt.Printf("你已经试了 %d 次了，请等 %d 分钟再试。", le.Attempts, le.Timeout)
    }
}
```

---

## 三、`Is` vs `As` 对照


| 维度           | `errors.Is(err, target)`               | `errors.As(err, &ptr)`                      |
| ------------ | -------------------------------------- | ------------------------------------------- |
| 第二个参数        | `error`（常是 **package-level sentinel**） | `&ptr`，`ptr` 为 **pointer to concrete type** |
| 适合           | `EOF`、`Canceled`、`ErrXxx`              | `*ValidationError` 等                        |
| 读 **fields** | 一般不读（只有字符串）                            | 转成 `ptr` 后用 **fields**                      |
| 与 `%w`       | 可 **unwrap** 到哨兵                       | 可 **unwrap** 到自定义类型                         |


**口诀**：**sentinel → `Is`，带 struct 数据的错误 → `As`。**

---

## 四、`panic` — 何时用、何时别用

**是什么**：立刻中止当前 **goroutine** 的正常控制流，开始 **stack unwinding（栈展开）**；若未被 `**recover`** 接住，**process** 退出。

**适合（相对少见）**：

- **programming error / broken invariant**：`must never happen`，例如内部状态机错乱；很多库用 `panic("index out of range")` 类断言失败。
- `init` 里致命配置缺失（仍有不少人改为 `log.Fatal` 后退出）。

**不适合（业务日常）**：

- 用户输入非法、HTTP **4xx**、查库为空 → 用 `**return ..., err`**，不要 `**panic**`。
- 可预期的「失败」一律走 `**error` return value**，便于调用方处理和测试。

`**recover`**：几乎只写在 `**defer**` 里，且通常限定在顶层（例如 HTTP **middleware** 防单个 **handler panic** 拖死 **process**）；不要在业务函数里到处 `**recover`** 当 `try/catch`。

---

## 五、`defer` — 典型用法

**语义**：函数 **return** 前（含 `**panic` 正在展开**时）按 **LIFO** 执行注册的调用。

**常见场景**：

1. **配对释放**：`mu.Lock()` / `defer mu.Unlock()`；`f, _ := os.Open(...)` / `defer f.Close()`。
2. **出口统一逻辑**：`defer func() { log.Println("done") }()`（若闭包要读返回值，可用 **named return values** + `defer` 里赋值，进阶再练）。
3. **与 `panic`/`recover` 配合**：`defer func() { if r := recover(); r != nil { ... } }()`。

**注意**：

- `**defer` 的参数在 `defer` 语句执行时求值**，不是函数返回那一刻再算外层变量（闭包除外）。
- 热路径里滥用 `defer` 曾有性能讨论；对 HTTP **handler** 里一两处 `defer` 日志**完全不用担心**。

---

## 六、和 mini-search 的对应关系

- **校验失败**：`*ValidationError` → **handler** 里 `**errors.As`** → **400** + JSON。
- **其它内部错**：**500**，不把 **stack trace** 暴露给客户端；可 `**log` + `defer`** 打「请求结束」。
- **业务路径**：不要 `**panic`** 表达「参数错了」；用 `**error**`。

---

## 延伸阅读

- Go by Example：[Errors](https://gobyexample.com/errors)、[Panic](https://gobyexample.com/panic)、[Defer](https://gobyexample.com/defer)
- `pkg.go.dev/errors`：`Is`、`As`、`Unwrap`

