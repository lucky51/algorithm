package backtracking

import "github.com/lucky51/pkg/ext"

var csRes [][]int
var csTracks []int
var csTrackSums = 0

// combinationSum leetcode 39
//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target
//找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合
// 并以列表形式返回。你可以按 任意顺序 返回这些组合
// candidates 中的 同一个 数字可以 无限制重复被选取
// 如果至少一个数字的被选数量不同，则两种组合是不同的
func combinationSum(candidates []int, target int) [][]int {
	csRes = make([][]int, 0)
	csTracks = make([]int, 0)
	csBackTrack(candidates, 0, target)
	return csRes
}

func csBackTrack(candidates []int, start, target int) {
	if csTrackSums == target {
		csRes = append(csRes, ext.Clone(csTracks))
	}
	if csTrackSums > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		//select
		csTrackSums += candidates[i]
		csTracks = append(csTracks, candidates[i])
		csBackTrack(candidates, i, target)
		//unselect
		csTrackSums -= candidates[i]
		csTracks = csTracks[:len(csTracks)-1]
	}
}
