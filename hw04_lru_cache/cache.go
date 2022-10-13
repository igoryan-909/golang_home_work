package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

var mutex = &sync.Mutex{}

func (c *lruCache) Set(key Key, value interface{}) bool {
	cacheItem := cacheItem{key, value}
	mutex.Lock()
	defer mutex.Unlock()
	if c.items[key] != nil {
		c.items[key].Value = cacheItem
		c.queue.MoveToFront(c.items[key])
		return true
	}
	if c.capacity <= c.queue.Len() {
		c.queue.Remove(c.queue.Back())
	}
	c.items[key] = c.queue.PushFront(cacheItem)

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	if c.items[key] != nil {
		c.queue.MoveToFront(c.items[key])
		return c.items[key].Value.(cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	mutex.Lock()
	defer mutex.Unlock()
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
