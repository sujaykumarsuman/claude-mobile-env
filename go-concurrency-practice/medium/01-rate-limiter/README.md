# Medium 01 — Token Bucket Rate Limiter

Implement a token-bucket rate limiter usable from many goroutines.

## Type and methods to implement

```go
type Limiter struct { /* ... */ }

// NewLimiter creates a bucket that refills at `rate` tokens/sec, with
// a maximum of `burst` tokens. Start with `burst` tokens available.
func NewLimiter(rate float64, burst int) *Limiter

// Allow returns true and consumes one token if a token is available,
// otherwise returns false immediately.
func (l *Limiter) Allow() bool

// Wait blocks until a token is available or ctx is canceled. Returns
// ctx.Err() on cancellation, nil on success.
func (l *Limiter) Wait(ctx context.Context) error

// Stop releases any resources (e.g. tickers). Safe to call once.
func (l *Limiter) Stop()
```

## Requirements

- Safe for concurrent use from many goroutines.
- `Allow` must be non-blocking.
- `Wait` must respect `ctx` cancellation promptly.
- Token count must never exceed `burst`.
- `Stop` must not leak the refill goroutine.

## Hints

- A buffered channel of size `burst` is a clean model: each element is
  a token. Refill = a `time.Ticker` goroutine that tries to push a
  token, non-blockingly skipping if the channel is full.
- `Allow` = `select { case <-tokens: return true; default: return false }`.
- `Wait` = `select` on `<-tokens` and `<-ctx.Done()`.
- `Stop` closes a `done` channel that the refill goroutine selects on
  alongside the ticker. Use `sync.Once` so `Stop` is idempotent.

## Stretch goals

- Add `Reserve(n int)` returning the duration to wait for `n` tokens.
- Replace the channel with an arithmetic model (last refill time +
  current tokens) protected by `sync.Mutex`; compare overhead with
  the channel-based version.
- Verify burst behavior: under steady load at `rate`, queueing delay
  should approach zero; under bursts up to `burst`, no delay.
