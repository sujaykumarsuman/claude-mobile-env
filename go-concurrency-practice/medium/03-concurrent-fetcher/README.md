# Medium 03 — Concurrent Fetcher with Bounded Concurrency

Fetch `N` URLs concurrently with a cap on simultaneous in-flight
requests, returning per-URL results in input order. Respect a global
`context.Context` for cancellation/timeout.

## Types and function to implement

```go
type Result struct {
	URL   string
	Body  []byte
	Err   error
}

// Fetcher abstracts HTTP fetching so tests can use a fake.
type Fetcher interface {
	Fetch(ctx context.Context, url string) ([]byte, error)
}

// FetchAll fetches every URL concurrently, capping in-flight requests
// at `concurrency`. Results are in input order. If ctx is canceled,
// in-flight fetches must abort and FetchAll returns promptly.
func FetchAll(ctx context.Context, f Fetcher, urls []string, concurrency int) []Result
```

## Requirements

- Never more than `concurrency` in-flight calls to `f.Fetch`.
- Output `Results[i]` corresponds to `urls[i]` (order preserved).
- A single slow or hanging fetch must not stall the others.
- `ctx` cancellation must propagate into in-flight `f.Fetch` calls.
- `concurrency <= 0` treated as `1`. Empty `urls` returns an empty slice.

## Hints

- A semaphore as a buffered channel of size `concurrency`: acquire by
  sending, release by receiving in a `defer`.
- Spawn one goroutine per URL but block on the semaphore before doing
  work. Write each result to `results[i]` (distinct index ⇒ no race).
- Use `sync.WaitGroup` to wait for all to finish.
- Forward the caller's `ctx` directly into `f.Fetch`.

## Stretch goals

- Add a per-URL retry policy with exponential backoff that aborts on
  `ctx.Done()`.
- Add a streaming API: `func StreamAll(...) <-chan Result` that emits
  results as they arrive (order no longer guaranteed; include the
  original index in `Result`).
- Add a deduper: if the same URL appears twice, fetch it once.
