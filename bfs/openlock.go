/*
* LeetCode 752 解开密码锁的最小次数
* 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字：'0','1','2','3','4','5','6','7','8','9'
* 每次旋转只能旋转一个拨轮的一位数字
* 锁的初始化数字为 '0000' ，一个代表四个拨轮的数字的字符串。
* 条件:
* 列表 deadends 包含了一组死亡数字,一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会永久锁定，无法再次被旋转。
* 字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何都不能解锁，返回 -1
 */
package bfs

import (
	"github.com/lucky51/pkg/queue"
)

// openLock 单向扩散的方式
func openLock(deadends [][]rune, target string) int {
	deadendslocksDict := make(map[string]byte, 0)
	for _, d := range deadends {
		deadendslocksDict[string(d)] = 1
	}
	visited := make(map[string]bool)

	step := 0
	q := queue.NewQueue[[]rune]()
	start := []rune{'0', '0', '0', '0'}
	visited["0000"] = true
	q.Offer(&start)
	for !q.IsEmpty() {
		len := int(q.Size())
		for i := 0; i < len; i++ {
			cur := q.Poll()
			if _, ok := deadendslocksDict[string(*cur)]; ok {
				continue
			}
			if string(*cur) == target {
				return step
			}
			// 选取将当前元素的所有相邻节点
			for j := 0; j < 4; j++ {
				// up
				up := scrollUp(*cur, j)
				// 访问过的不处理
				if _, ok := visited[string(up)]; !ok {
					q.Offer(&up)
					visited[string(up)] = true
				}

				//down
				down := scrollDown(*cur, j)
				// 访问过的不处理
				if _, ok := visited[string(down)]; !ok {
					q.Offer(&down)
					visited[string(down)] = true
				}

			}
		}
		step++
	}
	// 无解
	return -1
}

func scrollUp(s []rune, index int) []rune {
	cp := make([]rune, 4)
	copy(cp, s)
	if cp[index] == '9' {
		cp[index] = '0'
	} else {
		// c, _ := strconv.Atoi(string(cp[index]))
		// cstr := []rune(fmt.Sprintf("%d", c+1))
		// cp[index] = cstr[0]
		cp[index] = (cp[index]+1)%48 + 48

	}
	return cp
}

func scrollDown(s []rune, index int) []rune {
	cp := make([]rune, 4)
	copy(cp, s)
	if cp[index] == '0' {
		cp[index] = '9'
	} else {
		cp[index] = (cp[index]-1)%48 + 48

	}
	return cp
}

func contains(s [][]rune, target []rune) bool {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == string(target) {
			return true
		}
	}
	return false
}

// openLockDouble 双向扩散的方式，优化算法
func openLockDouble(deadends map[string]struct{}, target string) int {
	q1, q2 := make([][]rune, 0), make([][]rune, 0)
	q1 = append(q1, []rune{'0', '0', '0', '0'})
	q2 = append(q2, []rune(target))
	visited := make(map[string]struct{})
	visited["0000"] = struct{}{}
	step := 0
	for len(q1) > 0 && len(q2) > 0 {
		temp := make([][]rune, 0)
		//扩散q1
		for i := 0; i < len(q1); i++ {
			if _, ok := deadends[string(q1[i])]; ok {
				continue
			}
			// q2中含有q1的元素，则代表完成
			if contains(q2, q1[i]) {
				return step
			}
			visited[string(q1[i])] = struct{}{}
			for j := 0; j < 4; j++ {
				down := scrollDown(q1[i], j)
				if _, ok := visited[string(down)]; !ok {
					temp = append(temp, down)
				}

				up := scrollUp(q1[i], j)
				if _, ok := visited[string(up)]; !ok {
					temp = append(temp, up)
				}
			}
		}
		q1, q2 = q2, temp
		step++
	}
	return -1
}
