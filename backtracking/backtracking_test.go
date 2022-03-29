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

func TestSubSets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func TestSlicePointerAddress(t *testing.T) {
	a := make([]int, 0, 4)
	a = append(a, 1)
	b := a
	b = append(b, 66)
	fmt.Println(a[1:])
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
}

func TestSubsetsWithDup(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
func TestCombinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
func TestRune(t *testing.T) {
	fmt.Printf("%d", byte('.'))
	fmt.Println(46 == '.')
}
func TestSudoKu(t *testing.T) {
	b := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(b)
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			fmt.Printf("%s,", string(rune(b[i][j])))
		}
		fmt.Println("")
	}
}
func TestGenerateParenthesis(t *testing.T) {
	res := generateParenthesis(3)
	fmt.Println(res)
}
