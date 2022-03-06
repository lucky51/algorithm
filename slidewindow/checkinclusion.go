/*
* 字符串排列问题
* 输入两个字符串 S和T ，请你用算法判断S是否包含T的排列，也就是要判断S中是否存在一个子串T的一种全排列
* 比如输入 S = "helloworld" T = "oow" 算法返回 true ,因为S包含一个子串 "owo" 是 T 的排列
* 这里开始误解了 子串全排列的意思，以为是三个字母包含就符合题意，其实是错的 ！！！！！ 注意
 */
package slidewindow

import (
	"math"
)

// checkInclusion 判断s是否含有子串s ，问题描述如上
func checkInclusion(s, t string) bool {
	need, window := make(map[byte]byte), make(map[byte]byte)
	valid := 0
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	for right < len(s) {

		if _, ok := need[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				valid++
			}

		}
		right++
		for right-left >= len(t) {
			if valid == len(need) {
				return true
			}
			if _, ok := need[s[left]]; ok {
				window[s[left]]--
				if window[s[left]] == need[s[left]] {
					valid--
				}
			}
			left++
		}

	}
	return false
}

// minWindow 最小覆盖子串问题，给你两个字符串 S 和T ，请你在S中找到包含T总全部字母的最短子串
// 如果 S中没有这样的子串，则算法返回空串，如果存在这样一个子串，则可以认为答案是惟一的。
func minWindow(s, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	valid := 0
	left, right := 0, 0
	start, length := 0, math.MaxInt
	for right < len(s) {
		if _, ok := need[s[right]]; ok {
			window[s[right]]++
			if need[s[right]] == window[s[right]] {
				valid++
			}
		}
		right++
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			if n, ok := need[s[left]]; ok && n > 0 {
				if window[s[left]] == need[s[left]] {
					valid--
				}
				if window[s[left]] > 0 {
					window[s[left]]--
				}
			}
			left++

		}
	}
	if length == math.MaxInt {
		return ""
	} else {
		return s[start : start+length]
	}

}
