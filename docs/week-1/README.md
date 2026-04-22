# Week 1 学习资料索引

本目录是 **笔记与课表**，不是可执行代码；代码以仓库根目录 `cmd/`、`internal/` 为准。

## 文件一览（按用途）

| 文件 | 用途 | 备注 |
|------|------|------|
| [day1.md](day1.md) | **Day 1 执行手册**（v23 长日）+ 当天延伸问答 | 原文末长 Summary 已并入 `architecture-web-handler-ioc.md` §三～§五；`day1` 文末仅保留跳转，避免双份维护。 |
| [day2.md](day2.md) | **Day 2（v23）**：Error + Validation、`errors.As`、JSON 错误体、`defer` 日志 | 与 `handbook-v2` Phase 1 Day 2 对齐；代码以仓库 `internal/` 为准。 |
| [errors-panic-defer.md](errors-panic-defer.md) | **`errors.Is` / `errors.As`、`panic`、`defer` 场景备忘** | 中文叙述 + 英文术语；与 `day2` 互补。 |
| [go-basic.md](go-basic.md) | **Go 基础预习索引** | 链到 `struct` / `slice` / `maps` / `memory` / `day1`。 |
| [day02-legacy-20day-struct-json-query.md](day02-legacy-20day-struct-json-query.md) | 旧 20 天线的「Struct + JSON + Query」Day 2 文字备份 | 非主日程；struct 深挖仍看 `struct.md`。 |
| [architecture-web-handler-ioc.md](architecture-web-handler-ioc.md) | **Day 1 补充**：Handler vs Service、谁调 `ServeHTTP`、DI、**composition root**、`NewSearchHandler` 同包、长 import 路径 | **架构叙事真源**（含原 `day1` Summary 合并内容）；不重复贴业务代码。 |
| [struct.md](struct.md) | Struct / pointer / escape 的 **专题深挖**（Go by Example 路线） | 与 `memory-pointers-escape-and-map.md` 有重叠：**struct 细节以本文件为准**，总览表看 memory 文件。 |
| [slice.md](slice.md) | Slice 专题 | 独立。 |
| [maps.md](maps.md) | Map 与 **comma ok** 专题 | 与 `memory-pointers-escape-and-map.md` §6 同主题：**详细版以本文件为准**。 |
| [memory-pointers-escape-and-map.md](memory-pointers-escape-and-map.md) | 指针、逃逸、map comma ok **一页总览** + 和项目挂钩 | **不写重复长例子**；map 与 struct 分别指向 `maps.md` / `struct.md`。 |
| [环境变量.md](环境变量.md) | `PORT`、`os.Getenv`、12-factor | 与 `day1.md` 里 config 说明互补，**不重复删**，config 以代码为准。 |
| [questions.md](questions.md) | **问答索引** + 长文（gRPC、interface、receiver 等） | 顶部索引链到各专题；避免再写第三个「地图 comma ok」长文。 |

## 去重原则（以后照做）

1. **同一概念只保留一份「详细长文」**：另一处用 **3 行摘要 + 链接**。  
2. **课表（dayN）** 里优先放「当天步骤与验收」；超长原理挪到 `topic-*.md` 或 `questions.md`。  
3. **模块路径** 以仓库根目录 `go.mod` 的 `module ...` 为准；文档里出现旧路径时以 `go.mod` 为真源。

## Records 打卡

每日一行/一篇见 `docs/records/`（例如 `day-01-log.md`），链回本目录对应长文即可。
