# Easy 02 — Ping Pong

Two goroutines pass a "ball" back and forth using unbuffered channels.
Practice synchronous handoff and clean termination.

## Functions to implement

```go
// PingPong runs n volleys between a "ping" goroutine and a "pong" goroutine.
// Each volley appends one entry to the returned slice, in order, alternating
// "ping" then "pong". The slice has exactly 2*n entries.
func PingPong(n int) []string
```

## Requirements

- Use exactly two goroutines plus the main goroutine.
- Communication between them must use **unbuffered** channels.
- The output must be deterministic: `["ping", "pong", "ping", "pong", ...]`.
- All goroutines must exit cleanly — no leaks.
- `n == 0` returns an empty slice; no goroutines should hang.

## Hints

- Two unbuffered channels (`pingCh`, `pongCh`), or one channel with a
  baton/turn pattern.
- The simplest correct shape: `ping` reads from `pingCh`, records "ping",
  then sends on `pongCh`. `pong` does the mirror. The main goroutine
  kicks things off by sending the first token.
- Use `sync.WaitGroup` to know when both goroutines have returned.
- Recording into the slice can be done by each goroutine sending its
  label on a results channel, and the main goroutine collecting them.

## Stretch goals

- Add a `context.Context` so cancellation aborts mid-rally and both
  goroutines exit promptly.
- Replace channels with a `sync.Cond` + shared turn variable and compare
  ergonomics.
