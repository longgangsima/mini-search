# Day 2（旧 20 天 / 3 周 docx 线）· Struct + JSON + Query Params

> **已弃用为主日程**：主跟 **v23** 请以 [day2.md](day2.md)（Error + Validation）为准。  
> 本文件仅保留旧课表文字，便于对照；**Struct / JSON** 专题仍可看 [struct.md](struct.md)。

---

## Day 2 · Struct + JSON + Query Params（2–3h）

### 学什么

- Go by Example：Structs, Methods,JSON  
- 重点：`json:"name"` struct tag  
- `encoding/json`：Marshal / Unmarshal  
- `net/url` 解析 query params

### 项目任务

- 定义 `SearchRequest`：`Q`（query）, `Store`, `Page`  
- 定义 `SearchResponse`：`Results []Product`, `Total int`  
- 添加 `GET /search`，返回 mock 数据  
- 从 `r.URL.Query()` 读取 `q` / `store` / `page` 并打印

### 记忆激活

- Walmart.ca 的 search endpoint 大致也是：handler 接请求、解析 params、返回 JSON
