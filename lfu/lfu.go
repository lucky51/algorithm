package lfu

//TODO: lfu unimplement
import "container/list"

type lfuCache[T any] struct {
	keyfreqh map[string]int
	freqkeyh map[int]*list.List
	kv       map[string]lfuCacheItem[T]
	capacity int
}
type lfuCacheItem[T any] struct {
	data T
	ref  *list.Element
}

// Get
func (c *lfuCache[T]) Get(key string) T {
	var res T
	if v, ok := c.kv[key]; ok {
		if ok {
			oldfreq := c.keyfreqh[key]
			if ls, ok := c.freqkeyh[oldfreq]; ok {
				ls.Remove(v.ref)
				if ls1, ok1 := c.freqkeyh[oldfreq+1]; ok1 {
					v.ref = ls1.PushFront(v.ref.Value)
				} else {
					ls := list.New()
					v.ref = ls.PushFront(key)
					c.freqkeyh[oldfreq+1] = ls
				}
			}

		}
	}
	return res
}

// Put
func (c *lfuCache[T]) Put(key string) {

}

func NewLfuCache[T any](capacity int) *lfuCache[T] {
	return &lfuCache[T]{
		keyfreqh: make(map[string]int),
		freqkeyh: make(map[int]*list.List),
		kv:       make(map[string]lfuCacheItem[T], capacity),
		capacity: capacity,
	}
}
