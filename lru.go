package lrucache

import (
	"errors"
)

// This package implements a simple LRU cache that uses a linked-list
// for its internal representation.

type Entry struct {
	key   int
	value interface{}
}

type LRUCache struct {
	size int
	head *Entry
}

func NewLRUCache(size int) (*LRUCache, error) {
	if size <= 0 {
		return nil, errors.New("need a size greater than 0!")
	}
	lru := &LRUCache{
		size: size,
		head: &Entry{},
	}
	return lru, nil
}

// Add(key, value) - add entry to the LRU cache
func (c *LRUCache) Add(key int, value int) error {
	return nil
}

// Get(key) - get entry from the cache
func (c *LRUCache) Get(key int) error {
	return nil
}

// Remove(key) - remove entry from the cache
func (c *LRUCache) Remove(key int) error {
	return nil
}
