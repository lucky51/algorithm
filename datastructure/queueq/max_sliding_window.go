package queueq

import (
	"github.com/lucky51/pkg/queue"
)

var window = queue.NewQueue[int]()

func qpush(n int) {
	for !window.IsEmpty() && window.PeekLast() < n {
		window.PollLast()
	}
	window.Offer(n)
}

// maxSlidingWindow Leetcode 239
// 给你输入一个数组 nums 和一个正整数 k,有一个大小为 k 的窗口在 nums 上从左到右滑动，请你输出每次窗口中k个元素的最大值
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if i < k-1 {
			qpush(nums[i])
		} else {
			qpush(nums[i])
			res = append(res, window.Peek())
			if nums[i-k+1] == window.Peek() {
				window.Poll()
			}

		}
	}
	return res
}

var window1 []int = make([]int, 0)

func qpush1(n int) {
	for len(window1) > 0 && window1[len(window1)-1] < n {
		window1 = window1[:len(window1)-1]
	}
	window1 = append(window1, n)
}

// 这个方法在leecode上没通过，执行的结果不一样
func maxSlidingWindowLeetcode(nums []int, k int) []int {
	res := make([]int, 0)
	n := len(nums)

	for i := 0; i < n; i++ {
		if i < k-1 {
			qpush1(nums[i])
		} else {
			qpush1(nums[i])
			res = append(res, window1[0])
			if nums[i-k+1] == window1[0] {
				window1 = window1[1:]
			}
		}
	}
	return res
}
