package bfs

import (
	"fmt"
	"github.com/lucky51/pkg/queue"
)

// maxAreaOfIsland leetCode  695. 岛屿的最大面积
// 给你一个大小为 m x n 的二进制矩阵 grid 。
//岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设grid 的四个边缘都被 0（代表水）包围着。
//岛屿的面积是岛上值为 1 的单元格的数目。
//计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
func maxAreaOfIsland(grid [][]int) int {
	maxNums := 0
	used := make(map[string]bool)
	q := queue.NewQueue[[]int]()
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				//q.Offer([]int{i, j})
				qPush(used, q, i, j)
				currSums := 0
				for !q.IsEmpty() {
					cur := q.Poll()
					currSums++
					r, c := cur[0], cur[1]
					if r-1 >= 0 && grid[r-1][c] == 1 {
						qPush(used, q, r-1, c)
						//if isExists(used, r-1, c) {
						//
						//	q.Offer([]int{r - 1, c})
						//}
					}
					if r+1 < len(grid) && grid[r+1][c] == 1 {
						qPush(used, q, r+1, c)
						//if isExists(used, r+1, c) {
						//
						//	q.Offer([]int{r + 1, c})
						//}
					}
					if c-1 >= 0 && grid[r][c-1] == 1 {
						qPush(used, q, r, c-1)
						//if !isExists(used, r, c-1) {
						//
						//	q.Offer([]int{r, c - 1})
						//}
					}
					if c+1 < len(grid[r]) && grid[r][c+1] == 1 {
						qPush(used, q, r, c+1)
						//if !isExists(used, r, c+1) {
						//
						//	q.Offer([]int{r, c + 1})
						//}
					}
				}
				if currSums > maxNums {
					maxNums = currSums
				}
			}
		}
	}
	return maxNums
}
func qPush(used map[string]bool, q *queue.Queue[[]int], r, c int) {
	key := fmt.Sprintf("%d-%d", r, c)
	if _, ok := used[key]; !ok {
		used[key] = true
		q.Offer([]int{r, c})
	}
}
func isExists(m map[string]bool, r, c int) bool {
	key := fmt.Sprintf("%d-%d", r, c)
	_, ok := m[key]
	return ok
}
