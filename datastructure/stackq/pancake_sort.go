package stackq

var pancakeSortRes = make([]int, 0)

// 969. leetcode 煎饼排序 这是到懵逼题，leetcode执行和预期输出不一样，但是可以提交通过。
// 假设盘子上有n块面积大小不一的烧饼，你如何用一把锅铲进行若干次翻转，让这些烧饼的大小有序
// 输入:[3,2,4,1]
// 输出:[4,2,4,3]
// 解释:
// 我们执行4次煎饼翻转，k值分别为 4,2,4,3
// 初始状态 A =[3,2,4,1]
// 第一次翻转后 (k=4) : A=[1,4,2,3]
// 第二次翻转后 (k=2) : A=[4,1,2,3]
// 第三次翻转后 (k=4) : A=[3,2,1,4]
// 第四次翻转后 (k=3) : A=[1,2,3,4] ,此时已完成排序
func pancakeSort(nums []int) []int {
	sort1(nums, len(nums))
	return pancakeSortRes
}

// leetcode 官网题解
func pancakeSort1(arr []int) (ans []int) {
	for n := len(arr); n > 1; n-- {
		index := 0
		for i, v := range arr[:n] {
			if v > arr[index] {
				index = i
			}
		}
		if index == n-1 {
			continue
		}
		for i, m := 0, index; i < (m+1)/2; i++ {
			arr[i], arr[m-i] = arr[m-i], arr[i]
		}
		for i := 0; i < n/2; i++ {
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
		ans = append(ans, index+1, n)
	}
	return
}

func sort1(nums []int, n int) {
	if n == 1 {
		return
	}
	var maxVal, maxIndex = -1, 0
	for i := 0; i < n; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			maxIndex = i

		}
	}

	reverse(nums, 0, maxIndex)
	pancakeSortRes = append(pancakeSortRes, maxIndex+1)
	reverse(nums, 0, n-1)
	pancakeSortRes = append(pancakeSortRes, n)

	sort1(nums, n-1)
}
func reverse(nums []int, start, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
