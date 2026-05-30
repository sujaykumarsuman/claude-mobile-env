package main

import "fmt"

// ParallelSum returns the sum of nums computed concurrently with `workers`
// goroutines. workers <= 0 should be treated as 1.
func ParallelSum(nums []int64, workers int) int64 {
	// TODO: implement
	_ = nums
	_ = workers
	return 0
}

func main() {
	nums := make([]int64, 1_000_000)
	for i := range nums {
		nums[i] = int64(i + 1)
	}
	fmt.Println(ParallelSum(nums, 8))
}
