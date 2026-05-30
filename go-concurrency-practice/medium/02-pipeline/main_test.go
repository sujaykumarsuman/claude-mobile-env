package main

import (
	"context"
	"runtime"
	"testing"
	"time"
)

func TestSumOfSquares(t *testing.T) {
	cases := []struct {
		n    int
		want int64
	}{
		{0, 0},
		{1, 1},
		{10, 385},
		{100, 338350},
	}
	for _, c := range cases {
		got := SumOfSquares(context.Background(), c.n)
		if got != c.want {
			t.Fatalf("n=%d got=%d want=%d", c.n, got, c.want)
		}
	}
}

func TestCancellationPropagates(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	gen := Generate(ctx, 1_000_000_000)
	sq := Square(ctx, gen)
	cancel()
	// After cancel, Square's output channel should drain and close.
	timeout := time.After(500 * time.Millisecond)
	for {
		select {
		case _, ok := <-sq:
			if !ok {
				return // closed — good
			}
		case <-timeout:
			t.Fatal("Square did not close after ctx cancellation")
		}
	}
}

func TestNoGoroutineLeak(t *testing.T) {
	before := runtime.NumGoroutine()
	for i := 0; i < 50; i++ {
		SumOfSquares(context.Background(), 100)
	}
	time.Sleep(50 * time.Millisecond)
	after := runtime.NumGoroutine()
	if after > before+5 {
		t.Fatalf("possible leak: before=%d after=%d", before, after)
	}
}
