package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type JobID string

type Job struct {
	ID       JobID
	Priority int
	MaxRetry int
	Backoff  time.Duration
	Run      func(ctx context.Context) error
}

type Result struct {
	ID       JobID
	Attempts int
	Err      error
}

var ErrStopped = errors.New("scheduler stopped")

type Scheduler struct {
	// TODO: fields
}

func New(workers int) *Scheduler {
	// TODO
	_ = workers
	return &Scheduler{}
}

func (s *Scheduler) Start() {
	// TODO
}

func (s *Scheduler) Submit(j Job) error {
	// TODO
	_ = j
	return ErrStopped
}

func (s *Scheduler) Cancel(id JobID) {
	// TODO
	_ = id
}

func (s *Scheduler) Results() <-chan Result {
	// TODO
	return nil
}

func (s *Scheduler) Stop() {
	// TODO
}

func main() {
	s := New(2)
	s.Start()
	go func() {
		for r := range s.Results() {
			fmt.Printf("%s: attempts=%d err=%v\n", r.ID, r.Attempts, r.Err)
		}
	}()
	_ = s.Submit(Job{
		ID:       "hello",
		Priority: 1,
		Run:      func(ctx context.Context) error { fmt.Println("hi"); return nil },
	})
	time.Sleep(100 * time.Millisecond)
	s.Stop()
}
