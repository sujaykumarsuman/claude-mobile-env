package main

import (
	"context"
	"fmt"
)

type Fetcher interface {
	Fetch(ctx context.Context, url string) (body string, links []string, err error)
}

type Page struct {
	URL   string
	Depth int
	Body  string
	Err   error
}

// Crawl fetches the seed and (transitively) the URLs it links to, up
// to maxDepth, with at most `workers` concurrent fetches. Each URL is
// fetched at most once.
func Crawl(ctx context.Context, f Fetcher, seed string, maxDepth, workers int) []Page {
	// TODO: implement
	_ = ctx
	_ = f
	_ = seed
	_ = maxDepth
	_ = workers
	return nil
}

// fakeFetcher serves a tiny static graph for the main demo.
type fakeFetcher map[string]struct {
	body  string
	links []string
}

func (f fakeFetcher) Fetch(_ context.Context, url string) (string, []string, error) {
	if p, ok := f[url]; ok {
		return p.body, p.links, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

func main() {
	f := fakeFetcher{
		"/":     {"home", []string{"/a", "/b"}},
		"/a":    {"about", []string{"/", "/a/sub"}},
		"/b":    {"blog", []string{"/", "/a"}},
		"/a/sub": {"sub", []string{"/"}},
	}
	for _, p := range Crawl(context.Background(), f, "/", 3, 4) {
		fmt.Printf("[d=%d] %s err=%v\n", p.Depth, p.URL, p.Err)
	}
}
