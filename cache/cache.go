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
	content  map[string]cacheEntry
	Add      func(string, []byte)
	Get      func(string) ([]byte, bool)
	reapLoop func()
	mutex    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		content: map[string]cacheEntry{},
		mutex:   &sync.Mutex{},
	}

	c.Add = func(key string, val []byte) {
		c.mutex.Lock()
		c.content[key] = cacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
		c.mutex.Unlock()
	}

	c.Get = func(key string) ([]byte, bool) {
		c.mutex.Lock()
		entry, found := c.content[key]
		c.mutex.Unlock()
		if !found {
			return []byte{}, false
		}

		return entry.val, true
	}

	c.reapLoop = func() {
		ticker := time.NewTicker(time.Duration(interval))
		for range ticker.C {
			for key, entry := range c.content {
				if time.Duration(time.Now().Compare(entry.createdAt)) < interval {
					c.mutex.Lock()
					delete(c.content, key)
					c.mutex.Unlock()
				}
			}
		}
	}

	go c.reapLoop()
	return c
}
