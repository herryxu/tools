
package datastructure

import (
	"sync"
	"time"
)

type LocalMap struct {
	data     map[string]cacheItem
	mu       sync.RWMutex
	expireCh chan string
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

func NewLocalMap() *LocalMap {
	c := &LocalMap{
		data:     make(map[string]cacheItem),
		expireCh: make(chan string),
	}
	go c.startCleanup()
	return c
}

func (c *LocalMap) Set(key string, value interface{}, expiration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expireTime := time.Now().Add(expiration)
	c.data[key] = cacheItem{
		value: value,
	}
	if expiration.Seconds() > 0 {
		c.data[key] = cacheItem{
			value:      value,
			expiration: expireTime,
		}
		go c.scheduleExpiration(key, expireTime)
	}
}

func (c *LocalMap) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.data[key]
	if ok {
		return item.value, true
	}
	return nil, false
}

func (c *LocalMap) CheckAndSet(key string, value interface{}, expiration time.Duration) bool {
	if _, ok := c.Get(key); !ok {
		c.Set(key, value, expiration)
		return true
	}
	return false
}
func (c *LocalMap) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *LocalMap) startCleanup() {
	for {
		key := <-c.expireCh
		c.Delete(key)
	}
}

func (c *LocalMap) scheduleExpiration(key string, expireTime time.Time) {
	duration := time.Until(expireTime)
	timer := time.NewTimer(duration)
	<-timer.C
	c.expireCh <- key
}
