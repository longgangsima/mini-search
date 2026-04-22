# Interview Alignment (Kafka + Go + Django/PostgreSQL)

This file maps your project work to high-frequency interview questions.

## Rule Zero (Do Not Skip)

Before interviews, verify implementation facts in code. Do not rely on memory.

Use this phrasing when needed:

> "I mainly implemented X. Y was designed by another engineer, and my understanding of the design is ..."

This is safer and stronger than over-claiming.

## Kafka Mapping

Expected interview questions:

- validator role in pipeline
- failure handling and offset commit timing
- partition and consumer group parallelism
- rebalance behavior
- at-least-once vs at-most-once tradeoff

Project actions that support these answers:

- capture topic and consumer group names in notes
- document partition key and why it was chosen
- test manual vs auto commit and record observed behavior
- define retry and DLQ flow clearly

Minimum proof artifacts:

- one architecture diagram
- one produce/consume demo run log
- one short note for rebalance experiment outcome

## Golang Mapping

Expected interview questions:

- goroutine vs thread
- channel coordination and timeout patterns
- `context.Context` purpose and propagation
- error handling with `%w`, `errors.Is`, `errors.As`

Project actions that support these answers:

- build fan-out/fan-in service flow with 3 downstream clients
- enforce timeout and cancellation through handler -> service -> client
- implement typed errors and HTTP status mapping
- add tests for timeout and error-path behavior

Minimum proof artifacts:

- one fan-out/fan-in code path
- one timeout cancellation test
- one error-wrapping and `errors.Is` test

## Django + PostgreSQL Mapping (Story Prep)

This project is Go-first, but interview prep should include your prior backend stack.

Prepare concise, factual stories for:

- optimistic concurrency control with version field
- avoiding lost updates with conditional update
- ORM performance tuning (`select_related` / `prefetch_related`)
- transaction isolation tradeoffs

Minimum proof artifacts:

- one concise pseudo-query example of optimistic locking
- one N+1 example and optimized form
- one note on when you used `READ COMMITTED` vs stricter levels

## 30-second Positioning Draft

Use your own words, but keep this structure:

1. Current focus: full-stack delivery with backend ownership
2. Recent backend depth: Go service concurrency + reliability patterns
3. Prior system depth: Django/PostgreSQL data model and concurrency controls
4. Impact: latency, correctness, and operability improvements
