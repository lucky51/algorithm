package arrayq

// leetcode 560.和为 K 的子数组
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数
//输入：nums = [1,1,1], k = 2
//输出：2
func subarraySum(nums []int, k int) int {
	var res = 0
	var prefixArr = make([]int, len(nums)+1)
	prefixArr[0] = 0
	for i := 0; i < len(nums); i++ {
		prefixArr[i+1] = prefixArr[i] + nums[i]
	}
	for j := 1; j < len(prefixArr); j++ {
		for x := 0; x < j; x++ {
			if prefixArr[j]-prefixArr[x] == k {
				res++
			}
		}
	}
	return res
}
