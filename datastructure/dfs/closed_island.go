package dfs

import (
	"fmt"
)

// TODO:未完成
var closedIsLandUsed = make(map[string]bool)

func closedIsland(grid [][]int) int {
	nums := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 0 {
				if dfs(grid, i, j, true) {
					nums++
				}
			}
		}
	}
	return nums
}

func getMemoKey(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}

func dfs(grid [][]int, r, c int, result bool) bool {
	closedIsLandUsed[fmt.Sprintf("%d-%d", r, c)] = true
	if r == 0 || c == 0 || r == len(grid)-1 || c == len(grid[r])-1 {
		if _, ok := closedIsLandUsed[fmt.Sprintf("%d-%d")]; ok {

		}
	} else {
		grid[r][c] = 1
	}

	if r-1 >= 0 && grid[r-1][c] == 0 {
		result = result && dfs(grid, r-1, c, result)
	}
	if r+1 < len(grid) && grid[r+1][c] == 0 {
		result = result && dfs(grid, r+1, c, result)
	}
	if c-1 >= 0 && grid[r][c-1] == 0 {
		result = result && dfs(grid, r, c-1, result)
	}
	if c+1 < len(grid[r]) && grid[r][c+1] == 0 {
		result = result && dfs(grid, r, c+1, result)
	}
	if r == 0 || c == 0 || r == len(grid)-1 || c == len(grid[r])-1 {
		if grid[r][c] == 0 {
			result = false
		}
	}
	return result
}
