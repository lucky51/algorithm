package treequestions

import (
	"fmt"
	"testing"

	"github.com/lucky51/pkg/tree"
)

func TestReverseTree(t *testing.T) {
	r := tree.TreeNode{
		Val: 10,
		Left: &tree.TreeNode{
			Val: 1,
			Left: &tree.TreeNode{
				Val: 2,
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val: 5,
			},
			Right: &tree.TreeNode{
				Val: 6,
			},
		},
	}
	tree.PreOrderTree(&r)
	reverseTree(&r)
	fmt.Println("reversed")
	tree.PreOrderTree(&r)
}
