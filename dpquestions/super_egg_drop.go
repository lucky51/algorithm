package dpquestions

import (
	"fmt"
	"math"
	"strconv"
)

// 鸡蛋掉落 leetcode 887
// superEggDrop 高空扔鸡蛋问题，k:鸡蛋数，n:楼层  ,leetcode 提交超时
func superEggDrop(k, n int) int {
	return eggdp(k, n, make(map[string]int))
}
func eggdp(k, n int, memo map[string]int) int {
	//base case
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}
	key := fmt.Sprintf("%d%d", k, n)
	res := math.MaxInt
	if res, ok := memo[key]; ok {
		return res
	}
	for i := 1; i < n+1; i++ {
		res = int(
			math.Min(
				float64(
					math.Max(
						float64(eggdp(k-1, i-1, memo)), //碎了
						float64(eggdp(k, n-i, memo)),   //没碎
					)+1,
				),
				float64(res),
			),
		)
	}
	memo[key] = res
	return res
}

// superEggDropOptimize 优化 高空扔鸡蛋问题
func superEggDropOptimize(k, n int) int {
	return eggdpOptimize(k, n, make(map[string]int))
}

func eggdpOptimize(k, n int, memo map[string]int) int {
	//base case
	if k == 1 {
		return n
	}
	if n == 0 {
		return n
	}
	key := fmt.Sprintf("%d%d", k, n)
	strconv.Itoa(1)
	if res, ok := memo[key]; ok {
		return res
	}
	res := math.MaxInt
	// 线性查找改为二分查找
	lo, hi := 1, n
	for lo <= hi {
		mid := (lo + hi) / 2
		hasBroke := eggdpOptimize(k-1, mid-1, memo)
		notBroke := eggdpOptimize(k, n-mid, memo)
		if hasBroke > notBroke {
			hi = mid - 1
			res = int(math.Min(float64(hasBroke+1), float64(res)))
		} else {
			lo = mid + 1
			res = int(math.Min(float64(notBroke+1), float64(res)))
		}
	}
	memo[key] = res
	return res
}

func superEggDropForDpTable(k, n int) int {
	// dp[k][m] = n  k=鸡蛋数，m=最大的尝试次数, n =楼层
	dp := make([][]int, k+1)
	for index := range dp {
		dp[index] = make([]int, n+1)
	}
	m := 0
	for dp[k][m] < n {
		m++
		for j := 1; j <= k; j++ {
			dp[j][m] = dp[j][m-1] + dp[j-1][m-1] + 1
		}

	}
	return m
}
