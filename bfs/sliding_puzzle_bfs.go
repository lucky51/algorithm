package bfs

import (
	"fmt"
	"github.com/lucky51/pkg/queue"
)

/*
这个位置在棋盘上是固定的，不随着内容变化
0 ,1 ,2
3, 4, 5
*/
var neighbors = [][]int{
	{1, 3},
	{0, 4, 2},
	{1, 5},
	{0, 4},
	{3, 1, 5},
	{4, 2},
}

// slidingPuzzle 滑动谜题 leecode 773
// 在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用
// 0来表示。一次 移动 定义为选择0与一个相邻的数字（上下左右）进行交换.
func slidingPuzzle(board [][]byte) int {
	target, start := []byte{1, 2, 3, 4, 5, 0}, make([]byte, 6)
	sIndex := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			start[sIndex] = board[i][j]
			sIndex++
		}
	}
	visited := make(map[string]bool)
	q := queue.NewQueue[[]byte]()
	q.Offer(start)
	step := 0
	for q.Size() > 0 {
		size := q.Size()
		//这个循环不能少
		for k := 0; k < int(size); k++ {
			cur := q.Poll()
			if isEqual(cur, target) {
				return step
			}
			zi := 0
			for zi = range cur {
				if cur[zi] == 0 {
					break
				}
			}
			for i := 0; i < len(neighbors[zi]); i++ {
				newCur := swapZero(cur, zi, neighbors[zi][i])
				var key = bytes2Str(newCur)
				if _, ok := visited[key]; ok {
					continue
				}
				q.Offer(newCur)
				visited[key] = true
			}
		}

		step++
	}
	return -1
}
func bytes2Str(b []byte) string {
	s := ""
	for i := 0; i < len(b); i++ {
		s = fmt.Sprintf("%s%d", s, b[i])
	}
	return s
}
func isEqual(item1, item2 []byte) bool {
	for i := 0; i < len(item1); i++ {
		if item1[i] != item2[i] {
			return false
		}
	}
	return true
}
func swapZero(cur []byte, i, j int) []byte {
	newCur := make([]byte, len(cur))
	copy(newCur, cur)
	newCur[i], newCur[j] = newCur[j], newCur[i]
	return newCur
}
