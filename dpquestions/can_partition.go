package dpquestions

// canPartition leetcode:416 分割等和子集
// 可以将此问题转换为经典的背包问题
func canPartition(nums []int) bool {
	sum := 0
	n := len(nums)
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	w := sum / 2

	dp := make([][]bool, n+1)
	for d := range dp {
		dp[d] = make([]bool, w+1)
		for di := range dp[d] {
			if di == 0 {
				dp[d][di] = true
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				// 装入 或者不装入背包的情况
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][w]

}
