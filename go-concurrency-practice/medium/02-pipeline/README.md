# Medium 02 — Multi-Stage Pipeline

Build a 3-stage pipeline using channels. Practice the canonical Go
pipeline pattern: each stage is a goroutine that reads from an input
channel and writes to an output channel, closing its output when input
is exhausted.

## Functions to implement

```go
// Generate emits ints 1..n on the returned channel, then closes it.
// Respect ctx cancellation: stop emitting if ctx is canceled.
func Generate(ctx context.Context, n int) <-chan int

// Square reads ints from in and writes their squares to the returned
// channel, closing it when in is closed or ctx is canceled.
func Square(ctx context.Context, in <-chan int) <-chan int

// Sum reads ints from in until it closes (or ctx cancels) and returns
// the total.
func Sum(ctx context.Context, in <-chan int) int64

// SumOfSquares wires Generate -> Square -> Sum and returns the total
// of squares 1^2 + 2^2 + ... + n^2.
func SumOfSquares(ctx context.Context, n int) int64
```

## Requirements

- Every stage closes its output channel when done — no exceptions.
- Cancellation via `ctx` must propagate through the whole pipeline; no
  goroutine should outlive the call after `ctx` is canceled.
- No goroutine leaks. If a downstream stage stops consuming, the
  upstream stage must still be able to exit (use `select` with
  `ctx.Done()` on every send).

## Hints

- The canonical send is:
  ```go
  select {
  case out <- v:
  case <-ctx.Done():
      return
  }
  ```
- Pair every channel-producing function with `defer close(out)`.
- For `Sum`, `range` over `in` is enough — but also `select` on
  `ctx.Done()` if you want to abort early.

## Stretch goals

- Add a generic `Map[T,U](ctx, in <-chan T, fn func(T) U) <-chan U`.
- Add a `FanOut[T](ctx, in <-chan T, n int) []<-chan T` and a matching
  `FanIn[T](ctx, ins ...<-chan T) <-chan T` and parallelize `Square`.
- Verify with `go test -race`.
