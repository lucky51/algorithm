package bfs

import (
	"fmt"
	"github.com/lucky51/pkg/queue"
)

var countSubIslandsMemo map[string]bool

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	countSubIslandsMemo = make(map[string]bool)
	grid1Dict := make(map[string]*[][]int)
	grid2Islands := make([][][]int, 0)
	q1, q2 := queue.NewQueue[[]int](), queue.NewQueue[[]int]()
	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[i]); j++ {
			if _, ok := countSubIslandsMemo[getMemoKey(1, i, j)]; !ok && grid1[i][j] == 1 {
				q1.Offer([]int{i, j})
				currIslands := make([][]int, 0)
				for !q1.IsEmpty() {
					curr := q1.Poll()
					currIslands = append(currIslands, []int{curr[0], curr[1]})
					grid1Dict[getMemoKey(1, curr[0], curr[1])] = &currIslands
					islandsQPush(grid1, q1, 1, curr[0], curr[1])
				}
			}
			if _, ok := countSubIslandsMemo[getMemoKey(2, i, j)]; !ok && grid2[i][j] == 1 {
				q2.Offer([]int{i, j})
				currIslands := make([][]int, 0)
				for !q2.IsEmpty() {
					curr := q2.Poll()
					currIslands = append(currIslands, []int{curr[0], curr[1]})
					islandsQPush(grid2, q2, 2, curr[0], curr[1])
				}
				grid2Islands = append(grid2Islands, currIslands)
			}
		}
	}
	for _, grid2Island := range grid2Islands {
		res++
		var prevExistsItems *[][]int
		for _, grid2IslandItem := range grid2Island {
			key := getMemoKey(1, grid2IslandItem[0], grid2IslandItem[1])
			if existsItems, exists := grid1Dict[key]; exists {
				if prevExistsItems == nil {
					prevExistsItems = existsItems
				}
				if prevExistsItems != existsItems {
					res--
					break
				}
				prevExistsItems = existsItems
			} else {
				res--
				break
			}
		}
	}
	return res
}

func getMemoKey(n, r, c int) string {
	return fmt.Sprintf("%d:%d-%d", n, r, c)
}
func islandsQPush(grid [][]int, q *queue.Queue[[]int], n, r, c int) {
	countSubIslandsMemo[getMemoKey(n, r, c)] = true

	if r-1 >= 0 && grid[r-1][c] == 1 {
		if _, ok := countSubIslandsMemo[getMemoKey(n, r-1, c)]; !ok {
			q.Offer([]int{r - 1, c})
		}
	}
	if r+1 < len(grid) && grid[r+1][c] == 1 {
		if _, ok := countSubIslandsMemo[getMemoKey(n, r+1, c)]; !ok {
			q.Offer([]int{r + 1, c})
		}
	}
	if c-1 >= 0 && grid[r][c-1] == 1 {
		if _, ok := countSubIslandsMemo[getMemoKey(n, r, c-1)]; !ok {
			q.Offer([]int{r, c - 1})
		}
	}
	if c+1 < len(grid[r]) && grid[r][c+1] == 1 {
		if _, ok := countSubIslandsMemo[getMemoKey(n, r, c+1)]; !ok {
			q.Offer([]int{r, c + 1})
		}
	}
}
