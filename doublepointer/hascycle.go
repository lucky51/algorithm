package doublepointer

import (
	slist "github.com/lucky51/pkg/list"
)

func hasCycle(head *slist.ListNode[int]) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
}
