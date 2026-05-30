# Hard 03 — Priority Job Scheduler with Retries

Build an in-memory job scheduler that:

- Accepts jobs with a priority (higher = sooner).
- Runs them on a fixed pool of `N` worker goroutines.
- Supports per-job retry with exponential backoff.
- Supports cancellation of an individual job (by ID) and of the
  whole scheduler.
- Reports each job's final outcome on a result channel.

## Types and methods to implement

```go
type JobID string

type Job struct {
	ID       JobID
	Priority int
	MaxRetry int                          // 0 = no retry
	Backoff  time.Duration                // base; double on each retry
	Run      func(ctx context.Context) error
}

type Result struct {
	ID       JobID
	Attempts int
	Err      error
}

type Scheduler struct { /* ... */ }

// New creates a scheduler with `workers` parallel runners.
func New(workers int) *Scheduler

// Start begins processing. Must be called before Submit.
func (s *Scheduler) Start()

// Submit enqueues a job. Returns an error if the scheduler is stopped.
func (s *Scheduler) Submit(j Job) error

// Cancel cancels a pending or in-flight job. Idempotent.
func (s *Scheduler) Cancel(id JobID)

// Results returns the channel on which job outcomes are emitted.
// It is closed when Stop completes.
func (s *Scheduler) Results() <-chan Result

// Stop signals shutdown and waits for in-flight jobs to finish (giving
// them ctx cancellation). After Stop, Results' channel closes.
func (s *Scheduler) Stop()
```

## Requirements

- Higher `Priority` runs before lower. Ties broken by FIFO of submission.
- Exactly `workers` worker goroutines.
- A failed `Run` triggers retry after `Backoff` (then `2*Backoff`, then
  `4*Backoff`, ...) up to `MaxRetry` extra attempts. Final outcome
  (success or last error) is sent on `Results()`.
- `Cancel(id)`:
  - If pending: remove from queue, emit a `Result{Err: ctx.Canceled-ish}`.
  - If in-flight: cancel the per-job `ctx` so `Run` can exit.
  - If unknown/already-done: no-op.
- `Stop()`:
  - Stops accepting new submissions (`Submit` returns an error).
  - Cancels in-flight jobs' contexts.
  - Drains and emits pending jobs' results as canceled.
  - Closes `Results()` exactly once, after every job has had its result
    emitted.
- Safe for concurrent `Submit`/`Cancel` from many goroutines. Race-free.

## Hints

- Use `container/heap` for the priority queue. To break ties FIFO, store
  a monotonically increasing sequence number on each item and compare
  `(priority desc, seq asc)`.
- The PQ is shared mutable state — guard with `sync.Mutex` and signal a
  `sync.Cond` (or a `chan struct{}` "wake-up" signal) when a new item is
  pushed.
- Cancellation map: `map[JobID]context.CancelFunc` guarded by a mutex.
  Store the cancel func at the moment a job is dispatched. Remove on
  completion.
- For retries: schedule re-enqueue with `time.AfterFunc`. Cancel the
  timer if the job is canceled before it fires.
- Lifecycle: track the count of "live" jobs (queued + in-flight +
  pending-retry). `Stop` waits until this hits zero (after canceling
  everything), then closes `Results()`.

## Stretch goals

- Add `Submit` with an estimated deadline; expire jobs that miss it.
- Add per-job timeout independent of `Stop`.
- Persist the queue to disk so jobs survive process restarts.
- Add `Stats()` reporting queue depth, in-flight count, retry count,
  success rate.
