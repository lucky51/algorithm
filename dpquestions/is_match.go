package dpquestions

import "fmt"

func isMatch(s, p string) bool {
	memo := make(map[string]bool)
	return dp1(s, p, 0, 0, memo)
}

func dp1(s, p string, i, j int, memo map[string]bool) bool {
	key := fmt.Sprintf("%d%d", i, j)
	//base case
	if i == len(s) {
		if (len(p)-j)%2 == 1 {
			return false
		}
		for ; j+1 < len(p); j += 2 {
			if p[j+1] != '*' {
				return false
			}
		}
		return true
	}
	if j == len(p) {
		return i == len(s)
	}
	if _, ok := memo[key]; ok {
		return memo[key]
	}
	res := false
	//首字母匹配
	if s[i] == p[j] || p[j] == '.' {
		// 含有通配符*
		if j+1 < len(p) && p[j+1] == '*' {
			res = dp1(s, p, i, j+2, memo) || dp1(s, p, i+1, j, memo)
		} else {
			res = dp1(s, p, i+1, j+1, memo)
		}
	} else {
		if j+1 < len(p) && p[j+1] == '*' {
			res = dp1(s, p, i, j+2, memo)
		}
	}
	memo[key] = res
	return res
}
