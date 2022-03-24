package backtracking

var combineres [][]int
var combinetracks []int

// backtrackingCombine leecode77 组合
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合
func backtrackingCombine(n, k int) [][]int {
	combineres = make([][]int, 0)
	combinetracks = make([]int, 0)
	combinebacktracking(1, n, k)
	return res
}

func combinebacktracking(start, n, k int) {
	//base case
	if len(combinetracks) == k {
		newtracks := make([]int, len(combinetracks))
		copy(newtracks, combinetracks)
		res = append(res, newtracks)
		return
	}
	for i := start; i <= n; i++ {
		//select
		combinetracks = append(combinetracks, i)
		combinebacktracking(i+1, n, k)
		// unselect
		combinetracks = combinetracks[:len(combinetracks)-1]
	}
}
