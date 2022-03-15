package lru

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	r := ring.New(10)
	for i := 0; i < 10; i++ {
		r.Value = i
		r = r.Next()
	}
	r.Do(func(a any) {
		fmt.Printf("%d,", a)
	})
	fmt.Println("len:", r.Len(), "%:", 1%10, r.Value)
	prev := r.Prev()
	fmt.Println("prev:", prev.Value, "next:", r.Next().Value)
	prev.Prev().Unlink(1)
	r.Do(func(a any) {
		fmt.Printf("%d,", a)
	})
}

func printRing(r *ring.Ring) {
	for r.Next() != nil {

	}
	fmt.Println("")
}

func TestLRU(t *testing.T) {
	c := NewLruCache[int](10)
	for i := 0; i < 10; i++ {
		c.Put(fmt.Sprintf("key%d", i), i)
	}

	c.Each(func(key string, val int) {
		fmt.Printf("%s=%d,", key, val)
	})
	fmt.Println("")
	for j := 0; j < 5; j++ {
		c.Put(fmt.Sprintf("key%d", j), j)
	}

	c.Each(func(key string, val int) {
		fmt.Printf("%s=%d,", key, val)
	})

	fmt.Println("")
}
