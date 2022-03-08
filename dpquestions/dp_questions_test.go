package dpquestions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5, 6, 10}, 20))
}

func TestLengthOfLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{1, 2, 3, 4, 3, 10}))
}

func TestLengthOfLISByBS(t *testing.T) {
	fmt.Println(lengthOfLISByBS([]int{1, 2, 3, 4, 3, 10}))
}

func TestMaxEnvelopes(t *testing.T) {
	ass := assert.New(t)
	ass.Equal(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}), 3, "should be equal 3")
	//ass.Equal(maxEnvelopes([][]int{{1, 2}, {2, 3}, {2, 10}, {1, 2}}), 3, "should be equals")
	// ass.Equal(maxEnvelopes([][]int{}), 0, "should be equal zero")
	//ass.Equal(maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}), 1, "should be equal 1")
	//fmt.Println(maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}))

}

func TestQuickSort(t *testing.T) {
	var arr = [][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}}

	QuickSort(arr)
	fmt.Println(arr)
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-3, 1, 3, -1, 2, -4, 2}))
	fmt.Println(maxSubArray([]int{1}))
}

func TestMaxSubArraySlim(t *testing.T) {
	fmt.Println(maxSubArraySlim([]int{-3, 1, 3, -1, 2, -4, 2}))
	fmt.Println(maxSubArraySlim([]int{1}))
}
