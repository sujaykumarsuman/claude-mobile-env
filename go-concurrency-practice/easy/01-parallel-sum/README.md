# Easy 01 — Parallel Sum

Compute the sum of a slice of `int64` by splitting the work across `N`
goroutines and combining the partial sums.

## Function to implement

```go
func ParallelSum(nums []int64, workers int) int64
```

## Requirements

- Split `nums` into roughly equal chunks, one per worker.
- Each worker computes the sum of its chunk concurrently.
- Combine partial results into a single total and return it.
- If `workers <= 0`, treat it as `1`. If `len(nums) == 0`, return `0`.
- No data races — verify with `go test -race ./...`.

## Hints

- `sync.WaitGroup` to wait for all workers to finish.
- Two reasonable shapes:
  1. Each worker sends its partial sum on a buffered channel; the caller
     drains the channel after `Wait`.
  2. Each worker writes into a pre-sized `partials[i]` slot indexed by
     worker id (no shared mutable state — distinct indices).
- Avoid `sum += ...` on a shared variable without a mutex; that's a race.

## Stretch goals

- Benchmark `ParallelSum` against a serial sum for slices of 1M, 10M
  elements. At what size does parallelism win on your machine?
- Make the worker count auto-tune via `runtime.NumCPU()`.
