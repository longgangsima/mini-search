# Go 基础预习（Week 1）

「当天写项目前先过一遍」的**索引入口**；正文仍在各专题文件里，避免一份里双份维护。


| 主题                        | 文件                                                                     |
| ------------------------- | ---------------------------------------------------------------------- |
| Struct / 指针示例 / 匿名 struct | [struct.md](struct.md)                                                 |
| Slice                     | [slice.md](slice.md)                                                   |
| Map                       | [maps.md](maps.md)                                                     |
| 指针、逃逸、comma ok 总览         | [memory-pointers-escape-and-map.md](memory-pointers-escape-and-map.md) |
| Day 1 课表「学什么」+ HTTP 骨架    | [day1.md](day1.md)                                                     |


若你打算把 `struct.md` 全文迁到本文件：请先把 `struct.md` 全文复制进本文件，再批量改 [README.md](README.md) 与各处的 `struct.md` 链接，最后删除旧文件。**不要**只删 `struct.md` 而不粘贴——工作区会丢稿；已提交版本可用 `git restore docs/week-1/struct.md` 救回。