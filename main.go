package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5, 6, 10}, 20))
}

func coinChange(coins []int, amount int) int {
	// base case
	dp := make([]int, amount+1)
	for idx := range dp {
		dp[idx] = amount + 1
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}

			dp[i] = int(math.Min(float64(dp[i]), 1+float64(dp[i-coins[j]])))
		}
	}
	// 如果amount对应的值为初始值amount+1，无解
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
