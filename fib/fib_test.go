package fib

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestRecursiveFib(t *testing.T) {
	fmt.Println(fib(12))
}

func TestLoopFib(t *testing.T) {
	fmt.Println(fibloop(12))
}

func TestLength(t *testing.T) {
	a := byte(3)
	fmt.Printf("%b\n", a)
	fmt.Printf("len:%d\n", unsafe.Sizeof(a))
	fmt.Printf("int len:%d\n", unsafe.Sizeof(new(int)))
	fmt.Printf("bool len:%d\n", unsafe.Sizeof(true))
}
