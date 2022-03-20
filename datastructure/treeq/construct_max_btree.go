package treeq

import (
	"fmt"

	"github.com/lucky51/pkg/tree"
)

// leetcode 654
// constructMaximumBinaryTree 最大二叉树
func constructMaximumBinaryTree(nums []int) *tree.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &tree.TreeNode{}
	maxnum, maxindex := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxnum {
			maxnum = nums[i]
			maxindex = i
		}
	}
	root.Val = maxnum
	root.Left = constructMaximumBinaryTree(nums[:maxindex])
	root.Right = constructMaximumBinaryTree(nums[maxindex+1:])
	return root
}

// leetcode 105
// 从前序与中序遍历序列构造二叉树
func buildTree(preorder, inorder []int) *tree.TreeNode {
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

// build 按区间构建
func build(preorder, inorder []int, prestart, preend, instart, inend int) *tree.TreeNode {
	if prestart > preend {
		return nil
	}
	rootval := preorder[prestart]
	index := 0
	for index = instart; index <= inend; index++ {
		if inorder[index] == rootval {
			break
		}
	}
	root := &tree.TreeNode{
		Val: rootval,
	}
	leftSize := index - instart
	root.Left = build(preorder, inorder, prestart+1, leftSize+prestart, instart, index-1)
	root.Right = build(preorder, inorder, leftSize+prestart+1, preend, index+1, inend)
	return root
}

// 从中序、后序遍历顺序构造
func buildTreePost(inorder, postorder []int) *tree.TreeNode {
	return buildPost(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func buildPost(inorder, postorder []int, instart, inend, poststart, postend int) *tree.TreeNode {
	if poststart > postend {
		return nil
	}
	fmt.Println(postend)
	rootVal := postorder[postend]
	root := &tree.TreeNode{
		Val: rootVal,
	}
	index := 0
	for index = instart; index <= inend; index++ {
		if inorder[index] == rootVal {
			break
		}
	}
	rightSize := inend - index
	root.Left = buildPost(inorder, postorder, instart, index-1, poststart, postend-rightSize-1)
	root.Right = buildPost(inorder, postorder, index+1, inend, postend-rightSize, postend-1)
	return root
}
