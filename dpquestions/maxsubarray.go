package dpquestions

import (
	"math"
)

// maxSubArray 获取最大子数组的和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := math.MinInt
	//base case
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for j := 1; j < len(nums); j++ {

		if nums[j] > dp[j-1]+nums[j] {

			dp[j] = nums[j]
		} else {
			dp[j] = dp[j-1] + nums[j]
		}
	}
	for _, item := range dp {
		if item > res {
			res = item
		}
	}
	return res
}

// maxSubArraySlim 获取最大子数组的和
func maxSubArraySlim(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp0, dp1, res := nums[0], 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp0+nums[i] {
			dp1 = nums[i]
		} else {
			dp1 = dp0 + nums[i]
		}
		if dp1 > res {
			res = dp1
		}

		dp0 = dp1
	}
	return res
}
