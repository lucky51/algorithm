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
