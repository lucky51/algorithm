package doublepointer

import (
	"sort"
)

// twoSum 保证存在两个数字的和等于target，返回这两个数字的索引
func twoSumByHash(nums []int, target int) []int {
	indexer := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		indexer[nums[i]] = i
	}
	for j := 0; j < len(nums); j++ {
		other := target - nums[j]
		if otherIndex, ok := indexer[other]; ok && otherIndex != j {
			return []int{j, otherIndex}
		}
	}
	return []int{-1, -1}
}

// twoSumByPointer 采用双指针的方式进行查找，要先保证有序
func towSumTarget(nums []int, target int) [][]int {
	var res [][]int = make([][]int, 0)
	sort.Ints(nums)
	lo, hi := 0, len(nums)-1

	for lo < hi {
		//fmt.Println("lo,hi", lo, hi)
		lt, rh := nums[lo], nums[hi]
		sum := lt + rh
		if sum < target {
			for lo < hi && lt == nums[lo] {
				lo++
			}
		} else if sum > target {
			for lo < hi && rh == nums[hi] {
				hi--
			}
		} else {
			res = append(res, []int{lo, hi})
			for lo < hi && lt == nums[lo] {
				lo++
			}
			for lo < hi && rh == nums[hi] {
				hi--
			}
		}
	}
	return res
}

// nSumTarget n个数求得指定target的值
func nSumTarget(nums []int, n, start, target int) [][]int {
	sort.Ints(nums)
	var nSumTargetRes = make([][]int, 0)
	sz := len(nums)
	if n < 2 || sz < n {
		return nSumTargetRes
	}
	if n == 2 {
		lo, hi := start, len(nums)-1

		for lo < hi {
			left, right := nums[lo], nums[hi]
			sum := left + right
			if sum < target {
				for lo < hi && left == nums[lo] {
					lo++
				}
			} else if sum > target {
				for lo < hi && right == nums[hi] {
					hi--
				}
			} else {
				nSumTargetRes = append(nSumTargetRes, []int{lo, hi})
				for lo < hi && left == nums[lo] {
					lo++
				}
				for lo < hi && right == nums[hi] {
					hi--
				}
			}

		}
		//towSumResult := towSumTarget(nums, target)
		//nSumTargetRes = append(nSumTargetRes, towSumResult...)
	} else {

		for i := start; i < sz; i++ {
			tempSlice := nSumTarget(nums, n-1, i+1, target-nums[i])
			for j := 0; j < len(tempSlice); j++ {
				//for tj := range tempSlice[j] {
				//	tempSlice[j][tj] = nums[tempSlice[j][tj]]
				//}
				tempSlice[j] = append(tempSlice[j], i)
			}
			nSumTargetRes = append(nSumTargetRes, tempSlice...)
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}

	}
	return nSumTargetRes
}
