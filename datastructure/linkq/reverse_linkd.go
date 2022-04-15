package linkq

import "github.com/lucky51/pkg/list"

// reverseLinked 全链表反转 leetcode 206
func reverseLinked(head *list.Node[int]) *list.Node[int] {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseLinked(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

var successor *list.Node[int]

// reverseLinkedN 反转前N个元素的链表
func reverseLinkedN(head *list.Node[int], n int) *list.Node[int] {
	//base case
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseLinkedN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}

// reverseLinkedBetween 按区间反转链表
func reverseLinkedBetween(head *list.Node[int], m, n int) *list.Node[int] {
	//base case
	if m == 1 {
		return reverseLinkedN(head, n)
	}
	head.Next = reverseLinkedBetween(head.Next, m-1, n-1)
	return head
}

// reverse 反转链表
func reverse(a, b *list.Node[int]) *list.Node[int] {
	var pre, cur, next *list.Node[int] = nil, a, a
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// reverseKGroup 根据k个一组进行反转链表
func reverseKGroup(head *list.Node[int], k int) *list.Node[int] {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)

	a.Next = reverseKGroup(b, k)
	return newHead
}
