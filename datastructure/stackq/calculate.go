package stackq

import (
	"fmt"

	"github.com/lucky51/pkg/stack"
)

func calculate(s string) int {
	var sign byte = '+'
	num := 0
	stk := stack.NewStack[int]()

	for i := 0; i < len(s); i++ {
		if isDigital(s[i]) {
			num = num*10 + int(s[i]-'0')
			fmt.Println("num:", num)
		} else {
			fmt.Println(string(s[i]))
			if s[i] == ' ' {
				continue
			}
			switch sign {
			case '+':
				fmt.Println("push:", num)
				stk.Push(num)
			case '-':
				stk.Push(-num)
			case '*':
				prev, err := stk.Pop()
				if err != nil {
					num = prev.Data * num
					stk.Push(num)
				}
			case '/':
				prev, err := stk.Pop()
				if err != nil {
					num = prev.Data / num
					stk.Push(num)
				}
			}
			sign = s[i]
			num = 0
		}
	}
	num = 0
	fmt.Println("size:", stk.Size())
	for !stk.IsEmpty() {
		curr, _ := stk.Pop()
		num += curr.Data
	}
	return num
}

func isDigital(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
