package linkq

import (
	"fmt"
	"testing"

	"github.com/lucky51/pkg/list"
)

func TestIsPalindrome(t *testing.T) {
	head := &list.Node[int]{
		Data: 1,
		Next: &list.Node[int]{
			Data: 2,
			Next: &list.Node[int]{
				Data: 2,
				Next: &list.Node[int]{
					Data: 1,
				},
			},
		},
	}
	fmt.Println(isPalindrome(head))

}
