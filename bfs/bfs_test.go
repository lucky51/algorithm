package bfs

import (
	"fmt"
	"testing"
)

func TestSlidingPuzzle(t *testing.T) {
	board := [][]byte{
		{1, 2, 3},
		{4, 0, 5},
	}
	fmt.Println(slidingPuzzle(board))

}
