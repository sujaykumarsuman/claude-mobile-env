package main

import (
	"context"
	"fmt"
	"time"
)

type Limiter struct {
	// TODO: fields
}

// NewLimiter constructs a token-bucket limiter at `rate` tokens/sec
// with maximum `burst` tokens, starting full.
func NewLimiter(rate float64, burst int) *Limiter {
	// TODO: implement
	_ = rate
	_ = burst
	return &Limiter{}
}

func (l *Limiter) Allow() bool {
	// TODO: implement (non-blocking)
	return false
}

func (l *Limiter) Wait(ctx context.Context) error {
	// TODO: implement (blocking with ctx cancellation)
	_ = ctx
	return nil
}

func (l *Limiter) Stop() {
	// TODO: stop refill goroutine; must be idempotent
}

func main() {
	l := NewLimiter(5, 3)
	defer l.Stop()
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		if err := l.Wait(ctx); err != nil {
			return
		}
		fmt.Printf("%d at %v\n", i, time.Now().Format("15:04:05.000"))
	}
}
