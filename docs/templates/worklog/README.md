# Worklog（约定说明）

**个人滚动日志不进仓库**：在本地创建目录 **`docs/worklog/`**（已在 `.gitignore`），在其中维护 **`CURRENT.local.md`**、**`HISTORY.local.md`**、**`CONTEXT.local.md`** 等；这些文件不会被 `git push`。

本目录 **`docs/templates/worklog/`** 提供可随仓库分发的**用法说明**与**空白模板**；需要时把其中的 `*.md` **复制**到你的 **`docs/worklog/`** 即可。

---

**日粒度叙事**：进行中的「当天」**只写在 `CURRENT.local.md`**；`HISTORY.local.md` **当天不收工就不出现该日**（无占位）。收工后再把 `CURRENT` 正文 **复制或摘要 append** 进 `HISTORY`。课表与讲义仍在 `docs/week-1/`、`docs/handbook-v2/`。

This folder pattern answers:

1. what is the current state right now
2. what happened in previous work sessions

## Files（在本地 `docs/worklog/`）

### `CURRENT.local.md`

The live handoff file — **today / this session’s scratch pad**（计划、进行中、课表级打卡全文都可写在此）。收工、归档到 `HISTORY` 后，再重置为下一天的短占位。

### `HISTORY.local.md`

The archive file. Append session summaries when you roll `CURRENT` forward.

### `CONTEXT.local.md`

Stable background that should survive across many sessions.

## Optional templates（本目录中的文件）

- `daily-copy-template.md` — 「Recommended / Minimal」打卡结构，可拷贝进 `HISTORY` 或当 checklist。
- `weekly-review-template.md` — 周复盘。
- `kafka-reality-checklist.md` — Day 18 前后 Kafka 与面试前事实核对。

### `daily-template.md`（也可放在本地 `docs/worklog/templates/`）

Skeleton for rolling `CURRENT.local.md` forward（线上交接文件仍为 `*.local.md`，由 gitignore 排除）。

## Suggested Rule

**未收工的当天**：正文**只在 `CURRENT`**，`HISTORY` **不要**提前写该日标题或指针。**收工后**：在 `HISTORY` 末尾新开该日小节并粘贴正文（或摘要），然后清空 `CURRENT`、改日期为次日。

## Daily Update Checklist

When closing or handing off a session:

1. copy or summarize `CURRENT.local.md` into `HISTORY.local.md` for that date
2. replace `CURRENT.local.md` with the new live handoff state
3. update `CONTEXT.local.md` only if long-lived project background has changed
4. keep repo-facing curriculum in `docs/week-1/` and `docs/handbook-v2/`

## Naming

- rolling handoff log
- working journal
- current state + history archive

中文：滚动工作日志、交接日志、当前状态与历史归档。