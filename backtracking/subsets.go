package backtracking

import (
	"container/list"
)

var subnums = make([][]int, 0)
var tracks *list.List = list.New()

// subsets leetcode 78题
// 输入无重复元素的数组，其中每个元素最多用一次，请你返回nums的所有子集
func subsets(nums []int) [][]int {
	subbacktracking(nums, 0)
	return subnums
}

func listToSlice(head *list.List) []int {
	res := make([]int, 0)
	if head == nil {
		return res
	}
	temp := head.Front()
	for temp != nil {
		res = append(res, temp.Value.(int))
		temp = temp.Next()
	}
	return res
}

func subbacktracking(nums []int, start int) {
	// newtracks := make([]int, len(tracks))
	// copy(newtracks, tracks)
	subnums = append(subnums, listToSlice(tracks))
	for i := start; i < len(nums); i++ {
		tracks.PushBack(nums[i])
		//tracks = append(tracks, nums[i])
		subbacktracking(nums, i+1)
		tracks.Remove(tracks.Back())
		//	tracks = tracks[:len(tracks)-1]
	}
}
