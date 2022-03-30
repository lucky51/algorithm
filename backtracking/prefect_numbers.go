package backtracking

import (
	"fmt"
	"strconv"
)

var grResStr string
var grRes int

// getPerfectNumber 获取完美数
//给出一个数，只包含1， 2， 3的数字称为完美数，如果当前数字不符合完美数的要求，那么返回可以构造的最大的不超过当前数字的完美数。
// 输入 213 => 返回 213
func getPerfectNumber(num int) int {
	grBackTrack(num, []int{1, 2, 3})
	return grRes
}
func grBackTrack(target int, nums []int) {
	//base case
	grRes1, _ := strconv.Atoi(grResStr)
	if grRes1 > target {
		return
	} else {
		if grRes1 > grRes {
			grRes = grRes1
		}
	}
	for i := 0; i < len(nums); i++ {
		//select
		grResStr = fmt.Sprintf("%s%d", grResStr, nums[i])
		grBackTrack(target, nums)
		//unselect
		grResStr = grResStr[:len(grResStr)-1]
	}
}
