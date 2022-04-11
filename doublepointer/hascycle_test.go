package doublepointer

import (
	"fmt"
	"testing"

	slist "github.com/lucky51/pkg/list"
)

func getIntgerPointer(i int) *int {
	return &i
}

func TestHasCycle(t *testing.T) {

	head := &slist.ListNode[int]{}
	current := head
	for _, item := range []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13} {
		current.Next = new(slist.ListNode[int])
		current.Data = item
		current = current.Next
	}

	fmt.Println("not has cycle :", hasCycle(head))
	current.Next = head.Next.Next
	fmt.Println("should have cycle:", hasCycle(head))
}

func TestSumOfTwoNumbers(t *testing.T) {
	fmt.Println(sumOfTwoNumbers([]int{1, 5, 8, 9, 10, 20, 25, 30, 38}, 35))
}

func TestReverseArray(t *testing.T) {
	fmt.Println(reverseArray([]int{0, 1, 5, 8, 9, 10, 20, 25, 30, 38}))
}
func TestNSumTarget(t *testing.T) {
	fmt.Println(nSumTarget([]int{1, 2, 4, 5, 7}, 3, 0, 7))
}

func TestTwoSumTarget(t *testing.T) {
	fmt.Println(towSumTarget([]int{1, 2, 3, 4}, 4))
}
func TestStrToNum(t *testing.T) {
	for j := '0'; j < '9'; j++ {
		fmt.Printf("%d,", j)
	}
	s := "456"
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		// 字符转数字即 该字符的ASCII码值减去0的ASCII所得到的的数字
		n = 10*n + int(c-'0')
	}
	fmt.Println(n)
}
