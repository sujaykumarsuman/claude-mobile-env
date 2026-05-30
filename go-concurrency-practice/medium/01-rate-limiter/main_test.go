package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestAllowBurst(t *testing.T) {
	l := NewLimiter(1, 3)
	defer l.Stop()
	for i := 0; i < 3; i++ {
		if !l.Allow() {
			t.Fatalf("Allow %d: want true (within burst)", i)
		}
	}
	if l.Allow() {
		t.Fatal("Allow after burst: want false")
	}
}

func TestWaitContextCancel(t *testing.T) {
	l := NewLimiter(0.1, 1)
	defer l.Stop()
	if err := l.Wait(context.Background()); err != nil { // consume the initial token
		t.Fatalf("first Wait: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	start := time.Now()
	err := l.Wait(ctx)
	if err == nil {
		t.Fatal("expected ctx cancellation error")
	}
	if time.Since(start) > 200*time.Millisecond {
		t.Fatalf("Wait took %v; should respect ctx cancellation promptly", time.Since(start))
	}
}

func TestConcurrentSafety(t *testing.T) {
	l := NewLimiter(1000, 100)
	defer l.Stop()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				l.Allow()
			}
		}()
	}
	wg.Wait()
}

func TestRefillRate(t *testing.T) {
	l := NewLimiter(100, 1) // ~10ms per token
	defer l.Stop()
	// Drain initial token.
	l.Allow()
	start := time.Now()
	for i := 0; i < 5; i++ {
		if err := l.Wait(context.Background()); err != nil {
			t.Fatal(err)
		}
	}
	elapsed := time.Since(start)
	if elapsed < 30*time.Millisecond || elapsed > 200*time.Millisecond {
		t.Fatalf("5 tokens at 100/sec took %v; expected ~50ms", elapsed)
	}
}

func TestStopIdempotent(t *testing.T) {
	l := NewLimiter(10, 5)
	l.Stop()
	l.Stop()
}
