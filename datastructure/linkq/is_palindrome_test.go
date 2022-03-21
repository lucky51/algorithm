package linkq

import (
	"fmt"
	"testing"

	"github.com/lucky51/pkg/list"
)

func TestIsPalindrome(t *testing.T) {
	head := &list.ListNode[int]{
		Data: 1,
		Next: &list.ListNode[int]{
			Data: 2,
			Next: &list.ListNode[int]{
				Data: 2,
				Next: &list.ListNode[int]{
					Data: 1,
				},
			},
		},
	}
	fmt.Println(isPalindrome(head))

}
