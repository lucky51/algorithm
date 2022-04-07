package dfs

import (
	"fmt"
	"github.com/lucky51/pkg/ext"
)

// maxAreaOfIsland see bfs.maxAreaOfIsland
func maxAreaOfIsland(grid [][]int) int {
	memo := make(map[string]bool)
	var sumMax = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if _, ok := memo[fmt.Sprintf("%d-%d", i, j)]; !ok && grid[i][j] == 1 {
				sumMax = ext.Max(sumMax, maxAreaOfIslandDFS(memo, grid, i, j, 0))
			}
		}
	}
	return sumMax
}

// maxAreaOfIslandDFS 深度优先遍历
func maxAreaOfIslandDFS(memo map[string]bool, grid [][]int, r, c, sum int) int {
	key := fmt.Sprintf("%d-%d", r, c)
	memo[key] = true
	sum++
	if _, ok := memo[fmt.Sprintf("%d-%d", r-1, c)]; !ok && r-1 >= 0 && grid[r-1][c] == 1 {
		sum = maxAreaOfIslandDFS(memo, grid, r-1, c, sum)
	}
	if _, ok := memo[fmt.Sprintf("%d-%d", r+1, c)]; !ok && r+1 < len(grid) && grid[r+1][c] == 1 {
		sum = maxAreaOfIslandDFS(memo, grid, r+1, c, sum)
	}
	if _, ok := memo[fmt.Sprintf("%d-%d", r, c-1)]; !ok && c-1 >= 0 && grid[r][c-1] == 1 {
		sum = maxAreaOfIslandDFS(memo, grid, r, c-1, sum)
	}
	if _, ok := memo[fmt.Sprintf("%d-%d", r, c+1)]; !ok && c+1 < len(grid[r]) && grid[r][c+1] == 1 {
		sum = maxAreaOfIslandDFS(memo, grid, r, c+1, sum)
	}
	return sum
}
