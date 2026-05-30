package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBasicGetSet(t *testing.T) {
	c := New[string, int](10, time.Second, 100*time.Millisecond)
	defer c.Close()
	c.Set("a", 1)
	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Fatalf("Get a = (%d,%v)", v, ok)
	}
	if _, ok := c.Get("missing"); ok {
		t.Fatal("missing key reported present")
	}
}

func TestLRUEviction(t *testing.T) {
	c := New[string, int](3, time.Second, time.Second)
	defer c.Close()
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	// Touch a so b is now the LRU.
	if _, ok := c.Get("a"); !ok {
		t.Fatal("a should be present")
	}
	c.Set("d", 4) // should evict b
	if _, ok := c.Get("b"); ok {
		t.Fatal("b should have been evicted")
	}
	for _, k := range []string{"a", "c", "d"} {
		if _, ok := c.Get(k); !ok {
			t.Fatalf("%s should be present", k)
		}
	}
}

func TestTTLLazyExpiration(t *testing.T) {
	c := New[string, int](10, 30*time.Millisecond, time.Hour) // no sweep
	defer c.Close()
	c.Set("a", 1)
	time.Sleep(60 * time.Millisecond)
	if _, ok := c.Get("a"); ok {
		t.Fatal("a should have expired (lazy)")
	}
}

func TestTTLBackgroundSweep(t *testing.T) {
	c := New[string, int](10, 20*time.Millisecond, 10*time.Millisecond)
	defer c.Close()
	c.Set("a", 1)
	c.Set("b", 2)
	time.Sleep(80 * time.Millisecond)
	if n := c.Len(); n != 0 {
		t.Fatalf("Len after sweep = %d, want 0", n)
	}
}

func TestConcurrentAccess(t *testing.T) {
	c := New[int, int](100, time.Second, 50*time.Millisecond)
	defer c.Close()
	var wg sync.WaitGroup
	for w := 0; w < 16; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				k := (w*1000 + i) % 200
				c.Set(k, i)
				c.Get(k)
				if i%10 == 0 {
					c.Delete(k)
				}
			}
		}(w)
	}
	wg.Wait()
}

func TestCloseIdempotent(t *testing.T) {
	c := New[string, int](10, time.Second, 50*time.Millisecond)
	c.Close()
	c.Close()
}

func ExampleCache_Set() {
	c := New[string, int](2, time.Second, time.Second)
	defer c.Close()
	c.Set("x", 7)
	v, _ := c.Get("x")
	fmt.Println(v)
	// Output: 7
}
