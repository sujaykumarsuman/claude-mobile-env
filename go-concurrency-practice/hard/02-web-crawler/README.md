# Hard 02 — Bounded Concurrent Web Crawler

Crawl a graph of pages starting from a seed URL, in parallel, with:

- A max-depth limit (depth 0 = seed only).
- A cap on simultaneous workers.
- Deduplication: each URL fetched at most once.
- Cancellation via `context.Context`.
- No goroutine leaks; the call returns when the frontier is exhausted
  (or `ctx` is canceled).

## Types and function to implement

```go
// Fetcher returns the body and the list of outbound URLs found in it.
type Fetcher interface {
	Fetch(ctx context.Context, url string) (body string, links []string, err error)
}

type Page struct {
	URL   string
	Depth int
	Body  string
	Err   error
}

// Crawl returns one Page per URL visited (in any order). It fetches
// each URL at most once, no deeper than maxDepth, with at most
// `workers` concurrent fetches.
func Crawl(ctx context.Context, f Fetcher, seed string, maxDepth, workers int) []Page
```

## Requirements

- At most `workers` goroutines calling `f.Fetch` simultaneously.
- A URL is fetched exactly once, even if discovered through multiple
  parents.
- `Crawl` returns when (a) no more URLs are pending, or (b) `ctx` is
  canceled. It must not deadlock or leak goroutines in either case.
- The seed has depth 0. Links found on a page at depth `d` are at depth
  `d+1` and are only enqueued if `d+1 <= maxDepth`.

## The deadlock trap

The naive shape is:

```go
// workers read from `frontier`, push discovered links back to `frontier`
```

This deadlocks if `frontier` is unbuffered (workers block trying to push
while no-one is reading) and is hard to terminate cleanly (when do you
close `frontier`? Only when every worker is idle AND queue is empty —
not when "no one is sending").

Two clean shapes:

1. **Pending counter**: an atomic counter of `pending` URLs. Increment
   when enqueuing, decrement when a fetch completes. When it reaches
   zero, close the frontier channel.
2. **Coordinator goroutine**: a central loop that owns the seen-set,
   reads "discovered links" on one channel, and dispatches work on
   another. Track in-flight count; terminate when in-flight == 0 and
   queue empty.

## Hints

- Seen-set: `map[string]bool` guarded by a mutex, OR owned solely by the
  coordinator goroutine (no mutex needed).
- Use a buffered work channel + a semaphore (`chan struct{}` of size
  `workers`) if you prefer one goroutine per URL.
- Every send must `select` on `ctx.Done()` to avoid blocking on a full
  channel after cancellation.
- Use `sync.WaitGroup` to know when all fetches are done.

## Stretch goals

- Add a same-host filter (only follow links within the seed's host).
- Add per-URL retries with backoff that respect `ctx`.
- Stream results: `func CrawlStream(...) <-chan Page` that emits each
  page as it's fetched.
- Add a "robots.txt"-style allow predicate.
