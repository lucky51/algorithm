package dpquestions

import (
	pkgext "github.com/lucky51/pkg/ext"
	"github.com/lucky51/pkg/tree"
)

// rob 相邻的房子不能被同时取出，尽可能多的取钱
func rob(nums []int) int {
	return robdp(nums, 0)
}

func robdp(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	return pkgext.Max(
		robdp(nums, start+2)+nums[start],
		robdp(nums, start+1),
	)
}

// robdpt
func robdpt(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)
	//dp[i] = nums[i] + dp[i+2] ,dp[i+1]
	for i := n - 1; i >= 0; i-- {
		dp[i] = pkgext.Max(nums[i]+dp[i+2], dp[i+1])
	}
	return dp[0]
}

// robcycle 环形排列问题
func robcycle(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return pkgext.Max(
		robrange(nums, 1, n-1),
		robrange(nums, 0, n-2),
	)
}

// robrange
func robrange(nums []int, start, end int) int {
	// dpi ,dpi_1 ,dpi_2
	dpi, dpi_1, dpi_2 := 0, 0, 0
	for i := end; i >= start; i-- {
		dpi = pkgext.Max(nums[i]+dpi_2, dpi_1)
		dpi_2 = dpi_1
		dpi_1 = dpi
	}
	return dpi
}

// robTree 树形排列
func robTree(t *tree.Node[int]) int {
	if t == nil {
		return 0
	}
	// todo:添加memo备忘，以减少子问题重复计算
	do_it, not_do := t.Val, 0
	if t.Left != nil {
		do_it += robTree(t.Left.Left)
		do_it += robTree(t.Left.Right)

	}
	if t.Right != nil {
		do_it += robTree(t.Right.Left)
		do_it += robTree(t.Right.Right)
	}
	not_do = robTree(t.Left) + robTree(t.Right)
	res := pkgext.Max(do_it, not_do)
	return res
}
