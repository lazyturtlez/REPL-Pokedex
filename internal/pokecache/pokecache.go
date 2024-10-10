package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	m *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	c :=  &Cache{
		cache: make(map[string]cacheEntry),
		m: &sync.Mutex{},
	}
	go c.ReapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.m.Lock()
	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now().UTC(),
	}
	c.m.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.m.Lock()
	entry, exists := c.cache[key]
	c.m.Unlock()
	if !exists {
		return []byte{}, false
	}
	
	return entry.val, true
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Reap(interval)
	}
}

func (c *Cache) Reap(interval time.Duration) {
	c.m.Lock()
	defer c.m.Unlock()
	timePassed := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(timePassed) {
			delete(c.cache, k)
		}
	}
	
}

