package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	content map[string]cacheEntry
	mutex   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		content: map[string]cacheEntry{},
		mutex:   &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	entry, found := c.content[key]
	c.mutex.Unlock()
	if !found {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	c.content[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mutex.Unlock()
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.content {
			if time.Duration(time.Now().Compare(entry.createdAt)) < interval {
				delete(c.content, key)
			}
		}
		c.mutex.Unlock()
	}
}
