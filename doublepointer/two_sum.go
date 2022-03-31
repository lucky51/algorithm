package doublepointer

// twoSum 保证存在两个数字的和等于target，返回这两个书的索引
func twoSum(nums []int, target int) []int {
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
