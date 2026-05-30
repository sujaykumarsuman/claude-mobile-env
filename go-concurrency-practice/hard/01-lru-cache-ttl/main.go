package main

import (
	"fmt"
	"time"
)

type Cache[K comparable, V any] struct {
	// TODO: fields (map, list, mutex, ttl, capacity, done channel, ...)
}

func New[K comparable, V any](capacity int, defaultTTL, sweepEvery time.Duration) *Cache[K, V] {
	// TODO: implement; start sweeper goroutine.
	_ = capacity
	_ = defaultTTL
	_ = sweepEvery
	return &Cache[K, V]{}
}

func (c *Cache[K, V]) Set(key K, val V) {
	// TODO
	_ = key
	_ = val
}

func (c *Cache[K, V]) SetWithTTL(key K, val V, ttl time.Duration) {
	// TODO
	_ = key
	_ = val
	_ = ttl
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	// TODO
	var zero V
	_ = key
	return zero, false
}

func (c *Cache[K, V]) Delete(key K) {
	// TODO
	_ = key
}

func (c *Cache[K, V]) Len() int {
	// TODO
	return 0
}

func (c *Cache[K, V]) Close() {
	// TODO: stop sweeper, idempotent.
}

func main() {
	c := New[string, int](3, 100*time.Millisecond, 20*time.Millisecond)
	defer c.Close()
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	c.Set("d", 4) // should evict "a"
	if _, ok := c.Get("a"); ok {
		fmt.Println("expected a to be evicted")
	}
	fmt.Println(c.Len())
}
