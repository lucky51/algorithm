package bfs

import (
	"fmt"
	"testing"

	"github.com/lucky51/pkg/tree"
)

func TestMinDepth(t *testing.T) {
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
				Right: &tree.TreeNode{
					Val: 99,
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 6,
				Left: &tree.TreeNode{
					Val: 7,
					Right: &tree.TreeNode{
						Val: 100,
					},
				},
			},
		},
	}

	fmt.Println(minDepth(root))
}

func TestOpenLock(t *testing.T) {
	fmt.Println(openLock([][]rune{{'9', '9', '9', '8'}}, "9999"))
}

func TestOpenLockDouble(t *testing.T) {
	m := map[string]struct{}{
		"9998": {},
	}
	fmt.Println(openLockDouble(m, "9999"))
}
