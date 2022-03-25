package backtracking

import pkgExt "github.com/lucky51/pkg/ext"

var puTracks []int
var puRes [][]int
var puUsed []bool

func permuteUnique(nums []int) [][]int {
	puTracks = make([]int, 0)
	puRes = make([][]int, 0)
	puUsed = make([]bool, len(nums))
	puBackTrack(nums)
	return res
}

// puBackTrack LeetCode 47 全排列II
//  给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
func puBackTrack(nums []int) {
	if len(nums) == len(puTracks) {
		res = append(res, pkgExt.Clone(puTracks))
		return
	}
	for i := 0; i < len(nums); i++ {
		if puUsed[i] {
			continue
		}
		//为了保证相同元素的顺序选择，如果前边相邻相等元素没出现过则跳过
		//这个地方需要仔细理解，例如 1 2 '2 ''2 ,如果前边的2没有出现过，那我'2就不能出现，就是这块的逻辑
		if i > 0 && !puUsed[i-1] && nums[i] == nums[i-1] {
			continue
		}
		//select
		puUsed[i] = true
		puTracks = append(puTracks, nums[i])
		puBackTrack(nums)
		//unselect
		puUsed[i] = false
		puTracks = puTracks[:len(puTracks)-1]
	}

}
