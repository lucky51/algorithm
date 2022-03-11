package dpquestions

import "math"

// knapsack 背包问题,获取可以放置的最大价值
// 可装载重量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。其中第 i 个物品的重量为 wt[i]，价值为 val[i]
func knapsack(w, n int, wt, val []int) int {
	dp := make([][]int, n+1)
	for d := range dp {
		dp[d] = make([]int, w+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-wt[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-wt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[n][w]
}
