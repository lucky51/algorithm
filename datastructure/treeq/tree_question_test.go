package treeq

import (
	"fmt"
	"testing"

	"github.com/lucky51/pkg/tree"
)

func TestReverseTree(t *testing.T) {
	r := tree.Node[int]{
		Val: 10,
		Left: &tree.Node[int]{
			Val: 1,
			Left: &tree.Node[int]{
				Val: 2,
			},
			Right: &tree.Node[int]{
				Val: 3,
			},
		},
		Right: &tree.Node[int]{
			Val: 4,
			Left: &tree.Node[int]{
				Val: 5,
			},
			Right: &tree.Node[int]{
				Val: 6,
			},
		},
	}
	tree.PreOrderTree(&r)
	reverseTree(&r)
	fmt.Println("reversed")
	tree.PreOrderTree(&r)
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	r := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	tree.PreOrderTree(r)
}

func TestBuildTree(t *testing.T) {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	tree.PreOrderTree(root)
}

func TestBuildTreePost(t *testing.T) {
	root := buildTreePost([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	tree.PreOrderTree(root)
}

func TestNumTrees(t *testing.T) {
	fmt.Println(numTrees(3))
}
