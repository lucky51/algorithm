package stackq

import "github.com/lucky51/pkg/stack"

// nextGreaterElement 单调栈解题模板
// 输入一个数组,返回一个等长的数组,对应索引存储着下一个更大元素
// 如果没有更大的元素返回 -1
// 例如输入 nums=[2,1,2,4,3] ,返回 [4,3,4,-1,-1]
func nextGreaterElement(nums []int) []int {
	res := make([]int, len(nums))
	s := stack.NewStack[int]()
	for i := len(nums) - 1; i >= 0; i-- {
		for !s.IsEmpty() {
			top, _ := s.Peek()
			if *top.Data <= nums[i] {
				s.Pop()
			} else {
				break
			}
		}
		if s.IsEmpty() {
			res[i] = -1
		} else {
			top, _ := s.Peek()
			res[i] = *top.Data
		}
		s.Push(&nums[i])
	}
	return res
}

// dailyTemperatures 计算每一天至少要等多少天才能等到一个更暖和的天气
// 这本质还是一个单调栈的问题，但是数组中存放的是索引
func dailyTemperatures(nums []int) []int {
	res := make([]int, len(nums))
	s := stack.NewStack[int]()

	for i := len(nums) - 1; i >= 0; i-- {
		for !s.IsEmpty() {
			top, _ := s.Peek()
			if *top.Data <= nums[i] {
				s.Pop()
			} else {
				break
			}
		}
		if s.IsEmpty() {
			res[i] = 0
		} else {
			top, _ := s.Peek()
			res[i] = *top.Data - i
		}
		s.Push(&i)
	}
	return res
}

// nextGreaterElementCycle 计算环形数组的单调栈问题
// 方法1 可以构造一个2倍的数组长度，然后套用上边的方法
// 方法2 可以用 %来代替索引，防止索引越界
func nextGreaterElementCycle(nums []int) []int {
	s := stack.NewStack[int]()
	n := len(nums)
	res := make([]int, n)
	for i := 2 * (n - 1); i >= 0; i-- {
		for s.IsEmpty() {
			top, _ := s.Peek()
			if *top.Data <= nums[i%n] {
				s.Pop()
			} else {
				break
			}
		}
		if s.IsEmpty() {
			res[i%n] = -1
		} else {
			top, _ := s.Peek()
			res[i%n] = *top.Data
		}
		s.Push(&nums[i%n])
	}
	return res
}
