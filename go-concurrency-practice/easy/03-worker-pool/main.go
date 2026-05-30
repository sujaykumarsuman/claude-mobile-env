package main

import "fmt"

// Run dispatches each job to a worker pool of size `workers`. The result
// slice is in input order. workers <= 0 should be treated as 1.
func Run[J any, R any](workers int, jobs []J, do func(J) R) []R {
	// TODO: implement
	_ = workers
	_ = do
	out := make([]R, len(jobs))
	return out
}

func main() {
	jobs := []int{1, 2, 3, 4, 5}
	out := Run(3, jobs, func(n int) int { return n * n })
	fmt.Println(out)
}
