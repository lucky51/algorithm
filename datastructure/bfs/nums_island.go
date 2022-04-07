package bfs

import "github.com/lucky51/pkg/queue"

// numIslands bfs实现岛屿数量问题 leetcode 200
// dfs解法 reference dfs.numIslands
func numIslands(grid [][]byte) int {
	nums := 0
	q := queue.NewQueue[[]int]()
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				nums++
				q.Offer([]int{i, j})
				for !q.IsEmpty() {
					curr := q.Poll()
					r, c := curr[0], curr[1]
					grid[r][c] = 0
					if r-1 >= 0 && grid[r-1][c] == 1 {
						q.Offer([]int{r - 1, c})
					}
					if r+1 < len(grid) && grid[r+1][c] == 1 {
						q.Offer([]int{r + 1, c})
					}
					if c-1 >= 0 && grid[r][c-1] == 1 {
						q.Offer([]int{r, c - 1})
					}
					if c+1 < len(grid[r]) && grid[r][c+1] == 1 {
						q.Offer([]int{r, c + 1})
					}
				}
			}
		}
	}
	return nums
}
