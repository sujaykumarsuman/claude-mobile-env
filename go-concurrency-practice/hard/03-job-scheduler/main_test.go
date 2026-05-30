package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func collect(t *testing.T, s *Scheduler, want int, timeout time.Duration) []Result {
	t.Helper()
	out := make([]Result, 0, want)
	deadline := time.After(timeout)
	for len(out) < want {
		select {
		case r, ok := <-s.Results():
			if !ok {
				return out
			}
			out = append(out, r)
		case <-deadline:
			t.Fatalf("timed out waiting for results; got %d/%d", len(out), want)
		}
	}
	return out
}

func TestPriorityOrdering(t *testing.T) {
	s := New(1) // single worker so ordering is observable
	s.Start()
	defer s.Stop()

	var mu sync.Mutex
	order := []JobID{}
	mk := func(id JobID, prio int) Job {
		return Job{
			ID:       id,
			Priority: prio,
			Run: func(ctx context.Context) error {
				mu.Lock()
				order = append(order, id)
				mu.Unlock()
				return nil
			},
		}
	}
	// Submit low, medium, high — high should run first (after the
	// worker picks them all up).
	if err := s.Submit(mk("low", 1)); err != nil {
		t.Fatal(err)
	}
	if err := s.Submit(mk("med", 5)); err != nil {
		t.Fatal(err)
	}
	if err := s.Submit(mk("high", 10)); err != nil {
		t.Fatal(err)
	}
	collect(t, s, 3, 2*time.Second)
	mu.Lock()
	defer mu.Unlock()
	// First one might already have started; remaining two should be by priority.
	if len(order) != 3 {
		t.Fatalf("order=%v", order)
	}
	// At minimum, "high" should not be last.
	if order[2] == "high" {
		t.Fatalf("priority ignored: order=%v", order)
	}
}

func TestRetryWithBackoff(t *testing.T) {
	s := New(2)
	s.Start()
	defer s.Stop()

	var attempts int32
	want := errors.New("transient")
	if err := s.Submit(Job{
		ID:       "retry",
		MaxRetry: 3,
		Backoff:  5 * time.Millisecond,
		Run: func(ctx context.Context) error {
			n := atomic.AddInt32(&attempts, 1)
			if n < 3 {
				return want
			}
			return nil
		},
	}); err != nil {
		t.Fatal(err)
	}
	res := collect(t, s, 1, 2*time.Second)
	if res[0].Err != nil {
		t.Fatalf("Err=%v want nil (succeeds on 3rd attempt)", res[0].Err)
	}
	if res[0].Attempts != 3 {
		t.Fatalf("Attempts=%d want 3", res[0].Attempts)
	}
}

func TestRetryExhausted(t *testing.T) {
	s := New(2)
	s.Start()
	defer s.Stop()
	want := errors.New("permanent")
	_ = s.Submit(Job{
		ID:       "bad",
		MaxRetry: 2,
		Backoff:  2 * time.Millisecond,
		Run:      func(ctx context.Context) error { return want },
	})
	res := collect(t, s, 1, 2*time.Second)
	if !errors.Is(res[0].Err, want) {
		t.Fatalf("Err=%v want=%v", res[0].Err, want)
	}
	if res[0].Attempts != 3 { // 1 try + 2 retries
		t.Fatalf("Attempts=%d want 3", res[0].Attempts)
	}
}

func TestCancelInFlight(t *testing.T) {
	s := New(1)
	s.Start()
	defer s.Stop()

	started := make(chan struct{})
	_ = s.Submit(Job{
		ID:       "slow",
		Priority: 1,
		Run: func(ctx context.Context) error {
			close(started)
			<-ctx.Done()
			return ctx.Err()
		},
	})
	<-started
	s.Cancel("slow")
	res := collect(t, s, 1, 1*time.Second)
	if res[0].Err == nil {
		t.Fatal("expected non-nil err after cancel")
	}
}

func TestStopClosesResults(t *testing.T) {
	s := New(2)
	s.Start()
	_ = s.Submit(Job{ID: "a", Run: func(ctx context.Context) error { return nil }})
	_ = s.Submit(Job{ID: "b", Run: func(ctx context.Context) error { return nil }})
	go func() {
		time.Sleep(50 * time.Millisecond)
		s.Stop()
	}()
	count := 0
	for range s.Results() {
		count++
	}
	if count < 2 {
		t.Fatalf("got %d results before close, want >= 2", count)
	}
}

func TestSubmitAfterStop(t *testing.T) {
	s := New(1)
	s.Start()
	s.Stop()
	err := s.Submit(Job{ID: "late", Run: func(ctx context.Context) error { return nil }})
	if err == nil {
		t.Fatal("Submit after Stop should return error")
	}
}

func ExampleScheduler_Submit() {
	s := New(1)
	s.Start()
	_ = s.Submit(Job{ID: "x", Run: func(ctx context.Context) error {
		fmt.Println("ran x")
		return nil
	}})
	r := <-s.Results()
	fmt.Println(r.ID, r.Err)
	s.Stop()
	// Output:
	// ran x
	// x <nil>
}
