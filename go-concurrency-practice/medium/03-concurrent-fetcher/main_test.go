package main

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

type fakeFetcher struct {
	delay      time.Duration
	failOn     map[string]error
	inFlight   int32
	peak       int32
	totalCalls int32
}

func (f *fakeFetcher) Fetch(ctx context.Context, url string) ([]byte, error) {
	atomic.AddInt32(&f.totalCalls, 1)
	cur := atomic.AddInt32(&f.inFlight, 1)
	defer atomic.AddInt32(&f.inFlight, -1)
	for {
		p := atomic.LoadInt32(&f.peak)
		if cur <= p || atomic.CompareAndSwapInt32(&f.peak, p, cur) {
			break
		}
	}
	select {
	case <-time.After(f.delay):
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	if err, ok := f.failOn[url]; ok {
		return nil, err
	}
	return []byte(url + "-body"), nil
}

func TestFetchAllOrderAndConcurrency(t *testing.T) {
	urls := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	f := &fakeFetcher{delay: 20 * time.Millisecond}
	res := FetchAll(context.Background(), f, urls, 3)
	if len(res) != len(urls) {
		t.Fatalf("len=%d want=%d", len(res), len(urls))
	}
	for i, r := range res {
		if r.URL != urls[i] {
			t.Fatalf("res[%d].URL=%q want=%q", i, r.URL, urls[i])
		}
		if r.Err != nil {
			t.Fatalf("res[%d].Err=%v", i, r.Err)
		}
		if string(r.Body) != urls[i]+"-body" {
			t.Fatalf("res[%d].Body=%q", i, r.Body)
		}
	}
	if f.peak > 3 {
		t.Fatalf("peak concurrency=%d, want <= 3", f.peak)
	}
}

func TestFetchAllPropagatesPerURLError(t *testing.T) {
	urls := []string{"ok", "boom", "ok2"}
	boom := errors.New("boom")
	f := &fakeFetcher{failOn: map[string]error{"boom": boom}}
	res := FetchAll(context.Background(), f, urls, 2)
	if res[0].Err != nil || res[2].Err != nil {
		t.Fatalf("unexpected errors on ok urls: %v %v", res[0].Err, res[2].Err)
	}
	if !errors.Is(res[1].Err, boom) {
		t.Fatalf("res[1].Err=%v want=%v", res[1].Err, boom)
	}
}

func TestFetchAllCancellation(t *testing.T) {
	urls := make([]string, 50)
	for i := range urls {
		urls[i] = "u"
	}
	f := &fakeFetcher{delay: 200 * time.Millisecond}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	start := time.Now()
	res := FetchAll(ctx, f, urls, 4)
	if len(res) != len(urls) {
		t.Fatalf("len=%d want=%d", len(res), len(urls))
	}
	if time.Since(start) > 500*time.Millisecond {
		t.Fatalf("FetchAll took %v; cancellation should be prompt", time.Since(start))
	}
}

func TestFetchAllEmpty(t *testing.T) {
	res := FetchAll(context.Background(), &fakeFetcher{}, nil, 4)
	if len(res) != 0 {
		t.Fatalf("want empty, got %v", res)
	}
}
