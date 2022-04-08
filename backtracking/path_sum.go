package backtracking

import (
	"github.com/lucky51/pkg/ext"
	"github.com/lucky51/pkg/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var resTemp []int
var pathSumRes [][]int

// pathSum 113. 路径总和 II
// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
// 叶子节点 是指没有子节点的节点。
func pathSum(root *tree.TreeNode, targetSum int) [][]int {
	pathSumRes = make([][]int, 0)
	resTemp = make([]int, 0)
	backtrack1(root, targetSum)
	return pathSumRes
}

func backtrack1(root *tree.TreeNode, targetSum int) {
	if root == nil {
		return
	}
	resTemp = append(resTemp, root.Val)
	targetSum = targetSum - root.Val
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		pathSumRes = append(pathSumRes, ext.Clone(resTemp))
	} else {
		backtrack1(root.Left, targetSum)
		backtrack1(root.Right, targetSum)
	}
	targetSum = targetSum + root.Val
	resTemp = resTemp[:len(resTemp)-1]
}
