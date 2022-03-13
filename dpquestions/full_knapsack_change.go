package dpquestions

// 零钱兑换II 给定不同面额的硬币coins 和一个总金额 amount，写一个函数来计算可以凑成总金额
// 的硬币组合数，假设每一种面额的硬币有无限个
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for d := range dp {
		dp[d] = make([]int, amount+1)
		dp[d][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[n][amount]
}
