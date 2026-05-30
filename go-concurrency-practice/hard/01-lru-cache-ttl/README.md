# Hard 01 — Concurrent LRU Cache with TTL

Build a thread-safe LRU cache where each entry has a TTL. When capacity
is reached, evict the least-recently-used entry. A background goroutine
should periodically sweep expired entries.

## Type and methods to implement

```go
type Cache[K comparable, V any] struct { /* ... */ }

// New creates a cache with given capacity (max entries), default ttl,
// and sweep interval for the background expiration goroutine.
func New[K comparable, V any](capacity int, defaultTTL, sweepEvery time.Duration) *Cache[K, V]

// Set inserts/updates key with the default TTL.
func (c *Cache[K, V]) Set(key K, val V)

// SetWithTTL inserts/updates key with a custom TTL.
func (c *Cache[K, V]) SetWithTTL(key K, val V, ttl time.Duration)

// Get returns (value, true) if key is present and not expired, marking
// it most-recently-used. Returns (zero, false) otherwise.
func (c *Cache[K, V]) Get(key K) (V, bool)

// Delete removes a key. No-op if absent.
func (c *Cache[K, V]) Delete(key K)

// Len returns the number of non-expired entries (approximate is fine
// between sweeps; exact after a sweep).
func (c *Cache[K, V]) Len() int

// Close stops the background sweeper. Idempotent.
func (c *Cache[K, V]) Close()
```

## Requirements

- All operations safe for concurrent use.
- O(1) `Get`, `Set`, `Delete` (amortized).
- LRU eviction when `Len() > capacity` after an insertion.
- TTL enforced both lazily (in `Get`) and proactively (background
  sweep).
- `Close` stops the sweeper goroutine; further `Set`/`Get` may either
  panic or no-op — pick a policy and document it.
- No goroutine leaks. Race-free under `go test -race`.

## Hints

- Classic LRU = hash map + doubly-linked list. The map holds
  `*entry`, where `entry` is a list node holding `key, value, expiresAt`.
- Wrap the whole structure in a single `sync.Mutex`. Don't try clever
  per-entry locks until you've got the simple version green.
- Background sweep: a `time.Ticker` in a separate goroutine; on each
  tick, lock, walk the list/map, drop expired entries. Stop on a `done`
  channel.
- Use `sync.Once` for `Close`.
- The standard library has `container/list` if you don't want to roll
  your own doubly-linked list.

## Stretch goals

- Add metrics: hits, misses, evictions, expirations, exposed via
  `Stats()`.
- Add `GetOrLoad(key, loader func() (V, error))` with single-flight
  semantics: concurrent misses for the same key call `loader` once.
- Benchmark mutex-based vs `sync.RWMutex` vs sharded map under
  read-heavy / write-heavy workloads.
