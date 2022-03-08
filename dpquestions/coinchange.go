package dpquestions

import "math"

// coinChange 自底向上的dp方法计算找零钱需要最少的硬币数
func coinChange(coins []int, amount int) int {
	// 初始化dp数组,求得amount对应的值，数组的最大长度应为 amount +1
	dp := make([]int, amount+1)
	// 循环初始化dp数组的初始值为 amount +1 ，因为amount金额对应最大的数量为amount ，amount + 1 则可代替用于比较中不可能到达的值
	for cindex := range dp {
		dp[cindex] = amount + 1
	}
	// base case
	dp[0] = 0
	// 遍历dp ，根据子问题赋值
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			// 求得对应钱数的子问题+1
			if i-coin < 0 {
				// 子问题无解，跳过
				continue
			}
			dp[i] = int(math.Min(float64(dp[i]), 1+float64(dp[i-coin])))
		}
	}
	// 如果amount的值等于初始值amount +1 ，则无解
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
