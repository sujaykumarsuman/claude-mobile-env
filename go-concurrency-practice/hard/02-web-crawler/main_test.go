package main

import (
	"context"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type instrumentedFetcher struct {
	mu      sync.Mutex
	graph   map[string][]string
	calls   map[string]int
	delay   time.Duration
	inFlight int32
	peak    int32
}

func (f *instrumentedFetcher) Fetch(ctx context.Context, url string) (string, []string, error) {
	cur := atomic.AddInt32(&f.inFlight, 1)
	defer atomic.AddInt32(&f.inFlight, -1)
	for {
		p := atomic.LoadInt32(&f.peak)
		if cur <= p || atomic.CompareAndSwapInt32(&f.peak, p, cur) {
			break
		}
	}
	f.mu.Lock()
	f.calls[url]++
	links := f.graph[url]
	f.mu.Unlock()
	if f.delay > 0 {
		select {
		case <-time.After(f.delay):
		case <-ctx.Done():
			return "", nil, ctx.Err()
		}
	}
	return "body:" + url, links, nil
}

func TestCrawlVisitsAll(t *testing.T) {
	f := &instrumentedFetcher{
		graph: map[string][]string{
			"/":  {"/a", "/b"},
			"/a": {"/b", "/c"},
			"/b": {"/c"},
			"/c": {},
		},
		calls: map[string]int{},
	}
	pages := Crawl(context.Background(), f, "/", 5, 4)
	urls := map[string]int{}
	for _, p := range pages {
		urls[p.URL]++
	}
	for _, want := range []string{"/", "/a", "/b", "/c"} {
		if urls[want] != 1 {
			t.Fatalf("URL %s visited %d times, want 1", want, urls[want])
		}
	}
}

func TestCrawlDepthLimit(t *testing.T) {
	f := &instrumentedFetcher{
		graph: map[string][]string{
			"/":  {"/a"},
			"/a": {"/b"},
			"/b": {"/c"},
			"/c": {},
		},
		calls: map[string]int{},
	}
	pages := Crawl(context.Background(), f, "/", 1, 2)
	visited := []string{}
	for _, p := range pages {
		visited = append(visited, p.URL)
	}
	sort.Strings(visited)
	want := []string{"/", "/a"} // depth 0 = "/", depth 1 = "/a", "/b" is depth 2
	if len(visited) != len(want) || visited[0] != want[0] || visited[1] != want[1] {
		t.Fatalf("visited=%v want=%v", visited, want)
	}
}

func TestCrawlConcurrencyBound(t *testing.T) {
	graph := map[string][]string{"/": {}}
	// Build a fan-out graph: / -> 50 children.
	children := make([]string, 50)
	for i := range children {
		u := "/" + string(rune('a'+i%26)) + string(rune('a'+i/26))
		children[i] = u
		graph[u] = nil
	}
	graph["/"] = children
	f := &instrumentedFetcher{graph: graph, calls: map[string]int{}, delay: 5 * time.Millisecond}
	_ = Crawl(context.Background(), f, "/", 2, 4)
	if f.peak > 4 {
		t.Fatalf("peak concurrency = %d, want <= 4", f.peak)
	}
}

func TestCrawlCancellation(t *testing.T) {
	graph := map[string][]string{"/": {}}
	for i := 0; i < 200; i++ {
		u := "/u" + string(rune('a'+i%26))
		graph["/"] = append(graph["/"], u)
		graph[u] = nil
	}
	f := &instrumentedFetcher{graph: graph, calls: map[string]int{}, delay: 50 * time.Millisecond}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	before := runtime.NumGoroutine()
	start := time.Now()
	_ = Crawl(ctx, f, "/", 3, 4)
	if time.Since(start) > 500*time.Millisecond {
		t.Fatalf("Crawl took %v after cancellation; should be prompt", time.Since(start))
	}
	time.Sleep(50 * time.Millisecond)
	after := runtime.NumGoroutine()
	if after > before+5 {
		t.Fatalf("possible leak: before=%d after=%d", before, after)
	}
}
