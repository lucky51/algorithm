package backtracking

var res [][]int = make([][]int, 0)

// permute 全排列问题
func permute(nums []int) [][]int {
	track := make([]int, 0)
	backtrack(nums, track)
	return res
}

// contains
func contains(s []int, value int) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == value {
			return true
		}
	}
	return false
}

func backtrack(nums []int, track []int) {
	if len(track) == len(nums) {
		tem := make([]int, len(nums))
		copy(tem, track)
		res = append(res, tem)
		return
	}
	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}
		// 做选择
		track = append(track, nums[i])
		// 进入下一层决策树
		backtrack(nums, track)
		// 取消选择
		track = track[:len(track)-1]
	}

}
