package dpquestions

// lengthOfLIS 求解最长递增子序列
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	dp := make([]int, len(nums))
	//base case
	for index := range dp {
		dp[index] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// 判断第 i 个元素是否比前一个元素大
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	for _, n := range dp {
		if n > res {
			res = n
		}
	}
	return res
}

// lengthOfLISByBS 用二分查找解决
func lengthOfLISByBS(nums []int) int {
	// 用于记录牌堆顶部最小的元素
	top := make([]int, len(nums))
	// 牌堆数
	piles := 0
	for i := 0; i < len(nums); i++ {
		//抽取当前牌
		poker := nums[i]
		// 二分查找，定位左边界问题
		left, right := 0, piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}
		// 二分查找中直到left到达堆中最大值，就证明没有找到符合的条件，则新建堆
		if left == piles {
			piles++
		}
		top[left] = poker
	}

	return piles
}
