package lfu

import "testing"

func TestLfu(t *testing.T) {
	c := NewLfuCache[int](5)
	c.Put("key1", 1)
	c.Put("key2", 2)
	c.Put("key3", 3)
	c.Put("key4", 4)
	c.Put("key5", 5)

	c.Print()
	c.Get("key1")
	c.Get("key2")
	c.Print()
	c.Put("key6", 6)
	c.Put("key7", 7)
	c.Print()
	c.Get("key7")
	c.Get("key1")
	c.Get("key7")
	c.Get("key1")
	c.Print()
	c.Put("key3", 3)
	c.Print()
}
