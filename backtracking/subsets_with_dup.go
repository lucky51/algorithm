package backtracking

import (
	"fmt"
	"sort"
)

var reswithdup [][]int
var trackswithdup []int

func subsetsWithDup(nums []int) [][]int {
	reswithdup = make([][]int, 0)
	trackswithdup = make([]int, 0)
	sort.Ints(nums)
	fmt.Println(nums)
	backtrackwithdup(nums, 0)
	return reswithdup
}

func backtrackwithdup(nums []int, start int) {
	trackswithdupcpy := make([]int, len(trackswithdup))
	copy(trackswithdupcpy, trackswithdup)

	reswithdup = append(reswithdup, trackswithdupcpy)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i-1] == nums[i] {
			continue
		}
		trackswithdup = append(trackswithdup, nums[i])
		backtrackwithdup(nums, i+1)
		trackswithdup = trackswithdup[:len(trackswithdup)-1]
	}
}

func qsort(nums []int) {
	povit := nums[0]
	lt, rh := 0, len(nums)-1
	for rh > lt && nums[rh] >= povit {
		rh--
	}
	nums[lt] = nums[rh]

}
