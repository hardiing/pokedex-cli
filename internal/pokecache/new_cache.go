package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	m        map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{}
	c.m = make(map[string]cacheEntry)
	c.interval = interval
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.m[key]
	if ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.m {
			if time.Since(v.createdAt) > c.interval {
				delete(c.m, k)
			}
		}
		c.mu.Unlock()
	}
}
