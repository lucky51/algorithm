package treeq

import "github.com/lucky51/pkg/tree"

// reverseTree 翻转二叉树
func reverseTree(root *tree.Node[int]) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	reverseTree(root.Right)
	reverseTree(root.Left)
}
