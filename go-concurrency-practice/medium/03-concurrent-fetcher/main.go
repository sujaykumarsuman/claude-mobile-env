package main

import (
	"context"
	"fmt"
)

type Result struct {
	URL  string
	Body []byte
	Err  error
}

type Fetcher interface {
	Fetch(ctx context.Context, url string) ([]byte, error)
}

// FetchAll fetches all URLs concurrently with at most `concurrency`
// in-flight requests. Results are in input order.
func FetchAll(ctx context.Context, f Fetcher, urls []string, concurrency int) []Result {
	// TODO: implement
	_ = ctx
	_ = f
	_ = concurrency
	out := make([]Result, len(urls))
	for i, u := range urls {
		out[i] = Result{URL: u}
	}
	return out
}

type stubFetcher struct{}

func (stubFetcher) Fetch(_ context.Context, url string) ([]byte, error) {
	return []byte("body of " + url), nil
}

func main() {
	urls := []string{"https://a", "https://b", "https://c"}
	for _, r := range FetchAll(context.Background(), stubFetcher{}, urls, 2) {
		fmt.Printf("%s -> %s (err=%v)\n", r.URL, r.Body, r.Err)
	}
}
