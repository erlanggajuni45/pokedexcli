package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mx           *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheEntries: make(map[string]cacheEntry),
		mx:           &sync.RWMutex{},
	}
	go c.realLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	entry, exists := c.cacheEntries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) realLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mx.Lock()

		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mx.Unlock()
	}
}
