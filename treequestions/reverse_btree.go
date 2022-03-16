package treequestions

import "github.com/lucky51/pkg/tree"

// reverseTree 翻转二叉树
func reverseTree(root *tree.TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	reverseTree(root.Right)
	reverseTree(root.Left)
}
