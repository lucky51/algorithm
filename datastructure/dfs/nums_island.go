package dfs

func numsIslandDFS(grid [][]byte, r, c int) {
	grid[r][c] = 0
	gSize := len(grid)
	cSize := len(grid[0])
	if r-1 >= 0 && grid[r-1][c] == 1 {
		numsIslandDFS(grid, r-1, c)
	}
	if r+1 < gSize && grid[r+1][c] == 1 {
		numsIslandDFS(grid, r+1, c)
	}
	if c-1 >= 0 && grid[r][c-1] == 1 {
		numsIslandDFS(grid, r, c-1)
	}
	if c+1 < cSize && grid[r][c+1] == 1 {
		numsIslandDFS(grid, r, c+1)
	}
}

// numIslands leetcode 200. 岛屿数量
//给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//此外，你可以假设该网格的四条边均被水包围。
func numIslands(grid [][]byte) int {
	nums := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				nums++
				numsIslandDFS(grid, i, j)
			}
		}
	}
	return nums
}
