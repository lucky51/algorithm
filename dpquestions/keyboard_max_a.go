package dpquestions

import (
	"fmt"
	"math"
)

var maxAMemo = make(map[string]int)

//maxA   n = 剩余的操作数, a_nums = 屏幕上显示A的数量,copy = 缓冲区中A的数量
//求屏幕上可以显示A数量的最大值 ，此种算法效率低 n^3 ,加上备忘录模式效率也很低下
func maxA(n, anums, cp int) int {
	//base case
	if n <= 0 {
		return anums
	}
	key := fmt.Sprintf("%d%d%d", n, anums, cp)
	if _, ok := maxAMemo[key]; ok {
		return maxAMemo[key]
	}
	maxAMemo[key] = int(
		math.Max(
			math.Max(
				float64(maxA(n-1, anums+1, cp)),
				float64(maxA(n-1, anums+cp, cp)),
			),
			float64(maxA(n-2, anums, anums)),
		),
	)
	return maxAMemo[key]
}

// A,A,A,<C,A><C,V>,<C,P>,<C,P>,<C,P>....
func maxANew(N int) int {
	dp := make([]int, N+1)
	//base case
	dp[0] = 0
	for i := 1; i <= N; i++ {
		dp[i] = dp[i-1] + 1
		for j := 2; j < i; j++ {
			dp[i] = int(math.Max(float64(dp[i]), float64((i-j+1)*dp[j-2])))
		}
	}
	return dp[N]
}
