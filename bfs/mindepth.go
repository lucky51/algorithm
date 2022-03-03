package bfs

import (
	"github.com/lucky51/pkg/queue"
	"github.com/lucky51/pkg/tree"
)

// minDepth 求最小深度
func minDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	q := queue.NewQueue[tree.TreeNode]()
	q.Offer(root)
	depth := 1
	for !q.IsEmpty() {
		len := int(q.Size())
		for i := 0; i < len; i++ {
			cur := q.Poll()
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				q.Offer(cur.Left)
			}
			if cur.Right != nil {
				q.Offer(cur.Right)
			}
		}
		depth += 1
	}
	return depth
}
