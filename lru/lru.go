package lrucache

import (
	"container/list"
	"errors"
)

// This package implements a simple LRU cache that uses a linked-list
// for its internal representation.

type Entry struct {
	key   interface{}
	value interface{}
}

// This is not thread-safe
type LRUCache struct {
	size       int
	entryStore map[interface{}]*list.Element // Keep track of entries
	entryOrder *list.List                    // Keep track of ordering
}

// Factory function to create new LRU cache
func NewLRUCache(size int) (*LRUCache, error) {
	if size <= 0 {
		return nil, errors.New("need a size greater than 0")
	}
	lru := &LRUCache{
		size:       size,
		entryStore: make(map[interface{}]*list.Element),
		entryOrder: list.New(),
	}
	return lru, nil
}

// Add(key, value) - add entry to the LRU cache
func (c *LRUCache) Add(key interface{}, value interface{}) {
	// Check if the value already exists in here
	if val, ok := c.entryStore[key]; ok {
		// Update the ordering
		c.entryOrder.MoveToFront(val)
	}
	// Store in the entry and store in the ordering
	elem := c.entryOrder.PushFront(&Entry{key, value})
	c.entryStore[key] = elem
	if c.entryOrder.Len() > c.size {
		c.evictOldest()
	}
}

// Get(key) - get entry from the cache
func (c *LRUCache) Get(key interface{}) (interface{}, error) {
	if _, ok := c.entryStore[key]; !ok {
		return 0, errors.New("couldn't find key in cache")
	}
	val := c.entryStore[key]
	c.entryOrder.MoveToFront(val)
	return val.Value.(*Entry).value, nil
}

// Remove(key) - remove entry from the cache
func (c *LRUCache) Remove(key interface{}) error {
	return nil
}

// evictOldest() - removes the oldest entry from the cache
func (c *LRUCache) evictOldest() {
	entry := c.entryOrder.Remove(c.entryOrder.Back())
	delete(c.entryStore, entry.(Entry).key)
}
