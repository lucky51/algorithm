/*
有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。
这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。
如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
求所能获得硬币的最大数量。
*/
package backtracking

import "math"

var scoreRes = math.MinInt
var scoreMemo = make(map[int]int)

// leecode 312 戳气球
func maxCoins(nums []int) int {
	backtrackCoins(nums, 0)
	return scoreRes
}

func backtrackCoins(nums []int, score int) {
	if len(nums) == 0 {
		scoreRes = int(math.Max(float64(score), float64(scoreRes)))
		return
	}
	if _, ok := scoreMemo[score]; ok {
		return
	}
	for i := 0; i < len(nums); i++ {

		l, r := 0, 0
		if i == 0 {
			l = 1
		} else {
			l = nums[i-1]
		}
		if i == len(nums)-1 {
			r = 1
		} else {
			r = nums[i+1]
		}
		s := l * nums[i] * r
		temp := make([]int, len(nums))
		copy(temp, nums)
		backtrackCoins(append(temp[:i], temp[i+1:]...), s+score)
	}
}
