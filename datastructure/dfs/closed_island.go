package dfs

// leetcode 1254. 统计封闭岛屿的数目 TODO: leetcode 还未提交
//二维矩阵 grid由 0（土地）和 1（水）组成。岛是由最大的4个方向连通的 0组成的群，封闭岛是一个完全 由1包围（左、上、右、下）的岛。
//请返回 封闭岛屿 的数目
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

func dfs(grid [][]int, r, c int, result bool) bool {
	grid[r][c] = 2
	if r-1 >= 0 && grid[r-1][c] == 0 {
		if isBorder(grid, r-1, c) {
			result = false
		} else {
			// result放到后边，避免提前短路
			result = dfs(grid, r-1, c, result) && result
		}
	}
	if r+1 < len(grid) && grid[r+1][c] == 0 {
		if isBorder(grid, r+1, c) {
			result = false
		} else {
			result = dfs(grid, r+1, c, result) && result
		}

	}
	if c-1 >= 0 && grid[r][c-1] == 0 {
		if isBorder(grid, r, c-1) {
			result = false
		} else {
			result = dfs(grid, r, c-1, result) && result
		}

	}
	if c+1 < len(grid[r]) && grid[r][c+1] == 0 {
		if isBorder(grid, r, c+1) {
			result = false
		} else {
			result = dfs(grid, r, c+1, result) && result
		}
	}
	return result
}

func isBorder(grid [][]int, r, c int) bool {
	return r == 0 || c == 0 || r == len(grid)-1 || c == len(grid[r])-1
}
