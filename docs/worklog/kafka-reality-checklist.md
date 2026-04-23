# Kafka Reality Checklist (Pre-Interview)

Fill this file from real code and actual run logs.

## Topology Facts

- Producer service name:
- Consumer service name:
- Topic name(s):
- Consumer group id:
- Partition count:
- Partition key:
- Why this partition key:

## Delivery Semantics

- Commit strategy: auto / manual
- Commit timing: before processing / after processing
- Chosen semantic: at-most-once / at-least-once
- Why this is acceptable for your business case:

## Failure Handling

- Retry count:
- Retry backoff:
- DLQ topic:
- DLQ trigger condition:
- DLQ reprocessing owner/process:

## Rebalance Behavior

- What events trigger rebalance in your setup:
- Observed effect during rebalance:
- Session timeout / heartbeat settings:

## Idempotency

- Idempotency key:
- Duplicate handling logic:
- Side effects guarded against duplicates:

## Evidence Links (fill with real references)

- code reference 1:
- code reference 2:
- test case:
- run log/screenshot:

## Final Honesty Check

- Which parts did I implement directly?
- Which parts were designed by others but I understand?
- Which parts are still unknown and need clarification?
