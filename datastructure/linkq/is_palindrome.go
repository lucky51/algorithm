package linkq

import (
	"github.com/lucky51/pkg/list"
)

var left *list.ListNode[int]

func isPalindrome(head *list.ListNode[int]) bool {
	left = head
	return traverse(head)
}

func traverse(right *list.ListNode[int]) bool {
	if right == nil {
		return true
	}
	res := traverse(right.Next)
	res = res && left.Data == right.Data
	left = left.Next
	return res
}

// TODO: isPalindromeFastSlow 快慢指针确定中点位置,然后翻转比较
