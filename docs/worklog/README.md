# Worklog

**日粒度叙事**：进行中的「当天」**只写在 `CURRENT.local.md`**；`HISTORY.local.md` **当天不收工就不出现该日**（无占位）。收工后再把 `CURRENT` 正文 **复制或摘要 append** 进 `HISTORY`。课表与讲义仍在 `docs/week-1/`、`docs/handbook-v2/`。

This folder is a rolling handoff journal for active work.

It is meant to answer two questions quickly:

1. what is the current state right now
2. what happened in previous work sessions

## Files

### `CURRENT.local.md`

The live handoff file — **today / this session’s scratch pad**（计划、进行中、课表级打卡全文都可写在此）。收工、归档到 `HISTORY` 后，再重置为下一天的短占位。

Use it for:

- what was done most recently
- what is currently in progress
- the exact next step
- blockers, assumptions, and reminders

This file is expected to be **rolled forward** after each session (archive → `HISTORY`, then reset).

### `HISTORY.local.md`

The archive file.

Use it for short session summaries when `CURRENT.local.md` is rolled forward.

Each time a new work session starts:

1. read `CURRENT.local.md`
2. continue the work
3. when you stop, summarize the completed session into `HISTORY.local.md`
4. replace `CURRENT.local.md` with the new handoff state for the next return

### `CONTEXT.local.md`

The stable background file.

Use it for:

- current project overview
- important links and file paths
- longer-lived decisions
- environment or setup reminders

This file should change much less often than `CURRENT.local.md`.

### Optional templates (same directory)

- `daily-copy-template.md` — 「Recommended / Minimal」打卡结构，可拷贝进 `HISTORY` 或当 checklist。
- `weekly-review-template.md` — 周复盘。
- `kafka-reality-checklist.md` — Day 18 前后 Kafka 与面试前事实核对。

### `templates/daily-template.md`

Skeleton for rolling `CURRENT.local.md` forward (the live file stays `*.local.md` and remains ignored).

## Suggested Rule

**未收工的当天**：正文**只在 `CURRENT`**，`HISTORY` **不要**提前写该日标题或指针。**收工后**：在 `HISTORY` 末尾新开该日小节并粘贴正文（或摘要），然后清空 `CURRENT`、改日期为次日。

Keep `CURRENT.local.md` detailed enough to resume immediately.

Keep `HISTORY.local.md` compressed enough to scan over time.

Keep `CONTEXT.local.md` limited to stable background that should survive across many sessions.

## Daily Update Checklist

When closing or handing off a session:

1. copy or summarize `CURRENT.local.md` into `HISTORY.local.md` for that date
2. replace `CURRENT.local.md` with the new live handoff state
3. update `CONTEXT.local.md` only if long-lived project background has changed
4. keep repo-facing curriculum in `docs/week-1/` and `docs/handbook-v2/`

## Naming

This pattern is best described as:

- rolling handoff log
- working journal
- current state + history archive

中文更自然可以叫：

- 滚动工作日志
- 交接日志
- 当前状态与历史归档
