package stackq

import (
	"container/list"
	"fmt"
	"github.com/lucky51/pkg/stack"
)

func calculate(s string) int {
	l := list.New()
	for i := range s {
		l.PushBack(s[i])
	}
	return calculateHelper(l)
}

func isDigital(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
func calculateHelper(s *list.List) int {
	var sign byte = '+'
	num := 0
	stk := stack.NewStack[int]()
	temp := s.Front()
loop:
	for temp != nil {
		curr := temp.Value.(uint8)
		s.Remove(temp)
		condition := isDigital(curr)
		if condition {
			num = num*10 + int(curr-'0')
		}
		if curr == '(' {
			num = calculateHelper(s)
		}
		if (curr != ' ' && !condition) || s.Len() == 0 {
			switch sign {
			case '+':
				stk.Push(num)
			case '-':
				stk.Push(-num)
			case '*':
				prev, err := stk.Pop()
				if err == nil {
					num = prev.Data * num
					stk.Push(num)
				}
			case '/':
				prev, err := stk.Pop()
				if err == nil {
					num = prev.Data / num
					stk.Push(num)
				}
			}
			sign = curr
			num = 0
		}
		if curr == ')' {
			break loop
		}
		temp = s.Front()
	}
	for !stk.IsEmpty() {
		curr, _ := stk.Pop()
		num += curr.Data
	}
	return num
}

func simpleCalculate(s string) int {
	ops := []int{1}
	sign := 1
	nums := 0
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && s[i] != ' ' && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			nums += num * sign
		}
	}
	return nums
}

func calculate1(s string) int {
	stack := make([]int, 0)
	n := len(s)
	sign := '+'
	nums := 0
	for i := 0; i < n; i++ {
		//digital
		if s[i] >= '0' && s[i] <= '9' {
			nums += nums*10 + int(s[i]-'0')
		}
		fmt.Println("char:", string(s[i]))
		if ((s[i] < '0' || s[i] > '9') && s[i] != ' ') || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, nums)
			case '-':
				stack = append(stack, -nums)
			case '*':
				v := stack[len(stack)-1] * nums
				stack[len(stack)-1] = v
			case '/':
				v := stack[len(stack)-1] / nums
				fmt.Println("/ v:", v)
				stack[len(stack)-1] = v
			}
			sign = rune(s[i])
			fmt.Println("sign:", string(sign), sign == '/')
			nums = 0
		}
	}
	nums = 0
	for j := 0; j < len(stack); j++ {
		nums += stack[j]
	}
	return nums
}
