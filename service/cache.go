package service

import "sync"

type CacheManager struct {
	cacheMap map[string]interface{}
	lock     *sync.RWMutex
}

func (c *CacheManager)GetCache(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	elem, ok := c.cacheMap[key]
	if !ok {
		return nil
	}
	return elem
}

func (c *CacheManager) SetCache(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cacheMap[key] = value
}

func NewCacheManager() (*CacheManager) {
	return &CacheManager{cacheMap:make(map[string]interface{}), lock: &sync.RWMutex{}}
}