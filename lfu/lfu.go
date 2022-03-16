package lfu

import (
	"container/list"
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
)

type lfuCache[T any] struct {
	keyfreqh map[string]int
	freqkeyh map[int]*list.List
	kv       map[string]*lfuCacheItem[T]
	capacity int
	minfreq  int
}

type lfuCacheItem[T any] struct {
	data T
	ref  *list.Element
}

func (c *lfuCache[T]) Print() {
	fmt.Println("==============print================")
	fmt.Println("key - value")
	for k, v := range c.kv {
		fmt.Printf("%s,%v ->", k, v.data)
	}
	fmt.Println("\nfreq key")
	for fk, fv := range c.freqkeyh {
		fmt.Printf("%d=>", fk)
		for f := fv.Front(); f != nil; f = f.Next() {
			fmt.Printf("%s,", f.Value)
		}
	}
	fmt.Println("\nkey freq")
	for kk, kv := range c.keyfreqh {
		fmt.Printf("%s,%d ->", kk, kv)
	}
	fmt.Printf("\n=======cap:%d,minfreq:%d=========== \r\n", c.capacity, c.minfreq)
}

// Get
func (c *lfuCache[T]) Get(key string) (T, error) {
	var res T
	if v, ok := c.kv[key]; ok {
		//
		c.increaseFreq(v)
		res = v.data
	}
	return res, ErrNotFound
}

// increaseFreq increase freq
func (c *lfuCache[T]) increaseFreq(item *lfuCacheItem[T]) {
	key := item.ref.Value.(string)
	oldfreq := c.keyfreqh[key]
	c.keyfreqh[key]++
	ls, _ := c.freqkeyh[oldfreq]
	ls.Remove(item.ref)
	if ls.Len() == 0 {
		delete(c.freqkeyh, oldfreq)
		if c.minfreq == oldfreq {
			c.minfreq = oldfreq + 1
		}
	}
	if ls1, ok1 := c.freqkeyh[oldfreq+1]; ok1 {
		item.ref = ls1.PushFront(key)
	} else {
		ls := list.New()
		item.ref = ls.PushFront(key)
		c.freqkeyh[oldfreq+1] = ls
	}
}
func (c *lfuCache[T]) removeMinFreq() {
	minfreq := c.freqkeyh[c.minfreq]
	if minfreq.Len() > 0 {
		last := minfreq.Back()
		minfreq.Remove(last)
		delete(c.kv, last.Value.(string))
		delete(c.keyfreqh, last.Value.(string))
	} else {
		delete(c.freqkeyh, c.minfreq)
	}
}

// Put
func (c *lfuCache[T]) Put(key string, v T) {
	if v1, ok := c.kv[key]; ok {
		c.increaseFreq(v1)
		v1.data = v
	} else {
		if len(c.kv) >= c.capacity {
			c.removeMinFreq()
		}
		citem := &lfuCacheItem[T]{
			data: v,
		}
		c.kv[key] = citem
		c.keyfreqh[key] = 1
		c.minfreq = 1
		if minfreq, ok := c.freqkeyh[c.minfreq]; ok {
			citem.ref = minfreq.PushFront(key)
		} else {
			minfreq = list.New()
			citem.ref = minfreq.PushFront(key)
			c.freqkeyh[c.minfreq] = minfreq
		}
	}
}

func NewLfuCache[T any](capacity int) *lfuCache[T] {
	return &lfuCache[T]{
		keyfreqh: make(map[string]int),
		freqkeyh: make(map[int]*list.List),
		kv:       make(map[string]*lfuCacheItem[T], capacity),
		capacity: capacity,
	}
}
