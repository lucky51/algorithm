package treeq

import "github.com/lucky51/pkg/tree"

// lowestCommonAncestor 最近公共祖先
func lowestCommonAncestor(root, p, q *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if root == q || root == p {
		return root
	}
	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	// left 和right都不为nil ，则一定是 pq,又由于为后序遍历所以一定是最小的公共祖先
	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	} else {
		return left
	}
}
