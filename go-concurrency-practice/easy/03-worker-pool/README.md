# Easy 03 — Worker Pool

Build a fixed-size worker pool that consumes jobs from a channel, runs
each through a user-supplied function, and returns the results.

## Function to implement

```go
// Run dispatches every job in `jobs` to one of `workers` goroutines.
// `do` transforms a job into a result. The returned slice contains the
// results in the SAME ORDER as the input jobs.
func Run[J any, R any](workers int, jobs []J, do func(J) R) []R
```

## Requirements

- Exactly `workers` worker goroutines, not one-per-job.
- Inputs are dispatched via a channel; workers pull until the channel
  closes.
- Output order must match input order, even though workers finish out of
  order.
- `workers <= 0` is invalid — pick a sensible default (e.g. 1).
- No goroutine leaks: every spawned worker must exit before `Run`
  returns.

## Hints

- Send `(index, job)` pairs into the work channel; workers write results
  into `results[index]` (distinct indices ⇒ no race).
- Or: have workers send `(index, result)` on a results channel; the main
  goroutine assembles the output slice.
- `sync.WaitGroup` to know when all workers have drained the work
  channel and exited.
- Close the work channel *after* enqueuing all jobs, *before* `Wait`.

## Stretch goals

- Add a `context.Context` so cancelling stops in-flight work.
- Return an error if `do` panics on any job, using `recover` inside the
  worker.
- Switch to a streaming API: `func Stream[J,R](workers int, jobs <-chan J, do func(J) R) <-chan R`.
