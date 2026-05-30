# Go Concurrency Practice

Nine self-contained problems for practicing Go concurrency patterns: goroutines,
channels, `sync` primitives, `context`, and cancellation. Each folder is an
independent Go module — you can `cd` into one and work without touching the
others.

## Layout

```
easy/
  01-parallel-sum/        WaitGroup + fan-in
  02-ping-pong/           Unbuffered channels + synchronization
  03-worker-pool/         Fixed worker pool consuming a job channel
medium/
  01-rate-limiter/        Token bucket using time.Ticker + channels
  02-pipeline/            Multi-stage pipeline with proper channel closing
  03-concurrent-fetcher/  Bounded concurrency, ordered results, context cancel
hard/
  01-lru-cache-ttl/       Thread-safe LRU with TTL and background eviction
  02-web-crawler/         BFS crawler: depth limit, dedup, max workers, ctx
  03-job-scheduler/       Priority queue + workers + retries + cancellation
```

## How to use

Each problem folder contains:

- `README.md` — the problem statement, requirements, hints, and a stretch goal.
- `main.go` — function/type stubs with `// TODO` markers. Implement these.
- `main_test.go` — a few tests to verify your implementation. Run with
  `go test ./...` inside the folder.

Recommended workflow per problem:

```bash
cd easy/01-parallel-sum
go test ./...        # see failing tests
# implement in main.go
go test ./...        # iterate until green
go test -race ./...  # check for data races
```

## Suggested order

Work top-down: `easy/01 → easy/02 → easy/03 → medium/01 → ...`. Later
problems reuse patterns from earlier ones (fan-out/fan-in, `context`
cancellation, `sync.Mutex` vs channels for state).

## Tips

- Run with `-race` often. Most concurrency bugs surface only under the
  race detector.
- Prefer `context.Context` for cancellation over ad-hoc `done` channels
  once you're past the easy tier.
- "Don't communicate by sharing memory; share memory by communicating" —
  but `sync.Mutex` is the right tool when you genuinely have shared state
  (e.g. the LRU cache).
- Always close channels from the sender side, never the receiver.
