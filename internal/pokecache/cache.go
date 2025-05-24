package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cachemap map[string]cacheEntry
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{}
	cache.cachemap = make(map[string]cacheEntry)
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		time.Now(),
		val,
	}
	c.cachemap[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cachemap[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		c.mu.Lock()
		for key, entry := range c.cachemap {
			if time.Since(entry.createdAt) > interval {
				delete(c.cachemap, key)
			}
		}
		c.mu.Unlock()
	}
}
