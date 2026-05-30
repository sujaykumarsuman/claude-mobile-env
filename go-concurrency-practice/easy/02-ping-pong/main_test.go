package main

import (
	"reflect"
	"testing"
	"time"
)

func TestPingPong(t *testing.T) {
	cases := []struct {
		n    int
		want []string
	}{
		{0, []string{}},
		{1, []string{"ping", "pong"}},
		{3, []string{"ping", "pong", "ping", "pong", "ping", "pong"}},
	}
	for _, c := range cases {
		done := make(chan []string, 1)
		go func(n int) { done <- PingPong(n) }(c.n)
		select {
		case got := <-done:
			if c.n == 0 && len(got) == 0 {
				continue
			}
			if !reflect.DeepEqual(got, c.want) {
				t.Fatalf("n=%d got=%v want=%v", c.n, got, c.want)
			}
		case <-time.After(2 * time.Second):
			t.Fatalf("n=%d timed out — goroutines likely deadlocked", c.n)
		}
	}
}
