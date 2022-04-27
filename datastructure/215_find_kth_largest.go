package datastructure

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	kqsort(nums)
	fmt.Println(nums)
	kIndex := 0
	for kIndex = len(nums) - 1; kIndex >= 0 && k > 0; kIndex-- {
		fmt.Println(k, kIndex)
		k--
	}
	fmt.Println("ssdssd:", kIndex, nums[kIndex])
	return nums[kIndex]
}

func kqsort(nums []int) {
	if len(nums) < 2 {
		return
	}
	pivot := nums[0]
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	kqsort(nums[:left])
	kqsort(nums[left+1:])
}
