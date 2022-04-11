package stackq

import (
	"fmt"
	"testing"

	"github.com/lucky51/pkg/stack"
)

func TestCalculate(t *testing.T) {
	fmt.Println(calculate("1+2+2"))
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
