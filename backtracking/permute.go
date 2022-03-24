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

var resnew [][]int
var tracksnew []int
var used map[int]bool

func permuteNew(nums []int) [][]int {
	resnew = make([][]int, 0)
	tracksnew = make([]int, 0)
	used = make(map[int]bool, 0)
	backtrackNew(nums)
	return res
}
func backtrackNew(nums []int) {
	if len(nums) == len(tracksnew) {
		// 到达子节点
		tracksnewCopy := make([]int, len(tracksnew))
		copy(tracksnewCopy, tracksnew)
		res = append(res, tracksnewCopy)
		return
	}
	for i := 0; i < len(nums); i++ {
		if v, ok := used[i]; ok && v {
			continue
		}
		//select
		tracksnew = append(tracksnew, nums[i])
		used[i] = true
		backtrackNew(nums)
		used[i] = false
		tracksnew = tracksnew[:len(tracksnew)-1]
	}
}
