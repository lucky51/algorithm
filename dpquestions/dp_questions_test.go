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

func TestLongestCommonSubsequence(t *testing.T) {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequenceSlim("abcde", "ace"))
	fmt.Println(longestCommonSubsequenceSlim("abc", "abc"))
}
func TestMinDistance(t *testing.T) {
	fmt.Println(minDistance("sea", "eat"))
}

func TestMinimumDeleteSum(t *testing.T) {
	fmt.Println(minimumDeleteSum("delete", "leet"))
}
func TestAscII(t *testing.T) {
	var s = "."
	fmt.Println(s[0] == '.')
	fmt.Println(10 << 4) //1010
}
func TestEditMinDistance(t *testing.T) {
	fmt.Println(editminDistance("dinitrophenylhydrazine", "benzalphenylhydrazone"))
	fmt.Println(editMinDistancebyDPTable("dinitrophenylhydrazine", "benzalphenylhydrazone"))
}

func TestLongestPalindRomeSubSeq(t *testing.T) {
	/*
		[1 1 1 3 5]
		[0 1 1 3 3]
		[0 0 1 1 1]
		[0 0 0 1 1]
		[0 0 0 0 1]

	*/
	//	fmt.Println(longestPalindRomeSubseq("bbbab"))
	//fmt.Println(longestPalindRomeSubseqSlime("bbbab"))
	fmt.Println(minInsertions("zzazz"))
	fmt.Println(minInsertionsSlime("zzazz"))
}

func TestIsMatch(t *testing.T) {
	fmt.Println(isMatch("zaaab", ".a*b"))
}

func TestMaxA(t *testing.T) {
	fmt.Println(maxA(10, 0, 0))
	fmt.Println(maxANew(10))
}

func TestSuperEggDrop(t *testing.T) {
	//fmt.Println(superEggDrop(2, 6))
	fmt.Println(superEggDropOptimize(2, 6))
	fmt.Println(superEggDropForDpTable(2, 6))
}

func TestDpMaxCoins(t *testing.T) {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}
func TestKnapsack(t *testing.T) {
	fmt.Println(knapsack(4, 3, []int{2, 1, 3}, []int{4, 2, 3}))
}
func TestCanPartition(t *testing.T) {
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}

func TestChange(t *testing.T) {
	fmt.Println(change(5, []int{1, 2, 5}))
}
func TestRob(t *testing.T) {
	fmt.Println(rob([]int{1, 2, 5}))
	fmt.Println(robdpt([]int{1, 2, 5}))
}
