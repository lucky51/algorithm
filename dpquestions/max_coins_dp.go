package dpquestions

import (
	"fmt"
	"math"
)

// 戳气球问题，获取最大金币数
// maxCoins reference backtracking/max_coins.go
func maxCoins(nums []int) int {
	n := len(nums)
	points := make([]int, n+2)
	points[0] = 1

	for z := 0; z < n; z++ {
		points[z+1] = nums[z]
	}
	points[n+1] = 1
	fmt.Println(points)
	// dp[i][j] 表示 i j之间的金币数 （i,k)
	dp := make([][]int, n+2)
	for d := range dp {
		dp[d] = make([]int, n+2)
	}
	for i := n; i >= 0; i-- {
		for j := i + 2; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = int(
					math.Max(
						float64(dp[i][j]),
						float64(dp[i][k]+dp[k][j]+points[k]*points[i]*points[j]),
					),
				)
			}
		}
	}
	return dp[0][n+1]
}
