/*
* leetcode 354. 俄罗斯套娃信封问题
* 给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
*当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
*开始用到了冒泡排序发现，提交记录超时，又去学习快排，快排对于一个数组排序思考了很久!!! 需要根据第一个元素升序，第二个元素降序。
 */
package dpquestions

// maxEnvelopes 求解信封的最大嵌套数
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	//	res := 0
	// 按照宽度升序，如果宽度相同则按照高度降序
	// for k := 0; k < len(envelopes); k++ {
	// 	for l := 0; l < len(envelopes)-k-1; l++ {
	// 		if envelopes[l][0] > envelopes[l+1][0] {
	// 			envelopes[l], envelopes[l+1] = envelopes[l+1], envelopes[l]
	// 		} else if envelopes[l][0] == envelopes[l+1][0] {
	// 			if envelopes[l][1] < envelopes[l+1][1] {
	// 				envelopes[l], envelopes[l+1] = envelopes[l+1], envelopes[l]
	// 			}
	// 		}
	// 	}
	// }
	QuickSort(envelopes)
	// 使用二分查找方式
	top := make([]int, len(envelopes))
	piles := 0
	for z := 0; z < len(envelopes); z++ {
		poker := envelopes[z]
		left, right := 0, piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] >= poker[1] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left == piles {
			piles++
		}
		top[left] = poker[1]
	}

	return piles
	// 对排序好的数组元素中的（长），转化为获取最大递增子序列问题
	// dp := make([]int, len(envelopes))
	// for index := range dp {
	// 	dp[index] = 1
	// }
	// for i := 0; i < len(envelopes); i++ {
	// 	for j := 0; j < i; j++ {
	// 		if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
	// 			if dp[j]+1 > dp[i] {
	// 				dp[i] = dp[j] + 1
	// 			}
	// 		}
	// 	}
	// }
	// for _, n := range dp {
	// 	if n > res {
	// 		res = n
	// 	}
	// }
	// return res

}

func QuickSort(intList [][]int) {
	// 如果长度小于等于1就直接结束
	if len(intList) <= 1 {
		return
	}
	// 1. 将第一个值选定为基准值
	flag := intList[0]
	// 定义左右指针
	left, right := 0, len(intList)-1

	for i := 1; i <= right; {
		if intList[i][0] > flag[0] {
			intList[i], intList[right] = intList[right], intList[i]
			right--
		} else if intList[i][0] == flag[0] && intList[i][1] < flag[1] {
			intList[i], intList[right] = intList[right], intList[i]
			right--
		} else {
			intList[i], intList[left] = intList[left], intList[i]
			i++
			left++
		}
	}
	// 递归
	QuickSort(intList[:left])
	QuickSort(intList[left+1:])
}
