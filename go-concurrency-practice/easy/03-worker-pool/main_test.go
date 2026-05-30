package main

import (
	"reflect"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunOrderPreserved(t *testing.T) {
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := Run(4, jobs, func(n int) int {
		// Sleep different amounts so workers finish out of order.
		time.Sleep(time.Duration(10-n) * time.Millisecond)
		return n * 2
	})
	want := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v want=%v", got, want)
	}
}

func TestRunUsesPoolNotPerJob(t *testing.T) {
	const workers = 3
	var inFlight, peak int32
	jobs := make([]int, 100)
	for i := range jobs {
		jobs[i] = i
	}
	Run(workers, jobs, func(n int) int {
		cur := atomic.AddInt32(&inFlight, 1)
		defer atomic.AddInt32(&inFlight, -1)
		for {
			p := atomic.LoadInt32(&peak)
			if cur <= p || atomic.CompareAndSwapInt32(&peak, p, cur) {
				break
			}
		}
		time.Sleep(time.Millisecond)
		return n
	})
	if peak > workers {
		t.Fatalf("peak concurrency=%d, expected <= %d", peak, workers)
	}
}

func TestRunZeroWorkersTreatedAsOne(t *testing.T) {
	got := Run(0, []int{1, 2, 3}, func(n int) int { return n })
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("got=%v", got)
	}
}
