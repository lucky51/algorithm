package lru

import (
	"container/list"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

// lruCache lru cache
type lruCache[T any] struct {
	l        *list.List
	h        map[string]*list.Element
	capacity int
}
type lruCacheItem[T any] struct {
	key  string
	data T
}

// Each of items
func (c *lruCache[T]) Each(f func(key string, val T)) {
	temp := c.l.Front()
	for i := 0; i < c.l.Len(); i++ {
		if temp == nil {
			return
		}
		cur := (temp.Value).(*lruCacheItem[T])
		f(cur.key, cur.data)
		temp = temp.Next()
	}
}

// get item
func (c *lruCache[T]) Get(key string) (T, error) {
	var res T
	if item, ok := c.h[key]; ok {
		c.l.MoveToFront(item)
		return ((*item).Value).(*lruCacheItem[T]).data, nil
	}
	return res, ErrNotFound
}

// put a element
func (c *lruCache[T]) Put(key string, val T) error {
	if cur, ok := c.h[key]; ok {
		cur.Value.(*lruCacheItem[T]).data = val
		c.l.MoveToFront(cur)
	} else {
		if c.l.Len() >= c.capacity {
			b := c.l.Back()
			c.l.Remove(b)
			item := b.Value.(*lruCacheItem[T])
			delete(c.h, item.key)
		}
		c.h[key] = c.l.PushFront(&lruCacheItem[T]{
			data: val,
			key:  key,
		})
	}
	return nil
}

// delete a item
func (c *lruCache[T]) Delete(key string) (T, error) {
	var res T
	if item, ok := c.h[key]; ok {
		delete(c.h, key)
		return c.l.Remove(item).(lruCacheItem[T]).data, nil
	}
	return res, ErrNotFound
}

// NewLruCache create lru cache
func NewLruCache[T any](capacity int) *lruCache[T] {
	return &lruCache[T]{
		l:        list.New(),
		capacity: capacity,
		h:        make(map[string]*list.Element),
	}
}
