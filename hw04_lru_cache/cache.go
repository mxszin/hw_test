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
	lock     sync.Mutex
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	newCacheItem := cacheItem{
		key:   key,
		value: value,
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	if listItem, ok := c.items[key]; ok {
		listItem.Value = newCacheItem
		c.queue.MoveToFront(listItem)
		return true
	}

	c.items[key] = c.queue.PushFront(newCacheItem)

	if c.queue.Len() > c.capacity {
		listItem := c.queue.Back()
		c.queue.Remove(listItem)
		delete(c.items, listItem.Value.(cacheItem).key)
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if listItem, ok := c.items[key]; ok {
		c.queue.MoveToFront(listItem)
		return listItem.Value.(cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
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
