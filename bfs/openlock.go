package bfs

import (
	"github.com/lucky51/pkg/queue"
)

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
		// c, _ := strconv.Atoi(string(cp[index]))
		// cstr := []rune(fmt.Sprintf("%d", c-1))
		// cp[index] = cstr[0]
		cp[index] = (cp[index]-1)%48 + 48

	}
	return cp
}
