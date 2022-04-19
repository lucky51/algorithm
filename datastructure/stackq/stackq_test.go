package stackq

import (
	"fmt"
	"testing"

	"github.com/lucky51/pkg/stack"
)

func TestCalculate(t *testing.T) {
	fmt.Println(calculate("(((2+2)+(1+2))/2)-1*2"))
}
func TestSimpleCalculate(t *testing.T) {
	fmt.Println(simpleCalculate("6-(2+3)"))
}
func TestCalculate1(t *testing.T) {
	fmt.Println(calculate1(" 42"))
}
func TestStr2Bytes(t *testing.T) {
	var s = "123+"
	by := []byte(s)
	fmt.Println(by)
}
func TestStack(t *testing.T) {
	//TODO 这个pkg包中的 stack实现的有问题
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println("size:", s.Size())
	s.Pop()
}

func TestPancakeSort(T *testing.T) {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
}
