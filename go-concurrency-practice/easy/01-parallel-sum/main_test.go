package main

import "testing"

func TestParallelSum(t *testing.T) {
	cases := []struct {
		name    string
		nums    []int64
		workers int
		want    int64
	}{
		{"empty", nil, 4, 0},
		{"single", []int64{42}, 4, 42},
		{"1..100", makeRange(1, 100), 4, 5050},
		{"1..1000 with 1 worker", makeRange(1, 1000), 1, 500500},
		{"1..1000 with 7 workers (uneven)", makeRange(1, 1000), 7, 500500},
		{"zero workers treated as one", makeRange(1, 50), 0, 1275},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ParallelSum(c.nums, c.workers)
			if got != c.want {
				t.Fatalf("ParallelSum=%d want=%d", got, c.want)
			}
		})
	}
}

func makeRange(lo, hi int64) []int64 {
	out := make([]int64, 0, hi-lo+1)
	for i := lo; i <= hi; i++ {
		out = append(out, i)
	}
	return out
}
