package backtracking

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	result := permute([]int{1, 2, 3})
	fmt.Println(result)
}

func TestSliceRemoveLast(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5}
	a = a[0 : len(a)-1]
	fmt.Println(a)
}

func TestNQueens(t *testing.T) {
	res := solveQueens(4)
	for i, _ := range res {
		for j, _ := range res[i] {
			fmt.Println(res[i][j])
		}
		fmt.Println("=============")
	}
}

func TestMaxScore(t *testing.T) {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}
