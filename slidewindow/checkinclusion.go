/*
* 字符串排列问题
* 输入两个字符串 S和T ，请你用算法判断S是否包含T的排列，也就是要判断S中是否存在一个子串T的一种全排列
* 比如输入 S = "helloworld" T = "oow" 算法返回 true ,因为S包含一个子串 "owo" 是 T 的排列
* 这里开始误解了 子串全排列的意思，以为是三个字母包含就符合题意，其实是错的 ！！！！！ 注意
 */
package slidewindow

import (
	"fmt"
	"math"
)

// minWindow 最小覆盖子串问题，给你两个字符串 S 和T ，请你在S中找到包含T总全部字母的最短子串
// 如果 S中没有这样的子串，则算法返回空串，如果存在这样一个子串，则可以认为答案是惟一的。
func minWindow(s, t string) string {
	// map的value要用int ，leetcode最后的测试数据会很大超过byte
	//need, window := make(map[byte]byte), make(map[byte]byte)
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	valid := 0
	start, ln := 0, math.MaxInt
	left, right := 0, 0
	for right < len(s) {
		if _, ok := need[s[right]]; ok {
			window[s[right]]++
			if need[s[right]] == window[s[right]] {
				valid++
			}
		}
		right++
		for valid == len(need) {
			if right-left < ln {
				ln = right - left
				start = left
			}
			if _, ok := need[s[left]]; ok {
				if _, exists := window[s[left]]; exists {
					if window[s[left]] == need[s[left]] {
						valid--
					}
					window[s[left]]--
					if window[s[left]] < 1 {
						delete(window, s[left])
					}
				}
			}
			left++
		}
	}
	if ln == math.MaxInt {
		return ""
	} else {
		return s[start : start+ln]
	}
}

func minWindow1(s, t string) string {
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
			if _, ok := need[s[left]]; ok {
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

// checkInclusion 判断s是否含有子串s ，问题描述如上
func checkInclusion(s1, s2 string) bool {
	need, window := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < len(s2); i++ {
		need[s2[i]]++
	}
	valid := 0
	left, right := 0, 0
	for right < len(s1) {
		if _, ok := need[s1[right]]; ok {
			window[s1[right]]++
			if need[s1[right]] == window[s1[right]] {
				valid++
			}
		}
		right++
		for right-left >= len(s2) {
			if valid == len(need) {
				fmt.Println(need, window, valid)
				return true
			}
			if _, ok := need[s1[left]]; ok {
				if _, exists := window[s1[left]]; exists {
					// 需要先判断在移除
					if need[s1[left]] == window[s1[left]] {
						valid--
					}
					window[s1[left]]--
					if window[s1[left]] < 1 {
						// 需要删掉元素，要不然影响
						delete(window, s1[left])
					}

				}

			}
			left++
		}
	}
	return false
}

// findAnagrams 找到字符串中所有字母异位词
func findAnagrams(s, p string) []int {
	if len(p) > len(s) {
		return nil
	}
	need, window := make(map[byte]int), make(map[byte]int)
	result := make([]int, 0)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s) {
		if _, ok := need[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				valid++
			}
		}
		right++
		for right-left >= len(p) {
			if valid == len(need) {
				result = append(result, left)
				//valid = 0
			}

			if _, ok := need[s[left]]; ok {
				if _, exists := window[s[left]]; exists {
					if window[s[left]] == need[s[left]] {
						valid--
					}
					window[s[left]]--
					if window[s[left]] < 1 {
						delete(window, s[left])
					}
				}
			}
			left++
		}
	}

	return result
}

// lengthOfLongestSubstring 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	window := make(map[byte]int)
	result, left, right := 0, 0, 0
	for right < len(s) {
		window[s[right]]++
		for window[s[right]] > 1 {
			// 收缩左边
			window[s[left]]--
			left++
		}
		if right-left > result {
			result = right - left
		}
		right++
	}
	return result + 1
}
